package object

import "sync"

// Environment represents a scope for variable bindings
type Environment struct {
	store     map[string]Object
	constants map[string]bool // tracks which variables are constants
	outer     *Environment    // parent scope
	global    *Environment    // reference to global (root) environment
	mu        sync.RWMutex
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
	e.mu.RLock()
	obj, ok := e.store[name]
	outer := e.outer
	e.mu.RUnlock()
	if ok {
		return obj, true
	}
	if outer != nil {
		return outer.Get(name)
	}
	return nil, false
}

// Set assigns a variable in the environment
func (e *Environment) Set(name string, val Object) Object {
	e.mu.Lock()
	e.store[name] = val
	e.mu.Unlock()
	return val
}

// SetConstant assigns a constant in the environment
func (e *Environment) SetConstant(name string, val Object) Object {
	e.mu.Lock()
	e.store[name] = val
	e.constants[name] = true
	e.mu.Unlock()
	return val
}

// SetGlobal assigns a variable in the global environment
func (e *Environment) SetGlobal(name string, val Object) Object {
	global := e.GetGlobal()
	global.mu.Lock()
	global.store[name] = val
	global.mu.Unlock()
	return val
}

// IsConstant checks if a variable is a constant
func (e *Environment) IsConstant(name string) bool {
	e.mu.RLock()
	constant, ok := e.constants[name]
	outer := e.outer
	e.mu.RUnlock()
	if ok && constant {
		return true
	}
	if outer != nil {
		return outer.IsConstant(name)
	}
	return false
}

// Update updates a variable in the environment (searches outer scopes)
func (e *Environment) Update(name string, val Object) Object {
	e.mu.Lock()
	_, ok := e.store[name]
	if ok {
		e.store[name] = val
		e.mu.Unlock()
		return val
	}
	outer := e.outer
	e.mu.Unlock()
	if outer != nil {
		return outer.Update(name, val)
	}
	e.mu.Lock()
	e.store[name] = val
	e.mu.Unlock()
	return val
}

// All returns all variables in the current scope (not including outer scopes)
func (e *Environment) All() map[string]Object {
	e.mu.RLock()
	defer e.mu.RUnlock()
	out := make(map[string]Object, len(e.store))
	for k, v := range e.store {
		out[k] = v
	}
	return out
}
