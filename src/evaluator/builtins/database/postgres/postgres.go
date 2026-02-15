package postgres

import (
	"BanglaCode/src/object"
	"database/sql"
	"fmt"
	"sync"

	_ "github.com/lib/pq" // PostgreSQL driver
)

var (
	connections   = make(map[string]*sql.DB)
	connectionsMu sync.RWMutex
)

// Connect creates a new PostgreSQL connection
func Connect(config *object.Map) (*object.DBConnection, error) {
	// Extract connection parameters
	host := extractString(config, "host", "localhost")
	port := extractNumber(config, "port", 5432)
	database := extractString(config, "database", "postgres")
	user := extractString(config, "user", "postgres")
	password := extractString(config, "password", "")
	sslmode := extractString(config, "sslmode", "disable")

	// Build connection string
	connStr := fmt.Sprintf(
		"host=%s port=%d dbname=%s user=%s password=%s sslmode=%s",
		host, int(port), database, user, password, sslmode,
	)

	// Open connection
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to open connection: %v", err)
	}

	// Test connection
	if err := db.Ping(); err != nil {
		db.Close()
		return nil, fmt.Errorf("failed to ping database: %v", err)
	}

	// Configure connection pool for the underlying *sql.DB
	maxConns := extractNumber(config, "max_conns", 10)
	db.SetMaxOpenConns(int(maxConns))
	db.SetMaxIdleConns(int(maxConns / 2))

	// Generate unique connection ID
	connID := generateConnID()

	// Store connection globally
	connectionsMu.Lock()
	connections[connID] = db
	connectionsMu.Unlock()

	// Create metadata
	metadata := make(map[string]object.Object)
	metadata["host"] = &object.String{Value: host}
	metadata["port"] = &object.Number{Value: port}
	metadata["database"] = &object.String{Value: database}
	metadata["user"] = &object.String{Value: user}

	// Create DBConnection object
	conn := &object.DBConnection{
		ID:       connID,
		DBType:   "postgres",
		Native:   db,
		Metadata: metadata,
	}

	return conn, nil
}

// Close closes a PostgreSQL connection
func Close(conn *object.DBConnection) error {
	if conn.DBType != "postgres" {
		return fmt.Errorf("expected postgres connection, got %s", conn.DBType)
	}

	db, ok := conn.Native.(*sql.DB)
	if !ok {
		return fmt.Errorf("invalid native connection type")
	}

	// Remove from global registry
	connectionsMu.Lock()
	delete(connections, conn.ID)
	connectionsMu.Unlock()

	return db.Close()
}

// Query executes a SELECT query
func Query(conn *object.DBConnection, query string) (*object.DBResult, error) {
	db, ok := conn.Native.(*sql.DB)
	if !ok {
		return nil, fmt.Errorf("invalid native connection type")
	}

	rows, err := db.Query(query)
	if err != nil {
		return &object.DBResult{Error: &object.Error{Message: err.Error()}}, nil
	}
	defer rows.Close()

	// Convert rows to BanglaCode objects
	result, err := convertRows(rows)
	if err != nil {
		return &object.DBResult{Error: &object.Error{Message: err.Error()}}, nil
	}

	return result, nil
}

// Exec executes an INSERT, UPDATE, or DELETE statement
func Exec(conn *object.DBConnection, query string) (*object.DBResult, error) {
	db, ok := conn.Native.(*sql.DB)
	if !ok {
		return nil, fmt.Errorf("invalid native connection type")
	}

	result, err := db.Exec(query)
	if err != nil {
		return &object.DBResult{Error: &object.Error{Message: err.Error()}}, nil
	}

	rowsAffected, _ := result.RowsAffected()
	lastInsertID, _ := result.LastInsertId()

	return &object.DBResult{
		Rows:         []map[string]object.Object{},
		RowsAffected: rowsAffected,
		LastInsertID: lastInsertID,
	}, nil
}

// PreparedQuery executes a parameterized query (SQL injection safe)
func PreparedQuery(conn *object.DBConnection, query string, params []object.Object) (*object.DBResult, error) {
	db, ok := conn.Native.(*sql.DB)
	if !ok {
		return nil, fmt.Errorf("invalid native connection type")
	}

	// Convert BanglaCode objects to Go values
	args := make([]interface{}, len(params))
	for i, param := range params {
		args[i] = objectToGoValue(param)
	}

	// Check if query is SELECT or DML
	if isSelectQuery(query) {
		rows, err := db.Query(query, args...)
		if err != nil {
			return &object.DBResult{Error: &object.Error{Message: err.Error()}}, nil
		}
		defer rows.Close()

		result, err := convertRows(rows)
		if err != nil {
			return &object.DBResult{Error: &object.Error{Message: err.Error()}}, nil
		}

		return result, nil
	}

	// Execute DML (INSERT/UPDATE/DELETE)
	result, err := db.Exec(query, args...)
	if err != nil {
		return &object.DBResult{Error: &object.Error{Message: err.Error()}}, nil
	}

	rowsAffected, _ := result.RowsAffected()
	lastInsertID, _ := result.LastInsertId()

	return &object.DBResult{
		Rows:         []map[string]object.Object{},
		RowsAffected: rowsAffected,
		LastInsertID: lastInsertID,
	}, nil
}

