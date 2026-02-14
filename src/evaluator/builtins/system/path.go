package system

import (
	"BanglaCode/src/object"
	"path/filepath"
	"strings"
)

func init() {
	// ==================== Path Operations ====================

	// path_joro (পাথ জোড়ো) - Join path components
	registerBuiltin("path_joro", func(args ...object.Object) object.Object {
		if len(args) < 1 {
			return newError("path_joro requires at least 1 argument")
		}

		// Convert all arguments to strings
		parts := make([]string, len(args))
		for i, arg := range args {
			if arg.Type() != object.STRING_OBJ {
				return newError("all arguments must be STRING")
			}
			parts[i] = arg.(*object.String).Value
		}

		// Use filepath.Join for cross-platform path joining
		result := filepath.Join(parts...)
		return &object.String{Value: result}
	})

	// sompurno_path (সম্পূর্ণ পাথ) - Get absolute path
	registerBuiltin("sompurno_path", func(args ...object.Object) object.Object {
		if len(args) != 1 {
			return newError("sompurno_path requires 1 argument (path)")
		}
		if args[0].Type() != object.STRING_OBJ {
			return newError("path must be STRING, got %s", args[0].Type())
		}

		path := args[0].(*object.String).Value
		absPath, err := filepath.Abs(path)
		if err != nil {
			return newError("failed to get absolute path: %s", err.Error())
		}

		return &object.String{Value: absPath}
	})

	// path_naam (পাথ নাম) - Get base name (file name) from path
	registerBuiltin("path_naam", func(args ...object.Object) object.Object {
		if len(args) != 1 {
			return newError("path_naam requires 1 argument (path)")
		}
		if args[0].Type() != object.STRING_OBJ {
			return newError("path must be STRING, got %s", args[0].Type())
		}

		path := args[0].(*object.String).Value
		return &object.String{Value: filepath.Base(path)}
	})

	// directory_naam (ডিরেক্টরি নাম) - Get directory name from path
	registerBuiltin("directory_naam", func(args ...object.Object) object.Object {
		if len(args) != 1 {
			return newError("directory_naam requires 1 argument (path)")
		}
		if args[0].Type() != object.STRING_OBJ {
			return newError("path must be STRING, got %s", args[0].Type())
		}

		path := args[0].(*object.String).Value
		return &object.String{Value: filepath.Dir(path)}
	})

	// file_ext (ফাইল এক্সটেনশন) - Get file extension
	registerBuiltin("file_ext", func(args ...object.Object) object.Object {
		if len(args) != 1 {
			return newError("file_ext requires 1 argument (path)")
		}
		if args[0].Type() != object.STRING_OBJ {
			return newError("path must be STRING, got %s", args[0].Type())
		}

		path := args[0].(*object.String).Value
		return &object.String{Value: filepath.Ext(path)}
	})

	// path_bichchhed (পাথ বিচ্ছেদ) - Split path into directory and file
	// Returns array: [directory, file]
	registerBuiltin("path_bichchhed", func(args ...object.Object) object.Object {
		if len(args) != 1 {
			return newError("path_bichchhed requires 1 argument (path)")
		}
		if args[0].Type() != object.STRING_OBJ {
			return newError("path must be STRING, got %s", args[0].Type())
		}

		path := args[0].(*object.String).Value
		dir, file := filepath.Split(path)

		return &object.Array{
			Elements: []object.Object{
				&object.String{Value: strings.TrimSuffix(dir, string(filepath.Separator))},
				&object.String{Value: file},
			},
		}
	})

	// path_match (পাথ ম্যাচ) - Match path against glob pattern
	registerBuiltin("path_match", func(args ...object.Object) object.Object {
		if len(args) != 2 {
			return newError("path_match requires 2 arguments (pattern, path)")
		}
		if args[0].Type() != object.STRING_OBJ || args[1].Type() != object.STRING_OBJ {
			return newError("both arguments must be STRING")
		}

		pattern := args[0].(*object.String).Value
		path := args[1].(*object.String).Value

		matched, err := filepath.Match(pattern, path)
		if err != nil {
			return newError("invalid pattern: %s", err.Error())
		}

		if matched {
			return object.TRUE
		}
		return object.FALSE
	})
}
