import CodeBlock from "@/components/CodeBlock";
import DocNavigation from "@/components/DocNavigation";

export default function Methods() {
  return (
    <div>
      <div className="flex items-center gap-2 text-sm text-muted-foreground mb-4">
        <span className="px-2 py-1 bg-primary/10 text-primary rounded-full text-xs font-medium">
          Functions & OOP
        </span>
      </div>

      <h1>Methods</h1>

      <p className="lead text-xl text-muted-foreground mt-4">
        Methods are functions defined inside a class that operate on instance data
        using the <code>ei</code> (this) keyword.
      </p>

      <h2>Defining Methods</h2>

      <p>
        Methods are defined using the <code>kaj</code> keyword inside a class:
      </p>

      <CodeBlock
        code={`sreni Calculator {
    shuru() {
        ei.result = 0;
    }

    kaj add(n) {
        ei.result = ei.result + n;
        ferao ei;  // Return this for method chaining
    }

    kaj subtract(n) {
        ei.result = ei.result - n;
        ferao ei;
    }

    kaj multiply(n) {
        ei.result = ei.result * n;
        ferao ei;
    }

    kaj divide(n) {
        jodi (n != 0) {
            ei.result = ei.result / n;
        } nahole {
            dekho("Cannot divide by zero!");
        }
        ferao ei;
    }

    kaj getResult() {
        ferao ei.result;
    }

    kaj reset() {
        ei.result = 0;
        ferao ei;
    }
}

dhoro calc = notun Calculator();
dhoro result = calc.add(10).multiply(2).subtract(5).getResult();
dekho(result);  // 15`}
      />

      <h2>Accessing Instance Data with &apos;ei&apos;</h2>

      <p>
        The <code>ei</code> keyword (Bengali for &quot;this&quot;) refers to the current instance:
      </p>

      <CodeBlock
        code={`sreni Person {
    shuru(naam) {
        ei.naam = naam;
        ei.friends = [];
    }

    kaj addFriend(friend) {
        // 'ei' accesses this instance's properties
        dhokao(ei.friends, friend);
        dekho(ei.naam, "added", friend.naam, "as a friend");
    }

    kaj listFriends() {
        dekho(ei.naam + "'s friends:");
        ghuriye (dhoro i = 0; i < dorghyo(ei.friends); i = i + 1) {
            dekho(" -", ei.friends[i].naam);
        }
    }
}

dhoro rahim = notun Person("Rahim");
dhoro karim = notun Person("Karim");
dhoro jamil = notun Person("Jamil");

rahim.addFriend(karim);
rahim.addFriend(jamil);
rahim.listFriends();`}
      />

      <h2>Method Parameters</h2>

      <CodeBlock
        code={`sreni ShoppingCart {
    shuru() {
        ei.items = [];
        ei.total = 0;
    }

    kaj addItem(name, price, quantity) {
        dhokao(ei.items, {
            name: name,
            price: price,
            quantity: quantity
        });
        ei.total = ei.total + (price * quantity);
        dekho("Added:", quantity, "x", name, "@ Tk.", price);
    }

    kaj removeItem(name) {
        ghuriye (dhoro i = 0; i < dorghyo(ei.items); i = i + 1) {
            jodi (ei.items[i].name == name) {
                dhoro item = ei.items[i];
                ei.total = ei.total - (item.price * item.quantity);
                // Remove item logic would go here
                dekho("Removed:", name);
                ferao sotti;
            }
        }
        ferao mittha;
    }

    kaj applyDiscount(percentage) {
        dhoro discount = ei.total * (percentage / 100);
        ei.total = ei.total - discount;
        dekho("Applied", percentage, "% discount. Saved Tk.", discount);
    }

    kaj checkout() {
        dekho("===== Receipt =====");
        ghuriye (dhoro i = 0; i < dorghyo(ei.items); i = i + 1) {
            dhoro item = ei.items[i];
            dekho(item.quantity, "x", item.name, ":", "Tk.", item.price * item.quantity);
        }
        dekho("-------------------");
        dekho("Total: Tk.", ei.total);
    }
}

dhoro cart = notun ShoppingCart();
cart.addItem("Rice (5kg)", 350, 2);
cart.addItem("Oil (1L)", 180, 1);
cart.addItem("Sugar (1kg)", 85, 3);
cart.applyDiscount(10);
cart.checkout();`}
      />

      <h2>Methods Calling Other Methods</h2>

      <CodeBlock
        code={`sreni StringHelper {
    shuru(text) {
        ei.text = text;
    }

    kaj toUpperCase() {
        ei.text = boroHater(ei.text);
        ferao ei;
    }

    kaj toLowerCase() {
        ei.text = chotoHater(ei.text);
        ferao ei;
    }

    kaj trim() {
        ei.text = chhanto(ei.text);
        ferao ei;
    }

    kaj reverse() {
        dhoro chars = [];
        ghuriye (dhoro i = dorghyo(ei.text) - 1; i >= 0; i = i - 1) {
            dhokao(chars, ei.text[i]);
        }
        ei.text = joro(chars, "");
        ferao ei;
    }

    kaj process() {
        // Call multiple methods internally
        ei.trim();
        ei.toLowerCase();
        ferao ei;
    }

    kaj get() {
        ferao ei.text;
    }
}

dhoro helper = notun StringHelper("  HELLO WORLD  ");
dekho(helper.process().get());  // "hello world"

dhoro helper2 = notun StringHelper("Hello");
dekho(helper2.reverse().toUpperCase().get());  // "OLLEH"`}
      />

      <h2>Getter and Setter Pattern</h2>

      <CodeBlock
        code={`sreni Temperature {
    shuru(celsius) {
        ei.celsius = celsius;
    }

    // Getter methods
    kaj getCelsius() {
        ferao ei.celsius;
    }

    kaj getFahrenheit() {
        ferao (ei.celsius * 9/5) + 32;
    }

    kaj getKelvin() {
        ferao ei.celsius + 273.15;
    }

    // Setter methods
    kaj setCelsius(value) {
        ei.celsius = value;
    }

    kaj setFahrenheit(value) {
        ei.celsius = (value - 32) * 5/9;
    }

    kaj setKelvin(value) {
        ei.celsius = value - 273.15;
    }

    kaj display() {
        dekho("Temperature:");
        dekho(" - Celsius:", ei.getCelsius());
        dekho(" - Fahrenheit:", ei.getFahrenheit());
        dekho(" - Kelvin:", ei.getKelvin());
    }
}

dhoro temp = notun Temperature(25);
temp.display();

temp.setFahrenheit(98.6);  // Body temperature
temp.display();`}
      />

      <h2>Static-Like Behavior</h2>

      <p>
        BanglaCode doesn&apos;t have static methods, but you can simulate them with module functions:
      </p>

      <CodeBlock
        code={`// Math utility class
sreni MathUtils {
    shuru() {}

    kaj square(n) {
        ferao n * n;
    }

    kaj cube(n) {
        ferao n * n * n;
    }

    kaj factorial(n) {
        jodi (n <= 1) {
            ferao 1;
        }
        ferao n * ei.factorial(n - 1);
    }

    kaj isPrime(n) {
        jodi (n < 2) {
            ferao mittha;
        }
        ghuriye (dhoro i = 2; i * i <= n; i = i + 1) {
            jodi (n % i == 0) {
                ferao mittha;
            }
        }
        ferao sotti;
    }
}

dhoro math = notun MathUtils();
dekho(math.square(5));     // 25
dekho(math.factorial(5));  // 120
dekho(math.isPrime(17));   // sotti`}
      />

      <h2>Method Overloading Pattern</h2>

      <p>
        BanglaCode doesn&apos;t support method overloading, but you can use parameter checking:
      </p>

      <CodeBlock
        code={`sreni Logger {
    shuru(prefix) {
        ei.prefix = prefix;
    }

    kaj log(message, level) {
        // Default level to "INFO" if not provided
        jodi (level == khali) {
            level = "INFO";
        }

        dekho("[" + level + "]", ei.prefix + ":", message);
    }

    kaj info(message) {
        ei.log(message, "INFO");
    }

    kaj warn(message) {
        ei.log(message, "WARN");
    }

    kaj error(message) {
        ei.log(message, "ERROR");
    }
}

dhoro logger = notun Logger("App");
logger.info("Application started");
logger.warn("Low memory");
logger.error("Connection failed");`}
      />

      <h2>Builder Pattern</h2>

      <CodeBlock
        code={`sreni QueryBuilder {
    shuru(table) {
        ei.table = table;
        ei.columns = "*";
        ei.whereClause = "";
        ei.orderClause = "";
        ei.limitValue = khali;
    }

    kaj select(columns) {
        ei.columns = joro(columns, ", ");
        ferao ei;
    }

    kaj where(condition) {
        ei.whereClause = " WHERE " + condition;
        ferao ei;
    }

    kaj orderBy(column, direction) {
        ei.orderClause = " ORDER BY " + column + " " + direction;
        ferao ei;
    }

    kaj limit(n) {
        ei.limitValue = n;
        ferao ei;
    }

    kaj build() {
        dhoro query = "SELECT " + ei.columns + " FROM " + ei.table;
        query = query + ei.whereClause;
        query = query + ei.orderClause;
        jodi (ei.limitValue != khali) {
            query = query + " LIMIT " + lipi(ei.limitValue);
        }
        ferao query;
    }
}

dhoro query = notun QueryBuilder("users")
    .select(["id", "name", "email"])
    .where("active = 1")
    .orderBy("name", "ASC")
    .limit(10)
    .build();

dekho(query);
// SELECT id, name, email FROM users WHERE active = 1 ORDER BY name ASC LIMIT 10`}
      />

      <h2>Object Utility Methods</h2>

      <p>
        BanglaCode provides built-in utility methods for working with objects (maps):
      </p>

      <h3>maan - Get Object Values</h3>

      <p>
        Returns an array containing all values from an object:
      </p>

      <CodeBlock
        code={`dhoro person = {
    "naam": "Rahim",
    "age": 25,
    "city": "Dhaka"
};

dhoro values = maan(person);
dekho(values);  // ["Rahim", 25, "Dhaka"]

// Use values in higher-order functions
dhoro doubled = manchitro(values, kaj(val) {
    jodi (dhoron(val) == "NUMBER") {
        ferao val * 2;
    }
    ferao val;
});
dekho(doubled);  // ["Rahim", 50, "Dhaka"]`}
      />

      <h3>jora - Get Object Entries</h3>

      <p>
        Returns an array of [key, value] pairs from an object:
      </p>

      <CodeBlock
        code={`dhoro settings = {
    "theme": "dark",
    "notifications": sotti,
    "fontSize": 14
};

dhoro entries = jora(settings);
dekho(entries);
// [["theme", "dark"], ["notifications", sotti], ["fontSize", 14]]

// Iterate over entries
proti(entries, kaj(entry) {
    dekho(entry[0], ":", entry[1]);
});
// Output:
// theme : dark
// notifications : sotti
// fontSize : 14

// Filter entries and rebuild object
dhoro filtered = chhanno(entries, kaj(entry) {
    ferao entry[0] != "fontSize";  // Exclude fontSize
});
dekho(filtered);
// [["theme", "dark"], ["notifications", sotti]]`}
      />

      <h3>mishra - Merge Objects</h3>

      <p>
        Merges one or more source objects into a target object. Modifies the target in-place.
      </p>

      <CodeBlock
        code={`// Basic merge
dhoro obj1 = {"a": 1, "b": 2};
dhoro obj2 = {"c": 3};
dhoro obj3 = {"d": 4};

mishra(obj1, obj2, obj3);
dekho(obj1);  // {"a": 1, "b": 2, "c": 3, "d": 4}

// Later values override earlier ones
dhoro target = {"name": "Rahim", "age": 25};
dhoro update = {"age": 26, "city": "Dhaka"};

mishra(target, update);
dekho(target);  // {"name": "Rahim", "age": 26, "city": "Dhaka"}

// Practical: Merge default config with user config
dhoro defaultConfig = {
    "theme": "light",
    "fontSize": 12,
    "language": "bn"
};

dhoro userConfig = {
    "theme": "dark",
    "fontSize": 14
};

dhoro finalConfig = {"theme": "light", "fontSize": 12, "language": "bn"};
mishra(finalConfig, userConfig);
dekho(finalConfig);
// {"theme": "dark", "fontSize": 14, "language": "bn"}`}
      />

      <h2>Common Object Patterns</h2>

      <h3>Transform Object Values</h3>

      <CodeBlock
        code={`dhoro prices = {
    "apple": 50,
    "banana": 30,
    "mango": 80
};

// Apply discount to all prices
dhoro values = maan(prices);
dhoro discounted = manchitro(values, kaj(price) {
    ferao price * 0.9;  // 10% discount
});

dekho(discounted);  // [45, 27, 72]`}
      />

      <h3>Filter Object Properties</h3>

      <CodeBlock
        code={`dhoro user = {
    "name": "Rahim",
    "email": "rahim@example.com",
    "password": "secret123",
    "phone": "01712345678",
    "_internal": "data"
};

// Remove private fields (starting with _) and sensitive fields
dhoro filtered = chhanno(jora(user), kaj(entry) {
    dhoro key = entry[0];
    ferao key != "password" ebong key[0] != "_";
});

dekho(filtered);
// [["name", "Rahim"], ["email", "rahim@example.com"], ["phone", "01712345678"]]`}
      />

      <h3>Build Object Dynamically</h3>

      <CodeBlock
        code={`// Reduce key-value pairs into object
dhoro entries = [["id", 1], ["name", "Rahim"], ["active", sotti]];

dhoro obj = sonkuchito(entries, kaj(acc, entry) {
    acc[entry[0]] = entry[1];
    ferao acc;
}, {});

dekho(obj);  // {"id": 1, "name": "Rahim", "active": sotti}`}
      />

      <DocNavigation currentPath="/docs/methods" />
    </div>
  );
}
