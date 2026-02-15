package main

import (
	"BanglaCode/src/evaluator"
	"BanglaCode/src/lexer"
	"BanglaCode/src/object"
	"BanglaCode/src/parser"
	"BanglaCode/src/repl"
	"fmt"
	"os"
	"os/user"
	"path/filepath"
)

func main() {
	// Check for command line arguments
	if len(os.Args) == 1 {
		// No arguments - start REPL
		user, err := user.Current()
		if err != nil {
			panic(err)
		}
		fmt.Printf("Namaskar %s! Welcome to BanglaCode!\n", user.Username)
		repl.Start(os.Stdin, os.Stdout)
		return
	}

	// Check for flags
	if os.Args[1] == "--help" || os.Args[1] == "-h" {
		printHelp()
		return
	}

	if os.Args[1] == "--version" || os.Args[1] == "-v" {
		printVersion()
		return
	}

	// Execute file
	filename := os.Args[1]
	runFile(filename)
}

func printHelp() {
	fmt.Println("\033[1;36m╔══════════════════════════════════════════════════════════════════╗")
	fmt.Println("║                           BanglaCode                             ║")
	fmt.Println("║           A Programming Language in Bengali (Banglish)           ║")
	fmt.Println("╠══════════════════════════════════════════════════════════════════╣\033[0m")
	fmt.Println("\033[1;36m║\033[0m  👨‍💻 \033[1mAuthor:\033[0m  \033[1;35mAnkan Saha\033[0m                                          \033[1;36m║\033[0m")
	fmt.Println("\033[1;36m║\033[0m  🌍 \033[1mFrom:\033[0m    \033[1;37mWest Bengal, India\033[0m                                   \033[1;36m║\033[0m")
	fmt.Println("\033[1;36m╠══════════════════════════════════════════════════════════════════╣\033[0m")
	fmt.Println("")
	fmt.Println("\033[1;33m▸ Usage:\033[0m")
	fmt.Println("  \033[1;32mbanglacode\033[0m                  Start interactive REPL")
	fmt.Println("  \033[1;32mbanglacode <file>\033[0m           Execute a BanglaCode file")
	fmt.Println("  \033[1;32mbanglacode --help, -h\033[0m       Show this help message")
	fmt.Println("  \033[1;32mbanglacode --version, -v\033[0m    Show version information")
	fmt.Println("")
	fmt.Println("\033[1;33m▸ Supported File Extensions:\033[0m")
	fmt.Println("  \033[1;36m.bang\033[0m   \033[1;36m.bangla\033[0m   \033[1;36m.bong\033[0m")
	fmt.Println("")
	fmt.Println("\033[1;33m▸ Examples:\033[0m")
	fmt.Println("  \033[0;34m$\033[0m banglacode                  \033[2m# Start REPL\033[0m")
	fmt.Println("  \033[0;34m$\033[0m banglacode hello.bang       \033[2m# Run hello.bang file\033[0m")
	fmt.Println("  \033[0;34m$\033[0m banglacode app.bangla       \033[2m# Run app.bangla file\033[0m")
	fmt.Println("  \033[0;34m$\033[0m banglacode server.bong      \033[2m# Run server.bong file\033[0m")
	fmt.Println("")
	fmt.Println("\033[1;36m╚══════════════════════════════════════════════════════════════════╝\033[0m")
	fmt.Println("  📄 For more information, see \033[1;35mSYNTAX.md\033[0m")
	fmt.Println("  🔗 GitHub: \033[1;34mhttps://github.com/nexoral/BanglaCode\033[0m")
}

func printVersion() {
	fmt.Println("\033[1;36m╔════════════════════════════════════════════════════════╗")
	fmt.Println("║               BanglaCode v7.0.0                        ║")
	fmt.Println("║     A Programming Language in Bengali (Banglish)      ║")
	fmt.Println("╠════════════════════════════════════════════════════════╣\033[0m")
	fmt.Println("\033[1;36m║\033[0m  📦 \033[1mVersion:\033[0m      \033[1;32m7.0.0\033[0m                                 \033[1;36m║\033[0m")
	fmt.Println("\033[1;36m║\033[0m  👨‍💻 \033[1mAuthor:\033[0m       \033[1;35mAnkan Saha\033[0m                            \033[1;36m║\033[0m")
	fmt.Println("\033[1;36m║\033[0m  🌍 \033[1mFrom:\033[0m         \033[1;37mWest Bengal, India\033[0m                    \033[1;36m║\033[0m")
	fmt.Println("\033[1;36m║\033[0m  🔗 \033[1mGitHub:\033[0m       \033[1;34mhttps://github.com/nexoral/BanglaCode\033[0m \033[1;36m║\033[0m")
	fmt.Println("\033[1;36m║\033[0m  📄 \033[1mLicense:\033[0m      \033[1;37mMIT License\033[0m                           \033[1;36m║\033[0m")
	fmt.Println("\033[1;36m╚════════════════════════════════════════════════════════╝\033[0m")
}

func runFile(filename string) {
	// Validate file extension (warning only, not enforced)
	ext := filepath.Ext(filename)
	if ext != ".bang" && ext != ".bangla" && ext != ".bong" {
		fmt.Fprintf(os.Stderr, "\033[33mWarning: '%s' does not have a standard BanglaCode extension (.bang, .bangla, .bong)\033[0m\n", filename)
	}

	// Read file
	content, err := os.ReadFile(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}

	// Set current directory for module imports
	absPath, _ := filepath.Abs(filename)
	evaluator.SetCurrentDir(filepath.Dir(absPath))

	// Create environment
	env := object.NewEnvironment()

	// Lex
	l := lexer.New(string(content))

	// Parse
	p := parser.New(l)
	program := p.ParseProgram()

	if len(p.Errors()) != 0 {
		fmt.Fprintln(os.Stderr, "\033[31mParser errors:\033[0m")
		for _, msg := range p.Errors() {
			fmt.Fprintf(os.Stderr, "\t%s\n", msg)
		}
		os.Exit(1)
	}

	// Evaluate
	result := evaluator.Eval(program, env)

	if result != nil && result.Type() == object.ERROR_OBJ {
		fmt.Fprintf(os.Stderr, "\033[31m%s\033[0m\n", result.Inspect())
		os.Exit(1)
	}
}
