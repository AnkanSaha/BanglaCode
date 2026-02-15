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

const Version = "7.0.3"

const PROMPT = "\033[1;33m>> \033[0m"

// Color codes
const (
	Reset   = "\033[0m"
	Bold    = "\033[1m"
	Dim     = "\033[2m"
	Cyan    = "\033[1;36m"
	Green   = "\033[1;32m"
	Yellow  = "\033[1;33m"
	Blue    = "\033[1;34m"
	Magenta = "\033[1;35m"
	Red     = "\033[1;31m"
	White   = "\033[1;37m"
)

func printBanner(out io.Writer) {
	width := 70
	line := strings.Repeat("═", width-2)

	fmt.Fprintln(out)
	fmt.Fprintf(out, "%s╔%s╗%s\n", Cyan, line, Reset)
	fmt.Fprintf(out, "%s║%s", Cyan, Reset)
	fmt.Fprintf(out, "  %s ____                   _        _____          _       %s", Yellow, Reset)
	fmt.Fprintf(out, "       %s║%s\n", Cyan, Reset)
	fmt.Fprintf(out, "%s║%s", Cyan, Reset)
	fmt.Fprintf(out, "  %s|  _ \\                 | |      / ____|        | |      %s", Yellow, Reset)
	fmt.Fprintf(out, "       %s║%s\n", Cyan, Reset)
	fmt.Fprintf(out, "%s║%s", Cyan, Reset)
	fmt.Fprintf(out, "  %s| |_) | __ _ _ __   __ | | __ _| |     ___   __| | ___  %s", Yellow, Reset)
	fmt.Fprintf(out, "       %s║%s\n", Cyan, Reset)
	fmt.Fprintf(out, "%s║%s", Cyan, Reset)
	fmt.Fprintf(out, "  %s|  _ < / _' | '_ \\ / _' |/ _' | |    / _ \\ / _' |/ _ \\ %s", Yellow, Reset)
	fmt.Fprintf(out, "       %s║%s\n", Cyan, Reset)
	fmt.Fprintf(out, "%s║%s", Cyan, Reset)
	fmt.Fprintf(out, "  %s| |_) | (_| | | | | (_| | (_| | |___| (_) | (_| |  __/ %s", Yellow, Reset)
	fmt.Fprintf(out, "       %s║%s\n", Cyan, Reset)
	fmt.Fprintf(out, "%s║%s", Cyan, Reset)
	fmt.Fprintf(out, "  %s|____/ \\__,_|_| |_|\\__, |\\__,_|\\_____\\___/ \\__,_|\\___| %s", Yellow, Reset)
	fmt.Fprintf(out, "       %s║%s\n", Cyan, Reset)
	fmt.Fprintf(out, "%s║%s", Cyan, Reset)
	fmt.Fprintf(out, "  %s                    __/ |                              %s", Yellow, Reset)
	fmt.Fprintf(out, "       %s║%s\n", Cyan, Reset)
	fmt.Fprintf(out, "%s║%s", Cyan, Reset)
	fmt.Fprintf(out, "  %s                   |___/                               %s", Yellow, Reset)
	fmt.Fprintf(out, "       %s║%s\n", Cyan, Reset)
	fmt.Fprintf(out, "%s╠%s╣%s\n", Cyan, line, Reset)

	// Tagline
	tagline := "A PROGRAMMING LANGUAGE IN BENGALI (BANGLISH)"
	padding := (width - 2 - len(tagline)) / 2
	fmt.Fprintf(out, "%s║%s%s%s%s%s%s║%s\n", Cyan, Reset, strings.Repeat(" ", padding), White, tagline, strings.Repeat(" ", width-2-padding-len(tagline)), Cyan, Reset)

	fmt.Fprintf(out, "%s╠%s╣%s\n", Cyan, line, Reset)

	// Info section
	fmt.Fprintf(out, "%s║%s  📦 %sVersion:%s      %s%-55s%s║%s\n", Cyan, Reset, Bold, Reset, Green, Version, Cyan, Reset)
	fmt.Fprintf(out, "%s║%s  👨‍💻 %sAuthor:%s       %s%-55s%s║%s\n", Cyan, Reset, Bold, Reset, Magenta, "Ankan Saha", Cyan, Reset)
	fmt.Fprintf(out, "%s║%s  🌍 %sFrom:%s         %s%-55s%s║%s\n", Cyan, Reset, Bold, Reset, White, "West Bengal, India", Cyan, Reset)
	fmt.Fprintf(out, "%s║%s  🔗 %sGitHub:%s       %s%-55s%s║%s\n", Cyan, Reset, Bold, Reset, Blue, "https://github.com/nexoral/BanglaCode", Cyan, Reset)
	fmt.Fprintf(out, "%s║%s  📄 %sLicense:%s      %s%-55s%s║%s\n", Cyan, Reset, Bold, Reset, White, "MIT License", Cyan, Reset)

	fmt.Fprintf(out, "%s╠%s╣%s\n", Cyan, line, Reset)

	// Commands section
	fmt.Fprintf(out, "%s║%s  %ssahajjo%s   │ Show help & keywords                              %s║%s\n", Cyan, Reset, Blue, Reset, Cyan, Reset)
	fmt.Fprintf(out, "%s║%s  %smochho%s    │ Clear screen                                      %s║%s\n", Cyan, Reset, Blue, Reset, Cyan, Reset)
	fmt.Fprintf(out, "%s║%s  %sbaire%s     │ Exit REPL                                         %s║%s\n", Cyan, Reset, Blue, Reset, Cyan, Reset)

	fmt.Fprintf(out, "%s╚%s╝%s\n", Cyan, line, Reset)
	fmt.Fprintln(out)
}

