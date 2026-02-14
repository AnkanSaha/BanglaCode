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

## File Extensions

BanglaCode supports three file extensions:
- **`.bang`** - Primary extension (recommended)
- **`.bangla`** - Alternative extension (বাংলা)
- **`.bong`** - Alternative extension (বং)

All three extensions provide identical functionality with full syntax highlighting and IntelliSense support in VS Code.

## Coding Standards (MUST FOLLOW)

### 🎯 Coding Mindset: Think Like a 7-Year Experienced System Engineer

**CRITICAL: When writing code for BanglaCode, you MUST adopt the mindset and expertise of a senior system engineer with 7+ years of experience.**

#### What This Means:

**1. Deep System Knowledge**
- You understand **how systems work internally** - not just surface-level APIs
- You know how Go's runtime works: goroutines, channels, garbage collector, memory allocation
- You understand interpreter architecture: lexing, parsing, AST traversal, evaluation
- You think about **data flow** through the entire system
- You consider **memory layout**, **CPU cache**, and **allocation patterns**

**2. Performance is Non-Negotiable**
- Performance isn't an afterthought - it's **designed in from the start**
- You benchmark before and after changes
- You know the cost of operations:
  - Map lookup vs array access
  - Pointer receiver vs value receiver
  - Channel operations vs direct access
  - Allocation cost and GC pressure
- You write code that's **fast by default**, not code that needs optimization later
- You avoid unnecessary allocations, loops, and function calls
- You use profiling tools (`go test -bench`, `pprof`) to validate performance

**3. Production-Grade Thinking**
- You write code as if it will run in production **tomorrow**
- No experiments, no "let's see if this works", no temporary solutions
- Every line of code is **intentional** and **optimized**
- You consider edge cases, error conditions, and failure modes
- You think about **scalability** - will this work with 1M operations?

**4. Systems Thinking**
- You understand how components interact and affect each other
- You consider the **entire pipeline**: Lexer → Parser → AST → Evaluator
- You think about **data structures** that minimize allocations and maximize cache hits
- You design for **composability** and **reusability**
- You avoid tight coupling and circular dependencies

**5. Expert-Level Code Organization**
- You instinctively break large systems into small, focused components
- You create **clean abstractions** that hide complexity
- You write self-documenting code with clear intent
- You follow established patterns consistently across the codebase

#### How to Apply This:

✅ **DO:**
- Ask yourself: "How would a senior engineer implement this?"
- Understand the **full context** before writing code
- Profile and measure performance impact
- Use appropriate data structures (arrays for index access, maps for lookups)
- Minimize allocations (reuse objects, use pointers wisely)
- Write concurrent code correctly (proper channel usage, no race conditions)
- Think about the **entire system**, not just your immediate task

❌ **DON'T:**
- Write naive implementations that "just work"
- Ignore performance implications
- Create abstractions without understanding the cost
- Add features without considering system-wide impact
- Write code you wouldn't want to maintain in 2 years

#### Example: Senior Engineer vs Junior Engineer

**Junior Engineer Approach (AVOID):**
```go
// Just make it work
func processItems(items []string) {
    for _, item := range items {
        result := expensiveOperation(item)  // Allocates every time
        results = append(results, result)    // May reallocate multiple times
    }
}
```

**Senior Engineer Approach (FOLLOW):**
```go
// Optimized, intentional, production-ready
func processItems(items []string) []Result {
    // Pre-allocate to avoid reallocations
    results := make([]Result, 0, len(items))

    // Reuse buffer to minimize allocations
    var buf bytes.Buffer

    for i := range items {
        // Use index to avoid copying strings
        processItemOptimized(&items[i], &buf, &results)
        buf.Reset()
    }

    return results
}
```

**Remember:** You are not learning - you are an expert. Write code that demonstrates 7+ years of system engineering experience, deep performance knowledge, and production-grade thinking.

---

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

#### 🚨 STRICT RULE: FOLDER-BASED ORGANIZATION - GROUP RELATED COMPONENTS 🚨

**Not only split into multiple files, but also organize into FOLDERS for clean architecture:**

