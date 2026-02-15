import Link from "next/link";
import CodeBlock from "@/components/CodeBlock";
import DocNavigation from "@/components/DocNavigation";

export default function Database() {
  return (
    <div>
      <div className="flex items-center gap-2 text-sm text-muted-foreground mb-4">
        <span className="px-2 py-1 bg-primary/10 text-primary rounded-full text-xs font-medium">
          Advanced
        </span>
        <span className="px-2 py-1 bg-green-500/10 text-green-600 rounded-full text-xs font-medium">
          Production-Ready
        </span>
      </div>

      <h1>Database Connectivity</h1>

      <p className="lead text-xl text-muted-foreground mt-4">
        BanglaCode provides production-grade database connectors for PostgreSQL, MySQL, MongoDB, and Redis
        with built-in connection pooling, async/await support, and SQL injection protection.
      </p>

      <h2>Supported Databases</h2>

      <p>
        BanglaCode supports four popular databases with both synchronous and asynchronous APIs:
      </p>

      <ul>
        <li><strong>PostgreSQL</strong> - Advanced relational database with ACID compliance</li>
        <li><strong>MySQL</strong> - Popular relational database for web applications</li>
        <li><strong>MongoDB</strong> - NoSQL document database for flexible schemas</li>
        <li><strong>Redis</strong> - In-memory data store for caching and real-time apps</li>
      </ul>

      <h2>Connection Pooling (50-100x Faster)</h2>

      <p>
        Connection pooling is <strong>crucial for performance</strong>. Creating a new database connection
        for every query is extremely slow (~10-50ms per connection). With connection pooling, you reuse
        existing connections, making queries 50-100x faster (~0.1-1ms per query).
      </p>

      <h3>Connection Pool Example</h3>

      <CodeBlock
        filename="database_pool.bang"
        code={`// Create connection pool (reuses connections for maximum performance)
dhoro pool = db_pool_banao("postgres", {
    "host": "localhost",
    "port": 5432,
    "database": "myapp",
    "user": "admin",
    "password": "secret"
}, 10); // Max 10 connections

// Get connection from pool (very fast, ~0.1ms)
dhoro conn = db_pool_nao(pool);

// Execute query
dhoro users = db_query(conn, "SELECT * FROM users WHERE age > 25");

// Iterate results
ghuriye (dhoro i = 0; i < dorghyo(users["rows"]); i = i + 1) {
    dhoro user = users["rows"][i];
    dekho("User:", user["name"], "Age:", user["age"]);
}

// Return connection to pool (important! Connection is reused)
db_pool_ferot(pool, conn);

// Close pool when done
db_pool_bondho(pool);`}
      />

      <h2>PostgreSQL</h2>

      <p>
        PostgreSQL is a powerful, open-source relational database with advanced features like ACID transactions,
        foreign keys, and complex queries.
      </p>

      <h3>Basic PostgreSQL Connection</h3>

      <CodeBlock
        filename="postgres_basic.bang"
        code={`// Connect to PostgreSQL
dhoro conn = db_jukto("postgres", {
    "host": "localhost",
    "port": 5432,
    "database": "myapp",
    "user": "postgres",
    "password": "password"
});

// Execute SELECT query
dhoro result = db_query(conn, "SELECT * FROM users");

dekho("Total users:", dorghyo(result["rows"]));

// Iterate through results
ghuriye (dhoro i = 0; i < dorghyo(result["rows"]); i = i + 1) {
    dhoro user = result["rows"][i];
    dekho("User ID:", user["id"], "Name:", user["name"]);
}

// Close connection
db_bandho(conn);`}
      />

      <h3>Prepared Statements (SQL Injection Safe)</h3>

      <p>
        Always use prepared statements (<code>db_proshno</code>) to prevent SQL injection attacks.
        Never concatenate user input directly into SQL queries!
      </p>

      <CodeBlock
        filename="postgres_prepared.bang"
        code={`// BAD - Vulnerable to SQL injection!
// dhoro query = "SELECT * FROM users WHERE name = '" + userName + "'";

// GOOD - SQL injection safe!
dhoro result = db_proshno(conn, "SELECT * FROM users WHERE name = $1", ["Rahim"]);

// Insert with parameters
dhoro insertResult = db_proshno(conn,
    "INSERT INTO users (name, email, age) VALUES ($1, $2, $3)",
    ["Rahim Ahmed", "rahim@example.com", 30]
);

dekho("Rows affected:", insertResult["rows_affected"]);
dekho("Last insert ID:", insertResult["last_insert_id"]);`}
      />

      <h3>Transactions</h3>

      <p>
        Transactions ensure that multiple database operations either all succeed or all fail together,
        maintaining data integrity.
      </p>

      <CodeBlock
        filename="postgres_transaction.bang"
        code={`// Begin transaction
dhoro tx = db_transaction_shuru_postgres(conn);

chesta {
    // Transfer money between accounts (atomic operation)
    db_exec_postgres(conn, "UPDATE accounts SET balance = balance - 100 WHERE id = 1");
    db_exec_postgres(conn, "UPDATE accounts SET balance = balance + 100 WHERE id = 2");

    // Commit if all operations succeed
    db_commit_postgres(tx);
    dekho("Transaction completed successfully!");

} dhoro_bhul (error) {
    // Rollback if any operation fails
    db_rollback_postgres(tx);
    dekho("Transaction failed and rolled back:", error);
}`}
      />

      <h3>PostgreSQL Functions</h3>

      <div className="overflow-x-auto my-4">
        <table>
          <thead>
            <tr>
              <th>Function</th>
              <th>Parameters</th>
              <th>Description</th>
            </tr>
          </thead>
          <tbody>
            <tr>
              <td><code>db_jukto_postgres</code></td>
              <td><code>config</code></td>
              <td>Connect to PostgreSQL database</td>
            </tr>
            <tr>
              <td><code>db_query_postgres</code></td>
              <td><code>conn, sql</code></td>
              <td>Execute SELECT query</td>
            </tr>
            <tr>
              <td><code>db_exec_postgres</code></td>
              <td><code>conn, sql</code></td>
              <td>Execute INSERT/UPDATE/DELETE</td>
            </tr>
            <tr>
              <td><code>db_proshno_postgres</code></td>
              <td><code>conn, sql, params</code></td>
              <td>Prepared statement (SQL injection safe)</td>
            </tr>
            <tr>
              <td><code>db_transaction_shuru_postgres</code></td>
              <td><code>conn</code></td>
              <td>Begin transaction</td>
            </tr>
            <tr>
              <td><code>db_commit_postgres</code></td>
              <td><code>tx</code></td>
              <td>Commit transaction</td>
            </tr>
            <tr>
              <td><code>db_rollback_postgres</code></td>
              <td><code>tx</code></td>
              <td>Rollback transaction</td>
            </tr>
          </tbody>
        </table>
      </div>

      <h2>MySQL</h2>

      <p>
        MySQL is a popular relational database used by millions of websites and applications.
      </p>

      <CodeBlock
        filename="mysql_example.bang"
        code={`// Connect to MySQL
dhoro conn = db_jukto("mysql", {
    "host": "localhost",
    "port": 3306,
    "database": "webapp",
    "user": "root",
    "password": "password"
});

// Complete CRUD operations
// CREATE - Insert new user
db_proshno(conn, "INSERT INTO users (name, email) VALUES (?, ?)",
    ["Karim", "karim@example.com"]);

// READ - Query users
dhoro users = db_query(conn, "SELECT * FROM users");
ghuriye (dhoro i = 0; i < dorghyo(users["rows"]); i = i + 1) {
    dekho("User:", users["rows"][i]["name"]);
}

// UPDATE - Update user email
db_proshno(conn, "UPDATE users SET email = ? WHERE name = ?",
    ["new@example.com", "Karim"]);

// DELETE - Delete user
db_proshno(conn, "DELETE FROM users WHERE name = ?", ["Karim"]);

db_bandho(conn);`}
      />

      <h2>MongoDB</h2>

      <p>
        MongoDB is a NoSQL document database that stores data in flexible JSON-like documents.
        Perfect for applications with evolving schemas or hierarchical data.
      </p>

      <h3>Find Documents</h3>

      <CodeBlock
        filename="mongodb_find.bang"
        code={`// Connect to MongoDB
dhoro conn = db_jukto("mongodb", {
    "host": "localhost",
    "port": 27017,
    "database": "mydb"
});

// Find documents matching filter
dhoro users = db_khojo_mongodb(conn, "users", {
    "age": {"$gt": 25},         // Age greater than 25
    "city": "Dhaka"              // City is Dhaka
});

dekho("Found", dorghyo(users["rows"]), "users");

// Display results
ghuriye (dhoro i = 0; i < dorghyo(users["rows"]); i = i + 1) {
    dhoro user = users["rows"][i];
    dekho("Name:", user["name"], "Age:", user["age"], "City:", user["city"]);
}`}
      />

      <h3>Insert Documents</h3>

      <CodeBlock
        filename="mongodb_insert.bang"
        code={`// Insert single document
db_dhokao_mongodb(conn, "users", {
    "name": "Rahim Ahmed",
    "age": 30,
    "city": "Dhaka",
    "profession": "Engineer",
    "skills": ["Python", "Go", "BanglaCode"]
});

dekho("User inserted!");`}
      />

      <h3>Update and Delete</h3>

      <CodeBlock
        filename="mongodb_update.bang"
        code={`// Update documents
dhoro updateResult = db_update_mongodb(conn, "users",
    {"city": "Dhaka"},                    // Filter
    {"$set": {"country": "Bangladesh"}}   // Update
);

dekho("Updated", updateResult["rows_affected"], "documents");

// Delete documents
dhoro deleteResult = db_mujhe_mongodb(conn, "users", {
    "age": {"$lt": 18}  // Delete users under 18
});

dekho("Deleted", deleteResult["rows_affected"], "documents");`}
      />

      <h3>MongoDB Functions</h3>

      <div className="overflow-x-auto my-4">
        <table>
          <thead>
            <tr>
              <th>Function</th>
              <th>Parameters</th>
              <th>Description</th>
            </tr>
          </thead>
          <tbody>
            <tr>
              <td><code>db_jukto_mongodb</code></td>
              <td><code>config</code></td>
              <td>Connect to MongoDB</td>
            </tr>
            <tr>
              <td><code>db_khojo_mongodb</code></td>
              <td><code>conn, collection, filter</code></td>
              <td>Find documents</td>
            </tr>
            <tr>
              <td><code>db_dhokao_mongodb</code></td>
              <td><code>conn, collection, doc</code></td>
              <td>Insert document</td>
            </tr>
            <tr>
              <td><code>db_update_mongodb</code></td>
              <td><code>conn, collection, filter, update</code></td>
              <td>Update documents</td>
            </tr>
            <tr>
              <td><code>db_mujhe_mongodb</code></td>
              <td><code>conn, collection, filter</code></td>
              <td>Delete documents</td>
            </tr>
          </tbody>
        </table>
      </div>

      <h2>Redis</h2>

      <p>
        Redis is an in-memory data store used for caching, session management, real-time analytics,
        and message queues. Extremely fast (sub-millisecond response times).
      </p>

      <h3>Key-Value Operations</h3>

      <CodeBlock
        filename="redis_kv.bang"
        code={`// Connect to Redis
dhoro conn = db_jukto("redis", {
    "host": "localhost",
    "port": 6379
});

// Set key-value
db_set_redis(conn, "user:1", "Rahim Ahmed");

// Set with TTL (time to live - expires after 1 hour)
db_set_redis(conn, "session:abc123", "user_data", 3600);

// Get value
dhoro user = db_get_redis(conn, "user:1");
dekho("User:", user);

// Delete key
db_del_redis(conn, "user:1");

// Set expiration on existing key
db_expire_redis(conn, "session:abc123", 1800); // 30 minutes`}
      />

      <h3>List Operations (Queue)</h3>

      <CodeBlock
        filename="redis_list.bang"
        code={`// Redis lists work as queues (FIFO) or stacks (LIFO)

// Add tasks to queue
db_rpush_redis(conn, "tasks:queue", "Process order #123");
db_rpush_redis(conn, "tasks:queue", "Send email to user");
db_rpush_redis(conn, "tasks:queue", "Generate report");

// Process tasks from queue
dhoro task1 = db_lpop_redis(conn, "tasks:queue");
dekho("Processing:", task1); // Output: Process order #123

dhoro task2 = db_lpop_redis(conn, "tasks:queue");
dekho("Processing:", task2); // Output: Send email to user`}
      />

      <h3>Hash Operations</h3>

      <CodeBlock
        filename="redis_hash.bang"
        code={`// Store user profile as hash (like a map/object)
db_hset_redis(conn, "user:1:profile", "name", "Rahim Ahmed");
db_hset_redis(conn, "user:1:profile", "age", "30");
db_hset_redis(conn, "user:1:profile", "city", "Dhaka");

// Get single field
dhoro name = db_hget_redis(conn, "user:1:profile", "name");
dekho("Name:", name);

// Get all fields
dhoro profile = db_hgetall_redis(conn, "user:1:profile");
dekho("Full profile:", profile);`}
      />

      <h3>Redis Functions</h3>

      <div className="overflow-x-auto my-4">
        <table>
          <thead>
            <tr>
              <th>Function</th>
              <th>Parameters</th>
              <th>Description</th>
            </tr>
          </thead>
          <tbody>
            <tr>
              <td><code>db_jukto_redis</code></td>
              <td><code>config</code></td>
              <td>Connect to Redis</td>
            </tr>
            <tr>
              <td><code>db_set_redis</code></td>
              <td><code>conn, key, value, ttl?</code></td>
              <td>Set key-value (optional TTL in seconds)</td>
            </tr>
            <tr>
              <td><code>db_get_redis</code></td>
              <td><code>conn, key</code></td>
              <td>Get value by key</td>
            </tr>
            <tr>
              <td><code>db_del_redis</code></td>
              <td><code>conn, key</code></td>
              <td>Delete key</td>
            </tr>
            <tr>
              <td><code>db_expire_redis</code></td>
              <td><code>conn, key, seconds</code></td>
              <td>Set expiration time</td>
            </tr>
            <tr>
              <td><code>db_lpush_redis</code></td>
              <td><code>conn, key, value</code></td>
              <td>Push to list (left/front)</td>
            </tr>
            <tr>
              <td><code>db_rpush_redis</code></td>
              <td><code>conn, key, value</code></td>
              <td>Push to list (right/back)</td>
            </tr>
            <tr>
              <td><code>db_lpop_redis</code></td>
              <td><code>conn, key</code></td>
              <td>Pop from list (left/front)</td>
            </tr>
            <tr>
              <td><code>db_rpop_redis</code></td>
              <td><code>conn, key</code></td>
              <td>Pop from list (right/back)</td>
            </tr>
            <tr>
              <td><code>db_hset_redis</code></td>
              <td><code>conn, key, field, value</code></td>
              <td>Set hash field</td>
            </tr>
            <tr>
              <td><code>db_hget_redis</code></td>
              <td><code>conn, key, field</code></td>
              <td>Get hash field</td>
            </tr>
            <tr>
              <td><code>db_hgetall_redis</code></td>
              <td><code>conn, key</code></td>
              <td>Get all hash fields</td>
            </tr>
          </tbody>
        </table>
      </div>

      <h2>Universal Database Functions</h2>

      <p>
        These functions work with all supported databases (PostgreSQL, MySQL, MongoDB, Redis).
        The <code>db_jukto</code> function automatically routes to the correct database driver.
      </p>

      <div className="overflow-x-auto my-4">
        <table>
          <thead>
            <tr>
              <th>Function</th>
              <th>Parameters</th>
              <th>Description</th>
            </tr>
          </thead>
          <tbody>
            <tr>
              <td><code>db_jukto</code></td>
              <td><code>type, config</code></td>
              <td>Connect to database (type: &quot;postgres&quot;, &quot;mysql&quot;, &quot;mongodb&quot;, &quot;redis&quot;)</td>
            </tr>
            <tr>
              <td><code>db_jukto_async</code></td>
              <td><code>type, config</code></td>
              <td>Connect to database (async, returns promise)</td>
            </tr>
            <tr>
              <td><code>db_bandho</code></td>
              <td><code>conn</code></td>
              <td>Close connection</td>
            </tr>
            <tr>
              <td><code>db_bandho_async</code></td>
              <td><code>conn</code></td>
              <td>Close connection (async)</td>
            </tr>
            <tr>
              <td><code>db_query</code></td>
              <td><code>conn, sql</code></td>
              <td>Execute SELECT query (SQL databases only)</td>
            </tr>
            <tr>
              <td><code>db_query_async</code></td>
              <td><code>conn, sql</code></td>
              <td>Execute SELECT query async</td>
            </tr>
            <tr>
              <td><code>db_exec</code></td>
              <td><code>conn, sql</code></td>
              <td>Execute INSERT/UPDATE/DELETE (SQL databases only)</td>
            </tr>
            <tr>
              <td><code>db_exec_async</code></td>
              <td><code>conn, sql</code></td>
              <td>Execute INSERT/UPDATE/DELETE async</td>
            </tr>
            <tr>
              <td><code>db_proshno</code></td>
              <td><code>conn, sql, params</code></td>
              <td>Prepared statement (SQL injection safe)</td>
            </tr>
            <tr>
              <td><code>db_proshno_async</code></td>
              <td><code>conn, sql, params</code></td>
              <td>Prepared statement async</td>
            </tr>
          </tbody>
        </table>
      </div>

      <h2>Connection Pool Functions</h2>

      <div className="overflow-x-auto my-4">
        <table>
          <thead>
            <tr>
              <th>Function</th>
              <th>Parameters</th>
              <th>Description</th>
            </tr>
          </thead>
          <tbody>
            <tr>
              <td><code>db_pool_banao</code></td>
              <td><code>type, config, maxConns</code></td>
              <td>Create connection pool</td>
            </tr>
            <tr>
              <td><code>db_pool_nao</code></td>
              <td><code>pool</code></td>
              <td>Get connection from pool</td>
            </tr>
            <tr>
              <td><code>db_pool_ferot</code></td>
              <td><code>pool, conn</code></td>
              <td>Return connection to pool (important!)</td>
            </tr>
            <tr>
              <td><code>db_pool_bondho</code></td>
              <td><code>pool</code></td>
              <td>Close connection pool</td>
            </tr>
            <tr>
              <td><code>db_pool_tothyo</code></td>
              <td><code>pool</code></td>
              <td>Get pool statistics</td>
            </tr>
          </tbody>
        </table>
      </div>

      <h2>Async Database Queries</h2>

      <p>
        Use <code>proyash</code> and <code>opekha</code> for async database operations.
        This is essential for non-blocking I/O and high-performance applications.
      </p>

      <CodeBlock
        filename="async_database.bang"
        code={`// Async database function
proyash kaj fetchUsers() {
    // Connect async
    dhoro conn = opekha db_jukto_async("postgres", {
        "host": "localhost",
        "database": "myapp"
    });

    // Query async
    dhoro users = opekha db_query_async(conn, "SELECT * FROM users");

    // Close async
    opekha db_bandho_async(conn);

    ferao users;
}

// Call async function
dhoro result = opekha fetchUsers();
dekho("Fetched", dorghyo(result["rows"]), "users");

// Multiple concurrent queries
proyash kaj fetchMultiple() {
    dhoro conn = opekha db_jukto_async("postgres", {...});

    // Run multiple queries concurrently
    dhoro users = opekha db_query_async(conn, "SELECT * FROM users");
    dhoro posts = opekha db_query_async(conn, "SELECT * FROM posts");

    opekha db_bandho_async(conn);

    ferao {"users": users, "posts": posts};
}`}
      />

      <h2>Function Name Meanings</h2>

      <div className="overflow-x-auto my-4">
        <table>
          <thead>
            <tr>
              <th>Function</th>
              <th>Bengali</th>
              <th>Meaning</th>
            </tr>
          </thead>
          <tbody>
            <tr><td><code>jukto</code></td><td>যুক্ত</td><td>connect/join</td></tr>
            <tr><td><code>bandho</code></td><td>বন্ধ</td><td>close</td></tr>
            <tr><td><code>proshno</code></td><td>প্রশ্ন</td><td>question/query</td></tr>
            <tr><td><code>banao</code></td><td>বানাও</td><td>create/make</td></tr>
            <tr><td><code>nao</code></td><td>নাও</td><td>take/get</td></tr>
            <tr><td><code>ferot</code></td><td>ফেরত</td><td>return</td></tr>
            <tr><td><code>tothyo</code></td><td>তথ্য</td><td>information/data</td></tr>
            <tr><td><code>khojo</code></td><td>খোঁজো</td><td>search/find</td></tr>
            <tr><td><code>dhokao</code></td><td>ঢোকাও</td><td>insert/put</td></tr>
            <tr><td><code>mujhe</code></td><td>মুছে</td><td>delete/erase</td></tr>
          </tbody>
        </table>
      </div>

      <h2>Best Practices</h2>

      <ul>
        <li><strong>Always use connection pooling</strong> for production applications (50-100x faster)</li>
        <li><strong>Always use prepared statements</strong> (<code>db_proshno</code>) to prevent SQL injection</li>
        <li><strong>Always close connections</strong> or return them to the pool when done</li>
        <li>Use <strong>transactions</strong> for operations that must succeed or fail together</li>
        <li>Use <strong>async/await</strong> for non-blocking database operations</li>
        <li>Choose the right database for your needs:
          <ul>
            <li><strong>PostgreSQL/MySQL</strong> - Structured data, complex queries, ACID compliance</li>
            <li><strong>MongoDB</strong> - Flexible schemas, hierarchical data, rapid iteration</li>
            <li><strong>Redis</strong> - Caching, sessions, real-time data, message queues</li>
          </ul>
        </li>
      </ul>

      <h2>Performance Tips</h2>

      <ul>
        <li>Connection pooling reduces connection overhead by 50-100x (from ~10ms to ~0.1ms)</li>
        <li>Use <code>db_proshno</code> for repeated queries (2-5x faster than concatenated SQL)</li>
        <li>Index frequently queried columns in SQL databases</li>
        <li>Use Redis for frequently accessed data (sub-millisecond response times)</li>
        <li>Batch operations when possible to reduce round trips</li>
      </ul>

      <h2>Next Steps</h2>

      <p>
        Learn more about related topics:
      </p>

      <ul>
        <li><Link href="/docs/async-await" className="text-primary hover:underline">Async/Await</Link> - Understanding promises and async operations</li>
        <li><Link href="/docs/error-handling" className="text-primary hover:underline">Error Handling</Link> - Try/catch/finally for error handling</li>
        <li><Link href="/docs/builtins" className="text-primary hover:underline">Built-in Functions</Link> - All 130+ built-in functions</li>
        <li><Link href="/docs/examples" className="text-primary hover:underline">Examples</Link> - More code examples</li>
      </ul>

      <DocNavigation currentPath="/docs/database" />
    </div>
  );
}
