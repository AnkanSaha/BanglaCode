package mongodb

import (
	"BanglaCode/src/object"
	"fmt"
)

// Builtins holds all MongoDB built-in functions
var Builtins = make(map[string]*object.Builtin)

func init() {
	// Connection management
	registerBuiltin("db_jukto_mongodb", dbJuktoMongoDB)
	registerBuiltin("db_bandho_mongodb", dbBandhoMongoDB)

	// Document operations (synchronous)
	registerBuiltin("db_khojo_mongodb", dbKhojoMongoDB)
	registerBuiltin("db_dhokao_mongodb", dbDhokaoMongoDB)
	registerBuiltin("db_update_mongodb", dbUpdateMongoDB)
	registerBuiltin("db_mujhe_mongodb", dbMujheMongoDB)

	// Document operations (asynchronous)
	registerBuiltin("db_khojo_async_mongodb", dbKhojoAsyncMongoDB)
	registerBuiltin("db_dhokao_async_mongodb", dbDhokaoAsyncMongoDB)
	registerBuiltin("db_update_async_mongodb", dbUpdateAsyncMongoDB)
	registerBuiltin("db_mujhe_async_mongodb", dbMujheAsyncMongoDB)

	// Advanced operations
	registerBuiltin("db_aggregate_mongodb", dbAggregateMongoDB)
	registerBuiltin("db_findone_mongodb", dbFindOneMongoDB)
	registerBuiltin("db_count_mongodb", dbCountMongoDB)
	registerBuiltin("db_distinct_mongodb", dbDistinctMongoDB)
	registerBuiltin("db_khojo_options_mongodb", dbKhojoOptionsMongoDB)
	registerBuiltin("db_create_index_mongodb", dbCreateIndexMongoDB)
	registerBuiltin("db_insertmany_mongodb", dbInsertManyMongoDB)
}

func registerBuiltin(name string, fn object.BuiltinFunction) {
	Builtins[name] = &object.Builtin{Fn: fn}
}

func newError(format string, a ...interface{}) *object.Error {
	return &object.Error{Message: fmt.Sprintf(format, a...)}
}

// db_jukto_mongodb - Connect to MongoDB database
func dbJuktoMongoDB(args ...object.Object) object.Object {
	if len(args) != 1 {
		return newError("db_jukto_mongodb: wrong number of arguments. got=%d, want=1", len(args))
	}

	config, ok := args[0].(*object.Map)
	if !ok {
		return newError("db_jukto_mongodb: argument must be a map, got %s", args[0].Type())
	}

	conn, err := Connect(config)
	if err != nil {
		return newError("db_jukto_mongodb: %s", err.Error())
	}

	return conn
}

// db_bandho_mongodb - Close MongoDB connection
func dbBandhoMongoDB(args ...object.Object) object.Object {
	if len(args) != 1 {
		return newError("db_bandho_mongodb: wrong number of arguments. got=%d, want=1", len(args))
	}

	conn, ok := args[0].(*object.DBConnection)
	if !ok {
		return newError("db_bandho_mongodb: argument must be DB_CONNECTION, got %s", args[0].Type())
	}

	if err := Close(conn); err != nil {
		return newError("db_bandho_mongodb: %s", err.Error())
	}

	return object.TRUE
}

// db_khojo_mongodb - Find documents (synchronous)
// Usage: db_khojo_mongodb(conn, "users", {age: 25})
func dbKhojoMongoDB(args ...object.Object) object.Object {
	if len(args) != 3 {
		return newError("db_khojo_mongodb: wrong number of arguments. got=%d, want=3 (conn, collection, filter)", len(args))
	}

	conn, ok := args[0].(*object.DBConnection)
	if !ok {
		return newError("db_khojo_mongodb: first argument must be DB_CONNECTION, got %s", args[0].Type())
	}

	collectionName, ok := args[1].(*object.String)
	if !ok {
		return newError("db_khojo_mongodb: second argument must be STRING (collection name), got %s", args[1].Type())
	}

	filter, ok := args[2].(*object.Map)
	if !ok {
		return newError("db_khojo_mongodb: third argument must be MAP (filter), got %s", args[2].Type())
	}

	result, err := Find(conn, collectionName.Value, filter)
	if err != nil {
		return newError("db_khojo_mongodb: %s", err.Error())
	}

	if result.Error != nil {
		return result.Error
	}

	return dbResultToMap(result)
}

