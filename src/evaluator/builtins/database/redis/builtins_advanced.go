package redis

import (
	"BanglaCode/src/object"

	"github.com/redis/go-redis/v9"
)

// Advanced Redis built-in function implementations

// Sorted Sets operations

// db_zadd_redis - Add members to sorted set with scores
// Usage: db_zadd_redis(conn, "leaderboard", {"player1": 100, "player2": 95})
func dbZAddRedis(args ...object.Object) object.Object {
	if len(args) != 3 {
		return newError("db_zadd_redis: wrong number of arguments. got=%d, want=3 (conn, key, members)", len(args))
	}

	conn, ok := args[0].(*object.DBConnection)
	if !ok {
		return newError("db_zadd_redis: first argument must be DB_CONNECTION, got %s", args[0].Type())
	}

	key, ok := args[1].(*object.String)
	if !ok {
		return newError("db_zadd_redis: second argument must be STRING (key), got %s", args[1].Type())
	}

	membersMap, ok := args[2].(*object.Map)
	if !ok {
		return newError("db_zadd_redis: third argument must be MAP (members with scores), got %s", args[2].Type())
	}

	// Convert map to member-score pairs
	members := make(map[string]float64)
	for member, scoreObj := range membersMap.Pairs {
		score, ok := scoreObj.(*object.Number)
		if !ok {
			return newError("db_zadd_redis: score for member %s must be NUMBER, got %s", member, scoreObj.Type())
		}
		members[member] = score.Value
	}

	err := ZAdd(conn, key.Value, members)
	if err != nil {
		return newError("db_zadd_redis: %s", err.Error())
	}

	return object.TRUE
}

// db_zrange_redis - Get members from sorted set by index range
// Usage: db_zrange_redis(conn, "leaderboard", 0, 9) // Top 10
func dbZRangeRedis(args ...object.Object) object.Object {
	if len(args) != 4 {
		return newError("db_zrange_redis: wrong number of arguments. got=%d, want=4 (conn, key, start, stop)", len(args))
	}

	conn, ok := args[0].(*object.DBConnection)
	if !ok {
		return newError("db_zrange_redis: first argument must be DB_CONNECTION, got %s", args[0].Type())
	}

	key, ok := args[1].(*object.String)
	if !ok {
		return newError("db_zrange_redis: second argument must be STRING (key), got %s", args[1].Type())
	}

	start, ok := args[2].(*object.Number)
	if !ok {
		return newError("db_zrange_redis: third argument must be NUMBER (start), got %s", args[2].Type())
	}

	stop, ok := args[3].(*object.Number)
	if !ok {
		return newError("db_zrange_redis: fourth argument must be NUMBER (stop), got %s", args[3].Type())
	}

	members, err := ZRange(conn, key.Value, int64(start.Value), int64(stop.Value))
	if err != nil {
		return newError("db_zrange_redis: %s", err.Error())
	}

	// Convert to array
	elements := make([]object.Object, len(members))
	for i, member := range members {
		elements[i] = &object.String{Value: member}
	}

	return &object.Array{Elements: elements}
}

// db_zrank_redis - Get rank of member in sorted set
// Usage: db_zrank_redis(conn, "leaderboard", "player1")
func dbZRankRedis(args ...object.Object) object.Object {
	if len(args) != 3 {
		return newError("db_zrank_redis: wrong number of arguments. got=%d, want=3 (conn, key, member)", len(args))
	}

	conn, ok := args[0].(*object.DBConnection)
	if !ok {
		return newError("db_zrank_redis: first argument must be DB_CONNECTION, got %s", args[0].Type())
	}

	key, ok := args[1].(*object.String)
	if !ok {
		return newError("db_zrank_redis: second argument must be STRING (key), got %s", args[1].Type())
	}

	member, ok := args[2].(*object.String)
	if !ok {
		return newError("db_zrank_redis: third argument must be STRING (member), got %s", args[2].Type())
	}

	rank, err := ZRank(conn, key.Value, member.Value)
	if err != nil {
		if err == redis.Nil {
			return object.NULL
		}
		return newError("db_zrank_redis: %s", err.Error())
	}

	return &object.Number{Value: float64(rank)}
}

// db_zscore_redis - Get score of member in sorted set
// Usage: db_zscore_redis(conn, "leaderboard", "player1")
func dbZScoreRedis(args ...object.Object) object.Object {
	if len(args) != 3 {
		return newError("db_zscore_redis: wrong number of arguments. got=%d, want=3 (conn, key, member)", len(args))
	}

	conn, ok := args[0].(*object.DBConnection)
	if !ok {
		return newError("db_zscore_redis: first argument must be DB_CONNECTION, got %s", args[0].Type())
	}

	key, ok := args[1].(*object.String)
	if !ok {
		return newError("db_zscore_redis: second argument must be STRING (key), got %s", args[1].Type())
	}

	member, ok := args[2].(*object.String)
	if !ok {
		return newError("db_zscore_redis: third argument must be STRING (member), got %s", args[2].Type())
	}

	score, err := ZScore(conn, key.Value, member.Value)
	if err != nil {
		if err == redis.Nil {
			return object.NULL
		}
		return newError("db_zscore_redis: %s", err.Error())
	}

	return &object.Number{Value: score}
}

