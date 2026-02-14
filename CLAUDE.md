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

1. **Lexer** (`src/lexer/`) - Tokenizes source code into tokens. `token.go` contains keyword mappings (27 Bengali keywords).

2. **Parser** (`src/parser/`) - Builds AST using Pratt parsing (top-down operator precedence). `precedence.go` defines operator precedence levels.

3. **AST** (`src/ast/`) - Node definitions. Statements inherit from `Statement` interface, expressions from `Expression` interface.

4. **Object** (`src/object/`) - Runtime value types (Number, String, Boolean, Array, Map, Function, Class, Instance, Error, etc.). `environment.go` manages variable scopes with parent-child chain for closures.

5. **Evaluator** (`src/evaluator/`) - Tree-walking interpreter:
   - `evaluator.go` - Main `Eval()` switch on AST node types
   - `builtins.go` - 40+ built-in functions (`dekho`, `dorghyo`, `dhokao`, etc.)
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
- Maintain **clean architecture** and **modularity**
- Each package should have a single, clear responsibility
- Keep functions small and focused
- Avoid tight coupling between packages
- Follow existing project structure patterns

### Performance
- Write **optimized, fast code** - no unnecessary allocations or loops
- Avoid redundant operations
- Use appropriate data structures for the task
- Profile before optimizing, but write efficient code from the start

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
1. ✅ Implement in interpreter (`src/`)
2. ✅ Write tests in `test/`
3. ✅ Add syntax highlighting in `Extension/syntaxes/banglacode.tmLanguage.json`
4. ✅ Add snippet in `Extension/snippets/banglacode.json`
5. ✅ Update Documentation website (`Documentation/app/docs/`)
6. ✅ Add playground examples if applicable (`Documentation/app/playground/examples.ts`)
7. ✅ Update README.md and SYNTAX.md
