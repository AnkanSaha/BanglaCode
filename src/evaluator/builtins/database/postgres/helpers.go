package postgres

import (
	"BanglaCode/src/object"
	"fmt"
	"strings"
)

// ValidateConnection checks if a connection is valid for PostgreSQL operations
func ValidateConnection(conn *object.DBConnection) error {
	if conn == nil {
		return fmt.Errorf("connection is nil")
	}

	if conn.DBType != "postgres" {
		return fmt.Errorf("expected postgres connection, got %s", conn.DBType)
	}

	if conn.Native == nil {
		return fmt.Errorf("native connection is nil")
	}

	return nil
}

// BuildConnectionString builds a PostgreSQL connection string from config
func BuildConnectionString(host string, port int, database, user, password, sslmode string) string {
	return fmt.Sprintf(
		"host=%s port=%d dbname=%s user=%s password=%s sslmode=%s",
		host, port, database, user, password, sslmode,
	)
}

// IsPostgresError checks if an error is a PostgreSQL-specific error
func IsPostgresError(err error) bool {
	if err == nil {
		return false
	}

	// Check for common PostgreSQL error patterns
	errMsg := err.Error()
	return strings.Contains(errMsg, "pq:") ||
		strings.Contains(errMsg, "postgres") ||
		strings.Contains(errMsg, "PostgreSQL")
}

// ParsePostgresURL parses a PostgreSQL connection URL
// Format: postgres://user:password@host:port/database?sslmode=disable
func ParsePostgresURL(url string) (*object.Map, error) {
	// Simple URL parsing (production code would use net/url)
	if !strings.HasPrefix(url, "postgres://") && !strings.HasPrefix(url, "postgresql://") {
		return nil, fmt.Errorf("invalid PostgreSQL URL format")
	}

	// Create empty config map
	config := &object.Map{Pairs: make(map[string]object.Object)}

	// For now, return empty config (user should use map format)
	// Full URL parsing can be added later
	return config, nil
}

// SanitizeQuery performs basic query sanitization (removes dangerous patterns)
func SanitizeQuery(query string) string {
	// Remove leading/trailing whitespace
	query = strings.TrimSpace(query)

	// Remove comments (basic)
	lines := strings.Split(query, "\n")
	sanitized := []string{}
	for _, line := range lines {
		// Remove single-line comments
		if strings.HasPrefix(strings.TrimSpace(line), "--") {
			continue
		}
		sanitized = append(sanitized, line)
	}

	return strings.Join(sanitized, "\n")
}

// GetQueryType returns the type of SQL query (SELECT, INSERT, UPDATE, DELETE, etc.)
func GetQueryType(query string) string {
	query = strings.TrimSpace(strings.ToUpper(query))

	if strings.HasPrefix(query, "SELECT") {
		return "SELECT"
	} else if strings.HasPrefix(query, "INSERT") {
		return "INSERT"
	} else if strings.HasPrefix(query, "UPDATE") {
		return "UPDATE"
	} else if strings.HasPrefix(query, "DELETE") {
		return "DELETE"
	} else if strings.HasPrefix(query, "CREATE") {
		return "CREATE"
	} else if strings.HasPrefix(query, "ALTER") {
		return "ALTER"
	} else if strings.HasPrefix(query, "DROP") {
		return "DROP"
	}

	return "UNKNOWN"
}

// EstimateQueryComplexity estimates the complexity of a query (simple heuristic)
func EstimateQueryComplexity(query string) string {
	query = strings.ToUpper(query)

	// Count JOINs
	joinCount := strings.Count(query, "JOIN")

	// Count subqueries
	subqueryCount := strings.Count(query, "SELECT") - 1 // Subtract main SELECT

	// Count aggregations
	aggCount := strings.Count(query, "GROUP BY") + strings.Count(query, "HAVING")

	totalComplexity := joinCount + subqueryCount*2 + aggCount

	if totalComplexity == 0 {
		return "SIMPLE"
	} else if totalComplexity <= 3 {
		return "MEDIUM"
	} else {
		return "COMPLEX"
	}
}
