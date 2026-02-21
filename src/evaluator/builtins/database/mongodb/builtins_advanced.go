package mongodb

import (
	"BanglaCode/src/object"
)

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
