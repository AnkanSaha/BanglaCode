<p align="center">
  <img src="https://raw.githubusercontent.com/nexoral/BanglaCode/main/Documentation/public/banglacode.svg" alt="BanglaCode Logo" width="200"/>
</p>

<h1 align="center">BanglaCode</h1>

<p align="center">
  <strong>A Bengali-syntax programming language designed to make programming accessible to 300+ million Bengali speakers worldwide.</strong>
</p>

<p align="center">
  <a href="https://github.com/nexoral/BanglaCode/releases"><img src="https://img.shields.io/github/v/release/nexoral/BanglaCode?style=flat-square&color=blue" alt="Release"></a>
  <a href="https://github.com/nexoral/BanglaCode/blob/main/LICENSE"><img src="https://img.shields.io/github/license/nexoral/BanglaCode?style=flat-square&color=green" alt="License"></a>
  <a href="https://github.com/nexoral/BanglaCode/stargazers"><img src="https://img.shields.io/github/stars/nexoral/BanglaCode?style=flat-square" alt="Stars"></a>
  <a href="https://github.com/nexoral/BanglaCode/issues"><img src="https://img.shields.io/github/issues/nexoral/BanglaCode?style=flat-square" alt="Issues"></a>
  <a href="https://github.com/nexoral/BanglaCode/pulls"><img src="https://img.shields.io/github/issues-pr/nexoral/BanglaCode?style=flat-square" alt="Pull Requests"></a>
  <img src="https://img.shields.io/badge/Go-1.20+-00ADD8?style=flat-square&logo=go" alt="Go Version">
  <img src="https://img.shields.io/badge/Platform-Linux%20%7C%20macOS%20%7C%20Windows-lightgrey?style=flat-square" alt="Platform">
</p>

<p align="center">
  <a href="#-quick-start">Quick Start</a> •
  <a href="#-features">Features</a> •
  <a href="#-documentation">Documentation</a> •
  <a href="#-examples">Examples</a> •
  <a href="#-contributing">Contributing</a> •
  <a href="#-community">Community</a>
</p>

---

## Overview

**BanglaCode** is a statically-structured, dynamically-typed programming language that uses Bengali (Banglish) keywords while maintaining familiar C-style syntax. Built with Go for native performance, it bridges the gap between logical thinking and code implementation for Bengali-speaking developers and students.

> *"আমি একজন বাংলা মাধ্যমের ছাত্র। আমি logic তৈরি করতে পারি, কিন্তু সেই logic validate করতে Programming language এর syntax শিখতে হয়। যারা C syntax জানে, তাদের জন্য BanglaCode related হবে। আমি সেই barrier শেষ করতে চেয়েছি — যে ভাষা তুমি জানো, সেই ভাষাতেই logic run করো।"*
>
> — **Ankan**, Creator of BanglaCode

### Why BanglaCode?

| Challenge | Solution |
|-----------|----------|
| Language barrier in programming education | Bengali keywords (`dhoro`, `jodi`, `kaj`) that map directly to programming concepts |
| Slow interpreted languages | Go-powered interpreter with 3-4x faster execution than Python |
| Complex syntax for beginners | C-like structure familiar to students learning programming |
| Limited tooling for regional languages | Full IDE support with VSCode extension, syntax highlighting, and IntelliSense |

---

## Key Features

### Language Capabilities

- **Bengali Keywords** — Write code using familiar Bengali words in English script (Banglish)
- **Object-Oriented Programming** — Full support for classes, constructors, methods, and inheritance
- **Module System** — Import/export functionality for code organization and reusability
- **Error Handling** — Try/catch/finally blocks with custom error throwing
- **HTTP Server** — Built-in web server capabilities similar to Node.js
- **JSON Support** — Native JSON parsing and serialization
- **40+ Built-in Functions** — Comprehensive standard library for strings, arrays, math, files, and more

### Developer Experience

- **Interactive REPL** — Test and experiment with code in real-time
- **VSCode Extension** — Syntax highlighting, IntelliSense, 35+ snippets, and hover documentation
- **Clear Error Messages** — Helpful diagnostics with line and column information
- **Documentation Comments** — `@comment` annotations for IDE tooltips

