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
        BanglaCode provides 50+ built-in functions for I/O, type conversion, string manipulation,
        array operations, math, file handling, and HTTP.
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

      <DocNavigation currentPath="/docs/builtins" />
    </div>
  );
}
