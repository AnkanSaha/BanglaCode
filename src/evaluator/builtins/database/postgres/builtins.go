package postgres

import (
	"BanglaCode/src/object"
	"database/sql"
	"fmt"
	"sync"
)

// Builtins holds all PostgreSQL built-in functions
var Builtins = make(map[string]*object.Builtin)

func init() {
	// Connection management
	registerBuiltin("db_jukto_postgres", dbJuktoPostgres)
	registerBuiltin("db_bandho_postgres", dbBandhoPostgres)

	// Query operations (synchronous)
	registerBuiltin("db_query_postgres", dbQueryPostgres)
	registerBuiltin("db_exec_postgres", dbExecPostgres)
	registerBuiltin("db_proshno_postgres", dbProshnoPostgres)

	// Query operations (asynchronous)
	registerBuiltin("db_query_async_postgres", dbQueryAsyncPostgres)
	registerBuiltin("db_exec_async_postgres", dbExecAsyncPostgres)
	registerBuiltin("db_proshno_async_postgres", dbProshnoAsyncPostgres)

	// Transaction support
	registerBuiltin("db_transaction_shuru_postgres", dbTransactionShuruPostgres)
	registerBuiltin("db_commit_postgres", dbCommitPostgres)
	registerBuiltin("db_rollback_postgres", dbRollbackPostgres)

	// Bulk operations
	registerBuiltin("db_bulk_insert_postgres", dbBulkInsertPostgres)
}

func registerBuiltin(name string, fn object.BuiltinFunction) {
	Builtins[name] = &object.Builtin{Fn: fn}
}

// newError creates a new error object
func newError(format string, a ...interface{}) *object.Error {
	return &object.Error{Message: fmt.Sprintf(format, a...)}
}

// db_jukto_postgres - Connect to PostgreSQL database
// Usage: db_jukto("postgres", {host: "localhost", port: 5432, database: "mydb"})
func dbJuktoPostgres(args ...object.Object) object.Object {
	if len(args) != 1 {
		return newError("db_jukto_postgres: wrong number of arguments. got=%d, want=1", len(args))
	}

	config, ok := args[0].(*object.Map)
	if !ok {
		return newError("db_jukto_postgres: argument must be a map, got %s", args[0].Type())
	}

	conn, err := Connect(config)
	if err != nil {
		return newError("db_jukto_postgres: %s", err.Error())
	}

	return conn
}

// db_bandho_postgres - Close PostgreSQL connection
func dbBandhoPostgres(args ...object.Object) object.Object {
	if len(args) != 1 {
		return newError("db_bandho_postgres: wrong number of arguments. got=%d, want=1", len(args))
	}

	conn, ok := args[0].(*object.DBConnection)
	if !ok {
		return newError("db_bandho_postgres: argument must be DB_CONNECTION, got %s", args[0].Type())
	}

	if err := Close(conn); err != nil {
		return newError("db_bandho_postgres: %s", err.Error())
	}

	return object.TRUE
}

// db_query_postgres - Execute SELECT query (synchronous)
func dbQueryPostgres(args ...object.Object) object.Object {
	if len(args) != 2 {
		return newError("db_query_postgres: wrong number of arguments. got=%d, want=2", len(args))
	}

	conn, ok := args[0].(*object.DBConnection)
	if !ok {
		return newError("db_query_postgres: first argument must be DB_CONNECTION, got %s", args[0].Type())
	}

	query, ok := args[1].(*object.String)
	if !ok {
		return newError("db_query_postgres: second argument must be STRING, got %s", args[1].Type())
	}

	result, err := Query(conn, query.Value)
	if err != nil {
		return newError("db_query_postgres: %s", err.Error())
	}

	if result.Error != nil {
		return result.Error
	}

	// Convert DBResult to Map for easier access in BanglaCode
	return dbResultToMap(result)
}

// db_exec_postgres - Execute INSERT/UPDATE/DELETE (synchronous)
func dbExecPostgres(args ...object.Object) object.Object {
	if len(args) != 2 {
		return newError("db_exec_postgres: wrong number of arguments. got=%d, want=2", len(args))
	}

	conn, ok := args[0].(*object.DBConnection)
	if !ok {
		return newError("db_exec_postgres: first argument must be DB_CONNECTION, got %s", args[0].Type())
	}

	query, ok := args[1].(*object.String)
	if !ok {
		return newError("db_exec_postgres: second argument must be STRING, got %s", args[1].Type())
	}

	result, err := Exec(conn, query.Value)
	if err != nil {
		return newError("db_exec_postgres: %s", err.Error())
	}

	if result.Error != nil {
		return result.Error
	}

	return dbResultToMap(result)
}

