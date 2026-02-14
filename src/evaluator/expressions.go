package evaluator

import (
	"BanglaCode/src/ast"
	"BanglaCode/src/object"
	"math"
)

// evalUnaryExpression evaluates unary expressions (!, -, na)
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

// evalBangOperator evaluates the ! (not) operator
func evalBangOperator(right object.Object) object.Object {
	switch right {
	case object.TRUE:
		return object.FALSE
	case object.FALSE:
		return object.TRUE
	case object.NULL:
		return object.TRUE
	default:
		return object.FALSE
	}
}

// evalMinusOperator evaluates the - (negative) operator
func evalMinusOperator(right object.Object) object.Object {
	if right.Type() != object.NUMBER_OBJ {
		return newError("unknown operator: -%s", right.Type())
	}
	value := right.(*object.Number).Value
	return &object.Number{Value: -value}
}

// evalBinaryExpression evaluates binary expressions (+, -, *, /, etc.)
func evalBinaryExpression(operator string, left, right object.Object) object.Object {
	switch {
	case left.Type() == object.NUMBER_OBJ && right.Type() == object.NUMBER_OBJ:
		return evalNumberBinaryExpression(operator, left, right)
	case left.Type() == object.STRING_OBJ && right.Type() == object.STRING_OBJ:
		return evalStringBinaryExpression(operator, left, right)
	case left.Type() == object.STRING_OBJ && right.Type() == object.NUMBER_OBJ:
		return evalStringNumberBinaryExpression(operator, left, right)
	case operator == "==" || operator == "soman":
		return boolToObject(left == right)
	case operator == "!=" || operator == "osoman":
		return boolToObject(left != right)
	case operator == "ebong":
		return boolToObject(isTruthy(left) && isTruthy(right))
	case operator == "ba":
		return boolToObject(isTruthy(left) || isTruthy(right))
	default:
		return newError("unknown operator: %s %s %s", left.Type(), operator, right.Type())
	}
}

// boolToObject converts a Go bool to a BanglaCode Boolean object
func boolToObject(value bool) *object.Boolean {
	if value {
		return object.TRUE
	}
	return object.FALSE
}

// evalNumberBinaryExpression evaluates binary expressions on numbers
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
		return &object.Number{Value: float64(int64(leftVal) % int64(rightVal))}
	case "**":
		return &object.Number{Value: math.Pow(leftVal, rightVal)}
	case "<":
		return boolToObject(leftVal < rightVal)
	case ">":
		return boolToObject(leftVal > rightVal)
	case "<=":
		return boolToObject(leftVal <= rightVal)
	case ">=":
		return boolToObject(leftVal >= rightVal)
	case "==", "soman":
		return boolToObject(leftVal == rightVal)
	case "!=", "osoman":
		return boolToObject(leftVal != rightVal)
	default:
		return newError("unknown operator: %s %s %s", left.Type(), operator, right.Type())
	}
}

// evalStringBinaryExpression evaluates binary expressions on strings
func evalStringBinaryExpression(operator string, left, right object.Object) object.Object {
	leftVal := left.(*object.String).Value
	rightVal := right.(*object.String).Value

	switch operator {
	case "+":
		return &object.String{Value: leftVal + rightVal}
	case "==", "soman":
		return boolToObject(leftVal == rightVal)
	case "!=", "osoman":
		return boolToObject(leftVal != rightVal)
	default:
		return newError("unknown operator: %s %s %s", left.Type(), operator, right.Type())
	}
}

// evalStringNumberBinaryExpression handles string + number concatenation
func evalStringNumberBinaryExpression(operator string, left, right object.Object) object.Object {
	if operator == "+" {
		leftVal := left.(*object.String).Value
		rightVal := right.(*object.Number).Inspect()
		return &object.String{Value: leftVal + rightVal}
	}
	return newError("unknown operator: %s %s %s", left.Type(), operator, right.Type())
}

