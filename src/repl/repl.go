package repl

import (
	"BanglaCode/src/evaluator"
	"BanglaCode/src/lexer"
	"BanglaCode/src/object"
	"BanglaCode/src/parser"
	"bufio"
	"fmt"
	"io"
	"strings"
)

const PROMPT = ">> "

const BANNER = `
 ____                   _        _____          _
|  _ \                 | |      / ____|        | |
| |_) | __ _ _ __   __ | | __ _| |     ___   __| | ___
|  _ < / _' | '_ \ / _' |/ _' | |    / _ \ / _' |/ _ \
| |_) | (_| | | | | (_| | (_| | |___| (_) | (_| |  __/
|____/ \__,_|_| |_|\__, |\__,_|\_____\___/ \__,_|\___|
                    __/ |
                   |___/

Welcome to BanglaCode - Bengali Programming Language
Created by Ankan from West Bengal, India
Type 'exit' or Ctrl+C to quit
Type 'help' for a list of keywords
Type 'clear' to clear screen
`

const HELP = `
BanglaCode Keywords (কীওয়ার্ড):
  dhoro        - variable declaration (let/var)
  jodi         - if condition
  nahole       - else
  jotokkhon    - while loop
  ghuriye      - for loop
  kaj          - function definition
  ferao        - return statement
  class        - class definition
  notun        - new instance
  sotti        - true
  mittha       - false
  khali        - null
  ebong        - logical and (&&)
  ba           - logical or (||)
  na           - logical not (!)
  thamo        - break
  chharo       - continue
  ano          - import module
  pathao       - export function/class
  chesta       - try block
  dhoro_bhul   - catch block
  shesh        - finally block
  felo         - throw error

Built-in Functions (সাধারণ):
  dekho(...)       - দেখো - print values
  dorghyo(x)       - দৈর্ঘ্য - get length
  dhokao(arr, val) - ঢোকাও - add to array
  berKoro(arr)     - বের করো - remove last
  chabi(map)       - চাবি - get keys
  dhoron(x)        - ধরন - get type
  lipi(x)          - লিপি - to string
  sonkha(x)        - সংখ্যা - to number

String Functions (লেখা):
  boroHater(s)     - বড় হাতের - uppercase
  chotoHater(s)    - ছোট হাতের - lowercase
  chhanto(s)       - ছাঁটো - trim
  bhag(s, sep)     - ভাগ - split
  joro(arr, sep)   - জোড়ো - join
  khojo(s, sub)    - খোঁজো - indexOf
  angsho(s, start, end) - অংশ - substring
  bodlo(s, old, new)    - বদলো - replace

Array Functions (তালিকা):
  kato(arr, start, end) - কাটো - slice
  ulto(arr)        - উল্টো - reverse
  saja(arr)        - সাজা - sort
  ache(arr, val)   - আছে - includes

Math Functions (গণিত):
  borgomul(x)      - বর্গমূল - sqrt
  ghat(base, exp)  - ঘাত - pow
  niche(x)         - নিচে - floor
  upore(x)         - উপরে - ceil
  kache(x)         - কাছে - round
  niratek(x)       - নিরপেক্ষ - abs
  choto(...)       - ছোট - min
  boro(...)        - বড় - max
  lotto()          - লটো - random

Utility Functions (সহায়ক):
  somoy()          - সময় - timestamp
  ghum(ms)         - ঘুম - sleep
  nao(prompt)      - নাও - input
  bondho(code)     - বন্ধ - exit
  poro(path)       - পড়ো - read file
  lekho(path, content) - লেখো - write file
  server_chalu(port, handler) - সার্ভার চালু - HTTP server
  anun(url)        - আনুন - HTTP GET

REPL Commands:
  help         - show this help
  clear/cls    - clear screen
  exit/quit    - exit REPL

Example:
  >> dhoro naam = "Ankan"
  >> dekho("Namaskar", naam)
  Namaskar Ankan
`

// Start begins the REPL
func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()

	fmt.Fprint(out, BANNER)

	for {
		fmt.Fprint(out, PROMPT)

		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()

		// Handle special commands
		if line == "exit" || line == "quit" {
			fmt.Fprintln(out, "Dhonnobad! (Thank you!)")
			return
		}

		if line == "help" {
			fmt.Fprint(out, HELP)
			continue
		}

		if line == "clear" || line == "cls" {
			// Clear screen using ANSI escape codes
			fmt.Fprint(out, "\033[2J\033[H")
			continue
		}

		if line == "" {
			continue
		}

		// Handle multi-line input
		input := line
		if needsMoreInput(line) {
			input = readMultiLine(scanner, line)
		}

		l := lexer.New(input)
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}

		evaluated := evaluator.Eval(program, env)
		if evaluated != nil {
			if evaluated.Type() != object.NULL_OBJ && evaluated.Type() != object.ERROR_OBJ {
				io.WriteString(out, evaluated.Inspect())
				io.WriteString(out, "\n")
			} else if evaluated.Type() == object.ERROR_OBJ {
				io.WriteString(out, "\033[31m") // Red color
				io.WriteString(out, evaluated.Inspect())
				io.WriteString(out, "\033[0m\n") // Reset color
			}
		}
	}
}

func needsMoreInput(line string) bool {
	openBraces := strings.Count(line, "{")
	closeBraces := strings.Count(line, "}")
	openParens := strings.Count(line, "(")
	closeParens := strings.Count(line, ")")
	openBrackets := strings.Count(line, "[")
	closeBrackets := strings.Count(line, "]")

	return openBraces > closeBraces || openParens > closeParens || openBrackets > closeBrackets
}

func readMultiLine(scanner *bufio.Scanner, initial string) string {
	var builder strings.Builder
	builder.WriteString(initial)
	builder.WriteString("\n")

	for scanner.Scan() {
		line := scanner.Text()
		builder.WriteString(line)
		builder.WriteString("\n")

		if !needsMoreInput(builder.String()) {
			break
		}
		fmt.Print(".. ")
	}

	return builder.String()
}

func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, "\033[31m") // Red color
	io.WriteString(out, "Oops! Parser errors:\n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
	io.WriteString(out, "\033[0m") // Reset color
}
