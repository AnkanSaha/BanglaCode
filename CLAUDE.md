# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

BanglaCode is a Bengali-syntax programming language interpreter written in Go. It uses Banglish (Bengali words in English script) keywords like `dhoro` (let), `jodi` (if), `kaj` (function) to make programming accessible to Bengali speakers. The interpreter follows a classic tree-walking architecture: Source → Lexer → Parser → AST → Evaluator → Result.

## Build & Run Commands

```bash
# Build the interpreter
go build -o banglacode main.go

# Run a BanglaCode file
./banglacode examples/hello.bang

# Start interactive REPL
./banglacode

# Run with go directly
go run main.go examples/hello.bang

# Cross-compile for different platforms
GOOS=windows GOARCH=amd64 go build -o banglacode.exe .
GOOS=darwin GOARCH=arm64 go build -o banglacode .
GOOS=linux GOARCH=amd64 go build -o banglacode .
```

## Architecture

The interpreter pipeline flows through these components in order:

1. **Lexer** (`src/lexer/`) - Tokenizes source code into tokens. `token.go` contains keyword mappings (29 Bengali keywords).

2. **Parser** (`src/parser/`) - Builds AST using Pratt parsing (top-down operator precedence). `precedence.go` defines operator precedence levels.

3. **AST** (`src/ast/`) - Node definitions. Statements inherit from `Statement` interface, expressions from `Expression` interface.

4. **Object** (`src/object/`) - Runtime value types (Number, String, Boolean, Array, Map, Function, Class, Instance, Promise, Error, etc.). `environment.go` manages variable scopes with parent-child chain for closures.

5. **Evaluator** (`src/evaluator/`) - Tree-walking interpreter:
   - `evaluator.go` - Main `Eval()` switch on AST node types
   - `builtins.go` - 45+ built-in functions (`dekho`, `dorghyo`, `dhokao`, `ghumaao`, `anun_async`, etc.)
   - `async.go` - Async/await: promise management, async function execution, `proyash`/`opekha`
   - `classes.go` - OOP: class instantiation, method calls, `ei` (this) binding
   - `modules.go` - Import/export: `ano` (import), `pathao` (export), `hisabe` (alias)
   - `errors.go` - Try/catch/finally: `chesta`/`dhoro_bhul`/`shesh`

6. **REPL** (`src/repl/`) - Interactive shell with multi-line support and help system.

## Key Bengali Keywords

| Keyword | English | Usage |
|---------|---------|-------|
| `dhoro` | let/var | `dhoro x = 5;` |
| `sthir` | const | `sthir PI = 3.14;` (immutable constant) |
| `bishwo` | global | `bishwo count = 0;` (global variable) |
| `jodi`/`nahole` | if/else | `jodi (x > 0) { } nahole { }` |
| `jotokkhon` | while | `jotokkhon (x < 10) { }` |
| `ghuriye` | for | `ghuriye (dhoro i = 0; i < 5; i = i + 1) { }` |
| `kaj` | function | `kaj add(a, b) { ferao a + b; }` |
| `ferao` | return | `ferao result;` |
| `sreni`/`shuru` | class/constructor | `sreni Person { shuru(naam) { ei.naam = naam; } }` |
| `notun`/`ei` | new/this | `notun Person("Ankan")` |
| `sotti`/`mittha`/`khali` | true/false/null | Boolean and null literals |
| `ebong`/`ba`/`na` | and/or/not | Logical operators |
| `thamo`/`chharo` | break/continue | Loop control |
| `ano`/`pathao`/`hisabe` | import/export/as | Module system |
| `chesta`/`dhoro_bhul`/`shesh`/`felo` | try/catch/finally/throw | Error handling |
| `proyash`/`opekha` | async/await | Asynchronous programming: `proyash kaj fetchData() { ... }`, `opekha promise` |

## Adding New Features

**New built-in function:** Add to `builtins` map in `src/evaluator/builtins.go`

**New keyword:**
1. Add token constant and keyword mapping in `src/lexer/token.go`
2. Add AST node in `src/ast/statements.go` or `expressions.go`
3. Add parser case in `src/parser/statements.go` or `expressions.go`
4. Add evaluator case in `src/evaluator/evaluator.go`

