package ast

import (
	"BanglaCode/src/lexer"
	"bytes"
)

// ==================== Statement Nodes ====================

// VariableDeclaration represents: dhoro x = 5; or sthir x = 5; or bishwo x = 5;
type VariableDeclaration struct {
	Token      lexer.Token // the DHORO/STHIR/BISHWO token
	Name       *Identifier
	Value      Expression
	IsConstant bool // true for sthir declarations
	IsGlobal   bool // true for bishwo declarations
}

func (vd *VariableDeclaration) statementNode()       {}
func (vd *VariableDeclaration) TokenLiteral() string { return vd.Token.Literal }
func (vd *VariableDeclaration) String() string {
	var out bytes.Buffer
	if vd.IsConstant {
		out.WriteString("sthir ")
	} else if vd.IsGlobal {
		out.WriteString("bishwo ")
	} else {
		out.WriteString("dhoro ")
	}
	out.WriteString(vd.Name.String())
	out.WriteString(" = ")
	if vd.Value != nil {
		out.WriteString(vd.Value.String())
	}
	out.WriteString(";")
	return out.String()
}

// ExpressionStatement wraps an expression as a statement
type ExpressionStatement struct {
	Token      lexer.Token
	Expression Expression
}

func (es *ExpressionStatement) statementNode()       {}
func (es *ExpressionStatement) TokenLiteral() string { return es.Token.Literal }
func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}
	return ""
}

// BlockStatement represents a code block: { ... }
type BlockStatement struct {
	Token      lexer.Token // the '{' token
	Statements []Statement
}

func (bs *BlockStatement) statementNode()       {}
func (bs *BlockStatement) TokenLiteral() string { return bs.Token.Literal }
func (bs *BlockStatement) String() string {
	var out bytes.Buffer
	for _, s := range bs.Statements {
		out.WriteString(s.String())
	}
	return out.String()
}

// IfStatement represents: jodi (condition) { ... } nahole { ... }
type IfStatement struct {
	Token       lexer.Token // the JODI token
	Condition   Expression
	Consequence *BlockStatement
	Alternative *BlockStatement
}

func (is *IfStatement) statementNode()       {}
func (is *IfStatement) TokenLiteral() string { return is.Token.Literal }
func (is *IfStatement) String() string {
	var out bytes.Buffer
	out.WriteString("jodi ")
	out.WriteString(is.Condition.String())
	out.WriteString(" ")
	out.WriteString(is.Consequence.String())
	if is.Alternative != nil {
		out.WriteString(" nahole ")
		out.WriteString(is.Alternative.String())
	}
	return out.String()
}

// WhileStatement represents: jotokkhon (condition) { ... }
type WhileStatement struct {
	Token     lexer.Token // the JOTOKKHON token
	Condition Expression
	Body      *BlockStatement
}

func (ws *WhileStatement) statementNode()       {}
func (ws *WhileStatement) TokenLiteral() string { return ws.Token.Literal }
func (ws *WhileStatement) String() string {
	var out bytes.Buffer
	out.WriteString("jotokkhon ")
	out.WriteString(ws.Condition.String())
	out.WriteString(" ")
	out.WriteString(ws.Body.String())
	return out.String()
}

// ForStatement represents: ghuriye (init; condition; update) { ... }
type ForStatement struct {
	Token     lexer.Token // the GHURIYE token
	Init      Statement
	Condition Expression
	Update    Expression
	Body      *BlockStatement
}

func (fs *ForStatement) statementNode()       {}
func (fs *ForStatement) TokenLiteral() string { return fs.Token.Literal }
func (fs *ForStatement) String() string {
	var out bytes.Buffer
	out.WriteString("ghuriye (")
	if fs.Init != nil {
		out.WriteString(fs.Init.String())
	}
	out.WriteString("; ")
	if fs.Condition != nil {
		out.WriteString(fs.Condition.String())
	}
	out.WriteString("; ")
	if fs.Update != nil {
		out.WriteString(fs.Update.String())
	}
	out.WriteString(") ")
	out.WriteString(fs.Body.String())
	return out.String()
}

// ReturnStatement represents: ferao x;
type ReturnStatement struct {
	Token       lexer.Token // the FERAO token
	ReturnValue Expression
}

