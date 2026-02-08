export const EXAMPLES = {
  "hello.bang": {
    code: `// Hello World in BanglaCode
dekho("Hello, West Bengal!");
dekho("Namaskar!");

// Variables
dhoro naam = "Ankan";
dhoro boyosh = 25;
dekho("Amar naam", naam, "ebong ami", boyosh, "bochhor boyoshi");

// Type conversion
dekho("Type of naam:", dhoron(naam));
dekho("Type of boyosh:", dhoron(boyosh));
dekho("boyosh as lipi:", lipi(boyosh));`,
    output: [
      "Hello, West Bengal!",
      "Namaskar!",
      "Amar naam Ankan ebong ami 25 bochhor boyoshi",
      "Type of naam: string",
      "Type of boyosh: int",
      "boyosh as lipi: 25"
    ]
  },
  "fibonacci.bang": {
    code: `// Fibonacci Sequence
kaj fibonacci(n) {
    jodi (n <= 1) {
        ferao n;
    }
    ferao fibonacci(n - 1) + fibonacci(n - 2);
}

dekho("Fibonacci of 10:", fibonacci(10));

// Generate first 10 numbers
dekho("First 10 Fibonacci numbers:");
ghuriye (dhoro i = 0; i < 10; i = i + 1) {
    dekho(fibonacci(i));
}`,
    output: [
      "Fibonacci of 10: 55",
      "First 10 Fibonacci numbers:",
      "0", "1", "1", "2", "3", "5", "8", "13", "21", "34"
    ]
  },
  "homepage_demo.bang": {
    code: `// Loop example from Home Page
dhoro i = 0;
jotokkhon (i < 5) {
  dekho("Count: " + i);
  i = i + 1;
}`,
    output: [
      "Count: 0",
      "Count: 1",
      "Count: 2",
      "Count: 3",
      "Count: 4"
    ]
  },
  "conditions.bang": {
    code: `// Conditionals
dhoro score = 85;

jodi (score >= 90) {
    dekho("Grade: A");
} nahole jodi (score >= 80) {
    dekho("Grade: B");
} nahole {
    dekho("Grade: C");
}

dekho("Score:", score);`,
    output: [
      "Grade: B",
      "Score: 85"
    ]
  },
  "classes.bang": {
    code: `// Classes and Objects
sreni Manush {
    shuru(naam, boyosh) {
        ei.naam = naam;
        ei.boyosh = boyosh;
    }

    kaj porichoy() {
        dekho("Amar naam", ei.naam, "ebong ami", ei.boyosh, "bochhor boyoshi");
    }

    kaj birthday() {
        ei.boyosh = ei.boyosh + 1;
        dekho(ei.naam, "er ekhon", ei.boyosh, "bochhor");
    }
}

dhoro Ankan = notun Manush("Ankan", 25);
Ankan.porichoy();
Ankan.birthday();`,
    output: [
      "Amar naam Ankan ebong ami 25 bochhor boyoshi",
      "Ankan er ekhon 26 bochhor"
    ]
  }
};
