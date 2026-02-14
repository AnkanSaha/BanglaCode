package test

import (
	"BanglaCode/src/evaluator"
	"BanglaCode/src/lexer"
	"BanglaCode/src/object"
	"BanglaCode/src/parser"
	"testing"
)

func TestEvalNumberExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected float64
	}{
		{"5", 5},
		{"10", 10},
		{"-5", -5},
		{"-10", -10},
		{"5 + 5 + 5 + 5 - 10", 10},
		{"2 * 2 * 2 * 2 * 2", 32},
		{"-50 + 100 + -50", 0},
		{"5 * 2 + 10", 20},
		{"5 + 2 * 10", 25},
		{"20 + 2 * -10", 0},
		{"50 / 2 * 2 + 10", 60},
		{"2 * (5 + 10)", 30},
		{"3 * 3 * 3 + 10", 37},
		{"3 * (3 * 3) + 10", 37},
		{"(5 + 10 * 2 + 15 / 3) * 2 + -10", 50},
		{"10 % 3", 1},
		{"10.5 + 0.5", 11},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testNumberObject(t, evaluated, tt.expected)
	}
}

func TestEvalBooleanExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"sotti", true},
		{"mittha", false},
		{"1 < 2", true},
		{"1 > 2", false},
		{"1 < 1", false},
		{"1 > 1", false},
		{"1 == 1", true},
		{"1 != 1", false},
		{"1 == 2", false},
		{"1 != 2", true},
		{"1 <= 1", true},
		{"1 >= 1", true},
		{"1 <= 2", true},
		{"2 >= 1", true},
		{"sotti == sotti", true},
		{"mittha == mittha", true},
		{"sotti == mittha", false},
		{"sotti != mittha", true},
		{"mittha != sotti", true},
		{"(1 < 2) == sotti", true},
		{"(1 < 2) == mittha", false},
		{"(1 > 2) == sotti", false},
		{"(1 > 2) == mittha", true},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testBooleanObject(t, evaluated, tt.expected)
	}
}

func TestBangOperator(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"!sotti", false},
		{"!mittha", true},
		{"!5", false},
		{"!!sotti", true},
		{"!!mittha", false},
		{"!!5", true},
		{"na sotti", false},
		{"na mittha", true},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testBooleanObject(t, evaluated, tt.expected)
	}
}

func TestLogicalOperators(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"sotti ebong sotti", true},
		{"sotti ebong mittha", false},
		{"mittha ebong sotti", false},
		{"mittha ebong mittha", false},
		{"sotti ba sotti", true},
		{"sotti ba mittha", true},
		{"mittha ba sotti", true},
		{"mittha ba mittha", false},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testBooleanObject(t, evaluated, tt.expected)
	}
}

func TestIfElseExpressions(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{"jodi (sotti) { 10 }", 10.0},
		{"jodi (mittha) { 10 }", nil},
		{"jodi (1) { 10 }", 10.0},
		{"jodi (1 < 2) { 10 }", 10.0},
		{"jodi (1 > 2) { 10 }", nil},
		{"jodi (1 > 2) { 10 } nahole { 20 }", 20.0},
		{"jodi (1 < 2) { 10 } nahole { 20 }", 10.0},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		if integer, ok := tt.expected.(float64); ok {
			testNumberObject(t, evaluated, integer)
		} else {
			testNullObject(t, evaluated)
		}
	}
}

func TestEvalVariableDeclaration(t *testing.T) {
	tests := []struct {
		input    string
		expected float64
	}{
		{"dhoro a = 5; a;", 5},
		{"dhoro a = 5 * 5; a;", 25},
		{"dhoro a = 5; dhoro b = a; b;", 5},
		{"dhoro a = 5; dhoro b = a; dhoro c = a + b + 5; c;", 15},
	}

	for _, tt := range tests {
		testNumberObject(t, testEval(tt.input), tt.expected)
	}
}

func TestReturnStatements(t *testing.T) {
	tests := []struct {
		input    string
		expected float64
	}{
		{"ferao 10;", 10},
		{"ferao 10; 9;", 10},
		{"ferao 2 * 5; 9;", 10},
		{"9; ferao 2 * 5; 9;", 10},
		{`
jodi (10 > 1) {
  jodi (10 > 1) {
    ferao 10;
  }
  ferao 1;
}
`, 10},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testNumberObject(t, evaluated, tt.expected)
	}
}

