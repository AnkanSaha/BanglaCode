package builtins

import (
	"BanglaCode/src/object"
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strings"
	"time"
)

func init() {
	// Output - dekho (দেখো - see/show)
	Builtins["dekho"] = &object.Builtin{
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
	}

	// Type - dhoron (ধরন - type)
	Builtins["dhoron"] = &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}
			return &object.String{Value: string(args[0].Type())}
		},
	}

	// To string - lipi (লিপি - text/script)
	Builtins["lipi"] = &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}
			return &object.String{Value: args[0].Inspect()}
		},
	}

	// To string alias - jongate (জনগতে - join/concatenate helper for numbers)
	// Alias for lipi - commonly used for number to string conversion in concatenation
	Builtins["jongate"] = &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}
			return &object.String{Value: args[0].Inspect()}
		},
	}

	// To number - sonkha (সংখ্যা - number)
	Builtins["sonkha"] = &object.Builtin{
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
	}

	// Time - somoy (সময় - time)
	Builtins["somoy"] = &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			return &object.Number{Value: float64(time.Now().UnixMilli())}
		},
	}

	// Random - lotto (লটো - lottery/random)
	Builtins["lotto"] = &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			return &object.Number{Value: rand.Float64()}
		},
	}

	// Input - nao (নাও - take)
	Builtins["nao"] = &object.Builtin{
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
	}

	// Exit - bondho (বন্ধ - stop/close)
	Builtins["bondho"] = &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			code := 0
			if len(args) > 0 && args[0].Type() == object.NUMBER_OBJ {
				code = int(args[0].(*object.Number).Value)
			}
			os.Exit(code)
			return object.NULL
		},
	}

	// Sleep - ghum (ঘুম - sleep)
	Builtins["ghum"] = &object.Builtin{
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
	}
}
