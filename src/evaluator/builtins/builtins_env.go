package builtins

import (
	"BanglaCode/src/object"
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

// Global environment variable storage with thread-safe access
var (
	envVars   = make(map[string]string)
	envVarsMu sync.RWMutex
)

func init() {
	// Register environment variable functions
	Builtins["env_load"] = &object.Builtin{Fn: builtinEnvLoad}
	Builtins["env_load_auto"] = &object.Builtin{Fn: builtinEnvLoadAuto}
	Builtins["env_get"] = &object.Builtin{Fn: builtinEnvGet}
	Builtins["env_get_default"] = &object.Builtin{Fn: builtinEnvGetDefault}
	Builtins["env_set"] = &object.Builtin{Fn: builtinEnvSet}
	Builtins["env_all"] = &object.Builtin{Fn: builtinEnvAll}
	Builtins["env_clear"] = &object.Builtin{Fn: builtinEnvClear}
}

// env_load - Load environment variables from a .env file
// Usage: env_load(".env") or env_load(".env.prod")
func builtinEnvLoad(args ...object.Object) object.Object {
	if len(args) != 1 {
		return newError("env_load: wrong number of arguments. got=%d, want=1", len(args))
	}

	filename, ok := args[0].(*object.String)
	if !ok {
		return newError("env_load: argument must be STRING (filename), got %s", args[0].Type())
	}

	// Load and parse .env file
	if err := loadEnvFile(filename.Value); err != nil {
		return newError("env_load: %s", err.Error())
	}

	return object.TRUE
}

// env_load_auto - Automatically load .env file based on environment
// Usage: env_load_auto("uat") -> tries .env.uat, then .env
// Usage: env_load_auto("prod") -> tries .env.prod, then .env
// Usage: env_load_auto("") -> loads .env
func builtinEnvLoadAuto(args ...object.Object) object.Object {
	if len(args) != 1 {
		return newError("env_load_auto: wrong number of arguments. got=%d, want=1", len(args))
	}

	envName, ok := args[0].(*object.String)
	if !ok {
		return newError("env_load_auto: argument must be STRING (environment), got %s", args[0].Type())
	}

	// Build filename based on environment
	var filenames []string

	if envName.Value != "" {
		// Try .env.{environment} first
		filenames = append(filenames, fmt.Sprintf(".env.%s", envName.Value))
	}

	// Always fallback to .env
	filenames = append(filenames, ".env")

	// Try loading files in order
	var lastErr error
	for _, filename := range filenames {
		if err := loadEnvFile(filename); err == nil {
			// Successfully loaded
			return &object.String{Value: filename}
		} else {
			lastErr = err
		}
	}

	// None of the files loaded successfully
	if lastErr != nil {
		return newError("env_load_auto: failed to load any .env file: %s", lastErr.Error())
	}

	return newError("env_load_auto: no .env files found")
}

// env_get - Get environment variable value
// Usage: dhoro api_key = env_get("API_KEY");
func builtinEnvGet(args ...object.Object) object.Object {
	if len(args) != 1 {
		return newError("env_get: wrong number of arguments. got=%d, want=1", len(args))
	}

	key, ok := args[0].(*object.String)
	if !ok {
		return newError("env_get: argument must be STRING (key), got %s", args[0].Type())
	}

	// Thread-safe read
	envVarsMu.RLock()
	value, exists := envVars[key.Value]
	envVarsMu.RUnlock()

	if !exists {
		return newError("env_get: environment variable '%s' not found", key.Value)
	}

	return &object.String{Value: value}
}

// env_get_default - Get environment variable with default fallback
// Usage: dhoro db_host = env_get_default("DB_HOST", "localhost");
func builtinEnvGetDefault(args ...object.Object) object.Object {
	if len(args) != 2 {
		return newError("env_get_default: wrong number of arguments. got=%d, want=2", len(args))
	}

	key, ok := args[0].(*object.String)
	if !ok {
		return newError("env_get_default: first argument must be STRING (key), got %s", args[0].Type())
	}

	defaultVal, ok := args[1].(*object.String)
	if !ok {
		return newError("env_get_default: second argument must be STRING (default), got %s", args[1].Type())
	}

	// Thread-safe read
	envVarsMu.RLock()
	value, exists := envVars[key.Value]
	envVarsMu.RUnlock()

	if !exists {
		return defaultVal
	}

	return &object.String{Value: value}
}

// env_set - Set environment variable at runtime
// Usage: env_set("API_KEY", "secret123");
func builtinEnvSet(args ...object.Object) object.Object {
	if len(args) != 2 {
		return newError("env_set: wrong number of arguments. got=%d, want=2", len(args))
	}

	key, ok := args[0].(*object.String)
	if !ok {
		return newError("env_set: first argument must be STRING (key), got %s", args[0].Type())
	}

	value, ok := args[1].(*object.String)
	if !ok {
		return newError("env_set: second argument must be STRING (value), got %s", args[1].Type())
	}

	// Thread-safe write
	envVarsMu.Lock()
	envVars[key.Value] = value.Value
	envVarsMu.Unlock()

	return object.TRUE
}

// env_all - Get all environment variables as a map
// Usage: dhoro all_env = env_all();
func builtinEnvAll(args ...object.Object) object.Object {
	if len(args) != 0 {
		return newError("env_all: wrong number of arguments. got=%d, want=0", len(args))
	}

	// Thread-safe read and copy
	envVarsMu.RLock()
	pairs := make(map[string]object.Object, len(envVars))
	for k, v := range envVars {
		pairs[k] = &object.String{Value: v}
	}
	envVarsMu.RUnlock()

	return &object.Map{Pairs: pairs}
}

// env_clear - Clear all loaded environment variables
// Usage: env_clear();
func builtinEnvClear(args ...object.Object) object.Object {
	if len(args) != 0 {
		return newError("env_clear: wrong number of arguments. got=%d, want=0", len(args))
	}

	// Thread-safe clear
	envVarsMu.Lock()
	envVars = make(map[string]string)
	envVarsMu.Unlock()

	return object.TRUE
}

// Helper function to load and parse .env file
// Optimized for performance with minimal allocations
func loadEnvFile(filename string) error {
	// Resolve absolute path
	absPath, err := filepath.Abs(filename)
	if err != nil {
		return fmt.Errorf("invalid path: %v", err)
	}

	// Check if file exists
	if _, err := os.Stat(absPath); os.IsNotExist(err) {
		return fmt.Errorf("file not found: %s", filename)
	}

	// Open file for reading
	file, err := os.Open(absPath)
	if err != nil {
		return fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	// Pre-allocate map with reasonable capacity
	newVars := make(map[string]string, 32)

	// Parse file line by line (memory-efficient for large files)
	scanner := bufio.NewScanner(file)
	lineNum := 0

	for scanner.Scan() {
		lineNum++
		line := strings.TrimSpace(scanner.Text())

		// Skip empty lines and comments
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		// Parse KEY=VALUE format
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			return fmt.Errorf("invalid format at line %d: %s", lineNum, line)
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		// Remove surrounding quotes if present
		value = strings.Trim(value, `"'`)

		// Validate key (must be valid identifier)
		if key == "" {
			return fmt.Errorf("empty key at line %d", lineNum)
		}

		newVars[key] = value
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error reading file: %v", err)
	}

	// Thread-safe merge into global env vars
	envVarsMu.Lock()
	for k, v := range newVars {
		envVars[k] = v
	}
	envVarsMu.Unlock()

	return nil
}