func (rs *ReturnStatement) statementNode()       {}
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }
func (rs *ReturnStatement) String() string {
	var out bytes.Buffer
	out.WriteString("ferao ")
	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}
	out.WriteString(";")
	return out.String()
}

// ClassDeclaration represents: sreni Manush { ... }
type ClassDeclaration struct {
	Token   lexer.Token // the SRENI token
	Name    *Identifier
	Methods []*FunctionLiteral
}

func (cd *ClassDeclaration) statementNode()       {}
func (cd *ClassDeclaration) TokenLiteral() string { return cd.Token.Literal }
func (cd *ClassDeclaration) String() string {
	var out bytes.Buffer
	out.WriteString("sreni ")
	out.WriteString(cd.Name.String())
	out.WriteString(" { ")
	for _, method := range cd.Methods {
		out.WriteString(method.String())
	}
	out.WriteString(" }")
	return out.String()
}

// BreakStatement represents: thamo;
type BreakStatement struct {
	Token lexer.Token // the THAMO token
}

func (bs *BreakStatement) statementNode()       {}
func (bs *BreakStatement) TokenLiteral() string { return bs.Token.Literal }
func (bs *BreakStatement) String() string       { return "thamo;" }

// ContinueStatement represents: chharo;
type ContinueStatement struct {
	Token lexer.Token // the CHHARO token
}

func (cs *ContinueStatement) statementNode()       {}
func (cs *ContinueStatement) TokenLiteral() string { return cs.Token.Literal }
func (cs *ContinueStatement) String() string       { return "chharo;" }

// ImportStatement represents: ano "module.bang" hisabe alias;
type ImportStatement struct {
	Token lexer.Token // the ANO token
	Path  *StringLiteral
	Alias *Identifier // optional: ano "module.bang" hisabe mod;
}

func (is *ImportStatement) statementNode()       {}
func (is *ImportStatement) TokenLiteral() string { return is.Token.Literal }
func (is *ImportStatement) String() string {
	var out bytes.Buffer
	out.WriteString("ano ")
	out.WriteString("\"" + is.Path.Value + "\"")
	if is.Alias != nil {
		out.WriteString(" hisabe " + is.Alias.Value)
	}
	out.WriteString(";")
	return out.String()
}

// ExportStatement represents: pathao kaj funcName() { }
type ExportStatement struct {
	Token     lexer.Token // the PATHAO token
	Statement Statement   // the exported statement (function, class, or variable)
}

func (es *ExportStatement) statementNode()       {}
func (es *ExportStatement) TokenLiteral() string { return es.Token.Literal }
func (es *ExportStatement) String() string {
	var out bytes.Buffer
	out.WriteString("pathao ")
	out.WriteString(es.Statement.String())
	return out.String()
}

// TryCatchStatement represents: chesta { } dhoro_bhul (e) { } shesh { }
type TryCatchStatement struct {
	Token        lexer.Token // the CHESTA token
	TryBlock     *BlockStatement
	CatchParam   *Identifier // parameter in catch block
	CatchBlock   *BlockStatement
	FinallyBlock *BlockStatement // optional
}

func (tcs *TryCatchStatement) statementNode()       {}
func (tcs *TryCatchStatement) TokenLiteral() string { return tcs.Token.Literal }
func (tcs *TryCatchStatement) String() string {
	var out bytes.Buffer
	out.WriteString("chesta ")
	out.WriteString(tcs.TryBlock.String())
	if tcs.CatchBlock != nil {
		out.WriteString(" dhoro_bhul (")
		if tcs.CatchParam != nil {
			out.WriteString(tcs.CatchParam.Value)
		}
		out.WriteString(") ")
		out.WriteString(tcs.CatchBlock.String())
	}
	if tcs.FinallyBlock != nil {
		out.WriteString(" shesh ")
		out.WriteString(tcs.FinallyBlock.String())
	}
	return out.String()
}

// ThrowStatement represents: felo "error message";
type ThrowStatement struct {
	Token lexer.Token // the FELO token
	Value Expression
}

func (ts *ThrowStatement) statementNode()       {}
func (ts *ThrowStatement) TokenLiteral() string { return ts.Token.Literal }
func (ts *ThrowStatement) String() string {
	var out bytes.Buffer
	out.WriteString("felo ")
	if ts.Value != nil {
		out.WriteString(ts.Value.String())
	}
	out.WriteString(";")
	return out.String()
}
