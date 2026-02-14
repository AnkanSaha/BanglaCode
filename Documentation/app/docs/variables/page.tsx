import CodeBlock from "@/components/CodeBlock";
import DocNavigation from "@/components/DocNavigation";

export default function Variables() {
  return (
    <div>
      <div className="flex items-center gap-2 text-sm text-muted-foreground mb-4">
        <span className="px-2 py-1 bg-primary/10 text-primary rounded-full text-xs font-medium">
          Language Basics
        </span>
      </div>

      <h1>Variables</h1>

      <p className="lead text-xl text-muted-foreground mt-4">
        BanglaCode provides three ways to declare variables: <code>dhoro</code> (regular),{" "}
        <code>sthir</code> (constant), and <code>bishwo</code> (global). BanglaCode is dynamically
        typed, so you don&apos;t need to specify types.
      </p>

      <h2>Variable Declaration Keywords</h2>

      <div className="overflow-x-auto my-6">
        <table className="w-full border-collapse">
          <thead>
            <tr className="border-b border-border">
              <th className="text-left py-3 px-4 font-semibold">Keyword</th>
              <th className="text-left py-3 px-4 font-semibold">Bengali</th>
              <th className="text-left py-3 px-4 font-semibold">English</th>
              <th className="text-left py-3 px-4 font-semibold">Description</th>
            </tr>
          </thead>
          <tbody>
            <tr className="border-b border-border/50">
              <td className="py-3 px-4"><code>dhoro</code></td>
              <td className="py-3 px-4">ধরো</td>
              <td className="py-3 px-4">let</td>
              <td className="py-3 px-4">Regular mutable variable</td>
            </tr>
            <tr className="border-b border-border/50">
              <td className="py-3 px-4"><code>sthir</code></td>
              <td className="py-3 px-4">স্থির</td>
              <td className="py-3 px-4">const</td>
              <td className="py-3 px-4">Immutable constant (cannot be changed)</td>
            </tr>
            <tr className="border-b border-border/50">
              <td className="py-3 px-4"><code>bishwo</code></td>
              <td className="py-3 px-4">বিশ্ব</td>
              <td className="py-3 px-4">global</td>
              <td className="py-3 px-4">Global variable (accessible everywhere)</td>
            </tr>
          </tbody>
        </table>
      </div>

      <h2>Regular Variables (dhoro)</h2>

      <p>
        Use the <code>dhoro</code> keyword for regular mutable variables:
      </p>

      <CodeBlock
        code={`// Basic variable declaration
dhoro naam = "Rahim";
dhoro boyosh = 25;
dhoro active = sotti;

// Variables can be reassigned
dhoro message = "Hello";
message = 42;        // Now holds a number
message = [1, 2, 3]; // Now holds an array`}
      />

      <h2>Constants (sthir)</h2>

      <p>
        Use the <code>sthir</code> (স্থির = fixed/constant) keyword for values that should never change.
        Attempting to reassign a constant will result in an error:
      </p>

      <CodeBlock
        code={`// Declare constants
sthir PI = 3.14159;
sthir MAX_SIZE = 100;
sthir APP_NAME = "BanglaCode";

// Use constants
dhoro area = PI * 5 * 5;
dekho(area);  // 78.53975

// This will cause an error!
// PI = 3.14;  // Error: 'PI' ekti sthir (constant), eitake bodlano jabe na`}
      />

      <div className="bg-amber-500/10 border border-amber-500/20 rounded-lg p-4 my-4">
        <p className="text-amber-200 font-medium mb-2">⚠️ Constant Naming Convention</p>
        <p className="text-muted-foreground text-sm">
          By convention, use UPPERCASE names for constants to make them easily identifiable.
        </p>
      </div>

      <h2>Global Variables (bishwo)</h2>

      <p>
        Use the <code>bishwo</code> (বিশ্ব = world/global) keyword for variables that need to be
        accessible from any scope, including inside functions:
      </p>

      <CodeBlock
        code={`// Declare a global variable
bishwo ganok = 0;

kaj barao() {
    // Can access and modify global variable
    ganok = ganok + 1;
}

barao();
barao();
barao();

dekho(ganok);  // 3`}
      />

      <h3>Global vs Regular Variables</h3>

      <p>
        The key difference is that regular variables inside functions create new local bindings,
        while global variables are shared across all scopes:
      </p>

      <CodeBlock
        code={`// Regular variable - function creates new local binding
dhoro x = 10;

kaj test1() {
    dhoro x = 20;  // New local variable
    dekho(x);      // 20
}

test1();
dekho(x);  // 10 (outer x unchanged)

// Global variable - shared across all scopes
bishwo counter = 0;

kaj test2() {
    counter = counter + 1;  // Modifies global
}

test2();
test2();
dekho(counter);  // 2`}
      />

      <h2>Naming Rules</h2>

      <p>Variable names must follow these rules:</p>

      <ul>
        <li>Must start with a letter (a-z, A-Z) or underscore (_)</li>
        <li>Can contain letters, digits (0-9), and underscores</li>
        <li>Cannot be a reserved keyword</li>
        <li>Are case-sensitive (<code>naam</code> and <code>Naam</code> are different)</li>
      </ul>

      <CodeBlock
        code={`// Valid variable names
dhoro firstName = "Rahim";
dhoro _private = "secret";
dhoro count2 = 100;
sthir MAX_VALUE = 1000;
bishwo app_state = "running";

// Invalid variable names (will cause errors)
// dhoro 2count = 5;      // Cannot start with digit
// dhoro first-name = ""; // Hyphens not allowed
// dhoro dhoro = 5;       // Cannot use reserved keywords`}
      />

      <h2>Assignment Operators</h2>

      <p>
        BanglaCode supports compound assignment operators for regular variables:
      </p>

      <CodeBlock
        code={`dhoro x = 10;

// Simple assignment
x = 20;

// Compound assignment operators
x += 5;   // x = x + 5  (x is now 25)
x -= 3;   // x = x - 3  (x is now 22)
x *= 2;   // x = x * 2  (x is now 44)
x /= 4;   // x = x / 4  (x is now 11)

dekho(x);  // Output: 11

// Note: Compound operators don't work on constants
// sthir Y = 10;
// Y += 5;  // Error!`}
      />

      <h2>Variable Scope</h2>

      <p>
        Variables have function scope. A variable declared inside a function is only accessible
        within that function:
      </p>

      <CodeBlock
        code={`dhoro outer = "I am outer";

kaj myFunction() {
    dhoro inner = "I am inner";
    dekho(outer);  // Works - can access outer variable
    dekho(inner);  // Works - can access local variable
}

myFunction();
dekho(outer);  // Works
// dekho(inner);  // Error! inner is not defined here`}
      />

      <h3>Closures</h3>

      <p>
        Functions capture variables from their surrounding scope:
      </p>

      <CodeBlock
        code={`kaj makeCounter() {
    dhoro count = 0;

    ferao kaj() {
        count = count + 1;
        ferao count;
    };
}

dhoro counter = makeCounter();
dekho(counter());  // 1
dekho(counter());  // 2
dekho(counter());  // 3`}
      />

      <h2>Combining All Three</h2>

      <p>
        Here&apos;s an example using all three variable types together:
      </p>

      <CodeBlock
        code={`// Constants for configuration
sthir PI = 3.14159;
sthir RATE = 0.05;

// Global for shared state
bishwo total = 0;

kaj calculateArea(radius) {
    // Local variable for computation
    dhoro area = PI * radius * radius;
    
    // Update global total
    total = total + area;
    
    ferao area;
}

dekho(calculateArea(5));   // 78.53975
dekho(calculateArea(10));  // 314.159
dekho("Total:", total);    // Total: 392.69875`}
      />

      <h2>Type Checking</h2>

      <p>
        Use the <code>dhoron</code> (type) built-in function to check a variable&apos;s type:
      </p>

      <CodeBlock
        code={`dhoro num = 42;
sthir STR = "hello";
bishwo flag = sotti;

dekho(dhoron(num));   // "int" or "float"
dekho(dhoron(STR));   // "string"
dekho(dhoron(flag));  // "boolean"`}
      />

      <h2>Type Conversion</h2>

      <p>
        BanglaCode provides built-in functions for type conversion:
      </p>

      <CodeBlock
        code={`// Convert to string
dhoro num = 42;
dhoro str = lipi(num);  // "42"

// Convert to number
dhoro text = "123";
dhoro number = sonkha(text);  // 123

// Check results
dekho(dhoron(str));     // "string"
dekho(dhoron(number));  // "int"`}
      />

      <h2>Best Practices</h2>

      <ul>
        <li>Use <code>sthir</code> for values that should never change (configuration, mathematical constants)</li>
        <li>Use <code>bishwo</code> sparingly - only when truly needed for shared state</li>
        <li>Prefer <code>dhoro</code> for most variables</li>
        <li>Use descriptive names that explain the variable&apos;s purpose</li>
        <li>Use UPPERCASE for constants, camelCase or snake_case for others</li>
      </ul>

      <DocNavigation currentPath="/docs/variables" />
    </div>
  );
}
