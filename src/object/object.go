package object

import (
	"BanglaCode/src/ast"
	"bytes"
	"fmt"
	"strings"
	"sync"
)

// ObjectType represents the type of an object
type ObjectType string

const (
	NUMBER_OBJ        = "NUMBER"
	STRING_OBJ        = "STRING"
	BOOLEAN_OBJ       = "BOOLEAN"
	NULL_OBJ          = "NULL"
	RETURN_OBJ        = "RETURN"
	ERROR_OBJ         = "ERROR"
	FUNCTION_OBJ      = "FUNCTION"
	BUILTIN_OBJ       = "BUILTIN"
	ARRAY_OBJ         = "ARRAY"
	MAP_OBJ           = "MAP"
	CLASS_OBJ         = "CLASS"
	INSTANCE_OBJ      = "INSTANCE"
	BREAK_OBJ         = "BREAK"
	CONTINUE_OBJ      = "CONTINUE"
	EXCEPTION_OBJ     = "EXCEPTION"
	MODULE_OBJ        = "MODULE"
	PROMISE_OBJ       = "PROMISE"
	DB_CONNECTION_OBJ = "DB_CONNECTION"
	DB_RESULT_OBJ     = "DB_RESULT"
	DB_POOL_OBJ       = "DB_POOL"
)

// Object represents any runtime value
type Object interface {
	Type() ObjectType
	Inspect() string
}

// Number represents a numeric value (int or float)
type Number struct {
	Value float64
}

func (n *Number) Type() ObjectType { return NUMBER_OBJ }
func (n *Number) Inspect() string  { return fmt.Sprintf("%g", n.Value) }

// String represents a string value
type String struct {
	Value string
}

func (s *String) Type() ObjectType { return STRING_OBJ }
func (s *String) Inspect() string  { return s.Value }

// Boolean represents a boolean value
type Boolean struct {
	Value bool
}

func (b *Boolean) Type() ObjectType { return BOOLEAN_OBJ }
func (b *Boolean) Inspect() string  { return fmt.Sprintf("%t", b.Value) }

// Null represents a null value (khali)
type Null struct{}

func (n *Null) Type() ObjectType { return NULL_OBJ }
func (n *Null) Inspect() string  { return "khali" }

// ReturnValue wraps a value for return statements
type ReturnValue struct {
	Value Object
}

func (rv *ReturnValue) Type() ObjectType { return RETURN_OBJ }
func (rv *ReturnValue) Inspect() string  { return rv.Value.Inspect() }

// Error represents a runtime error
type Error struct {
	Message string
	Line    int
	Column  int
}

func (e *Error) Type() ObjectType { return ERROR_OBJ }
func (e *Error) Inspect() string {
	if e.Line > 0 {
		return fmt.Sprintf("Error [line %d, col %d]: %s", e.Line, e.Column, e.Message)
	}
	return "Error: " + e.Message
}

// Function represents a user-defined function
type Function struct {
	Parameters    []*ast.Identifier
	RestParameter *ast.Identifier // optional rest parameter (...args)
	Body          *ast.BlockStatement
	Env           *Environment
	Name          string
	IsAsync       bool // true for async functions (proyash kaj)
}

func (f *Function) Type() ObjectType { return FUNCTION_OBJ }
func (f *Function) Inspect() string {
	var out bytes.Buffer
	params := []string{}
	for _, p := range f.Parameters {
		params = append(params, p.String())
	}
	if f.RestParameter != nil {
		params = append(params, "..."+f.RestParameter.String())
	}
	out.WriteString("kaj")
	if f.Name != "" {
		out.WriteString(" " + f.Name)
	}
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") {\n")
	out.WriteString(f.Body.String())
	out.WriteString("\n}")
	return out.String()
}

// BuiltinFunction represents a built-in function
type BuiltinFunction func(args ...Object) Object

// Builtin wraps a built-in function
type Builtin struct {
	Fn BuiltinFunction
}

func (b *Builtin) Type() ObjectType { return BUILTIN_OBJ }
func (b *Builtin) Inspect() string  { return "builtin function" }

// Array represents an array/list
type Array struct {
	Elements []Object
}

func (a *Array) Type() ObjectType { return ARRAY_OBJ }
func (a *Array) Inspect() string {
	var out bytes.Buffer
	elements := []string{}
	for _, e := range a.Elements {
		elements = append(elements, e.Inspect())
	}
	out.WriteString("[")
	out.WriteString(strings.Join(elements, ", "))
	out.WriteString("]")
	return out.String()
}

// Map represents a hash map/object
type Map struct {
	Pairs map[string]Object
}

func (m *Map) Type() ObjectType { return MAP_OBJ }
func (m *Map) Inspect() string {
	var out bytes.Buffer
	pairs := []string{}
	for key, value := range m.Pairs {
		pairs = append(pairs, fmt.Sprintf("%s: %s", key, value.Inspect()))
	}
	out.WriteString("{")
	out.WriteString(strings.Join(pairs, ", "))
	out.WriteString("}")
	return out.String()
}

// Class represents a class definition
type Class struct {
	Name    string
	Methods map[string]*Function
}

func (c *Class) Type() ObjectType { return CLASS_OBJ }
func (c *Class) Inspect() string  { return "sreni " + c.Name }

// Instance represents an instance of a class
type Instance struct {
	Class      *Class
	Properties map[string]Object
}

func (i *Instance) Type() ObjectType { return INSTANCE_OBJ }
func (i *Instance) Inspect() string  { return i.Class.Name + " er udahoron" }

// Break represents a break statement
type Break struct{}

