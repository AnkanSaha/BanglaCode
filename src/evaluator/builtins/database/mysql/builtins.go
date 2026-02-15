package mysql

import (
	"BanglaCode/src/object"
	"database/sql"
	"fmt"
	"sync"
)

// Builtins holds all MySQL built-in functions
var Builtins = make(map[string]*object.Builtin)

func init() {
	// Connection management
	registerBuiltin("db_jukto_mysql", dbJuktoMySQL)
	registerBuiltin("db_bandho_mysql", dbBandhoMySQL)

	// Query operations (synchronous)
	registerBuiltin("db_query_mysql", dbQueryMySQL)
	registerBuiltin("db_exec_mysql", dbExecMySQL)
	registerBuiltin("db_proshno_mysql", dbProshnoMySQL)

	// Query operations (asynchronous)
	registerBuiltin("db_query_async_mysql", dbQueryAsyncMySQL)
	registerBuiltin("db_exec_async_mysql", dbExecAsyncMySQL)
	registerBuiltin("db_proshno_async_mysql", dbProshnoAsyncMySQL)

	// Transaction support
	registerBuiltin("db_transaction_shuru_mysql", dbTransactionShuruMySQL)
	registerBuiltin("db_commit_mysql", dbCommitMySQL)
	registerBuiltin("db_rollback_mysql", dbRollbackMySQL)

	// Bulk operations
	registerBuiltin("db_bulk_insert_mysql", dbBulkInsertMySQL)
}

func registerBuiltin(name string, fn object.BuiltinFunction) {
	Builtins[name] = &object.Builtin{Fn: fn}
}

func newError(format string, a ...interface{}) *object.Error {
	return &object.Error{Message: fmt.Sprintf(format, a...)}
}

// db_jukto_mysql - Connect to MySQL database
func dbJuktoMySQL(args ...object.Object) object.Object {
	if len(args) != 1 {
		return newError("db_jukto_mysql: wrong number of arguments. got=%d, want=1", len(args))
	}

	config, ok := args[0].(*object.Map)
	if !ok {
		return newError("db_jukto_mysql: argument must be a map, got %s", args[0].Type())
	}

	conn, err := Connect(config)
	if err != nil {
		return newError("db_jukto_mysql: %s", err.Error())
	}

	return conn
}

// db_bandho_mysql - Close MySQL connection
func dbBandhoMySQL(args ...object.Object) object.Object {
	if len(args) != 1 {
		return newError("db_bandho_mysql: wrong number of arguments. got=%d, want=1", len(args))
	}

	conn, ok := args[0].(*object.DBConnection)
	if !ok {
		return newError("db_bandho_mysql: argument must be DB_CONNECTION, got %s", args[0].Type())
	}

	if err := Close(conn); err != nil {
		return newError("db_bandho_mysql: %s", err.Error())
	}

	return object.TRUE
}

// db_query_mysql - Execute SELECT query (synchronous)
func dbQueryMySQL(args ...object.Object) object.Object {
	if len(args) != 2 {
		return newError("db_query_mysql: wrong number of arguments. got=%d, want=2", len(args))
	}

	conn, ok := args[0].(*object.DBConnection)
	if !ok {
		return newError("db_query_mysql: first argument must be DB_CONNECTION, got %s", args[0].Type())
	}

	query, ok := args[1].(*object.String)
	if !ok {
		return newError("db_query_mysql: second argument must be STRING, got %s", args[1].Type())
	}

	result, err := Query(conn, query.Value)
	if err != nil {
		return newError("db_query_mysql: %s", err.Error())
	}

	if result.Error != nil {
		return result.Error
	}

	return dbResultToMap(result)
}

// db_exec_mysql - Execute INSERT/UPDATE/DELETE (synchronous)
func dbExecMySQL(args ...object.Object) object.Object {
	if len(args) != 2 {
		return newError("db_exec_mysql: wrong number of arguments. got=%d, want=2", len(args))
	}

	conn, ok := args[0].(*object.DBConnection)
	if !ok {
		return newError("db_exec_mysql: first argument must be DB_CONNECTION, got %s", args[0].Type())
	}

	query, ok := args[1].(*object.String)
	if !ok {
		return newError("db_exec_mysql: second argument must be STRING, got %s", args[1].Type())
	}

	result, err := Exec(conn, query.Value)
	if err != nil {
		return newError("db_exec_mysql: %s", err.Error())
	}

	if result.Error != nil {
		return result.Error
	}

	return dbResultToMap(result)
}

