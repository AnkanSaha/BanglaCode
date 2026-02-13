package test

import (
	"BanglaCode/src/object"
	"testing"
)

func TestNumberObject(t *testing.T) {
	num := &object.Number{Value: 42.5}

	if num.Type() != object.NUMBER_OBJ {
		t.Errorf("num.Type() wrong. got=%s", num.Type())
	}

	if num.Inspect() != "42.5" {
		t.Errorf("num.Inspect() wrong. got=%s", num.Inspect())
	}
}

func TestNumberIntegerObject(t *testing.T) {
	num := &object.Number{Value: 100}

	if num.Type() != object.NUMBER_OBJ {
		t.Errorf("num.Type() wrong. got=%s", num.Type())
	}

	if num.Inspect() != "100" {
		t.Errorf("num.Inspect() wrong. got=%s", num.Inspect())
	}
}

func TestStringObject(t *testing.T) {
	str := &object.String{Value: "hello world"}

	if str.Type() != object.STRING_OBJ {
		t.Errorf("str.Type() wrong. got=%s", str.Type())
	}

	if str.Inspect() != "hello world" {
		t.Errorf("str.Inspect() wrong. got=%s", str.Inspect())
	}
}

func TestBooleanObject(t *testing.T) {
	tests := []struct {
		value    bool
		expected string
	}{
		{true, "true"},
		{false, "false"},
	}

	for _, tt := range tests {
		b := &object.Boolean{Value: tt.value}

		if b.Type() != object.BOOLEAN_OBJ {
			t.Errorf("b.Type() wrong. got=%s", b.Type())
		}

		if b.Inspect() != tt.expected {
			t.Errorf("b.Inspect() wrong. got=%s", b.Inspect())
		}
	}
}

func TestNullObject(t *testing.T) {
	null := &object.Null{}

	if null.Type() != object.NULL_OBJ {
		t.Errorf("null.Type() wrong. got=%s", null.Type())
	}

	if null.Inspect() != "khali" {
		t.Errorf("null.Inspect() wrong. got=%s", null.Inspect())
	}
}

func TestReturnValueObject(t *testing.T) {
	rv := &object.ReturnValue{Value: &object.Number{Value: 5}}

	if rv.Type() != object.RETURN_OBJ {
		t.Errorf("rv.Type() wrong. got=%s", rv.Type())
	}

	if rv.Inspect() != "5" {
		t.Errorf("rv.Inspect() wrong. got=%s", rv.Inspect())
	}
}

func TestErrorObject(t *testing.T) {
	err := &object.Error{Message: "something went wrong", Line: 10, Column: 5}

	if err.Type() != object.ERROR_OBJ {
		t.Errorf("err.Type() wrong. got=%s", err.Type())
	}

	expected := "Error [line 10, col 5]: something went wrong"
	if err.Inspect() != expected {
		t.Errorf("err.Inspect() wrong. got=%s, expected=%s", err.Inspect(), expected)
	}
}

func TestErrorObjectWithoutPosition(t *testing.T) {
	err := &object.Error{Message: "something went wrong"}

	expected := "Error: something went wrong"
	if err.Inspect() != expected {
		t.Errorf("err.Inspect() wrong. got=%s, expected=%s", err.Inspect(), expected)
	}
}

func TestArrayObject(t *testing.T) {
	arr := &object.Array{
		Elements: []object.Object{
			&object.Number{Value: 1},
			&object.Number{Value: 2},
			&object.String{Value: "hello"},
		},
	}

	if arr.Type() != object.ARRAY_OBJ {
		t.Errorf("arr.Type() wrong. got=%s", arr.Type())
	}

	expected := "[1, 2, hello]"
	if arr.Inspect() != expected {
		t.Errorf("arr.Inspect() wrong. got=%s, expected=%s", arr.Inspect(), expected)
	}
}

func TestEmptyArrayObject(t *testing.T) {
	arr := &object.Array{Elements: []object.Object{}}

	if arr.Inspect() != "[]" {
		t.Errorf("arr.Inspect() wrong. got=%s", arr.Inspect())
	}
}

func TestMapObject(t *testing.T) {
	m := &object.Map{
		Pairs: map[string]object.Object{
			"name": &object.String{Value: "Ankan"},
		},
	}

	if m.Type() != object.MAP_OBJ {
		t.Errorf("m.Type() wrong. got=%s", m.Type())
	}

	// Note: Map iteration order is not guaranteed, so we check presence
	inspect := m.Inspect()
	if inspect != "{name: Ankan}" {
		t.Errorf("m.Inspect() wrong. got=%s", inspect)
	}
}

func TestClassObject(t *testing.T) {
	class := &object.Class{
		Name:    "Person",
		Methods: make(map[string]*object.Function),
	}

	if class.Type() != object.CLASS_OBJ {
		t.Errorf("class.Type() wrong. got=%s", class.Type())
	}

	if class.Inspect() != "sreni Person" {
		t.Errorf("class.Inspect() wrong. got=%s", class.Inspect())
	}
}

