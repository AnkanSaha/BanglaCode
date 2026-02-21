package mongodb

import (
	"BanglaCode/src/object"
)

// Async operations with promise-based responses

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

// dbResultToMap converts a DBResult to a BanglaCode Map
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