**New object type:** Define in `src/object/object.go` implementing the `Object` interface (`Type()` and `Inspect()` methods)

## File Extension

BanglaCode source files use `.bang` extension.

## Coding Standards (MUST FOLLOW)

### Code Quality
- Write **production-ready code** only - no experimental or half-done solutions
- Follow **SOLID principles**:
  - Single Responsibility: Each function/struct does one thing
  - Open/Closed: Open for extension, closed for modification
  - Liskov Substitution: Subtypes must be substitutable for base types
  - Interface Segregation: Small, specific interfaces over large ones
  - Dependency Inversion: Depend on abstractions, not concretions
- **No hacks or workarounds** - implement proper solutions
- Keep code **clean, readable, and self-documenting**
- Use meaningful variable/function names that explain intent

### Architecture

#### 🚨 STRICT RULE: NO LARGE FILES - ALWAYS BREAK INTO MULTIPLE FILES 🚨

**When building ANY feature in `src/`, you MUST create multiple files:**

- ❌ **NEVER** create one large file with all logic
- ✅ **ALWAYS** break into multiple focused files using imports and grouping
- ✅ **ALWAYS** use Go's import system to organize related components
- ✅ **EASY TO UNDERSTAND** is mandatory - small files are easier to read and maintain

**Maximum File Size Limits (STRICT):**
- 📏 **500 lines MAX** per file - if you approach this, split immediately
- 📏 **300 lines IDEAL** - aim for this size for optimal readability
- 📏 **50 lines per function** - break down large functions into smaller ones

**How to Break Code into Multiple Files:**

1. **Group by feature/component:**
   ```
   src/evaluator/
   ├── async.go           # Only async/await logic
   ├── async_helpers.go   # Async helper functions
   ├── async_builtins.go  # Async built-in functions (if many)
   └── async_test.go      # Async tests
   ```

2. **Use imports to connect files in same package:**
   ```go
   // In async_helpers.go
   package evaluator

   func createPromise() *object.Promise {
       // Helper shared across async files
   }

   // In async.go
   package evaluator
   // No import needed - same package!
   // Can directly use createPromise()
   ```

3. **When a feature grows, split further:**
   ```
   Before (BAD - 800 lines):
   ├── builtins.go        # 800 lines - TOO BIG!

   After (GOOD - broken down):
   ├── builtins.go        # 150 lines - core infrastructure
   ├── builtins_string.go # 120 lines - string functions
   ├── builtins_array.go  # 130 lines - array functions
   ├── builtins_math.go   # 100 lines - math functions
   ├── builtins_async.go  # 150 lines - async functions
   └── builtins_io.go     # 150 lines - I/O functions
   ```

**Benefits of Multiple Files:**
- ✅ **Easy to navigate** - find code faster
- ✅ **Easy to understand** - each file has clear purpose
- ✅ **Easy to test** - test files match implementation files
- ✅ **Easy to review** - smaller diffs in PRs
- ✅ **Easy to maintain** - changes are isolated
- ✅ **Better performance** - Go compiler can parallelize builds

**General Architecture Rules:**
- Maintain **clean architecture** and **modularity**
- Each package should have a single, clear responsibility
- Keep functions small and focused (ideally < 50 lines per function)
- Avoid tight coupling between packages
- Follow existing project structure patterns
- Use clear file naming that describes the component: `<feature>.go`, `<component>_test.go`

### Performance (HIGHEST PRIORITY)
- **Performance is the FIRST priority** when adding or modifying features
- **Simple syntax + performance** is the highest priority combination:
  - Always choose the simpler syntax if performance is equivalent
  - Never sacrifice performance for complex abstractions
  - Benchmark new features to ensure they don't degrade performance
- Write **optimized, fast code** - no unnecessary allocations or loops
- Avoid redundant operations and memory allocations
- Use appropriate data structures for the task:
  - Prefer arrays over maps when index access is needed
  - Use buffered channels (size 1) for promise communication to prevent blocking
  - Reuse objects when possible instead of creating new ones
- Profile before optimizing, but write efficient code from the start
- **Measure performance impact** of new features:
  ```bash
  # Benchmark before and after
  go test -bench=. -benchmem ./test/
  ```
