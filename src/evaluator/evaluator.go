// Package evaluator implements the tree-walking interpreter for BanglaCode.
// It evaluates AST nodes and produces object values.
package evaluator

import (
	"BanglaCode/src/ast"
	"BanglaCode/src/object"
)

func init() {
	// Set up EvalFunc for builtins that need to call back into the evaluator
	EvalFunc = evalFunctionCall
}

// evalFunctionCall evaluates a function with the given arguments
// Used by builtins that need to call back into the evaluator
func evalFunctionCall(handler *object.Function, args []object.Object) object.Object {
	env := object.NewEnclosedEnvironment(handler.Env)
	for i, param := range handler.Parameters {
		if i < len(args) {
			env.Set(param.Value, args[i])
		}
	}
	result := Eval(handler.Body, env)
	return unwrapReturnValue(result)
}

// Eval evaluates an AST node and returns the resulting object
func Eval(node ast.Node, env *object.Environment) object.Object {
	switch node := node.(type) {

	// ==================== Statements ====================

	case *ast.Program:
		return evalProgram(node.Statements, env)

	case *ast.ExpressionStatement:
		return Eval(node.Expression, env)

	case *ast.BlockStatement:
		return evalBlockStatement(node, env)

	case *ast.VariableDeclaration:
		val := Eval(node.Value, env)
		if isError(val) {
			return val
		}
		if node.IsConstant {
			env.SetConstant(node.Name.Value, val)
		} else if node.IsGlobal {
			env.SetGlobal(node.Name.Value, val)
		} else {
			env.Set(node.Name.Value, val)
		}
		return val

	case *ast.IfStatement:
		return evalIfStatement(node, env)

	case *ast.WhileStatement:
		return evalWhileStatement(node, env)

	case *ast.ForStatement:
		return evalForStatement(node, env)

	case *ast.ReturnStatement:
		val := Eval(node.ReturnValue, env)
		if isError(val) {
			return val
		}
		return &object.ReturnValue{Value: val}

	case *ast.BreakStatement:
		return object.BREAK

	case *ast.ContinueStatement:
		return object.CONTINUE

	// ==================== Classes & Modules ====================

	case *ast.ClassDeclaration:
		return evalClassDeclaration(node, env)

	case *ast.ImportStatement:
		return evalImportStatement(node, env)

	case *ast.ExportStatement:
		return evalExportStatement(node, env)

	// ==================== Error Handling ====================

	case *ast.TryCatchStatement:
		return evalTryCatchStatement(node, env)

	case *ast.ThrowStatement:
		return evalThrowStatement(node, env)

	// ==================== Literals ====================

	case *ast.NumberLiteral:
		return &object.Number{Value: node.Value}

	case *ast.StringLiteral:
		return &object.String{Value: node.Value}

	case *ast.BooleanLiteral:
		return object.NativeBoolToBooleanObject(node.Value)

	case *ast.NullLiteral:
		return object.NULL

	case *ast.ArrayLiteral:
		elements := evalExpressions(node.Elements, env)
		if len(elements) == 1 && isError(elements[0]) {
			return elements[0]
		}
		return &object.Array{Elements: elements}

	case *ast.MapLiteral:
		return evalMapLiteral(node, env)

	// ==================== Expressions ====================

	case *ast.Identifier:
		return evalIdentifier(node, env)

	case *ast.UnaryExpression:
		right := Eval(node.Right, env)
		if isError(right) {
			return right
		}
		return evalUnaryExpression(node.Operator, right)

	case *ast.BinaryExpression:
		left := Eval(node.Left, env)
		if isError(left) {
			return left
		}
		right := Eval(node.Right, env)
		if isError(right) {
			return right
		}
		return evalBinaryExpression(node.Operator, left, right)

	case *ast.AssignmentExpression:
		return evalAssignmentExpression(node, env)

	case *ast.CallExpression:
		function := Eval(node.Function, env)
		if isError(function) {
			return function
		}
		args := evalExpressions(node.Arguments, env)
		if len(args) == 1 && isError(args[0]) {
			return args[0]
		}
		return applyFunctionWithPosition(function, args, env, node.Token.Line, node.Token.Column, node.Function)

	case *ast.MemberExpression:
		return evalMemberExpression(node, env)

	case *ast.FunctionLiteral:
		params := node.Parameters
		body := node.Body
		name := ""
		if node.Name != nil {
			name = node.Name.Value
		}
		fn := &object.Function{Parameters: params, Env: env, Body: body, Name: name}
		if name != "" {
			env.Set(name, fn)
		}
		return fn

	case *ast.NewExpression:
		return evalNewExpression(node, env)
	}

	return nil
}
