package test

import (
	"BanglaCode/src/ast"
	"BanglaCode/src/lexer"
	"BanglaCode/src/parser"
	"testing"
)

func TestParseVariableDeclaration(t *testing.T) {
	tests := []struct {
		input              string
		expectedIdentifier string
		expectedValue      interface{}
	}{
		{"dhoro x = 5;", "x", 5.0},
		{"dhoro y = 10.5;", "y", 10.5},
		{"dhoro foo = sotti;", "foo", true},
		{"dhoro bar = mittha;", "bar", false},
	}

	for _, tt := range tests {
		l := lexer.New(tt.input)
		p := parser.New(l)
		program := p.ParseProgram()
		checkParserErrors(t, p)

		if len(program.Statements) != 1 {
			t.Fatalf("program.Statements does not contain 1 statement. got=%d",
				len(program.Statements))
		}

		stmt, ok := program.Statements[0].(*ast.VariableDeclaration)
		if !ok {
			t.Fatalf("program.Statements[0] is not *ast.VariableDeclaration. got=%T",
				program.Statements[0])
		}

		if stmt.Name.Value != tt.expectedIdentifier {
			t.Errorf("stmt.Name.Value not '%s'. got=%s",
				tt.expectedIdentifier, stmt.Name.Value)
		}

		if !testLiteralExpression(t, stmt.Value, tt.expectedValue) {
			return
		}
	}
}

func TestReturnStatement(t *testing.T) {
	tests := []struct {
		input         string
		expectedValue interface{}
	}{
		{"ferao 5;", 5.0},
		{"ferao sotti;", true},
		{"ferao x;", "x"},
	}

	for _, tt := range tests {
		l := lexer.New(tt.input)
		p := parser.New(l)
		program := p.ParseProgram()
		checkParserErrors(t, p)

		if len(program.Statements) != 1 {
			t.Fatalf("program.Statements does not contain 1 statement. got=%d",
				len(program.Statements))
		}

		stmt, ok := program.Statements[0].(*ast.ReturnStatement)
		if !ok {
			t.Fatalf("program.Statements[0] is not *ast.ReturnStatement. got=%T",
				program.Statements[0])
		}

		if stmt.TokenLiteral() != "ferao" {
			t.Errorf("stmt.TokenLiteral not 'ferao'. got=%s", stmt.TokenLiteral())
		}
	}
}

func TestIdentifierExpression(t *testing.T) {
	input := "foobar;"

	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("program has not enough statements. got=%d",
			len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not *ast.ExpressionStatement. got=%T",
			program.Statements[0])
	}

	ident, ok := stmt.Expression.(*ast.Identifier)
	if !ok {
		t.Fatalf("exp not *ast.Identifier. got=%T", stmt.Expression)
	}

	if ident.Value != "foobar" {
		t.Errorf("ident.Value not %s. got=%s", "foobar", ident.Value)
	}

	if ident.TokenLiteral() != "foobar" {
		t.Errorf("ident.TokenLiteral not %s. got=%s", "foobar", ident.TokenLiteral())
	}
}

func TestNumberLiteralExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected float64
	}{
		{"5;", 5},
		{"10.5;", 10.5},
		{"3.14159;", 3.14159},
	}

	for _, tt := range tests {
		l := lexer.New(tt.input)
		p := parser.New(l)
		program := p.ParseProgram()
		checkParserErrors(t, p)

		if len(program.Statements) != 1 {
			t.Fatalf("program has not enough statements. got=%d",
				len(program.Statements))
		}

		stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
		if !ok {
			t.Fatalf("program.Statements[0] is not *ast.ExpressionStatement. got=%T",
				program.Statements[0])
		}

		num, ok := stmt.Expression.(*ast.NumberLiteral)
		if !ok {
			t.Fatalf("exp not *ast.NumberLiteral. got=%T", stmt.Expression)
		}

		if num.Value != tt.expected {
			t.Errorf("num.Value not %f. got=%f", tt.expected, num.Value)
		}
	}
}

