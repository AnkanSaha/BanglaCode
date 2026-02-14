package system

import (
	"BanglaCode/src/object"
	"os"
	"os/user"
	"runtime"
	"sync"
)

// Performance optimization: Cache system info that doesn't change
var (
	cachedHostname string
	cachedUsername string
	cachedHomeDir  string
	cachedOS       string
	cachedArch     string
	cachedCPUCount int
	cacheOnce      sync.Once
)

func initSystemCache() {
	cacheOnce.Do(func() {
		// Cache hostname
		if h, err := os.Hostname(); err == nil {
			cachedHostname = h
		}

		// Cache user info
		if u, err := user.Current(); err == nil {
			cachedUsername = u.Username
			cachedHomeDir = u.HomeDir
		}

		// Cache OS and architecture (these never change during execution)
		cachedOS = runtime.GOOS
		cachedArch = runtime.GOARCH
		cachedCPUCount = runtime.NumCPU()
	})
}

func init() {
	// Initialize cache on startup
	initSystemCache()

	// ==================== System Information ====================

	// os_naam (ওএস নাম) - Get operating system name
	registerBuiltin("os_naam", func(args ...object.Object) object.Object {
		return &object.String{Value: cachedOS}
	})

	// bibhag (বিভাগ) - Get system architecture
	registerBuiltin("bibhag", func(args ...object.Object) object.Object {
		return &object.String{Value: cachedArch}
	})

	// hostname (হোস্টনাম) - Get system hostname (cached for performance)
	registerBuiltin("hostname", func(args ...object.Object) object.Object {
		if cachedHostname == "" {
			h, err := os.Hostname()
			if err != nil {
				return newError("failed to get hostname: %s", err.Error())
			}
			return &object.String{Value: h}
		}
		return &object.String{Value: cachedHostname}
	})

	// cpu_sonkha (সিপিইউ সংখ্যা) - Get number of CPU cores
	registerBuiltin("cpu_sonkha", func(args ...object.Object) object.Object {
		return &object.Number{Value: float64(cachedCPUCount)}
	})

	// ==================== User Information ====================

	// bebosthok_naam (ব্যবস্থাপক নাম) - Get current username (cached)
	registerBuiltin("bebosthok_naam", func(args ...object.Object) object.Object {
		if cachedUsername == "" {
			u, err := user.Current()
			if err != nil {
				return newError("failed to get username: %s", err.Error())
			}
			return &object.String{Value: u.Username}
		}
		return &object.String{Value: cachedUsername}
	})

	// bari_directory (বাড়ি ডিরেক্টরি) - Get user home directory (cached)
	registerBuiltin("bari_directory", func(args ...object.Object) object.Object {
		if cachedHomeDir == "" {
			dir, err := os.UserHomeDir()
			if err != nil {
				return newError("failed to get home directory: %s", err.Error())
			}
			return &object.String{Value: dir}
		}
		return &object.String{Value: cachedHomeDir}
	})
}