### Performance

| Metric | BanglaCode | Python | Improvement |
|--------|------------|--------|-------------|
| Startup Time | ~5ms | ~30ms | 6x faster |
| Loop (1M iterations) | ~50ms | ~200ms | 4x faster |
| Memory Usage | Low | Higher | More efficient |
| Recursion | Very Fast | Stack-limited | No GIL limitations |

---

## Quick Start

### Prerequisites

- **Go 1.20** or higher
- **Git** (for cloning the repository)

### Installation

#### From Source

```bash
# Clone the repository
git clone https://github.com/nexoral/BanglaCode.git
cd BanglaCode

# Build the interpreter
go build -o banglacode main.go

# Verify installation
./banglacode --version
```

#### Using Go Install

```bash
go install github.com/nexoral/BanglaCode@latest
```

### Your First Program

Create a file named `hello.bang`:

```banglacode
// @comment: আমার প্রথম BanglaCode প্রোগ্রাম
dhoro naam = "World";
dekho("Namaskar,", naam, "!");

// Variables and arithmetic
dhoro a = 10;
dhoro b = 20;
dekho("Sum:", a + b);
```

Run it:

```bash
./banglacode hello.bang
```

Output:
```
Namaskar, World !
Sum: 30
```

### Interactive REPL

Start the interactive shell to experiment:

```bash
./banglacode
```

```
BanglaCode v3.3.0 - Interactive Mode
Type 'sahajjo' for help, 'baire' to exit

>> dhoro x = 42;
>> dekho(x * 2);
84
>> baire
```

---

## Language Syntax

### Variables and Data Types

```banglacode
// Numbers
dhoro integer = 42;
dhoro decimal = 3.14159;

// Strings
dhoro text = "Hello, BanglaCode!";

// Booleans
dhoro isActive = sotti;    // true
dhoro isDisabled = mittha; // false

// Null
dhoro empty = khali;

// Arrays
dhoro numbers = [1, 2, 3, 4, 5];
dhoro mixed = ["text", 42, sotti];

// Maps/Objects
dhoro person = {
    "naam": "Ankan",
    "boyosh": 25,
    "city": "Kolkata"
};
```

### Control Flow

```banglacode
// If-Else
jodi (score >= 90) {
    dekho("Excellent!");
} nahole jodi (score >= 60) {
    dekho("Good!");
} nahole {
    dekho("Keep trying!");
}

// While Loop
dhoro count = 0;
jotokkhon (count < 5) {
    dekho(count);
    count = count + 1;
}

// For Loop
ghuriye (dhoro i = 0; i < 10; i = i + 1) {
    jodi (i == 5) {
        chharo;  // continue
    }
    dekho(i);
}
```

### Functions

```banglacode
// Function definition
kaj greet(naam) {
    ferao "Namaskar, " + naam + "!";
}

// Function call
dhoro message = greet("Ankan");
dekho(message);

// Recursive function
kaj factorial(n) {
    jodi (n <= 1) {
        ferao 1;
    }
    ferao n * factorial(n - 1);
}

dekho("5! =", factorial(5));  // Output: 5! = 120
```

### Classes and OOP

```banglacode
sreni BankAccount {
    shuru(owner, balance) {
        ei.owner = owner;
        ei.balance = balance;
    }

    kaj deposit(amount) {
        ei.balance = ei.balance + amount;
        dekho("Deposited:", amount);
    }

    kaj withdraw(amount) {
        jodi (amount > ei.balance) {
            dekho("Insufficient funds!");
            ferao mittha;
        }
        ei.balance = ei.balance - amount;
        ferao sotti;
    }

    kaj getBalance() {
        ferao ei.balance;
    }
}

dhoro account = notun BankAccount("Ankan", 1000);
account.deposit(500);
dekho("Balance:", account.getBalance());  // Output: Balance: 1500
```

### Modules

