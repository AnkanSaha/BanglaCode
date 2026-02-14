package ast

import (
	"BanglaCode/src/lexer"
	"bytes"
	"strings"
)

// ==================== Expression Nodes ====================

// Identifier represents a variable name
type Identifier struct {
	Token lexer.Token
	Value string
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
func (i *Identifier) String() string       { return i.Value }

// BinaryExpression represents operations like a + b, a == b
type BinaryExpression struct {
	Token    lexer.Token // the operator token
	Left     Expression
	Operator string
	Right    Expression
}

func (be *BinaryExpression) expressionNode()      {}
func (be *BinaryExpression) TokenLiteral() string { return be.Token.Literal }
func (be *BinaryExpression) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(be.Left.String())
	out.WriteString(" " + be.Operator + " ")
	out.WriteString(be.Right.String())
	out.WriteString(")")
	return out.String()
}

// UnaryExpression represents !x, -x, na x
type UnaryExpression struct {
	Token    lexer.Token // the operator token
	Operator string
	Right    Expression
}

func (ue *UnaryExpression) expressionNode()      {}
func (ue *UnaryExpression) TokenLiteral() string { return ue.Token.Literal }
func (ue *UnaryExpression) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(ue.Operator)
	out.WriteString(ue.Right.String())
	out.WriteString(")")
	return out.String()
}

// AssignmentExpression represents x = 5, x += 5
type AssignmentExpression struct {
	Token    lexer.Token // the = token
	Name     Expression
	Operator string // =, +=, -=, *=, /=
	Value    Expression
}

func (ae *AssignmentExpression) expressionNode()      {}
func (ae *AssignmentExpression) TokenLiteral() string { return ae.Token.Literal }
func (ae *AssignmentExpression) String() string {
	var out bytes.Buffer
	out.WriteString(ae.Name.String())
	out.WriteString(" " + ae.Operator + " ")
	out.WriteString(ae.Value.String())
	return out.String()
}

// CallExpression represents function calls: add(1, 2)
type CallExpression struct {
	Token     lexer.Token // the '(' token
	Function  Expression  // Identifier or FunctionLiteral
	Arguments []Expression
}

func (ce *CallExpression) expressionNode()      {}
func (ce *CallExpression) TokenLiteral() string { return ce.Token.Literal }
func (ce *CallExpression) String() string {
	var out bytes.Buffer
	args := []string{}
	for _, a := range ce.Arguments {
		args = append(args, a.String())
	}
	out.WriteString(ce.Function.String())
	out.WriteString("(")
	out.WriteString(strings.Join(args, ", "))
	out.WriteString(")")
	return out.String()
}

// MemberExpression represents property access: obj.prop, arr[0]
type MemberExpression struct {
	Token    lexer.Token // the '.' or '[' token
	Object   Expression
	Property Expression
	Computed bool // true for arr[0], false for obj.prop
}

func (me *MemberExpression) expressionNode()      {}
func (me *MemberExpression) TokenLiteral() string { return me.Token.Literal }
func (me *MemberExpression) String() string {
	var out bytes.Buffer
	out.WriteString(me.Object.String())
	if me.Computed {
		out.WriteString("[")
		out.WriteString(me.Property.String())
		out.WriteString("]")
	} else {
		out.WriteString(".")
		out.WriteString(me.Property.String())
	}
	return out.String()
}

// NewExpression represents: notun Manush(args)
type NewExpression struct {
	Token     lexer.Token // the NOTUN token
	Class     Expression
	Arguments []Expression
}

func (ne *NewExpression) expressionNode()      {}
func (ne *NewExpression) TokenLiteral() string { return ne.Token.Literal }
func (ne *NewExpression) String() string {
	var out bytes.Buffer
	args := []string{}
	for _, a := range ne.Arguments {
		args = append(args, a.String())
	}
	out.WriteString("notun ")
	out.WriteString(ne.Class.String())
	out.WriteString("(")
	out.WriteString(strings.Join(args, ", "))
	out.WriteString(")")
	return out.String()
}

// SpreadElement represents ...expression (spread operator)
type SpreadElement struct {
	Token    lexer.Token // the '...' token
	Argument Expression
}

func (se *SpreadElement) expressionNode()      {}
func (se *SpreadElement) TokenLiteral() string { return se.Token.Literal }
func (se *SpreadElement) String() string {
	var out bytes.Buffer
	out.WriteString("...")
	out.WriteString(se.Argument.String())
	return out.String()
}

// AwaitExpression represents: opekha promise
type AwaitExpression struct {
	Token      lexer.Token // the OPEKHA token
	Expression Expression
}

func (ae *AwaitExpression) expressionNode()      {}
func (ae *AwaitExpression) TokenLiteral() string { return ae.Token.Literal }
func (ae *AwaitExpression) String() string {
	var out bytes.Buffer
	out.WriteString("opekha ")
	out.WriteString(ae.Expression.String())
	return out.String()
}