func TestStringLiteralExpression(t *testing.T) {
	input := `"hello world";`

	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	stmt := program.Statements[0].(*ast.ExpressionStatement)
	str, ok := stmt.Expression.(*ast.StringLiteral)
	if !ok {
		t.Fatalf("exp not *ast.StringLiteral. got=%T", stmt.Expression)
	}

	if str.Value != "hello world" {
		t.Errorf("str.Value not %s. got=%s", "hello world", str.Value)
	}
}

func TestBooleanExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"sotti;", true},
		{"mittha;", false},
	}

	for _, tt := range tests {
		l := lexer.New(tt.input)
		p := parser.New(l)
		program := p.ParseProgram()
		checkParserErrors(t, p)

		if len(program.Statements) != 1 {
			t.Fatalf("program has not enough statements. got=%d",
				len(program.Statements))
		}

		stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
		if !ok {
			t.Fatalf("program.Statements[0] is not *ast.ExpressionStatement. got=%T",
				program.Statements[0])
		}

		boolean, ok := stmt.Expression.(*ast.BooleanLiteral)
		if !ok {
			t.Fatalf("exp not *ast.BooleanLiteral. got=%T", stmt.Expression)
		}

		if boolean.Value != tt.expected {
			t.Errorf("boolean.Value not %t. got=%t", tt.expected, boolean.Value)
		}
	}
}

func TestNullExpression(t *testing.T) {
	input := "khali;"

	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not *ast.ExpressionStatement. got=%T",
			program.Statements[0])
	}

	_, ok = stmt.Expression.(*ast.NullLiteral)
	if !ok {
		t.Fatalf("exp not *ast.NullLiteral. got=%T", stmt.Expression)
	}
}

func TestUnaryExpression(t *testing.T) {
	tests := []struct {
		input    string
		operator string
		value    interface{}
	}{
		{"!sotti;", "!", true},
		{"-15;", "-", 15.0},
		{"na sotti;", "na", true},
	}

	for _, tt := range tests {
		l := lexer.New(tt.input)
		p := parser.New(l)
		program := p.ParseProgram()
		checkParserErrors(t, p)

		if len(program.Statements) != 1 {
			t.Fatalf("program.Statements does not contain %d statements. got=%d\n",
				1, len(program.Statements))
		}

		stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
		if !ok {
			t.Fatalf("program.Statements[0] is not *ast.ExpressionStatement. got=%T",
				program.Statements[0])
		}

		exp, ok := stmt.Expression.(*ast.UnaryExpression)
		if !ok {
			t.Fatalf("stmt is not *ast.UnaryExpression. got=%T", stmt.Expression)
		}

		if exp.Operator != tt.operator {
			t.Fatalf("exp.Operator is not '%s'. got=%s",
				tt.operator, exp.Operator)
		}

		if !testLiteralExpression(t, exp.Right, tt.value) {
			return
		}
	}
}

func TestBinaryExpressions(t *testing.T) {
	tests := []struct {
		input      string
		leftValue  interface{}
		operator   string
		rightValue interface{}
	}{
		{"5 + 5;", 5.0, "+", 5.0},
		{"5 - 5;", 5.0, "-", 5.0},
		{"5 * 5;", 5.0, "*", 5.0},
		{"5 / 5;", 5.0, "/", 5.0},
		{"5 % 5;", 5.0, "%", 5.0},
		{"5 > 5;", 5.0, ">", 5.0},
		{"5 < 5;", 5.0, "<", 5.0},
		{"5 == 5;", 5.0, "==", 5.0},
		{"5 != 5;", 5.0, "!=", 5.0},
		{"5 >= 5;", 5.0, ">=", 5.0},
		{"5 <= 5;", 5.0, "<=", 5.0},
		{"sotti ebong sotti;", true, "ebong", true},
		{"sotti ba mittha;", true, "ba", false},
	}

	for _, tt := range tests {
		l := lexer.New(tt.input)
		p := parser.New(l)
		program := p.ParseProgram()
		checkParserErrors(t, p)

		if len(program.Statements) != 1 {
			t.Fatalf("program.Statements does not contain %d statements. got=%d\n",
				1, len(program.Statements))
		}

		stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
		if !ok {
			t.Fatalf("program.Statements[0] is not *ast.ExpressionStatement. got=%T",
				program.Statements[0])
		}

		exp, ok := stmt.Expression.(*ast.BinaryExpression)
		if !ok {
			t.Fatalf("exp is not *ast.BinaryExpression. got=%T", stmt.Expression)
		}

		if !testLiteralExpression(t, exp.Left, tt.leftValue) {
			return
		}

		if exp.Operator != tt.operator {
			t.Fatalf("exp.Operator is not '%s'. got=%s",
				tt.operator, exp.Operator)
		}

		if !testLiteralExpression(t, exp.Right, tt.rightValue) {
			return
		}
	}
}