// db_proshno_postgres - Execute parameterized query (SQL injection safe)
func dbProshnoPostgres(args ...object.Object) object.Object {
	if len(args) != 3 {
		return newError("db_proshno_postgres: wrong number of arguments. got=%d, want=3", len(args))
	}

	conn, ok := args[0].(*object.DBConnection)
	if !ok {
		return newError("db_proshno_postgres: first argument must be DB_CONNECTION, got %s", args[0].Type())
	}

	query, ok := args[1].(*object.String)
	if !ok {
		return newError("db_proshno_postgres: second argument must be STRING, got %s", args[1].Type())
	}

	params, ok := args[2].(*object.Array)
	if !ok {
		return newError("db_proshno_postgres: third argument must be ARRAY, got %s", args[2].Type())
	}

	result, err := PreparedQuery(conn, query.Value, params.Elements)
	if err != nil {
		return newError("db_proshno_postgres: %s", err.Error())
	}

	if result.Error != nil {
		return result.Error
	}

	return dbResultToMap(result)
}

// db_query_async_postgres - Execute SELECT query (asynchronous)
func dbQueryAsyncPostgres(args ...object.Object) object.Object {
	if len(args) != 2 {
		return newError("db_query_async_postgres: wrong number of arguments. got=%d, want=2", len(args))
	}

	conn, ok := args[0].(*object.DBConnection)
	if !ok {
		return newError("db_query_async_postgres: first argument must be DB_CONNECTION, got %s", args[0].Type())
	}

	query, ok := args[1].(*object.String)
	if !ok {
		return newError("db_query_async_postgres: second argument must be STRING, got %s", args[1].Type())
	}

	// Create promise
	promise := object.CreatePromise()

	// Execute query asynchronously
	go func() {
		result, err := Query(conn, query.Value)
		if err != nil {
			object.RejectPromise(promise, newError("db_query_async_postgres: %s", err.Error()))
			return
		}

		if result.Error != nil {
			object.RejectPromise(promise, result.Error)
			return
		}

		object.ResolvePromise(promise, dbResultToMap(result))
	}()

	return promise
}

// db_exec_async_postgres - Execute INSERT/UPDATE/DELETE (asynchronous)
func dbExecAsyncPostgres(args ...object.Object) object.Object {
	if len(args) != 2 {
		return newError("db_exec_async_postgres: wrong number of arguments. got=%d, want=2", len(args))
	}

	conn, ok := args[0].(*object.DBConnection)
	if !ok {
		return newError("db_exec_async_postgres: first argument must be DB_CONNECTION, got %s", args[0].Type())
	}

	query, ok := args[1].(*object.String)
	if !ok {
		return newError("db_exec_async_postgres: second argument must be STRING, got %s", args[1].Type())
	}

	// Create promise
	promise := object.CreatePromise()

	// Execute query asynchronously
	go func() {
		result, err := Exec(conn, query.Value)
		if err != nil {
			object.RejectPromise(promise, newError("db_exec_async_postgres: %s", err.Error()))
			return
		}

		if result.Error != nil {
			object.RejectPromise(promise, result.Error)
			return
		}

		object.ResolvePromise(promise, dbResultToMap(result))
	}()

	return promise
}

// db_proshno_async_postgres - Execute parameterized query (async)
func dbProshnoAsyncPostgres(args ...object.Object) object.Object {
	if len(args) != 3 {
		return newError("db_proshno_async_postgres: wrong number of arguments. got=%d, want=3", len(args))
	}

	conn, ok := args[0].(*object.DBConnection)
	if !ok {
		return newError("db_proshno_async_postgres: first argument must be DB_CONNECTION, got %s", args[0].Type())
	}

	query, ok := args[1].(*object.String)
	if !ok {
		return newError("db_proshno_async_postgres: second argument must be STRING, got %s", args[1].Type())
	}

	params, ok := args[2].(*object.Array)
	if !ok {
		return newError("db_proshno_async_postgres: third argument must be ARRAY, got %s", args[2].Type())
	}

	// Create promise
	promise := object.CreatePromise()

	// Execute query asynchronously
	go func() {
		result, err := PreparedQuery(conn, query.Value, params.Elements)
		if err != nil {
			object.RejectPromise(promise, newError("db_proshno_async_postgres: %s", err.Error()))
			return
		}

		if result.Error != nil {
			object.RejectPromise(promise, result.Error)
			return
		}

		object.ResolvePromise(promise, dbResultToMap(result))
	}()

	return promise
}

// Transaction support

var transactions = make(map[string]*sql.Tx)
var transactionsMu sync.RWMutex
var txIDCounter int64

