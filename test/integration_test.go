package test

import (
	"BanglaCode/src/evaluator"
	"BanglaCode/src/lexer"
	"BanglaCode/src/object"
	"BanglaCode/src/parser"
	"testing"
)

func runProgram(input string) object.Object {
	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()
	env := object.NewEnvironment()
	return evaluator.Eval(program, env)
}

// Test a complete factorial program
func TestIntegrationFactorial(t *testing.T) {
	input := `
	kaj factorial(n) {
		jodi (n <= 1) {
			ferao 1;
		}
		ferao n * factorial(n - 1);
	}
	factorial(10);
	`

	result := runProgram(input)
	num, ok := result.(*object.Number)
	if !ok {
		t.Fatalf("expected Number, got %T", result)
	}
	if num.Value != 3628800 {
		t.Errorf("expected 3628800, got %f", num.Value)
	}
}

// Test fibonacci sequence
func TestIntegrationFibonacci(t *testing.T) {
	input := `
	kaj fib(n) {
		jodi (n <= 1) {
			ferao n;
		}
		ferao fib(n - 1) + fib(n - 2);
	}
	fib(10);
	`

	result := runProgram(input)
	num, ok := result.(*object.Number)
	if !ok {
		t.Fatalf("expected Number, got %T", result)
	}
	if num.Value != 55 {
		t.Errorf("expected 55, got %f", num.Value)
	}
}

// Test array manipulation using array mutation for scope compatibility
func TestIntegrationArrayManipulation(t *testing.T) {
	input := `
	dhoro arr = [1, 2, 3, 4, 5];
	dhoro result = [0];
	ghuriye (dhoro i = 0; i < dorghyo(arr); i = i + 1) {
		result[0] = result[0] + arr[i];
	}
	result[0];
	`

	result := runProgram(input)
	num, ok := result.(*object.Number)
	if !ok {
		t.Fatalf("expected Number, got %T", result)
	}
	if num.Value != 15 {
		t.Errorf("expected 15, got %f", num.Value)
	}
}

// Test map/object manipulation
func TestIntegrationMapManipulation(t *testing.T) {
	input := `
	dhoro person = {
		naam: "Ankan",
		boyes: 25,
		city: "Kolkata"
	};
	person.boyes = person.boyes + 1;
	person.boyes;
	`

	result := runProgram(input)
	num, ok := result.(*object.Number)
	if !ok {
		t.Fatalf("expected Number, got %T", result)
	}
	if num.Value != 26 {
		t.Errorf("expected 26, got %f", num.Value)
	}
}

// Test nested loops with break using array mutation for scope compatibility
func TestIntegrationNestedLoopsWithBreak(t *testing.T) {
	input := `
	dhoro count = [0];
	ghuriye (dhoro i = 0; i < 10; i = i + 1) {
		ghuriye (dhoro j = 0; j < 10; j = j + 1) {
			count[0] = count[0] + 1;
			jodi (j >= 4) {
				thamo;
			}
		}
	}
	count[0];
	`

	result := runProgram(input)
	num, ok := result.(*object.Number)
	if !ok {
		t.Fatalf("expected Number, got %T", result)
	}
	// Each outer iteration runs inner loop 5 times (j=0,1,2,3,4 then break)
	// 10 * 5 = 50
	if num.Value != 50 {
		t.Errorf("expected 50, got %f", num.Value)
	}
}

// Test class with multiple methods
func TestIntegrationClassWithMethods(t *testing.T) {
	input := `
	sreni Calculator {
		shuru(initial) {
			ei.value = initial;
		}
		kaj add(x) {
			ei.value = ei.value + x;
			ferao ei;
		}
		kaj subtract(x) {
			ei.value = ei.value - x;
			ferao ei;
		}
		kaj multiply(x) {
			ei.value = ei.value * x;
			ferao ei;
		}
		kaj getValue() {
			ferao ei.value;
		}
	}

	dhoro calc = notun Calculator(10);
	calc.add(5).multiply(2).subtract(10).getValue();
	`

	result := runProgram(input)
	num, ok := result.(*object.Number)
	if !ok {
		t.Fatalf("expected Number, got %T", result)
	}
	// (10 + 5) * 2 - 10 = 30 - 10 = 20
	if num.Value != 20 {
		t.Errorf("expected 20, got %f", num.Value)
	}
}