func (b *Break) Type() ObjectType { return BREAK_OBJ }
func (b *Break) Inspect() string  { return "thamo" }

// Continue represents a continue statement
type Continue struct{}

func (c *Continue) Type() ObjectType { return CONTINUE_OBJ }
func (c *Continue) Inspect() string  { return "chharo" }

// Exception represents a thrown exception
type Exception struct {
	Message string
	Value   Object
}

func (e *Exception) Type() ObjectType { return EXCEPTION_OBJ }
func (e *Exception) Inspect() string {
	if e.Value != nil {
		return "Exception: " + e.Value.Inspect()
	}
	return "Exception: " + e.Message
}

// Module represents an imported module
type Module struct {
	Name    string
	Exports map[string]Object
}

func (m *Module) Type() ObjectType { return MODULE_OBJ }
func (m *Module) Inspect() string  { return "module " + m.Name }

// PromiseState represents the state of a promise
type PromiseState string

const (
	PROMISE_PENDING  PromiseState = "PENDING"
	PROMISE_RESOLVED PromiseState = "RESOLVED"
	PROMISE_REJECTED PromiseState = "REJECTED"
)

// Promise represents an async operation that will complete in the future
type Promise struct {
	State      PromiseState
	Value      Object      // resolved value
	Error      Object      // rejection error
	ResultChan chan Object // for goroutine communication
	ErrorChan  chan Object // for error communication
	Mu         sync.RWMutex
}

func (p *Promise) Type() ObjectType { return PROMISE_OBJ }
func (p *Promise) Inspect() string {
	p.Mu.RLock()
	defer p.Mu.RUnlock()
	switch p.State {
	case PROMISE_RESOLVED:
		if p.Value != nil {
			return fmt.Sprintf("Promise(RESOLVED: %s)", p.Value.Inspect())
		}
		return "Promise(RESOLVED)"
	case PROMISE_REJECTED:
		if p.Error != nil {
			return fmt.Sprintf("Promise(REJECTED: %s)", p.Error.Inspect())
		}
		return "Promise(REJECTED)"
	default:
		return "Promise(PENDING)"
	}
}

// CreatePromise creates a new pending promise with channels
func CreatePromise() *Promise {
	return &Promise{
		State:      PROMISE_PENDING,
		ResultChan: make(chan Object, 1),
		ErrorChan:  make(chan Object, 1),
	}
}

// ResolvePromise resolves a promise with a value
func ResolvePromise(promise *Promise, value Object) {
	promise.Mu.Lock()
	promise.State = PROMISE_RESOLVED
	promise.Value = value
	promise.Mu.Unlock()
	promise.ResultChan <- value
}

// RejectPromise rejects a promise with an error
func RejectPromise(promise *Promise, err Object) {
	promise.Mu.Lock()
	promise.State = PROMISE_REJECTED
	promise.Error = err
	promise.Mu.Unlock()
	promise.ErrorChan <- err
}

// DBConnection represents a database connection
type DBConnection struct {
	ID       string            // Unique connection identifier
	DBType   string            // Database type: "postgres", "mysql", "mongodb", "redis"
	Native   interface{}       // Underlying driver connection (*sql.DB, *mongo.Client, *redis.Client)
	PoolID   string            // Reference to connection pool (if pooled)
	Metadata map[string]Object // Connection metadata (host, port, database, etc.)
	Mu       sync.RWMutex      // Thread-safe access
}

func (d *DBConnection) Type() ObjectType { return DB_CONNECTION_OBJ }
func (d *DBConnection) Inspect() string {
	d.Mu.RLock()
	defer d.Mu.RUnlock()
	return fmt.Sprintf("DB_CONNECTION(%s, id=%s)", d.DBType, d.ID)
}

// DBResult represents a database query result
type DBResult struct {
	Rows         []map[string]Object // Result rows as maps (column name -> value)
	RowsAffected int64               // Rows affected by INSERT/UPDATE/DELETE
	LastInsertID int64               // Last inserted ID (for SQL databases)
	Error        *Error              // Query error (if any)
}

func (d *DBResult) Type() ObjectType { return DB_RESULT_OBJ }
func (d *DBResult) Inspect() string {
	if d.Error != nil {
		return fmt.Sprintf("DB_RESULT(error: %s)", d.Error.Message)
	}
	return fmt.Sprintf("DB_RESULT(rows=%d, affected=%d)", len(d.Rows), d.RowsAffected)
}

// DBPool represents a connection pool
type DBPool struct {
	ID          string            // Unique pool identifier
	DBType      string            // Database type
	MaxConns    int               // Maximum connections
	ActiveConns int               // Currently active connections
	Config      map[string]Object // Pool configuration
	Conns       []*DBConnection   // Pool of reusable connections
	Mu          sync.RWMutex      // Thread-safe pool management
}

func (d *DBPool) Type() ObjectType { return DB_POOL_OBJ }
func (d *DBPool) Inspect() string {
	d.Mu.RLock()
	defer d.Mu.RUnlock()
	return fmt.Sprintf("DB_POOL(%s, active=%d/%d)", d.DBType, d.ActiveConns, d.MaxConns)
}

// Singleton objects for common values
var (
	NULL     = &Null{}
	TRUE     = &Boolean{Value: true}
	FALSE    = &Boolean{Value: false}
	BREAK    = &Break{}
	CONTINUE = &Continue{}
)

// NativeBoolToBooleanObject converts a native bool to a Boolean object
func NativeBoolToBooleanObject(input bool) *Boolean {
	if input {
		return TRUE
	}
	return FALSE
}