- ❌ **NEVER** keep all files in one flat directory
- ✅ **ALWAYS** group related files into logical folders
- ✅ **EASY TO NAVIGATE** - folder structure reflects architecture

**Folder Organization Rules:**

1. **Group by domain/feature:**
   ```
   src/evaluator/
   ├── core/
   │   ├── evaluator.go      # Main Eval() logic
   │   ├── expressions.go    # Expression evaluation
   │   └── statements.go     # Statement evaluation
   ├── builtins/
   │   ├── builtins.go       # Core infrastructure
   │   ├── util.go           # Utility functions (dekho, dhoron, etc.)
   │   ├── math.go           # Math functions
   │   ├── string.go         # String operations
   │   └── array.go          # Array operations
   ├── async/
   │   ├── async.go          # Async/await core logic
   │   ├── promise.go        # Promise management
   │   └── builtins.go       # Async built-in functions
   ├── io/
   │   ├── file.go           # File I/O operations
   │   └── builtins.go       # I/O built-in functions
   ├── network/
   │   ├── http.go           # HTTP client/server
   │   ├── json.go           # JSON parsing
   │   └── builtins.go       # Network built-in functions
   ├── oop/
   │   ├── classes.go        # Class instantiation
   │   └── instances.go      # Instance methods
   ├── modules/
   │   └── modules.go        # Import/export system
   ├── errors/
   │   ├── errors.go         # Error handling
   │   └── exceptions.go     # Try/catch/finally
   └── helpers/
       └── helpers.go        # Shared helper functions
   ```

2. **Each folder should:**
   - Have a **single, clear responsibility** (IO, Network, Async, etc.)
   - Contain **related files** working together
   - Be **independently understandable** - clear what the folder does
   - Have **minimal dependencies** on other folders

3. **Folder naming conventions:**
   - Use **lowercase** names
   - Use **singular** form (e.g., `async/` not `asyncs/`)
   - Use **domain terms** (e.g., `network/` not `net_stuff/`)
   - Be **descriptive** and **concise**

4. **When to create a new folder:**
   - When you have **3+ related files** for a feature
   - When a feature has **distinct responsibilities** (e.g., client + server in `network/`)
   - When files share **common domain logic** (e.g., all HTTP-related in `network/`)
   - When it improves **navigability** and **understanding**

**Examples of Good Folder Structure:**

✅ **GOOD - Organized by domain:**
```
src/evaluator/
├── builtins/      # All built-in functions
│   ├── math.go
│   ├── string.go
│   └── array.go
├── async/         # All async/await logic
│   ├── async.go
│   └── promise.go
└── network/       # All network operations
    ├── http.go
    └── json.go
```

❌ **BAD - All files in one directory:**
```
src/evaluator/
├── evaluator.go
├── builtins_math.go
├── builtins_string.go
├── builtins_array.go
├── async.go
├── promise.go
├── http.go
└── json.go         # Hard to navigate!
```

**Benefits of Folder Organization:**
- ✅ **Clear architecture** - folder structure shows design
- ✅ **Easy to find code** - know exactly where to look
- ✅ **Scalable** - can add new folders without cluttering
- ✅ **Team-friendly** - multiple people can work on different folders
- ✅ **Better imports** - `import "evaluator/async"` is clearer than `import "evaluator"`

**General Architecture Rules:**
- Maintain **clean architecture** and **modularity**
- Each package/folder should have a single, clear responsibility
- Keep functions small and focused (ideally < 50 lines per function)
- Avoid tight coupling between packages/folders
- Follow existing project structure patterns
- Use clear file naming that describes the component: `<feature>.go`, `<component>_test.go`
- Use clear folder naming that describes the domain: `async/`, `network/`, `io/`

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

**🚨 CRITICAL RULE: ALWAYS UPDATE DOCUMENTATION WEBSITE 🚨**

**WHENEVER YOU CHANGE ANYTHING** in the interpreter, VS Code extension, or add/modify features, you **MUST IMMEDIATELY** update the Documentation website in `Documentation/` folder. This is **NON-NEGOTIABLE** and **MANDATORY**.

