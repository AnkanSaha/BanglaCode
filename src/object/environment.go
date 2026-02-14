package object

// Environment represents a scope for variable bindings
type Environment struct {
	store     map[string]Object
	constants map[string]bool // tracks which variables are constants
	outer     *Environment    // parent scope
	global    *Environment    // reference to global (root) environment
}

// NewEnvironment creates a new environment
func NewEnvironment() *Environment {
	s := make(map[string]Object)
	c := make(map[string]bool)
	env := &Environment{store: s, constants: c, outer: nil, global: nil}
	env.global = env // root environment is its own global
	return env
}

// NewEnclosedEnvironment creates a new environment with an outer scope
func NewEnclosedEnvironment(outer *Environment) *Environment {
	env := &Environment{
		store:     make(map[string]Object),
		constants: make(map[string]bool),
		outer:     outer,
		global:    outer.GetGlobal(),
	}
	return env
}

// GetGlobal returns the global (root) environment
func (e *Environment) GetGlobal() *Environment {
	if e.global != nil {
		return e.global
	}
	return e
}

// Get retrieves a variable from the environment
func (e *Environment) Get(name string) (Object, bool) {
	obj, ok := e.store[name]
	if !ok && e.outer != nil {
		obj, ok = e.outer.Get(name)
	}
	return obj, ok
}

// Set assigns a variable in the environment
func (e *Environment) Set(name string, val Object) Object {
	e.store[name] = val
	return val
}

// SetConstant assigns a constant in the environment
func (e *Environment) SetConstant(name string, val Object) Object {
	e.store[name] = val
	e.constants[name] = true
	return val
}

// SetGlobal assigns a variable in the global environment
func (e *Environment) SetGlobal(name string, val Object) Object {
	global := e.GetGlobal()
	global.store[name] = val
	return val
}

// IsConstant checks if a variable is a constant
func (e *Environment) IsConstant(name string) bool {
	if constant, ok := e.constants[name]; ok && constant {
		return true
	}
	if e.outer != nil {
		return e.outer.IsConstant(name)
	}
	return false
}

// Update updates a variable in the environment (searches outer scopes)
func (e *Environment) Update(name string, val Object) Object {
	_, ok := e.store[name]
	if ok {
		e.store[name] = val
		return val
	}
	if e.outer != nil {
		return e.outer.Update(name, val)
	}
	// If variable doesn't exist, create it in current scope
	e.store[name] = val
	return val
}

// All returns all variables in the current scope (not including outer scopes)
func (e *Environment) All() map[string]Object {
	return e.store
}