// db_proshno_mysql - Execute parameterized query (SQL injection safe)
func dbProshnoMySQL(args ...object.Object) object.Object {
	if len(args) != 3 {
		return newError("db_proshno_mysql: wrong number of arguments. got=%d, want=3", len(args))
	}

	conn, ok := args[0].(*object.DBConnection)
	if !ok {
		return newError("db_proshno_mysql: first argument must be DB_CONNECTION, got %s", args[0].Type())
	}

	query, ok := args[1].(*object.String)
	if !ok {
		return newError("db_proshno_mysql: second argument must be STRING, got %s", args[1].Type())
	}

	params, ok := args[2].(*object.Array)
	if !ok {
		return newError("db_proshno_mysql: third argument must be ARRAY, got %s", args[2].Type())
	}

	result, err := PreparedQuery(conn, query.Value, params.Elements)
	if err != nil {
		return newError("db_proshno_mysql: %s", err.Error())
	}

	if result.Error != nil {
		return result.Error
	}

	return dbResultToMap(result)
}

// Async functions

func dbQueryAsyncMySQL(args ...object.Object) object.Object {
	if len(args) != 2 {
		return newError("db_query_async_mysql: wrong number of arguments. got=%d, want=2", len(args))
	}

	conn, ok := args[0].(*object.DBConnection)
	if !ok {
		return newError("db_query_async_mysql: first argument must be DB_CONNECTION, got %s", args[0].Type())
	}

	query, ok := args[1].(*object.String)
	if !ok {
		return newError("db_query_async_mysql: second argument must be STRING, got %s", args[1].Type())
	}

	promise := object.CreatePromise()

	go func() {
		result, err := Query(conn, query.Value)
		if err != nil {
			object.RejectPromise(promise, newError("db_query_async_mysql: %s", err.Error()))
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

func dbExecAsyncMySQL(args ...object.Object) object.Object {
	if len(args) != 2 {
		return newError("db_exec_async_mysql: wrong number of arguments. got=%d, want=2", len(args))
	}

	conn, ok := args[0].(*object.DBConnection)
	if !ok {
		return newError("db_exec_async_mysql: first argument must be DB_CONNECTION, got %s", args[0].Type())
	}

	query, ok := args[1].(*object.String)
	if !ok {
		return newError("db_exec_async_mysql: second argument must be STRING, got %s", args[1].Type())
	}

	promise := object.CreatePromise()

	go func() {
		result, err := Exec(conn, query.Value)
		if err != nil {
			object.RejectPromise(promise, newError("db_exec_async_mysql: %s", err.Error()))
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

func dbProshnoAsyncMySQL(args ...object.Object) object.Object {
	if len(args) != 3 {
		return newError("db_proshno_async_mysql: wrong number of arguments. got=%d, want=3", len(args))
	}

	conn, ok := args[0].(*object.DBConnection)
	if !ok {
		return newError("db_proshno_async_mysql: first argument must be DB_CONNECTION, got %s", args[0].Type())
	}

	query, ok := args[1].(*object.String)
	if !ok {
		return newError("db_proshno_async_mysql: second argument must be STRING, got %s", args[1].Type())
	}

	params, ok := args[2].(*object.Array)
	if !ok {
		return newError("db_proshno_async_mysql: third argument must be ARRAY, got %s", args[2].Type())
	}

	promise := object.CreatePromise()

	go func() {
		result, err := PreparedQuery(conn, query.Value, params.Elements)
		if err != nil {
			object.RejectPromise(promise, newError("db_proshno_async_mysql: %s", err.Error()))
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

func dbTransactionShuruMySQL(args ...object.Object) object.Object {
	if len(args) != 1 {
		return newError("db_transaction_shuru_mysql: wrong number of arguments. got=%d, want=1", len(args))
	}

	conn, ok := args[0].(*object.DBConnection)
	if !ok {
		return newError("db_transaction_shuru_mysql: argument must be DB_CONNECTION, got %s", args[0].Type())
	}

	tx, err := BeginTransaction(conn)
	if err != nil {
		return newError("db_transaction_shuru_mysql: %s", err.Error())
	}

	txIDCounter++
	txID := fmt.Sprintf("tx-mysql-%d", txIDCounter)

	transactionsMu.Lock()
	transactions[txID] = tx
	transactionsMu.Unlock()

	return &object.String{Value: txID}
}

func dbCommitMySQL(args ...object.Object) object.Object {
	if len(args) != 1 {
		return newError("db_commit_mysql: wrong number of arguments. got=%d, want=1", len(args))
	}

	txID, ok := args[0].(*object.String)
	if !ok {
		return newError("db_commit_mysql: argument must be STRING (transaction ID), got %s", args[0].Type())
	}

	transactionsMu.Lock()
	tx, exists := transactions[txID.Value]
	if !exists {
		transactionsMu.Unlock()
		return newError("db_commit_mysql: transaction %s not found", txID.Value)
	}
	delete(transactions, txID.Value)
	transactionsMu.Unlock()

	if err := Commit(tx); err != nil {
		return newError("db_commit_mysql: %s", err.Error())
	}

	return object.TRUE
}

func dbRollbackMySQL(args ...object.Object) object.Object {
	if len(args) != 1 {
		return newError("db_rollback_mysql: wrong number of arguments. got=%d, want=1", len(args))
	}

	txID, ok := args[0].(*object.String)
	if !ok {
		return newError("db_rollback_mysql: argument must be STRING (transaction ID), got %s", args[0].Type())
	}

	transactionsMu.Lock()
	tx, exists := transactions[txID.Value]
	if !exists {
		transactionsMu.Unlock()
		return newError("db_rollback_mysql: transaction %s not found", txID.Value)
	}
	delete(transactions, txID.Value)
	transactionsMu.Unlock()

	if err := Rollback(tx); err != nil {
		return newError("db_rollback_mysql: %s", err.Error())
	}

	return object.TRUE
}

// Bulk operations

// db_bulk_insert_mysql - Perform efficient bulk insert
// Usage: db_bulk_insert_mysql(conn, "users", ["name", "age"], [["Alice", 25], ["Bob", 30], ["Charlie", 35]])
func dbBulkInsertMySQL(args ...object.Object) object.Object {
	if len(args) != 4 {
		return newError("db_bulk_insert_mysql: wrong number of arguments. got=%d, want=4 (conn, table, columns, rows)", len(args))
	}

	conn, ok := args[0].(*object.DBConnection)
	if !ok {
		return newError("db_bulk_insert_mysql: first argument must be DB_CONNECTION, got %s", args[0].Type())
	}

	table, ok := args[1].(*object.String)
	if !ok {
		return newError("db_bulk_insert_mysql: second argument must be STRING (table name), got %s", args[1].Type())
	}

	columnsArray, ok := args[2].(*object.Array)
	if !ok {
		return newError("db_bulk_insert_mysql: third argument must be ARRAY (column names), got %s", args[2].Type())
	}

	// Convert column names array to string slice
	columns := make([]string, len(columnsArray.Elements))
	for i, elem := range columnsArray.Elements {
		colName, ok := elem.(*object.String)
		if !ok {
			return newError("db_bulk_insert_mysql: column name must be STRING, got %s", elem.Type())
		}
		columns[i] = colName.Value
	}

	rows, ok := args[3].(*object.Array)
	if !ok {
		return newError("db_bulk_insert_mysql: fourth argument must be ARRAY (rows), got %s", args[3].Type())
	}

	result, err := BulkInsert(conn, table.Value, columns, rows)
	if err != nil {
		return newError("db_bulk_insert_mysql: %s", err.Error())
	}

	if result.Error != nil {
		return result.Error
	}

	return dbResultToMap(result)
}

func dbResultToMap(result *object.DBResult) *object.Map {
	pairs := make(map[string]object.Object)

	rowsArray := &object.Array{Elements: make([]object.Object, len(result.Rows))}
	for i, row := range result.Rows {
		rowsArray.Elements[i] = &object.Map{Pairs: row}
	}

	pairs["rows"] = rowsArray
	pairs["rows_affected"] = &object.Number{Value: float64(result.RowsAffected)}
	pairs["last_insert_id"] = &object.Number{Value: float64(result.LastInsertID)}

	return &object.Map{Pairs: pairs}
}
