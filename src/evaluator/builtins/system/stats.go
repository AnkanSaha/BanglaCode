package system

import (
	"BanglaCode/src/object"
	"fmt"
	"os"
	"runtime"
)

func init() {
	// ==================== System Statistics ====================

	// memory_total (মেমরি টোটাল) - Get total system memory in bytes
	registerBuiltin("memory_total", func(args ...object.Object) object.Object {
		var totalMem uint64

		switch runtime.GOOS {
		case "linux":
			// Read /proc/meminfo on Linux
			data, err := os.ReadFile("/proc/meminfo")
			if err == nil {
				var memTotal uint64
				_, _ = fmt.Sscanf(string(data), "MemTotal: %d kB", &memTotal)
				totalMem = memTotal * 1024 // Convert KB to bytes
			}
		case "darwin":
			// Use sysctl on macOS - placeholder
			totalMem = 0
		default:
			// Windows or other platforms - use runtime stats as approximation
			var m runtime.MemStats
			runtime.ReadMemStats(&m)
			totalMem = m.Sys
		}

		return &object.Number{Value: float64(totalMem)}
	})

	// memory_bebohrito (মেমরি ব্যবহৃত) - Get used memory in bytes
	registerBuiltin("memory_bebohrito", func(args ...object.Object) object.Object {
		var m runtime.MemStats
		runtime.ReadMemStats(&m)

		// Return heap allocated memory
		return &object.Number{Value: float64(m.Alloc)}
	})

	// memory_mukt (মেমরি মুক্ত) - Get free memory in bytes
	registerBuiltin("memory_mukt", func(args ...object.Object) object.Object {
		var freeMem uint64

		if runtime.GOOS == "linux" {
			// Read /proc/meminfo on Linux
			data, err := os.ReadFile("/proc/meminfo")
			if err == nil {
				// Parse for MemAvailable (better than MemFree)
				lines := string(data)
				var memAvail uint64
				_, _ = fmt.Sscanf(lines, "MemTotal: %*d kB\nMemFree: %d kB", &memAvail)
				freeMem = memAvail * 1024
			}
		}

		if freeMem == 0 {
			return newError("free memory not available on this platform")
		}

		return &object.Number{Value: float64(freeMem)}
	})

	// cpu_bebohrito (সিপিইউ ব্যবহৃত) - Get CPU usage percentage
	registerBuiltin("cpu_bebohrito", func(args ...object.Object) object.Object {
		// CPU usage requires sampling over time
		// This is a placeholder - real implementation would sample /proc/stat
		return newError("cpu_bebohrito not fully implemented yet")
	})

	// disk_akar (ডিস্ক আকার) - Get disk size in bytes
	registerBuiltin("disk_akar", func(args ...object.Object) object.Object {
		// Default to root filesystem
		path := "/"
		if len(args) > 0 {
			if args[0].Type() != object.STRING_OBJ {
				return newError("path must be STRING, got %s", args[0].Type())
			}
			path = args[0].(*object.String).Value
		}

		total, _, _, err := getDiskStats(path)
		if err != nil {
			return newError("failed to get disk stats: %s", err.Error())
		}

		return &object.Number{Value: float64(total)}
	})

	// disk_bebohrito (ডিস্ক ব্যবহৃত) - Get disk used space in bytes
	registerBuiltin("disk_bebohrito", func(args ...object.Object) object.Object {
		// Default to root filesystem
		path := "/"
		if len(args) > 0 {
			if args[0].Type() != object.STRING_OBJ {
				return newError("path must be STRING, got %s", args[0].Type())
			}
			path = args[0].(*object.String).Value
		}

		total, free, _, err := getDiskStats(path)
		if err != nil {
			return newError("failed to get disk stats: %s", err.Error())
		}

		// Used = total - free
		usedSize := total - free

		return &object.Number{Value: float64(usedSize)}
	})

	// disk_mukt (ডিস্ক মুক্ত) - Get disk free space in bytes
	registerBuiltin("disk_mukt", func(args ...object.Object) object.Object {
		// Default to root filesystem
		path := "/"
		if len(args) > 0 {
			if args[0].Type() != object.STRING_OBJ {
				return newError("path must be STRING, got %s", args[0].Type())
			}
			path = args[0].(*object.String).Value
		}

		_, _, avail, err := getDiskStats(path)
		if err != nil {
			return newError("failed to get disk stats: %s", err.Error())
		}

		// Available to non-root users
		return &object.Number{Value: float64(avail)}
	})
}
