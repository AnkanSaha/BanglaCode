package test

import (
	"BanglaCode/src/object"
	"testing"
)

// ==================== Object.maan (values) Tests ====================

func TestObjectMaan(t *testing.T) {
	tests := []struct {
		input           string
		expectedLength  int
		expectedContent string // Semicolon-separated values for verification
	}{
		// Basic object values
		{
			`dhoro obj = {"a": 1, "b": 2, "c": 3};
			 maan(obj)`,
			3,
			"1;2;3",
		},
		// Empty object
		{
			`dhoro obj = {};
			 maan(obj)`,
			0,
			"",
		},
		// Single key-value pair
		{
			`dhoro obj = {"x": 42};
			 maan(obj)`,
			1,
			"42",
		},
		// String values
		{
			`dhoro obj = {"naam": "Ankan", "city": "Dhaka"};
			 maan(obj)`,
			2,
			"",
		},
		// Mixed types
		{
			`dhoro obj = {"num": 10, "str": "hello", "bool": sotti};
			 maan(obj)`,
			3,
			"",
		},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		arr, ok := evaluated.(*object.Array)
		if !ok {
			t.Errorf("expected Array, got %T (%+v)", evaluated, evaluated)
			continue
		}

		if len(arr.Elements) != tt.expectedLength {
			t.Errorf("wrong array length. expected=%d, got=%d",
				tt.expectedLength, len(arr.Elements))
		}
	}
}

func TestObjectMaanErrors(t *testing.T) {
	tests := []struct {
		input           string
		expectedErrType string
	}{
		// Non-object argument
		{
			`maan([1, 2, 3])`,
			"must be MAP",
		},
		// Non-object argument (number)
		{
			`maan(42)`,
			"must be MAP",
		},
		// Wrong number of arguments
		{
			`maan()`,
			"wrong number of arguments",
		},
		// Too many arguments
		{
			`maan({"a": 1}, {"b": 2})`,
			"wrong number of arguments",
		},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		if _, ok := evaluated.(*object.Error); !ok {
			t.Errorf("expected error, got %T (%+v)", evaluated, evaluated)
		}
	}
}

// ==================== Object.jora (entries) Tests ====================

func TestObjectJora(t *testing.T) {
	tests := []struct {
		input          string
		expectedLength int
	}{
		// Basic object entries
		{
			`dhoro obj = {"a": 1, "b": 2};
			 jora(obj)`,
			2,
		},
		// Empty object
		{
			`dhoro obj = {};
			 jora(obj)`,
			0,
		},
		// Single entry
		{
			`dhoro obj = {"x": 42};
			 jora(obj)`,
			1,
		},
		// Multiple entries
		{
			`dhoro obj = {"naam": "Ankan", "age": 25, "city": "Dhaka"};
			 jora(obj)`,
			3,
		},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		arr, ok := evaluated.(*object.Array)
		if !ok {
			t.Errorf("expected Array, got %T (%+v)", evaluated, evaluated)
			continue
		}

		if len(arr.Elements) != tt.expectedLength {
			t.Errorf("wrong array length. expected=%d, got=%d",
				tt.expectedLength, len(arr.Elements))
			continue
		}

		// Each entry should be an array [key, value]
		for i, entry := range arr.Elements {
			entryArr, ok := entry.(*object.Array)
			if !ok {
				t.Errorf("entry[%d] is not Array. got=%T", i, entry)
				continue
			}
			if len(entryArr.Elements) != 2 {
				t.Errorf("entry[%d] should have 2 elements. got=%d",
					i, len(entryArr.Elements))
			}
		}
	}
}

func TestObjectJoraErrors(t *testing.T) {
	tests := []struct {
		input           string
		expectedErrType string
	}{
		// Non-object argument
		{
			`jora([1, 2, 3])`,
			"must be MAP",
		},
		// Non-object argument (string)
		{
			`jora("string")`,
			"must be MAP",
		},
		// Wrong number of arguments
		{
			`jora()`,
			"wrong number of arguments",
		},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		if _, ok := evaluated.(*object.Error); !ok {
			t.Errorf("expected error, got %T (%+v)", evaluated, evaluated)
		}
	}
}

