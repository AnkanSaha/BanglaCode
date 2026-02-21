import CodeBlock from "@/components/CodeBlock";
import DocNavigation from "@/components/DocNavigation";

export default function DataTypes() {
  return (
    <div>
      <div className="flex items-center gap-2 text-sm text-muted-foreground mb-4">
        <span className="px-2 py-1 bg-primary/10 text-primary rounded-full text-xs font-medium">
          Language Basics
        </span>
      </div>

      <h1>Data Types</h1>

      <p className="lead text-xl text-muted-foreground mt-4">
        BanglaCode is dynamically typed and supports several built-in data types including
        numbers, strings, booleans, arrays, maps, functions, and classes.
      </p>

      <h2>Type Overview</h2>

      <div className="overflow-x-auto my-6">
        <table>
          <thead>
            <tr>
              <th>Type</th>
              <th>Description</th>
              <th>Example</th>
              <th><code>dhoron()</code> Returns</th>
            </tr>
          </thead>
          <tbody>
            <tr>
              <td><strong>Number (Integer)</strong></td>
              <td>Whole numbers</td>
              <td><code>42</code>, <code>-5</code></td>
              <td><code>&quot;int&quot;</code></td>
            </tr>
            <tr>
              <td><strong>Number (Float)</strong></td>
              <td>Decimal numbers</td>
              <td><code>3.14</code>, <code>-0.5</code></td>
              <td><code>&quot;float&quot;</code></td>
            </tr>
            <tr>
              <td><strong>String</strong></td>
              <td>Text</td>
              <td><code>&quot;hello&quot;</code></td>
              <td><code>&quot;string&quot;</code></td>
            </tr>
            <tr>
              <td><strong>Boolean</strong></td>
              <td>True/False</td>
              <td><code>sotti</code>, <code>mittha</code></td>
              <td><code>&quot;boolean&quot;</code></td>
            </tr>
            <tr>
              <td><strong>Null</strong></td>
              <td>No value</td>
              <td><code>khali</code></td>
              <td><code>&quot;null&quot;</code></td>
            </tr>
            <tr>
              <td><strong>Array</strong></td>
              <td>Ordered list</td>
              <td><code>[1, 2, 3]</code></td>
              <td><code>&quot;array&quot;</code></td>
            </tr>
            <tr>
              <td><strong>Map</strong></td>
              <td>Key-value pairs</td>
              <td><code>{`{x: 1, y: 2}`}</code></td>
              <td><code>&quot;map&quot;</code></td>
            </tr>
            <tr>
              <td><strong>Function</strong></td>
              <td>Callable</td>
              <td><code>kaj(x) {`{}`}</code></td>
              <td><code>&quot;function&quot;</code></td>
            </tr>
            <tr>
              <td><strong>Class</strong></td>
              <td>Blueprint</td>
              <td><code>sreni X {`{}`}</code></td>
              <td><code>&quot;class&quot;</code></td>
            </tr>
            <tr>
              <td><strong>Instance</strong></td>
              <td>Object instance</td>
              <td><code>notun X()</code></td>
              <td><code>&quot;instance&quot;</code></td>
            </tr>
          </tbody>
        </table>
      </div>

      <h2>Numbers</h2>

      <p>
        BanglaCode supports both integers and floating-point numbers:
      </p>

      <CodeBlock
        code={`// Integers
dhoro age = 25;
dhoro negative = -10;
dhoro zero = 0;

// Floating-point numbers
dhoro pi = 3.14159;
dhoro temperature = -5.5;
dhoro scientific = 1.5e2;  // 150

// Arithmetic operations
dhoro sum = 10 + 5;      // 15
dhoro diff = 10 - 3;     // 7
dhoro product = 4 * 3;   // 12
dhoro quotient = 15 / 4; // 3.75
dhoro remainder = 17 % 5; // 2
dhoro power = 2 ** 10;   // 1024

// Mixed integer and float
dhoro result = 5 + 3.5;  // 8.5`}
      />

      <h2>Strings</h2>

      <p>
        Strings can be defined with double or single quotes:
      </p>

      <CodeBlock
        code={`// String literals
dhoro greeting = "Namaskar";
dhoro name = 'Rahim';

// String concatenation
dhoro fullName = "Rahim" + " " + "Khan";
dekho(fullName);  // "Rahim Khan"

// String with numbers
dhoro message = "Age: " + lipi(25);

// Multi-word strings
dhoro sentence = "This is a complete sentence.";

// Empty string
dhoro empty = "";

// String length
dekho(dorghyo(greeting));  // 8

// Access character by index
dekho(greeting[0]);  // "N"`}
      />

      <h3>String Operations</h3>

      <CodeBlock
        code={`dhoro text = "Hello World";

// Length
dekho(dorghyo(text));        // 11

// Case conversion
dekho(boroHater(text));      // "HELLO WORLD"
dekho(chotoHater(text));     // "hello world"

// Substring
dekho(angsho(text, 0, 5));   // "Hello"
dekho(angsho(text, 6));      // "World"

// Find substring
dekho(khojo(text, "World")); // 6
dekho(khojo(text, "xyz"));   // -1

// Replace
dekho(bodlo(text, "World", "BanglaCode")); // "Hello BanglaCode"

// Trim whitespace
dhoro padded = "  hello  ";
dekho(chhanto(padded));      // "hello"

// Split string
dhoro csv = "a,b,c,d";
dekho(bhag(csv, ","));       // ["a", "b", "c", "d"]`}
      />

      <h3>Template Literals</h3>

      <p>
        Template literals use backticks (`) and allow embedded expressions with ${`{...}`} syntax:
      </p>

      <CodeBlock
        code={`// Basic interpolation
dhoro name = "Rahim";
dhoro greeting = \`Hello \${name}!\`;
dekho(greeting);  // "Hello Rahim!"

// Multiple expressions
dhoro age = 25;
dhoro city = "Dhaka";
dhoro intro = \`My name is \${name}, I'm \${age} years old, from \${city}\`;
dekho(intro);  // "My name is Rahim, I'm 25 years old, from Dhaka"

// Arithmetic in expressions
dhoro num1 = 10;
dhoro num2 = 20;
dekho(\`Sum: \${num1 + num2}\`);        // "Sum: 30"
dekho(\`Product: \${num1 * num2}\`);    // "Product: 200"

// Function calls in expressions
kaj square(x) {
    ferao x * x;
}
dekho(\`5 squared is \${square(5)}\`);  // "5 squared is 25"

// Array and object access
dhoro colors = ["Red", "Green", "Blue"];
dekho(\`First color: \${colors[0]}\`);  // "First color: Red"

dhoro person = {"naam": "Rahim", "age": 25};
dekho(\`Person: \${person["naam"]}\`);  // "Person: Rahim"

// Conditional expressions
dhoro x = 10;
dekho(\`Value is \${jodi (x > 5) { ferao "big"; } nahole { ferao "small"; }}\`);

// Chaining template literals
dhoro greeting1 = \`Hello\`;
dhoro greeting2 = \`\${greeting1} \${name}!\`;
dekho(greeting2);  // "Hello Rahim!"

// Multiple interpolations
dekho(\`\${1} + \${2} = \${1 + 2}\`);  // "1 + 2 = 3"`}
      />

      <h2>Booleans</h2>

      <p>
        Boolean values use Bengali keywords <code>sotti</code> (true) and <code>mittha</code> (false):
      </p>

      <CodeBlock
        code={`// Boolean literals
dhoro isActive = sotti;    // true
dhoro isComplete = mittha; // false

// Boolean from comparison
dhoro isEqual = 5 == 5;      // sotti
dhoro isGreater = 10 > 5;    // sotti
dhoro isLess = 3 < 1;        // mittha

// Logical operations
dhoro andResult = sotti ebong mittha;  // mittha (AND)
dhoro orResult = sotti ba mittha;      // sotti (OR)
dhoro notResult = na sotti;            // mittha (NOT)

// Using ! for NOT
dhoro negated = !mittha;  // sotti`}
      />

      <h2>Null</h2>

      <p>
        The <code>khali</code> keyword represents the absence of a value:
      </p>

      <CodeBlock
        code={`// Null value
dhoro empty = khali;

// Check for null
jodi (empty == khali) {
    dekho("Value is null");
}

// Functions without return value return null
kaj doSomething() {
    dekho("Doing something");
    // No return statement
}

dhoro result = doSomething();
dekho(result == khali);  // sotti`}
      />

      <h2>Arrays</h2>

      <p>
        Arrays are ordered collections that can hold mixed types:
      </p>

      <CodeBlock
        code={`// Array literals
dhoro numbers = [1, 2, 3, 4, 5];
dhoro mixed = ["hello", 42, sotti, khali];
dhoro nested = [[1, 2], [3, 4], [5, 6]];
dhoro empty = [];

// Access by index (0-based)
dekho(numbers[0]);   // 1
dekho(numbers[4]);   // 5
dekho(nested[1][0]); // 3

// Modify elements
numbers[0] = 100;
dekho(numbers);  // [100, 2, 3, 4, 5]

// Array length
dekho(dorghyo(numbers));  // 5

// Add element (push)
dhokao(numbers, 6);
dekho(numbers);  // [100, 2, 3, 4, 5, 6]

// Remove last element (pop)
dhoro last = berKoro(numbers);
dekho(last);     // 6
dekho(numbers);  // [100, 2, 3, 4, 5]`}
      />

      <h3>Array Operations</h3>

      <CodeBlock
        code={`dhoro arr = [3, 1, 4, 1, 5, 9, 2, 6];

// Slice
dekho(kato(arr, 2, 5));    // [4, 1, 5]
dekho(kato(arr, 5));       // [9, 2, 6]

// Reverse
dekho(ulto(arr));          // [6, 2, 9, 5, 1, 4, 1, 3]

// Check if element exists
dekho(ache(arr, 4));       // sotti
dekho(ache(arr, 10));      // mittha

// Sort
dekho(saja([3, 1, 4, 1, 5])); // [1, 1, 3, 4, 5]

// Join to string
dhoro words = ["hello", "world"];
dekho(joro(words, " "));   // "hello world"`}
      />

      <h2>Maps (Objects)</h2>

      <p>
        Maps are collections of key-value pairs:
      </p>

      <CodeBlock
        code={`// Map literal
dhoro person = {
    naam: "Rahim",
    boyosh: 25,
    city: "Dhaka",
    active: sotti
};

// Access properties
dekho(person.naam);      // "Rahim"
dekho(person["boyosh"]); // 25

// Modify properties
person.boyosh = 26;
person["city"] = "Kolkata";

// Add new properties
person.email = "rahim@example.com";

// Nested maps
dhoro company = {
    name: "TechCorp",
    address: {
        street: "123 Main St",
        city: "Dhaka"
    }
};

dekho(company.address.city);  // "Dhaka"

// Get all keys
dekho(chabi(person));  // ["naam", "boyosh", "city", "active", "email"]`}
      />

      <h2>Functions</h2>

      <p>
        Functions are first-class values in BanglaCode:
      </p>

      <CodeBlock
        code={`// Named function
kaj add(a, b) {
    ferao a + b;
}

// Function expression (anonymous)
dhoro multiply = kaj(x, y) {
    ferao x * y;
};

// Functions are values
dhoro operation = add;
dekho(operation(5, 3));  // 8

// Higher-order functions
kaj applyTwice(fn, x) {
    ferao fn(fn(x));
}

kaj double(n) {
    ferao n * 2;
}

dekho(applyTwice(double, 5));  // 20

// Check type
dekho(dhoron(add));  // "function"`}
      />

      <h2>Classes and Instances</h2>

      <p>
        Classes define blueprints for objects:
      </p>

      <CodeBlock
        code={`// Class definition
sreni Car {
    shuru(brand, model) {
        ei.brand = brand;
        ei.model = model;
    }

    kaj getInfo() {
        ferao ei.brand + " " + ei.model;
    }
}

// Create instance
dhoro myCar = notun Car("Toyota", "Corolla");

// Check types
dekho(dhoron(Car));    // "class"
dekho(dhoron(myCar));  // "instance"

// Access properties and methods
dekho(myCar.brand);      // "Toyota"
dekho(myCar.getInfo());  // "Toyota Corolla"`}
      />

      <h2>Type Conversion</h2>

      <CodeBlock
        code={`// To string
dekho(lipi(42));        // "42"
dekho(lipi(3.14));      // "3.14"
dekho(lipi(sotti));     // "true"
dekho(lipi([1, 2, 3])); // "[1, 2, 3]"

// To number
dekho(sonkha("42"));    // 42
dekho(sonkha("3.14"));  // 3.14
dekho(sonkha("abc"));   // Error or NaN

// Type checking
dekho(dhoron(42));      // "int"
dekho(dhoron(3.14));    // "float"
dekho(dhoron("hi"));    // "string"`}
      />

      <DocNavigation currentPath="/docs/data-types" />
    </div>
  );
}
