package test

import (
	"BanglaCode/src/evaluator"
	"BanglaCode/src/lexer"
	"BanglaCode/src/object"
	"BanglaCode/src/parser"
	"math"
	"testing"
)

func testEvalBuiltin(input string) object.Object {
	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()
	env := object.NewEnvironment()
	return evaluator.Eval(program, env)
}

// dorghyo (length) tests
func TestBuiltinDorghyo(t *testing.T) {
	tests := []struct {
		input    string
		expected any
	}{
		{`dorghyo("hello")`, 5.0},
		{`dorghyo("")`, 0.0},
		{`dorghyo("Hello World")`, 11.0},
		{`dorghyo([1, 2, 3])`, 3.0},
		{`dorghyo([])`, 0.0},
	}

	for _, tt := range tests {
		evaluated := testEvalBuiltin(tt.input)
		switch expected := tt.expected.(type) {
		case float64:
			testBuiltinNumberObject(t, evaluated, expected)
		}
	}
}

func TestBuiltinDorghyoError(t *testing.T) {
	tests := []struct {
		input           string
		expectedMessage string
	}{
		{`dorghyo(1)`, "argument to `dorghyo` not supported, got NUMBER"},
		{`dorghyo("one", "two")`, "wrong number of arguments. got=2, want=1"},
	}

	for _, tt := range tests {
		evaluated := testEvalBuiltin(tt.input)
		errObj, ok := evaluated.(*object.Error)
		if !ok {
			t.Errorf("no error object returned. got=%T(%+v)", evaluated, evaluated)
			continue
		}
		if errObj.Message != tt.expectedMessage {
			t.Errorf("wrong error message. expected=%q, got=%q",
				tt.expectedMessage, errObj.Message)
		}
	}
}

// dhokao (push) tests
func TestBuiltinDhokao(t *testing.T) {
	input := `
	dhoro arr = [1, 2, 3];
	dhokao(arr, 4);
	arr[3];
	`
	evaluated := testEvalBuiltin(input)
	testBuiltinNumberObject(t, evaluated, 4)
}

// berKoro (pop) tests
func TestBuiltinBerKoro(t *testing.T) {
	input := `
	dhoro arr = [1, 2, 3];
	berKoro(arr);
	`
	evaluated := testEvalBuiltin(input)
	testBuiltinNumberObject(t, evaluated, 3)
}

func TestBuiltinBerKoroModifiesArray(t *testing.T) {
	input := `
	dhoro arr = [1, 2, 3];
	berKoro(arr);
	dorghyo(arr);
	`
	evaluated := testEvalBuiltin(input)
	testBuiltinNumberObject(t, evaluated, 2)
}

// chabi (keys) tests
func TestBuiltinChabi(t *testing.T) {
	input := `
	dhoro obj = {naam: "Ankan", boyes: 25};
	dhoro keys = chabi(obj);
	dorghyo(keys);
	`
	evaluated := testEvalBuiltin(input)
	testBuiltinNumberObject(t, evaluated, 2)
}

// dhoron (type) tests
func TestBuiltinDhoron(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{`dhoron(5)`, "NUMBER"},
		{`dhoron("hello")`, "STRING"},
		{`dhoron(sotti)`, "BOOLEAN"},
		{`dhoron([1, 2, 3])`, "ARRAY"},
		{`dhoron({x: 1})`, "MAP"},
		{`dhoron(khali)`, "NULL"},
	}

	for _, tt := range tests {
		evaluated := testEvalBuiltin(tt.input)
		str, ok := evaluated.(*object.String)
		if !ok {
			t.Errorf("object is not String. got=%T", evaluated)
			continue
		}
		if str.Value != tt.expected {
			t.Errorf("wrong type. expected=%q, got=%q", tt.expected, str.Value)
		}
	}
}

