package mongodb

import (
	"BanglaCode/src/object"
	"fmt"
	"strings"
)

// ValidateConnection checks if a connection is valid for MongoDB operations
func ValidateConnection(conn *object.DBConnection) error {
	if conn == nil {
		return fmt.Errorf("connection is nil")
	}

	if conn.DBType != "mongodb" {
		return fmt.Errorf("expected mongodb connection, got %s", conn.DBType)
	}

	if conn.Native == nil {
		return fmt.Errorf("native connection is nil")
	}

	return nil
}

// BuildMongoURI builds a MongoDB connection URI
func BuildMongoURI(username, password, host string, port int, database string) string {
	if username != "" && password != "" {
		return fmt.Sprintf("mongodb://%s:%s@%s:%d/%s",
			username, password, host, port, database)
	}
	return fmt.Sprintf("mongodb://%s:%d", host, port)
}

// IsMongoDBError checks if an error is MongoDB-specific
func IsMongoDBError(err error) bool {
	if err == nil {
		return false
	}

	errMsg := err.Error()
	return strings.Contains(errMsg, "mongo") ||
		strings.Contains(errMsg, "MongoDB") ||
		strings.Contains(errMsg, "BSON")
}

// ValidateFilter validates a MongoDB filter map
func ValidateFilter(filter *object.Map) error {
	if filter == nil {
		return fmt.Errorf("filter cannot be nil")
	}

	// Basic validation - check for valid operators
	for key := range filter.Pairs {
		if strings.HasPrefix(key, "$") {
			// MongoDB operator - validate it's a known one
			validOps := []string{"$eq", "$gt", "$gte", "$lt", "$lte", "$ne", "$in", "$nin", "$set", "$unset", "$inc"}
			valid := false
			for _, op := range validOps {
				if key == op {
					valid = true
					break
				}
			}
			if !valid {
				return fmt.Errorf("unknown MongoDB operator: %s", key)
			}
		}
	}

	return nil
}
