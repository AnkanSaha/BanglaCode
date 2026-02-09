# Contributing to BanglaCode

First off, thank you for considering contributing to BanglaCode!

BanglaCode is a community-driven project that aims to make programming accessible to Bengali speakers. Every contribution, no matter how small, helps improve the language and its ecosystem.

## Table of Contents

- [How Can I Contribute?](#-how-can-i-contribute)
- [Reporting Bugs](#reporting-bugs)
- [Suggesting Enhancements](#suggesting-enhancements)
- [Pull Requests](#pull-requests)
- [Development Setup](#%EF%B8%8F-development-setup)
- [Adding New Features](#-adding-new-features)
- [Testing](#-testing)
- [Documentation Standards](#-documentation-standards)
- [Code of Conduct](#-code-of-conduct)

## Quick Links

| Resource | Description |
|----------|-------------|
| [CODE_OF_CONDUCT.md](CODE_OF_CONDUCT.md) | Community guidelines |
| [ARCHITECTURE.md](ARCHITECTURE.md) | Technical architecture |
| [ROADMAP.md](ROADMAP.md) | Future development plans |
| [SECURITY.md](SECURITY.md) | Security policy |
| [SUPPORT.md](SUPPORT.md) | Getting help |

## 🌟 How Can I Contribute?

### Reporting Bugs

Before creating bug reports, please check the existing issues to avoid duplicates. When you create a bug report, include as many details as possible:

- **Use a clear and descriptive title**
- **Describe the exact steps to reproduce the problem**
- **Provide specific examples** (code snippets, error messages)
- **Describe the behavior you observed** and what you expected
- **Include your environment details** (OS, Go version, BanglaCode version)

**Bug Report Template:**

```markdown
**Environment:**
- OS: [e.g., Ubuntu 22.04, Windows 11, macOS 14]
- Go Version: [e.g., 1.21.5]
- BanglaCode Version: [e.g., 3.3.0]

**Description:**
A clear description of the bug.

**Steps to Reproduce:**
1. Create file with code: `...`
2. Run `./banglacode file.bang`
3. See error

**Expected Behavior:**
What you expected to happen.

**Actual Behavior:**
What actually happened.

**Code Sample:**
```banglacode
dhoro x = 5;
// ... minimal reproducible example
```

**Error Message:**
```
[paste error message here]
```
```

### Suggesting Enhancements

Enhancement suggestions are welcome! This includes:

- **New built-in functions**
- **Language features** (new keywords, operators, syntax)
- **Performance improvements**
- **Better error messages**
- **Documentation improvements**
- **VSCode extension features**

When suggesting an enhancement:

1. **Check if it's already suggested** in existing issues
2. **Provide a clear use case** - explain why this enhancement is useful
3. **Consider backward compatibility** - will this break existing code?
4. **Propose an implementation** - if you have ideas on how to implement it

**Enhancement Template:**

```markdown
**Feature Description:**
Brief description of the feature.

**Use Case:**
Why is this feature needed? What problem does it solve?

**Proposed Syntax/Behavior:**
```banglacode
// Example code showing how the feature would work
```

**Alternatives Considered:**
Other ways this could be implemented.
```

### Pull Requests

We actively welcome your pull requests! Here's the process:

1. **Fork the repository** and create your branch from `main`
2. **Make your changes** following our coding standards
3. **Test your changes** thoroughly
4. **Update documentation** if needed
5. **Write clear commit messages**
6. **Submit a pull request**

#### Pull Request Guidelines

- **One feature per PR** - keep PRs focused and reviewable
- **Write tests** for new functionality
- **Update SYNTAX.md** if you add/change language features
- **Update README.md** if needed
- **Follow Go conventions** and existing code style
- **Add examples** for new features in the `examples/` directory

#### Commit Message Format

Use clear, descriptive commit messages:

```
feat: add string interpolation support
fix: correct error line number reporting
docs: update CONTRIBUTING.md with PR guidelines
test: add tests for array slicing
refactor: simplify lexer token handling
perf: optimize recursive function calls
```

**Prefixes:**
- `feat:` - New feature
- `fix:` - Bug fix
- `docs:` - Documentation changes
- `test:` - Adding/updating tests
- `refactor:` - Code refactoring
- `perf:` - Performance improvements
- `chore:` - Maintenance tasks

## 🏗️ Development Setup

### Prerequisites

- **Go 1.20 or higher**
- **Git**
- **A code editor** (VSCode recommended)

### Getting Started

```bash
# Clone your fork
git clone https://github.com/YOUR_USERNAME/BanglaCode.git
cd BanglaCode

# Add upstream remote
git remote add upstream https://github.com/nexoral/BanglaCode.git

# Create a new branch
git checkout -b feature/my-new-feature

# Build the project
go build -o banglacode main.go

# Run tests (if available)
go test ./...

# Test your changes
./banglacode examples/hello.bang
```

### Project Structure

```
BanglaCode/
├── main.go              # Entry point
├── go.mod               # Go module file
├── src/
│   ├── lexer/          # Tokenization (lexer.go, token.go)
│   ├── parser/         # Parsing (parser.go)
│   ├── ast/            # Abstract Syntax Tree (ast.go)
│   ├── object/         # Runtime objects (object.go, environment.go)
│   ├── evaluator/      # Interpreter (evaluator.go, builtins.go)
│   └── repl/           # Interactive shell (repl.go)
├── examples/           # Example programs
├── Extension/          # VSCode extension
└── Documentation/      # Documentation website
```

### Key Components

#### 1. Lexer (`src/lexer/`)
- **Purpose:** Converts source code into tokens
- **Files:** `lexer.go` (tokenizer), `token.go` (token definitions)
- **Example:** `dhoro x = 5;` → `[DHORO, IDENTIFIER(x), ASSIGN, NUMBER(5), SEMICOLON]`

#### 2. Parser (`src/parser/`)
- **Purpose:** Converts tokens into an Abstract Syntax Tree (AST)
- **Files:** `parser.go`
- **Example:** Tokens → AST nodes (expressions, statements)

#### 3. AST (`src/ast/`)
- **Purpose:** Represents program structure as a tree
- **Files:** `ast.go`
- **Nodes:** Program, LetStatement, FunctionLiteral, etc.

#### 4. Evaluator (`src/evaluator/`)
- **Purpose:** Executes the AST (interpreter)
- **Files:** `evaluator.go` (core), `builtins.go` (built-in functions)
- **Example:** Walks AST and executes code

#### 5. Object System (`src/object/`)
- **Purpose:** Runtime value representation
- **Files:** `object.go` (value types), `environment.go` (variable scopes)
- **Types:** Integer, String, Boolean, Array, Function, etc.

#### 6. REPL (`src/repl/`)
- **Purpose:** Interactive shell
- **Files:** `repl.go`

## 🔧 Adding New Features

### Adding a New Built-in Function

1. **Define the function** in `src/evaluator/builtins.go`:

```go
"myFunc": {
    Fn: func(args ...object.Object) object.Object {
        // Validate arguments
        if len(args) != 1 {
            return newError("wrong number of arguments. got=%d, want=1", len(args))
        }
        
        // Implement logic
        // ...
        
        return &object.String{Value: result}
    },
},
```

2. **Update documentation** in README.md and SYNTAX.md
3. **Add an example** in `examples/builtins_demo.bang`
4. **Test it** in the REPL

### Adding a New Keyword

1. **Add token type** in `src/lexer/token.go`:

```go
const (
    // ... existing tokens
    MYNEWKEYWORD = "MYNEWKEYWORD"
)

var keywords = map[string]TokenType{
    // ... existing keywords
    "mynewkeyword": MYNEWKEYWORD,
}
```

2. **Update parser** in `src/parser/parser.go` to recognize and parse the new syntax
3. **Update evaluator** in `src/evaluator/evaluator.go` to handle the new AST node
4. **Add tests** and examples
5. **Update documentation**

### Adding a New Data Type

1. **Define object type** in `src/object/object.go`:

```go
type MyNewType struct {
    Value interface{}
}

func (m *MyNewType) Type() ObjectType { return MYNEWTYPE_OBJ }
func (m *MyNewType) Inspect() string  { return fmt.Sprintf("%v", m.Value) }
```

2. **Update evaluator** to handle the new type
3. **Add conversion/manipulation functions**
4. **Document the new type**

## 🧪 Testing

While BanglaCode doesn't have extensive automated tests yet, you should manually test:

1. **Run existing examples:** `./banglacode examples/*.bang`
2. **Test in REPL:** Try various edge cases interactively
3. **Test your specific changes** with custom scripts
4. **Check error messages** are clear and helpful

### Testing Checklist

- [ ] Feature works as expected
- [ ] Edge cases handled (null, empty arrays, etc.)
- [ ] Error messages are clear
- [ ] No crashes or panics
- [ ] Documentation is updated
- [ ] Examples added/updated

## 📝 Documentation Standards

### Code Comments

- **Document exported functions** with clear descriptions
- **Explain complex logic** with inline comments
- **Use `@comment` tags** in BanglaCode source for VSCode tooltips

Example:

```go
// EvalProgram evaluates a parsed program by walking the AST
// and executing each statement sequentially.
func EvalProgram(program *ast.Program, env *object.Environment) object.Object {
    // ...
}
```

### Markdown Documentation

- **Use clear headings** and structure
- **Include code examples** for every feature
- **Provide both English and Bengali** explanations where appropriate
- **Keep language simple** - remember the target audience (students)

## 🌍 Internationalization

BanglaCode aims to be accessible to Bengali speakers:

- **Keywords** should have clear Bengali meanings
- **Error messages** should be helpful (consider Bengali versions in future)
- **Documentation** should include Bengali explanations
- **Examples** should use Bengali variable names and contexts

## Code of Conduct

Please read our full [Code of Conduct](CODE_OF_CONDUCT.md) before contributing.

### Summary

- **Be respectful** and welcoming
- **Accept constructive criticism** gracefully
- **Focus on what's best** for the community
- **Show empathy** towards other community members

Harassment, discrimination, or offensive behavior will not be tolerated.

## 🎯 Priority Areas

We especially welcome contributions in these areas:

1. **Performance Optimization** - Make BanglaCode faster
2. **Error Messages** - More helpful and descriptive errors
3. **Standard Library** - More built-in functions
4. **Documentation** - Tutorials, guides, Bengali translations
5. **VSCode Extension** - Better IDE support
6. **Examples** - More real-world example programs
7. **Testing** - Automated test suite
8. **Website/Playground** - Online interpreter

## Getting Help

- **Questions?** Open a GitHub Discussion or check [SUPPORT.md](SUPPORT.md)
- **Need guidance?** Comment on an issue
- **Architecture questions?** See [ARCHITECTURE.md](ARCHITECTURE.md)
- **Security concerns?** See [SECURITY.md](SECURITY.md)

## Contributor Recognition

All contributors are recognized in [AUTHORS.md](AUTHORS.md). Significant contributions may lead to:

- Core Contributor status
- Maintainer role (see [GOVERNANCE.md](GOVERNANCE.md))
- Featured on the documentation website

## License

By contributing, you agree that your contributions will be licensed under the GNU General Public License v3.0.

---

## Additional Resources

| Document | Description |
|----------|-------------|
| [ARCHITECTURE.md](ARCHITECTURE.md) | How BanglaCode works internally |
| [ROADMAP.md](ROADMAP.md) | Where the project is heading |
| [GOVERNANCE.md](GOVERNANCE.md) | Project governance and decision-making |
| [AUTHORS.md](AUTHORS.md) | Contributors and maintainers |

---

**Thank you for helping make programming accessible to Bengali speakers!**

**আপনার অবদানের জন্য ধন্যবাদ!**

*Made with care from West Bengal, India*
