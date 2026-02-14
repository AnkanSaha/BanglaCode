package system

import (
	"BanglaCode/src/object"
	"fmt"
	"os"
	"runtime"
	"time"
)

func init() {
	// ==================== Time Operations ====================

	// shomoy_ekhon (সময় এখন) - Get current Unix timestamp
	registerBuiltin("shomoy_ekhon", func(args ...object.Object) object.Object {
		return &object.Number{Value: float64(time.Now().Unix())}
	})

	// shomoy_format (সময় ফরম্যাট) - Format timestamp to string
	// Args: timestamp (number), format (string, optional)
	registerBuiltin("shomoy_format", func(args ...object.Object) object.Object {
		if len(args) < 1 {
			return newError("shomoy_format requires at least 1 argument (timestamp)")
		}
		if args[0].Type() != object.NUMBER_OBJ {
			return newError("timestamp must be NUMBER, got %s", args[0].Type())
		}

		timestamp := int64(args[0].(*object.Number).Value)
		t := time.Unix(timestamp, 0)

		// Default format: RFC3339
		format := time.RFC3339
		if len(args) > 1 {
			if args[1].Type() != object.STRING_OBJ {
				return newError("format must be STRING, got %s", args[1].Type())
			}
			format = args[1].(*object.String).Value
		}

		return &object.String{Value: t.Format(format)}
	})

	// shomoy_parse (সময় পার্স) - Parse time string to Unix timestamp
	// Args: timestring (string), format (string, optional)
	registerBuiltin("shomoy_parse", func(args ...object.Object) object.Object {
		if len(args) < 1 {
			return newError("shomoy_parse requires at least 1 argument (time string)")
		}
		if args[0].Type() != object.STRING_OBJ {
			return newError("time string must be STRING, got %s", args[0].Type())
		}

		timeStr := args[0].(*object.String).Value

		// Default format: RFC3339
		format := time.RFC3339
		if len(args) > 1 {
			if args[1].Type() != object.STRING_OBJ {
				return newError("format must be STRING, got %s", args[1].Type())
			}
			format = args[1].(*object.String).Value
		}

		t, err := time.Parse(format, timeStr)
		if err != nil {
			return newError("failed to parse time: %s", err.Error())
		}

		return &object.Number{Value: float64(t.Unix())}
	})

	// uptime (আপটাইম) - Get system uptime in seconds
	registerBuiltin("uptime", func(args ...object.Object) object.Object {
		var uptime int64

		// Platform-specific implementation
		if runtime.GOOS == "linux" {
			// Read /proc/uptime on Linux
			data, err := os.ReadFile("/proc/uptime")
			if err == nil {
				var uptimeFloat float64
				_, _ = fmt.Sscanf(string(data), "%f", &uptimeFloat)
				uptime = int64(uptimeFloat)
			}
		} else if runtime.GOOS == "darwin" || runtime.GOOS == "freebsd" {
			// Use sysctl on macOS/BSD
			// This is a simplified version - actual implementation would use sysctl
			// For now, return 0 as placeholder
			uptime = 0
		} else {
			// Windows or other platforms
			uptime = 0
		}

		if uptime == 0 {
			return newError("uptime not supported on this platform")
		}

		return &object.Number{Value: float64(uptime)}
	})

	// boot_shomoy (বুট সময়) - Get system boot time (Unix timestamp)
	registerBuiltin("boot_shomoy", func(args ...object.Object) object.Object {
		var bootTime int64

		// Platform-specific implementation
		if runtime.GOOS == "linux" {
			// Read /proc/uptime and calculate boot time
			data, err := os.ReadFile("/proc/uptime")
			if err == nil {
				var uptimeFloat float64
				_, _ = fmt.Sscanf(string(data), "%f", &uptimeFloat)
				bootTime = time.Now().Unix() - int64(uptimeFloat)
			}
		} else {
			// Other platforms
			bootTime = 0
		}

		if bootTime == 0 {
			return newError("boot time not supported on this platform")
		}

		return &object.Number{Value: float64(bootTime)}
	})

	// timezone (টাইমজোন) - Get current timezone
	registerBuiltin("timezone", func(args ...object.Object) object.Object {
		zone, _ := time.Now().Zone()
		return &object.String{Value: zone}
	})
}