**This includes:**
- ✅ Adding new keywords, built-in functions, or syntax
- ✅ Modifying existing features or behavior
- ✅ Fixing bugs that change functionality
- ✅ Adding new examples or use cases
- ✅ Performance improvements worth documenting
- ✅ Breaking changes or deprecations
- ✅ **ANY change that affects how users write BanglaCode**

**Why this is critical:**
- The Documentation website is the **primary learning resource** for users
- Outdated documentation causes confusion and frustration
- Users will not discover new features if they're not documented
- Documentation must **always match** the current codebase state

**If you add a feature but DON'T update documentation, the feature is INCOMPLETE.**

---

When adding **any new feature** (keyword, built-in function, syntax, control flow), you **MUST** also update the Documentation website in `Documentation/` folder:

| Feature Type | Files to Update |
|--------------|-----------------|
| New keyword/syntax | `Documentation/app/docs/syntax/page.tsx` |
| New built-in function | `Documentation/app/docs/functions/page.tsx` |
| New control flow (if/while/for) | `Documentation/app/docs/control-flow/page.tsx` |
| New OOP feature (class/method) | `Documentation/app/docs/oop/page.tsx` |
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
├── lib/
│   └── docs-config.ts               # Documentation navigation config
└── components/                      # Shared UI components
```

**Checklist for every new feature (ALL ITEMS MANDATORY):**
1. ✅ Implement in interpreter (`src/`) - **break into separate files for each component**
2. ✅ Write tests in `test/`
3. ✅ Add syntax highlighting in `Extension/syntaxes/banglacode.tmLanguage.json`
4. ✅ Add snippet in `Extension/snippets/banglacode.json`
5. ✅ **Update Documentation website (`Documentation/app/docs/`) - MANDATORY, NOT OPTIONAL**
6. ✅ Update README.md and SYNTAX.md
7. ✅ **Benchmark performance impact** - ensure no regression

**⚠️ WARNING:** If you skip step 5 (Documentation website update), the feature is considered **INCOMPLETE** and **UNFINISHED**. Users will not know about the feature, making it effectively useless. ALWAYS update documentation.

### Core Principles Summary (CRITICAL)

When writing code for BanglaCode, always follow these principles in order of priority:

1. **🚀 PERFORMANCE FIRST** - Performance is non-negotiable
   - Every feature must be benchmarked
   - Simple syntax + high performance = ideal
   - Never sacrifice speed for abstraction

2. **📚 ALWAYS UPDATE DOCUMENTATION** - Documentation is MANDATORY, not optional
   - **EVERY change must be documented** in `Documentation/` folder
   - Users cannot use features they don't know exist
   - Outdated documentation is worse than no documentation
   - A feature without documentation is an **incomplete feature**
   - Update docs **immediately** when making changes, not "later"
   - If you change code but skip documentation, **the task is NOT done**

3. **📁 NO LARGE FILES - MULTIPLE FILES MANDATORY** - STRICT enforcement
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

4. **📦 COMPONENT-BASED DESIGN** - One file = one component
   - Each component/feature gets its own dedicated file
   - Clear file naming: `<feature>.go`, `<feature>_helpers.go`, `<feature>_test.go`
   - Related files grouped by prefix (e.g., `async.go`, `async_helpers.go`, `async_builtins.go`)

5. **🏗️ CLEAN ARCHITECTURE** - Separation of concerns
   - Each file has ONE responsibility
   - Minimal coupling between components
   - Follow SOLID principles
   - Use interfaces for abstraction

6. **✨ SIMPLE SYNTAX** - User experience matters
   - Bengali keywords that are intuitive
   - Consistent with existing patterns
   - Easy to read and understand

**Example of IDEAL Component Design with Folder Organization:**
```
src/evaluator/
├── core/
│   ├── evaluator.go     # Main Eval() switch (200 lines)
│   ├── expressions.go   # Expression evaluation (300 lines)
│   └── statements.go    # Statement evaluation (280 lines)
├── builtins/
│   ├── builtins.go      # Core infrastructure (12 lines)
│   ├── util.go          # Utility functions (135 lines)
│   ├── math.go          # Math functions (133 lines)
│   ├── string.go        # String functions (175 lines)
│   └── array.go         # Array functions (161 lines)
├── async/
│   ├── async.go         # Async/await core logic (132 lines)
│   ├── promise.go       # Promise management (can be extracted from async.go)
│   └── builtins.go      # Async built-in functions (101 lines)
├── io/
│   └── builtins.go      # File I/O functions (98 lines)
├── network/
│   ├── http.go          # HTTP client/server (can be extracted)
│   ├── json.go          # JSON parsing (can be extracted)
│   └── builtins.go      # Network built-in functions (326 lines)
├── oop/
│   └── classes.go       # OOP features (177 lines)
├── modules/
│   └── modules.go       # Import/export system (220 lines)
├── errors/
│   └── errors.go        # Error handling (54 lines)
└── helpers/
    └── helpers.go       # Shared helper functions (86 lines)
