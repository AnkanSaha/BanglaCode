package redis

import (
	"BanglaCode/src/object"
)

// Advanced Redis built-in function implementations (Counter, Pub/Sub, Utility operations)

// Counter operations

// db_incr_redis - Increment a counter by 1
// Usage: db_incr_redis(conn, "page_views")
func dbIncrRedis(args ...object.Object) object.Object {
	if len(args) != 2 {
		return newError("db_incr_redis: wrong number of arguments. got=%d, want=2 (conn, key)", len(args))
	}

	conn, ok := args[0].(*object.DBConnection)
	if !ok {
		return newError("db_incr_redis: first argument must be DB_CONNECTION, got %s", args[0].Type())
	}

	key, ok := args[1].(*object.String)
	if !ok {
		return newError("db_incr_redis: second argument must be STRING (key), got %s", args[1].Type())
	}

	value, err := Incr(conn, key.Value)
	if err != nil {
		return newError("db_incr_redis: %s", err.Error())
	}

	return &object.Number{Value: float64(value)}
}

// db_decr_redis - Decrement a counter by 1
// Usage: db_decr_redis(conn, "items_left")
func dbDecrRedis(args ...object.Object) object.Object {
	if len(args) != 2 {
		return newError("db_decr_redis: wrong number of arguments. got=%d, want=2 (conn, key)", len(args))
	}

	conn, ok := args[0].(*object.DBConnection)
	if !ok {
		return newError("db_decr_redis: first argument must be DB_CONNECTION, got %s", args[0].Type())
	}

	key, ok := args[1].(*object.String)
	if !ok {
		return newError("db_decr_redis: second argument must be STRING (key), got %s", args[1].Type())
	}

	value, err := Decr(conn, key.Value)
	if err != nil {
		return newError("db_decr_redis: %s", err.Error())
	}

	return &object.Number{Value: float64(value)}
}

// db_incrby_redis - Increment a counter by specific amount
// Usage: db_incrby_redis(conn, "score", 10)
func dbIncrByRedis(args ...object.Object) object.Object {
	if len(args) != 3 {
		return newError("db_incrby_redis: wrong number of arguments. got=%d, want=3 (conn, key, increment)", len(args))
	}

	conn, ok := args[0].(*object.DBConnection)
	if !ok {
		return newError("db_incrby_redis: first argument must be DB_CONNECTION, got %s", args[0].Type())
	}

	key, ok := args[1].(*object.String)
	if !ok {
		return newError("db_incrby_redis: second argument must be STRING (key), got %s", args[1].Type())
	}

	increment, ok := args[2].(*object.Number)
	if !ok {
		return newError("db_incrby_redis: third argument must be NUMBER (increment), got %s", args[2].Type())
	}

	value, err := IncrBy(conn, key.Value, int64(increment.Value))
	if err != nil {
		return newError("db_incrby_redis: %s", err.Error())
	}

	return &object.Number{Value: float64(value)}
}

// db_decrby_redis - Decrement a counter by specific amount
// Usage: db_decrby_redis(conn, "balance", 50)
func dbDecrByRedis(args ...object.Object) object.Object {
	if len(args) != 3 {
		return newError("db_decrby_redis: wrong number of arguments. got=%d, want=3 (conn, key, decrement)", len(args))
	}

	conn, ok := args[0].(*object.DBConnection)
	if !ok {
		return newError("db_decrby_redis: first argument must be DB_CONNECTION, got %s", args[0].Type())
	}

	key, ok := args[1].(*object.String)
	if !ok {
		return newError("db_decrby_redis: second argument must be STRING (key), got %s", args[1].Type())
	}

	decrement, ok := args[2].(*object.Number)
	if !ok {
		return newError("db_decrby_redis: third argument must be NUMBER (decrement), got %s", args[2].Type())
	}

	value, err := DecrBy(conn, key.Value, int64(decrement.Value))
	if err != nil {
		return newError("db_decrby_redis: %s", err.Error())
	}

	return &object.Number{Value: float64(value)}
}

// db_incrbyfloat_redis - Increment a counter by float amount
// Usage: db_incrbyfloat_redis(conn, "temperature", 0.5)
func dbIncrByFloatRedis(args ...object.Object) object.Object {
	if len(args) != 3 {
		return newError("db_incrbyfloat_redis: wrong number of arguments. got=%d, want=3 (conn, key, increment)", len(args))
	}

	conn, ok := args[0].(*object.DBConnection)
	if !ok {
		return newError("db_incrbyfloat_redis: first argument must be DB_CONNECTION, got %s", args[0].Type())
	}

	key, ok := args[1].(*object.String)
	if !ok {
		return newError("db_incrbyfloat_redis: second argument must be STRING (key), got %s", args[1].Type())
	}

	increment, ok := args[2].(*object.Number)
	if !ok {
		return newError("db_incrbyfloat_redis: third argument must be NUMBER (increment), got %s", args[2].Type())
	}

	value, err := IncrByFloat(conn, key.Value, increment.Value)
	if err != nil {
		return newError("db_incrbyfloat_redis: %s", err.Error())
	}

	return &object.Number{Value: value}
}

// Pub/Sub operations

