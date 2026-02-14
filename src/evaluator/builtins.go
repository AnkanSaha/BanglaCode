package evaluator

import (
	"BanglaCode/src/object"
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"math"
	"math/rand"
	"net/http"
	"os"
	"sort"
	"strings"
	"time"
)

// EvalFunc is a function pointer for evaluating AST nodes (set by evaluator.go to avoid circular dependency)
var EvalFunc func(handler *object.Function, args []object.Object) object.Object

var builtins = map[string]*object.Builtin{
	// Output - dekho (দেখো - see/show)
	"dekho": {
		Fn: func(args ...object.Object) object.Object {
			for i, arg := range args {
				if i > 0 {
					fmt.Print(" ")
				}
				fmt.Print(arg.Inspect())
			}
			fmt.Println()
			return object.NULL
		},
	},
	// Length - dorghyo (দৈর্ঘ্য - length)
	"dorghyo": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}

			switch arg := args[0].(type) {
			case *object.String:
				return &object.Number{Value: float64(len(arg.Value))}
			case *object.Array:
				return &object.Number{Value: float64(len(arg.Elements))}
			default:
				return newError("argument to `dorghyo` not supported, got %s", args[0].Type())
			}
		},
	},
	// Push - dhokao (ঢোকাও - insert)
	"dhokao": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				return newError("wrong number of arguments. got=%d, want=2", len(args))
			}
			if args[0].Type() != object.ARRAY_OBJ {
				return newError("argument to `dhokao` must be ARRAY, got %s", args[0].Type())
			}

			arr := args[0].(*object.Array)
			arr.Elements = append(arr.Elements, args[1])
			return arr
		},
	},
	// Pop - beroKoro (বের করো - take out)
	"berKoro": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}
			if args[0].Type() != object.ARRAY_OBJ {
				return newError("argument to `berKoro` must be ARRAY, got %s", args[0].Type())
			}

			arr := args[0].(*object.Array)
			length := len(arr.Elements)
			if length > 0 {
				lastElement := arr.Elements[length-1]
				arr.Elements = arr.Elements[:length-1]
				return lastElement
			}
			return object.NULL
		},
	},
	// Keys - chabi (চাবি - keys)
	"chabi": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}
			if args[0].Type() != object.MAP_OBJ {
				return newError("argument to `chabi` must be MAP, got %s", args[0].Type())
			}

			mapObj := args[0].(*object.Map)
			keys := make([]object.Object, 0, len(mapObj.Pairs))
			for key := range mapObj.Pairs {
				keys = append(keys, &object.String{Value: key})
			}
			return &object.Array{Elements: keys}
		},
	},
	// Type - dhoron (ধরন - type)
	"dhoron": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}
			return &object.String{Value: string(args[0].Type())}
		},
	},
	// To string - lipi (লিপি - text/script)
	"lipi": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}
			return &object.String{Value: args[0].Inspect()}
		},
	},
	// To number - sonkha (সংখ্যা - number)
	"sonkha": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}

			switch arg := args[0].(type) {
			case *object.Number:
				return arg
			case *object.String:
				var num float64
				_, err := fmt.Sscanf(arg.Value, "%f", &num)
				if err != nil {
					return newError("cannot convert string to number: %s", arg.Value)
				}
				return &object.Number{Value: num}
			case *object.Boolean:
				if arg.Value {
					return &object.Number{Value: 1}
				}
				return &object.Number{Value: 0}
			default:
				return newError("cannot convert %s to number", arg.Type())
			}
		},
	},
	// Math functions
	// Square root - borgomul (বর্গমূল)
	"borgomul": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}
			if args[0].Type() != object.NUMBER_OBJ {
				return newError("argument to `borgomul` must be NUMBER, got %s", args[0].Type())
			}
			num := args[0].(*object.Number).Value
			return &object.Number{Value: math.Sqrt(num)}
		},
	},
	// Power - ghat (ঘাত)
	"ghat": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				return newError("wrong number of arguments. got=%d, want=2", len(args))
			}
			if args[0].Type() != object.NUMBER_OBJ || args[1].Type() != object.NUMBER_OBJ {
				return newError("arguments to `ghat` must be NUMBERs")
			}
			base := args[0].(*object.Number).Value
			exp := args[1].(*object.Number).Value
			return &object.Number{Value: math.Pow(base, exp)}
		},
	},
	// Floor - niche (নিচে - down)
	"niche": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}
			if args[0].Type() != object.NUMBER_OBJ {
				return newError("argument to `niche` must be NUMBER, got %s", args[0].Type())
			}
			num := args[0].(*object.Number).Value
			return &object.Number{Value: math.Floor(num)}
		},
	},
	// Ceil - upore (উপরে - up)
	"upore": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}
			if args[0].Type() != object.NUMBER_OBJ {
				return newError("argument to `upore` must be NUMBER, got %s", args[0].Type())
			}
			num := args[0].(*object.Number).Value
			return &object.Number{Value: math.Ceil(num)}
		},
	},
	// Round - kache (কাছে - near)
	"kache": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}
			if args[0].Type() != object.NUMBER_OBJ {
				return newError("argument to `kache` must be NUMBER, got %s", args[0].Type())
			}
			num := args[0].(*object.Number).Value
			return &object.Number{Value: math.Round(num)}
		},
	},
	// Absolute - niratek (নিরপেক্ষ - absolute)
	"niratek": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}
			if args[0].Type() != object.NUMBER_OBJ {
				return newError("argument to `niratek` must be NUMBER, got %s", args[0].Type())
			}
			num := args[0].(*object.Number).Value
			return &object.Number{Value: math.Abs(num)}
		},
	},
	// Min - sobcheye_choto (সবচেয়ে ছোট - smallest)
	"choto": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) < 2 {
				return newError("wrong number of arguments. got=%d, want at least 2", len(args))
			}
			minVal := math.Inf(1)
			for _, arg := range args {
				if arg.Type() != object.NUMBER_OBJ {
					return newError("all arguments to `choto` must be NUMBERs")
				}
				val := arg.(*object.Number).Value
				if val < minVal {
					minVal = val
				}
			}
			return &object.Number{Value: minVal}
		},
	},
	// Max - sobcheye_boro (সবচেয়ে বড় - biggest)
	"boro": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) < 2 {
				return newError("wrong number of arguments. got=%d, want at least 2", len(args))
			}
			maxVal := math.Inf(-1)
			for _, arg := range args {
				if arg.Type() != object.NUMBER_OBJ {
					return newError("all arguments to `boro` must be NUMBERs")
				}
				val := arg.(*object.Number).Value
				if val > maxVal {
					maxVal = val
				}
			}
			return &object.Number{Value: maxVal}
		},
	},
	// String functions
	// Upper -boro_hater (বড় হাতের - uppercase)
	"boroHater": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}
			if args[0].Type() != object.STRING_OBJ {
				return newError("argument to `boroHater` must be STRING, got %s", args[0].Type())
			}
			str := args[0].(*object.String).Value
			return &object.String{Value: strings.ToUpper(str)}
		},
	},
	// Lower - choto_hater (ছোট হাতের - lowercase)
	"chotoHater": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}
			if args[0].Type() != object.STRING_OBJ {
				return newError("argument to `chotoHater` must be STRING, got %s", args[0].Type())
			}
			str := args[0].(*object.String).Value
			return &object.String{Value: strings.ToLower(str)}
		},
	},
	// Split - bhag (ভাগ - divide)
	"bhag": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				return newError("wrong number of arguments. got=%d, want=2", len(args))
			}
			if args[0].Type() != object.STRING_OBJ || args[1].Type() != object.STRING_OBJ {
				return newError("arguments to `bhag` must be STRINGs")
			}
			str := args[0].(*object.String).Value
			sep := args[1].(*object.String).Value
			parts := strings.Split(str, sep)
			elements := make([]object.Object, len(parts))
			for i, p := range parts {
				elements[i] = &object.String{Value: p}
			}
			return &object.Array{Elements: elements}
		},
	},
	// Join - joro (জোড়ো - join)
	"joro": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				return newError("wrong number of arguments. got=%d, want=2", len(args))
			}
			if args[0].Type() != object.ARRAY_OBJ {
				return newError("first argument to `joro` must be ARRAY, got %s", args[0].Type())
			}
			if args[1].Type() != object.STRING_OBJ {
				return newError("second argument to `joro` must be STRING, got %s", args[1].Type())
			}
			arr := args[0].(*object.Array)
			sep := args[1].(*object.String).Value
			parts := make([]string, len(arr.Elements))
			for i, el := range arr.Elements {
				parts[i] = el.Inspect()
			}
			return &object.String{Value: strings.Join(parts, sep)}
		},
	},
	// Trim - chhanto (ছাঁটো - trim)
	"chhanto": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}
			if args[0].Type() != object.STRING_OBJ {
				return newError("argument to `chhanto` must be STRING, got %s", args[0].Type())
			}
			str := args[0].(*object.String).Value
			return &object.String{Value: strings.TrimSpace(str)}
		},
	},
	// Index of - khojo (খোঁজো - search)
	"khojo": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				return newError("wrong number of arguments. got=%d, want=2", len(args))
			}
			if args[0].Type() != object.STRING_OBJ || args[1].Type() != object.STRING_OBJ {
				return newError("arguments to `khojo` must be STRINGs")
			}
			str := args[0].(*object.String).Value
			substr := args[1].(*object.String).Value
			return &object.Number{Value: float64(strings.Index(str, substr))}
		},
	},
	// Substring - angsho (অংশ - portion)
	"angsho": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) < 2 || len(args) > 3 {
				return newError("wrong number of arguments. got=%d, want=2 or 3", len(args))
			}
			if args[0].Type() != object.STRING_OBJ {
				return newError("first argument to `angsho` must be STRING, got %s", args[0].Type())
			}
			if args[1].Type() != object.NUMBER_OBJ {
				return newError("second argument to `angsho` must be NUMBER, got %s", args[1].Type())
			}
			str := args[0].(*object.String).Value
			start := int(args[1].(*object.Number).Value)
			end := len(str)
			if len(args) == 3 {
				if args[2].Type() != object.NUMBER_OBJ {
					return newError("third argument to `angsho` must be NUMBER, got %s", args[2].Type())
				}
				end = int(args[2].(*object.Number).Value)
			}
			if start < 0 {
				start = 0
			}
			if end > len(str) {
				end = len(str)
			}
			if start > end {
				return &object.String{Value: ""}
			}
			return &object.String{Value: str[start:end]}
		},
	},
	// Replace - bodlo (বদলো - change)
	"bodlo": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 3 {
				return newError("wrong number of arguments. got=%d, want=3", len(args))
			}
			if args[0].Type() != object.STRING_OBJ || args[1].Type() != object.STRING_OBJ || args[2].Type() != object.STRING_OBJ {
				return newError("all arguments to `bodlo` must be STRINGs")
			}
			str := args[0].(*object.String).Value
			old := args[1].(*object.String).Value
			new := args[2].(*object.String).Value
			return &object.String{Value: strings.ReplaceAll(str, old, new)}
		},
	},
	// Array functions
	// Slice - kato (কাটো - cut)
	"kato": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) < 2 || len(args) > 3 {
				return newError("wrong number of arguments. got=%d, want=2 or 3", len(args))
			}
			if args[0].Type() != object.ARRAY_OBJ {
				return newError("first argument to `kato` must be ARRAY, got %s", args[0].Type())
			}
			if args[1].Type() != object.NUMBER_OBJ {
				return newError("second argument to `kato` must be NUMBER, got %s", args[1].Type())
			}
			arr := args[0].(*object.Array)
			start := int(args[1].(*object.Number).Value)
			end := len(arr.Elements)
			if len(args) == 3 {
				if args[2].Type() != object.NUMBER_OBJ {
					return newError("third argument to `kato` must be NUMBER, got %s", args[2].Type())
				}
				end = int(args[2].(*object.Number).Value)
			}
			if start < 0 {
				start = 0
			}
			if end > len(arr.Elements) {
				end = len(arr.Elements)
			}
			if start > end {
				return &object.Array{Elements: []object.Object{}}
			}
			newElements := make([]object.Object, end-start)
			copy(newElements, arr.Elements[start:end])
			return &object.Array{Elements: newElements}
		},
	},
	// Reverse - ulto (উল্টো - reverse)
	"ulto": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}
			if args[0].Type() != object.ARRAY_OBJ {
				return newError("argument to `ulto` must be ARRAY, got %s", args[0].Type())
			}
			arr := args[0].(*object.Array)
			length := len(arr.Elements)
			newElements := make([]object.Object, length)
			for i := 0; i < length; i++ {
				newElements[i] = arr.Elements[length-1-i]
			}
			return &object.Array{Elements: newElements}
		},
	},
	// Includes - ache (আছে - exists)
	"ache": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				return newError("wrong number of arguments. got=%d, want=2", len(args))
			}
			if args[0].Type() != object.ARRAY_OBJ {
				return newError("first argument to `ache` must be ARRAY, got %s", args[0].Type())
			}
			arr := args[0].(*object.Array)
			target := args[1]
			for _, el := range arr.Elements {
				if objectsEqual(el, target) {
					return object.TRUE
				}
			}
			return object.FALSE
		},
	},
	// Sort - saja (সাজা - arrange)
	"saja": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}
			if args[0].Type() != object.ARRAY_OBJ {
				return newError("argument to `saja` must be ARRAY, got %s", args[0].Type())
			}
			arr := args[0].(*object.Array)
			newElements := make([]object.Object, len(arr.Elements))
			copy(newElements, arr.Elements)
			sort.Slice(newElements, func(i, j int) bool {
				if newElements[i].Type() == object.NUMBER_OBJ && newElements[j].Type() == object.NUMBER_OBJ {
					return newElements[i].(*object.Number).Value < newElements[j].(*object.Number).Value
				}
				return newElements[i].Inspect() < newElements[j].Inspect()
			})
			return &object.Array{Elements: newElements}
		},
	},
	// Utility functions
	// Time - somoy (সময় - time)
	"somoy": {
		Fn: func(args ...object.Object) object.Object {
			return &object.Number{Value: float64(time.Now().UnixMilli())}
		},
	},
	// Random - lotto (লটো - lottery/random)
	"lotto": {
		Fn: func(args ...object.Object) object.Object {
			return &object.Number{Value: rand.Float64()}
		},
	},
	// Input - nao (নাও - take)
	"nao": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) > 0 {
				fmt.Print(args[0].Inspect())
			}
			reader := bufio.NewReader(os.Stdin)
			text, err := reader.ReadString('\n')
			if err != nil {
				if err == io.EOF {
					return &object.String{Value: ""}
				}
				return newError("error reading input: %s", err.Error())
			}
			return &object.String{Value: strings.TrimSpace(text)}
		},
	},
	// Read file - ppioro (পড়ো - read)
	"poro": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}
			if args[0].Type() != object.STRING_OBJ {
				return newError("argument to `poro` must be STRING, got %s", args[0].Type())
			}
			path := args[0].(*object.String).Value
			content, err := os.ReadFile(path)
			if err != nil {
				return newError("error reading file: %s", err.Error())
			}
			return &object.String{Value: string(content)}
		},
	},
	// Write file - lekho (লেখো - write)
	"lekho": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				return newError("wrong number of arguments. got=%d, want=2", len(args))
			}
			if args[0].Type() != object.STRING_OBJ {
				return newError("first argument to `lekho` must be STRING, got %s", args[0].Type())
			}
			path := args[0].(*object.String).Value
			content := args[1].Inspect()
			err := os.WriteFile(path, []byte(content), 0644)
			if err != nil {
				return newError("error writing file: %s", err.Error())
			}
			return object.TRUE
		},
	},
	// Exit - bondho (বন্ধ - stop/close)
	"bondho": {
		Fn: func(args ...object.Object) object.Object {
			code := 0
			if len(args) > 0 && args[0].Type() == object.NUMBER_OBJ {
				code = int(args[0].(*object.Number).Value)
			}
			os.Exit(code)
			return object.NULL
		},
	},
	// Sleep - ghum (ঘুম - sleep)
	"ghum": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}
			if args[0].Type() != object.NUMBER_OBJ {
				return newError("argument to `ghum` must be NUMBER, got %s", args[0].Type())
			}
			ms := int64(args[0].(*object.Number).Value)
			time.Sleep(time.Duration(ms) * time.Millisecond)
			return object.NULL
		},
	},
	// HTTP Server - server_chalu (সার্ভার চালু - start server)
	"server_chalu": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				return newError("wrong number of arguments. got=%d, want=2", len(args))
			}
			if args[0].Type() != object.NUMBER_OBJ {
				return newError("first argument to `server_chalu` must be NUMBER (port), got %s", args[0].Type())
			}
			if args[1].Type() != object.FUNCTION_OBJ {
				return newError("second argument to `server_chalu` must be FUNCTION (handler), got %s", args[1].Type())
			}

			port := int(args[0].(*object.Number).Value)
			handler := args[1].(*object.Function)

			http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
				reqMap := &object.Map{Pairs: make(map[string]object.Object)}
				reqMap.Pairs["method"] = &object.String{Value: r.Method}
				reqMap.Pairs["path"] = &object.String{Value: r.URL.Path}
				reqMap.Pairs["query"] = &object.String{Value: r.URL.RawQuery}

				headersMap := &object.Map{Pairs: make(map[string]object.Object)}
				for k, v := range r.Header {
					if len(v) > 0 {
						headersMap.Pairs[k] = &object.String{Value: v[0]}
					}
				}
				reqMap.Pairs["headers"] = headersMap

				body, _ := io.ReadAll(r.Body)
				reqMap.Pairs["body"] = &object.String{Value: string(body)}

				resMap := &object.Map{Pairs: make(map[string]object.Object)}
				resMap.Pairs["status"] = &object.Number{Value: 200}
				resMap.Pairs["body"] = &object.String{Value: ""}
				resMap.Pairs["headers"] = &object.Map{Pairs: make(map[string]object.Object)}

				var result object.Object
				if EvalFunc != nil {
					result = EvalFunc(handler, []object.Object{reqMap, resMap})
				}

				if statusObj, ok := resMap.Pairs["status"]; ok {
					if status, ok := statusObj.(*object.Number); ok {
						w.WriteHeader(int(status.Value))
					}
				}

				if headersObj, ok := resMap.Pairs["headers"]; ok {
					if headers, ok := headersObj.(*object.Map); ok {
						for k, v := range headers.Pairs {
							w.Header().Set(k, v.Inspect())
						}
					}
				}

				if bodyObj, ok := resMap.Pairs["body"]; ok {
					fmt.Fprint(w, bodyObj.Inspect())
				} else if result != nil && result != object.NULL {
					fmt.Fprint(w, result.Inspect())
				}
			})

			fmt.Printf("🚀 Server cholche http://localhost:%d e\n", port)
			err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
			if err != nil {
				return newError("server error: %s", err.Error())
			}
			return object.NULL
		},
	},
	// HTTP GET - anun (আনুন - fetch/bring)
	"anun": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}
			if args[0].Type() != object.STRING_OBJ {
				return newError("argument to `anun` must be STRING, got %s", args[0].Type())
			}
			url := args[0].(*object.String).Value

			resp, err := http.Get(url)
			if err != nil {
				return newError("HTTP error: %s", err.Error())
			}
			defer resp.Body.Close()

			body, err := io.ReadAll(resp.Body)
			if err != nil {
				return newError("error reading response: %s", err.Error())
			}

			result := &object.Map{Pairs: make(map[string]object.Object)}
			result.Pairs["status"] = &object.Number{Value: float64(resp.StatusCode)}
			result.Pairs["body"] = &object.String{Value: string(body)}

			return result
		},
	},
	// JSON Parse - json_poro (JSON পড়ো - read JSON)
	"json_poro": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}
			if args[0].Type() != object.STRING_OBJ {
				return newError("argument to `json_poro` must be STRING, got %s", args[0].Type())
			}
			jsonStr := args[0].(*object.String).Value
			return parseJSON(jsonStr)
		},
	},
	// JSON Stringify - json_banao (JSON বানাও - make JSON)
	"json_banao": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}
			return &object.String{Value: stringifyJSON(args[0])}
		},
	},
	// Simple HTTP response helper - uttor (উত্তর - reply/response)
	"uttor": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) < 2 || len(args) > 4 {
				return newError("wrong number of arguments. got=%d, want=2-4 (res, body, [status], [contentType])", len(args))
			}
			if args[0].Type() != object.MAP_OBJ {
				return newError("first argument to `uttor` must be response MAP, got %s", args[0].Type())
			}
			resMap := args[0].(*object.Map)

			// Set body
			resMap.Pairs["body"] = args[1]

			// Set status (optional, default 200)
			if len(args) >= 3 {
				if args[2].Type() != object.NUMBER_OBJ {
					return newError("third argument to `uttor` must be NUMBER (status), got %s", args[2].Type())
				}
				resMap.Pairs["status"] = args[2]
			}

			// Set content-type (optional)
			if len(args) >= 4 {
				if args[3].Type() != object.STRING_OBJ {
					return newError("fourth argument to `uttor` must be STRING (contentType), got %s", args[3].Type())
				}
				if headersObj, ok := resMap.Pairs["headers"]; ok {
					if headers, ok := headersObj.(*object.Map); ok {
						headers.Pairs["Content-Type"] = args[3]
					}
				}
			}

			return resMap
		},
	},
	// JSON response helper - json_uttor (JSON উত্তর - JSON reply)
	"json_uttor": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) < 2 || len(args) > 3 {
				return newError("wrong number of arguments. got=%d, want=2-3 (res, data, [status])", len(args))
			}
			if args[0].Type() != object.MAP_OBJ {
				return newError("first argument to `json_uttor` must be response MAP, got %s", args[0].Type())
			}
			resMap := args[0].(*object.Map)

			// Convert data to JSON string
			jsonStr := stringifyJSON(args[1])
			resMap.Pairs["body"] = &object.String{Value: jsonStr}

			// Set status (optional, default 200)
			if len(args) >= 3 {
				if args[2].Type() != object.NUMBER_OBJ {
					return newError("third argument to `json_uttor` must be NUMBER (status), got %s", args[2].Type())
				}
				resMap.Pairs["status"] = args[2]
			}

			// Set content-type to JSON
			if headersObj, ok := resMap.Pairs["headers"]; ok {
				if headers, ok := headersObj.(*object.Map); ok {
					headers.Pairs["Content-Type"] = &object.String{Value: "application/json; charset=utf-8"}
				}
			}

			return resMap
		},
	},
	// Async built-in functions
	// ghumaao (ঘুমাও) - async sleep
	"ghumaao": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}
			if args[0].Type() != object.NUMBER_OBJ {
				return newError("argument to `ghumaao` must be NUMBER, got %s", args[0].Type())
			}

			ms := int64(args[0].(*object.Number).Value)
			promise := createPromise()

			go func() {
				time.Sleep(time.Duration(ms) * time.Millisecond)
				resolvePromise(promise, object.NULL)
			}()

			return promise
		},
	},
	// anun_async (আনুন_async) - async HTTP GET
	"anun_async": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}
			if args[0].Type() != object.STRING_OBJ {
				return newError("argument to `anun_async` must be STRING, got %s", args[0].Type())
			}

			url := args[0].(*object.String).Value
			promise := createPromise()

			go func() {
				resp, err := http.Get(url)
				if err != nil {
					rejectPromise(promise, newError("HTTP error: %s", err.Error()))
					return
				}
				defer resp.Body.Close()

				body, err := io.ReadAll(resp.Body)
				if err != nil {
					rejectPromise(promise, newError("error reading response: %s", err.Error()))
					return
				}

				result := &object.Map{Pairs: make(map[string]object.Object)}
				result.Pairs["status"] = &object.Number{Value: float64(resp.StatusCode)}
				result.Pairs["body"] = &object.String{Value: string(body)}

				resolvePromise(promise, result)
			}()

			return promise
		},
	},
	// poro_async (পড়ো_async) - async file read
	"poro_async": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}
			if args[0].Type() != object.STRING_OBJ {
				return newError("argument to `poro_async` must be STRING, got %s", args[0].Type())
			}

			path := args[0].(*object.String).Value
			promise := createPromise()

			go func() {
				content, err := os.ReadFile(path)
				if err != nil {
					rejectPromise(promise, newError("error reading file: %s", err.Error()))
					return
				}
				resolvePromise(promise, &object.String{Value: string(content)})
			}()

			return promise
		},
	},
	// lekho_async (লেখো_async) - async file write
	"lekho_async": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				return newError("wrong number of arguments. got=%d, want=2", len(args))
			}
			if args[0].Type() != object.STRING_OBJ {
				return newError("first argument to `lekho_async` must be STRING, got %s", args[0].Type())
			}

			path := args[0].(*object.String).Value
			content := args[1].Inspect()
			promise := createPromise()

			go func() {
				err := os.WriteFile(path, []byte(content), 0644)
				if err != nil {
					rejectPromise(promise, newError("error writing file: %s", err.Error()))
					return
				}
				resolvePromise(promise, object.TRUE)
			}()

			return promise
		},
	},
	// sob_proyash (সব_প্রয়াস) - Promise.all
	"sob_proyash": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}
			if args[0].Type() != object.ARRAY_OBJ {
				return newError("argument to `sob_proyash` must be ARRAY, got %s", args[0].Type())
			}

			arr := args[0].(*object.Array)
			promises := make([]*object.Promise, len(arr.Elements))

			// Validate all elements are promises
			for i, el := range arr.Elements {
				p, ok := el.(*object.Promise)
				if !ok {
					return newError("all elements must be promises, got %s at index %d", el.Type(), i)
				}
				promises[i] = p
			}

			resultPromise := createPromise()

			go func() {
				results := make([]object.Object, len(promises))
				for i, p := range promises {
					// Wait for each promise
					select {
					case res := <-p.ResultChan:
						results[i] = res
					case err := <-p.ErrorChan:
						rejectPromise(resultPromise, err)
						return
					}
				}
				resolvePromise(resultPromise, &object.Array{Elements: results})
			}()

			return resultPromise
		},
	},
}

