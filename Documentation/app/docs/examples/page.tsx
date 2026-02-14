import CodeBlock from "@/components/CodeBlock";
import DocNavigation from "@/components/DocNavigation";

export default function Examples() {
  return (
    <div>
      <div className="flex items-center gap-2 text-sm text-muted-foreground mb-4">
        <span className="px-2 py-1 bg-primary/10 text-primary rounded-full text-xs font-medium">
          Reference
        </span>
      </div>

      <h1>Code Examples</h1>

      <p className="lead text-xl text-muted-foreground mt-4">
        A collection of complete, working BanglaCode programs demonstrating
        various language features and common programming patterns.
      </p>

      <h2>Hello World</h2>

      <CodeBlock
        filename="hello.bang"
        code={`// The classic first program
dekho("Namaskar, BanglaCode!");
dekho("Hello, World!");

// With variables
dhoro naam = "Rahim";
dhoro greeting = "Shuvo Sokal";
dekho(greeting + ",", naam + "!");`}
      />

      <h2>Calculator</h2>

      <CodeBlock
        filename="calculator.bang"
        code={`sreni Calculator {
    shuru() {
        ei.history = [];
    }

    kaj add(a, b) {
        dhoro result = a + b;
        ei.saveHistory(a, "+", b, result);
        ferao result;
    }

    kaj subtract(a, b) {
        dhoro result = a - b;
        ei.saveHistory(a, "-", b, result);
        ferao result;
    }

    kaj multiply(a, b) {
        dhoro result = a * b;
        ei.saveHistory(a, "*", b, result);
        ferao result;
    }

    kaj divide(a, b) {
        jodi (b == 0) {
            felo "Cannot divide by zero!";
        }
        dhoro result = a / b;
        ei.saveHistory(a, "/", b, result);
        ferao result;
    }

    kaj saveHistory(a, op, b, result) {
        dhokao(ei.history, lipi(a) + " " + op + " " + lipi(b) + " = " + lipi(result));
    }

    kaj showHistory() {
        dekho("=== Calculation History ===");
        ghuriye (dhoro i = 0; i < dorghyo(ei.history); i = i + 1) {
            dekho(i + 1, ".", ei.history[i]);
        }
    }
}

dhoro calc = notun Calculator();
dekho(calc.add(10, 5));       // 15
dekho(calc.multiply(3, 4));   // 12
dekho(calc.divide(20, 4));    // 5
calc.showHistory();`}
      />

      <h2>Fibonacci Sequence</h2>

      <CodeBlock
        filename="fibonacci.bang"
        code={`// Recursive Fibonacci
kaj fibRecursive(n) {
    jodi (n <= 1) {
        ferao n;
    }
    ferao fibRecursive(n - 1) + fibRecursive(n - 2);
}

// Iterative Fibonacci (more efficient)
kaj fibIterative(n) {
    jodi (n <= 1) {
        ferao n;
    }

    dhoro a = 0;
    dhoro b = 1;

    ghuriye (dhoro i = 2; i <= n; i = i + 1) {
        dhoro temp = a + b;
        a = b;
        b = temp;
    }

    ferao b;
}

// Generate Fibonacci sequence
kaj fibSequence(count) {
    dhoro sequence = [];
    ghuriye (dhoro i = 0; i < count; i = i + 1) {
        dhokao(sequence, fibIterative(i));
    }
    ferao sequence;
}

dekho("First 15 Fibonacci numbers:");
dekho(fibSequence(15));`}
      />

      <h2>Prime Number Checker</h2>

      <CodeBlock
        filename="primes.bang"
        code={`kaj isPrime(n) {
    jodi (n < 2) {
        ferao mittha;
    }
    jodi (n == 2) {
        ferao sotti;
    }
    jodi (n % 2 == 0) {
        ferao mittha;
    }

    ghuriye (dhoro i = 3; i * i <= n; i = i + 2) {
        jodi (n % i == 0) {
            ferao mittha;
        }
    }

    ferao sotti;
}

kaj findPrimes(limit) {
    dhoro primes = [];
    ghuriye (dhoro n = 2; n <= limit; n = n + 1) {
        jodi (isPrime(n)) {
            dhokao(primes, n);
        }
    }
    ferao primes;
}

kaj primeFactors(n) {
    dhoro factors = [];
    dhoro d = 2;

    jotokkhon (d * d <= n) {
        jotokkhon (n % d == 0) {
            dhokao(factors, d);
            n = n / d;
        }
        d = d + 1;
    }

    jodi (n > 1) {
        dhokao(factors, n);
    }

    ferao factors;
}

dekho("Primes up to 50:", findPrimes(50));
dekho("Is 97 prime?", isPrime(97));
dekho("Prime factors of 84:", primeFactors(84));`}
      />

      <h2>Sorting Algorithms</h2>

      <CodeBlock
        filename="sorting.bang"
        code={`// Bubble Sort
kaj bubbleSort(arr) {
    dhoro n = dorghyo(arr);
    dhoro result = kato(arr, 0);  // Copy array

    ghuriye (dhoro i = 0; i < n - 1; i = i + 1) {
        ghuriye (dhoro j = 0; j < n - i - 1; j = j + 1) {
            jodi (result[j] > result[j + 1]) {
                // Swap
                dhoro temp = result[j];
                result[j] = result[j + 1];
                result[j + 1] = temp;
            }
        }
    }

    ferao result;
}

// Selection Sort
kaj selectionSort(arr) {
    dhoro n = dorghyo(arr);
    dhoro result = kato(arr, 0);

    ghuriye (dhoro i = 0; i < n - 1; i = i + 1) {
        dhoro minIdx = i;

        ghuriye (dhoro j = i + 1; j < n; j = j + 1) {
            jodi (result[j] < result[minIdx]) {
                minIdx = j;
            }
        }

        jodi (minIdx != i) {
            dhoro temp = result[i];
            result[i] = result[minIdx];
            result[minIdx] = temp;
        }
    }

    ferao result;
}

// Test
dhoro numbers = [64, 34, 25, 12, 22, 11, 90];
dekho("Original:", numbers);
dekho("Bubble Sort:", bubbleSort(numbers));
dekho("Selection Sort:", selectionSort(numbers));`}
      />

      <h2>Todo List Application</h2>

      <CodeBlock
        filename="todo.bang"
        code={`sreni TodoApp {
    shuru() {
        ei.todos = [];
        ei.nextId = 1;
    }

    kaj add(text) {
        dhoro todo = {
            id: ei.nextId,
            text: text,
            done: mittha,
            createdAt: somoy()
        };
        dhokao(ei.todos, todo);
        ei.nextId = ei.nextId + 1;
        dekho("Added:", text);
        ferao todo.id;
    }

    kaj complete(id) {
        ghuriye (dhoro i = 0; i < dorghyo(ei.todos); i = i + 1) {
            jodi (ei.todos[i].id == id) {
                ei.todos[i].done = sotti;
                dekho("Completed:", ei.todos[i].text);
                ferao sotti;
            }
        }
        dekho("Todo not found:", id);
        ferao mittha;
    }

    kaj remove(id) {
        dhoro newTodos = [];
        ghuriye (dhoro i = 0; i < dorghyo(ei.todos); i = i + 1) {
            jodi (ei.todos[i].id != id) {
                dhokao(newTodos, ei.todos[i]);
            }
        }
        ei.todos = newTodos;
    }

    kaj list() {
        dekho("\\n=== Todo List ===");
        jodi (dorghyo(ei.todos) == 0) {
            dekho("No todos!");
            ferao;
        }

        ghuriye (dhoro i = 0; i < dorghyo(ei.todos); i = i + 1) {
            dhoro t = ei.todos[i];
            dhoro status = t.done ? "[X]" : "[ ]";
            dekho(status, t.id + ".", t.text);
        }
    }

    kaj pending() {
        dhoro count = 0;
        ghuriye (dhoro i = 0; i < dorghyo(ei.todos); i = i + 1) {
            jodi (na ei.todos[i].done) {
                count = count + 1;
            }
        }
        ferao count;
    }
}

// Usage
dhoro app = notun TodoApp();
app.add("Learn BanglaCode");
app.add("Build a project");
app.add("Share with friends");
app.complete(1);
app.list();
dekho("\\nPending tasks:", app.pending());`}
      />

      <h2>Bank Account System</h2>

      <CodeBlock
        filename="bank.bang"
        code={`sreni BankAccount {
    shuru(accountNumber, owner, initialBalance) {
        ei.accountNumber = accountNumber;
        ei.owner = owner;
        ei.balance = initialBalance;
        ei.transactions = [];
    }

    kaj deposit(amount) {
        jodi (amount <= 0) {
            felo "Deposit amount must be positive";
        }
        ei.balance = ei.balance + amount;
        ei.recordTransaction("DEPOSIT", amount);
        dekho("Deposited Tk.", amount);
    }

    kaj withdraw(amount) {
        jodi (amount <= 0) {
            felo "Withdrawal amount must be positive";
        }
        jodi (amount > ei.balance) {
            felo "Insufficient funds";
        }
        ei.balance = ei.balance - amount;
        ei.recordTransaction("WITHDRAW", amount);
        dekho("Withdrew Tk.", amount);
    }

    kaj transfer(toAccount, amount) {
        ei.withdraw(amount);
        toAccount.deposit(amount);
        dekho("Transferred Tk.", amount, "to", toAccount.owner);
    }

    kaj recordTransaction(type, amount) {
        dhokao(ei.transactions, {
            type: type,
            amount: amount,
            balance: ei.balance,
            timestamp: somoy()
        });
    }

    kaj getBalance() {
        ferao ei.balance;
    }

    kaj printStatement() {
        dekho("\\n=== Account Statement ===");
        dekho("Account:", ei.accountNumber);
        dekho("Owner:", ei.owner);
        dekho("Current Balance: Tk.", ei.balance);
        dekho("\\nTransactions:");

        ghuriye (dhoro i = 0; i < dorghyo(ei.transactions); i = i + 1) {
            dhoro t = ei.transactions[i];
            dekho(" ", t.type, "Tk.", t.amount, "- Balance: Tk.", t.balance);
        }
    }
}

// Usage
dhoro account1 = notun BankAccount("001", "Rahim", 10000);
dhoro account2 = notun BankAccount("002", "Karim", 5000);

account1.deposit(5000);
account1.withdraw(2000);
account1.transfer(account2, 3000);

account1.printStatement();
account2.printStatement();`}
      />

      <h2>Simple HTTP API</h2>

      <CodeBlock
        filename="api.bang"
        code={`// Simple REST API for managing users
dhoro users = [];
dhoro nextId = 1;

kaj findUserById(id) {
    ghuriye (dhoro i = 0; i < dorghyo(users); i = i + 1) {
        jodi (users[i].id == id) {
            ferao users[i];
        }
    }
    ferao khali;
}

kaj deleteUserById(id) {
    dhoro newUsers = [];
    ghuriye (dhoro i = 0; i < dorghyo(users); i = i + 1) {
        jodi (users[i].id != id) {
            dhokao(newUsers, users[i]);
        }
    }
    users = newUsers;
}

server_chalu(8080, kaj(req, res) {
    // Enable CORS
    res.headers["Access-Control-Allow-Origin"] = "*";
    res.headers["Content-Type"] = "application/json";

    // GET /users - List all users
    jodi (req.path == "/users" ebong req.method == "GET") {
        json_uttor(res, {
            success: sotti,
            data: users,
            count: dorghyo(users)
        });
    }
    // POST /users - Create user
    nahole jodi (req.path == "/users" ebong req.method == "POST") {
        chesta {
            dhoro body = json_poro(req.body);
            dhoro user = {
                id: nextId,
                naam: body.naam,
                email: body.email,
                createdAt: somoy()
            };
            dhokao(users, user);
            nextId = nextId + 1;
            json_uttor(res, {success: sotti, data: user}, 201);
        } dhoro_bhul (e) {
            json_uttor(res, {success: mittha, error: lipi(e)}, 400);
        }
    }
    // GET /users/:id - Get single user
    nahole jodi (angsho(req.path, 0, 7) == "/users/" ebong req.method == "GET") {
        dhoro id = sonkha(angsho(req.path, 7));
        dhoro user = findUserById(id);
        jodi (user != khali) {
            json_uttor(res, {success: sotti, data: user});
        } nahole {
            json_uttor(res, {success: mittha, error: "User not found"}, 404);
        }
    }
    // DELETE /users/:id - Delete user
    nahole jodi (angsho(req.path, 0, 7) == "/users/" ebong req.method == "DELETE") {
        dhoro id = sonkha(angsho(req.path, 7));
        deleteUserById(id);
        json_uttor(res, {success: sotti, message: "User deleted"});
    }
    // 404 for other routes
    nahole {
        json_uttor(res, {success: mittha, error: "Not found"}, 404);
    }
});

dekho("API running on http://localhost:8080");`}
      />

      <h2>Async Data Fetcher</h2>

      <p>
        Demonstrates asynchronous programming with <code>proyash</code> (async) and <code>opekha</code> (await).
        Shows concurrent execution with <code>sob_proyash</code> (Promise.all).
      </p>

      <CodeBlock
        code={`// Simulate fetching data from different sources
proyash kaj fetchUser(id) {
    dekho("Fetching user", id, "...");
    opekha ghumaao(1000);  // Simulate network delay
    ferao {naam: "User " + lipi(id), id: id};
}

proyash kaj fetchPosts(userId) {
    dekho("Fetching posts for user", userId, "...");
    opekha ghumaao(1500);  // Simulate network delay
    ferao ["Post 1", "Post 2", "Post 3"];
}

proyash kaj fetchComments(userId) {
    dekho("Fetching comments for user", userId, "...");
    opekha ghumaao(800);  // Simulate network delay
    ferao ["Comment 1", "Comment 2"];
}

// Load dashboard - sequential approach (slow)
proyash kaj loadDashboardSlow() {
    dekho("=== Sequential Loading (slow) ===");
    dhoro start = somoy();

    dhoro user = opekha fetchUser(1);
    dhoro posts = opekha fetchPosts(user["id"]);
    dhoro comments = opekha fetchComments(user["id"]);

    dhoro elapsed = somoy() - start;
    dekho("User:", user["naam"]);
    dekho("Posts:", dorghyo(posts));
    dekho("Comments:", dorghyo(comments));
    dekho("Sequential time:", elapsed, "ms");  // ~3300ms
}

// Load dashboard - concurrent approach (fast)
proyash kaj loadDashboardFast() {
    dekho("");
    dekho("=== Concurrent Loading (fast) ===");
    dhoro start = somoy();

    // First get user
    dhoro user = opekha fetchUser(1);

    // Then fetch posts and comments concurrently
    dhoro results = opekha sob_proyash([
        fetchPosts(user["id"]),
        fetchComments(user["id"])
    ]);

    dhoro posts = results[0];
    dhoro comments = results[1];

    dhoro elapsed = somoy() - start;
    dekho("User:", user["naam"]);
    dekho("Posts:", dorghyo(posts));
    dekho("Comments:", dorghyo(comments));
    dekho("Concurrent time:", elapsed, "ms");  // ~2500ms (33% faster!)
}

// Run both approaches
proyash kaj main() {
    opekha loadDashboardSlow();
    opekha loadDashboardFast();
    dekho("");
    dekho("✓ Async demo complete!");
}

main();`}
      />

      <p className="mt-4 text-sm text-muted-foreground">
        This example shows how async/await enables concurrent execution, reducing total wait time
        from 3300ms (sequential) to 2500ms (concurrent) - a 33% performance improvement!
      </p>

      <DocNavigation currentPath="/docs/examples" />
    </div>
  );
}
