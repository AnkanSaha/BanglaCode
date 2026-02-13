// Comprehensive BanglaCode examples for the playground
export const EXAMPLES: Record<string, { name: string; description: string; code: string }> = {
  "hello.bang": {
    name: "Hello World",
    description: "Basic output and variables",
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
dekho("boyosh as lipi:", lipi(boyosh));`
  },

  "variables.bang": {
    name: "Variables & Data Types",
    description: "All data types in BanglaCode",
    code: `// Numbers (integers and floats)
dhoro purnoSonkha = 42;
dhoro dosomik = 3.14159;
dekho("Integer:", purnoSonkha);
dekho("Float:", dosomik);

// Strings
dhoro naam = "BanglaCode";
dhoro message = 'Namaskar Duniya!';
dekho("String 1:", naam);
dekho("String 2:", message);

// Booleans
dhoro satti = sotti;
dhoro mitha = mittha;
dekho("Boolean true:", satti);
dekho("Boolean false:", mitha);

// Null
dhoro khaliMaan = khali;
dekho("Null value:", khaliMaan);

// Arrays
dhoro numbers = [1, 2, 3, 4, 5];
dhoro fruits = ["Aam", "Kathal", "Lichu"];
dekho("Array:", numbers);
dekho("Fruits:", fruits);

// Maps (Objects)
dhoro person = {
    naam: "Ankan",
    boyosh: 25,
    city: "Kolkata"
};
dekho("Map:", person);
dekho("Name:", person.naam);`
  },

  "operators.bang": {
    name: "Operators",
    description: "Arithmetic, comparison, and logical operators",
    code: `// Arithmetic Operators
dhoro a = 10;
dhoro b = 3;

dekho("a =", a, ", b =", b);
dekho("a + b =", a + b);
dekho("a - b =", a - b);
dekho("a * b =", a * b);
dekho("a / b =", a / b);
dekho("a % b =", a % b);
dekho("a ** b =", a ** b);

// Comparison Operators
dekho("--- Comparisons ---");
dekho("5 == 5:", 5 == 5);
dekho("5 != 3:", 5 != 3);
dekho("5 < 10:", 5 < 10);
dekho("10 > 5:", 10 > 5);
dekho("5 <= 5:", 5 <= 5);
dekho("10 >= 5:", 10 >= 5);

// Logical Operators
dekho("--- Logical ---");
dekho("sotti ebong mittha:", sotti ebong mittha);
dekho("sotti ba mittha:", sotti ba mittha);
dekho("na sotti:", na sotti);

// Compound Assignment
dhoro x = 10;
x += 5;
dekho("After x += 5:", x);
x -= 3;
dekho("After x -= 3:", x);
x *= 2;
dekho("After x *= 2:", x);
x /= 4;
dekho("After x /= 4:", x);`
  },

  "conditionals.bang": {
    name: "Conditionals (If-Else)",
    description: "Conditional statements with jodi/nahole",
    code: `// Simple if
dhoro age = 20;

jodi (age >= 18) {
    dekho("You are an adult");
}

// If-else
dhoro score = 85;

jodi (score >= 90) {
    dekho("Grade: A");
} nahole {
    dekho("Grade: B or lower");
}

// If-else if-else chain
dhoro marks = 75;

jodi (marks >= 90) {
    dekho("Grade: A+");
} nahole jodi (marks >= 80) {
    dekho("Grade: A");
} nahole jodi (marks >= 70) {
    dekho("Grade: B");
} nahole jodi (marks >= 60) {
    dekho("Grade: C");
} nahole {
    dekho("Grade: F");
}

// Nested conditions
dhoro num = 15;

jodi (num > 0) {
    jodi (num % 2 == 0) {
        dekho(num, "is positive and even");
    } nahole {
        dekho(num, "is positive and odd");
    }
} nahole {
    dekho(num, "is not positive");
}`
  },

  "loops.bang": {
    name: "Loops",
    description: "While and for loops with break/continue",
    code: `// While loop (jotokkhon)
dekho("--- While Loop ---");
dhoro i = 0;
jotokkhon (i < 5) {
    dekho("Count:", i);
    i = i + 1;
}

// For loop (ghuriye)
dekho("--- For Loop ---");
ghuriye (dhoro j = 0; j < 5; j = j + 1) {
    dekho("Iteration:", j);
}

// Break (thamo)
dekho("--- Break Example ---");
dhoro k = 0;
jotokkhon (k < 10) {
    jodi (k == 5) {
        dekho("Breaking at", k);
        thamo;
    }
    dekho("k =", k);
    k = k + 1;
}

// Continue (chharo)
dekho("--- Continue Example ---");
ghuriye (dhoro m = 0; m < 5; m = m + 1) {
    jodi (m == 2) {
        chharo;
    }
    dekho("m =", m);
}

// Nested loops
dekho("--- Multiplication Table (3x3) ---");
ghuriye (dhoro x = 1; x <= 3; x = x + 1) {
    ghuriye (dhoro y = 1; y <= 3; y = y + 1) {
        dekho(x, "x", y, "=", x * y);
    }
}`
  },

  "functions.bang": {
    name: "Functions",
    description: "Function definitions and recursion",
    code: `// Simple function
kaj greet(naam) {
    dekho("Hello,", naam, "!");
}

greet("World");
greet("BanglaCode");

// Function with return value
kaj add(a, b) {
    ferao a + b;
}

dhoro result = add(5, 3);
dekho("5 + 3 =", result);

// Multiple parameters
kaj calculate(x, y, z) {
    ferao (x + y) * z;
}

dekho("(2 + 3) * 4 =", calculate(2, 3, 4));

// Recursive function - Factorial
kaj factorial(n) {
    jodi (n <= 1) {
        ferao 1;
    }
    ferao n * factorial(n - 1);
}

dekho("5! =", factorial(5));
dekho("10! =", factorial(10));

// Recursive function - Fibonacci
kaj fibonacci(n) {
    jodi (n <= 1) {
        ferao n;
    }
    ferao fibonacci(n - 1) + fibonacci(n - 2);
}

dekho("--- First 10 Fibonacci numbers ---");
ghuriye (dhoro i = 0; i < 10; i = i + 1) {
    dekho("F(" + i + ") =", fibonacci(i));
}

// Higher-order function (function as parameter)
kaj applyTwice(fn, x) {
    ferao fn(fn(x));
}

kaj double(n) {
    ferao n * 2;
}

dekho("Double applied twice to 5:", applyTwice(double, 5));`
  },

  "classes.bang": {
    name: "Classes & OOP",
    description: "Object-oriented programming with classes",
    code: `// Class definition
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

// Create instances
dhoro person1 = notun Manush("Ankan", 25);
dhoro person2 = notun Manush("Rahim", 30);

person1.porichoy();
person2.porichoy();

person1.birthday();
person1.porichoy();

// Another class example
sreni Rectangle {
    shuru(width, height) {
        ei.width = width;
        ei.height = height;
    }

    kaj area() {
        ferao ei.width * ei.height;
    }

    kaj perimeter() {
        ferao 2 * (ei.width + ei.height);
    }

    kaj describe() {
        dekho("Rectangle", ei.width, "x", ei.height);
        dekho("Area:", ei.area());
        dekho("Perimeter:", ei.perimeter());
    }
}

dhoro rect = notun Rectangle(10, 5);
rect.describe();

// Modify properties
rect.width = 15;
rect.describe();`
  },

  "arrays.bang": {
    name: "Arrays",
    description: "Array operations and built-in functions",
    code: `// Create arrays
dhoro numbers = [1, 2, 3, 4, 5];
dhoro fruits = ["Aam", "Kathal", "Lichu", "Kola"];
dhoro mixed = [1, "hello", sotti, khali];

dekho("Numbers:", numbers);
dekho("Fruits:", fruits);
dekho("Mixed:", mixed);

// Access elements
dekho("First fruit:", fruits[0]);
dekho("Last fruit:", fruits[3]);

// Modify elements
fruits[1] = "Jamun";
dekho("After modification:", fruits);

// Array length
dekho("Length of fruits:", dorghyo(fruits));

// Push - dhokao
dhokao(numbers, 6);
dhokao(numbers, 7);
dekho("After push:", numbers);

// Pop - berKoro
dhoro last = berKoro(numbers);
dekho("Popped:", last);
dekho("After pop:", numbers);

// Slice - kato
dhoro sliced = kato(numbers, 1, 4);
dekho("Sliced [1:4]:", sliced);

// Reverse - ulto
dhoro reversed = ulto(numbers);
dekho("Reversed:", reversed);

// Sort - saja
dhoro unsorted = [5, 2, 8, 1, 9, 3];
dhoro sorted = saja(unsorted);
dekho("Original:", unsorted);
dekho("Sorted:", sorted);

// Includes - ache
dekho("Has 3?:", ache(numbers, 3));
dekho("Has 99?:", ache(numbers, 99));

// Iterate over array
dekho("--- Iterating ---");
ghuriye (dhoro i = 0; i < dorghyo(fruits); i = i + 1) {
    dekho("Fruit", i + 1, ":", fruits[i]);
}`
  },

  "maps.bang": {
    name: "Maps (Objects)",
    description: "Key-value collections",
    code: `// Create maps
dhoro person = {
    naam: "Ankan",
    boyosh: 25,
    city: "Kolkata",
    isStudent: sotti
};

dekho("Person:", person);

// Access values
dekho("Name:", person.naam);
dekho("Age:", person["boyosh"]);

// Modify values
person.boyosh = 26;
person.country = "India";
dekho("After modification:", person);

// Get all keys
dhoro keys = chabi(person);
dekho("Keys:", keys);

// Iterate over map
dekho("--- Iterating ---");
ghuriye (dhoro i = 0; i < dorghyo(keys); i = i + 1) {
    dhoro key = keys[i];
    dekho(key, ":", person[key]);
}

// Nested maps
dhoro config = {
    app: "BanglaCode",
    version: "1.0.0",
    author: {
        name: "Ankan",
        location: "West Bengal"
    },
    features: ["Bengali syntax", "OOP", "Modules"]
};

dekho("App:", config.app);
dekho("Author name:", config.author.name);
dekho("Features:", config.features);`
  },

  "strings.bang": {
    name: "String Functions",
    description: "String manipulation built-ins",
    code: `dhoro str = "  Hello BanglaCode World  ";

dekho("Original:", str);
dekho("Length:", dorghyo(str));

// Trim
dekho("Trimmed:", chhanto(str));

// Case conversion
dekho("Uppercase:", boroHater("hello"));
dekho("Lowercase:", chotoHater("HELLO"));

// Split and join
dhoro csv = "apple,banana,cherry";
dhoro parts = bhag(csv, ",");
dekho("Split:", parts);

dhoro joined = joro(parts, " - ");
dekho("Joined:", joined);

// Find
dhoro text = "Hello World";
dekho("Index of 'World':", khojo(text, "World"));
dekho("Index of 'xyz':", khojo(text, "xyz"));

// Substring
dekho("Substring [0:5]:", angsho(text, 0, 5));
dekho("Substring [6:]:", angsho(text, 6));

// Replace
dhoro original = "hello hello hello";
dhoro replaced = bodlo(original, "hello", "hi");
dekho("Replaced:", replaced);

// String concatenation
dhoro greeting = "Namaskar";
dhoro name = "BanglaCode";
dekho(greeting + ", " + name + "!");`
  },

  "math.bang": {
    name: "Math Functions",
    description: "Mathematical built-in functions",
    code: `// Square root
dekho("sqrt(16) =", borgomul(16));
dekho("sqrt(2) =", borgomul(2));

// Power
dekho("2^10 =", ghat(2, 10));
dekho("3^4 =", ghat(3, 4));

// Rounding
dhoro num = 4.7;
dekho("Value:", num);
dekho("Floor:", niche(num));
dekho("Ceil:", upore(num));
dekho("Round:", kache(num));

dhoro num2 = 4.3;
dekho("Value:", num2);
dekho("Round:", kache(num2));

// Absolute value
dekho("Absolute of -42:", niratek(-42));
dekho("Absolute of 42:", niratek(42));

// Min and Max
dekho("Min(5, 2, 8, 1):", choto(5, 2, 8, 1));
dekho("Max(5, 2, 8, 1):", boro(5, 2, 8, 1));

// Random numbers
dekho("--- Random Numbers ---");
ghuriye (dhoro i = 0; i < 5; i = i + 1) {
    dekho("Random:", lotto());
}

// Random integer in range
kaj randomInt(min, max) {
    ferao niche(lotto() * (max - min + 1)) + min;
}

dekho("--- Random integers 1-10 ---");
ghuriye (dhoro i = 0; i < 5; i = i + 1) {
    dekho("Random int:", randomInt(1, 10));
}`
  },

  "error_handling.bang": {
    name: "Error Handling",
    description: "Try-catch-finally with throw",
    code: `// Basic try-catch
chesta {
    dekho("Trying something...");
    felo "Something went wrong!";
    dekho("This won't print");
} dhoro_bhul (err) {
    dekho("Caught error:", err);
}

// Try-catch-finally
dekho("--- With Finally ---");
chesta {
    dekho("Opening resource...");
    felo "Resource error!";
} dhoro_bhul (err) {
    dekho("Error:", err);
} shesh {
    dekho("Cleanup: Closing resource");
}

// Finally always runs
dekho("--- Finally without error ---");
chesta {
    dekho("Operation successful");
} dhoro_bhul (err) {
    dekho("Error:", err);
} shesh {
    dekho("Finally block runs anyway");
}

// Safe division function
kaj safeDivide(a, b) {
    chesta {
        jodi (b == 0) {
            felo "Division by zero!";
        }
        ferao a / b;
    } dhoro_bhul (err) {
        dekho("Warning:", err);
        ferao khali;
    }
}

dekho("10 / 2 =", safeDivide(10, 2));
dekho("10 / 0 =", safeDivide(10, 0));

// Nested try-catch
dekho("--- Nested Try-Catch ---");
chesta {
    dekho("Outer try");
    chesta {
        dekho("Inner try");
        felo "Inner error";
    } dhoro_bhul (innerErr) {
        dekho("Caught inner:", innerErr);
        felo "Re-throwing as outer error";
    }
} dhoro_bhul (outerErr) {
    dekho("Caught outer:", outerErr);
}`
  },

  "json.bang": {
    name: "JSON Operations",
    description: "Parse and stringify JSON",
    code: `// Parse JSON string
dhoro jsonStr = '{"naam": "Ankan", "boyosh": 25, "skills": ["Go", "JavaScript"]}';
dhoro data = json_poro(jsonStr);

dekho("Parsed object:", data);
dekho("Name:", data.naam);
dekho("Age:", data.boyosh);
dekho("Skills:", data.skills);
dekho("First skill:", data.skills[0]);

// Create JSON string
dhoro person = {
    naam: "Rahim",
    city: "Dhaka",
    active: sotti,
    scores: [85, 90, 78]
};

dhoro json = json_banao(person);
dekho("JSON string:", json);

// Array to JSON
dhoro arr = [1, 2, 3, "hello", sotti, khali];
dekho("Array as JSON:", json_banao(arr));

// Parse and modify
dhoro config = json_poro('{"theme": "dark", "fontSize": 14}');
config.fontSize = 16;
config.language = "bn";
dekho("Modified:", json_banao(config));`
  },

  "algorithms.bang": {
    name: "Algorithms",
    description: "Common algorithms implemented in BanglaCode",
    code: `// Bubble Sort
kaj bubbleSort(arr) {
    dhoro n = dorghyo(arr);
    ghuriye (dhoro i = 0; i < n - 1; i = i + 1) {
        ghuriye (dhoro j = 0; j < n - i - 1; j = j + 1) {
            jodi (arr[j] > arr[j + 1]) {
                dhoro temp = arr[j];
                arr[j] = arr[j + 1];
                arr[j + 1] = temp;
            }
        }
    }
    ferao arr;
}

dhoro unsorted = [64, 34, 25, 12, 22, 11, 90];
dekho("Before sort:", unsorted);
bubbleSort(unsorted);
dekho("After sort:", unsorted);

// Binary Search
kaj binarySearch(arr, target) {
    dhoro left = 0;
    dhoro right = dorghyo(arr) - 1;
    
    jotokkhon (left <= right) {
        dhoro mid = niche((left + right) / 2);
        
        jodi (arr[mid] == target) {
            ferao mid;
        } nahole jodi (arr[mid] < target) {
            left = mid + 1;
        } nahole {
            right = mid - 1;
        }
    }
    ferao -1;
}

dhoro sorted = [1, 3, 5, 7, 9, 11, 13, 15];
dekho("Array:", sorted);
dekho("Index of 7:", binarySearch(sorted, 7));
dekho("Index of 4:", binarySearch(sorted, 4));

// Prime number checker
kaj isPrime(n) {
    jodi (n <= 1) {
        ferao mittha;
    }
    jodi (n <= 3) {
        ferao sotti;
    }
    jodi (n % 2 == 0 ba n % 3 == 0) {
        ferao mittha;
    }
    dhoro i = 5;
    jotokkhon (i * i <= n) {
        jodi (n % i == 0 ba n % (i + 2) == 0) {
            ferao mittha;
        }
        i = i + 6;
    }
    ferao sotti;
}

dekho("--- Prime numbers 1-30 ---");
ghuriye (dhoro n = 1; n <= 30; n = n + 1) {
    jodi (isPrime(n)) {
        dekho(n, "is prime");
    }
}

// GCD (Greatest Common Divisor)
kaj gcd(a, b) {
    jotokkhon (b != 0) {
        dhoro temp = b;
        b = a % b;
        a = temp;
    }
    ferao a;
}

dekho("GCD(48, 18) =", gcd(48, 18));
dekho("GCD(54, 24) =", gcd(54, 24));`
  },

  "calculator.bang": {
    name: "Calculator",
    description: "A simple calculator class",
    code: `// Calculator class
sreni Calculator {
    shuru() {
        ei.result = 0;
        ei.history = [];
    }

    kaj add(n) {
        ei.result = ei.result + n;
        dhokao(ei.history, "+" + n);
        ferao ei;
    }

    kaj subtract(n) {
        ei.result = ei.result - n;
        dhokao(ei.history, "-" + n);
        ferao ei;
    }

    kaj multiply(n) {
        ei.result = ei.result * n;
        dhokao(ei.history, "*" + n);
        ferao ei;
    }

    kaj divide(n) {
        jodi (n == 0) {
            dekho("Error: Cannot divide by zero!");
            ferao ei;
        }
        ei.result = ei.result / n;
        dhokao(ei.history, "/" + n);
        ferao ei;
    }

    kaj clear() {
        ei.result = 0;
        ei.history = [];
        dekho("Calculator cleared");
        ferao ei;
    }

    kaj showResult() {
        dekho("Result:", ei.result);
        ferao ei;
    }

    kaj showHistory() {
        dekho("History:", joro(ei.history, " "));
        ferao ei;
    }
}

// Use the calculator with method chaining
dhoro calc = notun Calculator();

calc.add(10).showResult();
calc.multiply(5).showResult();
calc.subtract(20).showResult();
calc.divide(2).showResult();
calc.showHistory();

// Another calculation
calc.clear();
calc.add(100).divide(4).multiply(3).subtract(25).showResult();
calc.showHistory();`
  },

  "student_system.bang": {
    name: "Student System",
    description: "Complete OOP example with Student class",
    code: `// Student management system
sreni Student {
    shuru(naam, roll) {
        ei.naam = naam;
        ei.roll = roll;
        ei.marks = [];
    }

    kaj addMark(subject, mark) {
        dhokao(ei.marks, {subject: subject, mark: mark});
        dekho("Added", subject, ":", mark);
        ferao ei;
    }

    kaj calculateAverage() {
        jodi (dorghyo(ei.marks) == 0) {
            ferao 0;
        }
        dhoro total = 0;
        ghuriye (dhoro i = 0; i < dorghyo(ei.marks); i = i + 1) {
            total = total + ei.marks[i].mark;
        }
        ferao total / dorghyo(ei.marks);
    }

    kaj getGrade() {
        dhoro avg = ei.calculateAverage();
        jodi (avg >= 90) {
            ferao "A+";
        } nahole jodi (avg >= 80) {
            ferao "A";
        } nahole jodi (avg >= 70) {
            ferao "B";
        } nahole jodi (avg >= 60) {
            ferao "C";
        } nahole {
            ferao "F";
        }
    }

    kaj displayReport() {
        dekho("================================");
        dekho("STUDENT REPORT");
        dekho("================================");
        dekho("Name:", ei.naam);
        dekho("Roll:", ei.roll);
        dekho("--------------------------------");
        dekho("MARKS:");
        ghuriye (dhoro i = 0; i < dorghyo(ei.marks); i = i + 1) {
            dhoro m = ei.marks[i];
            dekho("  ", m.subject, ":", m.mark);
        }
        dekho("--------------------------------");
        dekho("Average:", ei.calculateAverage());
        dekho("Grade:", ei.getGrade());
        dekho("================================");
    }
}

// Create students
dhoro student1 = notun Student("Ankan", 101);
student1.addMark("Bangla", 85);
student1.addMark("English", 90);
student1.addMark("Math", 95);
student1.addMark("Science", 88);
student1.displayReport();

dhoro student2 = notun Student("Priya", 102);
student2.addMark("Bangla", 78);
student2.addMark("English", 82);
student2.addMark("Math", 65);
student2.addMark("Science", 72);
student2.displayReport();`
  }
};