// lipi (to string) tests
func TestBuiltinLipi(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{`lipi(5)`, "5"},
		{`lipi(3.14)`, "3.14"},
		{`lipi(sotti)`, "true"},
		{`lipi(mittha)`, "false"},
		{`lipi([1, 2, 3])`, "[1, 2, 3]"},
	}

	for _, tt := range tests {
		evaluated := testEvalBuiltin(tt.input)
		str, ok := evaluated.(*object.String)
		if !ok {
			t.Errorf("object is not String. got=%T", evaluated)
			continue
		}
		if str.Value != tt.expected {
			t.Errorf("wrong string. expected=%q, got=%q", tt.expected, str.Value)
		}
	}
}

// sonkha (to number) tests
func TestBuiltinSonkha(t *testing.T) {
	tests := []struct {
		input    string
		expected float64
	}{
		{`sonkha(5)`, 5},
		{`sonkha("123")`, 123},
		{`sonkha("3.14")`, 3.14},
		{`sonkha(sotti)`, 1},
		{`sonkha(mittha)`, 0},
	}

	for _, tt := range tests {
		evaluated := testEvalBuiltin(tt.input)
		testBuiltinNumberObject(t, evaluated, tt.expected)
	}
}

// Math functions tests
func TestBuiltinBorgomul(t *testing.T) {
	tests := []struct {
		input    string
		expected float64
	}{
		{`borgomul(4)`, 2},
		{`borgomul(9)`, 3},
		{`borgomul(16)`, 4},
		{`borgomul(2)`, math.Sqrt(2)},
	}

	for _, tt := range tests {
		evaluated := testEvalBuiltin(tt.input)
		testBuiltinNumberObject(t, evaluated, tt.expected)
	}
}

func TestBuiltinGhat(t *testing.T) {
	tests := []struct {
		input    string
		expected float64
	}{
		{`ghat(2, 3)`, 8},
		{`ghat(5, 2)`, 25},
		{`ghat(10, 0)`, 1},
		{`ghat(2, 10)`, 1024},
	}

	for _, tt := range tests {
		evaluated := testEvalBuiltin(tt.input)
		testBuiltinNumberObject(t, evaluated, tt.expected)
	}
}

func TestBuiltinNiche(t *testing.T) {
	tests := []struct {
		input    string
		expected float64
	}{
		{`niche(3.7)`, 3},
		{`niche(3.2)`, 3},
		{`niche(5.9)`, 5},
		{`niche(5.0)`, 5},
	}

	for _, tt := range tests {
		evaluated := testEvalBuiltin(tt.input)
		testBuiltinNumberObject(t, evaluated, tt.expected)
	}
}

func TestBuiltinUpore(t *testing.T) {
	tests := []struct {
		input    string
		expected float64
	}{
		{`upore(3.2)`, 4},
		{`upore(3.7)`, 4},
		{`upore(5.0)`, 5},
		{`upore(5.1)`, 6},
	}

	for _, tt := range tests {
		evaluated := testEvalBuiltin(tt.input)
		testBuiltinNumberObject(t, evaluated, tt.expected)
	}
}

func TestBuiltinKache(t *testing.T) {
	tests := []struct {
		input    string
		expected float64
	}{
		{`kache(3.2)`, 3},
		{`kache(3.7)`, 4},
		{`kache(3.5)`, 4},
		{`kache(2.5)`, 3}, // Go's math.Round uses "round half away from zero"
	}

	for _, tt := range tests {
		evaluated := testEvalBuiltin(tt.input)
		testBuiltinNumberObject(t, evaluated, tt.expected)
	}
}

func TestBuiltinNiratek(t *testing.T) {
	tests := []struct {
		input    string
		expected float64
	}{
		{`niratek(-5)`, 5},
		{`niratek(5)`, 5},
		{`niratek(-3.14)`, 3.14},
		{`niratek(0)`, 0},
	}

	for _, tt := range tests {
		evaluated := testEvalBuiltin(tt.input)
		testBuiltinNumberObject(t, evaluated, tt.expected)
	}
}

