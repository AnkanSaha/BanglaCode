package system

import (
	"BanglaCode/src/evaluator/builtins/system/env"
	"BanglaCode/src/evaluator/builtins/system/filesystem"
	"BanglaCode/src/evaluator/builtins/system/info"
	"BanglaCode/src/evaluator/builtins/system/network"
	"BanglaCode/src/evaluator/builtins/system/path"
	"BanglaCode/src/evaluator/builtins/system/process"
	"BanglaCode/src/evaluator/builtins/system/stats"
	"BanglaCode/src/evaluator/builtins/system/time"
	"BanglaCode/src/object"
)

// Builtins is the map that holds all system built-in functions
// Functions from all subdirectory packages are merged into this map
var Builtins = make(map[string]*object.Builtin, 60)

func init() {
	// Merge all subdirectory builtins into the main Builtins map
	// This allows all system functions to be accessed through system.Builtins

	// Filesystem operations (file I/O, directories, temp files, symlinks)
	for name, fn := range filesystem.Builtins {
		Builtins[name] = fn
	}

	// System statistics (memory, CPU, disk usage)
	for name, fn := range stats.Builtins {
		Builtins[name] = fn
	}

	// Path operations (join, split, absolute paths, etc.)
	for name, fn := range path.Builtins {
		Builtins[name] = fn
	}

	// Process execution (run commands, get PID, etc.)
	for name, fn := range process.Builtins {
		Builtins[name] = fn
	}

	// Network information (interfaces, IP addresses, etc.)
	for name, fn := range network.Builtins {
		Builtins[name] = fn
	}

	// Time operations (current time, formatting, parsing, uptime)
	for name, fn := range time.Builtins {
		Builtins[name] = fn
	}

	// Environment variables (get, set, list)
	for name, fn := range env.Builtins {
		Builtins[name] = fn
	}

	// System information (hostname, username, OS, architecture)
	for name, fn := range info.Builtins {
		Builtins[name] = fn
	}
}