func TestInstanceObject(t *testing.T) {
	class := &object.Class{Name: "Person"}
	instance := &object.Instance{
		Class:      class,
		Properties: make(map[string]object.Object),
	}

	if instance.Type() != object.INSTANCE_OBJ {
		t.Errorf("instance.Type() wrong. got=%s", instance.Type())
	}

	if instance.Inspect() != "Person er udahoron" {
		t.Errorf("instance.Inspect() wrong. got=%s", instance.Inspect())
	}
}

func TestBreakObject(t *testing.T) {
	brk := &object.Break{}

	if brk.Type() != object.BREAK_OBJ {
		t.Errorf("brk.Type() wrong. got=%s", brk.Type())
	}

	if brk.Inspect() != "thamo" {
		t.Errorf("brk.Inspect() wrong. got=%s", brk.Inspect())
	}
}

func TestContinueObject(t *testing.T) {
	cont := &object.Continue{}

	if cont.Type() != object.CONTINUE_OBJ {
		t.Errorf("cont.Type() wrong. got=%s", cont.Type())
	}

	if cont.Inspect() != "chharo" {
		t.Errorf("cont.Inspect() wrong. got=%s", cont.Inspect())
	}
}

func TestExceptionObject(t *testing.T) {
	exc := &object.Exception{Message: "error occurred"}

	if exc.Type() != object.EXCEPTION_OBJ {
		t.Errorf("exc.Type() wrong. got=%s", exc.Type())
	}

	if exc.Inspect() != "Exception: error occurred" {
		t.Errorf("exc.Inspect() wrong. got=%s", exc.Inspect())
	}
}

func TestExceptionObjectWithValue(t *testing.T) {
	exc := &object.Exception{
		Value: &object.String{Value: "custom error"},
	}

	if exc.Inspect() != "Exception: custom error" {
		t.Errorf("exc.Inspect() wrong. got=%s", exc.Inspect())
	}
}

func TestModuleObject(t *testing.T) {
	mod := &object.Module{
		Name:    "math",
		Exports: make(map[string]object.Object),
	}

	if mod.Type() != object.MODULE_OBJ {
		t.Errorf("mod.Type() wrong. got=%s", mod.Type())
	}

	if mod.Inspect() != "module math" {
		t.Errorf("mod.Inspect() wrong. got=%s", mod.Inspect())
	}
}

func TestBuiltinObject(t *testing.T) {
	builtin := &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			return object.NULL
		},
	}

	if builtin.Type() != object.BUILTIN_OBJ {
		t.Errorf("builtin.Type() wrong. got=%s", builtin.Type())
	}

	if builtin.Inspect() != "builtin function" {
		t.Errorf("builtin.Inspect() wrong. got=%s", builtin.Inspect())
	}
}

func TestSingletonObjects(t *testing.T) {
	// Test TRUE singleton
	if object.TRUE.Value != true {
		t.Errorf("TRUE.Value wrong. got=%t", object.TRUE.Value)
	}

	// Test FALSE singleton
	if object.FALSE.Value != false {
		t.Errorf("FALSE.Value wrong. got=%t", object.FALSE.Value)
	}

	// Test NULL singleton
	if object.NULL.Type() != object.NULL_OBJ {
		t.Errorf("NULL.Type() wrong. got=%s", object.NULL.Type())
	}

	// Test BREAK singleton
	if object.BREAK.Type() != object.BREAK_OBJ {
		t.Errorf("BREAK.Type() wrong. got=%s", object.BREAK.Type())
	}

	// Test CONTINUE singleton
	if object.CONTINUE.Type() != object.CONTINUE_OBJ {
		t.Errorf("CONTINUE.Type() wrong. got=%s", object.CONTINUE.Type())
	}
}

func TestNativeBoolToBooleanObject(t *testing.T) {
	tests := []struct {
		input    bool
		expected *object.Boolean
	}{
		{true, object.TRUE},
		{false, object.FALSE},
	}

	for _, tt := range tests {
		result := object.NativeBoolToBooleanObject(tt.input)
		if result != tt.expected {
			t.Errorf("NativeBoolToBooleanObject(%t) wrong. got=%v, expected=%v",
				tt.input, result, tt.expected)
		}
	}
}

// Environment tests

func TestEnvironment_NewEnvironment(t *testing.T) {
	env := object.NewEnvironment()
	if env == nil {
		t.Fatal("NewEnvironment() returned nil")
	}
}

func TestEnvironment_SetAndGet(t *testing.T) {
	env := object.NewEnvironment()

	// Set a value
	value := &object.Number{Value: 42}
	env.Set("x", value)

	// Get the value
	got, ok := env.Get("x")
	if !ok {
		t.Fatal("Get('x') returned not ok")
	}

	if got != value {
		t.Errorf("Get('x') wrong. got=%v, expected=%v", got, value)
	}
}

func TestEnvironment_GetNonExistent(t *testing.T) {
	env := object.NewEnvironment()

	_, ok := env.Get("nonexistent")
	if ok {
		t.Error("Get('nonexistent') should return not ok")
	}
}