- For interpreter operations:
  - Minimize AST node allocations
  - Cache frequently accessed values
  - Use pointer receivers for methods to avoid copies
  - Prefer iterative solutions over recursive when performance-critical

### Security
- **Before writing any feature**, analyze potential security risks:
  - Input validation and sanitization
  - Injection vulnerabilities (code injection in eval, file path traversal)
  - Resource exhaustion (infinite loops, memory bombs)
  - Unsafe file operations
- Implement security fixes as part of the feature, not as an afterthought

### Testing
- Write test cases for **all new features** in `test/` folder
- Test file naming: `<feature>_test.go`
- Include:
  - Unit tests for individual functions
  - Edge cases and error conditions
  - Integration tests for component interactions
- Run tests before considering any feature complete:
  ```bash
  go test ./test/...
  ```

### VS Code Extension (MANDATORY)
When adding **any new feature** (keyword, built-in function, syntax), you **MUST** also update the VS Code Extension in `Extension/` folder:

| Feature Type | Files to Update |
|--------------|-----------------|
| New keyword | `Extension/syntaxes/banglacode.tmLanguage.json` (syntax highlighting) |
| New built-in function | `Extension/syntaxes/banglacode.tmLanguage.json` (highlight as function) |
| New snippet | `Extension/snippets/banglacode.json` (add code snippet) |
| New syntax pattern | `Extension/language-configuration.json` (brackets, comments, etc.) |

**Extension folder structure:**
```
Extension/
├── syntaxes/banglacode.tmLanguage.json  # Syntax highlighting rules
├── snippets/banglacode.json              # Code snippets for autocomplete
├── language-configuration.json           # Language settings (brackets, comments)
├── package.json                          # Extension metadata & configuration
└── extension.js                          # Extension activation logic
```

### Documentation Website (MANDATORY)
When adding **any new feature** (keyword, built-in function, syntax, control flow), you **MUST** also update the Documentation website in `Documentation/` folder:

| Feature Type | Files to Update |
|--------------|-----------------|
| New keyword/syntax | `Documentation/app/docs/syntax/page.tsx` |
| New built-in function | `Documentation/app/docs/functions/page.tsx` |
| New control flow (if/while/for) | `Documentation/app/docs/control-flow/page.tsx` |
| New OOP feature (class/method) | `Documentation/app/docs/oop/page.tsx` |
| New example code | `Documentation/app/playground/examples.ts` |
| New documentation section | `Documentation/lib/docs-config.ts` (navigation config) |

**Documentation folder structure:**
```
Documentation/
├── app/
│   ├── docs/
│   │   ├── syntax/page.tsx          # Syntax documentation
│   │   ├── functions/page.tsx       # Built-in functions documentation
│   │   ├── control-flow/page.tsx    # Control flow documentation
│   │   ├── oop/page.tsx             # OOP documentation
│   │   └── installation/page.tsx    # Installation guide
│   └── playground/
│       ├── page.tsx                 # Interactive playground
│       └── examples.ts              # Code examples for playground
├── lib/
│   └── docs-config.ts               # Documentation navigation config
└── components/                      # Shared UI components
```

**Checklist for every new feature:**
1. ✅ Implement in interpreter (`src/`) - **break into separate files for each component**
2. ✅ Write tests in `test/`
3. ✅ Add syntax highlighting in `Extension/syntaxes/banglacode.tmLanguage.json`
4. ✅ Add snippet in `Extension/snippets/banglacode.json`
5. ✅ Update Documentation website (`Documentation/app/docs/`)
6. ✅ Add playground examples if applicable (`Documentation/app/playground/examples.ts`)
7. ✅ Update README.md and SYNTAX.md
8. ✅ **Benchmark performance impact** - ensure no regression

### Core Principles Summary (CRITICAL)

When writing code for BanglaCode, always follow these principles in order of priority:

1. **🚀 PERFORMANCE FIRST** - Performance is non-negotiable
   - Every feature must be benchmarked
   - Simple syntax + high performance = ideal
   - Never sacrifice speed for abstraction

