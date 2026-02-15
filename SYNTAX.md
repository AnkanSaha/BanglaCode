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
- [Database Connectivity](#database-functions)
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
| `sthir` | fixed/constant | const |
| `bishwo` | world/global | global |
| `jodi` | if | if |
| `nahole` | else | else |
| `jotokkhon` | as long as | while |
| `ghuriye` | rotate/turn | for |
| `kaj` | work/function | function |
| `ferao` | return | return |
| `sreni` | class/category | class |
| `shuru` | start/begin | init (constructor) |
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
| `hisabe` | as | as (import alias) |
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

BanglaCode has three variable declaration keywords:

### Regular Variables (`dhoro`)
Mutable variables using `dhoro`:

```banglacode
dhoro naam = "Ankan";
dhoro boyosh = 25;
dhoro isStudent = sotti;
```

### Constants (`sthir`)
Immutable constants using `sthir` (স্থির = fixed):

```banglacode
sthir PI = 3.14159;
sthir MAX_SIZE = 100;
sthir APP_NAME = "BanglaCode";

// Trying to reassign will cause an error:
// PI = 3.14;  // Error: 'PI' ekti sthir (constant), eitake bodlano jabe na
```

### Global Variables (`bishwo`)
Global variables using `bishwo` (বিশ্ব = world) that are accessible from any scope:

```banglacode
bishwo counter = 0;

kaj increment() {
    counter = counter + 1;  // Modifies the global variable
}

increment();
increment();
dekho(counter);  // Output: 2
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

### Rest Parameters (Variadic Functions)
Use `...` to collect any number of arguments into an array:

```banglacode
kaj sum(...numbers) {
    dhoro total = 0;
    ghuriye (dhoro i = 0; i < dorghyo(numbers); i = i + 1) {
        total = total + numbers[i];
    }
    ferao total;
}

dekho(sum(1, 2, 3));        // Output: 6
dekho(sum(1, 2, 3, 4, 5));  // Output: 15
dekho(sum());               // Output: 0
```

### Mixed Parameters with Rest
```banglacode
kaj greetAll(greeting, ...names) {
    ghuriye (dhoro i = 0; i < dorghyo(names); i = i + 1) {
        dekho(greeting, names[i]);
    }
}

greetAll("Hello", "Alice", "Bob", "Charlie");
// Output:
// Hello Alice
// Hello Bob
// Hello Charlie
```

### Spread Operator
Use `...` to expand an array into individual elements:

```banglacode
// Spread in function calls
dhoro nums = [1, 2, 3, 4, 5];
dekho(sum(...nums));  // Output: 15

// Combine with regular arguments
dekho(sum(10, ...nums));  // Output: 25

// Spread in array literals
dhoro arr1 = [1, 2];
dhoro arr2 = [3, 4];
dhoro combined = [...arr1, ...arr2];  // [1, 2, 3, 4]

// Spread with dekho
dhoro items = ["apple", "banana"];
dekho(...items);  // Output: apple banana
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
sreni Manush {
    shuru(naam, boyosh) {
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
sreni Rectangle {
    shuru(width, height) {
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

pathao sreni Calculator {
    shuru() {
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
ano "math_utils.bang" hisabe math;

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
    uttor(res, "Hello from BanglaCode!");
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

### Response Helpers

BanglaCode provides simple helper functions for sending responses:

#### uttor() - Simple Response
```banglacode
uttor(res, body);                    // 200 OK with text body
uttor(res, body, 201);               // Custom status code
uttor(res, body, 200, "text/html");  // Custom content-type
```

#### json_uttor() - JSON Response
```banglacode
json_uttor(res, data);       // 200 OK with JSON body (auto Content-Type)
json_uttor(res, data, 201);  // Custom status code
```

### Full Server Example

```banglacode
kaj handleRequest(req, res) {
    dekho("Request:", req["method"], req["path"]);
    
    jodi (req["path"] == "/") {
        uttor(res, "Welcome to BanglaCode Server!");
    } nahole jodi (req["path"] == "/api/hello") {
        json_uttor(res, {"message": "Namaskar!"});
    } nahole {
        json_uttor(res, {"error": "Not Found"}, 404);
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

// Parse JSON response
dhoro data = json_poro(response["body"]);
dekho("Parsed data:", data);
```

## JSON Functions

BanglaCode provides built-in functions for working with JSON data.

### json_poro() - Parse JSON
Convert a JSON string to a BanglaCode object:

```banglacode
dhoro jsonStr = "{\"naam\": \"Ankan\", \"boyosh\": 25, \"skills\": [\"Go\", \"JavaScript\"]}";
dhoro data = json_poro(jsonStr);

dekho(data["naam"]);      // Output: Ankan
dekho(data["boyosh"]);    // Output: 25
dekho(data["skills"][0]); // Output: Go
```

### json_banao() - Create JSON String
Convert a BanglaCode object to a JSON string:

```banglacode
dhoro person = {
    "naam": "Ankan",
    "city": "Kolkata",
    "active": sotti
};
dhoro jsonStr = json_banao(person);
dekho(jsonStr);  // Output: {"active":true,"city":"Kolkata","naam":"Ankan"}

// Works with arrays too
dhoro arr = [1, 2, 3, "hello"];
dekho(json_banao(arr));  // Output: [1,2,3,"hello"]
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

### Creating Maps (JS-like Syntax)
```banglacode
// Keys can be identifiers (JS-like) or strings
dhoro person = {
    naam: "Ankan",
    boyosh: 25,
    city: "Kolkata",
    isActive: sotti
};

// Nested objects
dhoro config = {
    app: "BanglaCode",
    version: "1.0.0",
    author: {
        name: "Ankan",
        location: "West Bengal"
    },
    features: ["easy syntax", "Bangla keywords", "modularity"]
};
```

### Accessing Values
```banglacode
// Dot notation (recommended)
dekho(person.naam);     // Output: Ankan
dekho(config.author.name);  // Output: Ankan

// Bracket notation
dekho(person["city"]);  // Output: Kolkata
```

### Importing JSON Files
```banglacode
// Import JSON files directly
ano "./config.json" hisabe config;

dekho(config.name);         // Access properties
dekho(config.author.name);  // Access nested properties
dekho(config.features);     // Access arrays
```

### Modifying Values
```banglacode
person.boyosh = 26;
person.country = "India";
```

### Getting Keys
```banglacode
dhoro sobChabi = chabi(person);
dekho(sobChabi);  // Output: ["naam", "boyosh", "city", "isActive", "country"]
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

### Database Functions

BanglaCode provides production-grade database connectors for **PostgreSQL, MySQL, MongoDB, and Redis** with connection pooling support.

#### Universal Functions (Work with all databases)

- `db_jukto(type, config)` - Connect to database (sync)
- `db_jukto_async(type, config)` - Connect to database (async)
- `db_bandho(conn)` - Close connection (sync)
- `db_bandho_async(conn)` - Close connection (async)
- `db_query(conn, sql)` - Execute SELECT query (sync)
- `db_query_async(conn, sql)` - Execute SELECT query (async)
- `db_exec(conn, sql)` - Execute INSERT/UPDATE/DELETE (sync)
- `db_exec_async(conn, sql)` - Execute INSERT/UPDATE/DELETE (async)
- `db_proshno(conn, sql, params)` - Prepared query (SQL injection safe)
- `db_proshno_async(conn, sql, params)` - Prepared query async

#### Connection Pool Functions

- `db_pool_banao(type, config, maxConns)` - Create connection pool
- `db_pool_nao(pool)` - Get connection from pool
- `db_pool_ferot(pool, conn)` - Return connection to pool
- `db_pool_bondho(pool)` - Close connection pool
- `db_pool_tothyo(pool)` - Get pool statistics

```banglacode
// PostgreSQL with connection pool (50-100x faster!)
dhoro pool = db_pool_banao("postgres", {
    "host": "localhost",
    "port": 5432,
    "database": "myapp",
    "user": "admin",
    "password": "secret"
}, 10); // Max 10 connections

// Get connection from pool
dhoro conn = db_pool_nao(pool);

// Execute query
dhoro users = db_query(conn, "SELECT * FROM users WHERE age > 25");
ghuriye (dhoro i = 0; i < dorghyo(users["rows"]); i = i + 1) {
    dekho("User:", users["rows"][i]["name"]);
}

// Prepared query (SQL injection safe)
dhoro result = db_proshno(conn, "INSERT INTO users (name, email) VALUES ($1, $2)",
    ["Rahim", "rahim@example.com"]);

// Return connection to pool (important for reuse!)
db_pool_ferot(pool, conn);

// Close pool when done
db_pool_bondho(pool);
```

#### PostgreSQL Specific Functions

- `db_jukto_postgres(config)` - PostgreSQL connection
- `db_query_postgres(conn, sql)` - Execute query
- `db_exec_postgres(conn, sql)` - Execute statement
- `db_proshno_postgres(conn, sql, params)` - Prepared statement
- `db_transaction_shuru_postgres(conn)` - Begin transaction
- `db_commit_postgres(tx)` - Commit transaction
- `db_rollback_postgres(tx)` - Rollback transaction
- `db_bulk_insert_postgres(conn, table, columns, rows)` - Efficient bulk insert

```banglacode
// Transaction example
dhoro tx = db_transaction_shuru_postgres(conn);

chesta {
    db_exec_postgres(conn, "INSERT INTO accounts (name, balance) VALUES ('User1', 1000)");
    db_exec_postgres(conn, "UPDATE accounts SET balance = balance - 100 WHERE name = 'User1'");

    // Commit if successful
    db_commit_postgres(tx);
    dekho("Transaction completed!");
} dhoro_bhul (error) {
    // Rollback on error
    db_rollback_postgres(tx);
    dekho("Transaction failed:", error);
}

// Bulk insert example (10-100x faster than individual inserts!)
dhoro users = [
    ["Alice", 25, "alice@example.com"],
    ["Bob", 30, "bob@example.com"],
    ["Charlie", 35, "charlie@example.com"]
];

dhoro result = db_bulk_insert_postgres(conn, "users", ["name", "age", "email"], users);
dekho("Inserted", result["rows_affected"], "users in bulk!");
```

#### MySQL Specific Functions

- `db_jukto_mysql(config)` - MySQL connection
- `db_query_mysql(conn, sql)` - Execute query
- `db_exec_mysql(conn, sql)` - Execute statement
- `db_proshno_mysql(conn, sql, params)` - Prepared statement
- `db_transaction_shuru_mysql(conn)` - Begin transaction
- `db_commit_mysql(tx)` - Commit transaction
- `db_rollback_mysql(tx)` - Rollback transaction
- `db_bulk_insert_mysql(conn, table, columns, rows)` - Efficient bulk insert

#### MongoDB Specific Functions

**Basic Operations:**
- `db_jukto_mongodb(config)` - MongoDB connection
- `db_khojo_mongodb(conn, collection, filter)` - Find documents
- `db_khojo_async_mongodb(conn, collection, filter)` - Find documents async
- `db_dhokao_mongodb(conn, collection, doc)` - Insert document
- `db_dhokao_async_mongodb(conn, collection, doc)` - Insert document async
- `db_update_mongodb(conn, collection, filter, update)` - Update documents
- `db_update_async_mongodb(conn, collection, filter, update)` - Update async
- `db_mujhe_mongodb(conn, collection, filter)` - Delete documents
- `db_mujhe_async_mongodb(conn, collection, filter)` - Delete async

**Advanced Operations:**
- `db_aggregate_mongodb(conn, collection, pipeline)` - Execute aggregation pipeline
- `db_findone_mongodb(conn, collection, filter)` - Find single document
- `db_count_mongodb(conn, collection, filter)` - Count matching documents
- `db_distinct_mongodb(conn, collection, field, filter)` - Get distinct values
- `db_khojo_options_mongodb(conn, collection, filter, options)` - Find with sort/limit/skip
- `db_create_index_mongodb(conn, collection, keys)` - Create index for performance
- `db_insertmany_mongodb(conn, collection, docs)` - Bulk insert documents

```banglacode
// MongoDB example
dhoro mongoConn = db_jukto("mongodb", {
    "host": "localhost",
    "port": 27017,
    "database": "mydb"
});

// Find documents
dhoro users = db_khojo_mongodb(mongoConn, "users", {
    "age": {"$gt": 25},
    "city": "Dhaka"
});

dekho("Found users:", dorghyo(users["rows"]));

// Insert document
db_dhokao_mongodb(mongoConn, "users", {
    "name": "Karim Ahmed",
    "age": 30,
    "city": "Dhaka",
    "profession": "Engineer"
});

// Update documents
db_update_mongodb(mongoConn, "users",
    {"city": "Dhaka"},
    {"$set": {"country": "Bangladesh"}}
);

// Advanced: Aggregation pipeline
dhoro pipeline = [
    {"$match": {"city": "Dhaka"}},
    {"$group": {"_id": "$profession", "count": {"$sum": 1}}},
    {"$sort": {"count": -1}}
];
dhoro stats = db_aggregate_mongodb(mongoConn, "users", pipeline);
dekho("User statistics:", stats);

// Advanced: Find with options (sort, limit, skip)
dhoro topUsers = db_khojo_options_mongodb(mongoConn, "users",
    {"age": {"$gte": 18}},
    {"sort": {"age": -1}, "limit": 10, "skip": 0}
);

// Create index for faster queries
db_create_index_mongodb(mongoConn, "users", {"email": 1});

db_bandho(mongoConn);
```

#### Redis Specific Functions

**String Operations:**
- `db_jukto_redis(config)` - Redis connection
- `db_set_redis(conn, key, value, ttl)` - Set key-value
- `db_set_async_redis(conn, key, value, ttl)` - Set key-value async
- `db_get_redis(conn, key)` - Get value
- `db_get_async_redis(conn, key)` - Get value async
- `db_del_redis(conn, key)` - Delete key
- `db_expire_redis(conn, key, seconds)` - Set expiration

**List Operations:**
- `db_lpush_redis(conn, key, value)` - List push left
- `db_rpush_redis(conn, key, value)` - List push right
- `db_lpop_redis(conn, key)` - List pop left
- `db_rpop_redis(conn, key)` - List pop right

**Hash Operations:**
- `db_hset_redis(conn, key, field, value)` - Hash set field
- `db_hget_redis(conn, key, field)` - Hash get field
- `db_hgetall_redis(conn, key)` - Hash get all fields

**Sorted Sets (Leaderboards, Rankings):**
- `db_zadd_redis(conn, key, members)` - Add members with scores
- `db_zrange_redis(conn, key, start, stop)` - Get members by rank range
- `db_zrank_redis(conn, key, member)` - Get member's rank
- `db_zscore_redis(conn, key, member)` - Get member's score
- `db_zrem_redis(conn, key, members)` - Remove members

**Sets (Unique Collections):**
- `db_sadd_redis(conn, key, members)` - Add members to set
- `db_smembers_redis(conn, key)` - Get all set members
- `db_sismember_redis(conn, key, member)` - Check if member exists
- `db_srem_redis(conn, key, members)` - Remove members from set
- `db_sinter_redis(conn, keys)` - Intersection of sets
- `db_sunion_redis(conn, keys)` - Union of sets
- `db_sdiff_redis(conn, keys)` - Difference of sets

**Counters (Atomic Operations):**
- `db_incr_redis(conn, key)` - Increment by 1
- `db_decr_redis(conn, key)` - Decrement by 1
- `db_incrby_redis(conn, key, amount)` - Increment by amount
- `db_decrby_redis(conn, key, amount)` - Decrement by amount
- `db_incrbyfloat_redis(conn, key, amount)` - Increment by float

**Pub/Sub (Message Queues):**
- `db_publish_redis(conn, channel, message)` - Publish message to channel

**Utilities:**
- `db_ttl_redis(conn, key)` - Get time-to-live in seconds
- `db_persist_redis(conn, key)` - Remove expiration
- `db_exists_redis(conn, keys)` - Check if keys exist
- `db_keys_redis(conn, pattern)` - Get keys matching pattern

```banglacode
// Redis caching example
dhoro redisConn = db_jukto("redis", {
    "host": "localhost",
    "port": 6379
});

// Set key-value with 1 hour TTL
db_set_redis(redisConn, "user:1", "Rahim Ahmed");
db_expire_redis(redisConn, "user:1", 3600);

// Get cached value
dhoro cachedUser = db_get_redis(redisConn, "user:1");
dekho("Cached user:", cachedUser);

// List operations
db_rpush_redis(redisConn, "queue:tasks", "Task 1");
db_rpush_redis(redisConn, "queue:tasks", "Task 2");
dhoro task = db_lpop_redis(redisConn, "queue:tasks");
dekho("Processing task:", task);

// Hash operations
db_hset_redis(redisConn, "user:1:profile", "name", "Rahim");
db_hset_redis(redisConn, "user:1:profile", "age", "30");
dhoro profile = db_hgetall_redis(redisConn, "user:1:profile");
dekho("Profile:", profile);

// Advanced: Sorted Sets (Leaderboard)
db_zadd_redis(redisConn, "leaderboard", {
    "player1": 100,
    "player2": 95,
    "player3": 120
});
dhoro topPlayers = db_zrange_redis(redisConn, "leaderboard", 0, 2);  // Top 3
dekho("Top players:", topPlayers);

// Advanced: Sets (Unique tags)
db_sadd_redis(redisConn, "tags:article1", ["bangla", "programming", "tutorial"]);
db_sadd_redis(redisConn, "tags:article2", ["bangla", "coding", "tutorial"]);
dhoro commonTags = db_sinter_redis(redisConn, ["tags:article1", "tags:article2"]);
dekho("Common tags:", commonTags);

// Advanced: Atomic counters
db_incr_redis(redisConn, "page:views");
db_incrby_redis(redisConn, "user:points", 50);
dhoro views = db_get_redis(redisConn, "page:views");
dekho("Page views:", views);

db_bandho(redisConn);
```

#### Async Database Queries

```banglacode
// Async database operations with proyash/opekha
proyash kaj fetchUsers() {
    dhoro conn = opekha db_jukto_async("postgres", {
        "host": "localhost",
        "database": "myapp"
    });

    dhoro users = opekha db_query_async(conn, "SELECT * FROM users");
    opekha db_bandho_async(conn);

    ferao users;
}

// Call async function
dhoro result = opekha fetchUsers();
dekho("Fetched", dorghyo(result["rows"]), "users");
```

## Environment Variables (.env Files)

BanglaCode provides built-in support for loading and managing environment variables from `.env` files. This is perfect for managing configuration across different environments (development, UAT, production).

### Environment Variable Functions

- `env_load(filename)` - Load specific .env file
- `env_load_auto(environment)` - Auto-load .env.{environment} or fallback to .env
- `env_get(key)` - Get environment variable value
- `env_get_default(key, default)` - Get with default fallback
- `env_set(key, value)` - Set environment variable at runtime
- `env_all()` - Get all environment variables as a map
- `env_clear()` - Clear all loaded environment variables

### Simple .env File Loading

```banglacode
// Load default .env file
env_load(".env");

// Get environment variables
dhoro api_key = env_get("API_KEY");
dhoro db_host = env_get("DB_HOST");
dhoro db_port = env_get("DB_PORT");

dekho("Connecting to:", db_host, ":", db_port);
```

### Multi-Environment Support

BanglaCode makes it easy to manage different environments (UAT, staging, production):

```banglacode
// Auto-load based on environment
// Tries .env.uat first, then falls back to .env
env_load_auto("uat");

// Or load production environment
env_load_auto("prod");  // Tries .env.prod, then .env

// Get values with defaults
dhoro api_url = env_get_default("API_URL", "http://localhost:3000");
dhoro debug_mode = env_get_default("DEBUG", "false");

dekho("API URL:", api_url);
dekho("Debug Mode:", debug_mode);
```

### .env File Format

Create a `.env` file in your project directory:

```
# Database Configuration
DB_HOST=localhost
DB_PORT=5432
DB_NAME=myapp
DB_USER=admin
DB_PASSWORD=secret123

# API Keys
API_KEY=your-api-key-here
SECRET_KEY=your-secret-key

# Environment
NODE_ENV=development
DEBUG=true
```

### Different Environment Files

**`.env`** (Default/Development):
```
API_URL=http://localhost:3000
DB_HOST=localhost
DEBUG=true
```

**`.env.uat`** (UAT/Testing):
```
API_URL=https://uat-api.example.com
DB_HOST=uat-db.example.com
DEBUG=true
```

**`.env.prod`** (Production):
```
API_URL=https://api.example.com
DB_HOST=prod-db.example.com
DEBUG=false
```

### Complete Example with Database Connection

```banglacode
// Load environment-specific configuration
env_load_auto("prod");  // Tries .env.prod, then .env

// Get database credentials from .env
dhoro db_host = env_get("DB_HOST");
dhoro db_port = env_get_default("DB_PORT", "5432");
dhoro db_name = env_get("DB_NAME");
dhoro db_user = env_get("DB_USER");
dhoro db_pass = env_get("DB_PASSWORD");

// Connect to database using env variables
dhoro conn = db_jukto("postgres", {
    "host": db_host,
    "port": jongate(db_port),  // Convert string to number
    "database": db_name,
    "user": db_user,
    "password": db_pass
});

// Use connection
dhoro users = db_query(conn, "SELECT * FROM users");
dekho("Found", dorghyo(users["rows"]), "users");

db_bandho(conn);
```

### Runtime Environment Variables

```banglacode
// Set environment variable at runtime
env_set("TEMP_TOKEN", "abc123");

// Get all environment variables
dhoro all_vars = env_all();
dekho("All environment variables:", all_vars);

// Clear all environment variables
env_clear();
```

### Best Practices

1. **Never commit .env files** - Add `.env*` to `.gitignore`
2. **Use .env.example** - Commit a template without secrets
3. **Different files per environment** - `.env.dev`, `.env.uat`, `.env.prod`
4. **Use defaults** - Always provide sensible defaults with `env_get_default()`
5. **Validate early** - Check required env vars at startup

### Example .env.example (Safe to commit)

```
# Database Configuration
DB_HOST=localhost
DB_PORT=5432
DB_NAME=your_database_name
DB_USER=your_username
DB_PASSWORD=your_password

# API Keys (replace with your actual keys)
API_KEY=your_api_key_here
SECRET_KEY=your_secret_key_here

# Environment
NODE_ENV=development
DEBUG=true
```

## Comments

Use `//` for single-line comments:

```banglacode
// This is a comment
dekho("Hello");  // This is also a comment
```

### Documentation Comments with @comment

Use `// @comment` to add documentation that appears in VSCode hover:

```banglacode
// @comment Adds two numbers and returns the result
// @comment Example: add(5, 3) returns 8
kaj add(a, b) {
    ferao a + b;
}

// @comment Represents a person with name and age
// @comment Use: dhoro p = notun Person("name", age);
sreni Person {
    shuru(naam) {
        ei.naam = naam;
    }
    
    // @comment Returns a greeting message
    kaj greet() {
        ferao "Hello " + ei.naam;
    }
}
```

When you hover over `add` or `Person` anywhere in your code (or imported modules), VSCode will show the documentation.

**Features:**
- Multiple `@comment` lines are combined
- Works with functions, classes, and methods
- Shown when hovering on imported functions too
- Ignored by the interpreter (treated as regular comments)

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
sreni Student {
    shuru(naam, roll) {
        ei.naam = naam;
        ei.roll = roll;
        ei.marks = [];
    }

    kaj addMark(subject, mark) {
        dhokao(ei.marks, {"subject": subject, "mark": mark});
    }

    kaj calculateAverage() {
        jodi (dorghyo(ei.marks) == 0) {
            ferao 0;
        }
        dhoro total = 0;
        ghuriye (dhoro i = 0; i < dorghyo(ei.marks); i = i + 1) {
            total = total + ei.marks[i]["mark"];
        }
        ferao total / dorghyo(ei.marks);
    }

    kaj displayReport() {
        dekho("Student Name:", ei.naam);
        dekho("Roll:", ei.roll);
        dekho("Marks:");
        ghuriye (dhoro i = 0; i < dorghyo(ei.marks); i = i + 1) {
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

## File Extensions

BanglaCode supports three file extensions - use any of them:

**Primary Extension:**
- `.bang` (recommended) - `hello.bang`, `calculator.bang`, `game.bang`

**Alternative Extensions:**
- `.bangla` (বাংলা) - `hello.bangla`, `calculator.bangla`, `game.bangla`
- `.bong` (বং) - `hello.bong`, `calculator.bong`, `game.bong`

All three extensions provide identical functionality with full IDE support.

## Getting Help

- In REPL, type `sahajjo` (or `help`) to see available keywords and functions
- Type `baire` (or `exit`) or press Ctrl+C to quit REPL
- Type `mochho` (or `clear`) to clear the screen
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