```banglacode
// math_utils.bang
pathao kaj add(a, b) {
    ferao a + b;
}

pathao kaj multiply(a, b) {
    ferao a * b;
}

// main.bang
ano "math_utils.bang";

dekho(add(5, 3));       // Output: 8
dekho(multiply(4, 7));  // Output: 28

// Import with alias
ano "math_utils.bang" hisabe math;
dekho(math.add(10, 20));  // Output: 30
```

### Error Handling

```banglacode
kaj divide(a, b) {
    jodi (b == 0) {
        felo "Division by zero error!";
    }
    ferao a / b;
}

chesta {
    dhoro result = divide(10, 0);
    dekho(result);
} dhoro_bhul (err) {
    dekho("Error caught:", err);
} shesh {
    dekho("Operation completed.");
}
```

### HTTP Server

```banglacode
kaj handleRequest(req, res) {
    jodi (req.path == "/") {
        uttor(res, "Welcome to BanglaCode Server!");
    } nahole jodi (req.path == "/api/data") {
        dhoro data = {
            "status": "success",
            "message": "Namaskar from BanglaCode!"
        };
        json_uttor(res, data);
    } nahole {
        uttor(res, "Not Found", 404);
    }
}

dekho("Server starting on port 3000...");
server_chalu(3000, handleRequest);
```

---

## Keywords Reference

| Keyword | Bengali | English Equivalent | Example |
|---------|---------|-------------------|---------|
| `dhoro` | ধরো | let/var | `dhoro x = 5;` |
| `jodi` | যদি | if | `jodi (x > 0) { }` |
| `nahole` | নাহলে | else | `nahole { }` |
| `jotokkhon` | যতক্ষণ | while | `jotokkhon (x < 10) { }` |
| `ghuriye` | ঘুরিয়ে | for | `ghuriye (dhoro i = 0; i < 5; i = i + 1) { }` |
| `kaj` | কাজ | function | `kaj add(a, b) { }` |
| `ferao` | ফেরাও | return | `ferao result;` |
| `sreni` | শ্রেণী | class | `sreni Person { }` |
| `shuru` | শুরু | constructor | `shuru(naam) { }` |
| `notun` | নতুন | new | `notun Person()` |
| `ei` | এই | this | `ei.naam = "Ankan";` |
| `sotti` | সত্যি | true | `dhoro flag = sotti;` |
| `mittha` | মিথ্যা | false | `dhoro flag = mittha;` |
| `khali` | খালি | null | `dhoro val = khali;` |
| `ebong` | এবং | and (&&) | `jodi (a ebong b) { }` |
| `ba` | বা | or (\|\|) | `jodi (a ba b) { }` |
| `na` | না | not (!) | `jodi (na flag) { }` |
| `thamo` | থামো | break | `thamo;` |
| `chharo` | ছাড়ো | continue | `chharo;` |
| `ano` | আনো | import | `ano "module.bang";` |
| `hisabe` | হিসাবে | as (alias) | `ano "x.bang" hisabe y;` |
| `pathao` | পাঠাও | export | `pathao kaj fn() { }` |
| `chesta` | চেষ্টা | try | `chesta { }` |
| `dhoro_bhul` | ধরো ভুল | catch | `dhoro_bhul (e) { }` |
| `shesh` | শেষ | finally | `shesh { }` |
| `felo` | ফেলো | throw | `felo "error";` |

---

## Built-in Functions

### Output
| Function | Bengali | Description |
|----------|---------|-------------|
| `dekho(...)` | দেখো | Print values to console |

### Type Operations
| Function | Bengali | Description |
|----------|---------|-------------|
| `dhoron(x)` | ধরন | Get type of value |
| `lipi(x)` | লিপি | Convert to string |
| `sonkha(x)` | সংখ্যা | Convert to number |
| `dorghyo(x)` | দৈর্ঘ্য | Get length of string/array |

