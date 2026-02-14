package test

import (
	"BanglaCode/src/object"
	"testing"
	"time"
)

func TestProcessGhum(t *testing.T) {
	start := time.Now()

	// Sleep for 100ms
	evaluated := testEvalBuiltin(`process_ghum(100)`)

	elapsed := time.Since(start)

	// Should return NULL
	if evaluated.Type() != "NULL" {
		t.Errorf("Expected NULL, got %s", evaluated.Type())
	}

	// Should have slept for at least 100ms
	if elapsed < 100*time.Millisecond {
		t.Errorf("Sleep duration too short: %v", elapsed)
	}
}

func TestProcessId(t *testing.T) {
	evaluated := testEvalBuiltin(`process_id()`)

	// Should return a number
	if evaluated.Type() != "NUMBER" {
		t.Errorf("Expected NUMBER, got %s", evaluated.Type())
		return
	}

	// PID should be positive
	num := evaluated.(*object.Number)
	if num.Value <= 0 {
		t.Errorf("Expected positive PID, got %f", num.Value)
	}
}

func TestProcessParentId(t *testing.T) {
	evaluated := testEvalBuiltin(`process_parent_id()`)

	// Should return a number
	if evaluated.Type() != "NUMBER" {
		t.Errorf("Expected NUMBER, got %s", evaluated.Type())
		return
	}

	// PPID should be positive
	num := evaluated.(*object.Number)
	if num.Value <= 0 {
		t.Errorf("Expected positive PPID, got %f", num.Value)
	}
}

func TestProcessArgs(t *testing.T) {
	evaluated := testEvalBuiltin(`process_args()`)

	// Should return an array
	if evaluated.Type() != "ARRAY" {
		t.Errorf("Expected ARRAY, got %s", evaluated.Type())
		return
	}

	// Should have at least one argument (the program name)
	arr := evaluated.(*object.Array)
	if len(arr.Elements) < 1 {
		t.Error("Expected at least 1 argument")
	}
}

func TestChalan(t *testing.T) {
	// Test simple command
	evaluated := testEvalBuiltin(`chalan("echo", ["hello"])`)

	// Should return a map
	if evaluated.Type() != "MAP" {
		t.Errorf("Expected MAP, got %s", evaluated.Type())
		return
	}

	m := evaluated.(*object.Map)

	// Check that output contains "hello"
	if output, ok := m.Pairs["output"]; ok {
		if str, ok := output.(*object.String); ok {
			if str.Value != "hello\n" && str.Value != "hello" {
				t.Errorf("Expected 'hello', got '%s'", str.Value)
			}
		}
	} else {
		t.Error("Output key not found in result")
	}

	// Check exit code is 0
	if code, ok := m.Pairs["code"]; ok {
		if num, ok := code.(*object.Number); ok {
			if num.Value != 0 {
				t.Errorf("Expected exit code 0, got %f", num.Value)
			}
		}
	}
}

func TestKajDirectory(t *testing.T) {
	evaluated := testEvalBuiltin(`kaj_directory()`)

	// Should return a string (current working directory)
	if evaluated.Type() != "STRING" {
		t.Errorf("Expected STRING, got %s", evaluated.Type())
		return
	}

	str := evaluated.(*object.String)
	if str.Value == "" {
		t.Error("Working directory should not be empty")
	}
}
