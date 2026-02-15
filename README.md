<div align="center">

<img src="https://raw.githubusercontent.com/nexoral/BanglaCode/main/Documentation/public/banglacode.svg" alt="BanglaCode Logo" width="180"/>

# BanglaCode

### 🇮🇳 A Full-Featured Educational Programming Language in Bengali

**Write code in Bengali. Think in Bengali. Learn in Bengali.**

> *An educational language powerful enough for real projects — inspired by BhaiLang & Vedic, but with production-grade features.*

[![Version](https://img.shields.io/github/v/release/nexoral/BanglaCode?style=for-the-badge&color=blue&logo=github)](https://github.com/nexoral/BanglaCode/releases)
[![License](https://img.shields.io/github/license/nexoral/BanglaCode?style=for-the-badge&color=green)](https://github.com/nexoral/BanglaCode/blob/main/LICENSE)
[![Stars](https://img.shields.io/github/stars/nexoral/BanglaCode?style=for-the-badge&color=yellow&logo=github)](https://github.com/nexoral/BanglaCode/stargazers)
[![Go Version](https://img.shields.io/badge/Go-1.20+-00ADD8?style=for-the-badge&logo=go&logoColor=white)](https://go.dev/)
[![Platform](https://img.shields.io/badge/Platform-Linux%20%7C%20macOS%20%7C%20Windows-lightgrey?style=for-the-badge)](https://github.com/nexoral/BanglaCode)

[🚀 Quick Start](#-quick-start) • [📚 Documentation](https://banglacode.dev) • [💡 Examples](#-examples) • [🎯 Features](#-why-banglacode) • [🤝 Contributing](#-contributing)

</div>

---

## 🌟 Why BanglaCode?

**BanglaCode** is an **educational programming language** designed for **300+ million Bengali speakers worldwide**. While it's not competing with industry-standard languages like JavaScript or Python, it goes far beyond typical toy languages by offering production-grade features for real-world learning and projects.

### 🎓 Educational, But Powerful

**Inspired by:** [BhaiLang](https://github.com/DulLabs/bhai-lang) (Hindi) and [Vedic](https://github.com/vedic-lang/vedic) (Sanskrit) — regional programming languages that introduced native-language coding to India.

**The Difference:** While BhaiLang and Vedic are excellent toy languages with basic features, **BanglaCode is a full-featured educational language** that enables you to:

✅ **Build real backends** - HTTP servers, REST APIs, WebSocket servers
✅ **Connect to databases** - PostgreSQL, MySQL, MongoDB, Redis with connection pooling
✅ **Write modular code** - Import/export system, code organization, reusable modules
✅ **Handle complex logic** - OOP, async/await, error handling, promises
✅ **Access system resources** - File I/O, networking, process management

**Perfect for:** Bengali-speaking students learning programming concepts, educators teaching CS fundamentals, and hobbyists building projects in their native language.

**Not a replacement for:** Production enterprise applications (use JavaScript, Python, Go, etc. for that). BanglaCode is a learning tool that happens to be powerful enough for real projects.

### 💪 Features Beyond Typical Educational Languages

<table>
<tr>
<td width="50%">

**🚀 Fast & Efficient**
- Go-powered interpreter
- Quick startup time
- Efficient execution
- Lightweight runtime

</td>
<td width="50%">

**🎯 Modern Language**
- Object-Oriented Programming
- Async/Await (Promises)
- Module System (Import/Export)
- Error Handling (Try/Catch/Finally)

</td>
</tr>
<tr>
<td width="50%">

**🔧 135+ Built-in Functions**
- String & Array operations
- Math & Utility functions
- HTTP server & JSON support
- **Networking (TCP, UDP, WebSocket)**
- **Database (PostgreSQL, MySQL, MongoDB, Redis)**
- **Environment Variables (.env files)**
- **Complete OS-level access**

</td>
<td width="50%">

**🛠️ Developer Experience**
- VS Code extension (IntelliSense)
- Interactive REPL
- Clear error messages
- 40+ code snippets

</td>
</tr>
</table>

### 🎯 Built for Bengali Minds

> *"আমি একজন বাংলা মাধ্যমের ছাত্র। আমি logic তৈরি করতে পারি, কিন্তু সেই logic validate করতে Programming language এর syntax শিখতে হয়। BanglaCode সেই barrier শেষ করেছে — যে ভাষা তুমি জানো, সেই ভাষাতেই logic run করো।"*
>
> **— Ankan Saha**, Creator of BanglaCode

**The Problem:** Bengali students can think logically but struggle with English-based programming syntax.

**The Solution:** BanglaCode bridges this gap with Bengali keywords (`dhoro`, `jodi`, `kaj`) while maintaining C-like structure familiar to CS students.

---

## 🚀 Quick Start

### One-Line Installation

**Linux / macOS:**
```bash
curl -fsSL https://raw.githubusercontent.com/nexoral/BanglaCode/main/Scripts/install.sh | bash
```

**Windows (PowerShell):**
```powershell
irm https://raw.githubusercontent.com/nexoral/BanglaCode/main/Scripts/install.ps1 | iex
```

### From Source

```bash
# Clone the repository
git clone https://github.com/nexoral/BanglaCode.git
cd BanglaCode

# Build the interpreter (requires Go 1.20+)
go build -o banglacode main.go

# Verify installation
./banglacode --version
```

### Your First Program

Create `hello.bang` (or `.bangla` or `.bong`):

```banglacode
// Simple variables
dhoro naam = "বাংলাদেশ";
dekho("Namaskar,", naam, "!");

// Functions
kaj factorial(n) {
    jodi (n <= 1) { ferao 1; }
    ferao n * factorial(n - 1);
}

dekho("10! =", factorial(10));  // Output: 10! = 3628800
```

Run it:
```bash
./banglacode hello.bang
```

**Output:**
```
Namaskar, বাংলাদেশ !
10! = 3628800
```

---

## 🎯 Language Features

### 1️⃣ Variables & Constants

```banglacode
// Mutable variables
dhoro counter = 0;
dhoro name = "Ankan";

// Immutable constants (cannot be reassigned)
sthir PI = 3.14159;
sthir MAX_SIZE = 1000;

// Global variables (accessible from any scope)
bishwo appVersion = "1.0.0";

// Data types: Number, String, Boolean, Array, Map, Null
dhoro num = 42;
dhoro text = "Hello";
dhoro flag = sotti;           // true
dhoro empty = khali;          // null
dhoro list = [1, 2, 3];
dhoro obj = {"key": "value"};
```

### 2️⃣ Control Flow

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

// For Loop with break/continue
ghuriye (dhoro i = 0; i < 10; i = i + 1) {
    jodi (i == 5) { chharo; }  // continue
    jodi (i == 8) { thamo; }   // break
    dekho(i);
}
```

### 3️⃣ Functions & Closures

```banglacode
// Function definition
kaj greet(name) {
    ferao "Namaskar, " + name + "!";
}

// Higher-order functions & closures
kaj makeCounter() {
    dhoro count = 0;

    ferao kaj() {
        count = count + 1;
        ferao count;
    };
}

dhoro counter = makeCounter();
dekho(counter());  // 1
dekho(counter());  // 2
dekho(counter());  // 3
```

### 4️⃣ Object-Oriented Programming

```banglacode
sreni Person {
    // Constructor
    shuru(naam, boyosh) {
        ei.naam = naam;
        ei.boyosh = boyosh;
    }

    // Methods
    kaj greet() {
        dekho("Namaskar! Ami", ei.naam);
    }

    kaj age() {
        ferao ei.boyosh;
    }
}

// Inheritance example
sreni Student {
    shuru(naam, boyosh, school) {
        ei.naam = naam;
        ei.boyosh = boyosh;
        ei.school = school;
    }

    kaj study() {
        dekho(ei.naam, "is studying at", ei.school);
    }
}

dhoro student = notun Student("Rahim", 15, "Dhaka High School");
student.study();
```

### 5️⃣ Async/Await (Promises)

```banglacode
// Async function with proyash keyword
proyash kaj fetchData(url) {
    dhoro response = opekha anun_async(url);
    ferao json_poro(response);
}

// Using async functions
proyash kaj main() {
    chesta {
        dhoro data = opekha fetchData("https://api.example.com/data");
        dekho("Fetched:", data);
    } dhoro_bhul (error) {
        dekho("Error:", error);
    }
}

// Call async function
main();
```

### 6️⃣ Module System

```banglacode
// math_utils.bang
pathao kaj add(a, b) {
    ferao a + b;
}

pathao kaj multiply(a, b) {
    ferao a * b;
}

pathao sthir PI = 3.14159;

// main.bang
ano "math_utils.bang";

dekho(add(5, 3));        // 8
dekho(multiply(4, 7));   // 28
dekho(PI);               // 3.14159

// Import with alias
ano "math_utils.bang" hisabe math;
dekho(math.add(10, 20)); // 30
```

### 7️⃣ Error Handling

```banglacode
kaj riskyOperation() {
    dhoro randomNum = lotto();

    jodi (randomNum < 0.5) {
        felo "Operation failed!";
    }

    ferao "Success!";
}

chesta {
    dhoro result = riskyOperation();
    dekho(result);
} dhoro_bhul (err) {
    dekho("Caught error:", err);
} shesh {
    dekho("Cleanup always runs");
}
```

### 8️⃣ HTTP Server

```banglacode
kaj handleRequest(req, res) {
    jodi (req.path == "/") {
        uttor(res, "Welcome to BanglaCode Server!");
    } nahole jodi (req.path == "/api/users") {
        dhoro users = [
            {"id": 1, "naam": "Ankan"},
            {"id": 2, "naam": "Rahim"}
        ];
        json_uttor(res, users);
    } nahole {
        uttor(res, "404 Not Found", 404);
    }
}

dekho("Server running on http://localhost:3000");
server_chalu(3000, handleRequest);
```

### 9️⃣ System-Level Access (NEW!)

BanglaCode provides **complete OS-level access** with 50+ system functions:

```banglacode
// File operations
dhoro size = file_akar("/path/to/file.txt");
dhoro perms = file_permission("/path/to/file.txt");
file_permission_set("/path/to/file.txt", "0755");

// Directory operations
dhoro files = directory_taliika("/home/user");
dhoro allFiles = directory_ghumao("/home/user");  // Recursive

// Process management
dhoro result = chalan("ls", ["-la"]);
dekho("Output:", result["output"]);
dekho("Exit code:", result["code"]);

// Process control
process_ghum(1000);                    // Sleep 1 second
dekho("PID:", process_id());
dekho("Parent PID:", process_parent_id());

// System information
dekho("OS:", os_naam());               // linux/darwin/windows
dekho("Architecture:", bibhag());      // amd64/arm64
dekho("CPUs:", cpu_sonkha());
dekho("Hostname:", hostname());

// Memory & Disk
dhoro totalMem = memory_total();
dhoro usedMem = memory_bebohrito();
dhoro freeMem = memory_mukt();
dhoro diskSize = disk_akar("/");

// Network information
dhoro interfaces = network_interface();
dhoro ips = ip_shokal();
dhoro mac = mac_address("eth0");

// Environment variables
dhoro path = poribesh("PATH");
poribesh_set("MY_VAR", "value");
dhoro allEnv = poribesh_shokal();

// Time & Uptime
dhoro currentTime = shomoy_ekhon();
dhoro systemUptime = uptime();
dhoro bootTime = boot_shomoy();

// Temporary files
dhoro tempDir = temp_directory();
dhoro tempFile = temp_file("prefix-");
dhoro tempFolder = temp_folder("prefix-");

// Symbolic links
symlink_banao("/target/path", "/link/path");
dhoro isSymlink = symlink_ki("/path/to/check");
dhoro linkTarget = symlink_poro("/path/to/symlink");
```

### 🔟 Database Connectivity (NEW!)

BanglaCode supports production-grade database connectors with **connection pooling** for maximum performance:

```banglacode
// PostgreSQL with connection pool (50-100x faster than new connections)
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

// Iterate results
ghuriye (dhoro i = 0; i < dorghyo(users["rows"]); i = i + 1) {
    dhoro user = users["rows"][i];
    dekho("User:", user["name"], "Age:", user["age"]);
}

// Prepared queries (SQL injection safe)
dhoro result = db_proshno(conn, "INSERT INTO users (name, email) VALUES ($1, $2)",
    ["Rahim", "rahim@example.com"]);

// Return connection to pool (reused for next query)
db_pool_ferot(pool, conn);

// MongoDB document operations
dhoro mongoConn = db_jukto("mongodb", {
    "host": "localhost",
    "database": "mydb"
});

dhoro docs = db_khojo_mongodb(mongoConn, "users", {
    "age": {"$gt": 25},
    "city": "Dhaka"
});

db_dhokao_mongodb(mongoConn, "users", {
    "name": "Karim",
    "age": 30,
    "city": "Dhaka"
});

// Redis caching
dhoro redisConn = db_jukto("redis", {"host": "localhost"});

db_set_redis(redisConn, "user:1", "Rahim Ahmed");
db_expire_redis(redisConn, "user:1", 3600); // 1 hour TTL

dhoro cachedUser = db_get_redis(redisConn, "user:1");
dekho("Cached user:", cachedUser);

// Async database queries
proyash kaj fetchUsers() {
    dhoro conn = opekha db_jukto_async("postgres", {...});
    dhoro users = opekha db_query_async(conn, "SELECT * FROM users");
    opekha db_bandho_async(conn);
    ferao users;
}

dhoro result = opekha fetchUsers();
dekho("Fetched users:", dorghyo(result["rows"]));
```

---

## 📚 130+ Built-in Functions

### 🖨️ Output & Input
- `dekho(...)` - Print to console
- `nao(prompt)` - Read user input

### 🔤 String Operations
- `boroHater(str)` - Uppercase
- `chotoHater(str)` - Lowercase
- `chhanto(str)` - Trim whitespace
- `bhag(str, sep)` - Split string
- `joro(arr, sep)` - Join array to string
- `khojo(str, substr)` - Find substring
- `angsho(str, start, end)` - Substring
- `bodlo(str, old, new)` - Replace
- `kato(str, len)` - String length

### 📦 Array Operations
- `dorghyo(arr)` - Array length
- `dhokao(arr, val)` - Push element
- `berKoro(arr)` - Pop element
- `kato(arr, start, end)` - Slice array
- `ulto(arr)` - Reverse array
- `saja(arr)` - Sort array
- `ache(arr, val)` - Contains check
- `chabi(map)` - Get map keys

### 🧮 Math Functions
- `borgomul(x)` - Square root
- `ghat(base, exp)` - Power
- `niche(x)` - Floor
- `upore(x)` - Ceiling
- `kache(x)` - Round
- `niratek(x)` - Absolute value
- `choto(...)` - Minimum
- `boro(...)` - Maximum
- `lotto()` - Random (0-1)

### 📄 File I/O
- `poro(path)` - Read file
- `lekho(path, content)` - Write file
- `file_akar(path)` - File size
- `file_permission(path)` - Get permissions
- `file_permission_set(path, mode)` - Set permissions
- `file_dhoron(path)` - File type
- `file_rename(old, new)` - Rename file
- `ache_ki(path)` - Check existence
- `folder_banao(path)` - Create directory
- `muke_felo(path)` - Delete file/directory

### 📁 Directory Operations
- `directory_taliika(path)` - List directory
- `directory_ghumao(path)` - Walk directory tree
- `directory_khali_ki(path)` - Is directory empty
- `directory_akar(path)` - Directory total size
- `kaj_directory()` - Current working directory
- `kaj_directory_bodol(path)` - Change directory

### ⚙️ Process Management
- `chalan(cmd, args)` - Execute command
- `process_id()` - Current PID
- `process_parent_id()` - Parent PID
- `process_args()` - Command-line arguments
- `process_ghum(ms)` - Sleep
- `process_maro(pid)` - Kill process
- `process_signal(pid, signal)` - Send signal
- `process_ache_ki(pid)` - Check if running
- `process_opekha(pid)` - Wait for process

### 💻 System Information
- `os_naam()` - Operating system name
- `bibhag()` - Architecture (amd64, arm64)
- `hostname()` - System hostname
- `cpu_sonkha()` - Number of CPUs
- `bebosthok_naam()` - Username
- `bari_directory()` - Home directory
- `memory_total()` - Total RAM
- `memory_bebohrito()` - Used RAM
- `memory_mukt()` - Free RAM
- `cpu_bebohrito()` - CPU usage %
- `disk_akar(path)` - Disk total size
- `disk_bebohrito(path)` - Disk used
- `disk_mukt(path)` - Disk free

### 🌐 Network Functions
- `network_interface()` - Network interfaces
- `ip_address(interface)` - IP address
- `ip_shokal()` - All IP addresses
- `mac_address(interface)` - MAC address
- `network_gateway()` - Default gateway
- `dns_server()` - DNS servers

### 🌍 HTTP & JSON
- `server_chalu(port, handler)` - Start HTTP server
- `anun(url)` - HTTP GET request
- `anun_async(url)` - Async HTTP GET
- `uttor(res, body, status, type)` - Send response
- `json_uttor(res, data, status)` - Send JSON
- `json_poro(str)` - Parse JSON
- `json_banao(obj)` - Stringify JSON

### 🌐 Networking (TCP, UDP, WebSocket)
**TCP Functions:**
- `tcp_server_chalu(port, handler)` - Start TCP server
- `tcp_jukto(host, port)` - Connect to TCP server (async)
- `tcp_pathao(conn, data)` - Send data on TCP connection
- `tcp_lekho(conn, data)` - Write data (alias)
- `tcp_shuno(conn)` - Read data (async)
- `tcp_bondho(conn)` - Close TCP connection

**UDP Functions:**
- `udp_server_chalu(port, handler)` - Start UDP server
- `udp_pathao(host, port, data)` - Send UDP packet (async)
- `udp_uttor(packet, data)` - Send UDP response
- `udp_shuno(port, handler)` - Listen for packets (alias)
- `udp_bondho(conn)` - Close UDP connection

**WebSocket Functions:**
- `websocket_server_chalu(port, handler)` - Start WebSocket server
- `websocket_jukto(url)` - Connect to WebSocket (async)
- `websocket_pathao(conn, message)` - Send message
- `websocket_bondho(conn)` - Close WebSocket connection

### 🗄️ Database Functions (NEW!)

BanglaCode provides production-grade database connectors with **connection pooling** and both **sync/async APIs**:

**Universal Functions (Work with all databases):**
- `db_jukto(type, config)` - Connect to database
- `db_jukto_async(type, config)` - Connect async
- `db_bandho(conn)` - Close connection
- `db_bandho_async(conn)` - Close async
- `db_query(conn, sql)` - Execute SELECT query
- `db_query_async(conn, sql)` - Query async
- `db_exec(conn, sql)` - Execute INSERT/UPDATE/DELETE
- `db_exec_async(conn, sql)` - Execute async
- `db_proshno(conn, sql, params)` - Prepared query (SQL injection safe)
- `db_proshno_async(conn, sql, params)` - Prepared query async

**Connection Pool Functions (50-100x faster):**
- `db_pool_banao(type, config, maxConns)` - Create connection pool
- `db_pool_nao(pool)` - Get connection from pool
- `db_pool_ferot(pool, conn)` - Return connection to pool
- `db_pool_bondho(pool)` - Close pool
- `db_pool_tothyo(pool)` - Get pool statistics

**PostgreSQL Specific:**
- `db_jukto_postgres(config)` - PostgreSQL connection
- `db_query_postgres(conn, sql)` - Execute query
- `db_exec_postgres(conn, sql)` - Execute statement
- `db_proshno_postgres(conn, sql, params)` - Prepared statement
- `db_transaction_shuru_postgres(conn)` - Begin transaction
- `db_commit_postgres(tx)` - Commit transaction
- `db_rollback_postgres(tx)` - Rollback transaction

**MySQL Specific:**
- `db_jukto_mysql(config)` - MySQL connection
- `db_query_mysql(conn, sql)` - Execute query
- `db_exec_mysql(conn, sql)` - Execute statement
- `db_proshno_mysql(conn, sql, params)` - Prepared statement
- `db_transaction_shuru_mysql(conn)` - Begin transaction
- `db_commit_mysql(tx)` - Commit transaction
- `db_rollback_mysql(tx)` - Rollback transaction

**MongoDB Specific:**
- `db_jukto_mongodb(config)` - MongoDB connection
- `db_khojo_mongodb(conn, collection, filter)` - Find documents
- `db_khojo_async_mongodb(conn, collection, filter)` - Find async
- `db_dhokao_mongodb(conn, collection, doc)` - Insert document
- `db_dhokao_async_mongodb(conn, collection, doc)` - Insert async
- `db_update_mongodb(conn, collection, filter, update)` - Update documents
- `db_update_async_mongodb(conn, collection, filter, update)` - Update async
- `db_mujhe_mongodb(conn, collection, filter)` - Delete documents
- `db_mujhe_async_mongodb(conn, collection, filter)` - Delete async

**Redis Specific:**
- `db_jukto_redis(config)` - Redis connection
- `db_set_redis(conn, key, value, ttl)` - Set key-value
- `db_set_async_redis(conn, key, value, ttl)` - Set async
- `db_get_redis(conn, key)` - Get value
- `db_get_async_redis(conn, key)` - Get async
- `db_del_redis(conn, key)` - Delete key
- `db_expire_redis(conn, key, seconds)` - Set expiration
- `db_lpush_redis(conn, key, value)` - List push left
- `db_rpush_redis(conn, key, value)` - List push right
- `db_lpop_redis(conn, key)` - List pop left
- `db_rpop_redis(conn, key)` - List pop right
- `db_hset_redis(conn, key, field, value)` - Hash set field
- `db_hget_redis(conn, key, field)` - Hash get field
- `db_hgetall_redis(conn, key)` - Hash get all fields

### ⏱️ Time Functions
- `somoy()` - Current timestamp (ms)
- `shomoy_ekhon()` - Unix timestamp
- `shomoy_format(timestamp, format)` - Format time
- `shomoy_parse(str, format)` - Parse time
- `uptime()` - System uptime (seconds)
- `boot_shomoy()` - Boot timestamp
- `timezone()` - System timezone

### 🔗 Environment & Path
- `poribesh(name)` - Get environment variable
- `poribesh_set(name, value)` - Set env var
- `poribesh_shokal()` - All env vars
- `poribesh_muke(name)` - Unset env var
- `path_joro(...)` - Join path components
- `sompurno_path(path)` - Absolute path
- `path_naam(path)` - Base name
- `directory_naam(path)` - Directory name
- `file_ext(path)` - File extension
- `path_match(pattern, path)` - Glob matching

### 📦 Temporary Files
- `temp_directory()` - System temp directory
- `temp_file(prefix)` - Create temp file
- `temp_folder(prefix)` - Create temp directory
- `temp_muche_felo()` - Clean temp files

### 🔗 Symbolic Links
- `symlink_banao(target, link)` - Create symlink
- `symlink_poro(link)` - Read symlink target
- `symlink_ki(path)` - Is symlink check
- `hardlink_banao(target, link)` - Create hardlink
- `link_sonkha(path)` - Number of links

### 🛠️ Utility
- `dhoron(x)` - Get type
- `lipi(x)` - Convert to string
- `sonkha(x)` - Convert to number
- `bondho(code)` - Exit program

---

## 🎨 Keywords Reference

### Core Keywords
| Bengali | Banglish | English | Usage |
|---------|----------|---------|-------|
| ধরো | `dhoro` | let/var | `dhoro x = 5;` |
| স্থির | `sthir` | const | `sthir PI = 3.14;` |
| বিশ্ব | `bishwo` | global | `bishwo count = 0;` |
| যদি | `jodi` | if | `jodi (x > 0) { }` |
| নাহলে | `nahole` | else | `nahole { }` |
| যতক্ষণ | `jotokkhon` | while | `jotokkhon (x < 10) { }` |
| ঘুরিয়ে | `ghuriye` | for | `ghuriye (dhoro i = 0; i < 5; i++) { }` |
| কাজ | `kaj` | function | `kaj add(a, b) { }` |
| ফেরাও | `ferao` | return | `ferao result;` |
| থামো | `thamo` | break | `thamo;` |
| ছাড়ো | `chharo` | continue | `chharo;` |

### OOP Keywords
| Bengali | Banglish | English | Usage |
|---------|----------|---------|-------|
| শ্রেণী | `sreni` | class | `sreni Person { }` |
| শুরু | `shuru` | constructor | `shuru(naam) { }` |
| নতুন | `notun` | new | `notun Person()` |
| এই | `ei` | this | `ei.naam = "Ankan";` |

### Module Keywords
| Bengali | Banglish | English | Usage |
|---------|----------|---------|-------|
| আনো | `ano` | import | `ano "module.bang";` |
| পাঠাও | `pathao` | export | `pathao kaj fn() { }` |
| হিসাবে | `hisabe` | as | `ano "x.bang" hisabe y;` |

### Async Keywords
| Bengali | Banglish | English | Usage |
|---------|----------|---------|-------|
| প্রয়াস | `proyash` | async | `proyash kaj fn() { }` |
| অপেক্ষা | `opekha` | await | `opekha promise` |

### Error Handling
| Bengali | Banglish | English | Usage |
|---------|----------|---------|-------|
| চেষ্টা | `chesta` | try | `chesta { }` |
| ধরো ভুল | `dhoro_bhul` | catch | `dhoro_bhul (e) { }` |
| শেষ | `shesh` | finally | `shesh { }` |
| ফেলো | `felo` | throw | `felo "error";` |

### Literals
| Bengali | Banglish | English | Value |
|---------|----------|---------|-------|
| সত্যি | `sotti` | true | Boolean true |
| মিথ্যা | `mittha` | false | Boolean false |
| খালি | `khali` | null | Null value |
| এবং | `ebong` | and | Logical AND |
| বা | `ba` | or | Logical OR |
| না | `na` | not | Logical NOT |

---

## 🏗️ Architecture

BanglaCode follows a **classic tree-walking interpreter** architecture:

```
Source Code (.bang/.bangla/.bong)
        ↓
    [LEXER] → Tokenization
        ↓
    [PARSER] → Syntax Analysis (Pratt Parsing)
        ↓
    [AST] → Abstract Syntax Tree
        ↓
    [EVALUATOR] → Tree-Walking Execution
        ↓
    Result / Output
```

### Project Structure

```
BanglaCode/
├── src/
│   ├── lexer/          # Tokenization (29 Bengali keywords)
│   ├── parser/         # Pratt parser (precedence climbing)
│   ├── ast/            # Abstract Syntax Tree nodes
│   ├── object/         # Runtime values & environment
│   └── evaluator/      # Tree-walking interpreter
│       ├── builtins/   # 130+ built-in functions
│       │   ├── system/   # 50+ OS-level functions
│       │   ├── network/  # TCP, UDP, WebSocket
│       │   └── database/ # PostgreSQL, MySQL, MongoDB, Redis (NEW!)
│       ├── async.go    # Async/await implementation
│       ├── classes.go  # OOP support
│       ├── modules.go  # Import/export system
│       └── errors.go   # Try/catch/finally
├── Extension/          # VS Code extension
├── Documentation/      # Next.js documentation site
├── examples/           # Sample programs
└── test/               # Test suite (100+ tests)
```

---

## 🎯 VS Code Extension

Get the full development experience with our official VS Code extension:

### Features
✅ **Syntax Highlighting** for `.bang`, `.bangla`, `.bong` files
✅ **IntelliSense** with auto-completion
✅ **40+ Code Snippets** for common patterns
✅ **Hover Documentation** for built-in functions
✅ **Error Highlighting** with diagnostics
✅ **Custom File Icons** for BanglaCode files

### Installation

**From VS Code Marketplace:**
1. Open VS Code
2. Press `Ctrl+Shift+X` (Extensions)
3. Search "BanglaCode"
4. Click Install

**From Source:**
```bash
cd Extension
npm install
npx vsce package
code --install-extension banglacode-*.vsix
```

---

## 📖 Documentation

| Resource | Description |
|----------|-------------|
| [🌐 Official Website](https://banglacode.dev) | Complete documentation & tutorials |
| [📘 SYNTAX.md](SYNTAX.md) | Language syntax reference |
| [🏗️ ARCHITECTURE.md](ARCHITECTURE.md) | Technical architecture deep-dive |
| [🗺️ ROADMAP.md](ROADMAP.md) | Future development plans |
| [🤝 CONTRIBUTING.md](CONTRIBUTING.md) | Contribution guidelines |
| [📜 CODE_OF_CONDUCT.md](CODE_OF_CONDUCT.md) | Community standards |
| [🔒 SECURITY.md](SECURITY.md) | Security policy |
| [📋 CHANGELOG.md](CHANGELOG.md) | Version history |

---

## 💡 Examples

Explore real-world programs in the `examples/` directory:

| File | Features Demonstrated |
|------|----------------------|
| `hello.bang` | Variables, functions, recursion |
| `classes.bang` | OOP, inheritance, methods |
| `async.bang` | Async/await, promises |
| `http_server.bang` | Web server, routing, JSON API |
| `modules_demo.bang` | Import/export, code organization |
| `error_handling.bang` | Try/catch/finally, custom errors |
| `file_operations.bang` | File I/O, directory traversal |
| `system_info.bang` | OS-level access, system stats |
| `loops.bang` | For/while loops, break/continue |
| `data_structures.bang` | Arrays, maps, nested structures |

**Run any example:**
```bash
./banglacode examples/http_server.bang
```

---

## 🚢 Production Deployment

### Cross-Platform Compilation

```bash
# Linux
GOOS=linux GOARCH=amd64 go build -o banglacode-linux main.go

# macOS
GOOS=darwin GOARCH=arm64 go build -o banglacode-macos main.go

# Windows
GOOS=windows GOARCH=amd64 go build -o banglacode.exe main.go
```

### Docker Support

```dockerfile
FROM golang:1.20-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o banglacode main.go

FROM alpine:latest
COPY --from=builder /app/banglacode /usr/local/bin/
CMD ["banglacode"]
```

---

## 🤝 Contributing

We welcome contributions! BanglaCode is built by and for the Bengali-speaking community.

### How to Contribute

1. **Fork** the repository
2. **Create** a feature branch (`git checkout -b feature/amazing-feature`)
3. **Commit** your changes (`git commit -m 'feat: add amazing feature'`)
4. **Push** to the branch (`git push origin feature/amazing-feature`)
5. **Open** a Pull Request

### Priority Areas

🎯 **High Priority:**
- Performance optimizations
- Additional built-in functions
- Better error messages in Bengali
- Bengali tutorials and documentation

🔧 **Medium Priority:**
- Online playground/REPL
- Package manager
- Standard library expansion
- IDE integrations (IntelliJ, Sublime)

📚 **Community:**
- Example programs
- Tutorial videos
- Translation improvements
- Bug reports and fixes

See [CONTRIBUTING.md](CONTRIBUTING.md) for detailed guidelines.

---

## 🌍 Community & Support

<table>
<tr>
<td width="50%">

### 💬 Get Help
- [GitHub Discussions](https://github.com/nexoral/BanglaCode/discussions)
- [GitHub Issues](https://github.com/nexoral/BanglaCode/issues)
- [Documentation](https://banglacode.dev)

</td>
<td width="50%">

### 📊 Project Stats
- **300M+** potential Bengali-speaking users
- **130+** built-in functions (database, networking, system access)
- **29** Bengali keywords
- **Educational** language with production-grade features

</td>
</tr>
</table>

---

## 📜 License

BanglaCode is open source software licensed under the **GNU General Public License v3.0**.

This means you can:
- ✅ Use commercially
- ✅ Modify
- ✅ Distribute
- ✅ Use privately

See [LICENSE](LICENSE) for full details.

---

## 🙏 Acknowledgments

BanglaCode is inspired by great programming languages and communities:

- **C** — Syntax discipline and performance
- **JavaScript** — Modern features and async/await
- **Go** — Simplicity, performance, and tooling
- **Python** — Beginner-friendly philosophy
- **The Bengali Community** — Making programming accessible to 300M+ speakers

Special thanks to all contributors who helped make this vision a reality!

---

## 👨‍💻 Author

<div align="center">

**Ankan Saha**
Creator & Lead Developer
West Bengal, India

*"Programming should be about logic, not language barriers."*

[![GitHub](https://img.shields.io/badge/GitHub-nexoral-181717?style=for-the-badge&logo=github)](https://github.com/nexoral)

</div>

---

<div align="center">

### আপনার প্রোগ্রামিং যাত্রা শুভ হোক!
*May your programming journey be successful!*

**Made with ❤️ for Bengali developers worldwide**

[⬆ Back to Top](#banglacode)

</div>