const HELP = `
` + Cyan + `╔════════════════════════════════════════════════════════════════════╗
║                        BanglaCode Help                             ║
╚════════════════════════════════════════════════════════════════════╝` + Reset + `

` + Yellow + `▸ Keywords:` + Reset + `
  ` + Green + `dhoro` + Reset + `        variable declaration (let/var)
  ` + Green + `jodi` + Reset + `         if condition
  ` + Green + `nahole` + Reset + `       else
  ` + Green + `jotokkhon` + Reset + `    while loop
  ` + Green + `ghuriye` + Reset + `      for loop
  ` + Green + `kaj` + Reset + `          function definition
  ` + Green + `ferao` + Reset + `        return statement
  ` + Green + `sreni` + Reset + `        class definition
  ` + Green + `notun` + Reset + `        new instance
  ` + Green + `shuru` + Reset + `        constructor
  ` + Green + `sotti` + Reset + `        true
  ` + Green + `mittha` + Reset + `       false
  ` + Green + `khali` + Reset + `        null
  ` + Green + `ebong` + Reset + `        logical and (&&)
  ` + Green + `ba` + Reset + `           logical or (||)
  ` + Green + `na` + Reset + `           logical not (!)
  ` + Green + `thamo` + Reset + `        break
  ` + Green + `chharo` + Reset + `       continue
  ` + Green + `ano` + Reset + `          import module
  ` + Green + `pathao` + Reset + `       export
  ` + Green + `chesta` + Reset + `       try block
  ` + Green + `dhoro_bhul` + Reset + `   catch block
  ` + Green + `felo` + Reset + `         throw error

` + Yellow + `▸ Built-in Functions:` + Reset + `
  ` + Blue + `dekho(...)` + Reset + `       print values
  ` + Blue + `dorghyo(x)` + Reset + `       get length
  ` + Blue + `dhokao(arr,v)` + Reset + `    push to array
  ` + Blue + `berKoro(arr)` + Reset + `     pop from array
  ` + Blue + `dhoron(x)` + Reset + `        get type
  ` + Blue + `lipi(x)` + Reset + `          to string
  ` + Blue + `sonkha(x)` + Reset + `        to number

` + Yellow + `▸ String Functions:` + Reset + `
  ` + Blue + `boroHater(s)` + Reset + `     uppercase
  ` + Blue + `chotoHater(s)` + Reset + `    lowercase
  ` + Blue + `chhanto(s)` + Reset + `       trim
  ` + Blue + `bhag(s,sep)` + Reset + `      split
  ` + Blue + `joro(arr,sep)` + Reset + `    join
  ` + Blue + `khojo(s,sub)` + Reset + `     indexOf

` + Yellow + `▸ Math Functions:` + Reset + `
  ` + Blue + `borgomul(x)` + Reset + `      sqrt
  ` + Blue + `ghat(b,e)` + Reset + `        power
  ` + Blue + `niche(x)` + Reset + `         floor
  ` + Blue + `upore(x)` + Reset + `         ceil
  ` + Blue + `kache(x)` + Reset + `         round
  ` + Blue + `lotto()` + Reset + `          random

` + Yellow + `▸ Utility Functions:` + Reset + `
  ` + Blue + `somoy()` + Reset + `          timestamp
  ` + Blue + `ghum(ms)` + Reset + `         sleep
  ` + Blue + `nao(prompt)` + Reset + `      user input
  ` + Blue + `poro(path)` + Reset + `       read file
  ` + Blue + `lekho(p,c)` + Reset + `       write file
  ` + Blue + `server_chalu` + Reset + `     HTTP server

` + Cyan + `╔════════════════════════════════════════════════════════════════════╗
║                            Example                                 ║
╚════════════════════════════════════════════════════════════════════╝` + Reset + `
  ` + Yellow + `>>` + Reset + ` dhoro naam = "Ankan"
  ` + Yellow + `>>` + Reset + ` dekho("Namaskar", naam)
  ` + Green + `Namaskar Ankan` + Reset + `

  ` + Yellow + `>>` + Reset + ` sreni Manush { kaj shuru(naam) { ei.naam = naam; } }
  ` + Yellow + `>>` + Reset + ` dhoro m = notun Manush("Rana")
  ` + Yellow + `>>` + Reset + ` dekho(m.naam)
  ` + Green + `Rana` + Reset + `
`