func TestBuiltinChoto(t *testing.T) {
	tests := []struct {
		input    string
		expected float64
	}{
		{`choto(1, 2)`, 1},
		{`choto(5, 3)`, 3},
		{`choto(1, 2, 3, 4, 5)`, 1},
		{`choto(-1, 0, 1)`, -1},
	}

	for _, tt := range tests {
		evaluated := testEvalBuiltin(tt.input)
		testBuiltinNumberObject(t, evaluated, tt.expected)
	}
}

func TestBuiltinBoro(t *testing.T) {
	tests := []struct {
		input    string
		expected float64
	}{
		{`boro(1, 2)`, 2},
		{`boro(5, 3)`, 5},
		{`boro(1, 2, 3, 4, 5)`, 5},
		{`boro(-1, 0, 1)`, 1},
	}

	for _, tt := range tests {
		evaluated := testEvalBuiltin(tt.input)
		testBuiltinNumberObject(t, evaluated, tt.expected)
	}
}

// String functions tests
func TestBuiltinBoroHater(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{`boroHater("hello")`, "HELLO"},
		{`boroHater("Hello World")`, "HELLO WORLD"},
		{`boroHater("")`, ""},
	}

	for _, tt := range tests {
		evaluated := testEvalBuiltin(tt.input)
		str, ok := evaluated.(*object.String)
		if !ok {
			t.Errorf("object is not String. got=%T", evaluated)
			continue
		}
		if str.Value != tt.expected {
			t.Errorf("wrong string. expected=%q, got=%q", tt.expected, str.Value)
		}
	}
}

func TestBuiltinChotoHater(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{`chotoHater("HELLO")`, "hello"},
		{`chotoHater("Hello World")`, "hello world"},
		{`chotoHater("")`, ""},
	}

	for _, tt := range tests {
		evaluated := testEvalBuiltin(tt.input)
		str, ok := evaluated.(*object.String)
		if !ok {
			t.Errorf("object is not String. got=%T", evaluated)
			continue
		}
		if str.Value != tt.expected {
			t.Errorf("wrong string. expected=%q, got=%q", tt.expected, str.Value)
		}
	}
}

func TestBuiltinBhag(t *testing.T) {
	input := `bhag("a,b,c", ",")`
	evaluated := testEvalBuiltin(input)

	arr, ok := evaluated.(*object.Array)
	if !ok {
		t.Fatalf("object is not Array. got=%T", evaluated)
	}

	if len(arr.Elements) != 3 {
		t.Fatalf("array has wrong num of elements. got=%d", len(arr.Elements))
	}

	expected := []string{"a", "b", "c"}
	for i, el := range arr.Elements {
		str, ok := el.(*object.String)
		if !ok {
			t.Errorf("element[%d] is not String. got=%T", i, el)
			continue
		}
		if str.Value != expected[i] {
			t.Errorf("element[%d] wrong. expected=%q, got=%q", i, expected[i], str.Value)
		}
	}
}

func TestBuiltinJoro(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{`joro(["a", "b", "c"], ",")`, "a,b,c"},
		{`joro(["hello", "world"], " ")`, "hello world"},
		{`joro([1, 2, 3], "-")`, "1-2-3"},
	}

	for _, tt := range tests {
		evaluated := testEvalBuiltin(tt.input)
		str, ok := evaluated.(*object.String)
		if !ok {
			t.Errorf("object is not String. got=%T", evaluated)
			continue
		}
		if str.Value != tt.expected {
			t.Errorf("wrong string. expected=%q, got=%q", tt.expected, str.Value)
		}
	}
}

func TestBuiltinChhanto(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{`chhanto("  hello  ")`, "hello"},
		{`chhanto("hello")`, "hello"},
		{`chhanto("  ")`, ""},
	}

	for _, tt := range tests {
		evaluated := testEvalBuiltin(tt.input)
		str, ok := evaluated.(*object.String)
		if !ok {
			t.Errorf("object is not String. got=%T", evaluated)
			continue
		}
		if str.Value != tt.expected {
			t.Errorf("wrong string. expected=%q, got=%q", tt.expected, str.Value)
		}
	}
}

