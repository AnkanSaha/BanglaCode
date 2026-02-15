package builtins

import (
	"BanglaCode/src/object"

	"BanglaCode/src/evaluator/builtins/database"
	"BanglaCode/src/evaluator/builtins/system"
)

// EvalFunc is a function pointer for evaluating AST nodes (set by evaluator.go to avoid circular dependency)
var EvalFunc func(handler *object.Function, args []object.Object) object.Object

// Builtins is the global map of built-in functions
// Individual built-in functions are registered in their respective files using init()
var Builtins = map[string]*object.Builtin{}

func init() {
	// Register system built-in functions
	for name, fn := range system.Builtins {
		Builtins[name] = fn
	}

	// Register database built-in functions
	for name, fn := range database.Builtins {
		Builtins[name] = fn
	}
}