```

**Current Structure (Acceptable, but can be improved):**
```
src/evaluator/
├── evaluator.go         # Main Eval() switch (191 lines)
├── async.go             # Async/await logic (132 lines)
├── classes.go           # OOP features (177 lines)
├── modules.go           # Import/export (220 lines)
├── builtins.go          # Built-in core infrastructure (12 lines)
├── builtins_util.go     # Utility functions (135 lines)
├── builtins_math.go     # Math functions (133 lines)
├── builtins_string.go   # String functions (175 lines)
├── builtins_array.go    # Array functions (161 lines)
├── builtins_async.go    # Async built-in functions (101 lines)
├── builtins_io.go       # I/O built-in functions (98 lines)
├── builtins_http.go     # HTTP/JSON functions (326 lines)
├── errors.go            # Error handling (54 lines)
├── expressions.go       # Expression evaluation (414 lines)
├── statements.go        # Statement evaluation (190 lines)
└── helpers.go           # Helper functions (86 lines)
```
Note: Current structure is acceptable for now, but as the project grows, reorganize into folders.

**Example of BAD vs GOOD vs BEST File Structure:**

❌ **BAD - One Large File:**
```
src/evaluator/
└── builtins.go       # 800 lines - TOO BIG! Hard to navigate and understand
```

✅ **GOOD - Multiple Focused Files (Current):**
```
src/evaluator/
├── builtins.go          # 12 lines - Core infrastructure
├── builtins_util.go     # 135 lines - Utility functions
├── builtins_string.go   # 175 lines - String manipulation
├── builtins_array.go    # 161 lines - Array operations
├── builtins_math.go     # 133 lines - Math functions
├── builtins_async.go    # 101 lines - Async operations
├── builtins_io.go       # 98 lines - File I/O
└── builtins_http.go     # 326 lines - HTTP/JSON
```

✅✅ **BEST - Folder-Based Organization (Recommended for growth):**
```
src/evaluator/
├── builtins/
│   ├── builtins.go      # 12 lines - Core infrastructure
│   ├── util.go          # 135 lines - Utility functions
│   ├── math.go          # 133 lines - Math functions
│   ├── string.go        # 175 lines - String manipulation
│   └── array.go         # 161 lines - Array operations
├── async/
│   └── builtins.go      # 101 lines - Async operations
├── io/
│   └── builtins.go      # 98 lines - File I/O
└── network/
    └── builtins.go      # 326 lines - HTTP/JSON
```

**When to move from GOOD to BEST:**
- When you have **5+ files** in a directory
- When files can be **logically grouped** by domain (network, io, async)
- When you're **adding new features** that fit into existing groups
- When navigating becomes **harder** due to many files

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
6. **✅ DO update Documentation website (`Documentation/`) - MANDATORY**

### Critical Rules:
- 📚 **ALWAYS update Documentation/** - for EVERY change, not just new features
- 📏 **File size limit: 500 lines MAX, 300 lines IDEAL**
- 🚀 **Performance first** - benchmark everything
- 📦 **Multiple files** - never one large file
- 🎯 **One responsibility** - per file, per function
- ✨ **Simple syntax** - intuitive Bengali keywords

**Remember:**
- **Documentation is NOT optional** - update it for every change or the task is incomplete
- Performance and simplicity are the core values of BanglaCode
- Small, focused files = easy to understand and maintain
- If a feature impacts performance negatively, it should be redesigned
- If a file grows beyond 500 lines, split it immediately
- **If you change code but don't update docs, users won't know about it**