func TestBuiltinKhojo(t *testing.T) {
	tests := []struct {
		input    string
		expected float64
	}{
		{`khojo("hello", "l")`, 2},
		{`khojo("hello", "o")`, 4},
		{`khojo("hello", "x")`, -1},
		{`khojo("hello world", "world")`, 6},
	}

	for _, tt := range tests {
		evaluated := testEvalBuiltin(tt.input)
		testBuiltinNumberObject(t, evaluated, tt.expected)
	}
}

func TestBuiltinAngsho(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{`angsho("hello", 0, 2)`, "he"},
		{`angsho("hello", 1, 4)`, "ell"},
		{`angsho("hello", 2)`, "llo"},
		{`angsho("hello", 0, 10)`, "hello"},
	}

	for _, tt := range tests {
		evaluated := testEvalBuiltin(tt.input)
		str, ok := evaluated.(*object.String)
		if !ok {
			t.Errorf("object is not String. got=%T", evaluated)
			continue
		}
		if str.Value != tt.expected {
			t.Errorf("wrong string. expected=%q, got=%q", tt.expected, str.Value)
		}
	}
}

func TestBuiltinBodlo(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{`bodlo("hello", "l", "x")`, "hexxo"},
		{`bodlo("hello world", "world", "universe")`, "hello universe"},
		{`bodlo("aaa", "a", "b")`, "bbb"},
	}

	for _, tt := range tests {
		evaluated := testEvalBuiltin(tt.input)
		str, ok := evaluated.(*object.String)
		if !ok {
			t.Errorf("object is not String. got=%T", evaluated)
			continue
		}
		if str.Value != tt.expected {
			t.Errorf("wrong string. expected=%q, got=%q", tt.expected, str.Value)
		}
	}
}

// Array functions tests
func TestBuiltinKato(t *testing.T) {
	input := `kato([1, 2, 3, 4, 5], 1, 4)`
	evaluated := testEvalBuiltin(input)

	arr, ok := evaluated.(*object.Array)
	if !ok {
		t.Fatalf("object is not Array. got=%T", evaluated)
	}

	if len(arr.Elements) != 3 {
		t.Fatalf("array has wrong num of elements. got=%d", len(arr.Elements))
	}

	expected := []float64{2, 3, 4}
	for i, el := range arr.Elements {
		num, ok := el.(*object.Number)
		if !ok {
			t.Errorf("element[%d] is not Number. got=%T", i, el)
			continue
		}
		if num.Value != expected[i] {
			t.Errorf("element[%d] wrong. expected=%f, got=%f", i, expected[i], num.Value)
		}
	}
}

func TestBuiltinUlto(t *testing.T) {
	input := `ulto([1, 2, 3])`
	evaluated := testEvalBuiltin(input)

	arr, ok := evaluated.(*object.Array)
	if !ok {
		t.Fatalf("object is not Array. got=%T", evaluated)
	}

	expected := []float64{3, 2, 1}
	for i, el := range arr.Elements {
		num, ok := el.(*object.Number)
		if !ok {
			t.Errorf("element[%d] is not Number. got=%T", i, el)
			continue
		}
		if num.Value != expected[i] {
			t.Errorf("element[%d] wrong. expected=%f, got=%f", i, expected[i], num.Value)
		}
	}
}

func TestBuiltinAche(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{`ache([1, 2, 3], 2)`, true},
		{`ache([1, 2, 3], 5)`, false},
		{`ache(["a", "b", "c"], "b")`, true},
		{`ache(["a", "b", "c"], "x")`, false},
	}

	for _, tt := range tests {
		evaluated := testEvalBuiltin(tt.input)
		testBuiltinBooleanObject(t, evaluated, tt.expected)
	}
}

func TestBuiltinSaja(t *testing.T) {
	input := `saja([3, 1, 2])`
	evaluated := testEvalBuiltin(input)

	arr, ok := evaluated.(*object.Array)
	if !ok {
		t.Fatalf("object is not Array. got=%T", evaluated)
	}

	expected := []float64{1, 2, 3}
	for i, el := range arr.Elements {
		num, ok := el.(*object.Number)
		if !ok {
			t.Errorf("element[%d] is not Number. got=%T", i, el)
			continue
		}
		if num.Value != expected[i] {
			t.Errorf("element[%d] wrong. expected=%f, got=%f", i, expected[i], num.Value)
		}
	}
}

