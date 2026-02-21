package test

import (
	"testing"
)

// ==================== Switch/Case Tests ====================

func TestSwitchStatement(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		// Match first case
		{
			`bikolpo (1) {
				khetre 1 { ferao "one"; }
				khetre 2 { ferao "two"; }
				manchito { ferao "other"; }
			}`,
			"one",
		},
		// Match middle case
		{
			`bikolpo (2) {
				khetre 1 { ferao "one"; }
				khetre 2 { ferao "two"; }
				khetre 3 { ferao "three"; }
				manchito { ferao "other"; }
			}`,
			"two",
		},
		// Match default case
		{
			`bikolpo (10) {
				khetre 1 { ferao "one"; }
				khetre 2 { ferao "two"; }
				manchito { ferao "other"; }
			}`,
			"other",
		},
		// Multiple case values (matching second case)
		{
			`bikolpo (5) {
				khetre 1 { ferao "one"; }
				khetre 2 { ferao "two"; }
				khetre 5 { ferao "five"; }
			}`,
			"five",
		},
		// String matching
		{
			`bikolpo ("hello") {
				khetre "hello" { ferao "found"; }
				khetre "world" { ferao "not found"; }
				manchito { ferao "default"; }
			}`,
			"found",
		},
		// Boolean matching
		{
			`bikolpo (sotti) {
				khetre sotti { ferao "true"; }
				khetre mittha { ferao "false"; }
				manchito { ferao "other"; }
			}`,
			"true",
		},
		// Multiple statements in case
		{
			`dhoro x = 0;
			 bikolpo (2) {
				khetre 1 { x = 10; }
				khetre 2 { x = 20; x = 25; }
				manchito { x = 100; }
			 }
			 x`,
			25.0,
		},
		// Nested operations in case
		{
			`bikolpo (3) {
				khetre 1 { ferao 10 + 5; }
				khetre 2 { ferao 20 + 5; }
				khetre 3 { ferao 30 + 5; }
				manchito { ferao 100; }
			}`,
			35.0,
		},
		// Break statement (thamo)
		{
			`dhoro x = 0;
			 bikolpo (1) {
				khetre 1 { x = 10; thamo; }
				khetre 2 { x = 20; }
				manchito { x = 100; }
			 }
			 x`,
			10.0,
		},
		// Expression in switch
		{
			`bikolpo (5 + 5) {
				khetre 10 { ferao "ten"; }
				khetre 20 { ferao "twenty"; }
				manchito { ferao "other"; }
			}`,
			"ten",
		},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)

		switch expected := tt.expected.(type) {
		case string:
			testStringObject(t, evaluated, expected)
		case float64:
			testNumberObject(t, evaluated, expected)
		}
	}
}

func TestSwitchWithBreak(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		// Case with explicit break
		{
			`dhoro x = 0;
			 bikolpo (1) {
				khetre 1 { x = 1; thamo; }
				khetre 2 { x = 2; }
			 }
			 x`,
			1.0,
		},
		// No break, continue to next case
		{
			`dhoro x = 0;
			 bikolpo (1) {
				khetre 1 { x = 1; }
				khetre 2 { x = 2; }
				manchito { x = 10; }
			 }
			 x`,
			1.0,
		},
		// Default case with break
		{
			`dhoro x = 0;
			 bikolpo (99) {
				khetre 1 { x = 1; }
				khetre 2 { x = 2; }
				manchito { x = 100; thamo; }
			 }
			 x`,
			100.0,
		},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testNumberObject(t, evaluated, tt.expected.(float64))
	}
}

func TestSwitchWithDifferentTypes(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		// Number matching
		{
			`bikolpo (42) {
				khetre 42 { ferao "answer"; }
				manchito { ferao "unknown"; }
			}`,
			"answer",
		},
		// String matching
		{
			`bikolpo ("BanglaCode") {
				khetre "BanglaCode" { ferao "programming"; }
				khetre "other" { ferao "not programming"; }
				manchito { ferao "default"; }
			}`,
			"programming",
		},
		// Boolean true
		{
			`bikolpo (sotti) {
				khetre sotti { ferao "yes"; }
				manchito { ferao "no"; }
			}`,
			"yes",
		},
		// Boolean false
		{
			`bikolpo (mittha) {
				khetre sotti { ferao "yes"; }
				khetre mittha { ferao "no"; }
				manchito { ferao "maybe"; }
			}`,
			"no",
		},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testStringObject(t, evaluated, tt.expected.(string))
	}
}

func TestSwitchWithComplexConditions(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		// Switch on expression result
		{
			`bikolpo (2 + 3) {
				khetre 5 { ferao "five"; }
				manchito { ferao "not five"; }
			}`,
			"five",
		},
		// Switch on variable
		{
			`dhoro x = 20;
			 bikolpo (x / 2) {
				khetre 10 { ferao "ten"; }
				manchito { ferao "not ten"; }
			}`,
			"ten",
		},
		// Switch on function call result
		{
			`kaj getValue() { ferao 3; }
			 bikolpo (getValue()) {
				khetre 1 { ferao "one"; }
				khetre 2 { ferao "two"; }
				khetre 3 { ferao "three"; }
				manchito { ferao "other"; }
			}`,
			"three",
		},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testStringObject(t, evaluated, tt.expected.(string))
	}
}

func TestSwitchEdgeCases(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		// Only default case
		{
			`bikolpo (1) {
				manchito { ferao "only default"; }
			}`,
			"only default",
		},
		// Single case
		{
			`bikolpo (1) {
				khetre 1 { ferao "match"; }
			}`,
			"match",
		},
		// No match, no default
		{
			`bikolpo (99) {
				khetre 1 { ferao "one"; }
				khetre 2 { ferao "two"; }
			}`,
			"",
		},
		// Case with null
		{
			`bikolpo (khali) {
				khetre khali { ferao "null"; }
				manchito { ferao "not null"; }
			}`,
			"null",
		},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		if tt.expected == "" {
			testNullObject(t, evaluated)
		} else {
			testStringObject(t, evaluated, tt.expected.(string))
		}
	}
}

func TestSwitchWithSideEffects(t *testing.T) {
	tests := []struct {
		input    string
		expected float64
	}{
		// Side effects in cases
		{
			`dhoro x = 0;
			 dhoro y = 0;
			 bikolpo (2) {
				khetre 1 { x = 1; }
				khetre 2 { x = 2; y = 20; }
				manchito { x = 100; }
			 }
			 y`,
			20,
		},
		// Modification of array in case
		{
			`dhoro arr = [1, 2, 3];
			 bikolpo (1) {
				khetre 1 { arr = [10, 20]; }
				manchito { arr = [100]; }
			 }
			 dorghyo(arr)`,
			2,
		},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testNumberObject(t, evaluated, tt.expected)
	}
}
