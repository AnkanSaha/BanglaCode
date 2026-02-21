package evaluator

import (
	"BanglaCode/src/ast"
	"BanglaCode/src/evaluator/builtins"
	"BanglaCode/src/object"
)

// evalProgram evaluates a complete program
func evalProgram(stmts []ast.Statement, env *object.Environment) object.Object {
	var result object.Object

	for _, statement := range stmts {
		result = Eval(statement, env)

		switch result := result.(type) {
		case *object.ReturnValue:
			return result.Value
		case *object.Error:
			return result
		}
	}

	return result
}

// evalBlockStatement evaluates a block of statements
func evalBlockStatement(block *ast.BlockStatement, env *object.Environment) object.Object {
	var result object.Object

	for _, statement := range block.Statements {
		result = Eval(statement, env)

		if result != nil {
			rt := result.Type()
			if rt == object.RETURN_OBJ || rt == object.ERROR_OBJ ||
				rt == object.BREAK_OBJ || rt == object.CONTINUE_OBJ ||
				rt == object.EXCEPTION_OBJ {
				return result
			}
		}
	}

	return result
}

// evalIfStatement evaluates if/else statements
func evalIfStatement(ie *ast.IfStatement, env *object.Environment) object.Object {
	condition := Eval(ie.Condition, env)
	if isError(condition) {
		return condition
	}

	if isTruthy(condition) {
		return Eval(ie.Consequence, env)
	} else if ie.Alternative != nil {
		return Eval(ie.Alternative, env)
	}

	return object.NULL
}

// evalWhileStatement evaluates while loops
func evalWhileStatement(ws *ast.WhileStatement, env *object.Environment) object.Object {
	for {
		condition := Eval(ws.Condition, env)
		if isError(condition) {
			return condition
		}

		if !isTruthy(condition) {
			break
		}

		result := Eval(ws.Body, env)

		if result != nil {
			switch result.Type() {
			case object.RETURN_OBJ:
				return result
			case object.ERROR_OBJ:
				return result
			case object.BREAK_OBJ:
				return object.NULL
			case object.CONTINUE_OBJ:
				continue
			case object.EXCEPTION_OBJ:
				return result
			}
		}
	}

	return object.NULL
}

// evalForStatement evaluates for loops
func evalForStatement(fs *ast.ForStatement, env *object.Environment) object.Object {
	// Create new scope for loop
	loopEnv := object.NewEnclosedEnvironment(env)

	// Initialize
	if fs.Init != nil {
		result := Eval(fs.Init, loopEnv)
		if isError(result) {
			return result
		}
	}

	for {
		// Check condition
		if fs.Condition != nil {
			condition := Eval(fs.Condition, loopEnv)
			if isError(condition) {
				return condition
			}
			if !isTruthy(condition) {
				break
			}
		}

		// Execute body
		result := Eval(fs.Body, loopEnv)

		if result != nil {
			switch result.Type() {
			case object.RETURN_OBJ:
				return result
			case object.ERROR_OBJ:
				return result
			case object.BREAK_OBJ:
				return object.NULL
			case object.CONTINUE_OBJ:
				// Continue to update step
			case object.EXCEPTION_OBJ:
				return result
			}
		}

		// Update
		if fs.Update != nil {
			result := Eval(fs.Update, loopEnv)
			if isError(result) {
				return result
			}
		}
	}

	return object.NULL
}

// evalIdentifier evaluates variable references
func evalIdentifier(node *ast.Identifier, env *object.Environment) object.Object {
	if val, ok := env.Get(node.Value); ok {
		return val
	}

	if builtin, ok := builtins.Builtins[node.Value]; ok {
		return builtin
	}

	return newErrorAt(node.Token.Line, node.Token.Column, "variable '%s' is not defined", node.Value)
}

// evalExpressions evaluates a list of expressions, handling spread elements
func evalExpressions(exps []ast.Expression, env *object.Environment) []object.Object {
	var result []object.Object

	for _, e := range exps {
		// Check for spread element
		if spread, ok := e.(*ast.SpreadElement); ok {
			evaluated := Eval(spread.Argument, env)
			if isError(evaluated) {
				return []object.Object{evaluated}
			}
			// Spread must be an array
			if arr, ok := evaluated.(*object.Array); ok {
				result = append(result, arr.Elements...)
			} else {
				return []object.Object{newError("spread operator requires an array, got %s", evaluated.Type())}
			}
		} else {
			evaluated := Eval(e, env)
			if isError(evaluated) {
				return []object.Object{evaluated}
			}
			result = append(result, evaluated)
		}
	}

	return result
}

// evalSwitchStatement evaluates a switch statement
func evalSwitchStatement(node *ast.SwitchStatement, env *object.Environment) object.Object {
	// Evaluate the switch expression once
	switchValue := Eval(node.Expr, env)
	if isError(switchValue) {
		return switchValue
	}

	// Try each case
	for _, caseClause := range node.Cases {
		caseValue := Eval(caseClause.Value, env)
		if isError(caseValue) {
			return caseValue
		}

		// Check if values are equal using objectsEqual from helpers
		if objectsEqual(switchValue, caseValue) {
			result := Eval(caseClause.Body, env)

			// If result is break, exit switch (return NULL)
			if result == object.BREAK {
				return object.NULL
			}

			return result
		}
	}

	// Execute default case if no match
	if node.Default != nil {
		return Eval(node.Default, env)
	}

	return object.NULL
}
