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
