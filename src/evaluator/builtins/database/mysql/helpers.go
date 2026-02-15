package mysql

import (
	"BanglaCode/src/object"
	"fmt"
	"strings"
)

// ValidateConnection checks if a connection is valid for MySQL operations
func ValidateConnection(conn *object.DBConnection) error {
	if conn == nil {
		return fmt.Errorf("connection is nil")
	}

	if conn.DBType != "mysql" {
		return fmt.Errorf("expected mysql connection, got %s", conn.DBType)
	}

	if conn.Native == nil {
		return fmt.Errorf("native connection is nil")
	}

	return nil
}

// BuildDSN builds a MySQL DSN (Data Source Name)
func BuildDSN(user, password, host string, port int, database, charset, parseTime string) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=%s",
		user, password, host, port, database, charset, parseTime,
	)
}

// IsMySQLError checks if an error is MySQL-specific
func IsMySQLError(err error) bool {
	if err == nil {
		return false
	}

	errMsg := err.Error()
	return strings.Contains(errMsg, "mysql:") ||
		strings.Contains(errMsg, "MySQL") ||
		strings.Contains(errMsg, "Error 1")
}

// GetQueryType returns the type of SQL query
func GetQueryType(query string) string {
	query = strings.TrimSpace(strings.ToUpper(query))

	switch {
	case strings.HasPrefix(query, "SELECT"):
		return "SELECT"
	case strings.HasPrefix(query, "INSERT"):
		return "INSERT"
	case strings.HasPrefix(query, "UPDATE"):
		return "UPDATE"
	case strings.HasPrefix(query, "DELETE"):
		return "DELETE"
	case strings.HasPrefix(query, "CREATE"):
		return "CREATE"
	case strings.HasPrefix(query, "ALTER"):
		return "ALTER"
	case strings.HasPrefix(query, "DROP"):
		return "DROP"
	case strings.HasPrefix(query, "SHOW"):
		return "SHOW"
	case strings.HasPrefix(query, "DESCRIBE"), strings.HasPrefix(query, "DESC"):
		return "DESCRIBE"
	}

	return "UNKNOWN"
}

// SanitizeQuery performs basic query sanitization
func SanitizeQuery(query string) string {
	query = strings.TrimSpace(query)

	// Remove single-line comments
	lines := strings.Split(query, "\n")
	sanitized := []string{}
	for _, line := range lines {
		if strings.HasPrefix(strings.TrimSpace(line), "--") {
			continue
		}
		sanitized = append(sanitized, line)
	}

	return strings.Join(sanitized, "\n")
}
