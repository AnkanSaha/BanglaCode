package test

import (
	"BanglaCode/src/object"
	"testing"
)

// ==================== Array.manchitro (map) Tests ====================

func TestArrayManchitro(t *testing.T) {
	tests := []struct {
		input    string
		expected []float64
	}{
		// Basic map
		{
			`dhoro arr = [1, 2, 3];
			 dhoro result = manchitro(arr, kaj(x) { ferao x * 2; });
			 result`,
			[]float64{2, 4, 6},
		},
		// Map with index parameter
		{
			`dhoro arr = [10, 20, 30];
			 dhoro result = manchitro(arr, kaj(x, i) { ferao x + i; });
			 result`,
			[]float64{10, 21, 32},
		},
		// Empty array
		{
			`dhoro arr = [];
			 dhoro result = manchitro(arr, kaj(x) { ferao x * 2; });
			 result`,
			[]float64{},
		},
		// Single element
		{
			`dhoro arr = [5];
			 dhoro result = manchitro(arr, kaj(x) { ferao x * 3; });
			 result`,
			[]float64{15},
		},
		// Negative numbers
		{
			`dhoro arr = [-1, -2, -3];
			 dhoro result = manchitro(arr, kaj(x) { ferao x * 2; });
			 result`,
			[]float64{-2, -4, -6},
		},
		// Floats
		{
			`dhoro arr = [1.5, 2.5, 3.5];
			 dhoro result = manchitro(arr, kaj(x) { ferao x + 0.5; });
			 result`,
			[]float64{2.0, 3.0, 4.0},
		},
		// Complex transformation
		{
			`dhoro arr = [1, 2, 3, 4];
			 dhoro result = manchitro(arr, kaj(x, i) { ferao x * (i + 1); });
			 result`,
			[]float64{1, 4, 9, 16},
		},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testArrayObject(t, evaluated, tt.expected, 0)
	}
}

func TestArrayManchitroErrors(t *testing.T) {
	tests := []struct {
		input           string
		expectedErrType string
	}{
		// Wrong number of arguments
		{
			`manchitro([1, 2, 3])`,
			"wrong number of arguments",
		},
		// Non-array argument
		{
			`manchitro(5, kaj(x) { ferao x; })`,
			"must be ARRAY",
		},
		// Non-function callback
		{
			`manchitro([1, 2, 3], 5)`,
			"must be FUNCTION",
		},
		// Error in callback
		{
			`dhoro arr = [1, 2, 3];
			 manchitro(arr, kaj(x) { ferao 1 / 0; })`,
			"ERROR",
		},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		if errObj, ok := evaluated.(*object.Error); !ok {
			t.Errorf("expected error, got %T (%+v)", evaluated, evaluated)
		} else if len(errObj.Message) == 0 {
			t.Errorf("expected error message containing %q", tt.expectedErrType)
		}
	}
}

// ==================== Array.chhanno (filter) Tests ====================

func TestArrayChhanno(t *testing.T) {
	tests := []struct {
		input    string
		expected []float64
	}{
		// Basic filter
		{
			`dhoro arr = [1, 2, 3, 4, 5];
			 dhoro result = chhanno(arr, kaj(x) { ferao x > 2; });
			 result`,
			[]float64{3, 4, 5},
		},
		// Filter with index
		{
			`dhoro arr = [10, 20, 30, 40];
			 dhoro result = chhanno(arr, kaj(x, i) { ferao i < 2; });
			 result`,
			[]float64{10, 20},
		},
		// Empty array
		{
			`dhoro arr = [];
			 dhoro result = chhanno(arr, kaj(x) { ferao x > 0; });
			 result`,
			[]float64{},
		},
		// All match
		{
			`dhoro arr = [1, 2, 3];
			 dhoro result = chhanno(arr, kaj(x) { ferao sotti; });
			 result`,
			[]float64{1, 2, 3},
		},
		// None match
		{
			`dhoro arr = [1, 2, 3];
			 dhoro result = chhanno(arr, kaj(x) { ferao mittha; });
			 result`,
			[]float64{},
		},
		// Filter even numbers
		{
			`dhoro arr = [1, 2, 3, 4, 5, 6];
			 dhoro result = chhanno(arr, kaj(x) { ferao x % 2 == 0; });
			 result`,
			[]float64{2, 4, 6},
		},
		// Filter with complex condition
		{
			`dhoro arr = [1, 2, 3, 4, 5];
			 dhoro result = chhanno(arr, kaj(x, i) { ferao x > 1 ebong i < 3; });
			 result`,
			[]float64{2, 3},
		},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testArrayObject(t, evaluated, tt.expected, 0)
	}
}

func TestArrayChhannoErrors(t *testing.T) {
	tests := []struct {
		input           string
		expectedErrType string
	}{
		// Wrong number of arguments
		{
			`chhanno([1, 2, 3])`,
			"wrong number of arguments",
		},
		// Non-array argument
		{
			`chhanno("string", kaj(x) { ferao x; })`,
			"must be ARRAY",
		},
		// Non-function callback
		{
			`chhanno([1, 2, 3], 5)`,
			"must be FUNCTION",
		},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		if errObj, ok := evaluated.(*object.Error); !ok {
			t.Errorf("expected error, got %T (%+v)", evaluated, evaluated)
		} else if len(errObj.Message) == 0 {
			t.Errorf("expected error message containing %q", tt.expectedErrType)
		}
	}
}

// ==================== Array.sonkuchito (reduce) Tests ====================