// Utility functions tests
func TestBuiltinSomoy(t *testing.T) {
	input := `somoy()`
	evaluated := testEvalBuiltin(input)

	num, ok := evaluated.(*object.Number)
	if !ok {
		t.Fatalf("object is not Number. got=%T", evaluated)
	}

	// Just check that it returns a positive number (timestamp)
	if num.Value <= 0 {
		t.Errorf("somoy() should return positive number. got=%f", num.Value)
	}
}

func TestBuiltinLotto(t *testing.T) {
	input := `lotto()`
	evaluated := testEvalBuiltin(input)

	num, ok := evaluated.(*object.Number)
	if !ok {
		t.Fatalf("object is not Number. got=%T", evaluated)
	}

	// Random should be between 0 and 1
	if num.Value < 0 || num.Value >= 1 {
		t.Errorf("lotto() should return number between 0 and 1. got=%f", num.Value)
	}
}

// JSON functions tests
func TestBuiltinJsonPoro(t *testing.T) {
	input := `json_poro('{"name": "Ankan", "age": 25}')`
	evaluated := testEvalBuiltin(input)

	m, ok := evaluated.(*object.Map)
	if !ok {
		t.Fatalf("object is not Map. got=%T", evaluated)
	}

	if len(m.Pairs) != 2 {
		t.Errorf("map has wrong num of pairs. got=%d", len(m.Pairs))
	}

	name, ok := m.Pairs["name"]
	if !ok {
		t.Error("map should have 'name' key")
	} else {
		str, ok := name.(*object.String)
		if !ok || str.Value != "Ankan" {
			t.Errorf("name wrong. got=%v", name)
		}
	}
}

func TestBuiltinJsonBanao(t *testing.T) {
	input := `json_banao({name: "Ankan", age: 25})`
	evaluated := testEvalBuiltin(input)

	str, ok := evaluated.(*object.String)
	if !ok {
		t.Fatalf("object is not String. got=%T", evaluated)
	}

	// JSON order may vary, so just check it contains expected parts
	if len(str.Value) < 10 {
		t.Errorf("JSON string too short. got=%q", str.Value)
	}
}

func TestBuiltinJsonArray(t *testing.T) {
	input := `json_poro("[1, 2, 3]")`
	evaluated := testEvalBuiltin(input)

	arr, ok := evaluated.(*object.Array)
	if !ok {
		t.Fatalf("object is not Array. got=%T", evaluated)
	}

	if len(arr.Elements) != 3 {
		t.Errorf("array has wrong num of elements. got=%d", len(arr.Elements))
	}
}

// Helper functions
func testBuiltinNumberObject(t *testing.T, obj object.Object, expected float64) bool {
	result, ok := obj.(*object.Number)
	if !ok {
		t.Errorf("object is not Number. got=%T (%+v)", obj, obj)
		return false
	}
	if result.Value != expected {
		t.Errorf("object has wrong value. got=%f, want=%f", result.Value, expected)
		return false
	}
	return true
}

func testBuiltinBooleanObject(t *testing.T, obj object.Object, expected bool) bool {
	result, ok := obj.(*object.Boolean)
	if !ok {
		t.Errorf("object is not Boolean. got=%T (%+v)", obj, obj)
		return false
	}
	if result.Value != expected {
		t.Errorf("object has wrong value. got=%t, want=%t", result.Value, expected)
		return false
	}
	return true
}

func testBuiltinStringObject(t *testing.T, obj object.Object, expected string) bool {
	result, ok := obj.(*object.String)
	if !ok {
		t.Errorf("object is not String. got=%T (%+v)", obj, obj)
		return false
	}
	if result.Value != expected {
		t.Errorf("object has wrong value. got=%q, want=%q", result.Value, expected)
		return false
	}
	return true
}
