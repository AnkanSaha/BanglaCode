package system

import (
	"BanglaCode/src/object"
	"fmt"
)

// Builtins is the map that holds all system built-in functions
// Each file in this package registers its functions in this map via init()
var Builtins = make(map[string]*object.Builtin, 60)

// registerBuiltin is a helper function to register a built-in function
// Used by all files in this package to add their functions to the Builtins map
func registerBuiltin(name string, fn object.BuiltinFunction) {
	Builtins[name] = &object.Builtin{Fn: fn}
}

// newError creates an error object with a formatted message
// Helper function used across all system built-in implementations
func newError(format string, a ...interface{}) *object.Error {
	return &object.Error{Message: fmt.Sprintf(format, a...)}
}