// evalAssignmentExpression evaluates assignment expressions (=, +=, -=, etc.)
func evalAssignmentExpression(ae *ast.AssignmentExpression, env *object.Environment) object.Object {
	// Handle member assignment (obj.prop = value or arr[idx] = value)
	if member, ok := ae.Name.(*ast.MemberExpression); ok {
		return evalMemberAssignment(member, ae.Operator, ae.Value, env)
	}

	// Handle simple variable assignment
	ident, ok := ae.Name.(*ast.Identifier)
	if !ok {
		return newError("invalid assignment target")
	}

	// Check if trying to reassign a constant
	if env.IsConstant(ident.Value) {
		return newErrorAt(ae.Token.Line, ae.Token.Column, "'%s' ekti sthir (constant), eitake bodlano jabe na", ident.Value)
	}

	value := Eval(ae.Value, env)
	if isError(value) {
		return value
	}

	// Handle compound assignment operators
	switch ae.Operator {
	case "=":
		env.Update(ident.Value, value)
		return value
	case "+=", "-=", "*=", "/=":
		current, ok := env.Get(ident.Value)
		if !ok {
			return newErrorAt(ae.Token.Line, ae.Token.Column, "variable '%s' is not defined", ident.Value)
		}

		// Calculate new value based on operator
		var result object.Object
		op := string(ae.Operator[0]) // Get first char: +, -, *, /
		result = evalBinaryExpression(op, current, value)

		if isError(result) {
			return result
		}

		env.Update(ident.Value, result)
		return result
	default:
		return newError("unknown assignment operator: %s", ae.Operator)
	}
}

// evalMemberAssignment handles assignment to object properties or array elements
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
			return newError("array index must be a number, got %s", index.Type())
		}
		idx := int(index.(*object.Number).Value)
		if idx < 0 || idx >= len(o.Elements) {
			return newError("array index out of bounds: %d", idx)
		}

		// Handle compound operators
		if operator != "=" {
			current := o.Elements[idx]
			op := string(operator[0])
			val = evalBinaryExpression(op, current, val)
			if isError(val) {
				return val
			}
		}

		o.Elements[idx] = val
		return val

	case *object.Map:
		var key string
		if member.Computed {
			keyObj := Eval(member.Property, env)
			if isError(keyObj) {
				return keyObj
			}
			key = getMapKey(keyObj)
		} else {
			if ident, ok := member.Property.(*ast.Identifier); ok {
				key = ident.Value
			} else {
				return newError("invalid map key")
			}
		}

		// Handle compound operators
		if operator != "=" {
			current, ok := o.Pairs[key]
			if !ok {
				return newError("key '%s' not found in map", key)
			}
			op := string(operator[0])
			val = evalBinaryExpression(op, current, val)
			if isError(val) {
				return val
			}
		}

		o.Pairs[key] = val
		return val

	case *object.Instance:
		if ident, ok := member.Property.(*ast.Identifier); ok {
			// Handle compound operators
			if operator != "=" {
				current, ok := o.Properties[ident.Value]
				if !ok {
					return newError("property '%s' not found", ident.Value)
				}
				op := string(operator[0])
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
		return newError("cannot assign to %s", obj.Type())
	}
}

// evalMemberExpression evaluates member access (obj.prop or arr[idx])
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

// evalArrayIndex evaluates array indexing
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

// evalMapLiteral evaluates map/object literals
func evalMapLiteral(node *ast.MapLiteral, env *object.Environment) object.Object {
	pairs := make(map[string]object.Object)

	for keyNode, valueNode := range node.Pairs {
		var keyStr string

		// Handle identifier keys as string keys (JS-like object syntax)
		if ident, ok := keyNode.(*ast.Identifier); ok {
			keyStr = ident.Value
		} else {
			key := Eval(keyNode, env)
			if isError(key) {
				return key
			}

			switch k := key.(type) {
			case *object.String:
				keyStr = k.Value
			case *object.Number:
				keyStr = k.Inspect()
			default:
				return newError("unusable as map key: %s", key.Type())
			}
		}

		value := Eval(valueNode, env)
		if isError(value) {
			return value
		}

		pairs[keyStr] = value
	}

	return &object.Map{Pairs: pairs}
}
