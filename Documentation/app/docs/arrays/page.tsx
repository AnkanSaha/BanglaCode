import CodeBlock from "@/components/CodeBlock";
import DocNavigation from "@/components/DocNavigation";

export default function Arrays() {
  return (
    <div>
      <div className="flex items-center gap-2 text-sm text-muted-foreground mb-4">
        <span className="px-2 py-1 bg-primary/10 text-primary rounded-full text-xs font-medium">
          Data Structures
        </span>
      </div>

      <h1>Arrays</h1>

      <p className="lead text-xl text-muted-foreground mt-4">
        Arrays in BanglaCode are ordered collections that can hold elements of any type.
        They are zero-indexed and dynamically sized.
      </p>

      <h2>Creating Arrays</h2>

      <CodeBlock
        code={`// Empty array
dhoro empty = [];

// Array with elements
dhoro numbers = [1, 2, 3, 4, 5];
dhoro names = ["Rahim", "Karim", "Jamil"];

// Mixed types
dhoro mixed = [42, "hello", sotti, khali, [1, 2]];

// Nested arrays
dhoro matrix = [
    [1, 2, 3],
    [4, 5, 6],
    [7, 8, 9]
];`}
      />

      <h2>Accessing Elements</h2>

      <p>
        Arrays are zero-indexed (first element is at index 0):
      </p>

      <CodeBlock
        code={`dhoro fruits = ["Apple", "Banana", "Mango", "Orange"];

// Access by index
dekho(fruits[0]);   // "Apple"
dekho(fruits[2]);   // "Mango"

// Access last element
dhoro lastIndex = dorghyo(fruits) - 1;
dekho(fruits[lastIndex]);  // "Orange"

// Nested array access
dhoro matrix = [[1, 2], [3, 4], [5, 6]];
dekho(matrix[1][0]);  // 3
dekho(matrix[2][1]);  // 6`}
      />

      <h2>Modifying Arrays</h2>

      <CodeBlock
        code={`dhoro arr = [10, 20, 30];

// Modify element
arr[1] = 25;
dekho(arr);  // [10, 25, 30]

// Add element with index
arr[3] = 40;
dekho(arr);  // [10, 25, 30, 40]`}
      />

      <h2>Array Built-in Functions</h2>

      <h3>dorghyo - Length</h3>

      <CodeBlock
        code={`dhoro arr = [1, 2, 3, 4, 5];
dekho(dorghyo(arr));  // 5

dhoro empty = [];
dekho(dorghyo(empty));  // 0`}
      />

      <h3>dhokao - Push (Add to End)</h3>

      <CodeBlock
        code={`dhoro arr = [1, 2, 3];

dhokao(arr, 4);
dekho(arr);  // [1, 2, 3, 4]

dhokao(arr, 5);
dhokao(arr, 6);
dekho(arr);  // [1, 2, 3, 4, 5, 6]`}
      />

      <h3>berKoro - Pop (Remove from End)</h3>

      <CodeBlock
        code={`dhoro arr = [1, 2, 3, 4, 5];

dhoro last = berKoro(arr);
dekho(last);  // 5
dekho(arr);   // [1, 2, 3, 4]

dhoro another = berKoro(arr);
dekho(another);  // 4
dekho(arr);      // [1, 2, 3]`}
      />

      <h3>kato - Slice</h3>

      <CodeBlock
        code={`dhoro arr = [0, 1, 2, 3, 4, 5, 6, 7, 8, 9];

// slice(start, end) - excludes end index
dekho(kato(arr, 2, 5));   // [2, 3, 4]
dekho(kato(arr, 0, 3));   // [0, 1, 2]

// slice(start) - from start to end
dekho(kato(arr, 5));      // [5, 6, 7, 8, 9]
dekho(kato(arr, 7));      // [7, 8, 9]

// Original array unchanged
dekho(arr);  // [0, 1, 2, 3, 4, 5, 6, 7, 8, 9]`}
      />

      <h3>ulto - Reverse</h3>

      <CodeBlock
        code={`dhoro arr = [1, 2, 3, 4, 5];
dhoro reversed = ulto(arr);

dekho(reversed);  // [5, 4, 3, 2, 1]
dekho(arr);       // [1, 2, 3, 4, 5] (original unchanged)`}
      />

      <h3>ache - Contains/Includes</h3>

      <CodeBlock
        code={`dhoro fruits = ["Apple", "Banana", "Mango"];

dekho(ache(fruits, "Banana"));  // sotti
dekho(ache(fruits, "Orange"));  // mittha

// Use in conditions
jodi (ache(fruits, "Mango")) {
    dekho("Mango is available!");
}`}
      />

      <h3>saja - Sort</h3>

      <CodeBlock
        code={`// Sort numbers
dhoro numbers = [3, 1, 4, 1, 5, 9, 2, 6];
dekho(saja(numbers));  // [1, 1, 2, 3, 4, 5, 6, 9]

// Sort strings
dhoro names = ["Karim", "Rahim", "Abdul", "Zahir"];
dekho(saja(names));  // ["Abdul", "Karim", "Rahim", "Zahir"]`}
      />

      <h3>joro - Join</h3>

      <CodeBlock
        code={`dhoro words = ["Hello", "World", "BanglaCode"];

dekho(joro(words, " "));    // "Hello World BanglaCode"
dekho(joro(words, ", "));   // "Hello, World, BanglaCode"
dekho(joro(words, "-"));    // "Hello-World-BanglaCode"
dekho(joro(words, ""));     // "HelloWorldBanglaCode"`}
      />

      <h2>Array Higher-Order Methods</h2>

      <h3>manchitro - Map</h3>

      <p>
        Transforms each element in an array using a callback function. Returns a new array with the transformed elements.
      </p>

      <CodeBlock
        code={`// Double each number
dhoro numbers = [1, 2, 3, 4, 5];
dhoro doubled = manchitro(numbers, kaj(x) {
    ferao x * 2;
});
dekho(doubled);  // [2, 4, 6, 8, 10]

// Extract property from objects
dhoro people = [
    {"naam": "Rahim", "age": 25},
    {"naam": "Karim", "age": 30}
];
dhoro names = manchitro(people, kaj(person) {
    ferao person["naam"];
});
dekho(names);  // ["Rahim", "Karim"]

// Using index parameter
dhoro squared = manchitro(numbers, kaj(x, i) {
    ferao x * x;
});
dekho(squared);  // [1, 4, 9, 16, 25]`}
      />

      <h3>chhanno - Filter</h3>

      <p>
        Filters array elements based on a condition. Returns a new array containing only elements where the callback returns true.
      </p>

      <CodeBlock
        code={`// Filter even numbers
dhoro numbers = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10];
dhoro evens = chhanno(numbers, kaj(x) {
    ferao x % 2 == 0;
});
dekho(evens);  // [2, 4, 6, 8, 10]

// Filter by condition
dhoro scores = [45, 78, 92, 35, 88, 55];
dhoro passing = chhanno(scores, kaj(score) {
    ferao score >= 50;
});
dekho(passing);  // [78, 92, 88, 55]

// Filter with index
dhoro colors = ["Red", "Green", "Blue", "Yellow"];
dhoro result = chhanno(colors, kaj(color, index) {
    ferao index % 2 == 0;  // Keep items at even indices
});
dekho(result);  // ["Red", "Blue"]`}
      />

      <h3>sonkuchito - Reduce</h3>

      <p>
        Reduces array to a single value by applying a function to accumulate the result. Takes an optional initial value.
      </p>

      <CodeBlock
        code={`// Sum all numbers
dhoro numbers = [1, 2, 3, 4, 5];
dhoro sum = sonkuchito(numbers, kaj(acc, x) {
    ferao acc + x;
}, 0);
dekho(sum);  // 15

// Calculate product
dhoro product = sonkuchito(numbers, kaj(acc, x) {
    ferao acc * x;
}, 1);
dekho(product);  // 120

// Build a string
dhoro words = ["Hello", "World", "BanglaCode"];
dhoro sentence = sonkuchito(words, kaj(acc, word) {
    jodi (acc == "") {
        ferao word;
    }
    ferao acc + " " + word;
}, "");
dekho(sentence);  // "Hello World BanglaCode"

// Count occurrences
dhoro arr = [1, 2, 2, 3, 3, 3, 4];
dhoro counts = sonkuchito(arr, kaj(acc, x) {
    acc[lipi(x)] = (acc[lipi(x)] || 0) + 1;
    ferao acc;
}, {});
dekho(counts);  // {"1": 1, "2": 2, "3": 3, "4": 1}`}
      />

      <h3>proti - ForEach</h3>

      <p>
        Executes a callback function for each element. Returns null. Useful for side effects like logging.
      </p>

      <CodeBlock
        code={`// Print each element
dhoro fruits = ["Apple", "Banana", "Mango"];
proti(fruits, kaj(fruit) {
    dekho("- " + fruit);
});
// Output:
// - Apple
// - Banana
// - Mango

// Enumerate with index
dhoro items = ["First", "Second", "Third"];
proti(items, kaj(item, index) {
    dekho((index + 1) + ". " + item);
});
// Output:
// 1. First
// 2. Second
// 3. Third

// Side effects (mutations)
dhoro counter = 0;
proti([10, 20, 30], kaj(x) {
    counter = counter + 1;
});
dekho("Processed:", counter);  // Processed: 3`}
      />

      <h2>Chaining Higher-Order Methods</h2>

      <CodeBlock
        code={`// Chain multiple operations
dhoro numbers = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10];

// Filter > Map > Reduce: Get sum of squares of even numbers
dhoro evens = chhanno(numbers, kaj(x) { ferao x % 2 == 0; });
dhoro squares = manchitro(evens, kaj(x) { ferao x * x; });
dhoro sumOfSquares = sonkuchito(squares, kaj(acc, x) { ferao acc + x; }, 0);
dekho(sumOfSquares);  // 220 (4 + 16 + 36 + 64 + 100)

// Practical: Process user data
dhoro users = [
    {"naam": "Rahim", "age": 25, "active": sotti},
    {"naam": "Karim", "age": 17, "active": mittha},
    {"naam": "Jamil", "age": 30, "active": sotti}
];

// Get names of active adults
dhoro activeAdults = chhanno(users, kaj(user) {
    ferao user["active"] ebong user["age"] >= 18;
});
dhoro activeNames = manchitro(activeAdults, kaj(user) {
    ferao user["naam"];
});
dekho(activeNames);  // ["Rahim", "Jamil"]`}
      />

      <h2>Iterating Over Arrays</h2>

      <CodeBlock
        code={`dhoro colors = ["Red", "Green", "Blue", "Yellow"];

// Using for loop
ghuriye (dhoro i = 0; i < dorghyo(colors); i = i + 1) {
    dekho(i + 1, "-", colors[i]);
}

// Process each element
dhoro numbers = [1, 2, 3, 4, 5];
dhoro sum = 0;

ghuriye (dhoro i = 0; i < dorghyo(numbers); i = i + 1) {
    sum = sum + numbers[i];
}

dekho("Sum:", sum);  // Sum: 15`}
      />

      <h2>Common Array Patterns</h2>

      <h3>Finding Maximum/Minimum</h3>

      <CodeBlock
        code={`dhoro numbers = [34, 12, 89, 45, 23, 67];

// Find maximum
dhoro max = numbers[0];
ghuriye (dhoro i = 1; i < dorghyo(numbers); i = i + 1) {
    jodi (numbers[i] > max) {
        max = numbers[i];
    }
}
dekho("Max:", max);  // 89

// Find minimum
dhoro min = numbers[0];
ghuriye (dhoro i = 1; i < dorghyo(numbers); i = i + 1) {
    jodi (numbers[i] < min) {
        min = numbers[i];
    }
}
dekho("Min:", min);  // 12

// Using built-ins
dekho("Max:", boro(34, 12, 89, 45, 23, 67));  // 89
dekho("Min:", choto(34, 12, 89, 45, 23, 67)); // 12`}
      />

      <h3>Filter Array</h3>

      <CodeBlock
        code={`dhoro numbers = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10];

// Filter even numbers
dhoro evens = [];
ghuriye (dhoro i = 0; i < dorghyo(numbers); i = i + 1) {
    jodi (numbers[i] % 2 == 0) {
        dhokao(evens, numbers[i]);
    }
}
dekho(evens);  // [2, 4, 6, 8, 10]`}
      />

      <h3>Map/Transform Array</h3>

      <CodeBlock
        code={`dhoro numbers = [1, 2, 3, 4, 5];

// Double each element
dhoro doubled = [];
ghuriye (dhoro i = 0; i < dorghyo(numbers); i = i + 1) {
    dhokao(doubled, numbers[i] * 2);
}
dekho(doubled);  // [2, 4, 6, 8, 10]`}
      />

      <h3>Remove Duplicates</h3>

      <CodeBlock
        code={`dhoro arr = [1, 2, 2, 3, 3, 3, 4, 5, 5];

dhoro unique = [];
ghuriye (dhoro i = 0; i < dorghyo(arr); i = i + 1) {
    jodi (na ache(unique, arr[i])) {
        dhokao(unique, arr[i]);
    }
}
dekho(unique);  // [1, 2, 3, 4, 5]`}
      />

      <h2>Multi-dimensional Arrays</h2>

      <CodeBlock
        code={`// 2D array (matrix)
dhoro matrix = [
    [1, 2, 3],
    [4, 5, 6],
    [7, 8, 9]
];

// Access elements
dekho(matrix[0][0]);  // 1
dekho(matrix[1][1]);  // 5
dekho(matrix[2][2]);  // 9

// Iterate over 2D array
ghuriye (dhoro i = 0; i < dorghyo(matrix); i = i + 1) {
    dhoro row = "";
    ghuriye (dhoro j = 0; j < dorghyo(matrix[i]); j = j + 1) {
        row = row + lipi(matrix[i][j]) + " ";
    }
    dekho(row);
}

// Sum all elements
dhoro total = 0;
ghuriye (dhoro i = 0; i < dorghyo(matrix); i = i + 1) {
    ghuriye (dhoro j = 0; j < dorghyo(matrix[i]); j = j + 1) {
        total = total + matrix[i][j];
    }
}
dekho("Sum:", total);  // 45`}
      />

      <h2>Practical Examples</h2>

      <h3>Stack Implementation</h3>

      <CodeBlock
        code={`dhoro stack = [];

// Push
dhokao(stack, "first");
dhokao(stack, "second");
dhokao(stack, "third");

// Pop
dekho(berKoro(stack));  // "third"
dekho(berKoro(stack));  // "second"

// Peek (without removing)
dekho(stack[dorghyo(stack) - 1]);  // "first"`}
      />

      <h3>Queue Implementation</h3>

      <CodeBlock
        code={`dhoro queue = [];

// Enqueue (add to end)
dhokao(queue, "first");
dhokao(queue, "second");
dhokao(queue, "third");

// Dequeue (remove from front using slice)
dhoro front = queue[0];
queue = kato(queue, 1);
dekho(front);  // "first"
dekho(queue);  // ["second", "third"]`}
      />

      <DocNavigation currentPath="/docs/arrays" />
    </div>
  );
}