// db_transaction_shuru_postgres - Begin transaction
func dbTransactionShuruPostgres(args ...object.Object) object.Object {
	if len(args) != 1 {
		return newError("db_transaction_shuru_postgres: wrong number of arguments. got=%d, want=1", len(args))
	}

	conn, ok := args[0].(*object.DBConnection)
	if !ok {
		return newError("db_transaction_shuru_postgres: argument must be DB_CONNECTION, got %s", args[0].Type())
	}

	tx, err := BeginTransaction(conn)
	if err != nil {
		return newError("db_transaction_shuru_postgres: %s", err.Error())
	}

	// Generate transaction ID
	txIDCounter++
	txID := fmt.Sprintf("tx-%d", txIDCounter)

	// Store transaction
	transactionsMu.Lock()
	transactions[txID] = tx
	transactionsMu.Unlock()

	// Return transaction ID as string
	return &object.String{Value: txID}
}

// db_commit_postgres - Commit transaction
func dbCommitPostgres(args ...object.Object) object.Object {
	if len(args) != 1 {
		return newError("db_commit_postgres: wrong number of arguments. got=%d, want=1", len(args))
	}

	txID, ok := args[0].(*object.String)
	if !ok {
		return newError("db_commit_postgres: argument must be STRING (transaction ID), got %s", args[0].Type())
	}

	transactionsMu.Lock()
	tx, exists := transactions[txID.Value]
	if !exists {
		transactionsMu.Unlock()
		return newError("db_commit_postgres: transaction %s not found", txID.Value)
	}
	delete(transactions, txID.Value)
	transactionsMu.Unlock()

	if err := Commit(tx); err != nil {
		return newError("db_commit_postgres: %s", err.Error())
	}

	return object.TRUE
}

// db_rollback_postgres - Rollback transaction
func dbRollbackPostgres(args ...object.Object) object.Object {
	if len(args) != 1 {
		return newError("db_rollback_postgres: wrong number of arguments. got=%d, want=1", len(args))
	}

	txID, ok := args[0].(*object.String)
	if !ok {
		return newError("db_rollback_postgres: argument must be STRING (transaction ID), got %s", args[0].Type())
	}

	transactionsMu.Lock()
	tx, exists := transactions[txID.Value]
	if !exists {
		transactionsMu.Unlock()
		return newError("db_rollback_postgres: transaction %s not found", txID.Value)
	}
	delete(transactions, txID.Value)
	transactionsMu.Unlock()

	if err := Rollback(tx); err != nil {
		return newError("db_rollback_postgres: %s", err.Error())
	}

	return object.TRUE
}

// Bulk operations

// db_bulk_insert_postgres - Perform efficient bulk insert
// Usage: db_bulk_insert_postgres(conn, "users", ["name", "age"], [["Alice", 25], ["Bob", 30], ["Charlie", 35]])
func dbBulkInsertPostgres(args ...object.Object) object.Object {
	if len(args) != 4 {
		return newError("db_bulk_insert_postgres: wrong number of arguments. got=%d, want=4 (conn, table, columns, rows)", len(args))
	}

	conn, ok := args[0].(*object.DBConnection)
	if !ok {
		return newError("db_bulk_insert_postgres: first argument must be DB_CONNECTION, got %s", args[0].Type())
	}

	table, ok := args[1].(*object.String)
	if !ok {
		return newError("db_bulk_insert_postgres: second argument must be STRING (table name), got %s", args[1].Type())
	}

	columnsArray, ok := args[2].(*object.Array)
	if !ok {
		return newError("db_bulk_insert_postgres: third argument must be ARRAY (column names), got %s", args[2].Type())
	}

	// Convert column names array to string slice
	columns := make([]string, len(columnsArray.Elements))
	for i, elem := range columnsArray.Elements {
		colName, ok := elem.(*object.String)
		if !ok {
			return newError("db_bulk_insert_postgres: column name must be STRING, got %s", elem.Type())
		}
		columns[i] = colName.Value
	}

	rows, ok := args[3].(*object.Array)
	if !ok {
		return newError("db_bulk_insert_postgres: fourth argument must be ARRAY (rows), got %s", args[3].Type())
	}

	result, err := BulkInsert(conn, table.Value, columns, rows)
	if err != nil {
		return newError("db_bulk_insert_postgres: %s", err.Error())
	}

	if result.Error != nil {
		return result.Error
	}

	return dbResultToMap(result)
}

// Helper: Convert DBResult to Map for easier access
func dbResultToMap(result *object.DBResult) *object.Map {
	pairs := make(map[string]object.Object)

	// Convert rows to array
	rowsArray := &object.Array{Elements: make([]object.Object, len(result.Rows))}
	for i, row := range result.Rows {
		rowsArray.Elements[i] = &object.Map{Pairs: row}
	}

	pairs["rows"] = rowsArray
	pairs["rows_affected"] = &object.Number{Value: float64(result.RowsAffected)}
	pairs["last_insert_id"] = &object.Number{Value: float64(result.LastInsertID)}

	return &object.Map{Pairs: pairs}
}