func TestFunctionObject(t *testing.T) {
	input := "kaj(x) { x + 2; };"

	evaluated := testEval(input)
	fn, ok := evaluated.(*object.Function)
	if !ok {
		t.Fatalf("object is not Function. got=%T (%+v)", evaluated, evaluated)
	}

	if len(fn.Parameters) != 1 {
		t.Fatalf("function has wrong parameters. Parameters=%+v",
			fn.Parameters)
	}

	if fn.Parameters[0].String() != "x" {
		t.Fatalf("parameter is not 'x'. got=%q", fn.Parameters[0])
	}
}

func TestFunctionApplication(t *testing.T) {
	tests := []struct {
		input    string
		expected float64
	}{
		{"dhoro identity = kaj(x) { x; }; identity(5);", 5},
		{"dhoro identity = kaj(x) { ferao x; }; identity(5);", 5},
		{"dhoro double = kaj(x) { x * 2; }; double(5);", 10},
		{"dhoro add = kaj(x, y) { x + y; }; add(5, 5);", 10},
		{"dhoro add = kaj(x, y) { x + y; }; add(5 + 5, add(5, 5));", 20},
		{"kaj(x) { x; }(5)", 5},
	}

	for _, tt := range tests {
		testNumberObject(t, testEval(tt.input), tt.expected)
	}
}

func TestNamedFunction(t *testing.T) {
	input := `
	kaj add(a, b) { ferao a + b; }
	add(3, 4);
	`

	evaluated := testEval(input)
	testNumberObject(t, evaluated, 7)
}

func TestRecursiveFunction(t *testing.T) {
	input := `
	kaj factorial(n) {
		jodi (n <= 1) {
			ferao 1;
		} nahole {
			ferao n * factorial(n - 1);
		}
	}
	factorial(5);
	`

	evaluated := testEval(input)
	testNumberObject(t, evaluated, 120)
}

func TestClosures(t *testing.T) {
	input := `
dhoro newAdder = kaj(x) {
  ferao kaj(y) { x + y };
};

dhoro addTwo = newAdder(2);
addTwo(2);`

	testNumberObject(t, testEval(input), 4)
}

func TestStringLiteral(t *testing.T) {
	input := `"Hello World!"`

	evaluated := testEval(input)
	str, ok := evaluated.(*object.String)
	if !ok {
		t.Fatalf("object is not String. got=%T (%+v)", evaluated, evaluated)
	}

	if str.Value != "Hello World!" {
		t.Errorf("String has wrong value. got=%q", str.Value)
	}
}

func TestStringConcatenation(t *testing.T) {
	input := `"Hello" + " " + "World!"`

	evaluated := testEval(input)
	str, ok := evaluated.(*object.String)
	if !ok {
		t.Fatalf("object is not String. got=%T (%+v)", evaluated, evaluated)
	}

	if str.Value != "Hello World!" {
		t.Errorf("String has wrong value. got=%q", str.Value)
	}
}

func TestStringNumberConcatenation(t *testing.T) {
	input := `"Value: " + 42`

	evaluated := testEval(input)
	str, ok := evaluated.(*object.String)
	if !ok {
		t.Fatalf("object is not String. got=%T (%+v)", evaluated, evaluated)
	}

	if str.Value != "Value: 42" {
		t.Errorf("String has wrong value. got=%q", str.Value)
	}
}

func TestArrayLiterals(t *testing.T) {
	input := "[1, 2 * 2, 3 + 3]"

	evaluated := testEval(input)
	result, ok := evaluated.(*object.Array)
	if !ok {
		t.Fatalf("object is not Array. got=%T (%+v)", evaluated, evaluated)
	}

	if len(result.Elements) != 3 {
		t.Fatalf("array has wrong num of elements. got=%d",
			len(result.Elements))
	}

	testNumberObject(t, result.Elements[0], 1)
	testNumberObject(t, result.Elements[1], 4)
	testNumberObject(t, result.Elements[2], 6)
}

func TestArrayIndexExpressions(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{"[1, 2, 3][0]", 1.0},
		{"[1, 2, 3][1]", 2.0},
		{"[1, 2, 3][2]", 3.0},
		{"dhoro i = 0; [1][i];", 1.0},
		{"[1, 2, 3][1 + 1];", 3.0},
		{"dhoro myArray = [1, 2, 3]; myArray[2];", 3.0},
		{"dhoro myArray = [1, 2, 3]; myArray[0] + myArray[1] + myArray[2];", 6.0},
		{"dhoro myArray = [1, 2, 3]; dhoro i = myArray[0]; myArray[i]", 2.0},
		{"[1, 2, 3][3]", nil},
		{"[1, 2, 3][-1]", nil},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		if integer, ok := tt.expected.(float64); ok {
			testNumberObject(t, evaluated, integer)
		} else {
			testNullObject(t, evaluated)
		}
	}
}

