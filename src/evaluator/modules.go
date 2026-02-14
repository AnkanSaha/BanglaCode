package evaluator

import (
	"BanglaCode/src/ast"
	"BanglaCode/src/lexer"
	"BanglaCode/src/object"
	"BanglaCode/src/parser"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

// Module cache to prevent circular imports
var (
	moduleCache = make(map[string]*object.Module)
	moduleMutex sync.RWMutex
	currentDir  = "."
)

// SetCurrentDir sets the current directory for resolving imports
func SetCurrentDir(dir string) {
	currentDir = dir
}

// evalImportStatement evaluates import statements
func evalImportStatement(is *ast.ImportStatement, env *object.Environment) object.Object {
	modulePath := is.Path.Value

	// Resolve relative path
	fullPath := filepath.Join(currentDir, modulePath)

	// Check if it's a JSON file
	if strings.HasSuffix(modulePath, ".json") {
		return evalJSONImport(fullPath, modulePath, is.Alias, env)
	}

	// Check module cache
	moduleMutex.RLock()
	if mod, ok := moduleCache[fullPath]; ok {
		moduleMutex.RUnlock()
		// Import exports into environment
		importModuleExports(mod, is.Alias, env)
		return mod
	}
	moduleMutex.RUnlock()

	// Read module file
	content, err := os.ReadFile(fullPath)
	if err != nil {
		return newError("cannot import module '%s': %s", modulePath, err.Error())
	}

	// Create module environment
	moduleEnv := object.NewEnvironment()

	// Parse module
	l := lexer.New(string(content))
	p := parser.New(l)
	program := p.ParseProgram()

	if len(p.Errors()) != 0 {
		return newError("parse error in module '%s': %s", modulePath, p.Errors()[0])
	}

	// Save current directory and set module directory
	oldDir := currentDir
	currentDir = filepath.Dir(fullPath)

	// Create module object first
	mod := &object.Module{
		Name:    modulePath,
		Exports: make(map[string]object.Object),
	}

	// Only evaluate export statements and function/class definitions
	// Skip other top-level code to prevent execution on import
	for _, stmt := range program.Statements {
		var result object.Object

		switch s := stmt.(type) {
		case *ast.ExportStatement:
			// Evaluate export statements (pathao kaj ...)
			result = evalExportStatement(s, moduleEnv)
			if isError(result) {
				currentDir = oldDir
				return result
			}

		case *ast.ExpressionStatement:
			// Only evaluate function literals (kaj declarations)
			if fnLit, ok := s.Expression.(*ast.FunctionLiteral); ok && fnLit.Name != nil {
				result = Eval(s, moduleEnv)
				if isError(result) {
					currentDir = oldDir
					return result
				}
			}
			// Skip other expression statements (function calls, etc.)

		case *ast.ClassDeclaration:
			// Evaluate class declarations
			result = Eval(s, moduleEnv)
			if isError(result) {
				currentDir = oldDir
				return result
			}

		case *ast.VariableDeclaration:
			// Skip non-exported variable/constant/global declarations
			// These should only be evaluated if explicitly exported with pathao
			continue

		default:
			// Skip all other statements (variable declarations, loops, etc.)
			// This prevents top-level code execution on import
			continue
		}
	}

	// Restore directory
	currentDir = oldDir

	// Get exports from module environment (__exports__ map)
	if exports, ok := moduleEnv.Get("__exports__"); ok {
		if exportsMap, ok := exports.(*object.Map); ok {
			for k, v := range exportsMap.Pairs {
				mod.Exports[k] = v
			}
		}
	}

	// Cache module
	moduleMutex.Lock()
	moduleCache[fullPath] = mod
	moduleMutex.Unlock()

	// Import exports into environment
	importModuleExports(mod, is.Alias, env)

	return mod
}

// importModuleExports imports module exports into the environment
func importModuleExports(mod *object.Module, alias *ast.Identifier, env *object.Environment) {
	if alias != nil {
		// Import as namespace: ano "math.bang" hisabe math;
		// Access via: math.add(1, 2)
		modMap := &object.Map{Pairs: make(map[string]object.Object)}
		for k, v := range mod.Exports {
			modMap.Pairs[k] = v
		}
		env.Set(alias.Value, modMap)
	} else {
		// Import directly into namespace
		for k, v := range mod.Exports {
			env.Set(k, v)
		}
	}
}

// evalExportStatement evaluates export statements
func evalExportStatement(es *ast.ExportStatement, env *object.Environment) object.Object {
	// Evaluate the statement being exported
	result := Eval(es.Statement, env)
	if isError(result) {
		return result
	}

	// Get or create exports map
	var exportsMap *object.Map
	if exports, ok := env.Get("__exports__"); ok {
		exportsMap = exports.(*object.Map)
	} else {
		exportsMap = &object.Map{Pairs: make(map[string]object.Object)}
		env.Set("__exports__", exportsMap)
	}

	// Add to exports based on statement type
	switch stmt := es.Statement.(type) {
	case *ast.VariableDeclaration:
		exportsMap.Pairs[stmt.Name.Value] = result
	case *ast.ExpressionStatement:
		if fn, ok := stmt.Expression.(*ast.FunctionLiteral); ok && fn.Name != nil {
			if val, ok := env.Get(fn.Name.Value); ok {
				exportsMap.Pairs[fn.Name.Value] = val
			}
		}
	case *ast.ClassDeclaration:
		exportsMap.Pairs[stmt.Name.Value] = result
	}

	return result
}

// evalJSONImport handles importing JSON files
func evalJSONImport(fullPath, modulePath string, alias *ast.Identifier, env *object.Environment) object.Object {
	content, err := os.ReadFile(fullPath)
	if err != nil {
		return newError("cannot import JSON '%s': %s", modulePath, err.Error())
	}

	var jsonData interface{}
	if err := json.Unmarshal(content, &jsonData); err != nil {
		return newError("invalid JSON in '%s': %s", modulePath, err.Error())
	}

	// Convert JSON to BanglaCode object
	obj := jsonToObject(jsonData)

	// If alias provided, set with alias name, otherwise error (JSON requires alias)
	if alias != nil {
		env.Set(alias.Value, obj)
	} else {
		return newError("JSON import requires alias: ano \"%s\" hisabe <name>;", modulePath)
	}

	return obj
}
