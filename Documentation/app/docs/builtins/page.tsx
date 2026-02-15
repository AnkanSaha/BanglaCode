import CodeBlock from "@/components/CodeBlock";
import DocNavigation from "@/components/DocNavigation";

export default function Builtins() {
  return (
    <div>
      <div className="flex items-center gap-2 text-sm text-muted-foreground mb-4">
        <span className="px-2 py-1 bg-primary/10 text-primary rounded-full text-xs font-medium">
          Reference
        </span>
      </div>

      <h1>Built-in Functions Reference</h1>

      <p className="lead text-xl text-muted-foreground mt-4">
        BanglaCode provides 135+ built-in functions for I/O, type conversion, string manipulation,
        array operations, math, file handling, HTTP, JSON, environment variables, networking (TCP/UDP/WebSocket),
        database connectivity (PostgreSQL/MySQL/MongoDB/Redis), and complete OS-level system access.
      </p>

      <h2>Input/Output</h2>

      <div className="overflow-x-auto my-4">
        <table>
          <thead>
            <tr>
              <th>Function</th>
              <th>Parameters</th>
              <th>Returns</th>
              <th>Description</th>
            </tr>
          </thead>
          <tbody>
            <tr>
              <td><code>dekho</code></td>
              <td><code>args...</code></td>
              <td><code>khali</code></td>
              <td>Print values to console</td>
            </tr>
            <tr>
              <td><code>nao</code></td>
              <td><code>[prompt]</code></td>
              <td><code>string</code></td>
              <td>Read user input from console</td>
            </tr>
          </tbody>
        </table>
      </div>

      <CodeBlock
        code={`// Print output
dekho("Hello");
dekho("Name:", "Rahim", "Age:", 25);

// Read input
dhoro naam = nao("Enter your name: ");
dekho("Hello,", naam);`}
      />

      <h2>Type Functions</h2>

      <div className="overflow-x-auto my-4">
        <table>
          <thead>
            <tr>
              <th>Function</th>
              <th>Parameters</th>
              <th>Returns</th>
              <th>Description</th>
            </tr>
          </thead>
          <tbody>
            <tr>
              <td><code>dhoron</code></td>
              <td><code>value</code></td>
              <td><code>string</code></td>
              <td>Get type of value</td>
            </tr>
            <tr>
              <td><code>lipi</code></td>
              <td><code>value</code></td>
              <td><code>string</code></td>
              <td>Convert to string</td>
            </tr>
            <tr>
              <td><code>sonkha</code></td>
              <td><code>value</code></td>
              <td><code>number</code></td>
              <td>Convert to number</td>
            </tr>
          </tbody>
        </table>
      </div>

      <CodeBlock
        code={`// Type checking
dekho(dhoron(42));        // "int"
dekho(dhoron(3.14));      // "float"
dekho(dhoron("hello"));   // "string"
dekho(dhoron(sotti));     // "boolean"
dekho(dhoron([1,2,3]));   // "array"
dekho(dhoron({x:1}));     // "map"

// Convert to string
dekho(lipi(42));          // "42"
dekho(lipi(sotti));       // "true"

// Convert to number
dekho(sonkha("123"));     // 123
dekho(sonkha("3.14"));    // 3.14`}
      />

      <h2>String Functions</h2>

      <div className="overflow-x-auto my-4">
        <table>
          <thead>
            <tr>
              <th>Function</th>
              <th>Parameters</th>
              <th>Returns</th>
              <th>Description</th>
            </tr>
          </thead>
          <tbody>
            <tr>
              <td><code>dorghyo</code></td>
              <td><code>str</code></td>
              <td><code>int</code></td>
              <td>Get string length</td>
            </tr>
            <tr>
              <td><code>boroHater</code></td>
              <td><code>str</code></td>
              <td><code>string</code></td>
              <td>Convert to uppercase</td>
            </tr>
            <tr>
              <td><code>chotoHater</code></td>
              <td><code>str</code></td>
              <td><code>string</code></td>
              <td>Convert to lowercase</td>
            </tr>
            <tr>
              <td><code>bhag</code></td>
              <td><code>str, sep</code></td>
              <td><code>array</code></td>
              <td>Split string by separator</td>
            </tr>
            <tr>
              <td><code>joro</code></td>
              <td><code>arr, sep</code></td>
              <td><code>string</code></td>
              <td>Join array elements</td>
            </tr>
            <tr>
              <td><code>chhanto</code></td>
              <td><code>str</code></td>
              <td><code>string</code></td>
              <td>Trim whitespace</td>
            </tr>
            <tr>
              <td><code>khojo</code></td>
              <td><code>str, substr</code></td>
              <td><code>int</code></td>
              <td>Find index of substring (-1 if not found)</td>
            </tr>
            <tr>
              <td><code>angsho</code></td>
              <td><code>str, start, [end]</code></td>
              <td><code>string</code></td>
              <td>Get substring</td>
            </tr>
            <tr>
              <td><code>bodlo</code></td>
              <td><code>str, old, new</code></td>
              <td><code>string</code></td>
              <td>Replace all occurrences</td>
            </tr>
          </tbody>
        </table>
      </div>

      <CodeBlock
        code={`dhoro str = "  Hello World  ";

dekho(dorghyo(str));            // 15
dekho(chhanto(str));            // "Hello World"
dekho(boroHater("hello"));      // "HELLO"
dekho(chotoHater("HELLO"));     // "hello"
dekho(bhag("a,b,c", ","));      // ["a", "b", "c"]
dekho(joro(["a","b","c"], "-"));// "a-b-c"
dekho(khojo("hello", "ll"));    // 2
dekho(angsho("hello", 1, 4));   // "ell"
dekho(bodlo("aaa", "a", "b"));  // "bbb"`}
      />

      <h2>Array Functions</h2>

      <div className="overflow-x-auto my-4">
        <table>
          <thead>
            <tr>
              <th>Function</th>
              <th>Parameters</th>
              <th>Returns</th>
              <th>Description</th>
            </tr>
          </thead>
          <tbody>
            <tr>
              <td><code>dorghyo</code></td>
              <td><code>arr</code></td>
              <td><code>int</code></td>
              <td>Get array length</td>
            </tr>
            <tr>
              <td><code>dhokao</code></td>
              <td><code>arr, item</code></td>
              <td><code>array</code></td>
              <td>Push item to end</td>
            </tr>
            <tr>
              <td><code>berKoro</code></td>
              <td><code>arr</code></td>
              <td><code>any</code></td>
              <td>Pop item from end</td>
            </tr>
            <tr>
              <td><code>kato</code></td>
              <td><code>arr, start, [end]</code></td>
              <td><code>array</code></td>
              <td>Slice array</td>
            </tr>
            <tr>
              <td><code>ulto</code></td>
              <td><code>arr</code></td>
              <td><code>array</code></td>
              <td>Reverse array</td>
            </tr>
            <tr>
              <td><code>ache</code></td>
              <td><code>arr, item</code></td>
              <td><code>boolean</code></td>
              <td>Check if item exists</td>
            </tr>
            <tr>
              <td><code>saja</code></td>
              <td><code>arr</code></td>
              <td><code>array</code></td>
              <td>Sort array</td>
            </tr>
          </tbody>
        </table>
      </div>

      <CodeBlock
        code={`dhoro arr = [3, 1, 4, 1, 5];

dekho(dorghyo(arr));       // 5
dhokao(arr, 9);            // Push 9
dekho(berKoro(arr));       // 9 (and removes it)
dekho(kato(arr, 1, 3));    // [1, 4]
dekho(ulto(arr));          // [5, 1, 4, 1, 3]
dekho(ache(arr, 4));       // sotti
dekho(saja([3,1,2]));      // [1, 2, 3]`}
      />

      <h2>Math Functions</h2>

      <div className="overflow-x-auto my-4">
        <table>
          <thead>
            <tr>
              <th>Function</th>
              <th>Parameters</th>
              <th>Returns</th>
              <th>Description</th>
            </tr>
          </thead>
          <tbody>
            <tr>
              <td><code>borgomul</code></td>
              <td><code>n</code></td>
              <td><code>float</code></td>
              <td>Square root</td>
            </tr>
            <tr>
              <td><code>ghat</code></td>
              <td><code>base, exp</code></td>
              <td><code>number</code></td>
              <td>Power (base^exp)</td>
            </tr>
            <tr>
              <td><code>niche</code></td>
              <td><code>n</code></td>
              <td><code>int</code></td>
              <td>Floor (round down)</td>
            </tr>
            <tr>
              <td><code>upore</code></td>
              <td><code>n</code></td>
              <td><code>int</code></td>
              <td>Ceiling (round up)</td>
            </tr>
            <tr>
              <td><code>kache</code></td>
              <td><code>n</code></td>
              <td><code>int</code></td>
              <td>Round to nearest</td>
            </tr>
            <tr>
              <td><code>niratek</code></td>
              <td><code>n</code></td>
              <td><code>number</code></td>
              <td>Absolute value</td>
            </tr>
            <tr>
              <td><code>choto</code></td>
              <td><code>...nums</code></td>
              <td><code>number</code></td>
              <td>Minimum value</td>
            </tr>
            <tr>
              <td><code>boro</code></td>
              <td><code>...nums</code></td>
              <td><code>number</code></td>
              <td>Maximum value</td>
            </tr>
            <tr>
              <td><code>lotto</code></td>
              <td><code>(none)</code></td>
              <td><code>float</code></td>
              <td>Random number [0, 1)</td>
            </tr>
          </tbody>
        </table>
      </div>

      <CodeBlock
        code={`dekho(borgomul(16));       // 4
dekho(ghat(2, 10));        // 1024
dekho(niche(3.7));         // 3
dekho(upore(3.2));         // 4
dekho(kache(3.5));         // 4
dekho(niratek(-5));        // 5
dekho(choto(5, 2, 8, 1));  // 1
dekho(boro(5, 2, 8, 1));   // 8
dekho(lotto());            // 0.xxxxx (random)`}
      />

      <h2>Map Functions</h2>

      <div className="overflow-x-auto my-4">
        <table>
          <thead>
            <tr>
              <th>Function</th>
              <th>Parameters</th>
              <th>Returns</th>
              <th>Description</th>
            </tr>
          </thead>
          <tbody>
            <tr>
              <td><code>chabi</code></td>
              <td><code>map</code></td>
              <td><code>array</code></td>
              <td>Get all keys as array</td>
            </tr>
          </tbody>
        </table>
      </div>

      <CodeBlock
        code={`dhoro obj = {naam: "Rahim", boyosh: 25};
dekho(chabi(obj));  // ["naam", "boyosh"]`}
      />

      <h2>File I/O Functions</h2>

      <div className="overflow-x-auto my-4">
        <table>
          <thead>
            <tr>
              <th>Function</th>
              <th>Parameters</th>
              <th>Returns</th>
              <th>Description</th>
            </tr>
          </thead>
          <tbody>
            <tr>
              <td><code>poro</code></td>
              <td><code>path</code></td>
              <td><code>string</code></td>
              <td>Read file content</td>
            </tr>
            <tr>
              <td><code>lekho</code></td>
              <td><code>path, content</code></td>
              <td><code>boolean</code></td>
              <td>Write content to file</td>
            </tr>
          </tbody>
        </table>
      </div>

      <CodeBlock
        code={`// Read file
dhoro content = poro("data.txt");
dekho(content);

// Write file
lekho("output.txt", "Hello World");`}
      />

      <h2>JSON Functions</h2>

      <div className="overflow-x-auto my-4">
        <table>
          <thead>
            <tr>
              <th>Function</th>
              <th>Parameters</th>
              <th>Returns</th>
              <th>Description</th>
            </tr>
          </thead>
          <tbody>
            <tr>
              <td><code>json_poro</code></td>
              <td><code>str</code></td>
              <td><code>any</code></td>
              <td>Parse JSON string</td>
            </tr>
            <tr>
              <td><code>json_banao</code></td>
              <td><code>value</code></td>
              <td><code>string</code></td>
              <td>Convert to JSON string</td>
            </tr>
          </tbody>
        </table>
      </div>

      <CodeBlock
        code={`// Parse JSON
dhoro obj = json_poro('{"name":"Rahim","age":25}');
dekho(obj.name);  // "Rahim"

// Create JSON
dhoro json = json_banao({x: 1, y: 2});
dekho(json);  // {"x":1,"y":2}`}
      />

      <h2>Environment Variables</h2>

      <p>
        See the <a href="/docs/environment-variables" className="text-primary hover:underline">Environment Variables Documentation</a> for detailed examples and multi-environment setup.
      </p>

      <div className="overflow-x-auto my-4">
        <table>
          <thead>
            <tr>
              <th>Function</th>
              <th>Parameters</th>
              <th>Returns</th>
              <th>Description</th>
            </tr>
          </thead>
          <tbody>
            <tr>
              <td><code>env_load</code></td>
              <td><code>filename</code></td>
              <td><code>boolean</code></td>
              <td>Load environment variables from .env file</td>
            </tr>
            <tr>
              <td><code>env_load_auto</code></td>
              <td><code>environment</code></td>
              <td><code>string</code></td>
              <td>Auto-load .env.{'{environment}'} or fallback to .env</td>
            </tr>
            <tr>
              <td><code>env_get</code></td>
              <td><code>key</code></td>
              <td><code>string</code></td>
              <td>Get environment variable (error if not found)</td>
            </tr>
            <tr>
              <td><code>env_get_default</code></td>
              <td><code>key, default</code></td>
              <td><code>string</code></td>
              <td>Get environment variable with default fallback</td>
            </tr>
            <tr>
              <td><code>env_set</code></td>
              <td><code>key, value</code></td>
              <td><code>boolean</code></td>
              <td>Set environment variable at runtime</td>
            </tr>
            <tr>
              <td><code>env_all</code></td>
              <td><code>(none)</code></td>
              <td><code>map</code></td>
              <td>Get all environment variables as map</td>
            </tr>
            <tr>
              <td><code>env_clear</code></td>
              <td><code>(none)</code></td>
              <td><code>boolean</code></td>
              <td>Clear all loaded environment variables</td>
            </tr>
          </tbody>
        </table>
      </div>

      <CodeBlock
        code={`// Load .env file
env_load(".env");

// Get environment variables
dhoro api_key = env_get("API_KEY");
dhoro api_url = env_get_default("API_URL", "http://localhost:3000");

dekho("API URL:", api_url);

// Multi-environment support
env_load_auto("prod");  // Loads .env.prod or falls back to .env
dhoro db_host = env_get("DB_HOST");
dekho("Database Host:", db_host);

// Set runtime variable
env_set("SESSION_ID", "abc123");

// Get all variables
dhoro all = env_all();
dekho("All env vars:", all);`}
      />

      <h2>Utility Functions</h2>

      <div className="overflow-x-auto my-4">
        <table>
          <thead>
            <tr>
              <th>Function</th>
              <th>Parameters</th>
              <th>Returns</th>
              <th>Description</th>
            </tr>
          </thead>
          <tbody>
            <tr>
              <td><code>somoy</code></td>
              <td><code>(none)</code></td>
              <td><code>int</code></td>
              <td>Current time in milliseconds</td>
            </tr>
            <tr>
              <td><code>ghum</code></td>
              <td><code>ms</code></td>
              <td><code>khali</code></td>
              <td>Sleep/delay for milliseconds</td>
            </tr>
            <tr>
              <td><code>bondho</code></td>
              <td><code>[code]</code></td>
              <td><code>(exit)</code></td>
              <td>Exit program</td>
            </tr>
          </tbody>
        </table>
      </div>

      <CodeBlock
        code={`// Current time
dhoro now = somoy();
dekho("Timestamp:", now);

// Sleep 1 second
dekho("Waiting...");
ghum(1000);
dekho("Done!");

// Exit program
bondho(0);  // Exit with code 0`}
      />

      <h2>HTTP Functions</h2>

      <div className="overflow-x-auto my-4">
        <table>
          <thead>
            <tr>
              <th>Function</th>
              <th>Parameters</th>
              <th>Returns</th>
              <th>Description</th>
            </tr>
          </thead>
          <tbody>
            <tr>
              <td><code>server_chalu</code></td>
              <td><code>port, handler</code></td>
              <td><code>khali</code></td>
              <td>Start HTTP server</td>
            </tr>
            <tr>
              <td><code>anun</code></td>
              <td><code>url</code></td>
              <td><code>map</code></td>
              <td>HTTP GET request</td>
            </tr>
            <tr>
              <td><code>uttor</code></td>
              <td><code>res, body, [status], [type]</code></td>
              <td><code>map</code></td>
              <td>Send HTTP response</td>
            </tr>
            <tr>
              <td><code>json_uttor</code></td>
              <td><code>res, data, [status]</code></td>
              <td><code>map</code></td>
              <td>Send JSON response</td>
            </tr>
          </tbody>
        </table>
      </div>

      <CodeBlock
        code={`// Start server
server_chalu(8080, kaj(req, res) {
    json_uttor(res, {message: "Hello!"});
});

// Make HTTP request
dhoro response = anun("https://api.example.com/data");
dekho(response.body);`}
      />

      <h2>Async/Promise Functions</h2>

      <p>
        Asynchronous functions that return promises. Use with <code>proyash</code>/<code>opekha</code> keywords.
        See the <a href="/docs/async-await" className="text-primary hover:underline">Async/Await documentation</a> for detailed examples.
      </p>

      <div className="overflow-x-auto my-4">
        <table>
          <thead>
            <tr>
              <th>Function</th>
              <th>Parameters</th>
              <th>Returns</th>
              <th>Description</th>
            </tr>
          </thead>
          <tbody>
            <tr>
              <td><code>ghumaao</code></td>
              <td><code>milliseconds</code></td>
              <td><code>Promise</code></td>
              <td>Sleep for specified time (async)</td>
            </tr>
            <tr>
              <td><code>sob_proyash</code></td>
              <td><code>array</code></td>
              <td><code>Promise</code></td>
              <td>Wait for all promises concurrently (Promise.all)</td>
            </tr>
            <tr>
              <td><code>poro_async</code></td>
              <td><code>path</code></td>
              <td><code>Promise</code></td>
              <td>Read file asynchronously</td>
            </tr>
            <tr>
              <td><code>lekho_async</code></td>
              <td><code>path, content</code></td>
              <td><code>Promise</code></td>
              <td>Write file asynchronously</td>
            </tr>
            <tr>
              <td><code>anun_async</code></td>
              <td><code>url</code></td>
              <td><code>Promise</code></td>
              <td>HTTP GET request asynchronously</td>
            </tr>
          </tbody>
        </table>
      </div>

      <CodeBlock
        code={`// Sleep for 1 second
proyash kaj timer() {
    dekho("Start");
    opekha ghumaao(1000);
    dekho("1 second later");
}

// Run multiple operations concurrently
proyash kaj parallel() {
    dhoro results = opekha sob_proyash([
        ghumaao(500),
        poro_async("file.txt"),
        anun_async("https://api.example.com")
    ]);
    dekho("All operations complete!");
}

timer();
parallel();`}
      />

      <h2>Networking Functions</h2>

      <p className="text-muted-foreground">
        BanglaCode provides comprehensive networking capabilities for TCP, UDP, and WebSocket protocols,
        making network programming as easy as JavaScript/Node.js.
      </p>

      <h3>TCP Functions (6 functions)</h3>

      <div className="overflow-x-auto my-4">
        <table>
          <thead>
            <tr>
              <th>Function</th>
              <th>Parameters</th>
              <th>Returns</th>
              <th>Description</th>
            </tr>
          </thead>
          <tbody>
            <tr>
              <td><code>tcp_server_chalu</code></td>
              <td><code>port, handler</code></td>
              <td><code>khali</code></td>
              <td>Start TCP server with callback</td>
            </tr>
            <tr>
              <td><code>tcp_jukto</code></td>
              <td><code>host, port</code></td>
              <td><code>Promise</code></td>
              <td>Connect to TCP server (async)</td>
            </tr>
            <tr>
              <td><code>tcp_pathao</code></td>
              <td><code>connection, data</code></td>
              <td><code>khali</code></td>
              <td>Send data on TCP connection</td>
            </tr>
            <tr>
              <td><code>tcp_lekho</code></td>
              <td><code>connection, data</code></td>
              <td><code>khali</code></td>
              <td>Write data to TCP connection (alias)</td>
            </tr>
            <tr>
              <td><code>tcp_shuno</code></td>
              <td><code>connection</code></td>
              <td><code>Promise</code></td>
              <td>Read data from TCP connection (async)</td>
            </tr>
            <tr>
              <td><code>tcp_bondho</code></td>
              <td><code>connection</code></td>
              <td><code>khali</code></td>
              <td>Close TCP connection</td>
            </tr>
          </tbody>
        </table>
      </div>

      <CodeBlock
        code={`// TCP Server
tcp_server_chalu(8080, kaj(conn) {
    dekho("Client:", conn["remote_addr"]);
    dekho("Data:", conn["data"]);
    tcp_pathao(conn, "Echo: " + conn["data"]);
});

// TCP Client
proyash kaj client() {
    dhoro conn = opekha tcp_jukto("localhost", 8080);
    tcp_lekho(conn, "Hello!");
    dhoro response = opekha tcp_shuno(conn);
    dekho(response);
    tcp_bondho(conn);
}
client();`}
      />

      <h3>UDP Functions (5 functions)</h3>

      <div className="overflow-x-auto my-4">
        <table>
          <thead>
            <tr>
              <th>Function</th>
              <th>Parameters</th>
              <th>Returns</th>
              <th>Description</th>
            </tr>
          </thead>
          <tbody>
            <tr>
              <td><code>udp_server_chalu</code></td>
              <td><code>port, handler</code></td>
              <td><code>khali</code></td>
              <td>Start UDP server with callback</td>
            </tr>
            <tr>
              <td><code>udp_pathao</code></td>
              <td><code>host, port, data</code></td>
              <td><code>Promise</code></td>
              <td>Send UDP packet (async)</td>
            </tr>
            <tr>
              <td><code>udp_uttor</code></td>
              <td><code>packet, data</code></td>
              <td><code>khali</code></td>
              <td>Send UDP response to client</td>
            </tr>
            <tr>
              <td><code>udp_shuno</code></td>
              <td><code>port, handler</code></td>
              <td><code>khali</code></td>
              <td>Listen for UDP packets (alias)</td>
            </tr>
            <tr>
              <td><code>udp_bondho</code></td>
              <td><code>connection</code></td>
              <td><code>khali</code></td>
              <td>Close UDP connection</td>
            </tr>
          </tbody>
        </table>
      </div>

      <CodeBlock
        code={`// UDP Server
udp_server_chalu(9000, kaj(packet) {
    dekho("From:", packet["remote_addr"]);
    dekho("Data:", packet["data"]);
    udp_uttor(packet, "Got it!");
});

// UDP Client
proyash kaj send() {
    opekha udp_pathao("localhost", 9000, "Hello UDP!");
}
send();`}
      />

      <h3>WebSocket Functions (4 functions)</h3>

      <div className="overflow-x-auto my-4">
        <table>
          <thead>
            <tr>
              <th>Function</th>
              <th>Parameters</th>
              <th>Returns</th>
              <th>Description</th>
            </tr>
          </thead>
          <tbody>
            <tr>
              <td><code>websocket_server_chalu</code></td>
              <td><code>port, handler</code></td>
              <td><code>khali</code></td>
              <td>Start WebSocket server</td>
            </tr>
            <tr>
              <td><code>websocket_jukto</code></td>
              <td><code>url</code></td>
              <td><code>Promise</code></td>
              <td>Connect to WebSocket server (async)</td>
            </tr>
            <tr>
              <td><code>websocket_pathao</code></td>
              <td><code>connection, message</code></td>
              <td><code>khali</code></td>
              <td>Send WebSocket message</td>
            </tr>
            <tr>
              <td><code>websocket_bondho</code></td>
              <td><code>connection</code></td>
              <td><code>khali</code></td>
              <td>Close WebSocket connection</td>
            </tr>
          </tbody>
        </table>
      </div>

      <CodeBlock
        code={`// WebSocket Chat Server
websocket_server_chalu(3000, kaj(conn) {
    dekho("Message:", conn["message"]);
    websocket_pathao(conn, "Reply: " + conn["message"]);
});

// WebSocket Client
proyash kaj chat() {
    dhoro ws = opekha websocket_jukto("ws://localhost:3000");
    websocket_pathao(ws, "Hello WebSocket!");
    ghumaao(1000);
    websocket_bondho(ws);
}
chat();`}
      />

      <h2>Function Name Meanings</h2>

      <p>All built-in functions use Bengali words:</p>

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
            <tr><td><code>dekho</code></td><td>দেখো</td><td>look/see (print)</td></tr>
            <tr><td><code>nao</code></td><td>নাও</td><td>take (input)</td></tr>
            <tr><td><code>dhoron</code></td><td>ধরন</td><td>type/kind</td></tr>
            <tr><td><code>lipi</code></td><td>লিপি</td><td>script/writing (string)</td></tr>
            <tr><td><code>sonkha</code></td><td>সংখ্যা</td><td>number</td></tr>
            <tr><td><code>dorghyo</code></td><td>দৈর্ঘ্য</td><td>length</td></tr>
            <tr><td><code>dhokao</code></td><td>ঢোকাও</td><td>insert/push</td></tr>
            <tr><td><code>berKoro</code></td><td>বের করো</td><td>take out/pop</td></tr>
            <tr><td><code>boroHater</code></td><td>বড় হাতের</td><td>uppercase</td></tr>
            <tr><td><code>chotoHater</code></td><td>ছোট হাতের</td><td>lowercase</td></tr>
            <tr><td><code>borgomul</code></td><td>বর্গমূল</td><td>square root</td></tr>
            <tr><td><code>somoy</code></td><td>সময়</td><td>time</td></tr>
            <tr><td><code>poro</code></td><td>পড়ো</td><td>read</td></tr>
            <tr><td><code>lekho</code></td><td>লেখো</td><td>write</td></tr>
          </tbody>
        </table>
      </div>

      <h2>System Operations</h2>

      <p className="text-muted-foreground">
        BanglaCode provides 53 system-level functions for file operations, process management,
        network information, system statistics, environment variables, and temporary file handling.
      </p>

      <h3>File Metadata</h3>

      <div className="overflow-x-auto my-4">
        <table>
          <thead>
            <tr>
              <th>Function</th>
              <th>Parameters</th>
              <th>Returns</th>
              <th>Description</th>
            </tr>
          </thead>
          <tbody>
            <tr>
              <td><code>file_akar</code></td>
              <td><code>path</code></td>
              <td><code>number</code></td>
              <td>Get file size in bytes</td>
            </tr>
            <tr>
              <td><code>file_permission</code></td>
              <td><code>path</code></td>
              <td><code>string</code></td>
              <td>Get file permissions (e.g., "0644")</td>
            </tr>
            <tr>
              <td><code>file_permission_set</code></td>
              <td><code>path, perms</code></td>
              <td><code>khali</code></td>
              <td>Change file permissions</td>
            </tr>
            <tr>
              <td><code>file_malikan</code></td>
              <td><code>path</code></td>
              <td><code>map</code></td>
              <td>Get file owner (uid, gid, naam)</td>
            </tr>
            <tr>
              <td><code>file_shomoy_poribortito</code></td>
              <td><code>path</code></td>
              <td><code>number</code></td>
              <td>Get file modified time (Unix)</td>
            </tr>
            <tr>
              <td><code>file_dhoron</code></td>
              <td><code>path</code></td>
              <td><code>string</code></td>
              <td>Get file type (file/directory/symlink)</td>
            </tr>
            <tr>
              <td><code>file_rename</code></td>
              <td><code>old, new</code></td>
              <td><code>khali</code></td>
              <td>Rename or move file</td>
            </tr>
          </tbody>
        </table>
      </div>

      <h3>Directory Operations</h3>

      <div className="overflow-x-auto my-4">
        <table>
          <thead>
            <tr>
              <th>Function</th>
              <th>Parameters</th>
              <th>Returns</th>
              <th>Description</th>
            </tr>
          </thead>
          <tbody>
            <tr>
              <td><code>directory_taliika</code></td>
              <td><code>path</code></td>
              <td><code>array</code></td>
              <td>List directory contents</td>
            </tr>
            <tr>
              <td><code>directory_ghumao</code></td>
              <td><code>path</code></td>
              <td><code>array</code></td>
              <td>Recursive directory traversal</td>
            </tr>
            <tr>
              <td><code>directory_khali_ki</code></td>
              <td><code>path</code></td>
              <td><code>boolean</code></td>
              <td>Check if directory is empty</td>
            </tr>
            <tr>
              <td><code>directory_akar</code></td>
              <td><code>path</code></td>
              <td><code>number</code></td>
              <td>Get total directory size in bytes</td>
            </tr>
          </tbody>
        </table>
      </div>

      <CodeBlock
        code={`// File metadata example
dhoro size = file_akar("/path/to/file.txt");
dekho("Size:", size, "bytes");

dhoro perms = file_permission("/path/to/file.txt");
dekho("Permissions:", perms);

// Directory operations
dhoro files = directory_taliika("/home/user");
ghuriye (dhoro i = 0; i < dorghyo(files); i = i + 1) {
  dekho(files[i]);
}`}
      />

      <h3>Process Management</h3>

      <div className="overflow-x-auto my-4">
        <table>
          <thead>
            <tr>
              <th>Function</th>
              <th>Parameters</th>
              <th>Returns</th>
              <th>Description</th>
            </tr>
          </thead>
          <tbody>
            <tr>
              <td><code>chalan</code></td>
              <td><code>cmd, [args]</code></td>
              <td><code>map</code></td>
              <td>Execute system command</td>
            </tr>
            <tr>
              <td><code>process_ghum</code></td>
              <td><code>ms</code></td>
              <td><code>khali</code></td>
              <td>Sleep for milliseconds</td>
            </tr>
            <tr>
              <td><code>process_maro</code></td>
              <td><code>pid</code></td>
              <td><code>khali</code></td>
              <td>Kill process by PID</td>
            </tr>
            <tr>
              <td><code>process_chalu</code></td>
              <td><code>cmd, [args]</code></td>
              <td><code>map</code></td>
              <td>Start process in background</td>
            </tr>
            <tr>
              <td><code>process_opekha</code></td>
              <td><code>pid</code></td>
              <td><code>map</code></td>
              <td>Wait for process completion</td>
            </tr>
          </tbody>
        </table>
      </div>

      <h3>Network Information</h3>

      <div className="overflow-x-auto my-4">
        <table>
          <thead>
            <tr>
              <th>Function</th>
              <th>Parameters</th>
              <th>Returns</th>
              <th>Description</th>
            </tr>
          </thead>
          <tbody>
            <tr>
              <td><code>network_interface</code></td>
              <td>-</td>
              <td><code>array</code></td>
              <td>Get all network interfaces</td>
            </tr>
            <tr>
              <td><code>ip_address</code></td>
              <td><code>iface</code></td>
              <td><code>array</code></td>
              <td>Get IP addresses for interface</td>
            </tr>
            <tr>
              <td><code>ip_shokal</code></td>
              <td>-</td>
              <td><code>array</code></td>
              <td>Get all IP addresses</td>
            </tr>
            <tr>
              <td><code>mac_address</code></td>
              <td><code>iface</code></td>
              <td><code>string</code></td>
              <td>Get MAC address for interface</td>
            </tr>
          </tbody>
        </table>
      </div>

      <h3>System Statistics</h3>

      <div className="overflow-x-auto my-4">
        <table>
          <thead>
            <tr>
              <th>Function</th>
              <th>Parameters</th>
              <th>Returns</th>
              <th>Description</th>
            </tr>
          </thead>
          <tbody>
            <tr>
              <td><code>memory_total</code></td>
              <td>-</td>
              <td><code>number</code></td>
              <td>Total system memory in bytes</td>
            </tr>
            <tr>
              <td><code>memory_bebohrito</code></td>
              <td>-</td>
              <td><code>number</code></td>
              <td>Used memory in bytes</td>
            </tr>
            <tr>
              <td><code>disk_akar</code></td>
              <td><code>[path]</code></td>
              <td><code>number</code></td>
              <td>Total disk size in bytes</td>
            </tr>
            <tr>
              <td><code>disk_mukt</code></td>
              <td><code>[path]</code></td>
              <td><code>number</code></td>
              <td>Free disk space in bytes</td>
            </tr>
          </tbody>
        </table>
      </div>

      <h3>Time Operations</h3>

      <div className="overflow-x-auto my-4">
        <table>
          <thead>
            <tr>
              <th>Function</th>
              <th>Parameters</th>
              <th>Returns</th>
              <th>Description</th>
            </tr>
          </thead>
          <tbody>
            <tr>
              <td><code>shomoy_ekhon</code></td>
              <td>-</td>
              <td><code>number</code></td>
              <td>Current Unix timestamp</td>
            </tr>
            <tr>
              <td><code>shomoy_format</code></td>
              <td><code>timestamp, [format]</code></td>
              <td><code>string</code></td>
              <td>Format timestamp to string</td>
            </tr>
            <tr>
              <td><code>shomoy_parse</code></td>
              <td><code>str, [format]</code></td>
              <td><code>number</code></td>
              <td>Parse time string to timestamp</td>
            </tr>
            <tr>
              <td><code>timezone</code></td>
              <td>-</td>
              <td><code>string</code></td>
              <td>Get current timezone</td>
            </tr>
          </tbody>
        </table>
      </div>

      <h3>Temporary Files</h3>

      <div className="overflow-x-auto my-4">
        <table>
          <thead>
            <tr>
              <th>Function</th>
              <th>Parameters</th>
              <th>Returns</th>
              <th>Description</th>
            </tr>
          </thead>
          <tbody>
            <tr>
              <td><code>temp_directory</code></td>
              <td>-</td>
              <td><code>string</code></td>
              <td>Get system temp directory</td>
            </tr>
            <tr>
              <td><code>temp_file</code></td>
              <td><code>[prefix]</code></td>
              <td><code>string</code></td>
              <td>Create temporary file</td>
            </tr>
            <tr>
              <td><code>temp_folder</code></td>
              <td><code>[prefix]</code></td>
              <td><code>string</code></td>
              <td>Create temporary directory</td>
            </tr>
          </tbody>
        </table>
      </div>

      <h3>Symbolic Links</h3>

      <div className="overflow-x-auto my-4">
        <table>
          <thead>
            <tr>
              <th>Function</th>
              <th>Parameters</th>
              <th>Returns</th>
              <th>Description</th>
            </tr>
          </thead>
          <tbody>
            <tr>
              <td><code>symlink_banao</code></td>
              <td><code>target, link</code></td>
              <td><code>khali</code></td>
              <td>Create symbolic link</td>
            </tr>
            <tr>
              <td><code>symlink_poro</code></td>
              <td><code>link</code></td>
              <td><code>string</code></td>
              <td>Read symlink target</td>
            </tr>
            <tr>
              <td><code>symlink_ki</code></td>
              <td><code>path</code></td>
              <td><code>boolean</code></td>
              <td>Check if path is symlink</td>
            </tr>
          </tbody>
        </table>
      </div>

      <h2>Database Functions</h2>

      <p>
        Production-grade database connectors for PostgreSQL, MySQL, MongoDB, and Redis with connection pooling and async support.
        See the <a href="/docs/database" className="text-primary hover:underline">Database Documentation</a> for detailed examples.
      </p>

      <h3>Universal Database Functions</h3>

      <div className="overflow-x-auto my-4">
        <table>
          <thead>
            <tr>
              <th>Function</th>
              <th>Parameters</th>
              <th>Returns</th>
              <th>Description</th>
            </tr>
          </thead>
          <tbody>
            <tr>
              <td><code>db_jukto</code></td>
              <td><code>type, config</code></td>
              <td><code>connection</code></td>
              <td>Connect to database (postgres/mysql/mongodb/redis)</td>
            </tr>
            <tr>
              <td><code>db_jukto_async</code></td>
              <td><code>type, config</code></td>
              <td><code>promise</code></td>
              <td>Connect to database (async)</td>
            </tr>
            <tr>
              <td><code>db_bandho</code></td>
              <td><code>conn</code></td>
              <td><code>khali</code></td>
              <td>Close database connection</td>
            </tr>
            <tr>
              <td><code>db_query</code></td>
              <td><code>conn, sql</code></td>
              <td><code>result</code></td>
              <td>Execute SELECT query (SQL databases)</td>
            </tr>
            <tr>
              <td><code>db_exec</code></td>
              <td><code>conn, sql</code></td>
              <td><code>result</code></td>
              <td>Execute INSERT/UPDATE/DELETE</td>
            </tr>
            <tr>
              <td><code>db_proshno</code></td>
              <td><code>conn, sql, params</code></td>
              <td><code>result</code></td>
              <td>Prepared statement (SQL injection safe)</td>
            </tr>
          </tbody>
        </table>
      </div>

      <h3>Connection Pool Functions</h3>

      <div className="overflow-x-auto my-4">
        <table>
          <thead>
            <tr>
              <th>Function</th>
              <th>Parameters</th>
              <th>Returns</th>
              <th>Description</th>
            </tr>
          </thead>
          <tbody>
            <tr>
              <td><code>db_pool_banao</code></td>
              <td><code>type, config, maxConns</code></td>
              <td><code>pool</code></td>
              <td>Create connection pool (50-100x faster)</td>
            </tr>
            <tr>
              <td><code>db_pool_nao</code></td>
              <td><code>pool</code></td>
              <td><code>connection</code></td>
              <td>Get connection from pool</td>
            </tr>
            <tr>
              <td><code>db_pool_ferot</code></td>
              <td><code>pool, conn</code></td>
              <td><code>khali</code></td>
              <td>Return connection to pool</td>
            </tr>
            <tr>
              <td><code>db_pool_bondho</code></td>
              <td><code>pool</code></td>
              <td><code>khali</code></td>
              <td>Close connection pool</td>
            </tr>
            <tr>
              <td><code>db_pool_tothyo</code></td>
              <td><code>pool</code></td>
              <td><code>map</code></td>
              <td>Get pool statistics</td>
            </tr>
          </tbody>
        </table>
      </div>

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
              <td>Connect to PostgreSQL</td>
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
              <td>Prepared statement</td>
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

      <h3>MySQL Functions</h3>

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
              <td><code>db_jukto_mysql</code></td>
              <td><code>config</code></td>
              <td>Connect to MySQL</td>
            </tr>
            <tr>
              <td><code>db_query_mysql</code></td>
              <td><code>conn, sql</code></td>
              <td>Execute SELECT query</td>
            </tr>
            <tr>
              <td><code>db_exec_mysql</code></td>
              <td><code>conn, sql</code></td>
              <td>Execute INSERT/UPDATE/DELETE</td>
            </tr>
            <tr>
              <td><code>db_proshno_mysql</code></td>
              <td><code>conn, sql, params</code></td>
              <td>Prepared statement</td>
            </tr>
            <tr>
              <td><code>db_transaction_shuru_mysql</code></td>
              <td><code>conn</code></td>
              <td>Begin transaction</td>
            </tr>
            <tr>
              <td><code>db_commit_mysql</code></td>
              <td><code>tx</code></td>
              <td>Commit transaction</td>
            </tr>
            <tr>
              <td><code>db_rollback_mysql</code></td>
              <td><code>tx</code></td>
              <td>Rollback transaction</td>
            </tr>
          </tbody>
        </table>
      </div>

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
              <td>Set key-value (optional TTL)</td>
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

      <DocNavigation currentPath="/docs/builtins" />
    </div>
  );
}