### String Functions
| Function | Bengali | Description |
|----------|---------|-------------|
| `boroHater(str)` | বড় হাতের | Convert to uppercase |
| `chotoHater(str)` | ছোট হাতের | Convert to lowercase |
| `chhanto(str)` | ছাঁটো | Trim whitespace |
| `bhag(str, sep)` | ভাগ | Split string into array |
| `joro(arr, sep)` | জোড়ো | Join array into string |
| `khojo(str, substr)` | খোঁজো | Find substring index |
| `angsho(str, start, end)` | অংশ | Extract substring |
| `bodlo(str, old, new)` | বদলো | Replace substring |

### Array Functions
| Function | Bengali | Description |
|----------|---------|-------------|
| `dhokao(arr, val)` | ঢোকাও | Push element to array |
| `berKoro(arr)` | বের করো | Pop last element |
| `kato(arr, start, end)` | কাটো | Slice array |
| `ulto(arr)` | উল্টো | Reverse array |
| `saja(arr)` | সাজা | Sort array |
| `ache(arr, val)` | আছে | Check if contains |

### Math Functions
| Function | Bengali | Description |
|----------|---------|-------------|
| `borgomul(x)` | বর্গমূল | Square root |
| `ghat(base, exp)` | ঘাত | Power |
| `niche(x)` | নিচে | Floor |
| `upore(x)` | উপরে | Ceiling |
| `kache(x)` | কাছে | Round |
| `niratek(x)` | নিরপেক্ষ | Absolute value |
| `choto(...)` | ছোট | Minimum |
| `boro(...)` | বড় | Maximum |
| `lotto()` | লটো | Random (0-1) |

### File I/O
| Function | Bengali | Description |
|----------|---------|-------------|
| `poro(path)` | পড়ো | Read file contents |
| `lekho(path, content)` | লেখো | Write to file |

### JSON
| Function | Bengali | Description |
|----------|---------|-------------|
| `json_poro(str)` | JSON পড়ো | Parse JSON string |
| `json_banao(obj)` | JSON বানাও | Convert to JSON string |

### HTTP
| Function | Bengali | Description |
|----------|---------|-------------|
| `server_chalu(port, handler)` | সার্ভার চালু | Start HTTP server |
| `anun(url)` | আনুন | HTTP GET request |
| `uttor(res, body, status, type)` | উত্তর | Send response |
| `json_uttor(res, data, status)` | JSON উত্তর | Send JSON response |

### Utility
| Function | Bengali | Description |
|----------|---------|-------------|
| `somoy()` | সময় | Current timestamp (ms) |
| `ghum(ms)` | ঘুম | Sleep for milliseconds |
| `nao(prompt)` | নাও | Read user input |
| `bondho(code)` | বন্ধ | Exit program |

---

## Documentation

| Document | Description |
|----------|-------------|
| [SYNTAX.md](SYNTAX.md) | Complete language syntax reference |
| [CONTRIBUTING.md](CONTRIBUTING.md) | Contribution guidelines |
| [CODE_OF_CONDUCT.md](CODE_OF_CONDUCT.md) | Community code of conduct |
| [SECURITY.md](SECURITY.md) | Security policy |
| [CHANGELOG.md](CHANGELOG.md) | Version history |
| [ARCHITECTURE.md](ARCHITECTURE.md) | Technical architecture |
| [ROADMAP.md](ROADMAP.md) | Future development plans |
| [SUPPORT.md](SUPPORT.md) | Getting help |

---

## Examples

The `examples/` directory contains comprehensive sample programs:

| File | Description |
|------|-------------|
| `hello.bang` | Hello world and variables |
| `functions.bang` | Function definitions and recursion |
| `classes.bang` | Object-oriented programming |
| `loops.bang` | While and for loops |
| `data_structures.bang` | Arrays and maps |
| `modules_demo.bang` | Import/export system |
| `error_handling.bang` | Try/catch/finally |
| `http_server.bang` | Web server example |
| `builtins_demo.bang` | Built-in functions |
| `json.bang` | JSON handling |

Run any example:

```bash
./banglacode examples/functions.bang
```

---

## IDE Support

### VSCode Extension

Install the official BanglaCode extension for Visual Studio Code:

