package system

import (
	"BanglaCode/src/object"
	"os"
	"path/filepath"
)

func init() {
	// ==================== Directory Operations (NEW) ====================

	// directory_taliika (ডিরেক্টরি তালিকা) - List directory contents
	// Returns array of filenames
	registerBuiltin("directory_taliika", func(args ...object.Object) object.Object {
		if len(args) != 1 {
			return newError("directory_taliika requires 1 argument (path)")
		}
		if args[0].Type() != object.STRING_OBJ {
			return newError("path must be STRING, got %s", args[0].Type())
		}

		path := args[0].(*object.String).Value
		entries, err := os.ReadDir(path)
		if err != nil {
			return newError("failed to read directory: %s", err.Error())
		}

		// Pre-allocate array
		elements := make([]object.Object, 0, len(entries))
		for _, entry := range entries {
			elements = append(elements, &object.String{Value: entry.Name()})
		}

		return &object.Array{Elements: elements}
	})

	// directory_ghumao (ডিরেক্টরি ঘুরাও) - Recursive directory traversal
	// Returns array of all file paths in directory tree
	registerBuiltin("directory_ghumao", func(args ...object.Object) object.Object {
		if len(args) != 1 {
			return newError("directory_ghumao requires 1 argument (path)")
		}
		if args[0].Type() != object.STRING_OBJ {
			return newError("path must be STRING, got %s", args[0].Type())
		}

		root := args[0].(*object.String).Value
		files := make([]object.Object, 0, 100) // Pre-allocate

		err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				// Skip errors, don't fail entire walk
				return nil
			}
			files = append(files, &object.String{Value: path})
			return nil
		})

		if err != nil {
			return newError("failed to walk directory: %s", err.Error())
		}

		return &object.Array{Elements: files}
	})

	// directory_khali_ki (ডিরেক্টরি খালি কি) - Check if directory is empty
	registerBuiltin("directory_khali_ki", func(args ...object.Object) object.Object {
		if len(args) != 1 {
			return newError("directory_khali_ki requires 1 argument (path)")
		}
		if args[0].Type() != object.STRING_OBJ {
			return newError("path must be STRING, got %s", args[0].Type())
		}

		path := args[0].(*object.String).Value
		entries, err := os.ReadDir(path)
		if err != nil {
			return newError("failed to read directory: %s", err.Error())
		}

		if len(entries) == 0 {
			return object.TRUE
		}
		return object.FALSE
	})

	// directory_akar (ডিরেক্টরি আকার) - Get total size of directory
	// Returns total size in bytes
	registerBuiltin("directory_akar", func(args ...object.Object) object.Object {
		if len(args) != 1 {
			return newError("directory_akar requires 1 argument (path)")
		}
		if args[0].Type() != object.STRING_OBJ {
			return newError("path must be STRING, got %s", args[0].Type())
		}

		root := args[0].(*object.String).Value
		var totalSize int64

		err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				// Skip errors, don't fail entire walk
				return nil
			}
			if !info.IsDir() {
				totalSize += info.Size()
			}
			return nil
		})

		if err != nil {
			return newError("failed to calculate directory size: %s", err.Error())
		}

		return &object.Number{Value: float64(totalSize)}
	})
}