// db_zrem_redis - Remove members from sorted set
// Usage: db_zrem_redis(conn, "leaderboard", ["player1", "player2"])
func dbZRemRedis(args ...object.Object) object.Object {
	if len(args) != 3 {
		return newError("db_zrem_redis: wrong number of arguments. got=%d, want=3 (conn, key, members)", len(args))
	}

	conn, ok := args[0].(*object.DBConnection)
	if !ok {
		return newError("db_zrem_redis: first argument must be DB_CONNECTION, got %s", args[0].Type())
	}

	key, ok := args[1].(*object.String)
	if !ok {
		return newError("db_zrem_redis: second argument must be STRING (key), got %s", args[1].Type())
	}

	membersArray, ok := args[2].(*object.Array)
	if !ok {
		return newError("db_zrem_redis: third argument must be ARRAY (members), got %s", args[2].Type())
	}

	// Convert array to string slice
	members := make([]string, len(membersArray.Elements))
	for i, elem := range membersArray.Elements {
		member, ok := elem.(*object.String)
		if !ok {
			return newError("db_zrem_redis: array element must be STRING, got %s", elem.Type())
		}
		members[i] = member.Value
	}

	removed, err := ZRem(conn, key.Value, members...)
	if err != nil {
		return newError("db_zrem_redis: %s", err.Error())
	}

	return &object.Number{Value: float64(removed)}
}

// Sets operations

// db_sadd_redis - Add members to a set
// Usage: db_sadd_redis(conn, "users", ["user1", "user2", "user3"])
func dbSAddRedis(args ...object.Object) object.Object {
	if len(args) != 3 {
		return newError("db_sadd_redis: wrong number of arguments. got=%d, want=3 (conn, key, members)", len(args))
	}

	conn, ok := args[0].(*object.DBConnection)
	if !ok {
		return newError("db_sadd_redis: first argument must be DB_CONNECTION, got %s", args[0].Type())
	}

	key, ok := args[1].(*object.String)
	if !ok {
		return newError("db_sadd_redis: second argument must be STRING (key), got %s", args[1].Type())
	}

	membersArray, ok := args[2].(*object.Array)
	if !ok {
		return newError("db_sadd_redis: third argument must be ARRAY (members), got %s", args[2].Type())
	}

	// Convert array to string slice
	members := make([]string, len(membersArray.Elements))
	for i, elem := range membersArray.Elements {
		member, ok := elem.(*object.String)
		if !ok {
			return newError("db_sadd_redis: array element must be STRING, got %s", elem.Type())
		}
		members[i] = member.Value
	}

	added, err := SAdd(conn, key.Value, members...)
	if err != nil {
		return newError("db_sadd_redis: %s", err.Error())
	}

	return &object.Number{Value: float64(added)}
}

// db_smembers_redis - Get all members of a set
// Usage: db_smembers_redis(conn, "users")
func dbSMembersRedis(args ...object.Object) object.Object {
	if len(args) != 2 {
		return newError("db_smembers_redis: wrong number of arguments. got=%d, want=2 (conn, key)", len(args))
	}

	conn, ok := args[0].(*object.DBConnection)
	if !ok {
		return newError("db_smembers_redis: first argument must be DB_CONNECTION, got %s", args[0].Type())
	}

	key, ok := args[1].(*object.String)
	if !ok {
		return newError("db_smembers_redis: second argument must be STRING (key), got %s", args[1].Type())
	}

	members, err := SMembers(conn, key.Value)
	if err != nil {
		return newError("db_smembers_redis: %s", err.Error())
	}

	// Convert to array
	elements := make([]object.Object, len(members))
	for i, member := range members {
		elements[i] = &object.String{Value: member}
	}

	return &object.Array{Elements: elements}
}

// db_sismember_redis - Check if member is in set
// Usage: db_sismember_redis(conn, "users", "user1")
func dbSIsMemberRedis(args ...object.Object) object.Object {
	if len(args) != 3 {
		return newError("db_sismember_redis: wrong number of arguments. got=%d, want=3 (conn, key, member)", len(args))
	}

	conn, ok := args[0].(*object.DBConnection)
	if !ok {
		return newError("db_sismember_redis: first argument must be DB_CONNECTION, got %s", args[0].Type())
	}

	key, ok := args[1].(*object.String)
	if !ok {
		return newError("db_sismember_redis: second argument must be STRING (key), got %s", args[1].Type())
	}

	member, ok := args[2].(*object.String)
	if !ok {
		return newError("db_sismember_redis: third argument must be STRING (member), got %s", args[2].Type())
	}

	isMember, err := SIsMember(conn, key.Value, member.Value)
	if err != nil {
		return newError("db_sismember_redis: %s", err.Error())
	}

	return object.NativeBoolToBooleanObject(isMember)
}

