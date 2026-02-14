package system

import (
	"BanglaCode/src/object"
	"os"
)

func init() {
	// ==================== Symbolic Links ====================

	// symlink_banao (সিমলিংক বানাও) - Create symbolic link
	registerBuiltin("symlink_banao", func(args ...object.Object) object.Object {
		if len(args) != 2 {
			return newError("symlink_banao requires 2 arguments (target, linkname)")
		}
		if args[0].Type() != object.STRING_OBJ || args[1].Type() != object.STRING_OBJ {
			return newError("both arguments must be STRING")
		}

		target := args[0].(*object.String).Value
		linkname := args[1].(*object.String).Value

		if err := os.Symlink(target, linkname); err != nil {
			return newError("failed to create symlink: %s", err.Error())
		}

		return object.NULL
	})

	// symlink_poro (সিমলিংক পড়ো) - Read symlink target
	registerBuiltin("symlink_poro", func(args ...object.Object) object.Object {
		if len(args) != 1 {
			return newError("symlink_poro requires 1 argument (linkname)")
		}
		if args[0].Type() != object.STRING_OBJ {
			return newError("linkname must be STRING, got %s", args[0].Type())
		}

		linkname := args[0].(*object.String).Value
		target, err := os.Readlink(linkname)
		if err != nil {
			return newError("failed to read symlink: %s", err.Error())
		}

		return &object.String{Value: target}
	})

	// symlink_ki (সিমলিংক কি) - Check if path is a symlink
	registerBuiltin("symlink_ki", func(args ...object.Object) object.Object {
		if len(args) != 1 {
			return newError("symlink_ki requires 1 argument (path)")
		}
		if args[0].Type() != object.STRING_OBJ {
			return newError("path must be STRING, got %s", args[0].Type())
		}

		path := args[0].(*object.String).Value
		info, err := os.Lstat(path)
		if err != nil {
			return object.FALSE
		}

		if info.Mode()&os.ModeSymlink != 0 {
			return object.TRUE
		}
		return object.FALSE
	})

	// hardlink_banao (হার্ডলিংক বানাও) - Create hard link
	registerBuiltin("hardlink_banao", func(args ...object.Object) object.Object {
		if len(args) != 2 {
			return newError("hardlink_banao requires 2 arguments (target, linkname)")
		}
		if args[0].Type() != object.STRING_OBJ || args[1].Type() != object.STRING_OBJ {
			return newError("both arguments must be STRING")
		}

		target := args[0].(*object.String).Value
		linkname := args[1].(*object.String).Value

		if err := os.Link(target, linkname); err != nil {
			return newError("failed to create hard link: %s", err.Error())
		}

		return object.NULL
	})

	// link_sonkha (লিংক সংখ্যা) - Get number of hard links
	registerBuiltin("link_sonkha", func(args ...object.Object) object.Object {
		if len(args) != 1 {
			return newError("link_sonkha requires 1 argument (path)")
		}
		if args[0].Type() != object.STRING_OBJ {
			return newError("path must be STRING, got %s", args[0].Type())
		}

		path := args[0].(*object.String).Value
		info, err := os.Stat(path)
		if err != nil {
			return newError("failed to get file info: %s", err.Error())
		}

		// Get link count using platform-specific helper
		if nlink, ok := getFileLinkCount(info); ok {
			return &object.Number{Value: float64(nlink)}
		}

		// Fallback
		return &object.Number{Value: 1}
	})
}
