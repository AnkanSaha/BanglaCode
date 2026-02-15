package database

import (
	"BanglaCode/src/object"
	"database/sql"
	"fmt"
	"sync/atomic"
)

// Counter for generating unique connection IDs (thread-safe)
var connIDCounter int64

// newError creates a new error object with formatted message
func newError(format string, a ...interface{}) *object.Error {
	return &object.Error{Message: fmt.Sprintf(format, a...)}
}

// generateConnID generates a unique connection identifier
func generateConnID(dbType string) string {
	id := atomic.AddInt64(&connIDCounter, 1)
	return fmt.Sprintf("%s-%d", dbType, id)
}

// toObject converts a Go value to a BanglaCode object
func toObject(value interface{}) object.Object {
	if value == nil {
		return object.NULL
	}

	switch v := value.(type) {
	case int:
		return &object.Number{Value: float64(v)}
	case int32:
		return &object.Number{Value: float64(v)}
	case int64:
		return &object.Number{Value: float64(v)}
	case float32:
		return &object.Number{Value: float64(v)}
	case float64:
		return &object.Number{Value: v}
	case string:
		return &object.String{Value: v}
	case bool:
		return object.NativeBoolToBooleanObject(v)
	case []byte:
		return &object.String{Value: string(v)}
	default:
		// For unknown types, convert to string
		return &object.String{Value: fmt.Sprintf("%v", v)}
	}
}

// fromObject converts a BanglaCode object to a Go value
func fromObject(obj object.Object) interface{} {
	switch o := obj.(type) {
	case *object.Number:
		return o.Value
	case *object.String:
		return o.Value
	case *object.Boolean:
		return o.Value
	case *object.Null:
		return nil
	default:
		return o.Inspect()
	}
}

// convertRow converts a SQL row to a BanglaCode map
func convertRow(columns []string, values []interface{}) map[string]object.Object {
	row := make(map[string]object.Object, len(columns))
	for i, col := range columns {
		row[col] = toObject(values[i])
	}
	return row
}

// extractConfigString extracts a string value from connection config map
func extractConfigString(config *object.Map, key string, defaultValue string) string {
	if val, ok := config.Pairs[key]; ok {
		if str, ok := val.(*object.String); ok {
			return str.Value
		}
	}
	return defaultValue
}

// extractConfigNumber extracts a numeric value from connection config map
func extractConfigNumber(config *object.Map, key string, defaultValue float64) float64 {
	if val, ok := config.Pairs[key]; ok {
		if num, ok := val.(*object.Number); ok {
			return num.Value
		}
	}
	return defaultValue
}

// validateConnection checks if an object is a valid database connection
func validateConnection(obj object.Object, expectedType string) (*object.DBConnection, error) {
	conn, ok := obj.(*object.DBConnection)
	if !ok {
		return nil, fmt.Errorf("expected DB_CONNECTION, got %s", obj.Type())
	}

	if expectedType != "" && conn.DBType != expectedType {
		return nil, fmt.Errorf("expected %s connection, got %s", expectedType, conn.DBType)
	}

	return conn, nil
}

// convertSQLRows converts sql.Rows to a slice of maps (optimized for performance)
func convertSQLRows(rows *sql.Rows) ([]map[string]object.Object, error) {
	// Get column names once (reused for all rows)
	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	// Pre-allocate result slice with estimated capacity
	result := make([]map[string]object.Object, 0, 100)

	// Scan rows efficiently
	for rows.Next() {
		// Allocate value holders
		values := make([]interface{}, len(columns))
		valuePtrs := make([]interface{}, len(columns))
		for i := range values {
			valuePtrs[i] = &values[i]
		}

		// Scan row
		if err := rows.Scan(valuePtrs...); err != nil {
			return nil, err
		}

		// Convert to BanglaCode map
		row := convertRow(columns, values)
		result = append(result, row)
	}

	// Check for iteration errors
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

// createDBResult creates a DBResult from SQL query result
func createDBResult(rows []map[string]object.Object, rowsAffected, lastInsertID int64) *object.DBResult {
	return &object.DBResult{
		Rows:         rows,
		RowsAffected: rowsAffected,
		LastInsertID: lastInsertID,
		Error:        nil,
	}
}