func TestOperatorPrecedence(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"1 + 2 * 3;", "(1 + (2 * 3))"},
		{"1 * 2 + 3;", "((1 * 2) + 3)"},
		{"1 + 2 + 3;", "((1 + 2) + 3)"},
		{"1 + 2 * 3 + 4 / 5 - 6;", "(((1 + (2 * 3)) + (4 / 5)) - 6)"},
		{"-a * b;", "((-a) * b)"},
		{"!sotti;", "(!sotti)"},
		{"a + b + c;", "((a + b) + c)"},
		{"a + b - c;", "((a + b) - c)"},
		{"a * b * c;", "((a * b) * c)"},
		{"a * b / c;", "((a * b) / c)"},
		{"(a + b) * c;", "((a + b) * c)"},
		{"1 + (2 + 3) + 4;", "((1 + (2 + 3)) + 4)"},
	}

	for _, tt := range tests {
		l := lexer.New(tt.input)
		p := parser.New(l)
		program := p.ParseProgram()
		checkParserErrors(t, p)

		actual := program.String()
		if actual != tt.expected {
			t.Errorf("expected=%q, got=%q", tt.expected, actual)
		}
	}
}

func TestIfStatement(t *testing.T) {
	input := `jodi (x < y) { x; }`

	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("program.Statements does not contain %d statements. got=%d\n",
			1, len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.IfStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not *ast.IfStatement. got=%T",
			program.Statements[0])
	}

	if !testBinaryExpression(t, stmt.Condition, "x", "<", "y") {
		return
	}

	if len(stmt.Consequence.Statements) != 1 {
		t.Errorf("consequence is not 1 statement. got=%d\n",
			len(stmt.Consequence.Statements))
	}

	consequence, ok := stmt.Consequence.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("Statements[0] is not *ast.ExpressionStatement. got=%T",
			stmt.Consequence.Statements[0])
	}

	if !testIdentifier(t, consequence.Expression, "x") {
		return
	}

	if stmt.Alternative != nil {
		t.Errorf("stmt.Alternative was not nil. got=%+v", stmt.Alternative)
	}
}

func TestIfElseStatement(t *testing.T) {
	input := `jodi (x < y) { x; } nahole { y; }`

	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("program.Statements does not contain %d statements. got=%d\n",
			1, len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.IfStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not *ast.IfStatement. got=%T",
			program.Statements[0])
	}

	if !testBinaryExpression(t, stmt.Condition, "x", "<", "y") {
		return
	}

	if len(stmt.Consequence.Statements) != 1 {
		t.Errorf("consequence is not 1 statements. got=%d\n",
			len(stmt.Consequence.Statements))
	}

	if stmt.Alternative == nil {
		t.Fatalf("stmt.Alternative was nil")
	}

	if len(stmt.Alternative.Statements) != 1 {
		t.Errorf("alternative is not 1 statements. got=%d\n",
			len(stmt.Alternative.Statements))
	}
}

func TestWhileStatement(t *testing.T) {
	input := `jotokkhon (x < 10) { x = x + 1; }`

	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("program.Statements does not contain %d statements. got=%d\n",
			1, len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.WhileStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not *ast.WhileStatement. got=%T",
			program.Statements[0])
	}

	if !testBinaryExpression(t, stmt.Condition, "x", "<", 10.0) {
		return
	}

	if len(stmt.Body.Statements) != 1 {
		t.Errorf("body is not 1 statement. got=%d\n",
			len(stmt.Body.Statements))
	}
}

