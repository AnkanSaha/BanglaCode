package mongodb

import (
	"BanglaCode/src/object"
	"fmt"
)

// extractString extracts a string value from config map with default fallback
func extractString(config *object.Map, key string, defaultValue string) string {
	if val, ok := config.Pairs[key]; ok {
		if str, ok := val.(*object.String); ok {
			return str.Value
		}
	}
	return defaultValue
}

// extractNumber extracts a number value from config map with default fallback
func extractNumber(config *object.Map, key string, defaultValue float64) float64 {
	if val, ok := config.Pairs[key]; ok {
		if num, ok := val.(*object.Number); ok {
			return num.Value
		}
	}
	return defaultValue
}

var connIDCounter int64

// generateConnID generates a unique connection ID for tracking MongoDB clients
func generateConnID() string {
	connIDCounter++
	return fmt.Sprintf("mongodb-%d", connIDCounter)
}