func TestMapLiterals(t *testing.T) {
	input := `{naam: "Ankan", boyes: 25}`

	evaluated := testEval(input)
	result, ok := evaluated.(*object.Map)
	if !ok {
		t.Fatalf("object is not Map. got=%T (%+v)", evaluated, evaluated)
	}

	expected := map[string]interface{}{
		"naam":  "Ankan",
		"boyes": 25.0,
	}

	if len(result.Pairs) != len(expected) {
		t.Fatalf("Map has wrong num of pairs. got=%d", len(result.Pairs))
	}

	for key, val := range expected {
		v, ok := result.Pairs[key]
		if !ok {
			t.Errorf("no pair for given key: %s", key)
		}
		switch expected := val.(type) {
		case string:
			str, ok := v.(*object.String)
			if !ok {
				t.Errorf("value not string. got=%T", v)
			}
			if str.Value != expected {
				t.Errorf("wrong string value. got=%q", str.Value)
			}
		case float64:
			testNumberObject(t, v, expected)
		}
	}
}

func TestMapAccess(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{`{naam: "Ankan"}.naam`, "Ankan"},
		{`{x: 5}.x`, 5.0},
		{`{x: 5}["x"]`, 5.0},
		{`dhoro key = "x"; {x: 5}[key]`, 5.0},
		{`{}.foo`, nil},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		switch expected := tt.expected.(type) {
		case string:
			str, ok := evaluated.(*object.String)
			if !ok {
				t.Errorf("object is not String. got=%T", evaluated)
			} else if str.Value != expected {
				t.Errorf("wrong string. got=%q", str.Value)
			}
		case float64:
			testNumberObject(t, evaluated, expected)
		default:
			testNullObject(t, evaluated)
		}
	}
}

func TestWhileLoop(t *testing.T) {
	input := `
	dhoro i = 0;
	dhoro sum = 0;
	jotokkhon (i < 5) {
		sum = sum + i;
		i = i + 1;
	}
	sum;
	`

	evaluated := testEval(input)
	testNumberObject(t, evaluated, 10) // 0+1+2+3+4 = 10
}

func TestForLoop(t *testing.T) {
	// Use array mutation since it works across scopes
	input := `
	dhoro result = [0];
	ghuriye (dhoro i = 0; i < 5; i = i + 1) {
		result[0] = result[0] + i;
	}
	result[0];
	`

	evaluated := testEval(input)
	testNumberObject(t, evaluated, 10) // 0+1+2+3+4 = 10
}

func TestEvalBreakStatement(t *testing.T) {
	input := `
	dhoro i = 0;
	jotokkhon (sotti) {
		jodi (i >= 5) { thamo; }
		i = i + 1;
	}
	i;
	`

	evaluated := testEval(input)
	testNumberObject(t, evaluated, 5)
}

func TestEvalContinueStatement(t *testing.T) {
	// Use array mutation since it works across scopes
	input := `
	dhoro result = [0];
	ghuriye (dhoro i = 0; i < 10; i = i + 1) {
		jodi (i % 2 == 1) { chharo; }
		result[0] = result[0] + i;
	}
	result[0];
	`

	evaluated := testEval(input)
	testNumberObject(t, evaluated, 20) // 0+2+4+6+8 = 20
}

func TestEvalClassDeclaration(t *testing.T) {
	input := `
	sreni Person {
		shuru(naam) {
			ei.naam = naam;
		}
		kaj getName() {
			ferao ei.naam;
		}
	}
	dhoro p = notun Person("Ankan");
	p.getName();
	`

	evaluated := testEval(input)
	str, ok := evaluated.(*object.String)
	if !ok {
		t.Fatalf("object is not String. got=%T (%+v)", evaluated, evaluated)
	}
	if str.Value != "Ankan" {
		t.Errorf("wrong value. got=%q", str.Value)
	}
}

func TestClassProperties(t *testing.T) {
	input := `
	sreni Counter {
		shuru() {
			ei.count = 0;
		}
		kaj increment() {
			ei.count = ei.count + 1;
			ferao ei.count;
		}
	}
	dhoro c = notun Counter();
	c.increment();
	c.increment();
	c.increment();
	`

	evaluated := testEval(input)
	testNumberObject(t, evaluated, 3)
}