func TestForStatement(t *testing.T) {
	input := `ghuriye (dhoro i = 0; i < 5; i = i + 1) { x; }`

	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("program.Statements does not contain %d statements. got=%d\n",
			1, len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.ForStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not *ast.ForStatement. got=%T",
			program.Statements[0])
	}

	// Check init
	init, ok := stmt.Init.(*ast.VariableDeclaration)
	if !ok {
		t.Fatalf("init is not *ast.VariableDeclaration. got=%T", stmt.Init)
	}
	if init.Name.Value != "i" {
		t.Errorf("init.Name.Value not 'i'. got=%s", init.Name.Value)
	}

	// Check condition
	if !testBinaryExpression(t, stmt.Condition, "i", "<", 5.0) {
		return
	}

	// Check body
	if len(stmt.Body.Statements) != 1 {
		t.Errorf("body is not 1 statement. got=%d\n",
			len(stmt.Body.Statements))
	}
}

func TestFunctionLiteral(t *testing.T) {
	input := `kaj(a, b) { ferao a + b; };`

	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("program.Statements does not contain %d statements. got=%d\n",
			1, len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not *ast.ExpressionStatement. got=%T",
			program.Statements[0])
	}

	function, ok := stmt.Expression.(*ast.FunctionLiteral)
	if !ok {
		t.Fatalf("stmt.Expression is not *ast.FunctionLiteral. got=%T",
			stmt.Expression)
	}

	if len(function.Parameters) != 2 {
		t.Fatalf("function literal parameters wrong. want 2, got=%d\n",
			len(function.Parameters))
	}

	testLiteralExpression(t, function.Parameters[0], "a")
	testLiteralExpression(t, function.Parameters[1], "b")

	if len(function.Body.Statements) != 1 {
		t.Fatalf("function.Body.Statements has not 1 statement. got=%d\n",
			len(function.Body.Statements))
	}

	bodyStmt, ok := function.Body.Statements[0].(*ast.ReturnStatement)
	if !ok {
		t.Fatalf("function body stmt is not *ast.ReturnStatement. got=%T",
			function.Body.Statements[0])
	}

	testBinaryExpression(t, bodyStmt.ReturnValue, "a", "+", "b")
}

func TestNamedFunctionLiteral(t *testing.T) {
	input := `kaj add(a, b) { ferao a + b; }`

	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not *ast.ExpressionStatement. got=%T",
			program.Statements[0])
	}

	function, ok := stmt.Expression.(*ast.FunctionLiteral)
	if !ok {
		t.Fatalf("stmt.Expression is not *ast.FunctionLiteral. got=%T",
			stmt.Expression)
	}

	if function.Name == nil || function.Name.Value != "add" {
		t.Errorf("function name not 'add'. got=%v", function.Name)
	}
}

func TestCallExpression(t *testing.T) {
	input := "add(1, 2 * 3, 4 + 5);"

	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("program.Statements does not contain %d statements. got=%d\n",
			1, len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not *ast.ExpressionStatement. got=%T",
			program.Statements[0])
	}

	exp, ok := stmt.Expression.(*ast.CallExpression)
	if !ok {
		t.Fatalf("stmt.Expression is not *ast.CallExpression. got=%T",
			stmt.Expression)
	}

	if !testIdentifier(t, exp.Function, "add") {
		return
	}

	if len(exp.Arguments) != 3 {
		t.Fatalf("wrong length of arguments. got=%d", len(exp.Arguments))
	}

	testLiteralExpression(t, exp.Arguments[0], 1.0)
	testBinaryExpression(t, exp.Arguments[1], 2.0, "*", 3.0)
	testBinaryExpression(t, exp.Arguments[2], 4.0, "+", 5.0)
}

func TestArrayLiteral(t *testing.T) {
	input := "[1, 2 * 2, 3 + 3];"

	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	stmt := program.Statements[0].(*ast.ExpressionStatement)
	array, ok := stmt.Expression.(*ast.ArrayLiteral)
	if !ok {
		t.Fatalf("exp not *ast.ArrayLiteral. got=%T", stmt.Expression)
	}

	if len(array.Elements) != 3 {
		t.Fatalf("len(array.Elements) not 3. got=%d", len(array.Elements))
	}

	testNumberLiteral(t, array.Elements[0], 1)
	testBinaryExpression(t, array.Elements[1], 2.0, "*", 2.0)
	testBinaryExpression(t, array.Elements[2], 3.0, "+", 3.0)
}

