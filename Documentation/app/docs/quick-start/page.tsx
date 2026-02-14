import Link from "next/link";
import CodeBlock from "@/components/CodeBlock";
import DocNavigation from "@/components/DocNavigation";

export default function QuickStart() {
  return (
    <div>
      <div className="flex items-center gap-2 text-sm text-muted-foreground mb-4">
        <span className="px-2 py-1 bg-primary/10 text-primary rounded-full text-xs font-medium">
          Getting Started
        </span>
      </div>

      <h1>Quick Start Guide</h1>

      <p className="lead text-xl text-muted-foreground mt-4">
        Write your first BanglaCode program in 5 minutes. This guide covers the essential
        syntax to get you started.
      </p>

      <h2>Your First Program</h2>

      <p>
        Create a file called <code>hello.bang</code> and add the following code:
      </p>

      <CodeBlock
        filename="hello.bang"
        code={`// Your first BanglaCode program!
dekho("Namaskar, BanglaCode!");

// Variables
dhoro naam = "World";
dekho("Hello,", naam);`}
      />

      <p>Run it with:</p>

      <CodeBlock
        language="bash"
        showLineNumbers={false}
        code={`./banglacode hello.bang`}
      />

      <p>Output:</p>

      <CodeBlock
        language="output"
        showLineNumbers={false}
        code={`Namaskar, BanglaCode!
Hello, World`}
      />

      <h2>Variables</h2>

      <p>
        Use <code>dhoro</code> (meaning &quot;hold&quot;) to declare variables. BanglaCode is dynamically typed.
      </p>

      <CodeBlock
        code={`// Numbers
dhoro age = 25;
dhoro pi = 3.14159;

// Strings
dhoro name = "Rahim";
dhoro greeting = 'Namaskar';

// Booleans
dhoro isStudent = sotti;    // true
dhoro isWorking = mittha;   // false

// Null
dhoro data = khali;         // null

// Arrays
dhoro numbers = [1, 2, 3, 4, 5];
dhoro mixed = ["hello", 42, sotti];

// Maps (objects)
dhoro person = {
    naam: "Ankan",
    boyosh: 25,
    city: "Kolkata"
};`}
      />

      <h2>Basic Operations</h2>

      <CodeBlock
        code={`// Arithmetic
dhoro sum = 10 + 5;      // 15
dhoro diff = 10 - 5;     // 5
dhoro prod = 10 * 5;     // 50
dhoro quot = 10 / 5;     // 2
dhoro rem = 10 % 3;      // 1
dhoro power = 2 ** 3;    // 8

// String concatenation
dhoro fullName = "Rahim" + " " + "Khan";

// Comparison
dhoro isEqual = 5 == 5;      // sotti
dhoro isGreater = 10 > 5;    // sotti

// Logical operators
dhoro result = sotti ebong mittha;   // and -> mittha
dhoro result2 = sotti ba mittha;     // or -> sotti
dhoro result3 = na sotti;            // not -> mittha`}
      />

      <h2>Conditionals</h2>

      <p>
        Use <code>jodi</code> (if) and <code>nahole</code> (else) for conditional statements:
      </p>

      <CodeBlock
        code={`dhoro score = 85;

jodi (score >= 90) {
    dekho("Grade: A");
} nahole jodi (score >= 80) {
    dekho("Grade: B");
} nahole jodi (score >= 70) {
    dekho("Grade: C");
} nahole {
    dekho("Grade: F");
}`}
      />

      <h2>Loops</h2>

      <h3>While Loop</h3>

      <p>Use <code>jotokkhon</code> (while) for while loops:</p>

      <CodeBlock
        code={`dhoro i = 0;
jotokkhon (i < 5) {
    dekho("Count:", i);
    i = i + 1;
}`}
      />

      <h3>For Loop</h3>

      <p>Use <code>ghuriye</code> (rotating/looping) for for loops:</p>

      <CodeBlock
        code={`ghuriye (dhoro i = 1; i <= 5; i = i + 1) {
    dekho("Number:", i);
}`}
      />

      <h2>Functions</h2>

      <p>
        Use <code>kaj</code> (work/function) to define functions and <code>ferao</code> (return) to return values:
      </p>

      <CodeBlock
        code={`// Simple function
kaj greet(naam) {
    dekho("Namaskar,", naam);
}

greet("Rahim");  // Namaskar, Rahim

// Function with return value
kaj add(a, b) {
    ferao a + b;
}

dhoro result = add(5, 3);
dekho("Sum:", result);  // Sum: 8

// Recursive function
kaj factorial(n) {
    jodi (n <= 1) {
        ferao 1;
    }
    ferao n * factorial(n - 1);
}

dekho("5! =", factorial(5));  // 5! = 120`}
      />

      <h2>Classes</h2>

      <p>
        Use <code>sreni</code> (class) to define classes and <code>notun</code> (new) to create instances:
      </p>

      <CodeBlock
        code={`sreni Person {
    // Constructor
    shuru(naam, boyosh) {
        ei.naam = naam;
        ei.boyosh = boyosh;
    }

    // Method
    kaj introduce() {
        dekho("Hi, I am", ei.naam, "and I am", ei.boyosh, "years old");
    }

    kaj birthday() {
        ei.boyosh = ei.boyosh + 1;
        dekho(ei.naam, "is now", ei.boyosh);
    }
}

// Create instance
dhoro person = notun Person("Ankan", 25);
person.introduce();  // Hi, I am Ankan and I am 25 years old
person.birthday();   // Ankan is now 26`}
      />

      <h2>Error Handling</h2>

      <p>
        Use <code>chesta</code> (try), <code>dhoro_bhul</code> (catch error), and <code>felo</code> (throw):
      </p>

      <CodeBlock
        code={`chesta {
    dhoro result = 10 / 0;
} dhoro_bhul (error) {
    dekho("Error occurred:", error);
} shesh {
    dekho("Cleanup complete");
}

// Throw custom error
kaj validateAge(age) {
    jodi (age < 0) {
        felo "Age cannot be negative";
    }
    ferao sotti;
}`}
      />

      <h2>Built-in Functions</h2>

      <p>BanglaCode comes with 50+ built-in functions:</p>

      <CodeBlock
        code={`// Output
dekho("Hello World");

// Type checking
dekho(dhoron(42));       // "int"
dekho(dhoron("hello"));  // "string"

// String operations
dekho(dorghyo("hello"));           // 5
dekho(boroHater("hello"));         // "HELLO"
dekho(chotoHater("HELLO"));        // "hello"

// Array operations
dhoro arr = [3, 1, 4, 1, 5];
dhokao(arr, 9);                    // Push
dekho(dorghyo(arr));               // 6
dekho(saja(arr));                  // [1, 1, 3, 4, 5, 9]

// Math
dekho(borgomul(16));     // 4 (square root)
dekho(niratek(-5));      // 5 (absolute value)
dekho(kache(3.7));       // 4 (round)`}
      />

      <h2>Next Steps</h2>

      <p>
        You now know the basics! Explore the documentation to learn more about:
      </p>

      <ul>
        <li><Link href="/docs/keywords" className="text-primary hover:underline">All 27 Bengali keywords</Link></li>
        <li><Link href="/docs/builtins" className="text-primary hover:underline">50+ built-in functions</Link></li>
        <li><Link href="/docs/modules" className="text-primary hover:underline">Module system (import/export)</Link></li>
        <li><Link href="/docs/http-server" className="text-primary hover:underline">Building HTTP servers</Link></li>
      </ul>

      <DocNavigation currentPath="/docs/quick-start" />
    </div>
  );
}
