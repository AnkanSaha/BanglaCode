package redis

import (
	"BanglaCode/src/object"
	"fmt"
	"time"
)

// Builtins holds all Redis built-in functions
var Builtins = make(map[string]*object.Builtin)

func init() {
	// Connection management
	registerBuiltin("db_jukto_redis", dbJuktoRedis)
	registerBuiltin("db_bandho_redis", dbBandhoRedis)

	// String operations (synchronous)
	registerBuiltin("db_set_redis", dbSetRedis)
	registerBuiltin("db_get_redis", dbGetRedis)
	registerBuiltin("db_del_redis", dbDelRedis)
	registerBuiltin("db_expire_redis", dbExpireRedis)

	// String operations (asynchronous)
	registerBuiltin("db_set_async_redis", dbSetAsyncRedis)
	registerBuiltin("db_get_async_redis", dbGetAsyncRedis)

	// List operations
	registerBuiltin("db_lpush_redis", dbLPushRedis)
	registerBuiltin("db_rpush_redis", dbRPushRedis)
	registerBuiltin("db_lpop_redis", dbLPopRedis)
	registerBuiltin("db_rpop_redis", dbRPopRedis)

	// Hash operations
	registerBuiltin("db_hset_redis", dbHSetRedis)
	registerBuiltin("db_hget_redis", dbHGetRedis)
	registerBuiltin("db_hgetall_redis", dbHGetAllRedis)

	// Sorted Sets operations
	registerBuiltin("db_zadd_redis", dbZAddRedis)
	registerBuiltin("db_zrange_redis", dbZRangeRedis)
	registerBuiltin("db_zrank_redis", dbZRankRedis)
	registerBuiltin("db_zscore_redis", dbZScoreRedis)
	registerBuiltin("db_zrem_redis", dbZRemRedis)

	// Sets operations
	registerBuiltin("db_sadd_redis", dbSAddRedis)
	registerBuiltin("db_smembers_redis", dbSMembersRedis)
	registerBuiltin("db_sismember_redis", dbSIsMemberRedis)
	registerBuiltin("db_srem_redis", dbSRemRedis)
	registerBuiltin("db_sinter_redis", dbSInterRedis)
	registerBuiltin("db_sunion_redis", dbSUnionRedis)
	registerBuiltin("db_sdiff_redis", dbSDiffRedis)

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

func registerBuiltin(name string, fn object.BuiltinFunction) {
	Builtins[name] = &object.Builtin{Fn: fn}
}

func newError(format string, a ...interface{}) *object.Error {
	return &object.Error{Message: fmt.Sprintf(format, a...)}
}

// db_jukto_redis - Connect to Redis
func dbJuktoRedis(args ...object.Object) object.Object {
	if len(args) != 1 {
		return newError("db_jukto_redis: wrong number of arguments. got=%d, want=1", len(args))
	}

	config, ok := args[0].(*object.Map)
	if !ok {
		return newError("db_jukto_redis: argument must be a map, got %s", args[0].Type())
	}

	conn, err := Connect(config)
	if err != nil {
		return newError("db_jukto_redis: %s", err.Error())
	}

	return conn
}

// db_bandho_redis - Close Redis connection
func dbBandhoRedis(args ...object.Object) object.Object {
	if len(args) != 1 {
		return newError("db_bandho_redis: wrong number of arguments. got=%d, want=1", len(args))
	}

	conn, ok := args[0].(*object.DBConnection)
	if !ok {
		return newError("db_bandho_redis: argument must be DB_CONNECTION, got %s", args[0].Type())
	}

	if err := Close(conn); err != nil {
		return newError("db_bandho_redis: %s", err.Error())
	}

	return object.TRUE
}

// db_set_redis - Set key-value (synchronous)
// Usage: db_set_redis(conn, "key", "value") or db_set_redis(conn, "key", "value", ttl_seconds)
func dbSetRedis(args ...object.Object) object.Object {
	if len(args) < 3 || len(args) > 4 {
		return newError("db_set_redis: wrong number of arguments. got=%d, want=3 or 4", len(args))
	}

	conn, ok := args[0].(*object.DBConnection)
	if !ok {
		return newError("db_set_redis: first argument must be DB_CONNECTION, got %s", args[0].Type())
	}

	key, ok := args[1].(*object.String)
	if !ok {
		return newError("db_set_redis: second argument must be STRING (key), got %s", args[1].Type())
	}

	value, ok := args[2].(*object.String)
	if !ok {
		return newError("db_set_redis: third argument must be STRING (value), got %s", args[2].Type())
	}

	// Optional TTL
	var expiration time.Duration
	if len(args) == 4 {
		ttl, ok := args[3].(*object.Number)
		if !ok {
			return newError("db_set_redis: fourth argument must be NUMBER (TTL in seconds), got %s", args[3].Type())
		}
		expiration = time.Duration(ttl.Value) * time.Second
	}

	if err := Set(conn, key.Value, value.Value, expiration); err != nil {
		return newError("db_set_redis: %s", err.Error())
	}

	return object.TRUE
}

// db_get_redis - Get value by key (synchronous)
func dbGetRedis(args ...object.Object) object.Object {
	if len(args) != 2 {
		return newError("db_get_redis: wrong number of arguments. got=%d, want=2", len(args))
	}

	conn, ok := args[0].(*object.DBConnection)
	if !ok {
		return newError("db_get_redis: first argument must be DB_CONNECTION, got %s", args[0].Type())
	}

	key, ok := args[1].(*object.String)
	if !ok {
		return newError("db_get_redis: second argument must be STRING (key), got %s", args[1].Type())
	}

	value, err := Get(conn, key.Value)
	if err != nil {
		return newError("db_get_redis: %s", err.Error())
	}

	return &object.String{Value: value}
}

// db_del_redis - Delete key
func dbDelRedis(args ...object.Object) object.Object {
	if len(args) != 2 {
		return newError("db_del_redis: wrong number of arguments. got=%d, want=2", len(args))
	}

	conn, ok := args[0].(*object.DBConnection)
	if !ok {
		return newError("db_del_redis: first argument must be DB_CONNECTION, got %s", args[0].Type())
	}

	key, ok := args[1].(*object.String)
	if !ok {
		return newError("db_del_redis: second argument must be STRING (key), got %s", args[1].Type())
	}

	if err := Del(conn, key.Value); err != nil {
		return newError("db_del_redis: %s", err.Error())
	}

	return object.TRUE
}

// db_expire_redis - Set expiration on key
func dbExpireRedis(args ...object.Object) object.Object {
	if len(args) != 3 {
		return newError("db_expire_redis: wrong number of arguments. got=%d, want=3", len(args))
	}

	conn, ok := args[0].(*object.DBConnection)
	if !ok {
		return newError("db_expire_redis: first argument must be DB_CONNECTION, got %s", args[0].Type())
	}

	key, ok := args[1].(*object.String)
	if !ok {
		return newError("db_expire_redis: second argument must be STRING (key), got %s", args[1].Type())
	}

	seconds, ok := args[2].(*object.Number)
	if !ok {
		return newError("db_expire_redis: third argument must be NUMBER (seconds), got %s", args[2].Type())
	}

	if err := Expire(conn, key.Value, int(seconds.Value)); err != nil {
		return newError("db_expire_redis: %s", err.Error())
	}

	return object.TRUE
}

// Async functions

func dbSetAsyncRedis(args ...object.Object) object.Object {
	if len(args) < 3 || len(args) > 4 {
		return newError("db_set_async_redis: wrong number of arguments. got=%d, want=3 or 4", len(args))
	}

	conn, ok := args[0].(*object.DBConnection)
	if !ok {
		return newError("db_set_async_redis: first argument must be DB_CONNECTION, got %s", args[0].Type())
	}

	key, ok := args[1].(*object.String)
	if !ok {
		return newError("db_set_async_redis: second argument must be STRING, got %s", args[1].Type())
	}

	value, ok := args[2].(*object.String)
	if !ok {
		return newError("db_set_async_redis: third argument must be STRING, got %s", args[2].Type())
	}

	var expiration time.Duration
	if len(args) == 4 {
		ttl, ok := args[3].(*object.Number)
		if !ok {
			return newError("db_set_async_redis: fourth argument must be NUMBER, got %s", args[3].Type())
		}
		expiration = time.Duration(ttl.Value) * time.Second
	}

	promise := object.CreatePromise()

	go func() {
		if err := Set(conn, key.Value, value.Value, expiration); err != nil {
			object.RejectPromise(promise, newError("db_set_async_redis: %s", err.Error()))
			return
		}
		object.ResolvePromise(promise, object.TRUE)
	}()

	return promise
}

func dbGetAsyncRedis(args ...object.Object) object.Object {
	if len(args) != 2 {
		return newError("db_get_async_redis: wrong number of arguments. got=%d, want=2", len(args))
	}

	conn, ok := args[0].(*object.DBConnection)
	if !ok {
		return newError("db_get_async_redis: first argument must be DB_CONNECTION, got %s", args[0].Type())
	}

	key, ok := args[1].(*object.String)
	if !ok {
		return newError("db_get_async_redis: second argument must be STRING, got %s", args[1].Type())
	}

	promise := object.CreatePromise()

	go func() {
		value, err := Get(conn, key.Value)
		if err != nil {
			object.RejectPromise(promise, newError("db_get_async_redis: %s", err.Error()))
			return
		}
		object.ResolvePromise(promise, &object.String{Value: value})
	}()

	return promise
}

// List operations

func dbLPushRedis(args ...object.Object) object.Object {
	if len(args) != 3 {
		return newError("db_lpush_redis: wrong number of arguments. got=%d, want=3", len(args))
	}

	conn, ok := args[0].(*object.DBConnection)
	if !ok {
		return newError("db_lpush_redis: first argument must be DB_CONNECTION, got %s", args[0].Type())
	}

	key, ok := args[1].(*object.String)
	if !ok {
		return newError("db_lpush_redis: second argument must be STRING, got %s", args[1].Type())
	}

	value, ok := args[2].(*object.String)
	if !ok {
		return newError("db_lpush_redis: third argument must be STRING, got %s", args[2].Type())
	}

	length, err := LPush(conn, key.Value, value.Value)
	if err != nil {
		return newError("db_lpush_redis: %s", err.Error())
	}

	return &object.Number{Value: float64(length)}
}

func dbRPushRedis(args ...object.Object) object.Object {
	if len(args) != 3 {
		return newError("db_rpush_redis: wrong number of arguments. got=%d, want=3", len(args))
	}

	conn, ok := args[0].(*object.DBConnection)
	if !ok {
		return newError("db_rpush_redis: first argument must be DB_CONNECTION, got %s", args[0].Type())
	}

	key, ok := args[1].(*object.String)
	if !ok {
		return newError("db_rpush_redis: second argument must be STRING, got %s", args[1].Type())
	}

	value, ok := args[2].(*object.String)
	if !ok {
		return newError("db_rpush_redis: third argument must be STRING, got %s", args[2].Type())
	}

	length, err := RPush(conn, key.Value, value.Value)
	if err != nil {
		return newError("db_rpush_redis: %s", err.Error())
	}

	return &object.Number{Value: float64(length)}
}

func dbLPopRedis(args ...object.Object) object.Object {
	if len(args) != 2 {
		return newError("db_lpop_redis: wrong number of arguments. got=%d, want=2", len(args))
	}

	conn, ok := args[0].(*object.DBConnection)
	if !ok {
		return newError("db_lpop_redis: first argument must be DB_CONNECTION, got %s", args[0].Type())
	}

	key, ok := args[1].(*object.String)
	if !ok {
		return newError("db_lpop_redis: second argument must be STRING, got %s", args[1].Type())
	}

	value, err := LPop(conn, key.Value)
	if err != nil {
		return newError("db_lpop_redis: %s", err.Error())
	}

	return &object.String{Value: value}
}

func dbRPopRedis(args ...object.Object) object.Object {
	if len(args) != 2 {
		return newError("db_rpop_redis: wrong number of arguments. got=%d, want=2", len(args))
	}

	conn, ok := args[0].(*object.DBConnection)
	if !ok {
		return newError("db_rpop_redis: first argument must be DB_CONNECTION, got %s", args[0].Type())
	}

	key, ok := args[1].(*object.String)
	if !ok {
		return newError("db_rpop_redis: second argument must be STRING, got %s", args[1].Type())
	}

	value, err := RPop(conn, key.Value)
	if err != nil {
		return newError("db_rpop_redis: %s", err.Error())
	}

	return &object.String{Value: value}
}

// Hash operations

func dbHSetRedis(args ...object.Object) object.Object {
	if len(args) != 4 {
		return newError("db_hset_redis: wrong number of arguments. got=%d, want=4", len(args))
	}

	conn, ok := args[0].(*object.DBConnection)
	if !ok {
		return newError("db_hset_redis: first argument must be DB_CONNECTION, got %s", args[0].Type())
	}

	key, ok := args[1].(*object.String)
	if !ok {
		return newError("db_hset_redis: second argument must be STRING, got %s", args[1].Type())
	}

	field, ok := args[2].(*object.String)
	if !ok {
		return newError("db_hset_redis: third argument must be STRING, got %s", args[2].Type())
	}

	value, ok := args[3].(*object.String)
	if !ok {
		return newError("db_hset_redis: fourth argument must be STRING, got %s", args[3].Type())
	}

	if err := HSet(conn, key.Value, field.Value, value.Value); err != nil {
		return newError("db_hset_redis: %s", err.Error())
	}

	return object.TRUE
}

func dbHGetRedis(args ...object.Object) object.Object {
	if len(args) != 3 {
		return newError("db_hget_redis: wrong number of arguments. got=%d, want=3", len(args))
	}

	conn, ok := args[0].(*object.DBConnection)
	if !ok {
		return newError("db_hget_redis: first argument must be DB_CONNECTION, got %s", args[0].Type())
	}

	key, ok := args[1].(*object.String)
	if !ok {
		return newError("db_hget_redis: second argument must be STRING, got %s", args[1].Type())
	}

	field, ok := args[2].(*object.String)
	if !ok {
		return newError("db_hget_redis: third argument must be STRING, got %s", args[2].Type())
	}

	value, err := HGet(conn, key.Value, field.Value)
	if err != nil {
		return newError("db_hget_redis: %s", err.Error())
	}

	return &object.String{Value: value}
}

func dbHGetAllRedis(args ...object.Object) object.Object {
	if len(args) != 2 {
		return newError("db_hgetall_redis: wrong number of arguments. got=%d, want=2", len(args))
	}

	conn, ok := args[0].(*object.DBConnection)
	if !ok {
		return newError("db_hgetall_redis: first argument must be DB_CONNECTION, got %s", args[0].Type())
	}

	key, ok := args[1].(*object.String)
	if !ok {
		return newError("db_hgetall_redis: second argument must be STRING, got %s", args[1].Type())
	}

	hash, err := HGetAll(conn, key.Value)
	if err != nil {
		return newError("db_hgetall_redis: %s", err.Error())
	}

	// Convert to BanglaCode map
	pairs := make(map[string]object.Object)
	for field, value := range hash {
		pairs[field] = &object.String{Value: value}
	}

	return &object.Map{Pairs: pairs}
}