// db_publish_redis - Publish a message to a channel
// Usage: db_publish_redis(conn, "notifications", "New message!")
func dbPublishRedis(args ...object.Object) object.Object {
	if len(args) != 3 {
		return newError("db_publish_redis: wrong number of arguments. got=%d, want=3 (conn, channel, message)", len(args))
	}

	conn, ok := args[0].(*object.DBConnection)
	if !ok {
		return newError("db_publish_redis: first argument must be DB_CONNECTION, got %s", args[0].Type())
	}

	channel, ok := args[1].(*object.String)
	if !ok {
		return newError("db_publish_redis: second argument must be STRING (channel), got %s", args[1].Type())
	}

	message, ok := args[2].(*object.String)
	if !ok {
		return newError("db_publish_redis: third argument must be STRING (message), got %s", args[2].Type())
	}

	subscribers, err := Publish(conn, channel.Value, message.Value)
	if err != nil {
		return newError("db_publish_redis: %s", err.Error())
	}

	return &object.Number{Value: float64(subscribers)}
}

// Utility operations

// db_ttl_redis - Get time to live of a key
// Usage: db_ttl_redis(conn, "session:123")
func dbTTLRedis(args ...object.Object) object.Object {
	if len(args) != 2 {
		return newError("db_ttl_redis: wrong number of arguments. got=%d, want=2 (conn, key)", len(args))
	}

	conn, ok := args[0].(*object.DBConnection)
	if !ok {
		return newError("db_ttl_redis: first argument must be DB_CONNECTION, got %s", args[0].Type())
	}

	key, ok := args[1].(*object.String)
	if !ok {
		return newError("db_ttl_redis: second argument must be STRING (key), got %s", args[1].Type())
	}

	ttl, err := TTL(conn, key.Value)
	if err != nil {
		return newError("db_ttl_redis: %s", err.Error())
	}

	return &object.Number{Value: float64(ttl)}
}

// db_persist_redis - Remove expiration from a key
// Usage: db_persist_redis(conn, "permanent_data")
func dbPersistRedis(args ...object.Object) object.Object {
	if len(args) != 2 {
		return newError("db_persist_redis: wrong number of arguments. got=%d, want=2 (conn, key)", len(args))
	}

	conn, ok := args[0].(*object.DBConnection)
	if !ok {
		return newError("db_persist_redis: first argument must be DB_CONNECTION, got %s", args[0].Type())
	}

	key, ok := args[1].(*object.String)
	if !ok {
		return newError("db_persist_redis: second argument must be STRING (key), got %s", args[1].Type())
	}

	success, err := Persist(conn, key.Value)
	if err != nil {
		return newError("db_persist_redis: %s", err.Error())
	}

	return object.NativeBoolToBooleanObject(success)
}

// db_exists_redis - Check if keys exist
// Usage: db_exists_redis(conn, ["key1", "key2", "key3"])
func dbExistsRedis(args ...object.Object) object.Object {
	if len(args) != 2 {
		return newError("db_exists_redis: wrong number of arguments. got=%d, want=2 (conn, keys)", len(args))
	}

	conn, ok := args[0].(*object.DBConnection)
	if !ok {
		return newError("db_exists_redis: first argument must be DB_CONNECTION, got %s", args[0].Type())
	}

	keysArray, ok := args[1].(*object.Array)
	if !ok {
		return newError("db_exists_redis: second argument must be ARRAY (keys), got %s", args[1].Type())
	}

	// Convert array to string slice
	keys := make([]string, len(keysArray.Elements))
	for i, elem := range keysArray.Elements {
		key, ok := elem.(*object.String)
		if !ok {
			return newError("db_exists_redis: array element must be STRING, got %s", elem.Type())
		}
		keys[i] = key.Value
	}

	count, err := Exists(conn, keys...)
	if err != nil {
		return newError("db_exists_redis: %s", err.Error())
	}

	return &object.Number{Value: float64(count)}
}

// db_keys_redis - Get all keys matching a pattern
// Usage: db_keys_redis(conn, "user:*")
func dbKeysRedis(args ...object.Object) object.Object {
	if len(args) != 2 {
		return newError("db_keys_redis: wrong number of arguments. got=%d, want=2 (conn, pattern)", len(args))
	}

	conn, ok := args[0].(*object.DBConnection)
	if !ok {
		return newError("db_keys_redis: first argument must be DB_CONNECTION, got %s", args[0].Type())
	}

	pattern, ok := args[1].(*object.String)
	if !ok {
		return newError("db_keys_redis: second argument must be STRING (pattern), got %s", args[1].Type())
	}

	keys, err := Keys(conn, pattern.Value)
	if err != nil {
		return newError("db_keys_redis: %s", err.Error())
	}

	// Convert to array
	elements := make([]object.Object, len(keys))
	for i, key := range keys {
		elements[i] = &object.String{Value: key}
	}

	return &object.Array{Elements: elements}
}

// Register counter, pub/sub, and utility functions
func init() {
	// Counter operations
	registerBuiltin("db_incr_redis", dbIncrRedis)
	registerBuiltin("db_decr_redis", dbDecrRedis)
	registerBuiltin("db_incrby_redis", dbIncrByRedis)
	registerBuiltin("db_decrby_redis", dbDecrByRedis)
	registerBuiltin("db_incrbyfloat_redis", dbIncrByFloatRedis)
	// Pub/Sub operations
	registerBuiltin("db_publish_redis", dbPublishRedis)
	// Utility operations
	registerBuiltin("db_ttl_redis", dbTTLRedis)
	registerBuiltin("db_persist_redis", dbPersistRedis)
	registerBuiltin("db_exists_redis", dbExistsRedis)
	registerBuiltin("db_keys_redis", dbKeysRedis)
}
