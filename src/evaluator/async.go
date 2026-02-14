package evaluator

import (
	"BanglaCode/src/ast"
	"BanglaCode/src/object"
	"fmt"
)

// createPromise creates a new pending promise with channels
func createPromise() *object.Promise {
	return &object.Promise{
		State:      object.PROMISE_PENDING,
		ResultChan: make(chan object.Object, 1),
		ErrorChan:  make(chan object.Object, 1),
	}
}

// resolvePromise resolves a promise with a value
func resolvePromise(promise *object.Promise, value object.Object) {
	promise.Mu.Lock()
	promise.State = object.PROMISE_RESOLVED
	promise.Value = value
	promise.Mu.Unlock()
	promise.ResultChan <- value
}

// rejectPromise rejects a promise with an error
func rejectPromise(promise *object.Promise, err object.Object) {
	promise.Mu.Lock()
	promise.State = object.PROMISE_REJECTED
	promise.Error = err
	promise.Mu.Unlock()
	promise.ErrorChan <- err
}

// evalAsyncFunctionLiteral evaluates an async function literal and creates an async function object
func evalAsyncFunctionLiteral(node *ast.AsyncFunctionLiteral, env *object.Environment) object.Object {
	params := node.Parameters
	body := node.Body
	name := ""
	if node.Name != nil {
		name = node.Name.Value
	}

	fn := &object.Function{
		Parameters:    params,
		RestParameter: node.RestParameter,
		Env:           env,
		Body:          body,
		Name:          name,
		IsAsync:       true, // Mark as async
	}

	// If function has a name, bind it in the environment
	if name != "" {
		env.Set(name, fn)
	}

	return fn
}

// evalAsyncFunctionCall executes an async function in a goroutine and returns a promise
func evalAsyncFunctionCall(fn *object.Function, args []object.Object, env *object.Environment) object.Object {
	promise := createPromise()

	// Spawn goroutine to execute async function
	go func() {
		// Recover from panics in async functions
		defer func() {
			if r := recover(); r != nil {
				errorMsg := fmt.Sprintf("panic in async function: %v", r)
				rejectPromise(promise, &object.Error{Message: errorMsg})
			}
		}()

		// Create new environment for function execution
		extendedEnv := extendFunctionEnv(fn, args)

		// Execute function body
		result := Eval(fn.Body, extendedEnv)

		// Unwrap return value
		if returnValue, ok := result.(*object.ReturnValue); ok {
			result = returnValue.Value
		}

		// Check for errors or exceptions
		if err, ok := result.(*object.Error); ok {
			rejectPromise(promise, err)
			return
		}

		if exc, ok := result.(*object.Exception); ok {
			rejectPromise(promise, exc)
			return
		}

		// Resolve promise with result
		resolvePromise(promise, result)
	}()

	return promise
}

// evalAwaitExpression waits for a promise to resolve or reject
func evalAwaitExpression(node *ast.AwaitExpression, env *object.Environment) object.Object {
	// Evaluate the expression that should produce a promise
	value := Eval(node.Expression, env)
	if isError(value) {
		return value
	}

	// Value must be a promise
	promise, ok := value.(*object.Promise)
	if !ok {
		return newError("opekha (await) can only be used with promises, got %s", value.Type())
	}

	// Check current state
	promise.Mu.RLock()
	state := promise.State
	promise.Mu.RUnlock()

	// If already resolved/rejected, return immediately
	if state == object.PROMISE_RESOLVED {
		return promise.Value
	} else if state == object.PROMISE_REJECTED {
		return promise.Error
	}

	// Wait for promise to complete
	select {
	case result := <-promise.ResultChan:
		return result
	case err := <-promise.ErrorChan:
		return err
	}
}