// BeginTransaction starts a new database transaction
func BeginTransaction(conn *object.DBConnection) (*sql.Tx, error) {
	db, ok := conn.Native.(*sql.DB)
	if !ok {
		return nil, fmt.Errorf("invalid native connection type")
	}

	return db.Begin()
}

// Commit commits a transaction
func Commit(tx *sql.Tx) error {
	return tx.Commit()
}

// Rollback rolls back a transaction
func Rollback(tx *sql.Tx) error {
	return tx.Rollback()
}

// BulkInsert performs efficient bulk insert operation
// rows is an array of arrays, where each inner array represents a row
// Example: BulkInsert(conn, "users", ["name", "age"], [["Alice", 25], ["Bob", 30]])
func BulkInsert(conn *object.DBConnection, table string, columns []string, rows *object.Array) (*object.DBResult, error) {
	db, ok := conn.Native.(*sql.DB)
	if !ok {
		return nil, fmt.Errorf("invalid native connection type")
	}

	if len(rows.Elements) == 0 {
		return &object.DBResult{
			Rows:         []map[string]object.Object{},
			RowsAffected: 0,
		}, nil
	}

	// Build multi-row INSERT statement
	// INSERT INTO table (col1, col2) VALUES ($1, $2), ($3, $4), ...
	numCols := len(columns)
	numRows := len(rows.Elements)

	// Build column names
	colNames := ""
	for i, col := range columns {
		if i > 0 {
			colNames += ", "
		}
		colNames += col
	}

	// Build placeholders and collect values
	values := make([]interface{}, 0, numRows*numCols)
	placeholders := ""

	paramIndex := 1
	for i, rowObj := range rows.Elements {
		rowArray, ok := rowObj.(*object.Array)
		if !ok {
			return nil, fmt.Errorf("row %d must be an array, got %s", i, rowObj.Type())
		}

		if len(rowArray.Elements) != numCols {
			return nil, fmt.Errorf("row %d has %d columns, expected %d", i, len(rowArray.Elements), numCols)
		}

		if i > 0 {
			placeholders += ", "
		}
		placeholders += "("

		for j, valueObj := range rowArray.Elements {
			if j > 0 {
				placeholders += ", "
			}
			placeholders += fmt.Sprintf("$%d", paramIndex)
			paramIndex++
			values = append(values, objectToGoValue(valueObj))
		}

		placeholders += ")"
	}

	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES %s", table, colNames, placeholders)

	// Execute bulk insert
	result, err := db.Exec(query, values...)
	if err != nil {
		return &object.DBResult{Error: &object.Error{Message: err.Error()}}, nil
	}

	rowsAffected, _ := result.RowsAffected()
	lastInsertID, _ := result.LastInsertId()

	return &object.DBResult{
		Rows:         []map[string]object.Object{},
		RowsAffected: rowsAffected,
		LastInsertID: lastInsertID,
	}, nil
}

// Helper functions

// extractString extracts a string from config map with default value
func extractString(config *object.Map, key string, defaultValue string) string {
	if val, ok := config.Pairs[key]; ok {
		if str, ok := val.(*object.String); ok {
			return str.Value
		}
	}
	return defaultValue
}

// extractNumber extracts a number from config map with default value
func extractNumber(config *object.Map, key string, defaultValue float64) float64 {
	if val, ok := config.Pairs[key]; ok {
		if num, ok := val.(*object.Number); ok {
			return num.Value
		}
	}
	return defaultValue
}

// generateConnID generates a unique connection ID
var connIDCounter int64

func generateConnID() string {
	connIDCounter++
	return fmt.Sprintf("postgres-%d", connIDCounter)
}

// objectToGoValue converts a BanglaCode object to Go value
func objectToGoValue(obj object.Object) interface{} {
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

// convertRows converts SQL rows to DBResult
func convertRows(rows *sql.Rows) (*object.DBResult, error) {
	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	// Pre-allocate result slice
	result := make([]map[string]object.Object, 0, 100)

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
		row := make(map[string]object.Object, len(columns))
		for i, col := range columns {
			row[col] = goValueToObject(values[i])
		}

		result = append(result, row)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &object.DBResult{
		Rows:         result,
		RowsAffected: int64(len(result)),
	}, nil
}

// goValueToObject converts a Go value to BanglaCode object
func goValueToObject(value interface{}) object.Object {
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
		return &object.String{Value: fmt.Sprintf("%v", v)}
	}
}

// isSelectQuery checks if a query is a SELECT statement
func isSelectQuery(query string) bool {
	trimmed := ""
	for _, ch := range query {
		if ch != ' ' && ch != '\t' && ch != '\n' && ch != '\r' {
			trimmed += string(ch)
			if len(trimmed) >= 6 {
				break
			}
		}
	}

	if len(trimmed) >= 6 {
		prefix := trimmed[:6]
		return prefix == "SELECT" || prefix == "select"
	}
	return false
}
