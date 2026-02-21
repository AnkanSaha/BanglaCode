package redis

import (
	"BanglaCode/src/object"
)

// Redis Sets operations

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

// Register set functions
func init() {
	registerBuiltin("db_sadd_redis", dbSAddRedis)
	registerBuiltin("db_smembers_redis", dbSMembersRedis)
	registerBuiltin("db_sismember_redis", dbSIsMemberRedis)
	registerBuiltin("db_srem_redis", dbSRemRedis)
	registerBuiltin("db_sinter_redis", dbSInterRedis)
	registerBuiltin("db_sunion_redis", dbSUnionRedis)
	registerBuiltin("db_sdiff_redis", dbSDiffRedis)
}