// Start begins the REPL
func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()

	printBanner(out)

	for {
		fmt.Fprint(out, PROMPT)

		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()

		// Handle special commands (Banglish and English aliases)
		if line == "baire" || line == "exit" || line == "quit" {
			fmt.Fprintln(out, Green)
			fmt.Fprintln(out, "╔════════════════════════════════════════════╗")
			fmt.Fprintln(out, "║  Dhonnobad! Abar dekha hobe!               ║")
			fmt.Fprintln(out, "║  Thank you! See you again!                 ║")
			fmt.Fprintln(out, "╚════════════════════════════════════════════╝")
			fmt.Fprintln(out, Reset)
			return
		}

		if line == "sahajjo" || line == "help" {
			fmt.Fprint(out, HELP)
			continue
		}

		if line == "mochho" || line == "clear" || line == "cls" {
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
				io.WriteString(out, Red)
				io.WriteString(out, evaluated.Inspect())
				io.WriteString(out, Reset+"\n")
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
		fmt.Print(Yellow + ".. " + Reset)
	}

	return builder.String()
}

func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, Red)
	io.WriteString(out, "╔════════════════════════════════════════════╗\n")
	io.WriteString(out, "║  Bhul! Parser Errors                       ║\n")
	io.WriteString(out, "╚════════════════════════════════════════════╝\n")
	io.WriteString(out, "\033[0;31m") // Regular Red
	for _, msg := range errors {
		io.WriteString(out, "  ▸ "+msg+"\n")
	}
	io.WriteString(out, Reset)
}