2. **📁 NO LARGE FILES - MULTIPLE FILES MANDATORY** - STRICT enforcement
   - **NEVER write one large file** - always break into multiple focused files
   - Use Go's import system and same-package grouping extensively
   - **Maximum 500 lines per file** - split immediately if approaching this
   - **Ideal 300 lines per file** - easier to understand and navigate
   - Example: Instead of `builtins.go` (800 lines), create:
     - `builtins.go` (core infrastructure)
     - `builtins_string.go` (string functions)
     - `builtins_array.go` (array functions)
     - `builtins_math.go` (math functions)
     - `builtins_async.go` (async functions)
   - Files in the same package can access each other without imports
   - **Easy to understand = small, focused files**

3. **📦 COMPONENT-BASED DESIGN** - One file = one component
   - Each component/feature gets its own dedicated file
   - Clear file naming: `<feature>.go`, `<feature>_helpers.go`, `<feature>_test.go`
   - Related files grouped by prefix (e.g., `async.go`, `async_helpers.go`, `async_builtins.go`)

4. **🏗️ CLEAN ARCHITECTURE** - Separation of concerns
   - Each file has ONE responsibility
   - Minimal coupling between components
   - Follow SOLID principles
   - Use interfaces for abstraction

5. **✨ SIMPLE SYNTAX** - User experience matters
   - Bengali keywords that are intuitive
   - Consistent with existing patterns
   - Easy to read and understand

**Example of Good Component Design:**
```
src/evaluator/
├── evaluator.go      # Main Eval() switch - coordinates all evaluations (200 lines)
├── async.go          # Async/await logic only - promises, goroutines (150 lines)
├── async_helpers.go  # Async helper functions (100 lines)
├── classes.go        # OOP features only - classes, instances, methods (250 lines)
├── modules.go        # Import/export only - module loading, exports (180 lines)
├── builtins.go       # Built-in core infrastructure (150 lines)
├── builtins_string.go # String built-in functions (120 lines)
├── builtins_array.go  # Array built-in functions (130 lines)
├── builtins_math.go   # Math built-in functions (100 lines)
├── builtins_async.go  # Async built-in functions (150 lines)
├── builtins_io.go     # I/O built-in functions (150 lines)
├── errors.go         # Error handling only - try/catch/finally (200 lines)
├── expressions.go    # Expression evaluation only (300 lines)
└── statements.go     # Statement evaluation only (280 lines)
```

**Example of BAD vs GOOD File Structure:**

❌ **BAD - One Large File:**
```
src/evaluator/
└── builtins.go       # 800 lines - TOO BIG! Hard to navigate and understand
```

✅ **GOOD - Multiple Focused Files:**
```
src/evaluator/
├── builtins.go          # 150 lines - Core infrastructure, registration
├── builtins_string.go   # 120 lines - String manipulation functions
├── builtins_array.go    # 130 lines - Array operations
├── builtins_math.go     # 100 lines - Mathematical functions
├── builtins_async.go    # 150 lines - Async operations
└── builtins_io.go       # 150 lines - File I/O operations
```

**How Files Connect in Same Package:**
```go
// builtins_async.go
package evaluator

// This file defines async built-in functions
// Can use helpers from async.go without import (same package!)

func init() {
    // Register async built-ins
    builtins["ghumaao"] = &object.Builtin{
        Fn: func(args ...object.Object) object.Object {
            // Uses createPromise() from async.go directly
            promise := createPromise()
            // ... rest of implementation
        },
    }
}
```

---

## 🎯 FINAL REMINDERS (MUST READ)

### When Adding ANY Feature to `src/`:

1. **❌ DO NOT create one large file**
2. **✅ DO break into multiple files** (300-500 lines each)
3. **✅ DO use imports and Go's same-package grouping**
4. **✅ DO benchmark performance before and after**
5. **✅ DO make code easy to understand through small, focused files**

### Critical Rules:
- 📏 **File size limit: 500 lines MAX, 300 lines IDEAL**
- 🚀 **Performance first** - benchmark everything
- 📦 **Multiple files** - never one large file
- 🎯 **One responsibility** - per file, per function
- ✨ **Simple syntax** - intuitive Bengali keywords

**Remember:**
- Performance and simplicity are the core values of BanglaCode
- Small, focused files = easy to understand and maintain
- If a feature impacts performance negatively, it should be redesigned
- If a file grows beyond 500 lines, split it immediately