func TestMapLiteral(t *testing.T) {
	input := `{naam: "Ankan", boyes: 25};`

	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	stmt := program.Statements[0].(*ast.ExpressionStatement)
	mapLit, ok := stmt.Expression.(*ast.MapLiteral)
	if !ok {
		t.Fatalf("exp not *ast.MapLiteral. got=%T", stmt.Expression)
	}

	if len(mapLit.Pairs) != 2 {
		t.Errorf("mapLit.Pairs has wrong length. got=%d", len(mapLit.Pairs))
	}
}

func TestMemberExpression(t *testing.T) {
	tests := []struct {
		input    string
		object   string
		property string
		computed bool
	}{
		{"myArray[0];", "myArray", "0", true},
		{"obj.name;", "obj", "name", false},
	}

	for _, tt := range tests {
		l := lexer.New(tt.input)
		p := parser.New(l)
		program := p.ParseProgram()
		checkParserErrors(t, p)

		stmt := program.Statements[0].(*ast.ExpressionStatement)
		member, ok := stmt.Expression.(*ast.MemberExpression)
		if !ok {
			t.Fatalf("exp not *ast.MemberExpression. got=%T", stmt.Expression)
		}

		if member.Computed != tt.computed {
			t.Errorf("member.Computed not %t. got=%t", tt.computed, member.Computed)
		}
	}
}

func TestParseClassDeclaration(t *testing.T) {
	input := `sreni Person { shuru(naam) { ei.naam = naam; } kaj greet() { ferao "Hello"; } }`

	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("program.Statements does not contain %d statements. got=%d\n",
			1, len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.ClassDeclaration)
	if !ok {
		t.Fatalf("program.Statements[0] is not *ast.ClassDeclaration. got=%T",
			program.Statements[0])
	}

	if stmt.Name.Value != "Person" {
		t.Errorf("class name not 'Person'. got=%s", stmt.Name.Value)
	}

	if len(stmt.Methods) != 2 {
		t.Errorf("class has wrong number of methods. got=%d", len(stmt.Methods))
	}
}

func TestParseBreakStatement(t *testing.T) {
	input := "thamo;"

	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	stmt, ok := program.Statements[0].(*ast.BreakStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not *ast.BreakStatement. got=%T",
			program.Statements[0])
	}

	if stmt.TokenLiteral() != "thamo" {
		t.Errorf("stmt.TokenLiteral not 'thamo'. got=%s", stmt.TokenLiteral())
	}
}

func TestParseContinueStatement(t *testing.T) {
	input := "chharo;"

	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	stmt, ok := program.Statements[0].(*ast.ContinueStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not *ast.ContinueStatement. got=%T",
			program.Statements[0])
	}

	if stmt.TokenLiteral() != "chharo" {
		t.Errorf("stmt.TokenLiteral not 'chharo'. got=%s", stmt.TokenLiteral())
	}
}

func TestTryCatchStatement(t *testing.T) {
	input := `chesta { felo "error"; } dhoro_bhul (e) { dekho(e); } shesh { dekho("done"); }`

	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("program.Statements does not contain %d statements. got=%d\n",
			1, len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.TryCatchStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not *ast.TryCatchStatement. got=%T",
			program.Statements[0])
	}

	if stmt.TryBlock == nil {
		t.Error("TryBlock was nil")
	}

	if stmt.CatchBlock == nil {
		t.Error("CatchBlock was nil")
	}

	if stmt.CatchParam == nil || stmt.CatchParam.Value != "e" {
		t.Error("CatchParam was not 'e'")
	}

	if stmt.FinallyBlock == nil {
		t.Error("FinallyBlock was nil")
	}
}

func TestThrowStatement(t *testing.T) {
	input := `felo "error message";`

	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	stmt, ok := program.Statements[0].(*ast.ThrowStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not *ast.ThrowStatement. got=%T",
			program.Statements[0])
	}

	if stmt.TokenLiteral() != "felo" {
		t.Errorf("stmt.TokenLiteral not 'felo'. got=%s", stmt.TokenLiteral())
	}
}

