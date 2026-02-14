package system

import (
	"BanglaCode/src/object"
	"os"
)

func init() {
	// ==================== Environment Variables ====================

	// poribesh (পরিবেশ) - Get environment variable
	registerBuiltin("poribesh", func(args ...object.Object) object.Object {
		if len(args) != 1 {
			return newError("poribesh requires 1 argument (variable name)")
		}
		if args[0].Type() != object.STRING_OBJ {
			return newError("environment variable name must be STRING, got %s", args[0].Type())
		}

		name := args[0].(*object.String).Value
		value := os.Getenv(name)
		return &object.String{Value: value}
	})

	// poribesh_set (পরিবেশ সেট) - Set environment variable
	registerBuiltin("poribesh_set", func(args ...object.Object) object.Object {
		if len(args) != 2 {
			return newError("poribesh_set requires 2 arguments (name, value)")
		}
		if args[0].Type() != object.STRING_OBJ || args[1].Type() != object.STRING_OBJ {
			return newError("both arguments must be STRING")
		}

		name := args[0].(*object.String).Value
		value := args[1].(*object.String).Value

		if err := os.Setenv(name, value); err != nil {
			return newError("failed to set environment variable: %s", err.Error())
		}

		return object.NULL
	})

	// poribesh_shokal (পরিবেশ সকল) - Get all environment variables
	// Returns array of strings in "KEY=VALUE" format
	registerBuiltin("poribesh_shokal", func(args ...object.Object) object.Object {
		envVars := os.Environ()
		elements := make([]object.Object, len(envVars))
		for i, env := range envVars {
			elements[i] = &object.String{Value: env}
		}
		return &object.Array{Elements: elements}
	})

	// poribesh_muke (পরিবেশ মুছে) - Unset environment variable
	registerBuiltin("poribesh_muke", func(args ...object.Object) object.Object {
		if len(args) != 1 {
			return newError("poribesh_muke requires 1 argument (variable name)")
		}
		if args[0].Type() != object.STRING_OBJ {
			return newError("environment variable name must be STRING, got %s", args[0].Type())
		}

		name := args[0].(*object.String).Value
		if err := os.Unsetenv(name); err != nil {
			return newError("failed to unset environment variable: %s", err.Error())
		}

		return object.NULL
	})
}