// Test higher-order functions
func TestIntegrationHigherOrderFunctions(t *testing.T) {
	input := `
	kaj map(arr, f) {
		dhoro result = [];
		ghuriye (dhoro i = 0; i < dorghyo(arr); i = i + 1) {
			dhokao(result, f(arr[i]));
		}
		ferao result;
	}

	kaj double(x) {
		ferao x * 2;
	}

	dhoro nums = [1, 2, 3, 4, 5];
	dhoro doubled = map(nums, double);
	doubled[4];
	`

	result := runProgram(input)
	num, ok := result.(*object.Number)
	if !ok {
		t.Fatalf("expected Number, got %T", result)
	}
	if num.Value != 10 {
		t.Errorf("expected 10, got %f", num.Value)
	}
}

// Test closures using array mutation for scope compatibility
func TestIntegrationClosuresInLoops(t *testing.T) {
	input := `
	kaj makeCounter() {
		dhoro count = [0];
		ferao kaj() {
			count[0] = count[0] + 1;
			ferao count[0];
		};
	}

	dhoro counter = makeCounter();
	counter();
	counter();
	counter();
	`

	result := runProgram(input)
	num, ok := result.(*object.Number)
	if !ok {
		t.Fatalf("expected Number, got %T", result)
	}
	if num.Value != 3 {
		t.Errorf("expected 3, got %f", num.Value)
	}
}

// Test try-catch-finally flow using array mutation for scope compatibility
func TestIntegrationTryCatchFinallyFlow(t *testing.T) {
	input := `
	dhoro log = [""];

	kaj riskyOperation(shouldFail) {
		chesta {
			log[0] = log[0] + "try-";
			jodi (shouldFail) {
				felo "Operation failed";
			}
			log[0] = log[0] + "success-";
		} dhoro_bhul (e) {
			log[0] = log[0] + "catch-";
		} shesh {
			log[0] = log[0] + "finally";
		}
	}

	riskyOperation(sotti);
	log[0];
	`

	result := runProgram(input)
	str, ok := result.(*object.String)
	if !ok {
		t.Fatalf("expected String, got %T", result)
	}
	if str.Value != "try-catch-finally" {
		t.Errorf("expected 'try-catch-finally', got %q", str.Value)
	}
}

// Test string manipulation
func TestIntegrationStringManipulation(t *testing.T) {
	input := `
	dhoro text = "  Hello, World!  ";
	dhoro trimmed = chhanto(text);
	dhoro upper = boroHater(trimmed);
	dhoro replaced = bodlo(upper, "WORLD", "BANGLACODE");
	replaced;
	`

	result := runProgram(input)
	str, ok := result.(*object.String)
	if !ok {
		t.Fatalf("expected String, got %T", result)
	}
	if str.Value != "HELLO, BANGLACODE!" {
		t.Errorf("expected 'HELLO, BANGLACODE!', got %q", str.Value)
	}
}

// Test array with built-in functions
func TestIntegrationArrayBuiltins(t *testing.T) {
	input := `
	dhoro arr = [3, 1, 4, 1, 5, 9, 2, 6];
	dhoro sorted = saja(arr);
	dhoro reversed = ulto(sorted);
	dhoro sliced = kato(reversed, 0, 3);
	joro(sliced, "-");
	`

	result := runProgram(input)
	str, ok := result.(*object.String)
	if !ok {
		t.Fatalf("expected String, got %T", result)
	}
	if str.Value != "9-6-5" {
		t.Errorf("expected '9-6-5', got %q", str.Value)
	}
}

// Test multiple classes interacting
func TestIntegrationMultipleClasses(t *testing.T) {
	input := `
	sreni Point {
		shuru(x, y) {
			ei.x = x;
			ei.y = y;
		}
	}

	sreni Rectangle {
		shuru(topLeft, bottomRight) {
			ei.topLeft = topLeft;
			ei.bottomRight = bottomRight;
		}
		kaj width() {
			ferao ei.bottomRight.x - ei.topLeft.x;
		}
		kaj height() {
			ferao ei.bottomRight.y - ei.topLeft.y;
		}
		kaj area() {
			ferao ei.width() * ei.height();
		}
	}

	dhoro p1 = notun Point(0, 0);
	dhoro p2 = notun Point(10, 5);
	dhoro rect = notun Rectangle(p1, p2);
	rect.area();
	`

	result := runProgram(input)
	num, ok := result.(*object.Number)
	if !ok {
		t.Fatalf("expected Number, got %T", result)
	}
	if num.Value != 50 {
		t.Errorf("expected 50, got %f", num.Value)
	}
}

