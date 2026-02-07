# BanglaCode Syntax Guide

**BanglaCode** is a Bengali-syntax programming language that makes programming accessible to Bengali speakers. Created by **Ankan** from **West Bengal, India**, it uses Banglish (Bengali words written in English) keywords and combines C-like strict syntax with modern features.

## 🎯 Design Philosophy

BanglaCode is designed for **students** who:
- Know C syntax and want something familiar
- Want strict syntax discipline (semicolons required)
- Need a fast language (not Python-slow)
- Think in Bengali but code in English

## Table of Contents

- [Installation](#installation)
- [Quick Start](#quick-start)
- [Keywords](#keywords)
- [Data Types](#data-types)
- [Variables](#variables)
- [Operators](#operators)
- [Control Flow](#control-flow)
- [Loops](#loops)
- [Functions](#functions)
- [Classes and OOP](#classes-and-oop)
- [Modules (Import/Export)](#modules-importexport)
- [Error Handling](#error-handling)
- [HTTP Server](#http-server)
- [Arrays](#arrays)
- [Maps/Objects](#mapsobjects)
- [Built-in Functions](#built-in-functions)
- [Comments](#comments)
- [Examples](#examples)

## Installation

### Prerequisites
- Go 1.20 or higher

### Building from Source

```bash
cd /path/to/BanglaCode
go build -o banglacode main.go
```

### Running

```bash
# Start REPL (interactive mode)
./banglacode

# Run a file
./banglacode script.bang

# Or use go run
go run main.go
go run main.go examples/hello.bang
```

## Quick Start

Create a file `hello.bang`:

```banglacode
dekho("Hello, West Bengal!");
```

Run it:
```bash
./banglacode hello.bang
```

## Keywords

BanglaCode uses Banglish keywords that are intuitive for Bengali speakers:

| Keyword | Meaning | English Equivalent |
|---------|---------|-------------------|
| `dhoro` | hold/take | let/var |
| `jodi` | if | if |
| `nahole` | else | else |
| `jotokkhon` | as long as | while |
| `ghuriye` | rotate/turn | for |
| `kaj` | work/function | function |
| `ferao` | return | return |
| `class` | class | class |
| `notun` | new | new |
| `sotti` | truth | true |
| `mittha` | false | false |
| `khali` | empty | null/nil |
| `ebong` | and | && |
| `ba` | or | \|\| |
| `na` | not | ! |
| `thamo` | stop | break |
| `chharo` | leave | continue |
| `dekho` | see/show | print |
| `ei` | this | this |
| `ano` | import | import |
| `pathao` | export | export |
| `chesta` | try | try |
| `dhoro_bhul` | catch error | catch |
| `shesh` | finally/end | finally |
| `felo` | throw | throw |

## Data Types

BanglaCode supports the following data types:

### Numbers
Both integers and floating-point numbers:
```banglacode
dhoro age = 25;
dhoro price = 99.99;
dhoro pi = 3.14159;
```

### Strings
Enclosed in double or single quotes:
```banglacode
dhoro naam = "West Bengal";
dhoro message = 'Namaskar';
```

### Booleans
```banglacode
dhoro isTrue = sotti;
dhoro isFalse = mittha;
```

### Null
```banglacode
dhoro empty = khali;
```

### Arrays
```banglacode
dhoro numbers = [1, 2, 3, 4, 5];
dhoro fruits = ["Aam", "Kathal", "Lichu"];
dhoro mixed = [1, "hello", sotti, khali];
```

### Maps/Objects
```banglacode
dhoro person = {
    "naam": "Ankan",
    "boyosh": 25,
    "city": "Kolkata"
};
```

## Variables

Declare variables using `dhoro`:

```banglacode
dhoro naam = "Ankan";
dhoro boyosh = 25;
dhoro isStudent = sotti;
```

Variables are dynamically typed - no type declarations needed!

## Operators

### Arithmetic Operators
```banglacode
dhoro sum = 5 + 3;        // Addition
dhoro diff = 10 - 4;      // Subtraction
dhoro product = 6 * 7;    // Multiplication
dhoro quotient = 20 / 4;  // Division
dhoro remainder = 10 % 3; // Modulo
```

### Comparison Operators
```banglacode
5 == 5    // Equal to
5 != 3    // Not equal to
5 < 10    // Less than
10 > 5    // Greater than
5 <= 5    // Less than or equal to
10 >= 5   // Greater than or equal to
```

### Logical Operators
```banglacode
sotti ebong mittha   // AND (&&)
sotti ba mittha      // OR (||)
na sotti             // NOT (!)
```

### Assignment Operators
```banglacode
dhoro x = 10;
x = x + 5;    // or use:
x += 5;       // Compound addition
x -= 3;       // Compound subtraction
x *= 2;       // Compound multiplication
x /= 2;       // Compound division
```

## Control Flow

### If Statement
```banglacode
dhoro age = 20;

jodi (age >= 18) {
    dekho("Adult");
}
```

### If-Else Statement
```banglacode
jodi (age >= 18) {
    dekho("Adult");
} nahole {
    dekho("Minor");
}
```

### If-Else If-Else
```banglacode
dhoro marks = 85;

jodi (marks >= 90) {
    dekho("A+");
} nahole jodi (marks >= 80) {
    dekho("A");
} nahole jodi (marks >= 70) {
    dekho("B");
} nahole {
    dekho("C");
}
```

## Loops

### While Loop (`jotokkhon`)
```banglacode
dhoro i = 0;
jotokkhon (i < 5) {
    dekho(i);
    i = i + 1;
}
```

### For Loop (`ghuriye`)
```banglacode
ghuriye (dhoro i = 0; i < 5; i = i + 1) {
    dekho("Count:", i);
}
```

### Break (`thamo`)
```banglacode
dhoro i = 0;
jotokkhon (i < 10) {
    jodi (i == 5) {
        thamo;  // Exit loop
    }
    dekho(i);
    i = i + 1;
}
```

### Continue (`chharo`)
```banglacode
ghuriye (dhoro i = 0; i < 5; i = i + 1) {
    jodi (i == 2) {
        chharo;  // Skip this iteration
    }
    dekho(i);
}
```

## Functions

### Defining Functions
```banglacode
kaj greet(naam) {
    dekho("Hello,", naam);
}

greet("Ankan");
```

### Functions with Return Values
```banglacode
kaj add(a, b) {
    ferao a + b;
}

dhoro result = add(5, 3);
dekho("5 + 3 =", result);  // Output: 5 + 3 = 8
```

### Multiple Parameters
```banglacode
kaj calculate(x, y, z) {
    ferao (x + y) * z;
}

dekho(calculate(2, 3, 4));  // Output: 20
```

### Recursive Functions
```banglacode
kaj factorial(n) {
    jodi (n <= 1) {
        ferao 1;
    }
    ferao n * factorial(n - 1);
}

dekho("5! =", factorial(5));  // Output: 5! = 120
```

## Classes and OOP

### Defining Classes
```banglacode
class Manush {
    kaj init(naam, boyosh) {
        ei.naam = naam;
        ei.boyosh = boyosh;
    }

    kaj porichoy() {
        dekho("Amar naam", ei.naam);
    }
}
```

### Creating Instances
```banglacode
dhoro person = notun Manush("Ankan", 25);
person.porichoy();  // Output: Amar naam Ankan
```

### Accessing Properties
```banglacode
dekho(person.naam);    // Output: Ankan
dekho(person.boyosh);  // Output: 25
```

### Modifying Properties
```banglacode
person.boyosh = 26;
dekho(person.boyosh);  // Output: 26
```

### Complete Class Example
```banglacode
class Rectangle {
    kaj init(width, height) {
        ei.width = width;
        ei.height = height;
    }

    kaj area() {
        ferao ei.width * ei.height;
    }

    kaj perimeter() {
        ferao 2 * (ei.width + ei.height);
    }
}

dhoro rect = notun Rectangle(10, 5);
dekho("Area:", rect.area());           // Output: Area: 50
dekho("Perimeter:", rect.perimeter()); // Output: Perimeter: 30
```

## Modules (Import/Export)

BanglaCode supports a powerful module system for organizing code into reusable files.

### Exporting Functions and Classes

Use `pathao` (export) to make functions, classes, or variables available to other files:

```banglacode
// math_utils.bang

pathao kaj add(a, b) {
    ferao a + b;
}

pathao kaj multiply(a, b) {
    ferao a * b;
}

pathao class Calculator {
    kaj init() {
        ei.result = 0;
    }
    
    kaj add(n) {
        ei.result = ei.result + n;
        ferao ei;
    }
}
```

### Importing Modules

Use `ano` (import) to load modules:

```banglacode
// main.bang

// Import all exports directly into namespace
ano "math_utils.bang";

dekho(add(5, 3));      // Output: 8
dekho(multiply(4, 7)); // Output: 28

dhoro calc = notun Calculator();
calc.add(10).add(5);
dekho(calc.result);    // Output: 15
```

### Import with Alias

```banglacode
// Import as namespace (like Python's import x as y)
ano "math_utils.bang" as math;

dekho(math["add"](5, 3));  // Output: 8
```

## Error Handling

BanglaCode provides structured error handling with `chesta`/`dhoro_bhul`/`shesh` (try/catch/finally).

### Basic Try/Catch

```banglacode
chesta {
    // Code that might throw an error
    felo "Something went wrong!";
} dhoro_bhul (err) {
    // Handle the error
    dekho("Caught error:", err);
}
```

### Try/Catch/Finally

```banglacode
chesta {
    dekho("Opening file...");
    felo "File not found!";
} dhoro_bhul (err) {
    dekho("Error:", err);
} shesh {
    // Always runs, even if error occurred
    dekho("Cleanup: Closing file");
}
```

### Throwing Errors

Use `felo` (throw) to raise an error:

```banglacode
kaj divide(a, b) {
    jodi (b == 0) {
        felo "Division by zero!";
    }
    ferao a / b;
}

chesta {
    dekho(divide(10, 0));
} dhoro_bhul (err) {
    dekho("Error:", err);
}
```

### Safe Function Pattern

```banglacode
kaj safeDivide(a, b) {
    chesta {
        jodi (b == 0) {
            felo "Cannot divide by zero";
        }
        ferao a / b;
    } dhoro_bhul (err) {
        dekho("Warning:", err);
        ferao khali;
    }
}

dekho(safeDivide(10, 2));  // Output: 5
dekho(safeDivide(10, 0));  // Output: Warning: Cannot divide by zero, then khali
```

## HTTP Server

BanglaCode can create HTTP servers like Node.js, making it easy to build web applications.

### Basic Server

```banglacode
kaj handleRequest(req, res) {
    res["body"] = "Hello from BanglaCode!";
}

server_chalu(3000, handleRequest);
```

### Request Object

The request handler receives a `req` object with:
- `req["method"]` - HTTP method (GET, POST, etc.)
- `req["path"]` - URL path
- `req["query"]` - Query string
- `req["headers"]` - Request headers
- `req["body"]` - Request body

### Response Object

Set response properties:
- `res["status"]` - HTTP status code (default: 200)
- `res["body"]` - Response body
- `res["headers"]` - Response headers

### Full Server Example

```banglacode
kaj handleRequest(req, res) {
    dekho("Request:", req["method"], req["path"]);
    
    jodi (req["path"] == "/") {
        res["body"] = "Welcome to BanglaCode Server!";
        res["headers"]["Content-Type"] = "text/plain";
    } nahole jodi (req["path"] == "/api/hello") {
        res["body"] = "{\"message\": \"Namaskar!\"}";
        res["headers"]["Content-Type"] = "application/json";
    } nahole {
        res["status"] = 404;
        res["body"] = "Not Found";
    }
}

server_chalu(3000, handleRequest);
```

### HTTP Client

Make HTTP requests:

```banglacode
dhoro response = anun("https://api.example.com/data");
dekho("Status:", response["status"]);
dekho("Body:", response["body"]);
```

## Arrays

### Creating Arrays
```banglacode
dhoro fruits = ["Aam", "Kathal", "Lichu"];
dhoro numbers = [1, 2, 3, 4, 5];
```

### Accessing Elements
```banglacode
dekho(fruits[0]);  // Output: Aam
dekho(fruits[2]);  // Output: Lichu
```

### Modifying Elements
```banglacode
fruits[1] = "Kola";
dekho(fruits[1]);  // Output: Kola
```

### Array Operations
```banglacode
// Get length
dhoro len = dorghyo(fruits);

// Add element
dhokao(fruits, "Peyara");

// Remove last element
dhoro last = berKoro(fruits);

// Iterate over array
ghuriye (dhoro i = 0; i < dorghyo(fruits); i = i + 1) {
    dekho(fruits[i]);
}
```

## Maps/Objects

### Creating Maps
```banglacode
dhoro person = {
    "naam": "Ankan",
    "boyosh": 25,
    "city": "Kolkata"
};
```

### Accessing Values
```banglacode
dekho(person["naam"]);  // Output: Ankan
dekho(person["city"]);  // Output: Kolkata
```

### Modifying Values
```banglacode
person["boyosh"] = 26;
person["country"] = "West Bengal";
```

### Getting Keys
```banglacode
dhoro sobChabi = chabi(person);
dekho(sobChabi);  // Output: ["naam", "boyosh", "city", "country"]
```

### Iterating Over Maps
```banglacode
dhoro personKeys = chabi(person);
ghuriye (dhoro i = 0; i < dorghyo(personKeys); i = i + 1) {
    dhoro key = personKeys[i];
    dekho(key, ":", person[key]);
}
```

## Built-in Functions

### Output
- `dekho(...)` - Print values to console

```banglacode
dekho("Hello, World!");
dekho("Value:", 42);
dekho("Multiple", "values", "at", "once");
```

### Type Functions
- `dhoron(x)` - ধরন - Get type of value
- `lipi(x)` - লিপি - Convert to string
- `sonkha(x)` - সংখ্যা - Convert to number
- `dorghyo(x)` - দৈর্ঘ্য - Get length of string/array

```banglacode
dekho(dhoron(42));           // Output: NUMBER
dekho(dhoron("hello"));      // Output: STRING
dekho(lipi(123));            // Output: "123"
dekho(sonkha("456"));        // Output: 456
dekho(dorghyo("West Bengal")); // Output: 10
```

### Array Functions
- `dhokao(array, value)` - ঢোকাও - Add element to array
- `berKoro(array)` - বের করো - Remove and return last element
- `dorghyo(array)` - দৈর্ঘ্য - Get array length

```banglacode
dhoro arr = [1, 2, 3];
dhokao(arr, 4);
dekho(arr);           // Output: [1, 2, 3, 4]
dhoro last = berKoro(arr);
dekho(last);          // Output: 4
dekho(dorghyo(arr));  // Output: 3
```

### Map Functions
- `chabi(map)` - চাবি - Get array of keys

```banglacode
dhoro obj = {"a": 1, "b": 2};
dhoro k = chabi(obj);
dekho(k);  // Output: ["a", "b"]
```

### Math Functions
- `borgomul(x)` - বর্গমূল - Square root
- `ghat(base, exp)` - ঘাত - Power
- `niche(x)` - নিচে - Round down (floor)
- `upore(x)` - উপরে - Round up (ceil)
- `kache(x)` - কাছে - Round to nearest
- `niratek(x)` - নিরপেক্ষ - Absolute value
- `choto(a, b, ...)` - ছোট - Minimum value
- `boro(a, b, ...)` - বড় - Maximum value
- `lotto()` - লটো - Random number between 0 and 1

```banglacode
dekho(borgomul(16));     // Output: 4
dekho(ghat(2, 3));       // Output: 8
dekho(niche(4.7));       // Output: 4
dekho(upore(4.2));       // Output: 5
dekho(kache(4.5));       // Output: 5
dekho(niratek(-10));     // Output: 10
dekho(choto(5, 2, 8));   // Output: 2
dekho(boro(5, 2, 8));    // Output: 8
dekho(lotto());          // Output: 0.xxx (random)
```

### String Functions
- `boroHater(str)` - বড় হাতের - Convert to uppercase
- `chotoHater(str)` - ছোট হাতের - Convert to lowercase
- `chhanto(str)` - ছাঁটো - Remove leading/trailing whitespace
- `bhag(str, separator)` - ভাগ - Split string into array
- `joro(array, separator)` - জোড়ো - Join array into string
- `khojo(str, substring)` - খোঁজো - Find index of substring (-1 if not found)
- `angsho(str, start, end)` - অংশ - Extract substring
- `bodlo(str, old, new)` - বদলো - Replace all occurrences

```banglacode
dekho(boroHater("hello"));        // Output: HELLO
dekho(chotoHater("WORLD"));       // Output: world
dekho(chhanto("  hello  "));      // Output: hello
dekho(bhag("a,b,c", ","));        // Output: [a, b, c]
dekho(joro(["a", "b"], "-"));     // Output: a-b
dekho(khojo("hello", "ll"));      // Output: 2
dekho(angsho("hello", 1, 4));     // Output: ell
dekho(bodlo("hello", "l", "x"));  // Output: hexxo
```

### Additional Array Functions
- `kato(array, start, end)` - কাটো - Extract subarray
- `ulto(array)` - উল্টো - Reverse array (returns new array)
- `saja(array)` - সাজা - Sort array (returns new array)
- `ache(array, value)` - আছে - Check if value exists in array

```banglacode
dhoro arr = [3, 1, 4, 1, 5];
dekho(kato(arr, 1, 3));     // Output: [1, 4]
dekho(ulto(arr));           // Output: [5, 1, 4, 1, 3]
dekho(saja(arr));           // Output: [1, 1, 3, 4, 5]
dekho(ache(arr, 4));        // Output: sotti
dekho(ache(arr, 9));        // Output: mittha
```

### Utility Functions
- `somoy()` - সময় - Current timestamp in milliseconds
- `ghum(ms)` - ঘুম - Pause execution for milliseconds
- `nao(prompt)` - নাও - Read user input from console
- `bondho(code)` - বন্ধ - Exit program with code

```banglacode
dekho(somoy());             // Output: 1234567890123
ghum(1000);                 // Pauses for 1 second
dhoro naam = nao("Tomar naam ki: ");
dekho("Hello", naam);
```

### File Functions
- `poro(path)` - পড়ো - Read file contents as string
- `lekho(path, content)` - লেখো - Write string to file

```banglacode
// Write to file
lekho("output.txt", "Hello BanglaCode!");

// Read from file
dhoro content = poro("output.txt");
dekho(content);  // Output: Hello BanglaCode!
```

### HTTP Functions
- `server_chalu(port, handler)` - সার্ভার চালু - Start HTTP server
- `anun(url)` - আনুন - Make HTTP GET request

```banglacode
// HTTP GET request
dhoro response = anun("https://api.example.com/data");
dekho(response["status"]);  // HTTP status code
dekho(response["body"]);    // Response body
```

## Comments

Use `//` for single-line comments:

```banglacode
// This is a comment
dekho("Hello");  // This is also a comment
```

## Examples

### Example 1: Hello World
```banglacode
dekho("Hello, West Bengal!");
dekho("Namaskar!");
```

### Example 2: Calculator
```banglacode
kaj calculator(a, b, operation) {
    jodi (operation == "+") {
        ferao a + b;
    } nahole jodi (operation == "-") {
        ferao a - b;
    } nahole jodi (operation == "*") {
        ferao a * b;
    } nahole jodi (operation == "/") {
        jodi (b == 0) {
            dekho("Cannot divide by zero!");
            ferao khali;
        }
        ferao a / b;
    }
    ferao khali;
}

dekho("10 + 5 =", calculator(10, 5, "+"));
dekho("10 - 5 =", calculator(10, 5, "-"));
dekho("10 * 5 =", calculator(10, 5, "*"));
dekho("10 / 5 =", calculator(10, 5, "/"));
```

### Example 3: Fibonacci Sequence
```banglacode
kaj fibonacci(n) {
    jodi (n <= 1) {
        ferao n;
    }
    ferao fibonacci(n - 1) + fibonacci(n - 2);
}

dekho("Fibonacci sequence:");
ghuriye (dhoro i = 0; i < 10; i = i + 1) {
    dekho("F(", i, ") =", fibonacci(i));
}
```

### Example 4: Student Grade System
```banglacode
class Student {
    kaj init(naam, roll) {
        ei.naam = naam;
        ei.roll = roll;
        ei.marks = [];
    }

    kaj addMark(subject, mark) {
        push(ei.marks, {"subject": subject, "mark": mark});
    }

    kaj calculateAverage() {
        jodi (length(ei.marks) == 0) {
            ferao 0;
        }
        dhoro total = 0;
        ghuriye (dhoro i = 0; i < length(ei.marks); i = i + 1) {
            total = total + ei.marks[i]["mark"];
        }
        ferao total / length(ei.marks);
    }

    kaj displayReport() {
        dekho("Student Name:", ei.naam);
        dekho("Roll:", ei.roll);
        dekho("Marks:");
        ghuriye (dhoro i = 0; i < length(ei.marks); i = i + 1) {
            dhoro m = ei.marks[i];
            dekho("  ", m["subject"], ":", m["mark"]);
        }
        dekho("Average:", ei.calculateAverage());
    }
}

dhoro student = notun Student("Ankan", 101);
student.addMark("Bangla", 85);
student.addMark("English", 90);
student.addMark("Math", 95);
student.displayReport();
```

### Example 5: Prime Number Checker
```banglacode
kaj isPrime(n) {
    jodi (n <= 1) {
        ferao mittha;
    }
    jodi (n <= 3) {
        ferao sotti;
    }
    jodi (n % 2 == 0 ba n % 3 == 0) {
        ferao mittha;
    }

    dhoro i = 5;
    jotokkhon (i * i <= n) {
        jodi (n % i == 0 ba n % (i + 2) == 0) {
            ferao mittha;
        }
        i = i + 6;
    }
    ferao sotti;
}

dekho("Prime numbers from 1 to 50:");
ghuriye (dhoro num = 1; num <= 50; num = num + 1) {
    jodi (isPrime(num)) {
        dekho(num);
    }
}
```

## Tips and Best Practices

1. **Use meaningful variable names**: Use Bengali/Banglish words that make sense
   ```banglacode
   dhoro totalPrice = 100;  // Good
   dhoro x = 100;           // Less clear
   ```

2. **Add comments to explain complex logic**:
   ```banglacode
   // Calculate compound interest
   dhoro interest = principal * pow(1 + rate, time);
   ```

3. **Keep functions small and focused**: Each function should do one thing well

4. **Use consistent naming**: Stick to either Bengali or English words consistently

5. **Handle edge cases**: Check for null values, empty arrays, division by zero, etc.

## Error Handling

BanglaCode will display error messages for:
- Syntax errors (parse errors)
- Runtime errors (type mismatches, undefined variables)
- Division by zero
- Array index out of bounds

Example:
```banglacode
dhoro x = 10 / 0;  // Error: division by zero
```

## File Extension

Use `.bang` extension for BanglaCode source files:
- `hello.bang`
- `calculator.bang`
- `game.bang`

## Getting Help

- In REPL, type `help` to see available keywords and functions
- Type `exit` or press Ctrl+C to quit REPL
- Check the `examples/` directory for more code samples

## Contributing

BanglaCode is open for contributions! Feel free to:
- Report bugs
- Suggest new features
- Add more built-in functions
- Improve documentation
- Create tutorials in Bengali

---

**আপনার প্রোগ্রামিং যাত্রা শুভ হোক!** (May your programming journey be successful!)

**Made with ❤️ for West Bengal**
