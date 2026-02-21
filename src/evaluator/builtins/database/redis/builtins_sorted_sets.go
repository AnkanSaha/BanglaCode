package redis

import (
	"BanglaCode/src/object"

	"github.com/redis/go-redis/v9"
)

// Redis Sorted Sets operations

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

// Register sorted set functions
func init() {
	registerBuiltin("db_zadd_redis", dbZAddRedis)
	registerBuiltin("db_zrange_redis", dbZRangeRedis)
	registerBuiltin("db_zrank_redis", dbZRankRedis)
	registerBuiltin("db_zscore_redis", dbZScoreRedis)
	registerBuiltin("db_zrem_redis", dbZRemRedis)
}