// Test complex conditional logic
func TestIntegrationComplexConditional(t *testing.T) {
	input := `
	kaj classify(n) {
		jodi (n < 0) {
			ferao "negative";
		} nahole {
			jodi (n == 0) {
				ferao "zero";
			} nahole {
				jodi (n < 10) {
					ferao "small";
				} nahole {
					jodi (n < 100) {
						ferao "medium";
					} nahole {
						ferao "large";
					}
				}
			}
		}
	}

	classify(5) + "-" + classify(50) + "-" + classify(500);
	`

	result := runProgram(input)
	str, ok := result.(*object.String)
	if !ok {
		t.Fatalf("expected String, got %T", result)
	}
	if str.Value != "small-medium-large" {
		t.Errorf("expected 'small-medium-large', got %q", str.Value)
	}
}

// Test prime number checker
func TestIntegrationPrimeChecker(t *testing.T) {
	input := `
	kaj isPrime(n) {
		jodi (n <= 1) {
			ferao mittha;
		}
		jodi (n <= 3) {
			ferao sotti;
		}
		jodi (n % 2 == 0) {
			ferao mittha;
		}
		dhoro i = 3;
		jotokkhon (i * i <= n) {
			jodi (n % i == 0) {
				ferao mittha;
			}
			i = i + 2;
		}
		ferao sotti;
	}

	dhoro primes = [];
	ghuriye (dhoro n = 2; n <= 20; n = n + 1) {
		jodi (isPrime(n)) {
			dhokao(primes, n);
		}
	}
	joro(primes, ",");
	`

	result := runProgram(input)
	str, ok := result.(*object.String)
	if !ok {
		t.Fatalf("expected String, got %T", result)
	}
	if str.Value != "2,3,5,7,11,13,17,19" {
		t.Errorf("expected '2,3,5,7,11,13,17,19', got %q", str.Value)
	}
}

// Test bubble sort implementation
func TestIntegrationBubbleSort(t *testing.T) {
	input := `
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

	dhoro nums = [64, 34, 25, 12, 22, 11, 90];
	bubbleSort(nums);
	joro(nums, ",");
	`

	result := runProgram(input)
	str, ok := result.(*object.String)
	if !ok {
		t.Fatalf("expected String, got %T", result)
	}
	if str.Value != "11,12,22,25,34,64,90" {
		t.Errorf("expected '11,12,22,25,34,64,90', got %q", str.Value)
	}
}

// Test JSON serialization and deserialization
func TestIntegrationJsonRoundTrip(t *testing.T) {
	input := `
	dhoro obj = {
		name: "BanglaCode",
		version: 1.0,
		features: ["classes", "closures", "modules"]
	};

	dhoro jsonStr = json_banao(obj);
	dhoro parsed = json_poro(jsonStr);
	parsed.name;
	`

	result := runProgram(input)
	str, ok := result.(*object.String)
	if !ok {
		t.Fatalf("expected String, got %T", result)
	}
	if str.Value != "BanglaCode" {
		t.Errorf("expected 'BanglaCode', got %q", str.Value)
	}
}

// Test deeply nested function calls
func TestIntegrationDeepNesting(t *testing.T) {
	input := `
	kaj a(x) { ferao b(x + 1); }
	kaj b(x) { ferao c(x + 1); }
	kaj c(x) { ferao d(x + 1); }
	kaj d(x) { ferao e(x + 1); }
	kaj e(x) { ferao x + 1; }

	a(0);
	`

	result := runProgram(input)
	num, ok := result.(*object.Number)
	if !ok {
		t.Fatalf("expected Number, got %T", result)
	}
	if num.Value != 5 {
		t.Errorf("expected 5, got %f", num.Value)
	}
}

// Test math calculations
func TestIntegrationMathCalculations(t *testing.T) {
	input := `
	dhoro a = borgomul(16);      // 4
	dhoro b = ghat(2, 3);        // 8
	dhoro c = niratek(-5);       // 5
	dhoro d = kache(3.7);        // 4
	dhoro e = niche(3.7);        // 3
	dhoro f = upore(3.2);        // 4
	dhoro g = choto(a, b, c);    // 4
	dhoro h = boro(a, b, c);     // 8

	a + b + c + d + e + f + g + h;
	`

	result := runProgram(input)
	num, ok := result.(*object.Number)
	if !ok {
		t.Fatalf("expected Number, got %T", result)
	}
	// 4 + 8 + 5 + 4 + 3 + 4 + 4 + 8 = 40
	if num.Value != 40 {
		t.Errorf("expected 40, got %f", num.Value)
	}
}
