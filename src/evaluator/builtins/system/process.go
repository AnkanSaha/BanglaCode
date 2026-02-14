package system

import (
	"BanglaCode/src/object"
	"bytes"
	"os"
	"os/exec"
	"runtime"
	"syscall"
	"time"
)

func init() {
	// ==================== Process Execution ====================

	// chalan (চালান) - Execute system command
	// Returns: { "output": string, "error": string, "code": number }
	registerBuiltin("chalan", func(args ...object.Object) object.Object {
		if len(args) < 1 {
			return newError("chalan requires at least 1 argument (command)")
		}

		// First argument is the command
		if args[0].Type() != object.STRING_OBJ {
			return newError("command must be STRING, got %s", args[0].Type())
		}
		cmdStr := args[0].(*object.String).Value

		// Parse command and arguments
		var cmdArgs []string
		if len(args) > 1 {
			// Additional arguments provided as array or individual strings
			if args[1].Type() == object.ARRAY_OBJ {
				arr := args[1].(*object.Array)
				cmdArgs = make([]string, len(arr.Elements))
				for i, elem := range arr.Elements {
					if elem.Type() != object.STRING_OBJ {
						return newError("command arguments must be strings")
					}
					cmdArgs[i] = elem.(*object.String).Value
				}
			} else {
				// Individual string arguments
				cmdArgs = make([]string, len(args)-1)
				for i, arg := range args[1:] {
					if arg.Type() != object.STRING_OBJ {
						return newError("command arguments must be strings")
					}
					cmdArgs[i] = arg.(*object.String).Value
				}
			}
		}

		// Execute command
		var cmd *exec.Cmd
		if len(cmdArgs) > 0 {
			cmd = exec.Command(cmdStr, cmdArgs...)
		} else {
			// If no args, try to execute as shell command
			if runtime.GOOS == "windows" {
				cmd = exec.Command("cmd", "/C", cmdStr)
			} else {
				cmd = exec.Command("sh", "-c", cmdStr)
			}
		}

		// Capture output and error
		var stdout, stderr bytes.Buffer
		cmd.Stdout = &stdout
		cmd.Stderr = &stderr

		// Execute
		err := cmd.Run()
		exitCode := 0
		if err != nil {
			if exitErr, ok := err.(*exec.ExitError); ok {
				exitCode = exitErr.ExitCode()
			} else {
				exitCode = 1
			}
		}

		// Return result as map
		result := make(map[string]object.Object)
		result["output"] = &object.String{Value: stdout.String()}
		result["error"] = &object.String{Value: stderr.String()}
		result["code"] = &object.Number{Value: float64(exitCode)}

		return &object.Map{Pairs: result}
	})

	// ==================== Process Information ====================

	// process_id (প্রসেস আইডি) - Get current process ID
	registerBuiltin("process_id", func(args ...object.Object) object.Object {
		return &object.Number{Value: float64(os.Getpid())}
	})

	// process_args (প্রসেস আর্গস) - Get command-line arguments
	registerBuiltin("process_args", func(args ...object.Object) object.Object {
		cmdArgs := os.Args
		elements := make([]object.Object, len(cmdArgs))
		for i, arg := range cmdArgs {
			elements[i] = &object.String{Value: arg}
		}
		return &object.Array{Elements: elements}
	})

	// ==================== Process Management (NEW) ====================

	// process_ghum (প্রসেস ঘুম) - Sleep for specified milliseconds
	registerBuiltin("process_ghum", func(args ...object.Object) object.Object {
		if len(args) != 1 {
			return newError("process_ghum requires 1 argument (milliseconds)")
		}
		if args[0].Type() != object.NUMBER_OBJ {
			return newError("milliseconds must be NUMBER, got %s", args[0].Type())
		}

		ms := args[0].(*object.Number).Value
		time.Sleep(time.Duration(ms) * time.Millisecond)

		return object.NULL
	})

	// process_maro (প্রসেস মারো) - Kill process by PID
	registerBuiltin("process_maro", func(args ...object.Object) object.Object {
		if len(args) != 1 {
			return newError("process_maro requires 1 argument (process ID)")
		}
		if args[0].Type() != object.NUMBER_OBJ {
			return newError("process ID must be NUMBER, got %s", args[0].Type())
		}

		pid := int(args[0].(*object.Number).Value)
		process, err := os.FindProcess(pid)
		if err != nil {
			return newError("failed to find process: %s", err.Error())
		}

		if err := process.Kill(); err != nil {
			return newError("failed to kill process: %s", err.Error())
		}

		return object.NULL
	})

	// process_signal (প্রসেস সিগন্যাল) - Send signal to process
	registerBuiltin("process_signal", func(args ...object.Object) object.Object {
		if len(args) != 2 {
			return newError("process_signal requires 2 arguments (pid, signal)")
		}
		if args[0].Type() != object.NUMBER_OBJ || args[1].Type() != object.NUMBER_OBJ {
			return newError("both arguments must be NUMBER")
		}

		pid := int(args[0].(*object.Number).Value)
		sig := syscall.Signal(int(args[1].(*object.Number).Value))

		process, err := os.FindProcess(pid)
		if err != nil {
			return newError("failed to find process: %s", err.Error())
		}

		if err := process.Signal(sig); err != nil {
			return newError("failed to send signal: %s", err.Error())
		}

		return object.NULL
	})

	// process_ache_ki (প্রসেস আছে কি) - Check if process exists
	registerBuiltin("process_ache_ki", func(args ...object.Object) object.Object {
		if len(args) != 1 {
			return newError("process_ache_ki requires 1 argument (process ID)")
		}
		if args[0].Type() != object.NUMBER_OBJ {
			return newError("process ID must be NUMBER, got %s", args[0].Type())
		}

		pid := int(args[0].(*object.Number).Value)
		process, err := os.FindProcess(pid)
		if err != nil {
			return object.FALSE
		}

		// Try sending signal 0 (does nothing but checks if process exists)
		err = process.Signal(syscall.Signal(0))
		if err != nil {
			return object.FALSE
		}

		return object.TRUE
	})

	// process_parent_id (প্রসেস প্যারেন্ট আইডি) - Get parent process ID
	registerBuiltin("process_parent_id", func(args ...object.Object) object.Object {
		return &object.Number{Value: float64(os.Getppid())}
	})

	// process_chalu (প্রসেস চালু) - Start process in background
	// Returns: { "pid": number, "wait": function }
	registerBuiltin("process_chalu", func(args ...object.Object) object.Object {
		if len(args) < 1 {
			return newError("process_chalu requires at least 1 argument (command)")
		}

		if args[0].Type() != object.STRING_OBJ {
			return newError("command must be STRING, got %s", args[0].Type())
		}
		cmdStr := args[0].(*object.String).Value

		// Parse command arguments
		var cmdArgs []string
		if len(args) > 1 {
			if args[1].Type() == object.ARRAY_OBJ {
				arr := args[1].(*object.Array)
				cmdArgs = make([]string, len(arr.Elements))
				for i, elem := range arr.Elements {
					if elem.Type() != object.STRING_OBJ {
						return newError("command arguments must be strings")
					}
					cmdArgs[i] = elem.(*object.String).Value
				}
			}
		}

		// Create command
		var cmd *exec.Cmd
		if len(cmdArgs) > 0 {
			cmd = exec.Command(cmdStr, cmdArgs...)
		} else {
			if runtime.GOOS == "windows" {
				cmd = exec.Command("cmd", "/C", cmdStr)
			} else {
				cmd = exec.Command("sh", "-c", cmdStr)
			}
		}

		// Start process
		if err := cmd.Start(); err != nil {
			return newError("failed to start process: %s", err.Error())
		}

		// Return process info
		result := make(map[string]object.Object)
		result["pid"] = &object.Number{Value: float64(cmd.Process.Pid)}

		return &object.Map{Pairs: result}
	})

	// process_opekha (প্রসেস অপেক্ষা) - Wait for process by PID
	// Returns: { "code": number }
	registerBuiltin("process_opekha", func(args ...object.Object) object.Object {
		if len(args) != 1 {
			return newError("process_opekha requires 1 argument (process ID)")
		}
		if args[0].Type() != object.NUMBER_OBJ {
			return newError("process ID must be NUMBER, got %s", args[0].Type())
		}

		pid := int(args[0].(*object.Number).Value)
		process, err := os.FindProcess(pid)
		if err != nil {
			return newError("failed to find process: %s", err.Error())
		}

		state, err := process.Wait()
		if err != nil {
			return newError("failed to wait for process: %s", err.Error())
		}

		result := make(map[string]object.Object)
		result["code"] = &object.Number{Value: float64(state.ExitCode())}

		return &object.Map{Pairs: result}
	})

	// ==================== Working Directory ====================

	// kaj_directory (কাজ ডিরেক্টরি) - Get current working directory
	registerBuiltin("kaj_directory", func(args ...object.Object) object.Object {
		dir, err := os.Getwd()
		if err != nil {
			return newError("failed to get working directory: %s", err.Error())
		}
		return &object.String{Value: dir}
	})

	// kaj_directory_bodol (কাজ ডিরেক্টরি বদল) - Change working directory
	registerBuiltin("kaj_directory_bodol", func(args ...object.Object) object.Object {
		if len(args) != 1 {
			return newError("kaj_directory_bodol requires 1 argument (directory path)")
		}
		if args[0].Type() != object.STRING_OBJ {
			return newError("directory path must be STRING, got %s", args[0].Type())
		}

		dir := args[0].(*object.String).Value
		if err := os.Chdir(dir); err != nil {
			return newError("failed to change directory: %s", err.Error())
		}

		return object.NULL
	})
}
