package ast

import (
	"BanglaCode/src/lexer"
	"bytes"
	"strings"
)

// ==================== Literal Expressions ====================

// NumberLiteral represents a numeric value
type NumberLiteral struct {
	Token lexer.Token
	Value float64
}

func (nl *NumberLiteral) expressionNode()      {}
func (nl *NumberLiteral) TokenLiteral() string { return nl.Token.Literal }
func (nl *NumberLiteral) String() string       { return nl.Token.Literal }

// StringLiteral represents a string value
type StringLiteral struct {
	Token lexer.Token
	Value string
}

func (sl *StringLiteral) expressionNode()      {}
func (sl *StringLiteral) TokenLiteral() string { return sl.Token.Literal }
func (sl *StringLiteral) String() string       { return sl.Token.Literal }

// TemplateLiteral represents a template literal with interpolation: `hello ${name}`
type TemplateLiteral struct {
	Token lexer.Token // the backtick token
	Value string      // the raw template string including ${...} parts
}

func (tl *TemplateLiteral) expressionNode()      {}
func (tl *TemplateLiteral) TokenLiteral() string { return tl.Token.Literal }
func (tl *TemplateLiteral) String() string       { return "`" + tl.Value + "`" }

// BooleanLiteral represents sotti or mittha
type BooleanLiteral struct {
	Token lexer.Token
	Value bool
}

func (bl *BooleanLiteral) expressionNode()      {}
func (bl *BooleanLiteral) TokenLiteral() string { return bl.Token.Literal }
func (bl *BooleanLiteral) String() string       { return bl.Token.Literal }

// NullLiteral represents khali
type NullLiteral struct {
	Token lexer.Token
}

func (nl *NullLiteral) expressionNode()      {}
func (nl *NullLiteral) TokenLiteral() string { return nl.Token.Literal }
func (nl *NullLiteral) String() string       { return "khali" }

// ArrayLiteral represents [1, 2, 3]
type ArrayLiteral struct {
	Token    lexer.Token // the '[' token
	Elements []Expression
}

func (al *ArrayLiteral) expressionNode()      {}
func (al *ArrayLiteral) TokenLiteral() string { return al.Token.Literal }
func (al *ArrayLiteral) String() string {
	var out bytes.Buffer
	elements := []string{}
	for _, el := range al.Elements {
		elements = append(elements, el.String())
	}
	out.WriteString("[")
	out.WriteString(strings.Join(elements, ", "))
	out.WriteString("]")
	return out.String()
}

// MapLiteral represents {"key": "value"}
type MapLiteral struct {
	Token lexer.Token // the '{' token
	Pairs map[Expression]Expression
}

func (ml *MapLiteral) expressionNode()      {}
func (ml *MapLiteral) TokenLiteral() string { return ml.Token.Literal }
func (ml *MapLiteral) String() string {
	var out bytes.Buffer
	pairs := []string{}
	for key, value := range ml.Pairs {
		pairs = append(pairs, key.String()+":"+value.String())
	}
	out.WriteString("{")
	out.WriteString(strings.Join(pairs, ", "))
	out.WriteString("}")
	return out.String()
}

// FunctionLiteral represents: kaj(a, b) { ... }
type FunctionLiteral struct {
	Token         lexer.Token // the KAJ token
	Name          *Identifier // optional function name
	Parameters    []*Identifier
	RestParameter *Identifier // optional rest parameter (...args)
	Body          *BlockStatement
}

func (fl *FunctionLiteral) expressionNode()      {}
func (fl *FunctionLiteral) TokenLiteral() string { return fl.Token.Literal }
func (fl *FunctionLiteral) String() string {
	var out bytes.Buffer
	params := []string{}
	for _, p := range fl.Parameters {
		params = append(params, p.String())
	}
	if fl.RestParameter != nil {
		params = append(params, "..."+fl.RestParameter.String())
	}
	// Constructor shuru() is output without kaj prefix
	if fl.Name != nil && fl.Name.Value == "shuru" {
		out.WriteString("shuru(")
	} else {
		out.WriteString("kaj")
		if fl.Name != nil {
			out.WriteString(" " + fl.Name.String())
		}
		out.WriteString("(")
	}
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") ")
	out.WriteString(fl.Body.String())
	return out.String()
}

// AsyncFunctionLiteral represents: proyash kaj(a, b) { ... }
type AsyncFunctionLiteral struct {
	Token         lexer.Token // the PROYASH token
	Name          *Identifier // optional function name
	Parameters    []*Identifier
	RestParameter *Identifier // optional rest parameter (...args)
	Body          *BlockStatement
}

func (afl *AsyncFunctionLiteral) expressionNode()      {}
func (afl *AsyncFunctionLiteral) TokenLiteral() string { return afl.Token.Literal }
func (afl *AsyncFunctionLiteral) String() string {
	var out bytes.Buffer
	params := []string{}
	for _, p := range afl.Parameters {
		params = append(params, p.String())
	}
	if afl.RestParameter != nil {
		params = append(params, "..."+afl.RestParameter.String())
	}
	out.WriteString("proyash kaj")
	if afl.Name != nil {
		out.WriteString(" " + afl.Name.String())
	}
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") ")
	out.WriteString(afl.Body.String())
	return out.String()
}