func TestArraySonkuchito(t *testing.T) {
	tests := []struct {
		input    string
		expected float64
	}{
		// Basic reduce (sum)
		{
			`dhoro arr = [1, 2, 3, 4, 5];
			 sonkuchito(arr, kaj(acc, x) { ferao acc + x; }, 0)`,
			15,
		},
		// Reduce to product
		{
			`dhoro arr = [1, 2, 3, 4];
			 sonkuchito(arr, kaj(acc, x) { ferao acc * x; }, 1)`,
			24,
		},
		// Reduce without initial value (should use first element)
		{
			`dhoro arr = [1, 2, 3, 4];
			 sonkuchito(arr, kaj(acc, x) { ferao acc + x; })`,
			10,
		},
		// Single element array
		{
			`dhoro arr = [42];
			 sonkuchito(arr, kaj(acc, x) { ferao acc + x; }, 0)`,
			42,
		},
		// Reduce with initial value 10
		{
			`dhoro arr = [1, 2, 3];
			 sonkuchito(arr, kaj(acc, x) { ferao acc + x; }, 10)`,
			16,
		},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testNumberObject(t, evaluated, tt.expected)
	}
}

func TestArraySonkuchitoErrors(t *testing.T) {
	tests := []struct {
		input           string
		expectedErrType string
	}{
		// Empty array without initial value
		{
			`dhoro arr = [];
			 sonkuchito(arr, kaj(acc, x) { ferao acc + x; })`,
			"ERROR",
		},
		// Non-array argument
		{
			`sonkuchito(5, kaj(acc, x) { ferao acc + x; }, 0)`,
			"must be ARRAY",
		},
		// Non-function callback
		{
			`sonkuchito([1, 2, 3], 5, 0)`,
			"must be FUNCTION",
		},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		if _, ok := evaluated.(*object.Error); !ok {
			t.Errorf("expected error for input: %s, got %T (%+v)", tt.input, evaluated, evaluated)
		}
	}
}

// ==================== Array.proti (forEach) Tests ====================

func TestArrayProti(t *testing.T) {
	tests := []struct {
		input    string
		expected object.Object
	}{
		// Basic forEach with side effects
		{
			`dhoro count = 0;
			 proti([1, 2, 3], kaj(x) { count = count + 1; });
			 count`,
			&object.Number{Value: 3},
		},
		// forEach with empty array
		{
			`dhoro count = 0;
			 proti([], kaj(x) { count = count + 1; });
			 count`,
			&object.Number{Value: 0},
		},
		// forEach with index parameter
		{
			`dhoro sum = 0;
			 proti([10, 20, 30], kaj(x, i) { sum = sum + i; });
			 sum`,
			&object.Number{Value: 3}, // 0 + 1 + 2
		},
		// forEach returns NULL
		{
			`proti([1, 2, 3], kaj(x) { ferao x; })`,
			object.NULL,
		},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)

		switch expected := tt.expected.(type) {
		case *object.Number:
			testNumberObject(t, evaluated, expected.Value)
		case *object.Null:
			testNullObject(t, evaluated)
		}
	}
}

func TestArrayProtiErrors(t *testing.T) {
	tests := []struct {
		input           string
		expectedErrType string
	}{
		// Wrong number of arguments
		{
			`proti([1, 2, 3])`,
			"wrong number of arguments",
		},
		// Non-array argument
		{
			`proti("string", kaj(x) { ferao x; })`,
			"must be ARRAY",
		},
		// Non-function callback
		{
			`proti([1, 2, 3], 5)`,
			"must be FUNCTION",
		},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		if errObj, ok := evaluated.(*object.Error); !ok {
			t.Errorf("expected error, got %T (%+v)", evaluated, evaluated)
		} else if len(errObj.Message) == 0 {
			t.Errorf("expected error message containing %q", tt.expectedErrType)
		}
	}
}

// ==================== Array Method Chaining Tests ====================

func TestArrayMethodChaining(t *testing.T) {
	tests := []struct {
		input    string
		expected []float64
	}{
		// Map then filter
		{
			`dhoro arr = [1, 2, 3, 4];
			 dhoro doubled = manchitro(arr, kaj(x) { ferao x * 2; });
			 dhoro result = chhanno(doubled, kaj(x) { ferao x > 4; });
			 result`,
			[]float64{6, 8},
		},
		// Filter then map
		{
			`dhoro arr = [1, 2, 3, 4, 5];
			 dhoro filtered = chhanno(arr, kaj(x) { ferao x > 2; });
			 dhoro result = manchitro(filtered, kaj(x) { ferao x * 10; });
			 result`,
			[]float64{30, 40, 50},
		},
		// Map, filter, then reduce
		{
			`dhoro arr = [1, 2, 3, 4, 5];
			 dhoro mapped = manchitro(arr, kaj(x) { ferao x * 2; });
			 dhoro filtered = chhanno(mapped, kaj(x) { ferao x > 4; });
			 sonkuchito(filtered, kaj(acc, x) { ferao acc + x; }, 0)`,
			[]float64{30}, // [2,4,6,8,10] -> [6,8,10] -> 24
		},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		if arr, ok := evaluated.(*object.Array); ok {
			expected := tt.expected
			if len(arr.Elements) != len(expected) {
				t.Errorf("wrong array length. expected=%d, got=%d",
					len(expected), len(arr.Elements))
			}
		} else {
			// Last test returns a number
			testNumberObject(t, evaluated, 24)
		}
	}
}
