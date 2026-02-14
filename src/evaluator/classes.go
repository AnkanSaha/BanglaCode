package evaluator

import (
	"BanglaCode/src/ast"
	"BanglaCode/src/object"
)

// evalClassDeclaration evaluates class declarations
func evalClassDeclaration(cd *ast.ClassDeclaration, env *object.Environment) object.Object {
	class := &object.Class{
		Name:    cd.Name.Value,
		Methods: make(map[string]*object.Function),
	}

	// Create class environment for methods
	classEnv := object.NewEnclosedEnvironment(env)

	// Evaluate methods (Methods is a slice of FunctionLiterals)
	for _, method := range cd.Methods {
		methodName := ""
		if method.Name != nil {
			methodName = method.Name.Value
		}
		fn := &object.Function{
			Parameters: method.Parameters,
			Body:       method.Body,
			Env:        classEnv,
			Name:       methodName,
		}
		class.Methods[methodName] = fn
	}

	env.Set(cd.Name.Value, class)
	return class
}

// evalNewExpression evaluates new expressions (object instantiation)
func evalNewExpression(ne *ast.NewExpression, env *object.Environment) object.Object {
	// Get the class
	classObj := Eval(ne.Class, env)
	if isError(classObj) {
		return classObj
	}

	class, ok := classObj.(*object.Class)
	if !ok {
		return newError("'%s' is not a class", ne.Class.String())
	}

	// Create instance
	instance := &object.Instance{
		Class:      class,
		Properties: make(map[string]object.Object),
	}

	// Evaluate constructor arguments
	args := evalExpressions(ne.Arguments, env)
	if len(args) == 1 && isError(args[0]) {
		return args[0]
	}

	// Call constructor if exists (method named "shuru")
	if constructor, ok := class.Methods["shuru"]; ok {
		// Check argument count
		if len(args) != len(constructor.Parameters) {
			return newError("constructor expects %d argument(s), got %d",
				len(constructor.Parameters), len(args))
		}

		// Create constructor environment
		constructorEnv := object.NewEnclosedEnvironment(constructor.Env)
		constructorEnv.Set("ei", instance)

		// Bind parameters
		for i, param := range constructor.Parameters {
			constructorEnv.Set(param.Value, args[i])
		}

		// Execute constructor
		result := Eval(constructor.Body, constructorEnv)
		if isError(result) {
			return result
		}
		if isException(result) {
			return result
		}
	}

	return instance
}

// applyFunction applies a function to arguments (wrapper for backward compatibility)
func applyFunction(fn object.Object, args []object.Object, env *object.Environment) object.Object {
	return applyFunctionWithPosition(fn, args, env, 0, 0, nil)
}

// applyFunctionWithPosition applies a function with position info for error reporting
func applyFunctionWithPosition(fn object.Object, args []object.Object, env *object.Environment, line, col int, callExpr ast.Expression) object.Object {
	switch fn := fn.(type) {
	case *object.Function:
		// Check argument count (considering rest parameter)
		minParams := len(fn.Parameters)
		actual := len(args)

		if fn.RestParameter == nil {
			// No rest parameter: exact match required
			if actual != minParams {
				funcName := fn.Name
				if funcName == "" {
					funcName = "anonymous function"
				}
				return newErrorAt(line, col, "function '%s' expects %d argument(s) but got %d", funcName, minParams, actual)
			}
		} else {
			// Has rest parameter: at least minParams required
			if actual < minParams {
				funcName := fn.Name
				if funcName == "" {
					funcName = "anonymous function"
				}
				return newErrorAt(line, col, "function '%s' expects at least %d argument(s) but got %d", funcName, minParams, actual)
			}
		}

		// Check if function is async - if so, execute in goroutine and return promise
		if fn.IsAsync {
			return evalAsyncFunctionCall(fn, args, env)
		}

		// Regular synchronous function execution
		extendedEnv := extendFunctionEnv(fn, args)
		evaluated := Eval(fn.Body, extendedEnv)
		return unwrapReturnValue(evaluated)

	case *object.Builtin:
		return fn.Fn(args...)

	default:
		// Better error for null/undefined
		funcName := "unknown"
		if ident, ok := callExpr.(*ast.Identifier); ok {
			funcName = ident.Value
		}
		if fn == nil || fn.Type() == object.NULL_OBJ {
			return newErrorAt(line, col, "'%s' is not defined or is null", funcName)
		}
		return newErrorAt(line, col, "'%s' is not a function (got %s)", funcName, fn.Type())
	}
}

// extendFunctionEnv creates a new environment for function execution
func extendFunctionEnv(fn *object.Function, args []object.Object) *object.Environment {
	env := object.NewEnclosedEnvironment(fn.Env)

	// Bind regular parameters
	for paramIdx, param := range fn.Parameters {
		if paramIdx < len(args) {
			env.Set(param.Value, args[paramIdx])
		}
	}

	// Bind rest parameter if present
	if fn.RestParameter != nil {
		restArgs := []object.Object{}
		if len(args) > len(fn.Parameters) {
			restArgs = args[len(fn.Parameters):]
		}
		env.Set(fn.RestParameter.Value, &object.Array{Elements: restArgs})
	}

	// Check if we need to bind 'ei' (this)
	if ei, ok := fn.Env.Get("ei"); ok {
		env.Set("ei", ei)
	}

	return env
}