// db_srem_redis - Remove members from set
// Usage: db_srem_redis(conn, "users", ["user1", "user2"])
func dbSRemRedis(args ...object.Object) object.Object {
	if len(args) != 3 {
		return newError("db_srem_redis: wrong number of arguments. got=%d, want=3 (conn, key, members)", len(args))
	}

	conn, ok := args[0].(*object.DBConnection)
	if !ok {
		return newError("db_srem_redis: first argument must be DB_CONNECTION, got %s", args[0].Type())
	}

	key, ok := args[1].(*object.String)
	if !ok {
		return newError("db_srem_redis: second argument must be STRING (key), got %s", args[1].Type())
	}

	membersArray, ok := args[2].(*object.Array)
	if !ok {
		return newError("db_srem_redis: third argument must be ARRAY (members), got %s", args[2].Type())
	}

	// Convert array to string slice
	members := make([]string, len(membersArray.Elements))
	for i, elem := range membersArray.Elements {
		member, ok := elem.(*object.String)
		if !ok {
			return newError("db_srem_redis: array element must be STRING, got %s", elem.Type())
		}
		members[i] = member.Value
	}

	removed, err := SRem(conn, key.Value, members...)
	if err != nil {
		return newError("db_srem_redis: %s", err.Error())
	}

	return &object.Number{Value: float64(removed)}
}

// db_sinter_redis - Get intersection of sets
// Usage: db_sinter_redis(conn, ["set1", "set2", "set3"])
func dbSInterRedis(args ...object.Object) object.Object {
	if len(args) != 2 {
		return newError("db_sinter_redis: wrong number of arguments. got=%d, want=2 (conn, keys)", len(args))
	}

	conn, ok := args[0].(*object.DBConnection)
	if !ok {
		return newError("db_sinter_redis: first argument must be DB_CONNECTION, got %s", args[0].Type())
	}

	keysArray, ok := args[1].(*object.Array)
	if !ok {
		return newError("db_sinter_redis: second argument must be ARRAY (keys), got %s", args[1].Type())
	}

	// Convert array to string slice
	keys := make([]string, len(keysArray.Elements))
	for i, elem := range keysArray.Elements {
		key, ok := elem.(*object.String)
		if !ok {
			return newError("db_sinter_redis: array element must be STRING, got %s", elem.Type())
		}
		keys[i] = key.Value
	}

	result, err := SInter(conn, keys...)
	if err != nil {
		return newError("db_sinter_redis: %s", err.Error())
	}

	// Convert to array
	elements := make([]object.Object, len(result))
	for i, member := range result {
		elements[i] = &object.String{Value: member}
	}

	return &object.Array{Elements: elements}
}

// db_sunion_redis - Get union of sets
// Usage: db_sunion_redis(conn, ["set1", "set2"])
func dbSUnionRedis(args ...object.Object) object.Object {
	if len(args) != 2 {
		return newError("db_sunion_redis: wrong number of arguments. got=%d, want=2 (conn, keys)", len(args))
	}

	conn, ok := args[0].(*object.DBConnection)
	if !ok {
		return newError("db_sunion_redis: first argument must be DB_CONNECTION, got %s", args[0].Type())
	}

	keysArray, ok := args[1].(*object.Array)
	if !ok {
		return newError("db_sunion_redis: second argument must be ARRAY (keys), got %s", args[1].Type())
	}

	// Convert array to string slice
	keys := make([]string, len(keysArray.Elements))
	for i, elem := range keysArray.Elements {
		key, ok := elem.(*object.String)
		if !ok {
			return newError("db_sunion_redis: array element must be STRING, got %s", elem.Type())
		}
		keys[i] = key.Value
	}

	result, err := SUnion(conn, keys...)
	if err != nil {
		return newError("db_sunion_redis: %s", err.Error())
	}

	// Convert to array
	elements := make([]object.Object, len(result))
	for i, member := range result {
		elements[i] = &object.String{Value: member}
	}

	return &object.Array{Elements: elements}
}

// db_sdiff_redis - Get difference of sets
// Usage: db_sdiff_redis(conn, ["set1", "set2"])
func dbSDiffRedis(args ...object.Object) object.Object {
	if len(args) != 2 {
		return newError("db_sdiff_redis: wrong number of arguments. got=%d, want=2 (conn, keys)", len(args))
	}

	conn, ok := args[0].(*object.DBConnection)
	if !ok {
		return newError("db_sdiff_redis: first argument must be DB_CONNECTION, got %s", args[0].Type())
	}

	keysArray, ok := args[1].(*object.Array)
	if !ok {
		return newError("db_sdiff_redis: second argument must be ARRAY (keys), got %s", args[1].Type())
	}

	// Convert array to string slice
	keys := make([]string, len(keysArray.Elements))
	for i, elem := range keysArray.Elements {
		key, ok := elem.(*object.String)
		if !ok {
			return newError("db_sdiff_redis: array element must be STRING, got %s", elem.Type())
		}
		keys[i] = key.Value
	}

	result, err := SDiff(conn, keys...)
	if err != nil {
		return newError("db_sdiff_redis: %s", err.Error())
	}

	// Convert to array
	elements := make([]object.Object, len(result))
	for i, member := range result {
		elements[i] = &object.String{Value: member}
	}

	return &object.Array{Elements: elements}
}

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
