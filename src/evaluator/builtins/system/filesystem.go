package system

import (
	"BanglaCode/src/object"
	"fmt"
	"os"
	"os/user"
	"strconv"
	"syscall"
)

func init() {
	// ==================== File System Operations ====================

	// ache_ki (আছে কি) - Check if file or directory exists
	registerBuiltin("ache_ki", func(args ...object.Object) object.Object {
		if len(args) != 1 {
			return newError("ache_ki requires 1 argument (path)")
		}
		if args[0].Type() != object.STRING_OBJ {
			return newError("path must be STRING, got %s", args[0].Type())
		}

		path := args[0].(*object.String).Value
		_, err := os.Stat(path)
		if err == nil {
			return object.TRUE
		}
		if os.IsNotExist(err) {
			return object.FALSE
		}
		// Other error (permission denied, etc.)
		return object.FALSE
	})

	// folder_banao (ফোল্ডার বানাও) - Create directory
	registerBuiltin("folder_banao", func(args ...object.Object) object.Object {
		if len(args) != 1 {
			return newError("folder_banao requires 1 argument (path)")
		}
		if args[0].Type() != object.STRING_OBJ {
			return newError("path must be STRING, got %s", args[0].Type())
		}

		path := args[0].(*object.String).Value
		if err := os.Mkdir(path, 0755); err != nil {
			return newError("failed to create directory: %s", err.Error())
		}

		return object.NULL
	})

	// folder_banao_shokal (ফোল্ডার বানাও সকল) - Create directory with all parent directories
	registerBuiltin("folder_banao_shokal", func(args ...object.Object) object.Object {
		if len(args) != 1 {
			return newError("folder_banao_shokal requires 1 argument (path)")
		}
		if args[0].Type() != object.STRING_OBJ {
			return newError("path must be STRING, got %s", args[0].Type())
		}

		path := args[0].(*object.String).Value
		if err := os.MkdirAll(path, 0755); err != nil {
			return newError("failed to create directories: %s", err.Error())
		}

		return object.NULL
	})

	// muke_felo (মুছে ফেলো) - Remove file or directory
	registerBuiltin("muke_felo", func(args ...object.Object) object.Object {
		if len(args) != 1 {
			return newError("muke_felo requires 1 argument (path)")
		}
		if args[0].Type() != object.STRING_OBJ {
			return newError("path must be STRING, got %s", args[0].Type())
		}

		path := args[0].(*object.String).Value
		if err := os.RemoveAll(path); err != nil {
			return newError("failed to remove: %s", err.Error())
		}

		return object.NULL
	})

	// ==================== File Metadata (NEW) ====================

	// file_akar (ফাইল আকার) - Get file size in bytes
	registerBuiltin("file_akar", func(args ...object.Object) object.Object {
		if len(args) != 1 {
			return newError("file_akar requires 1 argument (path)")
		}
		if args[0].Type() != object.STRING_OBJ {
			return newError("path must be STRING, got %s", args[0].Type())
		}

		path := args[0].(*object.String).Value
		info, err := os.Stat(path)
		if err != nil {
			return newError("failed to get file info: %s", err.Error())
		}

		return &object.Number{Value: float64(info.Size())}
	})

	// file_permission (ফাইল পারমিশন) - Get file permissions
	registerBuiltin("file_permission", func(args ...object.Object) object.Object {
		if len(args) != 1 {
			return newError("file_permission requires 1 argument (path)")
		}
		if args[0].Type() != object.STRING_OBJ {
			return newError("path must be STRING, got %s", args[0].Type())
		}

		path := args[0].(*object.String).Value
		info, err := os.Stat(path)
		if err != nil {
			return newError("failed to get file info: %s", err.Error())
		}

		// Return permissions as octal string (e.g., "0644")
		perms := fmt.Sprintf("%04o", info.Mode().Perm())
		return &object.String{Value: perms}
	})

	// file_permission_set (ফাইল পারমিশন সেট) - Change file permissions
	registerBuiltin("file_permission_set", func(args ...object.Object) object.Object {
		if len(args) != 2 {
			return newError("file_permission_set requires 2 arguments (path, permissions)")
		}
		if args[0].Type() != object.STRING_OBJ {
			return newError("path must be STRING, got %s", args[0].Type())
		}

		path := args[0].(*object.String).Value

		// Parse permissions - can be number or string
		var mode os.FileMode
		switch arg := args[1].(type) {
		case *object.Number:
			mode = os.FileMode(uint32(arg.Value))
		case *object.String:
			// Parse octal string
			val, err := strconv.ParseUint(arg.Value, 8, 32)
			if err != nil {
				return newError("invalid permission format: %s", err.Error())
			}
			mode = os.FileMode(val)
		default:
			return newError("permissions must be NUMBER or STRING, got %s", args[1].Type())
		}

		if err := os.Chmod(path, mode); err != nil {
			return newError("failed to change permissions: %s", err.Error())
		}

		return object.NULL
	})

	// file_malikan (ফাইল মালিকান) - Get file owner info
	// Returns: { "uid": number, "gid": number, "naam": string }
	registerBuiltin("file_malikan", func(args ...object.Object) object.Object {
		if len(args) != 1 {
			return newError("file_malikan requires 1 argument (path)")
		}
		if args[0].Type() != object.STRING_OBJ {
			return newError("path must be STRING, got %s", args[0].Type())
		}

		path := args[0].(*object.String).Value
		info, err := os.Stat(path)
		if err != nil {
			return newError("failed to get file info: %s", err.Error())
		}

		result := make(map[string]object.Object)

		// Get UID and GID from system-specific stat
		if stat, ok := info.Sys().(*syscall.Stat_t); ok {
			result["uid"] = &object.Number{Value: float64(stat.Uid)}
			result["gid"] = &object.Number{Value: float64(stat.Gid)}

			// Try to get username from UID
			if u, err := user.LookupId(fmt.Sprintf("%d", stat.Uid)); err == nil {
				result["naam"] = &object.String{Value: u.Username}
			} else {
				result["naam"] = &object.String{Value: ""}
			}
		} else {
			// Platform doesn't support UID/GID
			result["uid"] = &object.Number{Value: 0}
			result["gid"] = &object.Number{Value: 0}
			result["naam"] = &object.String{Value: ""}
		}

		return &object.Map{Pairs: result}
	})

	// file_malikan_set (ফাইল মালিকান সেট) - Change file owner
	registerBuiltin("file_malikan_set", func(args ...object.Object) object.Object {
		if len(args) != 3 {
			return newError("file_malikan_set requires 3 arguments (path, uid, gid)")
		}
		if args[0].Type() != object.STRING_OBJ {
			return newError("path must be STRING, got %s", args[0].Type())
		}
		if args[1].Type() != object.NUMBER_OBJ || args[2].Type() != object.NUMBER_OBJ {
			return newError("uid and gid must be NUMBER")
		}

		path := args[0].(*object.String).Value
		uid := int(args[1].(*object.Number).Value)
		gid := int(args[2].(*object.Number).Value)

		if err := os.Chown(path, uid, gid); err != nil {
			return newError("failed to change owner: %s", err.Error())
		}

		return object.NULL
	})

	// file_shomoy_poribortito (ফাইল সময় পরিবর্তিত) - Get file modified time (Unix timestamp)
	registerBuiltin("file_shomoy_poribortito", func(args ...object.Object) object.Object {
		if len(args) != 1 {
			return newError("file_shomoy_poribortito requires 1 argument (path)")
		}
		if args[0].Type() != object.STRING_OBJ {
			return newError("path must be STRING, got %s", args[0].Type())
		}

		path := args[0].(*object.String).Value
		info, err := os.Stat(path)
		if err != nil {
			return newError("failed to get file info: %s", err.Error())
		}

		return &object.Number{Value: float64(info.ModTime().Unix())}
	})

	// file_shomoy_access (ফাইল সময় এক্সেস) - Get file access time (Unix timestamp)
	registerBuiltin("file_shomoy_access", func(args ...object.Object) object.Object {
		if len(args) != 1 {
			return newError("file_shomoy_access requires 1 argument (path)")
		}
		if args[0].Type() != object.STRING_OBJ {
			return newError("path must be STRING, got %s", args[0].Type())
		}

		path := args[0].(*object.String).Value
		info, err := os.Stat(path)
		if err != nil {
			return newError("failed to get file info: %s", err.Error())
		}

		// Get access time from system-specific stat
		if stat, ok := info.Sys().(*syscall.Stat_t); ok {
			return &object.Number{Value: float64(stat.Atim.Sec)}
		}

		// Fallback to mod time if access time not available
		return &object.Number{Value: float64(info.ModTime().Unix())}
	})

	// file_shomoy_tori (ফাইল সময় তৈরি) - Get file creation time (Unix timestamp)
	registerBuiltin("file_shomoy_tori", func(args ...object.Object) object.Object {
		if len(args) != 1 {
			return newError("file_shomoy_tori requires 1 argument (path)")
		}
		if args[0].Type() != object.STRING_OBJ {
			return newError("path must be STRING, got %s", args[0].Type())
		}

		path := args[0].(*object.String).Value
		info, err := os.Stat(path)
		if err != nil {
			return newError("failed to get file info: %s", err.Error())
		}

		// Get creation time from system-specific stat
		if stat, ok := info.Sys().(*syscall.Stat_t); ok {
			return &object.Number{Value: float64(stat.Ctim.Sec)}
		}

		// Fallback to mod time if creation time not available
		return &object.Number{Value: float64(info.ModTime().Unix())}
	})

	// file_dhoron (ফাইল ধরন) - Get file type
	// Returns: "file", "directory", "symlink", "other"
	registerBuiltin("file_dhoron", func(args ...object.Object) object.Object {
		if len(args) != 1 {
			return newError("file_dhoron requires 1 argument (path)")
		}
		if args[0].Type() != object.STRING_OBJ {
			return newError("path must be STRING, got %s", args[0].Type())
		}

		path := args[0].(*object.String).Value
		info, err := os.Lstat(path) // Use Lstat to detect symlinks
		if err != nil {
			return newError("failed to get file info: %s", err.Error())
		}

		mode := info.Mode()
		switch {
		case mode.IsRegular():
			return &object.String{Value: "file"}
		case mode.IsDir():
			return &object.String{Value: "directory"}
		case mode&os.ModeSymlink != 0:
			return &object.String{Value: "symlink"}
		default:
			return &object.String{Value: "other"}
		}
	})

	// file_rename (ফাইল রিনেম) - Rename or move file
	registerBuiltin("file_rename", func(args ...object.Object) object.Object {
		if len(args) != 2 {
			return newError("file_rename requires 2 arguments (oldpath, newpath)")
		}
		if args[0].Type() != object.STRING_OBJ || args[1].Type() != object.STRING_OBJ {
			return newError("both arguments must be STRING")
		}

		oldpath := args[0].(*object.String).Value
		newpath := args[1].(*object.String).Value

		if err := os.Rename(oldpath, newpath); err != nil {
			return newError("failed to rename file: %s", err.Error())
		}

		return object.NULL
	})
}
