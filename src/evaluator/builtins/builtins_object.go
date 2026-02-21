package builtins

import (
	"BanglaCode/src/object"
)

func init() {
	// Values - maan (মান - values)
	Builtins["maan"] = &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}
			if args[0].Type() != object.MAP_OBJ {
				return newError("argument to `maan` must be MAP, got %s", args[0].Type())
			}

			mapObj := args[0].(*object.Map)
			values := make([]object.Object, 0, len(mapObj.Pairs))

			for _, value := range mapObj.Pairs {
				values = append(values, value)
			}

			return &object.Array{Elements: values}
		},
	}

	// Entries - jora (জোড়া - pairs)
	Builtins["jora"] = &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}
			if args[0].Type() != object.MAP_OBJ {
				return newError("argument to `jora` must be MAP, got %s", args[0].Type())
			}

			mapObj := args[0].(*object.Map)
			entries := make([]object.Object, 0, len(mapObj.Pairs))

			for key, value := range mapObj.Pairs {
				entry := &object.Array{
					Elements: []object.Object{
						&object.String{Value: key},
						value,
					},
				}
				entries = append(entries, entry)
			}

			return &object.Array{Elements: entries}
		},
	}

	// Assign - mishra (মিশ্র - mix/merge)
	Builtins["mishra"] = &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) < 2 {
				return newError("wrong number of arguments. got=%d, want at least 2", len(args))
			}

			// First argument is target
			if args[0].Type() != object.MAP_OBJ {
				return newError("first argument to `mishra` must be MAP, got %s", args[0].Type())
			}

			target := args[0].(*object.Map)

			// Merge all source objects into target
			for i := 1; i < len(args); i++ {
				if args[i].Type() != object.MAP_OBJ {
					return newError("argument %d to `mishra` must be MAP, got %s", i+1, args[i].Type())
				}

				source := args[i].(*object.Map)
				for key, value := range source.Pairs {
					target.Pairs[key] = value
				}
			}

			return target
		},
	}
}