// db_dhokao_mongodb - Insert document (synchronous)
// Usage: db_dhokao_mongodb(conn, "users", {name: "Rahim", age: 30})
func dbDhokaoMongoDB(args ...object.Object) object.Object {
	if len(args) != 3 {
		return newError("db_dhokao_mongodb: wrong number of arguments. got=%d, want=3 (conn, collection, document)", len(args))
	}

	conn, ok := args[0].(*object.DBConnection)
	if !ok {
		return newError("db_dhokao_mongodb: first argument must be DB_CONNECTION, got %s", args[0].Type())
	}

	collectionName, ok := args[1].(*object.String)
	if !ok {
		return newError("db_dhokao_mongodb: second argument must be STRING (collection name), got %s", args[1].Type())
	}

	document, ok := args[2].(*object.Map)
	if !ok {
		return newError("db_dhokao_mongodb: third argument must be MAP (document), got %s", args[2].Type())
	}

	result, err := InsertOne(conn, collectionName.Value, document)
	if err != nil {
		return newError("db_dhokao_mongodb: %s", err.Error())
	}

	if result.Error != nil {
		return result.Error
	}

	return dbResultToMap(result)
}

// db_update_mongodb - Update documents (synchronous)
// Usage: db_update_mongodb(conn, "users", {name: "Rahim"}, {"$set": {age: 31}})
func dbUpdateMongoDB(args ...object.Object) object.Object {
	if len(args) != 4 {
		return newError("db_update_mongodb: wrong number of arguments. got=%d, want=4 (conn, collection, filter, update)", len(args))
	}

	conn, ok := args[0].(*object.DBConnection)
	if !ok {
		return newError("db_update_mongodb: first argument must be DB_CONNECTION, got %s", args[0].Type())
	}

	collectionName, ok := args[1].(*object.String)
	if !ok {
		return newError("db_update_mongodb: second argument must be STRING (collection name), got %s", args[1].Type())
	}

	filter, ok := args[2].(*object.Map)
	if !ok {
		return newError("db_update_mongodb: third argument must be MAP (filter), got %s", args[2].Type())
	}

	update, ok := args[3].(*object.Map)
	if !ok {
		return newError("db_update_mongodb: fourth argument must be MAP (update), got %s", args[3].Type())
	}

	result, err := UpdateMany(conn, collectionName.Value, filter, update)
	if err != nil {
		return newError("db_update_mongodb: %s", err.Error())
	}

	if result.Error != nil {
		return result.Error
	}

	return dbResultToMap(result)
}

// db_mujhe_mongodb - Delete documents (synchronous)
// Usage: db_mujhe_mongodb(conn, "users", {age: 30})
func dbMujheMongoDB(args ...object.Object) object.Object {
	if len(args) != 3 {
		return newError("db_mujhe_mongodb: wrong number of arguments. got=%d, want=3 (conn, collection, filter)", len(args))
	}

	conn, ok := args[0].(*object.DBConnection)
	if !ok {
		return newError("db_mujhe_mongodb: first argument must be DB_CONNECTION, got %s", args[0].Type())
	}

	collectionName, ok := args[1].(*object.String)
	if !ok {
		return newError("db_mujhe_mongodb: second argument must be STRING (collection name), got %s", args[1].Type())
	}

	filter, ok := args[2].(*object.Map)
	if !ok {
		return newError("db_mujhe_mongodb: third argument must be MAP (filter), got %s", args[2].Type())
	}

	result, err := DeleteMany(conn, collectionName.Value, filter)
	if err != nil {
		return newError("db_mujhe_mongodb: %s", err.Error())
	}

	if result.Error != nil {
		return result.Error
	}

	return dbResultToMap(result)
}

// Async functions

