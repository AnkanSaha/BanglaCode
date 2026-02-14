package test

import (
	"BanglaCode/src/object"
	"strings"
	"testing"
)

// TestConstantDeclaration tests sthir (constant) variable declaration
func TestConstantDeclaration(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		// Basic constant declaration
		{"sthir PI = 3.14; PI;", 3.14},
		{"sthir NAAM = \"BanglaCode\"; NAAM;", "BanglaCode"},
		{"sthir SHOTTO = sotti; SHOTTO;", true},
		{"sthir SHUNNO = 0; SHUNNO;", 0.0},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		switch expected := tt.expected.(type) {
		case float64:
			testNumberObject(t, evaluated, expected)
		case string:
			testStringObject(t, evaluated, expected)
		case bool:
			testBooleanObject(t, evaluated, expected)
		}
	}
}

// TestConstantReassignmentError tests that constants cannot be reassigned
func TestConstantReassignmentError(t *testing.T) {
	tests := []struct {
		input           string
		expectedMessage string
	}{
		{
			"sthir PI = 3.14; PI = 3.14159;",
			"sthir (constant)",
		},
		{
			"sthir x = 10; x = 20;",
			"sthir (constant)",
		},
		{
			"sthir naam = \"hello\"; naam = \"world\";",
			"sthir (constant)",
		},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		errObj, ok := evaluated.(*object.Error)
		if !ok {
			t.Errorf("expected error object for input: %s, got=%T (%+v)",
				tt.input, evaluated, evaluated)
			continue
		}

		if !strings.Contains(errObj.Message, tt.expectedMessage) {
			t.Errorf("error message should contain '%s', got='%s'",
				tt.expectedMessage, errObj.Message)
		}
	}
}

// TestGlobalVariableDeclaration tests bishwo (global) variable declaration
func TestGlobalVariableDeclaration(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		// Basic global declaration
		{"bishwo ganok = 0; ganok;", 0.0},
		{"bishwo msg = \"hello\"; msg;", "hello"},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		switch expected := tt.expected.(type) {
		case float64:
			testNumberObject(t, evaluated, expected)
		case string:
			testStringObject(t, evaluated, expected)
		}
	}
}

// TestGlobalVariableAccessFromNestedScope tests global variable access from inner scopes
func TestGlobalVariableAccessFromNestedScope(t *testing.T) {
	input := `
	bishwo ganok = 0;
	
	kaj barao() {
		ganok = ganok + 1;
	}
	
	barao();
	barao();
	barao();
	ganok;
	`

	evaluated := testEval(input)
	testNumberObject(t, evaluated, 3.0)
}

// TestGlobalVariableModificationFromFunction tests modifying global from function
func TestGlobalVariableModificationFromFunction(t *testing.T) {
	input := `
	bishwo result = "";
	
	kaj add_text(text) {
		result = result + text;
	}
	
	add_text("Hello");
	add_text(" ");
	add_text("World");
	result;
	`

	evaluated := testEval(input)
	testStringObject(t, evaluated, "Hello World")
}

// TestLocalVariableShadowing tests that local variables shadow outer scope
func TestLocalVariableShadowing(t *testing.T) {
	input := `
	dhoro x = 10;
	dhoro result = 0;
	
	jodi (sotti) {
		dhoro x = 20;
		result = x;
	}
	
	result;
	`

	evaluated := testEval(input)
	testNumberObject(t, evaluated, 20.0)
}

// TestLocalVariableDoesNotAffectOuter tests that local redeclaration shadows outer scope
// Note: In BanglaCode, if blocks share the same scope (like Python)
// To get true block scoping, use functions
func TestLocalVariableDoesNotAffectOuter(t *testing.T) {
	input := `
	dhoro x = 10;
	
	kaj test() {
		dhoro x = 20;
		ferao x;
	}
	
	dhoro inner = test();
	x;
	`

	evaluated := testEval(input)
	testNumberObject(t, evaluated, 10.0)
}

// TestMixedScopes tests interaction between dhoro, sthir, and bishwo
func TestMixedScopes(t *testing.T) {
	input := `
	sthir PI = 3.14;
	bishwo total = 0;
	dhoro radius = 5;
	
	kaj calculateArea(r) {
		dhoro area = PI * r * r;
		total = total + area;
		ferao area;
	}
	
	dhoro a1 = calculateArea(radius);
	dhoro a2 = calculateArea(10);
	
	total;
	`

	evaluated := testEval(input)
	// total should be PI * 5 * 5 + PI * 10 * 10 = 3.14 * 25 + 3.14 * 100 = 78.5 + 314 = 392.5
	testNumberObject(t, evaluated, 392.5)
}

// TestConstantInFunction tests constant access in function scope
func TestConstantInFunction(t *testing.T) {
	input := `
	sthir MULTIPLIER = 2;
	
	kaj double(x) {
		ferao x * MULTIPLIER;
	}
	
	double(5);
	`

	evaluated := testEval(input)
	testNumberObject(t, evaluated, 10.0)
}

// TestGlobalInNestedFunction tests global access in nested function calls
func TestGlobalInNestedFunction(t *testing.T) {
	input := `
	bishwo counter = 0;
	
	kaj outer() {
		kaj inner() {
			counter = counter + 1;
		}
		inner();
		inner();
	}
	
	outer();
	outer();
	counter;
	`

	evaluated := testEval(input)
	testNumberObject(t, evaluated, 4.0)
}
