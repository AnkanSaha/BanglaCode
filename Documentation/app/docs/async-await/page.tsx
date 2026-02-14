import CodeBlock from "@/components/CodeBlock";
import DocNavigation from "@/components/DocNavigation";

export default function AsyncAwait() {
  return (
    <div>
      <div className="flex items-center gap-2 text-sm text-muted-foreground mb-4">
        <span className="px-2 py-1 bg-primary/10 text-primary rounded-full text-xs font-medium">
          Advanced
        </span>
      </div>

      <h1>Async/Await</h1>

      <p className="lead text-xl text-muted-foreground mt-4">
        BanglaCode supports asynchronous programming using promises with the <code>proyash</code> (async) and <code>opekha</code> (await) keywords. This allows you to write non-blocking code that runs concurrently.
      </p>

      <div className="bg-blue-50 dark:bg-blue-950 border border-blue-200 dark:border-blue-800 rounded-lg p-4 my-6">
        <h4 className="text-blue-900 dark:text-blue-100 font-semibold mb-2">Key Concepts</h4>
        <ul className="text-blue-800 dark:text-blue-200 space-y-1">
          <li><code>proyash</code> (প্রয়াস) - Marks a function as asynchronous</li>
          <li><code>opekha</code> (অপেক্ষা) - Waits for a promise to resolve</li>
          <li><strong>Promise</strong> - A value that will be available in the future</li>
          <li><strong>Concurrent Execution</strong> - Multiple async operations run at the same time</li>
        </ul>
      </div>

      <h2>Async Functions</h2>

      <p>
        Declare an async function using the <code>proyash</code> keyword before <code>kaj</code>:
      </p>

      <CodeBlock
        code={`// Async function definition
proyash kaj fetchData() {
    dekho("Fetching data...");
    opekha ghumaao(1000);  // Sleep for 1 second
    dekho("Data fetched!");
    ferao "Data from server";
}

// Call async function - returns a promise
dhoro promise = fetchData();
dekho(promise);  // Promise(PENDING)`}
      />

      <h2>Await Keyword</h2>

      <p>
        Use <code>opekha</code> to wait for a promise to resolve or reject. Await can only be used inside async functions:
      </p>

      <CodeBlock
        code={`proyash kaj main() {
    dekho("Starting...");

    // Wait for async operation to complete
    dhoro result = opekha fetchData();
    dekho("Result:", result);  // Result: Data from server

    dekho("Finished!");
}

// Run the main async function
main();`}
      />

      <h3>Error Handling with Await</h3>

      <CodeBlock
        code={`proyash kaj riskyOperation() {
    opekha ghumaao(500);
    felo "Something went wrong!";
}

proyash kaj safeRun() {
    chesta {
        dhoro result = opekha riskyOperation();
        dekho("Success:", result);
    } dhoro_bhul (error) {
        dekho("Error caught:", error);
    }
}

safeRun();  // Error caught: Something went wrong!`}
      />

      <h2>Built-in Async Functions</h2>

      <h3>ghumaao (ঘুমাও) - Sleep</h3>

      <p>
        Pause execution for a specified number of milliseconds:
      </p>

      <CodeBlock
        code={`proyash kaj timer() {
    dekho("Start");
    opekha ghumaao(1000);  // Sleep for 1 second
    dekho("1 second later");
    opekha ghumaao(2000);  // Sleep for 2 seconds
    dekho("3 seconds total");
}

timer();`}
      />

      <h3>sob_proyash (সব_প্রয়াস) - Promise.all</h3>

      <p>
        Wait for multiple promises to complete <strong>concurrently</strong> (all at the same time):
      </p>

      <CodeBlock
        code={`proyash kaj fetchUser() {
    opekha ghumaao(1000);
    ferao "User data";
}

proyash kaj fetchPosts() {
    opekha ghumaao(1500);
    ferao "Posts data";
}

proyash kaj fetchComments() {
    opekha ghumaao(800);
    ferao "Comments data";
}

proyash kaj loadDashboard() {
    dekho("Loading dashboard...");

    // Run all three fetches concurrently
    dhoro results = opekha sob_proyash([
        fetchUser(),
        fetchPosts(),
        fetchComments()
    ]);

    dekho("All data loaded!");
    dekho("User:", results[0]);
    dekho("Posts:", results[1]);
    dekho("Comments:", results[2]);
}

loadDashboard();
// Takes ~1.5 seconds (longest operation)
// NOT 3.3 seconds (sum of all operations)`}
      />

      <div className="bg-green-50 dark:bg-green-950 border border-green-200 dark:border-green-800 rounded-lg p-4 my-6">
        <h4 className="text-green-900 dark:text-green-100 font-semibold mb-2">Performance Tip</h4>
        <p className="text-green-800 dark:text-green-200">
          <code>sob_proyash</code> runs all promises <strong>concurrently</strong>, not sequentially.
          This means 3 promises of 500ms each will complete in ~500ms total, not 1500ms!
        </p>
      </div>

      <h3>Async File I/O</h3>

      <CodeBlock
        code={`// poro_async - Read file asynchronously
proyash kaj readConfig() {
    dhoro content = opekha poro_async("config.json");
    dekho("Config:", content);
    ferao json_poro(content);
}

// lekho_async - Write file asynchronously
proyash kaj saveData(data) {
    dhoro json = json_banao(data);
    opekha lekho_async("output.json", json);
    dekho("File saved!");
}

// Use them
proyash kaj main() {
    dhoro config = opekha readConfig();
    opekha saveData(config);
}

main();`}
      />

      <h3>Async HTTP Requests</h3>

      <CodeBlock
        code={`// anun_async - Fetch data from URL
proyash kaj fetchAPI() {
    dekho("Fetching from API...");

    dhoro response = opekha anun_async("https://api.example.com/data");
    dekho("Status:", response["status"]);

    dhoro data = json_poro(response["body"]);
    dekho("Data:", data);

    ferao data;
}

fetchAPI();`}
      />

      <h2>Practical Examples</h2>

      <h3>Sequential vs Concurrent Execution</h3>

      <CodeBlock
        code={`// Sequential (slow) - waits for each operation
proyash kaj sequential() {
    dhoro start = somoy();

    opekha ghumaao(500);  // Wait 500ms
    opekha ghumaao(500);  // Wait 500ms more
    opekha ghumaao(500);  // Wait 500ms more

    dhoro elapsed = somoy() - start;
    dekho("Sequential time:", elapsed, "ms");  // ~1500ms
}

// Concurrent (fast) - all operations at once
proyash kaj concurrent() {
    dhoro start = somoy();

    dhoro p1 = ghumaao(500);
    dhoro p2 = ghumaao(500);
    dhoro p3 = ghumaao(500);

    opekha sob_proyash([p1, p2, p3]);

    dhoro elapsed = somoy() - start;
    dekho("Concurrent time:", elapsed, "ms");  // ~500ms
}

sequential();  // Takes ~1500ms
concurrent();  // Takes ~500ms (3x faster!)`}
      />

      <h3>Error Handling with Promise.all</h3>

      <CodeBlock
        code={`proyash kaj task1() {
    opekha ghumaao(500);
    ferao "Task 1 done";
}

proyash kaj task2() {
    opekha ghumaao(300);
    felo "Task 2 failed!";
}

proyash kaj task3() {
    opekha ghumaao(700);
    ferao "Task 3 done";
}

proyash kaj runAll() {
    chesta {
        // If any promise rejects, sob_proyash rejects immediately
        dhoro results = opekha sob_proyash([
            task1(),
            task2(),
            task3()
        ]);
        dekho("All tasks succeeded:", results);
    } dhoro_bhul (error) {
        dekho("One task failed:", error);  // Task 2 failed!
    }
}

runAll();`}
      />

      <h3>Chaining Async Operations</h3>

      <CodeBlock
        code={`proyash kaj getUserId() {
    opekha ghumaao(500);
    ferao 123;
}

proyash kaj getUserProfile(userId) {
    opekha ghumaao(500);
    ferao { naam: "Rahim", id: userId };
}

proyash kaj getUserPosts(userId) {
    opekha ghumaao(500);
    ferao ["Post 1", "Post 2", "Post 3"];
}

proyash kaj loadUserData() {
    // Step 1: Get user ID
    dhoro id = opekha getUserId();
    dekho("Got user ID:", id);

    // Step 2: Load profile and posts concurrently
    dhoro results = opekha sob_proyash([
        getUserProfile(id),
        getUserPosts(id)
    ]);

    dhoro profile = results[0];
    dhoro posts = results[1];

    dekho("Profile:", profile);
    dekho("Posts:", posts);
}

loadUserData();`}
      />

      <h2>Timeout Handling</h2>

      <p>
        All <code>opekha</code> operations have a built-in 30-second timeout to prevent infinite blocking:
      </p>

      <CodeBlock
        code={`proyash kaj neverResolves() {
    // This promise never resolves or rejects
    // After 30 seconds, await will timeout
}

proyash kaj testTimeout() {
    chesta {
        opekha neverResolves();
    } dhoro_bhul (error) {
        dekho(error);  // await timeout: promise did not resolve within 30 seconds
    }
}

testTimeout();`}
      />

      <h2>Best Practices</h2>

      <div className="space-y-4">
        <div className="border-l-4 border-blue-500 pl-4">
          <h4 className="font-semibold">✓ Use concurrent execution when possible</h4>
          <p className="text-sm text-muted-foreground">
            Run independent async operations concurrently with <code>sob_proyash</code> for better performance.
          </p>
        </div>

        <div className="border-l-4 border-blue-500 pl-4">
          <h4 className="font-semibold">✓ Always handle errors</h4>
          <p className="text-sm text-muted-foreground">
            Wrap async operations in <code>chesta/dhoro_bhul</code> blocks to catch rejections.
          </p>
        </div>

        <div className="border-l-4 border-blue-500 pl-4">
          <h4 className="font-semibold">✓ Keep async functions focused</h4>
          <p className="text-sm text-muted-foreground">
            Break complex async logic into smaller, composable async functions.
          </p>
        </div>

        <div className="border-l-4 border-red-500 pl-4">
          <h4 className="font-semibold">✗ Don&apos;t use await in loops unnecessarily</h4>
          <p className="text-sm text-muted-foreground">
            Sequential awaits in loops can be slow. Use <code>sob_proyash</code> to run iterations concurrently.
          </p>
        </div>
      </div>

      <h2>Summary</h2>

      <div className="bg-gray-50 dark:bg-gray-900 rounded-lg p-6 my-6">
        <table className="w-full">
          <thead>
            <tr className="border-b border-gray-200 dark:border-gray-700">
              <th className="text-left pb-2">Keyword</th>
              <th className="text-left pb-2">Bengali</th>
              <th className="text-left pb-2">Purpose</th>
            </tr>
          </thead>
          <tbody className="text-sm">
            <tr className="border-b border-gray-100 dark:border-gray-800">
              <td className="py-2"><code>proyash</code></td>
              <td className="py-2">প্রয়াস</td>
              <td className="py-2">Declare async function</td>
            </tr>
            <tr className="border-b border-gray-100 dark:border-gray-800">
              <td className="py-2"><code>opekha</code></td>
              <td className="py-2">অপেক্ষা</td>
              <td className="py-2">Wait for promise</td>
            </tr>
            <tr className="border-b border-gray-100 dark:border-gray-800">
              <td className="py-2"><code>ghumaao</code></td>
              <td className="py-2">ঘুমাও</td>
              <td className="py-2">Sleep for milliseconds</td>
            </tr>
            <tr className="border-b border-gray-100 dark:border-gray-800">
              <td className="py-2"><code>sob_proyash</code></td>
              <td className="py-2">সব_প্রয়াস</td>
              <td className="py-2">Wait for all promises (concurrent)</td>
            </tr>
            <tr className="border-b border-gray-100 dark:border-gray-800">
              <td className="py-2"><code>poro_async</code></td>
              <td className="py-2">পড়ো_async</td>
              <td className="py-2">Read file asynchronously</td>
            </tr>
            <tr className="border-b border-gray-100 dark:border-gray-800">
              <td className="py-2"><code>lekho_async</code></td>
              <td className="py-2">লেখো_async</td>
              <td className="py-2">Write file asynchronously</td>
            </tr>
            <tr>
              <td className="py-2"><code>anun_async</code></td>
              <td className="py-2">আনুন_async</td>
              <td className="py-2">HTTP GET asynchronously</td>
            </tr>
          </tbody>
        </table>
      </div>

      <DocNavigation currentPath="/docs/async-await" />
    </div>
  );
}
