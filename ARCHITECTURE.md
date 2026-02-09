# BanglaCode Architecture

This document describes the technical architecture of the BanglaCode programming language interpreter.

## Table of Contents

- [Overview](#overview)
- [High-Level Architecture](#high-level-architecture)
- [Component Details](#component-details)
- [Data Flow](#data-flow)
- [Design Decisions](#design-decisions)
- [Performance Considerations](#performance-considerations)
- [Extending the Language](#extending-the-language)

---

## Overview

BanglaCode is a **tree-walking interpreter** written in Go. It follows the classic interpreter pipeline:

```
Source Code → Lexer → Parser → AST → Evaluator → Result
```

### Key Characteristics

| Aspect | Implementation |
|--------|---------------|
| **Type System** | Dynamically typed |
| **Evaluation** | Tree-walking interpreter |
| **Memory Management** | Go's garbage collector |
| **Concurrency** | Go goroutines (for HTTP server) |
| **Module System** | File-based imports |

### Technology Stack

- **Language**: Go 1.20+
- **Dependencies**: Standard library only (no external dependencies for core)
- **Build**: Native Go toolchain
- **Testing**: Go testing framework

---

## High-Level Architecture

```
┌─────────────────────────────────────────────────────────────────────┐
│                           BanglaCode                                │
├─────────────────────────────────────────────────────────────────────┤
│                                                                     │
│  ┌─────────┐    ┌─────────┐    ┌─────────┐    ┌─────────────────┐  │
│  │  main   │───▶│  REPL   │───▶│  Lexer  │───▶│     Parser      │  │
│  │  .go    │    │         │    │         │    │                 │  │
│  └─────────┘    └─────────┘    └────┬────┘    └────────┬────────┘  │
│                                     │                  │            │
│                                     ▼                  ▼            │
│                              ┌─────────────────────────────────┐   │
│                              │              AST                │   │
│                              │    (Abstract Syntax Tree)       │   │
│                              └─────────────┬───────────────────┘   │
│                                            │                        │
│                                            ▼                        │
│  ┌─────────────────────────────────────────────────────────────┐   │
│  │                       Evaluator                              │   │
│  │  ┌──────────┐  ┌──────────┐  ┌──────────┐  ┌──────────────┐ │   │
│  │  │Expressions│  │Statements│  │ Classes  │  │   Modules    │ │   │
│  │  └──────────┘  └──────────┘  └──────────┘  └──────────────┘ │   │
│  │  ┌──────────┐  ┌──────────┐                                  │   │
│  │  │ Builtins │  │  Errors  │                                  │   │
│  │  └──────────┘  └──────────┘                                  │   │
│  └─────────────────────────────┬───────────────────────────────┘   │
│                                │                                    │
│                                ▼                                    │
│                         ┌─────────────┐                             │
│                         │   Object    │                             │
│                         │   System    │                             │
│                         └─────────────┘                             │
│                                                                     │
└─────────────────────────────────────────────────────────────────────┘
```

---

## Component Details

### 1. Entry Point (`main.go`)

The entry point handles:
- Command-line argument parsing
- File reading and execution
- REPL initialization
- Version and help display

```go
// Simplified main flow
func main() {
    if len(os.Args) == 1 {
        repl.Start()      // Interactive mode
    } else {
        runFile(os.Args[1]) // Execute file
    }
}
```

### 2. Lexer (`src/lexer/`)

The lexer (tokenizer) converts source code into a stream of tokens.

#### Files
- `lexer.go` — Main lexer implementation
- `token.go` — Token type definitions and keyword mapping

#### Token Types

```go
const (
    // Literals
    IDENT   = "IDENT"
    INT     = "INT"
    FLOAT   = "FLOAT"
    STRING  = "STRING"

    // Operators
    ASSIGN  = "="
    PLUS    = "+"
    MINUS   = "-"
    // ...

    // Keywords (Bengali)
    DHORO      = "DHORO"      // let/var
    JODI       = "JODI"       // if
    NAHOLE     = "NAHOLE"     // else
    JOTOKKHON  = "JOTOKKHON"  // while
    GHURIYE    = "GHURIYE"    // for
    KAJ        = "KAJ"        // function
    FERAO      = "FERAO"      // return
    // ...
)
```

#### Keyword Mapping

```go
var keywords = map[string]TokenType{
    "dhoro":      DHORO,
    "jodi":       JODI,
    "nahole":     NAHOLE,
    "jotokkhon":  JOTOKKHON,
    "ghuriye":    GHURIYE,
    "kaj":        KAJ,
    "ferao":      FERAO,
    "sreni":      SRENI,
    "shuru":      SHURU,
    "notun":      NOTUN,
    "ei":         EI,
    "sotti":      SOTTI,
    "mittha":     MITTHA,
    "khali":      KHALI,
    "ebong":      EBONG,
    "ba":         BA,
    "na":         NA,
    // ...
}
```

#### Lexer Process

```
Input: "dhoro x = 5;"

Output Tokens:
┌─────────┬───────┬──────┬────────┐
│  Type   │ Value │ Line │ Column │
├─────────┼───────┼──────┼────────┤
│ DHORO   │ dhoro │  1   │   1    │
│ IDENT   │ x     │  1   │   7    │
│ ASSIGN  │ =     │  1   │   9    │
│ INT     │ 5     │  1   │  11    │
│ SEMI    │ ;     │  1   │  12    │
│ EOF     │       │  1   │  13    │
└─────────┴───────┴──────┴────────┘
```

### 3. Parser (`src/parser/`)

The parser constructs an Abstract Syntax Tree (AST) from tokens using recursive descent parsing with operator precedence climbing.

#### Files
- `parser.go` — Main parser and entry point
- `expressions.go` — Expression parsing
- `statements.go` — Statement parsing
- `precedence.go` — Operator precedence definitions

#### Operator Precedence

```go
const (
    _ int = iota
    LOWEST
    EQUALS      // ==
    LESSGREATER // > or <
    SUM         // +
    PRODUCT     // *
    PREFIX      // -X or !X
    CALL        // myFunction(X)
    INDEX       // array[index]
)
```

#### Parsing Strategy

The parser uses Pratt parsing (top-down operator precedence) for expressions:

```go
func (p *Parser) parseExpression(precedence int) ast.Expression {
    prefix := p.prefixParseFns[p.curToken.Type]
    leftExp := prefix()

    for precedence < p.peekPrecedence() {
        infix := p.infixParseFns[p.peekToken.Type]
        leftExp = infix(leftExp)
    }

    return leftExp
}
```

### 4. AST (`src/ast/`)

The Abstract Syntax Tree represents the program structure.

#### Files
- `ast.go` — Base node interfaces
- `expressions.go` — Expression nodes
- `statements.go` — Statement nodes
- `literals.go` — Literal value nodes

#### Node Hierarchy

```
Node (interface)
├── Statement (interface)
│   ├── LetStatement        // dhoro x = 5;
│   ├── ReturnStatement     // ferao x;
│   ├── ExpressionStatement // x + 5;
│   ├── BlockStatement      // { ... }
│   ├── IfStatement         // jodi ... nahole
│   ├── WhileStatement      // jotokkhon
│   ├── ForStatement        // ghuriye
│   ├── ClassStatement      // sreni
│   ├── TryStatement        // chesta ... dhoro_bhul
│   ├── ThrowStatement      // felo
│   ├── ImportStatement     // ano
│   └── ExportStatement     // pathao
│
└── Expression (interface)
    ├── Identifier
    ├── IntegerLiteral
    ├── FloatLiteral
    ├── StringLiteral
    ├── BooleanLiteral
    ├── NullLiteral
    ├── ArrayLiteral
    ├── MapLiteral
    ├── PrefixExpression    // -x, !x, na x
    ├── InfixExpression     // x + y, x ebong y
    ├── IfExpression
    ├── FunctionLiteral     // kaj(x) { ... }
    ├── CallExpression      // fn(x)
    ├── IndexExpression     // arr[0]
    ├── MemberExpression    // obj.prop
    ├── AssignExpression    // x = 5
    └── NewExpression       // notun Class()
```

#### Example AST

For `dhoro x = 5 + 3;`:

```
Program
└── LetStatement
    ├── Name: Identifier("x")
    └── Value: InfixExpression
        ├── Operator: "+"
        ├── Left: IntegerLiteral(5)
        └── Right: IntegerLiteral(3)
```

### 5. Object System (`src/object/`)

The object system represents runtime values.

#### Files
- `object.go` — Object types and interfaces
- `environment.go` — Variable scope management

#### Object Types

```go
type ObjectType string

const (
    INTEGER_OBJ      = "INTEGER"
    FLOAT_OBJ        = "FLOAT"
    STRING_OBJ       = "STRING"
    BOOLEAN_OBJ      = "BOOLEAN"
    NULL_OBJ         = "NULL"
    ARRAY_OBJ        = "ARRAY"
    MAP_OBJ          = "MAP"
    FUNCTION_OBJ     = "FUNCTION"
    BUILTIN_OBJ      = "BUILTIN"
    CLASS_OBJ        = "CLASS"
    INSTANCE_OBJ     = "INSTANCE"
    ERROR_OBJ        = "ERROR"
    RETURN_VALUE_OBJ = "RETURN_VALUE"
    BREAK_OBJ        = "BREAK"
    CONTINUE_OBJ     = "CONTINUE"
    MODULE_OBJ       = "MODULE"
)
```

#### Environment (Scope)

```go
type Environment struct {
    store map[string]Object
    outer *Environment  // Parent scope for closures
}

func (e *Environment) Get(name string) (Object, bool) {
    obj, ok := e.store[name]
    if !ok && e.outer != nil {
        obj, ok = e.outer.Get(name)  // Check parent scope
    }
    return obj, ok
}
```

### 6. Evaluator (`src/evaluator/`)

The evaluator walks the AST and executes the program.

#### Files
- `evaluator.go` — Main evaluation loop
- `expressions.go` — Expression evaluation
- `statements.go` — Statement evaluation
- `builtins.go` — Built-in functions (40+)
- `classes.go` — OOP support
- `modules.go` — Import/export handling
- `errors.go` — Error creation and handling
- `helpers.go` — Utility functions

#### Evaluation Flow

```go
func Eval(node ast.Node, env *object.Environment) object.Object {
    switch node := node.(type) {
    case *ast.Program:
        return evalProgram(node, env)
    case *ast.LetStatement:
        val := Eval(node.Value, env)
        env.Set(node.Name.Value, val)
    case *ast.IfStatement:
        return evalIfStatement(node, env)
    case *ast.InfixExpression:
        left := Eval(node.Left, env)
        right := Eval(node.Right, env)
        return evalInfixExpression(node.Operator, left, right)
    // ... more cases
    }
}
```

#### Built-in Functions

```go
var builtins = map[string]*object.Builtin{
    "dekho": {
        Fn: func(args ...object.Object) object.Object {
            // Print implementation
        },
    },
    "dorghyo": {
        Fn: func(args ...object.Object) object.Object {
            // Length implementation
        },
    },
    // ... 40+ more functions
}
```

### 7. REPL (`src/repl/`)

The Read-Eval-Print Loop for interactive usage.

#### Features
- Multi-line input support
- Command history
- Built-in help system
- Clear screen command
- Graceful exit handling

```go
func Start(in io.Reader, out io.Writer) {
    scanner := bufio.NewScanner(in)
    env := object.NewEnvironment()

    for {
        fmt.Print(">> ")
        if !scanner.Scan() {
            return
        }

        line := scanner.Text()
        l := lexer.New(line)
        p := parser.New(l)
        program := p.ParseProgram()

        result := evaluator.Eval(program, env)
        if result != nil {
            io.WriteString(out, result.Inspect())
        }
    }
}
```

---

## Data Flow

### Complete Execution Pipeline

```
┌──────────────────────────────────────────────────────────────────┐
│                         Source Code                               │
│                    "dhoro x = 5 + 3;"                            │
└─────────────────────────────┬────────────────────────────────────┘
                              │
                              ▼
┌──────────────────────────────────────────────────────────────────┐
│                           Lexer                                   │
│  Converts source text into tokens                                │
│  [DHORO, IDENT(x), ASSIGN, INT(5), PLUS, INT(3), SEMICOLON]     │
└─────────────────────────────┬────────────────────────────────────┘
                              │
                              ▼
┌──────────────────────────────────────────────────────────────────┐
│                          Parser                                   │
│  Builds Abstract Syntax Tree from tokens                         │
│  LetStatement { Name: "x", Value: InfixExpr(5, +, 3) }          │
└─────────────────────────────┬────────────────────────────────────┘
                              │
                              ▼
┌──────────────────────────────────────────────────────────────────┐
│                        Evaluator                                  │
│  Walks AST and executes nodes                                    │
│  1. Evaluate InfixExpr(5, +, 3) → 8                             │
│  2. Bind "x" → 8 in environment                                  │
└─────────────────────────────┬────────────────────────────────────┘
                              │
                              ▼
┌──────────────────────────────────────────────────────────────────┐
│                       Environment                                 │
│  { "x": Integer(8) }                                             │
└──────────────────────────────────────────────────────────────────┘
```

---

## Design Decisions

### 1. Why Go?

| Reason | Benefit |
|--------|---------|
| **Performance** | Native compilation, efficient execution |
| **Simplicity** | Clean syntax, easy to understand |
| **Standard Library** | Rich stdlib reduces dependencies |
| **Concurrency** | Goroutines for HTTP server |
| **Cross-platform** | Easy cross-compilation |
| **Memory Safety** | Garbage collection, no manual memory management |

### 2. Why Tree-Walking Interpreter?

| Advantage | Trade-off |
|-----------|-----------|
| Simplicity | Slower than bytecode compilation |
| Debuggability | Higher memory usage |
| Rapid development | No optimization passes |
| Educational value | Suitable for scripting, not systems programming |

### 3. Why Banglish Keywords?

| Reason | Benefit |
|--------|---------|
| **Accessibility** | Familiar to Bengali speakers |
| **ASCII compatibility** | Works on all keyboards |
| **C-like syntax** | Familiar structure for programming students |
| **Searchability** | Easy to type and search |

### 4. Why Strict Semicolons?

- Teaches discipline and attention to detail
- Matches C/Java syntax students will encounter
- Simpler parser (no automatic semicolon insertion)
- Clear statement boundaries

---

## Performance Considerations

### Current Optimizations

1. **Efficient Token Lookup** — HashMap for keyword identification
2. **Environment Chain** — Fast variable lookup with scope chain
3. **Object Pooling** — Reuse common objects (NULL, TRUE, FALSE)
4. **Lazy Evaluation** — Short-circuit evaluation for `ebong`/`ba`

### Potential Future Optimizations

1. **Bytecode Compilation** — Compile AST to bytecode for faster execution
2. **JIT Compilation** — Just-in-time compilation for hot paths
3. **Constant Folding** — Pre-compute constant expressions
4. **Tail Call Optimization** — Optimize recursive functions
5. **Inline Caching** — Cache property lookups

### Memory Management

- Go's garbage collector handles memory
- Environments are garbage collected when no longer referenced
- Large strings and arrays are heap-allocated

---

## Extending the Language

### Adding a New Built-in Function

1. **Define the function** in `src/evaluator/builtins.go`:

```go
"myFunc": {
    Fn: func(args ...object.Object) object.Object {
        if len(args) != 1 {
            return newError("wrong number of arguments")
        }
        // Implementation
        return &object.String{Value: result}
    },
},
```

2. **Update documentation** in README.md and SYNTAX.md
3. **Add examples** in `examples/` directory

### Adding a New Keyword

1. **Add token** in `src/lexer/token.go`:
```go
const MYKEYWORD = "MYKEYWORD"

var keywords = map[string]TokenType{
    "mykeyword": MYKEYWORD,
}
```

2. **Add AST node** in `src/ast/`:
```go
type MyStatement struct {
    Token token.Token
    // Fields
}
```

3. **Update parser** in `src/parser/`:
```go
case token.MYKEYWORD:
    return p.parseMyStatement()
```

4. **Update evaluator** in `src/evaluator/`:
```go
case *ast.MyStatement:
    return evalMyStatement(node, env)
```

### Adding a New Data Type

1. **Define object type** in `src/object/object.go`:
```go
const MYTYPE_OBJ = "MYTYPE"

type MyType struct {
    Value interface{}
}

func (m *MyType) Type() ObjectType { return MYTYPE_OBJ }
func (m *MyType) Inspect() string  { return fmt.Sprintf("%v", m.Value) }
```

2. **Handle in evaluator** — Add cases for operations with the new type
3. **Add conversion functions** — Built-in functions to create/convert

---

## Directory Structure Summary

```
BanglaCode/
├── main.go                    # Entry point (CLI handling)
├── go.mod                     # Go module definition
├── src/
│   ├── lexer/
│   │   ├── lexer.go          # Tokenization logic
│   │   └── token.go          # Token definitions
│   ├── parser/
│   │   ├── parser.go         # Main parser
│   │   ├── expressions.go    # Expression parsing
│   │   ├── statements.go     # Statement parsing
│   │   └── precedence.go     # Operator precedence
│   ├── ast/
│   │   ├── ast.go            # AST interfaces
│   │   ├── expressions.go    # Expression nodes
│   │   ├── statements.go     # Statement nodes
│   │   └── literals.go       # Literal nodes
│   ├── object/
│   │   ├── object.go         # Runtime types
│   │   └── environment.go    # Scope management
│   ├── evaluator/
│   │   ├── evaluator.go      # Main evaluation
│   │   ├── builtins.go       # Built-in functions
│   │   ├── expressions.go    # Expression evaluation
│   │   ├── statements.go     # Statement evaluation
│   │   ├── classes.go        # OOP support
│   │   ├── modules.go        # Module system
│   │   ├── errors.go         # Error handling
│   │   └── helpers.go        # Utilities
│   └── repl/
│       └── repl.go           # Interactive shell
├── examples/                  # Example programs
├── Extension/                 # VSCode extension
└── Documentation/             # Docs website
```

---

## References

- [Writing An Interpreter In Go](https://interpreterbook.com/) — Thorsten Ball
- [Crafting Interpreters](https://craftinginterpreters.com/) — Robert Nystrom
- [Go Programming Language](https://golang.org/doc/)

---

**Last Updated**: 2024
**Architecture Version**: 3.x
