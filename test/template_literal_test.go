package test

import (
	"BanglaCode/src/object"
	"testing"
)

func testStringObjectHelper(t *testing.T, obj object.Object, expected string) bool {
	t.Helper()
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

// ==================== Template Literal Tests ====================

func TestTemplateLiteralBasic(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		// Simple interpolation
		{
			"dhoro name = \"Ankan\";\n`Hello ${name}`",
			"Hello Ankan",
		},
		// Multiple interpolations
		{
			"dhoro naam = \"Ankan\";\ndhoro age = 25;\n`${naam} is ${age} years old`",
			"Ankan is 25 years old",
		},
		// Number interpolation
		{
			"`The answer is ${42}`",
			"The answer is 42",
		},
		// Boolean interpolation
		{
			"`Is it true? ${sotti}`",
			"Is it true? sotti",
		},
		// Empty template
		{
			"``",
			"",
		},
		// Text only (no interpolation)
		{
			"`Hello World`",
			"Hello World",
		},
		// Interpolation at start
		{
			"dhoro name = \"Ankan\";\n`${name} is here`",
			"Ankan is here",
		},
		// Interpolation at end
		{
			"dhoro city = \"Dhaka\";\n`I live in ${city}`",
			"I live in Dhaka",
		},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testStringObjectHelper(t, evaluated, tt.expected)
	}
}

func TestTemplateLiteralExpressions(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		// Arithmetic in interpolation
		{
			"`5 + 3 = ${5 + 3}`",
			"5 + 3 = 8",
		},
		// Function call in interpolation
		{
			"kaj add(a, b) { ferao a + b; }\n`add(10, 20) = ${add(10, 20)}`",
			"add(10, 20) = 30",
		},
		// Array access in interpolation
		{
			"dhoro arr = [1, 2, 3];\n`First element: ${arr[0]}`",
			"First element: 1",
		},
		// Object access in interpolation
		{
			"dhoro obj = {\"name\": \"Ankan\"};\n`Name: ${obj[\"name\"]}`",
			"Name: Ankan",
		},
		// Conditional in interpolation (using function instead of if statement)
		{
			"kaj check(val) { jodi (val > 5) { ferao \"big\"; } nahole { ferao \"small\"; } }\ndhoro x = 10;\n`x is ${check(x)}`",
			"x is big",
		},
		// Multiple operations in interpolation
		{
			"dhoro a = 10;\ndhoro b = 5;\n`${a} + ${b} = ${a + b}`",
			"10 + 5 = 15",
		},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testStringObjectHelper(t, evaluated, tt.expected)
	}
}

func TestTemplateLiteralWithVariables(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		// Single variable
		{
			"dhoro x = \"hello\";\n`Word: ${x}`",
			"Word: hello",
		},
		// Multiple different variables
		{
			"dhoro first = \"John\";\ndhoro last = \"Doe\";\ndhoro age = 30;\n`${first} ${last} is ${age}`",
			"John Doe is 30",
		},
		// Repeated variable
		{
			"dhoro msg = \"repeat\";\n`${msg} ${msg} ${msg}`",
			"repeat repeat repeat",
		},
		// Variable in complex expression
		{
			"dhoro n = 5;\n`Double: ${n * 2}`",
			"Double: 10",
		},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testStringObjectHelper(t, evaluated, tt.expected)
	}
}

func TestTemplateLiteralNumbers(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		// Integer
		{
			"`Number: ${42}`",
			"Number: 42",
		},
		// Float
		{
			"`Float: ${3.14}`",
			"Float: 3.14",
		},
		// Negative number
		{
			"`Negative: ${-5}`",
			"Negative: -5",
		},
		// Zero
		{
			"`Zero: ${0}`",
			"Zero: 0",
		},
		// Arithmetic result
		{
			"`Result: ${2 * 3 + 4}`",
			"Result: 10",
		},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testStringObjectHelper(t, evaluated, tt.expected)
	}
}

func TestTemplateLiteralSpecialChars(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		// Dollar sign (when not followed by brace)
		{
			"dhoro val = 100;\n`Price: $${val}`",
			"Price: $100",
		},
		// Braces in template
		{
			"dhoro val = 5;\n`Count: {${val}}`",
			"Count: {5}",
		},
		// Unicode characters
		{
			"`Bengali: বাংলা`",
			"Bengali: বাংলা",
		},
		// Newline in template (if supported)
		{
			"`Line 1\nLine 2`",
			"Line 1\nLine 2",
		},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testStringObjectHelper(t, evaluated, tt.expected)
	}
}

func TestTemplateLiteralWithFunctions(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		// Simple function call
		{
			"kaj greet(name) { ferao `Hello ${name}!`; }\ngreet(\"World\")",
			"Hello World!",
		},
		// Function with arithmetic
		{
			"kaj times(n) { ferao `${n} * 2 = ${n * 2}`; }\ntimes(7)",
			"7 * 2 = 14",
		},
		// Array operations in template
		{
			"dhoro arr = [1, 2, 3, 4, 5];\n`Array length: ${dorghyo(arr)}`",
			"Array length: 5",
		},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testStringObjectHelper(t, evaluated, tt.expected)
	}
}

func TestTemplateLiteralEdgeCases(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		// Empty interpolation (should work)
		{
			"`Start ${\"\"} End`",
			"Start  End",
		},
		// Null value
		{
			"`Value: ${khali}`",
			"Value: ",
		},
		// Consecutive interpolations
		{
			"`${1}${2}${3}`",
			"123",
		},
		// Interpolation with spaces
		{
			"dhoro x = 42;\n` ${x} `",
			" 42 ",
		},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testStringObjectHelper(t, evaluated, tt.expected)
	}
}

func TestTemplateLiteralErrors(t *testing.T) {
	tests := []struct {
		input           string
		expectedErrType string
	}{
		// Undefined variable in interpolation
		{
			"`Value: ${undefined_var}`",
			"not defined",
		},
		// Invalid expression
		{
			"`Value: ${1 + }`",
			"",
		},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		if _, ok := evaluated.(*object.Error); !ok {
			t.Errorf("expected error for input: %s, got %T (%+v)", tt.input, evaluated, evaluated)
		}
	}
}

func TestTemplateLiteralComplexScenarios(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		// Template in array
		{
			"dhoro person = \"Alice\";\ndhoro arr = [`Hello ${person}`, \"world\"];\narr[0]",
			"Hello Alice",
		},
		// Template in object
		{
			"dhoro name = \"Bob\";\ndhoro obj = {\"greeting\": `Hi ${name}`};\nobj[\"greeting\"]",
			"Hi Bob",
		},
		// Nested template usage
		{
			"dhoro a = \"A\";\ndhoro b = \"B\";\ndhoro result = `${a}${b}`;\nresult",
			"AB",
		},
		// Template with all data types
		{
			"dhoro num = 42;\ndhoro str = \"text\";\ndhoro bool = sotti;\n`${num}, ${str}, ${bool}`",
			"42, text, sotti",
		},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testStringObjectHelper(t, evaluated, tt.expected)
	}
}
