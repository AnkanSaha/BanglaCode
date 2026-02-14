import CodeBlock from "@/components/CodeBlock";
import DocNavigation from "@/components/DocNavigation";

export default function Keywords() {
  const keywords = [
    { keyword: "dhoro", meaning: "hold/let", english: "var/let", description: "Variable declaration", example: 'dhoro x = 5;' },
    { keyword: "sthir", meaning: "fixed/constant", english: "const", description: "Constant declaration (immutable)", example: 'sthir PI = 3.14;' },
    { keyword: "bishwo", meaning: "world/global", english: "global", description: "Global variable declaration", example: 'bishwo count = 0;' },
    { keyword: "jodi", meaning: "if", english: "if", description: "Conditional statement", example: 'jodi (x > 0) { }' },
    { keyword: "nahole", meaning: "else/otherwise", english: "else", description: "Else branch", example: 'nahole { }' },
    { keyword: "jotokkhon", meaning: "as long as", english: "while", description: "While loop", example: 'jotokkhon (x < 10) { }' },
    { keyword: "ghuriye", meaning: "rotating/looping", english: "for", description: "For loop", example: 'ghuriye (dhoro i = 0; i < 5; i = i + 1) { }' },
    { keyword: "kaj", meaning: "work/task", english: "function", description: "Function declaration", example: 'kaj add(a, b) { ferao a + b; }' },
    { keyword: "ferao", meaning: "return/give back", english: "return", description: "Return from function", example: 'ferao result;' },
    { keyword: "sreni", meaning: "class/category", english: "class", description: "Class declaration", example: 'sreni Person { }' },
    { keyword: "shuru", meaning: "start/begin", english: "constructor", description: "Constructor method", example: 'shuru(naam) { ei.naam = naam; }' },
    { keyword: "notun", meaning: "new", english: "new", description: "Create instance", example: 'notun Person("Ankan")' },
    { keyword: "ei", meaning: "this", english: "this", description: "Current instance reference", example: 'ei.naam = "value";' },
    { keyword: "sotti", meaning: "truth", english: "true", description: "Boolean true", example: 'dhoro active = sotti;' },
    { keyword: "mittha", meaning: "false/lie", english: "false", description: "Boolean false", example: 'dhoro done = mittha;' },
    { keyword: "khali", meaning: "empty", english: "null", description: "Null value", example: 'dhoro data = khali;' },
    { keyword: "ebong", meaning: "and", english: "and", description: "Logical AND", example: 'x > 0 ebong y < 10' },
    { keyword: "ba", meaning: "or", english: "or", description: "Logical OR", example: 'x == 1 ba x == 2' },
    { keyword: "na", meaning: "not/no", english: "not", description: "Logical NOT", example: 'na sotti' },
    { keyword: "thamo", meaning: "stop", english: "break", description: "Break from loop", example: 'thamo;' },
    { keyword: "chharo", meaning: "skip/leave", english: "continue", description: "Continue to next iteration", example: 'chharo;' },
    { keyword: "ano", meaning: "bring", english: "import", description: "Import module", example: 'ano "math.bang" hisabe math;' },
    { keyword: "pathao", meaning: "send", english: "export", description: "Export symbol", example: 'pathao kaj add(a, b) { }' },
    { keyword: "hisabe", meaning: "as", english: "as", description: "Alias for imports", example: 'ano "mod.bang" hisabe m;' },
    { keyword: "chesta", meaning: "try/attempt", english: "try", description: "Try block", example: 'chesta { }' },
    { keyword: "dhoro_bhul", meaning: "catch error", english: "catch", description: "Catch block", example: 'dhoro_bhul (e) { }' },
    { keyword: "shesh", meaning: "end/finish", english: "finally", description: "Finally block", example: 'shesh { }' },
    { keyword: "felo", meaning: "throw", english: "throw", description: "Throw exception", example: 'felo "error message";' },
  ];

  return (
    <div>
      <div className="flex items-center gap-2 text-sm text-muted-foreground mb-4">
        <span className="px-2 py-1 bg-primary/10 text-primary rounded-full text-xs font-medium">
          Language Basics
        </span>
      </div>

      <h1>Keywords Reference</h1>

      <p className="lead text-xl text-muted-foreground mt-4">
        BanglaCode has 29 Bengali keywords written in Banglish (Bengali words using English/Latin script).
        This makes it easy to type on any keyboard while remaining familiar to Bengali speakers.
      </p>

      <h2>Complete Keyword Table</h2>

      <div className="overflow-x-auto my-6">
        <table className="min-w-full">
          <thead>
            <tr>
              <th className="text-left">Keyword</th>
              <th className="text-left">Bengali Meaning</th>
              <th className="text-left">English Equivalent</th>
              <th className="text-left">Description</th>
            </tr>
          </thead>
          <tbody>
            {keywords.map((kw, i) => (
              <tr key={i}>
                <td><code className="text-primary">{kw.keyword}</code></td>
                <td className="text-muted-foreground">{kw.meaning}</td>
                <td><code>{kw.english}</code></td>
                <td className="text-muted-foreground">{kw.description}</td>
              </tr>
            ))}
          </tbody>
        </table>
      </div>

      <h2>Keywords by Category</h2>

      <h3>Variable Declaration</h3>

      <CodeBlock
        code={`// dhoro - declare a mutable variable (like let/var)
dhoro naam = "Rahim";
dhoro boyosh = 25;

// sthir - declare a constant (cannot be changed)
sthir PI = 3.14159;
sthir MAX_SIZE = 100;

// bishwo - declare a global variable (accessible everywhere)
bishwo counter = 0;

kaj increment() {
    counter = counter + 1;  // Modifies global
}

increment();
dekho(counter);  // 1`}
      />

      <h3>Boolean & Null Literals</h3>

      <CodeBlock
        code={`// sotti - true
dhoro isValid = sotti;

// mittha - false
dhoro isComplete = mittha;

// khali - null
dhoro data = khali;`}
      />

      <h3>Logical Operators</h3>

      <CodeBlock
        code={`// ebong - logical AND
jodi (age >= 18 ebong hasLicense) {
    dekho("Can drive");
}

// ba - logical OR
jodi (isAdmin ba isModerator) {
    dekho("Has access");
}

// na - logical NOT
jodi (na isBlocked) {
    dekho("User is active");
}`}
      />

      <h3>Conditionals</h3>

      <CodeBlock
        code={`// jodi - if
// nahole - else
dhoro score = 85;

jodi (score >= 90) {
    dekho("Excellent!");
} nahole jodi (score >= 70) {
    dekho("Good job!");
} nahole {
    dekho("Keep trying!");
}`}
      />

      <h3>Loops</h3>

      <CodeBlock
        code={`// jotokkhon - while loop
dhoro i = 0;
jotokkhon (i < 5) {
    dekho(i);
    i = i + 1;
}

// ghuriye - for loop
ghuriye (dhoro j = 0; j < 5; j = j + 1) {
    dekho(j);
}

// thamo - break out of loop
// chharo - continue to next iteration
ghuriye (dhoro k = 0; k < 10; k = k + 1) {
    jodi (k == 3) {
        chharo;  // Skip 3
    }
    jodi (k == 7) {
        thamo;   // Stop at 7
    }
    dekho(k);
}`}
      />

      <h3>Functions</h3>

      <CodeBlock
        code={`// kaj - function declaration
// ferao - return value
kaj add(a, b) {
    ferao a + b;
}

kaj greet(naam) {
    dekho("Namaskar,", naam);
}

// Call functions
dhoro sum = add(5, 3);
greet("Rahim");`}
      />

      <h3>Classes & OOP</h3>

      <CodeBlock
        code={`// sreni - class
// shuru - constructor
// notun - new instance
// ei - this (current instance)

sreni Car {
    shuru(brand, model) {
        ei.brand = brand;
        ei.model = model;
    }

    kaj getInfo() {
        ferao ei.brand + " " + ei.model;
    }
}

dhoro myCar = notun Car("Toyota", "Corolla");
dekho(myCar.getInfo());`}
      />

      <h3>Modules</h3>

      <CodeBlock
        filename="math.bang"
        code={`// pathao - export
pathao kaj add(a, b) {
    ferao a + b;
}

pathao dhoro PI = 3.14159;`}
      />

      <CodeBlock
        filename="main.bang"
        code={`// ano - import
// hisabe - as (alias)
ano "math.bang" hisabe math;

dekho(math.add(5, 3));
dekho(math.PI);`}
      />

      <h3>Error Handling</h3>

      <CodeBlock
        code={`// chesta - try
// dhoro_bhul - catch
// shesh - finally
// felo - throw

chesta {
    dhoro result = riskyOperation();
} dhoro_bhul (error) {
    dekho("Error:", error);
} shesh {
    dekho("Cleanup done");
}

// Throw custom error
kaj validateAge(age) {
    jodi (age < 0) {
        felo "Age cannot be negative";
    }
}`}
      />

      <h2>Reserved Words</h2>

      <p>
        All 29 keywords are reserved and cannot be used as variable names, function names,
        or identifiers. Attempting to use a keyword as an identifier will result in a syntax error.
      </p>

      <DocNavigation currentPath="/docs/keywords" />
    </div>
  );
}