func TestNewExpression(t *testing.T) {
	input := `notun Person("Ankan");`

	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	stmt := program.Statements[0].(*ast.ExpressionStatement)
	newExp, ok := stmt.Expression.(*ast.NewExpression)
	if !ok {
		t.Fatalf("exp not *ast.NewExpression. got=%T", stmt.Expression)
	}

	if !testIdentifier(t, newExp.Class, "Person") {
		return
	}

	if len(newExp.Arguments) != 1 {
		t.Errorf("wrong number of arguments. got=%d", len(newExp.Arguments))
	}
}

func TestAssignmentExpression(t *testing.T) {
	tests := []struct {
		input    string
		operator string
	}{
		{"x = 5;", "="},
		{"x += 5;", "+="},
		{"x -= 5;", "-="},
		{"x *= 5;", "*="},
		{"x /= 5;", "/="},
	}

	for _, tt := range tests {
		l := lexer.New(tt.input)
		p := parser.New(l)
		program := p.ParseProgram()
		checkParserErrors(t, p)

		stmt := program.Statements[0].(*ast.ExpressionStatement)
		assign, ok := stmt.Expression.(*ast.AssignmentExpression)
		if !ok {
			t.Fatalf("exp not *ast.AssignmentExpression. got=%T", stmt.Expression)
		}

		if assign.Operator != tt.operator {
			t.Errorf("assign.Operator not '%s'. got=%s", tt.operator, assign.Operator)
		}
	}
}

// Helper functions

func checkParserErrors(t *testing.T, p *parser.Parser) {
	errors := p.Errors()
	if len(errors) == 0 {
		return
	}

	t.Errorf("parser has %d errors", len(errors))
	for _, msg := range errors {
		t.Errorf("parser error: %q", msg)
	}
	t.FailNow()
}

func testLiteralExpression(t *testing.T, exp ast.Expression, expected interface{}) bool {
	switch v := expected.(type) {
	case int:
		return testNumberLiteral(t, exp, float64(v))
	case float64:
		return testNumberLiteral(t, exp, v)
	case string:
		return testIdentifier(t, exp, v)
	case bool:
		return testBooleanLiteral(t, exp, v)
	}
	t.Errorf("type of exp not handled. got=%T", exp)
	return false
}

func testNumberLiteral(t *testing.T, exp ast.Expression, value float64) bool {
	num, ok := exp.(*ast.NumberLiteral)
	if !ok {
		t.Errorf("exp not *ast.NumberLiteral. got=%T", exp)
		return false
	}

	if num.Value != value {
		t.Errorf("num.Value not %f. got=%f", value, num.Value)
		return false
	}

	return true
}

func testIdentifier(t *testing.T, exp ast.Expression, value string) bool {
	ident, ok := exp.(*ast.Identifier)
	if !ok {
		t.Errorf("exp not *ast.Identifier. got=%T", exp)
		return false
	}

	if ident.Value != value {
		t.Errorf("ident.Value not %s. got=%s", value, ident.Value)
		return false
	}

	return true
}

func testBooleanLiteral(t *testing.T, exp ast.Expression, value bool) bool {
	bo, ok := exp.(*ast.BooleanLiteral)
	if !ok {
		t.Errorf("exp not *ast.BooleanLiteral. got=%T", exp)
		return false
	}

	if bo.Value != value {
		t.Errorf("bo.Value not %t. got=%t", value, bo.Value)
		return false
	}

	return true
}

func testBinaryExpression(t *testing.T, exp ast.Expression, left interface{}, operator string, right interface{}) bool {
	binExp, ok := exp.(*ast.BinaryExpression)
	if !ok {
		t.Errorf("exp not *ast.BinaryExpression. got=%T(%s)", exp, exp)
		return false
	}

	if !testLiteralExpression(t, binExp.Left, left) {
		return false
	}

	if binExp.Operator != operator {
		t.Errorf("binExp.Operator not '%s'. got=%s", operator, binExp.Operator)
		return false
	}

	if !testLiteralExpression(t, binExp.Right, right) {
		return false
	}

	return true
}