func TestEnvironment_EnclosedEnvironment(t *testing.T) {
	outer := object.NewEnvironment()
	inner := object.NewEnclosedEnvironment(outer)

	// Set value in outer scope
	outer.Set("x", &object.Number{Value: 10})

	// Get value from inner scope (should find in outer)
	got, ok := inner.Get("x")
	if !ok {
		t.Fatal("Get('x') from inner should find in outer")
	}

	num, ok := got.(*object.Number)
	if !ok || num.Value != 10 {
		t.Errorf("Expected 10, got %v", got)
	}
}

func TestEnvironment_InnerShadowsOuter(t *testing.T) {
	outer := object.NewEnvironment()
	inner := object.NewEnclosedEnvironment(outer)

	outer.Set("x", &object.Number{Value: 10})
	inner.Set("x", &object.Number{Value: 20})

	// Inner should return its own value
	got, _ := inner.Get("x")
	num := got.(*object.Number)
	if num.Value != 20 {
		t.Errorf("Expected inner value 20, got %f", num.Value)
	}

	// Outer should still have its own value
	got, _ = outer.Get("x")
	num = got.(*object.Number)
	if num.Value != 10 {
		t.Errorf("Expected outer value 10, got %f", num.Value)
	}
}

func TestEnvironment_Update(t *testing.T) {
	env := object.NewEnvironment()

	env.Set("x", &object.Number{Value: 10})
	env.Update("x", &object.Number{Value: 20})

	got, _ := env.Get("x")
	num := got.(*object.Number)
	if num.Value != 20 {
		t.Errorf("Expected updated value 20, got %f", num.Value)
	}
}

func TestEnvironment_UpdateInOuter(t *testing.T) {
	outer := object.NewEnvironment()
	inner := object.NewEnclosedEnvironment(outer)

	outer.Set("x", &object.Number{Value: 10})

	// Update from inner should update outer
	inner.Update("x", &object.Number{Value: 30})

	got, _ := outer.Get("x")
	num := got.(*object.Number)
	if num.Value != 30 {
		t.Errorf("Expected updated outer value 30, got %f", num.Value)
	}
}

func TestEnvironment_UpdateNonExistent(t *testing.T) {
	env := object.NewEnvironment()

	// Update non-existent should create it
	env.Update("y", &object.Number{Value: 50})

	got, ok := env.Get("y")
	if !ok {
		t.Fatal("Get('y') should exist after Update")
	}

	num := got.(*object.Number)
	if num.Value != 50 {
		t.Errorf("Expected 50, got %f", num.Value)
	}
}

func TestEnvironment_All(t *testing.T) {
	env := object.NewEnvironment()

	env.Set("x", &object.Number{Value: 1})
	env.Set("y", &object.Number{Value: 2})
	env.Set("z", &object.Number{Value: 3})

	all := env.All()
	if len(all) != 3 {
		t.Errorf("All() should return 3 items, got %d", len(all))
	}

	if _, ok := all["x"]; !ok {
		t.Error("All() should contain 'x'")
	}
	if _, ok := all["y"]; !ok {
		t.Error("All() should contain 'y'")
	}
	if _, ok := all["z"]; !ok {
		t.Error("All() should contain 'z'")
	}
}

func TestEnvironment_AllDoesNotIncludeOuter(t *testing.T) {
	outer := object.NewEnvironment()
	inner := object.NewEnclosedEnvironment(outer)

	outer.Set("x", &object.Number{Value: 1})
	inner.Set("y", &object.Number{Value: 2})

	all := inner.All()
	if len(all) != 1 {
		t.Errorf("All() should return 1 item from current scope, got %d", len(all))
	}

	if _, ok := all["y"]; !ok {
		t.Error("All() should contain 'y' from current scope")
	}

	if _, ok := all["x"]; ok {
		t.Error("All() should not contain 'x' from outer scope")
	}
}

func TestEnvironment_NestedScopes(t *testing.T) {
	global := object.NewEnvironment()
	local1 := object.NewEnclosedEnvironment(global)
	local2 := object.NewEnclosedEnvironment(local1)

	global.Set("global", &object.String{Value: "global"})
	local1.Set("local1", &object.String{Value: "local1"})
	local2.Set("local2", &object.String{Value: "local2"})

	// local2 should see all variables
	_, ok := local2.Get("global")
	if !ok {
		t.Error("local2 should see 'global'")
	}

	_, ok = local2.Get("local1")
	if !ok {
		t.Error("local2 should see 'local1'")
	}

	_, ok = local2.Get("local2")
	if !ok {
		t.Error("local2 should see 'local2'")
	}

	// local1 should not see local2
	_, ok = local1.Get("local2")
	if ok {
		t.Error("local1 should not see 'local2'")
	}

	// global should not see local variables
	_, ok = global.Get("local1")
	if ok {
		t.Error("global should not see 'local1'")
	}
}
