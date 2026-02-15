package redis

import (
	"BanglaCode/src/object"
	"fmt"
	"strings"
)

// ValidateConnection checks if a connection is valid for Redis operations
func ValidateConnection(conn *object.DBConnection) error {
	if conn == nil {
		return fmt.Errorf("connection is nil")
	}

	if conn.DBType != "redis" {
		return fmt.Errorf("expected redis connection, got %s", conn.DBType)
	}

	if conn.Native == nil {
		return fmt.Errorf("native connection is nil")
	}

	return nil
}

// BuildRedisAddr builds a Redis connection address
func BuildRedisAddr(host string, port int) string {
	return fmt.Sprintf("%s:%d", host, port)
}

// IsRedisError checks if an error is Redis-specific
func IsRedisError(err error) bool {
	if err == nil {
		return false
	}

	errMsg := err.Error()
	return strings.Contains(errMsg, "redis") ||
		strings.Contains(errMsg, "Redis") ||
		strings.Contains(errMsg, "REDIS")
}

// ValidateKey validates a Redis key
func ValidateKey(key string) error {
	if key == "" {
		return fmt.Errorf("key cannot be empty")
	}

	// Redis keys are binary safe, but we can add basic validation
	if len(key) > 512*1024*1024 { // 512MB limit
		return fmt.Errorf("key too long (max 512MB)")
	}

	return nil
}

// ValidateValue validates a Redis value
func ValidateValue(value string) error {
	// Redis values can be up to 512MB
	if len(value) > 512*1024*1024 {
		return fmt.Errorf("value too long (max 512MB)")
	}

	return nil
}

// ParseRedisURL parses a Redis connection URL
// Format: redis://user:password@host:port/db
func ParseRedisURL(url string) (*object.Map, error) {
	if !strings.HasPrefix(url, "redis://") {
		return nil, fmt.Errorf("invalid Redis URL format")
	}

	// Create empty config map
	config := &object.Map{Pairs: make(map[string]object.Object)}

	// For now, return empty config (user should use map format)
	// Full URL parsing can be added later
	return config, nil
}
