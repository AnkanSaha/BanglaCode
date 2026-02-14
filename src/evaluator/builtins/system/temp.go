package system

import (
	"BanglaCode/src/object"
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var (
	cachedTempDir string
	tempCacheOnce sync.Once
	// Track temp files/dirs created for cleanup
	tempFiles      []string
	tempFilesMutex sync.Mutex
)

func initTempCache() {
	tempCacheOnce.Do(func() {
		cachedTempDir = os.TempDir()
	})
}

func init() {
	initTempCache()

	// ==================== Temporary Files ====================

	// temp_directory (টেম্প ডিরেক্টরি) - Get system temp directory
	registerBuiltin("temp_directory", func(args ...object.Object) object.Object {
		return &object.String{Value: cachedTempDir}
	})

	// temp_file (টেম্প ফাইল) - Create temporary file
	// Returns path to created file
	registerBuiltin("temp_file", func(args ...object.Object) object.Object {
		// Optional prefix argument
		prefix := "banglacode-"
		if len(args) > 0 {
			if args[0].Type() != object.STRING_OBJ {
				return newError("prefix must be STRING, got %s", args[0].Type())
			}
			prefix = args[0].(*object.String).Value
		}

		file, err := os.CreateTemp("", prefix+"*")
		if err != nil {
			return newError("failed to create temp file: %s", err.Error())
		}

		path := file.Name()
		file.Close()

		// Track for cleanup
		tempFilesMutex.Lock()
		tempFiles = append(tempFiles, path)
		tempFilesMutex.Unlock()

		return &object.String{Value: path}
	})

	// temp_folder (টেম্প ফোল্ডার) - Create temporary directory
	// Returns path to created directory
	registerBuiltin("temp_folder", func(args ...object.Object) object.Object {
		// Optional prefix argument
		prefix := "banglacode-"
		if len(args) > 0 {
			if args[0].Type() != object.STRING_OBJ {
				return newError("prefix must be STRING, got %s", args[0].Type())
			}
			prefix = args[0].(*object.String).Value
		}

		dir, err := os.MkdirTemp("", prefix+"*")
		if err != nil {
			return newError("failed to create temp directory: %s", err.Error())
		}

		// Track for cleanup
		tempFilesMutex.Lock()
		tempFiles = append(tempFiles, dir)
		tempFilesMutex.Unlock()

		return &object.String{Value: dir}
	})

	// temp_path (টেম্প পাথ) - Generate unique temporary path
	// Does not create the file, just generates a unique name
	registerBuiltin("temp_path", func(args ...object.Object) object.Object {
		// Optional prefix argument
		prefix := "banglacode-"
		if len(args) > 0 {
			if args[0].Type() != object.STRING_OBJ {
				return newError("prefix must be STRING, got %s", args[0].Type())
			}
			prefix = args[0].(*object.String).Value
		}

		// Generate unique name using timestamp and random component
		name := fmt.Sprintf("%s%d", prefix, time.Now().UnixNano())
		path := filepath.Join(cachedTempDir, name)

		return &object.String{Value: path}
	})

	// temp_muche_felo (টেম্প মুছে ফেলো) - Clean up all temp files created
	registerBuiltin("temp_muche_felo", func(args ...object.Object) object.Object {
		tempFilesMutex.Lock()
		defer tempFilesMutex.Unlock()

		var errors []string
		for _, path := range tempFiles {
			if err := os.RemoveAll(path); err != nil {
				errors = append(errors, err.Error())
			}
		}

		// Clear the list
		tempFiles = nil

		if len(errors) > 0 {
			return newError("failed to clean some temp files: %v", errors)
		}

		return object.NULL
	})
}