// parseJSON converts a JSON string to BanglaCode objects
func parseJSON(jsonStr string) object.Object {
	var data interface{}
	if err := json.Unmarshal([]byte(jsonStr), &data); err != nil {
		return newError("JSON parse error: %s", err.Error())
	}
	return jsonToObject(data)
}

// jsonToObject recursively converts Go values to BanglaCode objects
func jsonToObject(data interface{}) object.Object {
	switch v := data.(type) {
	case nil:
		return object.NULL
	case bool:
		if v {
			return object.TRUE
		}
		return object.FALSE
	case float64:
		return &object.Number{Value: v}
	case string:
		return &object.String{Value: v}
	case []interface{}:
		elements := make([]object.Object, len(v))
		for i, item := range v {
			elements[i] = jsonToObject(item)
		}
		return &object.Array{Elements: elements}
	case map[string]interface{}:
		pairs := make(map[string]object.Object)
		for key, val := range v {
			pairs[key] = jsonToObject(val)
		}
		return &object.Map{Pairs: pairs}
	default:
		return newError("unsupported JSON type")
	}
}

// stringifyJSON converts a BanglaCode object to JSON string
func stringifyJSON(obj object.Object) string {
	data := objectToJSON(obj)
	bytes, err := json.Marshal(data)
	if err != nil {
		return "{}"
	}
	return string(bytes)
}

// objectToJSON recursively converts BanglaCode objects to Go values
func objectToJSON(obj object.Object) interface{} {
	switch v := obj.(type) {
	case *object.Null:
		return nil
	case *object.Boolean:
		return v.Value
	case *object.Number:
		return v.Value
	case *object.String:
		return v.Value
	case *object.Array:
		arr := make([]interface{}, len(v.Elements))
		for i, el := range v.Elements {
			arr[i] = objectToJSON(el)
		}
		return arr
	case *object.Map:
		m := make(map[string]interface{})
		for key, val := range v.Pairs {
			m[key] = objectToJSON(val)
		}
		return m
	default:
		return obj.Inspect()
	}
}