**Features:**
- Syntax highlighting for `.bang` files
- IntelliSense and auto-completion
- 35+ code snippets
- Hover documentation
- Custom file icons
- Error highlighting

**Installation:**

1. Open VSCode
2. Go to Extensions (Ctrl+Shift+X)
3. Search for "BanglaCode"
4. Click Install

Or install from VSIX:

```bash
cd Extension
npm install
npx vsce package
# Install the generated .vsix file
```

---

## Project Structure

```
BanglaCode/
├── main.go                 # Entry point and CLI
├── go.mod                  # Go module definition
├── VERSION                 # Version file
├── LICENSE                 # GPL-3.0 License
│
├── src/                    # Core interpreter
│   ├── lexer/              # Tokenization
│   │   ├── lexer.go        # Scanner implementation
│   │   └── token.go        # Token definitions
│   ├── parser/             # Syntax analysis
│   │   ├── parser.go       # Parser implementation
│   │   ├── expressions.go  # Expression parsing
│   │   ├── statements.go   # Statement parsing
│   │   └── precedence.go   # Operator precedence
│   ├── ast/                # Abstract Syntax Tree
│   │   ├── ast.go          # AST base
│   │   ├── expressions.go  # Expression nodes
│   │   ├── statements.go   # Statement nodes
│   │   └── literals.go     # Literal nodes
│   ├── object/             # Runtime values
│   │   ├── object.go       # Object types
│   │   └── environment.go  # Variable scopes
│   ├── evaluator/          # Interpreter
│   │   ├── evaluator.go    # Core evaluation
│   │   ├── builtins.go     # Built-in functions
│   │   ├── expressions.go  # Expression evaluation
│   │   ├── statements.go   # Statement evaluation
│   │   ├── classes.go      # OOP support
│   │   ├── modules.go      # Module system
│   │   └── errors.go       # Error handling
│   └── repl/               # Interactive shell
│       └── repl.go         # REPL implementation
│
├── examples/               # Example programs
├── Extension/              # VSCode extension
├── Documentation/          # Documentation website
├── Scripts/                # Build scripts
└── .github/                # CI/CD workflows
```

---

## Contributing

We welcome contributions from the community! See [CONTRIBUTING.md](CONTRIBUTING.md) for detailed guidelines.

### Quick Start for Contributors

```bash
# Fork and clone
git clone https://github.com/YOUR_USERNAME/BanglaCode.git
cd BanglaCode

# Create feature branch
git checkout -b feature/my-feature

# Make changes and test
go build -o banglacode main.go
./banglacode examples/hello.bang

# Commit and push
git commit -m "feat: add my feature"
git push origin feature/my-feature
```

### Priority Areas

- Performance optimization
- Additional built-in functions
- Better error messages
- Bengali documentation and tutorials
- Testing infrastructure
- Online playground

---

## Community

- **GitHub Issues**: [Bug reports and feature requests](https://github.com/nexoral/BanglaCode/issues)
- **GitHub Discussions**: [Questions and community chat](https://github.com/nexoral/BanglaCode/discussions)
- **Documentation**: [https://banglacode.dev](https://banglacode.dev)

---

## License

BanglaCode is open source software licensed under the [GNU General Public License v3.0](LICENSE).

---

## Acknowledgments

BanglaCode draws inspiration from:
- **C** — Strict syntax discipline
- **JavaScript** — Modern language features
- **Go** — Performance and simplicity
- **The Bengali-speaking community** — Making programming accessible to 300+ million speakers

---

## Author

<table>
  <tr>
    <td align="center">
      <strong>Ankan</strong><br/>
      Creator & Lead Developer<br/>
      West Bengal, India<br/><br/>
      <em>"Programming should be about logic, not language barriers."</em>
    </td>
  </tr>
</table>

---

<p align="center">
  <strong>আপনার প্রোগ্রামিং যাত্রা শুভ হোক!</strong><br/>
  <em>May your programming journey be successful!</em>
</p>

<p align="center">
  Made with care from West Bengal, India
</p>

<p align="center">
  <a href="#banglacode">Back to top</a>
</p>