func TestCompoundAssignment(t *testing.T) {
	tests := []struct {
		input    string
		expected float64
	}{
		{"dhoro x = 5; x += 3; x;", 8},
		{"dhoro x = 10; x -= 3; x;", 7},
		{"dhoro x = 5; x *= 2; x;", 10},
		{"dhoro x = 10; x /= 2; x;", 5},
	}

	for _, tt := range tests {
		testNumberObject(t, testEval(tt.input), tt.expected)
	}
}

func TestArrayMutation(t *testing.T) {
	input := `
	dhoro arr = [1, 2, 3];
	arr[1] = 10;
	arr[1];
	`

	evaluated := testEval(input)
	testNumberObject(t, evaluated, 10)
}

func TestMapMutation(t *testing.T) {
	input := `
	dhoro obj = {x: 1};
	obj.x = 10;
	obj.x;
	`

	evaluated := testEval(input)
	testNumberObject(t, evaluated, 10)
}

func TestTryCatch(t *testing.T) {
	// Use array mutation since it works across scopes
	input := `
	dhoro result = [0];
	chesta {
		felo "error";
	} dhoro_bhul (e) {
		result[0] = 1;
	}
	result[0];
	`

	evaluated := testEval(input)
	testNumberObject(t, evaluated, 1)
}

func TestTryCatchFinally(t *testing.T) {
	// Use array mutation since it works across scopes
	input := `
	dhoro result = [""];
	chesta {
		felo "error";
	} dhoro_bhul (e) {
		result[0] = result[0] + "catch";
	} shesh {
		result[0] = result[0] + "-finally";
	}
	result[0];
	`

	evaluated := testEval(input)
	str, ok := evaluated.(*object.String)
	if !ok {
		t.Fatalf("object is not String. got=%T (%+v)", evaluated, evaluated)
	}
	if str.Value != "catch-finally" {
		t.Errorf("wrong value. got=%q", str.Value)
	}
}

func TestErrorHandling(t *testing.T) {
	tests := []struct {
		input           string
		expectedMessage string
	}{
		{"5 + sotti;", "unknown operator: NUMBER + BOOLEAN"},
		{"-sotti", "unknown operator: -BOOLEAN"},
		{"sotti + mittha;", "unknown operator: BOOLEAN + BOOLEAN"},
		{"foobar", "variable 'foobar' is not defined"},
		{"5 / 0", "division by zero"},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)

		errObj, ok := evaluated.(*object.Error)
		if !ok {
			t.Errorf("no error object returned. got=%T(%+v)",
				evaluated, evaluated)
			continue
		}

		if errObj.Message != tt.expectedMessage {
			t.Errorf("wrong error message. expected=%q, got=%q",
				tt.expectedMessage, errObj.Message)
		}
	}
}

func TestNullValue(t *testing.T) {
	input := "khali"

	evaluated := testEval(input)
	testNullObject(t, evaluated)
}

// Helper functions

func testEval(input string) object.Object {
	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()
	env := object.NewEnvironment()

	return evaluator.Eval(program, env)
}

func testNumberObject(t *testing.T, obj object.Object, expected float64) bool {
	result, ok := obj.(*object.Number)
	if !ok {
		t.Errorf("object is not Number. got=%T (%+v)", obj, obj)
		return false
	}
	if result.Value != expected {
		t.Errorf("object has wrong value. got=%f, want=%f",
			result.Value, expected)
		return false
	}

	return true
}

func testBooleanObject(t *testing.T, obj object.Object, expected bool) bool {
	result, ok := obj.(*object.Boolean)
	if !ok {
		t.Errorf("object is not Boolean. got=%T (%+v)", obj, obj)
		return false
	}
	if result.Value != expected {
		t.Errorf("object has wrong value. got=%t, want=%t",
			result.Value, expected)
		return false
	}
	return true
}

func testNullObject(t *testing.T, obj object.Object) bool {
	if obj != object.NULL {
		t.Errorf("object is not NULL. got=%T (%+v)", obj, obj)
		return false
	}
	return true
}

func testStringObject(t *testing.T, obj object.Object, expected string) bool {
	t.Helper()
	result, ok := obj.(*object.String)
	if !ok {
		t.Errorf("object is not String. got=%T (%+v)", obj, obj)
		return false
	}
	if result.Value != expected {
		t.Errorf("object has wrong value. got=%s, want=%s", result.Value, expected)
		return false
	}
	return true
}
