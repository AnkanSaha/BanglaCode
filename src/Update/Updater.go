package update

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

// Update commands for each operating system
var updateCommands = map[string]updateCommand{
	"linux": {
		shell:   "bash",
		shellArg: "-c",
		script:  "curl -fsSL https://raw.githubusercontent.com/nexoral/BanglaCode/main/Scripts/install.sh | bash",
	},
	"darwin": {
		shell:   "bash",
		shellArg: "-c",
		script:  "curl -fsSL https://raw.githubusercontent.com/nexoral/BanglaCode/main/Scripts/install.sh | bash",
	},
	"windows": {
		shell:   "powershell",
		shellArg: "-Command",
		script:  "irm https://raw.githubusercontent.com/nexoral/BanglaCode/main/Scripts/install.ps1 | iex",
	},
}

type updateCommand struct {
	shell    string
	shellArg string
	script   string
}

// Update checks if the update command was passed and executes the appropriate update script
func Updater() {
	fmt.Println("\033[1;36m╔════════════════════════════════════════════════════════╗")
	fmt.Println("║              Updating BanglaCode...                    ║")
	fmt.Println("╚════════════════════════════════════════════════════════╝\033[0m")
	fmt.Println()

	// Detect operating system
	currentOS := runtime.GOOS
	cmd, exists := updateCommands[currentOS]

	if !exists {
		fmt.Fprintf(os.Stderr, "\033[31mError: Unsupported operating system '%s'\033[0m\n", currentOS)
		fmt.Fprintf(os.Stderr, "Supported systems: Linux, macOS (darwin), Windows\n")
		os.Exit(1)
	}

	fmt.Printf("Detected OS: \033[1;32m%s\033[0m\n", currentOS)
	fmt.Printf("Running update command...\n\n")

	// Execute the appropriate update command
	command := exec.Command(cmd.shell, cmd.shellArg, cmd.script)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	if err := command.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "\n\033[31mError: Update failed: %v\033[0m\n", err)
		os.Exit(1)
	}

	fmt.Println("\n\033[1;32m✓ Update completed successfully!\033[0m")
}

// checkUpdateCommand checks if the "update" argument is present in the command-line args
func CheckUpdateCommand(args []string, target string) bool {
	for _, arg := range args {
		if strings.ToLower(arg) == target {
			return true
		}
	}
	return false
}
