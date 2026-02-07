package evaluator

import (
	"BanglaCode/src/ast"
	"BanglaCode/src/lexer"
	"BanglaCode/src/object"
	"BanglaCode/src/parser"
	"fmt"
	"math"
	"os"
	"path/filepath"
	"sync"
)

// Module cache to prevent circular imports
var (
	moduleCache = make(map[string]*object.Module)
	moduleMutex sync.RWMutex
	currentDir  = "."
)

func init() {
	// Set up EvalFunc for builtins that need to call back into the evaluator
	EvalFunc = evalFunctionCall
}

// evalFunctionCall evaluates a function with the given arguments
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

// SetCurrentDir sets the current directory for resolving imports
func SetCurrentDir(dir string) {
	currentDir = dir
}

// Eval evaluates an AST node
func Eval(node ast.Node, env *object.Environment) object.Object {
	switch node := node.(type) {

	// Statements
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
		env.Set(node.Name.Value, val)
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

	case *ast.ClassDeclaration:
		return evalClassDeclaration(node, env)

	case *ast.BreakStatement:
		return object.BREAK

	case *ast.ContinueStatement:
		return object.CONTINUE

	case *ast.ImportStatement:
		return evalImportStatement(node, env)

	case *ast.ExportStatement:
		return evalExportStatement(node, env)

	case *ast.TryCatchStatement:
		return evalTryCatchStatement(node, env)

	case *ast.ThrowStatement:
		return evalThrowStatement(node, env)

	// Expressions
	case *ast.NumberLiteral:
		return &object.Number{Value: node.Value}

	case *ast.StringLiteral:
		return &object.String{Value: node.Value}

	case *ast.BooleanLiteral:
		return object.NativeBoolToBooleanObject(node.Value)

	case *ast.NullLiteral:
		return object.NULL

	case *ast.Identifier:
		return evalIdentifier(node, env)

	case *ast.ArrayLiteral:
		elements := evalExpressions(node.Elements, env)
		if len(elements) == 1 && isError(elements[0]) {
			return elements[0]
		}
		return &object.Array{Elements: elements}

	case *ast.MapLiteral:
		return evalMapLiteral(node, env)

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
		return applyFunction(function, args, env)

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

func evalWhileStatement(ws *ast.WhileStatement, env *object.Environment) object.Object {
	var result object.Object

	for {
		condition := Eval(ws.Condition, env)
		if isError(condition) {
			return condition
		}

		if !isTruthy(condition) {
			break
		}

		result = Eval(ws.Body, env)

		if result != nil {
			if result.Type() == object.RETURN_OBJ || result.Type() == object.ERROR_OBJ {
				return result
			}
			if result.Type() == object.BREAK_OBJ {
				break
			}
			if result.Type() == object.CONTINUE_OBJ {
				continue
			}
		}
	}

	return result
}

func evalForStatement(fs *ast.ForStatement, env *object.Environment) object.Object {
	var result object.Object

	// Create new scope for loop
	loopEnv := object.NewEnclosedEnvironment(env)

	// Initialize
	if fs.Init != nil {
		result = Eval(fs.Init, loopEnv)
		if isError(result) {
			return result
		}
	}

	// Loop
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
		result = Eval(fs.Body, loopEnv)

		if result != nil {
			if result.Type() == object.RETURN_OBJ || result.Type() == object.ERROR_OBJ {
				return result
			}
			if result.Type() == object.BREAK_OBJ {
				break
			}
			if result.Type() == object.CONTINUE_OBJ {
				// Update before continuing
				if fs.Update != nil {
					Eval(fs.Update, loopEnv)
				}
				continue
			}
		}

		// Update
		if fs.Update != nil {
			result = Eval(fs.Update, loopEnv)
			if isError(result) {
				return result
			}
		}
	}

	return result
}

func evalClassDeclaration(cd *ast.ClassDeclaration, env *object.Environment) object.Object {
	methods := make(map[string]*object.Function)

	for _, method := range cd.Methods {
		name := ""
		if method.Name != nil {
			name = method.Name.Value
		}
		fn := &object.Function{
			Parameters: method.Parameters,
			Body:       method.Body,
			Env:        env,
			Name:       name,
		}
		methods[name] = fn
	}

	class := &object.Class{
		Name:    cd.Name.Value,
		Methods: methods,
	}

	env.Set(cd.Name.Value, class)
	return class
}

func evalNewExpression(ne *ast.NewExpression, env *object.Environment) object.Object {
	class := Eval(ne.Class, env)
	if isError(class) {
		return class
	}

	classObj, ok := class.(*object.Class)
	if !ok {
		return newError("'%s' is not a class", ne.Class)
	}

	instance := &object.Instance{
		Class:      classObj,
		Properties: make(map[string]object.Object),
	}

	// Call init method if exists
	if initMethod, ok := classObj.Methods["init"]; ok {
		args := evalExpressions(ne.Arguments, env)
		if len(args) == 1 && isError(args[0]) {
			return args[0]
		}

		// Create instance environment with 'ei' (this) bound
		instanceEnv := object.NewEnclosedEnvironment(initMethod.Env)
		instanceEnv.Set("ei", instance)

		for paramIdx, param := range initMethod.Parameters {
			if paramIdx < len(args) {
				instanceEnv.Set(param.Value, args[paramIdx])
			}
		}

		result := Eval(initMethod.Body, instanceEnv)
		if isError(result) {
			return result
		}

		// Copy properties from instanceEnv to instance
		if ei, ok := instanceEnv.Get("ei"); ok {
			if inst, ok := ei.(*object.Instance); ok {
				instance.Properties = inst.Properties
			}
		}
	}

	return instance
}

func evalIdentifier(node *ast.Identifier, env *object.Environment) object.Object {
	if val, ok := env.Get(node.Value); ok {
		return val
	}

	if builtin, ok := builtins[node.Value]; ok {
		return builtin
	}

	return newError("%s", "identifier not found: " + node.Value)
}

func evalExpressions(exps []ast.Expression, env *object.Environment) []object.Object {
	var result []object.Object

	for _, e := range exps {
		evaluated := Eval(e, env)
		if isError(evaluated) {
			return []object.Object{evaluated}
		}
		result = append(result, evaluated)
	}

	return result
}

func evalMapLiteral(node *ast.MapLiteral, env *object.Environment) object.Object {
	pairs := make(map[string]object.Object)

	for keyNode, valueNode := range node.Pairs {
		key := Eval(keyNode, env)
		if isError(key) {
			return key
		}

		var keyStr string
		switch k := key.(type) {
		case *object.String:
			keyStr = k.Value
		case *object.Number:
			keyStr = k.Inspect()
		default:
			return newError("unusable as map key: %s", key.Type())
		}

		value := Eval(valueNode, env)
		if isError(value) {
			return value
		}

		pairs[keyStr] = value
	}

	return &object.Map{Pairs: pairs}
}

func evalUnaryExpression(operator string, right object.Object) object.Object {
	switch operator {
	case "!", "na":
		return evalBangOperator(right)
	case "-":
		return evalMinusOperator(right)
	default:
		return newError("unknown operator: %s%s", operator, right.Type())
	}
}

func evalBangOperator(right object.Object) object.Object {
	return object.NativeBoolToBooleanObject(!isTruthy(right))
}

func evalMinusOperator(right object.Object) object.Object {
	if right.Type() != object.NUMBER_OBJ {
		return newError("unknown operator: -%s", right.Type())
	}
	value := right.(*object.Number).Value
	return &object.Number{Value: -value}
}

func evalBinaryExpression(operator string, left, right object.Object) object.Object {
	switch {
	case left.Type() == object.NUMBER_OBJ && right.Type() == object.NUMBER_OBJ:
		return evalNumberBinaryExpression(operator, left, right)
	case left.Type() == object.STRING_OBJ && right.Type() == object.STRING_OBJ:
		return evalStringBinaryExpression(operator, left, right)
	case left.Type() == object.STRING_OBJ && right.Type() == object.NUMBER_OBJ:
		return evalStringNumberBinaryExpression(operator, left, right)
	case operator == "==":
		return object.NativeBoolToBooleanObject(objectsEqual(left, right))
	case operator == "!=":
		return object.NativeBoolToBooleanObject(!objectsEqual(left, right))
	case operator == "ebong":
		return object.NativeBoolToBooleanObject(isTruthy(left) && isTruthy(right))
	case operator == "ba":
		if isTruthy(left) {
			return left
		}
		return right
	default:
		return newError("unknown operator: %s %s %s", left.Type(), operator, right.Type())
	}
}

func evalNumberBinaryExpression(operator string, left, right object.Object) object.Object {
	leftVal := left.(*object.Number).Value
	rightVal := right.(*object.Number).Value

	switch operator {
	case "+":
		return &object.Number{Value: leftVal + rightVal}
	case "-":
		return &object.Number{Value: leftVal - rightVal}
	case "*":
		return &object.Number{Value: leftVal * rightVal}
	case "/":
		if rightVal == 0 {
			return newError("division by zero")
		}
		return &object.Number{Value: leftVal / rightVal}
	case "%":
		return &object.Number{Value: math.Mod(leftVal, rightVal)}
	case "<":
		return object.NativeBoolToBooleanObject(leftVal < rightVal)
	case ">":
		return object.NativeBoolToBooleanObject(leftVal > rightVal)
	case "<=":
		return object.NativeBoolToBooleanObject(leftVal <= rightVal)
	case ">=":
		return object.NativeBoolToBooleanObject(leftVal >= rightVal)
	case "==":
		return object.NativeBoolToBooleanObject(leftVal == rightVal)
	case "!=":
		return object.NativeBoolToBooleanObject(leftVal != rightVal)
	default:
		return newError("unknown operator: %s %s %s", left.Type(), operator, right.Type())
	}
}

func evalStringBinaryExpression(operator string, left, right object.Object) object.Object {
	leftVal := left.(*object.String).Value
	rightVal := right.(*object.String).Value

	switch operator {
	case "+":
		return &object.String{Value: leftVal + rightVal}
	case "==":
		return object.NativeBoolToBooleanObject(leftVal == rightVal)
	case "!=":
		return object.NativeBoolToBooleanObject(leftVal != rightVal)
	default:
		return newError("unknown operator: %s %s %s", left.Type(), operator, right.Type())
	}
}

func evalStringNumberBinaryExpression(operator string, left, right object.Object) object.Object {
	if operator == "+" {
		leftVal := left.(*object.String).Value
		rightVal := right.(*object.Number).Inspect()
		return &object.String{Value: leftVal + rightVal}
	}
	return newError("unknown operator: %s %s %s", left.Type(), operator, right.Type())
}

func evalAssignmentExpression(ae *ast.AssignmentExpression, env *object.Environment) object.Object {
	// Handle member expressions (obj.prop = val, arr[i] = val)
	if member, ok := ae.Name.(*ast.MemberExpression); ok {
		return evalMemberAssignment(member, ae.Operator, ae.Value, env)
	}

	// Handle simple identifiers
	ident, ok := ae.Name.(*ast.Identifier)
	if !ok {
		return newError("invalid assignment target")
	}

	val := Eval(ae.Value, env)
	if isError(val) {
		return val
	}

	// Handle compound assignment
	if ae.Operator != "=" {
		current, ok := env.Get(ident.Value)
		if !ok {
			return newError("%s", "identifier not found: " + ident.Value)
		}

		var operator string
		switch ae.Operator {
		case "+=":
			operator = "+"
		case "-=":
			operator = "-"
		case "*=":
			operator = "*"
		case "/=":
			operator = "/"
		}

		val = evalBinaryExpression(operator, current, val)
		if isError(val) {
			return val
		}
	}

	env.Update(ident.Value, val)
	return val
}

func evalMemberAssignment(member *ast.MemberExpression, operator string, value ast.Expression, env *object.Environment) object.Object {
	obj := Eval(member.Object, env)
	if isError(obj) {
		return obj
	}

	val := Eval(value, env)
	if isError(val) {
		return val
	}

	switch o := obj.(type) {
	case *object.Array:
		index := Eval(member.Property, env)
		if isError(index) {
			return index
		}
		if index.Type() != object.NUMBER_OBJ {
			return newError("array index must be a number")
		}
		idx := int(index.(*object.Number).Value)
		if idx < 0 || idx >= len(o.Elements) {
			return newError("array index out of bounds")
		}

		if operator != "=" {
			current := o.Elements[idx]
			var op string
			switch operator {
			case "+=":
				op = "+"
			case "-=":
				op = "-"
			case "*=":
				op = "*"
			case "/=":
				op = "/"
			}
			val = evalBinaryExpression(op, current, val)
			if isError(val) {
				return val
			}
		}

		o.Elements[idx] = val
		return val

	case *object.Map:
		var keyStr string
		if member.Computed {
			key := Eval(member.Property, env)
			if isError(key) {
				return key
			}
			keyStr = getMapKey(key)
		} else {
			if ident, ok := member.Property.(*ast.Identifier); ok {
				keyStr = ident.Value
			} else {
				return newError("invalid map key")
			}
		}

		if operator != "=" {
			current, ok := o.Pairs[keyStr]
			if !ok {
				current = object.NULL
			}
			var op string
			switch operator {
			case "+=":
				op = "+"
			case "-=":
				op = "-"
			case "*=":
				op = "*"
			case "/=":
				op = "/"
			}
			val = evalBinaryExpression(op, current, val)
			if isError(val) {
				return val
			}
		}

		o.Pairs[keyStr] = val
		return val

	case *object.Instance:
		if ident, ok := member.Property.(*ast.Identifier); ok {
			if operator != "=" {
				current, ok := o.Properties[ident.Value]
				if !ok {
					current = object.NULL
				}
				var op string
				switch operator {
				case "+=":
					op = "+"
				case "-=":
					op = "-"
				case "*=":
					op = "*"
				case "/=":
					op = "/"
				}
				val = evalBinaryExpression(op, current, val)
				if isError(val) {
					return val
				}
			}

			o.Properties[ident.Value] = val
			return val
		}
		return newError("invalid property name")

	default:
		return newError("cannot assign to property of %s", obj.Type())
	}
}

func evalMemberExpression(me *ast.MemberExpression, env *object.Environment) object.Object {
	obj := Eval(me.Object, env)
	if isError(obj) {
		return obj
	}

	switch o := obj.(type) {
	case *object.Array:
		index := Eval(me.Property, env)
		if isError(index) {
			return index
		}
		return evalArrayIndex(o, index)

	case *object.Map:
		var key string
		if me.Computed {
			keyObj := Eval(me.Property, env)
			if isError(keyObj) {
				return keyObj
			}
			key = getMapKey(keyObj)
		} else {
			if ident, ok := me.Property.(*ast.Identifier); ok {
				key = ident.Value
			} else {
				return newError("invalid map key")
			}
		}
		if val, ok := o.Pairs[key]; ok {
			return val
		}
		return object.NULL

	case *object.Instance:
		if ident, ok := me.Property.(*ast.Identifier); ok {
			// Check properties first
			if val, ok := o.Properties[ident.Value]; ok {
				return val
			}
			// Check methods
			if method, ok := o.Class.Methods[ident.Value]; ok {
				// Bind 'ei' (this) to instance by creating a new environment
				boundEnv := object.NewEnclosedEnvironment(method.Env)
				boundEnv.Set("ei", o)
				return &object.Function{
					Parameters: method.Parameters,
					Body:       method.Body,
					Env:        boundEnv,
					Name:       method.Name,
				}
			}
			return object.NULL
		}
		return newError("invalid property name")

	default:
		return newError("member access not supported on %s", obj.Type())
	}
}

func evalArrayIndex(array *object.Array, index object.Object) object.Object {
	if index.Type() != object.NUMBER_OBJ {
		return newError("array index must be a number, got %s", index.Type())
	}

	idx := int(index.(*object.Number).Value)
	max := len(array.Elements) - 1

	if idx < 0 || idx > max {
		return object.NULL
	}

	return array.Elements[idx]
}

func getMapKey(key object.Object) string {
	switch k := key.(type) {
	case *object.String:
		return k.Value
	case *object.Number:
		return k.Inspect()
	default:
		return ""
	}
}

func applyFunction(fn object.Object, args []object.Object, env *object.Environment) object.Object {
	switch fn := fn.(type) {
	case *object.Function:
		extendedEnv := extendFunctionEnv(fn, args)
		evaluated := Eval(fn.Body, extendedEnv)
		return unwrapReturnValue(evaluated)

	case *object.Builtin:
		return fn.Fn(args...)

	default:
		return newError("not a function: %s", fn.Type())
	}
}

func extendFunctionEnv(fn *object.Function, args []object.Object) *object.Environment {
	env := object.NewEnclosedEnvironment(fn.Env)

	for paramIdx, param := range fn.Parameters {
		if paramIdx < len(args) {
			env.Set(param.Value, args[paramIdx])
		}
	}

	// Check if we need to bind 'ei' (this)
	if ei, ok := fn.Env.Get("ei"); ok {
		env.Set("ei", ei)
	}

	return env
}

func unwrapReturnValue(obj object.Object) object.Object {
	if returnValue, ok := obj.(*object.ReturnValue); ok {
		return returnValue.Value
	}
	return obj
}

func isTruthy(obj object.Object) bool {
	switch obj {
	case object.NULL:
		return false
	case object.TRUE:
		return true
	case object.FALSE:
		return false
	default:
		return true
	}
}

func objectsEqual(left, right object.Object) bool {
	if left.Type() != right.Type() {
		return false
	}

	switch left := left.(type) {
	case *object.Number:
		return left.Value == right.(*object.Number).Value
	case *object.String:
		return left.Value == right.(*object.String).Value
	case *object.Boolean:
		return left.Value == right.(*object.Boolean).Value
	case *object.Null:
		return true
	default:
		return left == right
	}
}

func newError(format string, a ...interface{}) *object.Error {
	return &object.Error{Message: fmt.Sprintf(format, a...)}
}

func isError(obj object.Object) bool {
	if obj != nil {
		return obj.Type() == object.ERROR_OBJ
	}
	return false
}

func isException(obj object.Object) bool {
	if obj != nil {
		return obj.Type() == object.EXCEPTION_OBJ
	}
	return false
}

func evalImportStatement(is *ast.ImportStatement, env *object.Environment) object.Object {
	modulePath := is.Path.Value

	// Resolve relative path
	fullPath := filepath.Join(currentDir, modulePath)

	// Check module cache
	moduleMutex.RLock()
	if mod, ok := moduleCache[fullPath]; ok {
		moduleMutex.RUnlock()
		// Import exports into environment
		importModuleExports(mod, is.Alias, env)
		return mod
	}
	moduleMutex.RUnlock()

	// Read module file
	content, err := os.ReadFile(fullPath)
	if err != nil {
		return newError("cannot import module '%s': %s", modulePath, err.Error())
	}

	// Create module environment
	moduleEnv := object.NewEnvironment()

	// Parse module
	l := lexer.New(string(content))
	p := parser.New(l)
	program := p.ParseProgram()

	if len(p.Errors()) != 0 {
		return newError("parse error in module '%s': %s", modulePath, p.Errors()[0])
	}

	// Save current directory and set module directory
	oldDir := currentDir
	currentDir = filepath.Dir(fullPath)

	// Evaluate module
	result := Eval(program, moduleEnv)

	// Restore directory
	currentDir = oldDir

	if isError(result) {
		return result
	}

	// Create module object
	mod := &object.Module{
		Name:    modulePath,
		Exports: make(map[string]object.Object),
	}

	// Get exports from module environment
	if exports, ok := moduleEnv.Get("__exports__"); ok {
		if exportsMap, ok := exports.(*object.Map); ok {
			for k, v := range exportsMap.Pairs {
				mod.Exports[k] = v
			}
		}
	}

	// If no explicit exports, export all top-level functions and classes
	if len(mod.Exports) == 0 {
		for _, stmt := range program.Statements {
			switch s := stmt.(type) {
			case *ast.ExpressionStatement:
				if fn, ok := s.Expression.(*ast.FunctionLiteral); ok && fn.Name != nil {
					if val, ok := moduleEnv.Get(fn.Name.Value); ok {
						mod.Exports[fn.Name.Value] = val
					}
				}
			case *ast.ClassDeclaration:
				if val, ok := moduleEnv.Get(s.Name.Value); ok {
					mod.Exports[s.Name.Value] = val
				}
			case *ast.VariableDeclaration:
				if val, ok := moduleEnv.Get(s.Name.Value); ok {
					mod.Exports[s.Name.Value] = val
				}
			case *ast.ExportStatement:
				// Already handled by evalExportStatement
			}
		}
	}

	// Cache module
	moduleMutex.Lock()
	moduleCache[fullPath] = mod
	moduleMutex.Unlock()

	// Import exports into environment
	importModuleExports(mod, is.Alias, env)

	return mod
}

func importModuleExports(mod *object.Module, alias *ast.Identifier, env *object.Environment) {
	if alias != nil {
		// Import as namespace: anayat "math.bang" as math;
		// Access via: math.add(1, 2)
		modMap := &object.Map{Pairs: make(map[string]object.Object)}
		for k, v := range mod.Exports {
			modMap.Pairs[k] = v
		}
		env.Set(alias.Value, modMap)
	} else {
		// Import directly into namespace
		for k, v := range mod.Exports {
			env.Set(k, v)
		}
	}
}

func evalExportStatement(es *ast.ExportStatement, env *object.Environment) object.Object {
	// Evaluate the statement being exported
	result := Eval(es.Statement, env)
	if isError(result) {
		return result
	}

	// Get or create exports map
	var exportsMap *object.Map
	if exports, ok := env.Get("__exports__"); ok {
		exportsMap = exports.(*object.Map)
	} else {
		exportsMap = &object.Map{Pairs: make(map[string]object.Object)}
		env.Set("__exports__", exportsMap)
	}

	// Add to exports based on statement type
	switch stmt := es.Statement.(type) {
	case *ast.VariableDeclaration:
		exportsMap.Pairs[stmt.Name.Value] = result
	case *ast.ExpressionStatement:
		if fn, ok := stmt.Expression.(*ast.FunctionLiteral); ok && fn.Name != nil {
			if val, ok := env.Get(fn.Name.Value); ok {
				exportsMap.Pairs[fn.Name.Value] = val
			}
		}
	case *ast.ClassDeclaration:
		exportsMap.Pairs[stmt.Name.Value] = result
	}

	return result
}

func evalTryCatchStatement(tcs *ast.TryCatchStatement, env *object.Environment) object.Object {
	// Execute try block
	result := Eval(tcs.TryBlock, env)

	// Check for exception
	if isException(result) || isError(result) {
		if tcs.CatchBlock != nil {
			// Create catch scope with exception parameter
			catchEnv := object.NewEnclosedEnvironment(env)
			if tcs.CatchParam != nil {
				if exc, ok := result.(*object.Exception); ok {
					catchEnv.Set(tcs.CatchParam.Value, exc.Value)
				} else if err, ok := result.(*object.Error); ok {
					catchEnv.Set(tcs.CatchParam.Value, &object.String{Value: err.Message})
				}
			}

			// Execute catch block
			result = Eval(tcs.CatchBlock, catchEnv)
		}
	}

	// Execute finally block if present
	if tcs.FinallyBlock != nil {
		finallyResult := Eval(tcs.FinallyBlock, env)
		// Finally block errors override previous results
		if isError(finallyResult) || isException(finallyResult) {
			return finallyResult
		}
	}

	// Don't propagate caught exceptions
	if isException(result) && tcs.CatchBlock != nil {
		return object.NULL
	}

	return result
}

func evalThrowStatement(ts *ast.ThrowStatement, env *object.Environment) object.Object {
	val := Eval(ts.Value, env)
	if isError(val) {
		return val
	}

	return &object.Exception{
		Value: val,
	}
}