// ==================== Object.mishra (assign) Tests ====================

func TestObjectMishra(t *testing.T) {
	tests := []struct {
		input              string
		expectedKeyCount   int
		shouldContainValue string
	}{
		// Basic assign (merge)
		{
			`dhoro target = {"a": 1};
			 dhoro source = {"b": 2};
			 mishra(target, source);
			 target`,
			2,
			"",
		},
		// Merge multiple sources
		{
			`dhoro target = {"a": 1};
			 dhoro source1 = {"b": 2};
			 dhoro source2 = {"c": 3};
			 mishra(target, source1, source2);
			 target`,
			3,
			"",
		},
		// Overwrite existing key
		{
			`dhoro target = {"a": 1, "b": 2};
			 dhoro source = {"b": 20};
			 mishra(target, source);
			 target`,
			2,
			"",
		},
		// Empty source
		{
			`dhoro target = {"a": 1};
			 dhoro source = {};
			 mishra(target, source);
			 target`,
			1,
			"",
		},
		// Merge into empty object
		{
			`dhoro target = {};
			 dhoro source = {"a": 1, "b": 2};
			 mishra(target, source);
			 target`,
			2,
			"",
		},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		mapObj, ok := evaluated.(*object.Map)
		if !ok {
			t.Errorf("expected Map, got %T (%+v)", evaluated, evaluated)
			continue
		}

		if len(mapObj.Pairs) != tt.expectedKeyCount {
			t.Errorf("wrong map size. expected=%d, got=%d",
				tt.expectedKeyCount, len(mapObj.Pairs))
		}
	}
}

func TestObjectMishraReturnValue(t *testing.T) {
	// mishra returns the modified target object
	tests := []struct {
		input        string
		expectedType string
	}{
		{
			`dhoro target = {"a": 1};
			 mishra(target, {"b": 2})`,
			"MAP",
		},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		mapObj, ok := evaluated.(*object.Map)
		if !ok {
			t.Errorf("expected Map, got %T (%+v)", evaluated, evaluated)
			continue
		}

		// Should return the target object itself
		if mapObj.Type() != object.MAP_OBJ {
			t.Errorf("wrong object type. expected=%s, got=%s",
				object.MAP_OBJ, mapObj.Type())
		}
	}
}

func TestObjectMishraErrors(t *testing.T) {
	tests := []struct {
		input           string
		expectedErrType string
	}{
		// Non-object target
		{
			`mishra([1, 2, 3], {"a": 1})`,
			"must be MAP",
		},
		// Non-object source
		{
			`mishra({"a": 1}, [1, 2, 3])`,
			"must be MAP",
		},
		// Wrong number of arguments
		{
			`mishra({"a": 1})`,
			"wrong number of arguments",
		},
		// No arguments
		{
			`mishra()`,
			"wrong number of arguments",
		},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		if _, ok := evaluated.(*object.Error); !ok {
			t.Errorf("expected error, got %T (%+v)", evaluated, evaluated)
		}
	}
}

// ==================== Object Methods Integration Tests ====================

func TestObjectMethodsIntegration(t *testing.T) {
	tests := []struct {
		input          string
		expectedLength int
	}{
		// Get values and check length
		{
			`dhoro obj = {"a": 1, "b": 2, "c": 3};
			 dhoro values = maan(obj);
			 values`,
			3,
		},
		// Get entries and check length
		{
			`dhoro obj = {"a": 10, "b": 20, "c": 30};
			 dhoro entries = jora(obj);
			 entries`,
			3,
		},
		// Merge then get values
		{
			`dhoro obj1 = {"a": 1};
			 dhoro obj2 = {"b": 2};
			 mishra(obj1, obj2);
			 maan(obj1)`,
			2,
		},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		arr, ok := evaluated.(*object.Array)
		if !ok {
			t.Errorf("expected Array, got %T (%+v)", evaluated, evaluated)
			continue
		}

		if len(arr.Elements) != tt.expectedLength {
			t.Errorf("wrong array length. expected=%d, got=%d",
				tt.expectedLength, len(arr.Elements))
		}
	}
}