func dbKhojoAsyncMongoDB(args ...object.Object) object.Object {
	if len(args) != 3 {
		return newError("db_khojo_async_mongodb: wrong number of arguments. got=%d, want=3", len(args))
	}

	conn, ok := args[0].(*object.DBConnection)
	if !ok {
		return newError("db_khojo_async_mongodb: first argument must be DB_CONNECTION, got %s", args[0].Type())
	}

	collectionName, ok := args[1].(*object.String)
	if !ok {
		return newError("db_khojo_async_mongodb: second argument must be STRING, got %s", args[1].Type())
	}

	filter, ok := args[2].(*object.Map)
	if !ok {
		return newError("db_khojo_async_mongodb: third argument must be MAP, got %s", args[2].Type())
	}

	promise := object.CreatePromise()

	go func() {
		result, err := Find(conn, collectionName.Value, filter)
		if err != nil {
			object.RejectPromise(promise, newError("db_khojo_async_mongodb: %s", err.Error()))
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

func dbDhokaoAsyncMongoDB(args ...object.Object) object.Object {
	if len(args) != 3 {
		return newError("db_dhokao_async_mongodb: wrong number of arguments. got=%d, want=3", len(args))
	}

	conn, ok := args[0].(*object.DBConnection)
	if !ok {
		return newError("db_dhokao_async_mongodb: first argument must be DB_CONNECTION, got %s", args[0].Type())
	}

	collectionName, ok := args[1].(*object.String)
	if !ok {
		return newError("db_dhokao_async_mongodb: second argument must be STRING, got %s", args[1].Type())
	}

	document, ok := args[2].(*object.Map)
	if !ok {
		return newError("db_dhokao_async_mongodb: third argument must be MAP, got %s", args[2].Type())
	}

	promise := object.CreatePromise()

	go func() {
		result, err := InsertOne(conn, collectionName.Value, document)
		if err != nil {
			object.RejectPromise(promise, newError("db_dhokao_async_mongodb: %s", err.Error()))
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

func dbUpdateAsyncMongoDB(args ...object.Object) object.Object {
	if len(args) != 4 {
		return newError("db_update_async_mongodb: wrong number of arguments. got=%d, want=4", len(args))
	}

	conn, ok := args[0].(*object.DBConnection)
	if !ok {
		return newError("db_update_async_mongodb: first argument must be DB_CONNECTION, got %s", args[0].Type())
	}

	collectionName, ok := args[1].(*object.String)
	if !ok {
		return newError("db_update_async_mongodb: second argument must be STRING, got %s", args[1].Type())
	}

	filter, ok := args[2].(*object.Map)
	if !ok {
		return newError("db_update_async_mongodb: third argument must be MAP, got %s", args[2].Type())
	}

	update, ok := args[3].(*object.Map)
	if !ok {
		return newError("db_update_async_mongodb: fourth argument must be MAP, got %s", args[3].Type())
	}

	promise := object.CreatePromise()

	go func() {
		result, err := UpdateMany(conn, collectionName.Value, filter, update)
		if err != nil {
			object.RejectPromise(promise, newError("db_update_async_mongodb: %s", err.Error()))
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

func dbMujheAsyncMongoDB(args ...object.Object) object.Object {
	if len(args) != 3 {
		return newError("db_mujhe_async_mongodb: wrong number of arguments. got=%d, want=3", len(args))
	}

	conn, ok := args[0].(*object.DBConnection)
	if !ok {
		return newError("db_mujhe_async_mongodb: first argument must be DB_CONNECTION, got %s", args[0].Type())
	}

	collectionName, ok := args[1].(*object.String)
	if !ok {
		return newError("db_mujhe_async_mongodb: second argument must be STRING, got %s", args[1].Type())
	}

	filter, ok := args[2].(*object.Map)
	if !ok {
		return newError("db_mujhe_async_mongodb: third argument must be MAP, got %s", args[2].Type())
	}

	promise := object.CreatePromise()

	go func() {
		result, err := DeleteMany(conn, collectionName.Value, filter)
		if err != nil {
			object.RejectPromise(promise, newError("db_mujhe_async_mongodb: %s", err.Error()))
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

func dbResultToMap(result *object.DBResult) *object.Map {
	pairs := make(map[string]object.Object)

	rowsArray := &object.Array{Elements: make([]object.Object, len(result.Rows))}
	for i, row := range result.Rows {
		rowsArray.Elements[i] = &object.Map{Pairs: row}
	}

	pairs["rows"] = rowsArray
	pairs["rows_affected"] = &object.Number{Value: float64(result.RowsAffected)}

	return &object.Map{Pairs: pairs}
}

// Advanced MongoDB operations

// db_aggregate_mongodb - Execute aggregation pipeline
// Usage: db_aggregate_mongodb(conn, "users", [{$match: {age: {$gt: 25}}}, {$group: {_id: "$city", count: {$sum: 1}}}])
func dbAggregateMongoDB(args ...object.Object) object.Object {
	if len(args) != 3 {
		return newError("db_aggregate_mongodb: wrong number of arguments. got=%d, want=3 (conn, collection, pipeline)", len(args))
	}

	conn, ok := args[0].(*object.DBConnection)
	if !ok {
		return newError("db_aggregate_mongodb: first argument must be DB_CONNECTION, got %s", args[0].Type())
	}

	collectionName, ok := args[1].(*object.String)
	if !ok {
		return newError("db_aggregate_mongodb: second argument must be STRING (collection name), got %s", args[1].Type())
	}

	pipeline, ok := args[2].(*object.Array)
	if !ok {
		return newError("db_aggregate_mongodb: third argument must be ARRAY (pipeline), got %s", args[2].Type())
	}

	result, err := Aggregate(conn, collectionName.Value, pipeline)
	if err != nil {
		return newError("db_aggregate_mongodb: %s", err.Error())
	}

	if result.Error != nil {
		return result.Error
	}

	return dbResultToMap(result)
}

// db_findone_mongodb - Find a single document
// Usage: db_findone_mongodb(conn, "users", {name: "Rahim"})
func dbFindOneMongoDB(args ...object.Object) object.Object {
	if len(args) != 3 {
		return newError("db_findone_mongodb: wrong number of arguments. got=%d, want=3 (conn, collection, filter)", len(args))
	}

	conn, ok := args[0].(*object.DBConnection)
	if !ok {
		return newError("db_findone_mongodb: first argument must be DB_CONNECTION, got %s", args[0].Type())
	}

	collectionName, ok := args[1].(*object.String)
	if !ok {
		return newError("db_findone_mongodb: second argument must be STRING (collection name), got %s", args[1].Type())
	}

	filter, ok := args[2].(*object.Map)
	if !ok {
		return newError("db_findone_mongodb: third argument must be MAP (filter), got %s", args[2].Type())
	}

	result, err := FindOne(conn, collectionName.Value, filter)
	if err != nil {
		return newError("db_findone_mongodb: %s", err.Error())
	}

	if result.Error != nil {
		return result.Error
	}

	// Return single document or null if not found
	if len(result.Rows) == 0 {
		return object.NULL
	}

	return &object.Map{Pairs: result.Rows[0]}
}

// db_count_mongodb - Count documents matching filter
// Usage: db_count_mongodb(conn, "users", {age: {$gt: 25}})
func dbCountMongoDB(args ...object.Object) object.Object {
	if len(args) != 3 {
		return newError("db_count_mongodb: wrong number of arguments. got=%d, want=3 (conn, collection, filter)", len(args))
	}

	conn, ok := args[0].(*object.DBConnection)
	if !ok {
		return newError("db_count_mongodb: first argument must be DB_CONNECTION, got %s", args[0].Type())
	}

	collectionName, ok := args[1].(*object.String)
	if !ok {
		return newError("db_count_mongodb: second argument must be STRING (collection name), got %s", args[1].Type())
	}

	filter, ok := args[2].(*object.Map)
	if !ok {
		return newError("db_count_mongodb: third argument must be MAP (filter), got %s", args[2].Type())
	}

	count, err := Count(conn, collectionName.Value, filter)
	if err != nil {
		return newError("db_count_mongodb: %s", err.Error())
	}

	return &object.Number{Value: float64(count)}
}

// db_distinct_mongodb - Get distinct values for a field
// Usage: db_distinct_mongodb(conn, "users", "city", {})
func dbDistinctMongoDB(args ...object.Object) object.Object {
	if len(args) != 4 {
		return newError("db_distinct_mongodb: wrong number of arguments. got=%d, want=4 (conn, collection, field, filter)", len(args))
	}

	conn, ok := args[0].(*object.DBConnection)
	if !ok {
		return newError("db_distinct_mongodb: first argument must be DB_CONNECTION, got %s", args[0].Type())
	}

	collectionName, ok := args[1].(*object.String)
	if !ok {
		return newError("db_distinct_mongodb: second argument must be STRING (collection name), got %s", args[1].Type())
	}

	field, ok := args[2].(*object.String)
	if !ok {
		return newError("db_distinct_mongodb: third argument must be STRING (field name), got %s", args[2].Type())
	}

	filter, ok := args[3].(*object.Map)
	if !ok {
		return newError("db_distinct_mongodb: fourth argument must be MAP (filter), got %s", args[3].Type())
	}

	result, err := Distinct(conn, collectionName.Value, field.Value, filter)
	if err != nil {
		return newError("db_distinct_mongodb: %s", err.Error())
	}

	return result
}

// db_khojo_options_mongodb - Find with sort, limit, skip, projection
// Usage: db_khojo_options_mongodb(conn, "users", {age: {$gt: 25}}, {sort: {age: -1}, limit: 10})
func dbKhojoOptionsMongoDB(args ...object.Object) object.Object {
	if len(args) != 4 {
		return newError("db_khojo_options_mongodb: wrong number of arguments. got=%d, want=4 (conn, collection, filter, options)", len(args))
	}

	conn, ok := args[0].(*object.DBConnection)
	if !ok {
		return newError("db_khojo_options_mongodb: first argument must be DB_CONNECTION, got %s", args[0].Type())
	}

	collectionName, ok := args[1].(*object.String)
	if !ok {
		return newError("db_khojo_options_mongodb: second argument must be STRING (collection name), got %s", args[1].Type())
	}

	filter, ok := args[2].(*object.Map)
	if !ok {
		return newError("db_khojo_options_mongodb: third argument must be MAP (filter), got %s", args[2].Type())
	}

	options, ok := args[3].(*object.Map)
	if !ok {
		return newError("db_khojo_options_mongodb: fourth argument must be MAP (options), got %s", args[3].Type())
	}

	result, err := FindWithOptions(conn, collectionName.Value, filter, options)
	if err != nil {
		return newError("db_khojo_options_mongodb: %s", err.Error())
	}

	if result.Error != nil {
		return result.Error
	}

	return dbResultToMap(result)
}

// db_create_index_mongodb - Create an index
// Usage: db_create_index_mongodb(conn, "users", {email: 1}, sotti) // unique index
func dbCreateIndexMongoDB(args ...object.Object) object.Object {
	if len(args) != 4 {
		return newError("db_create_index_mongodb: wrong number of arguments. got=%d, want=4 (conn, collection, keys, unique)", len(args))
	}

	conn, ok := args[0].(*object.DBConnection)
	if !ok {
		return newError("db_create_index_mongodb: first argument must be DB_CONNECTION, got %s", args[0].Type())
	}

	collectionName, ok := args[1].(*object.String)
	if !ok {
		return newError("db_create_index_mongodb: second argument must be STRING (collection name), got %s", args[1].Type())
	}

	keys, ok := args[2].(*object.Map)
	if !ok {
		return newError("db_create_index_mongodb: third argument must be MAP (keys), got %s", args[2].Type())
	}

	unique, ok := args[3].(*object.Boolean)
	if !ok {
		return newError("db_create_index_mongodb: fourth argument must be BOOLEAN (unique), got %s", args[3].Type())
	}

	err := CreateIndex(conn, collectionName.Value, keys, unique.Value)
	if err != nil {
		return newError("db_create_index_mongodb: %s", err.Error())
	}

	return object.TRUE
}

// db_insertmany_mongodb - Insert multiple documents
// Usage: db_insertmany_mongodb(conn, "users", [{name: "A", age: 20}, {name: "B", age: 30}])
func dbInsertManyMongoDB(args ...object.Object) object.Object {
	if len(args) != 3 {
		return newError("db_insertmany_mongodb: wrong number of arguments. got=%d, want=3 (conn, collection, documents)", len(args))
	}

	conn, ok := args[0].(*object.DBConnection)
	if !ok {
		return newError("db_insertmany_mongodb: first argument must be DB_CONNECTION, got %s", args[0].Type())
	}

	collectionName, ok := args[1].(*object.String)
	if !ok {
		return newError("db_insertmany_mongodb: second argument must be STRING (collection name), got %s", args[1].Type())
	}

	documents, ok := args[2].(*object.Array)
	if !ok {
		return newError("db_insertmany_mongodb: third argument must be ARRAY (documents), got %s", args[2].Type())
	}

	result, err := InsertMany(conn, collectionName.Value, documents)
	if err != nil {
		return newError("db_insertmany_mongodb: %s", err.Error())
	}

	if result.Error != nil {
		return result.Error
	}

	return dbResultToMap(result)
}
