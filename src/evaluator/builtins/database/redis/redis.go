package redis

import (
	"BanglaCode/src/object"
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/redis/go-redis/v9"
)

var (
	clients   = make(map[string]*redis.Client)
	clientsMu sync.RWMutex
	ctx       = context.Background()
)

// Connect creates a new Redis connection
func Connect(config *object.Map) (*object.DBConnection, error) {
	// Extract connection parameters
	host := extractString(config, "host", "localhost")
	port := extractNumber(config, "port", 6379)
	password := extractString(config, "password", "")
	db := int(extractNumber(config, "db", 0))

	// Create Redis client options
	opts := &redis.Options{
		Addr:     fmt.Sprintf("%s:%d", host, int(port)),
		Password: password,
		DB:       db,
	}

	// Create Redis client
	client := redis.NewClient(opts)

	// Test connection
	if err := client.Ping(ctx).Err(); err != nil {
		return nil, fmt.Errorf("failed to ping Redis: %v", err)
	}

	// Generate unique connection ID
	connID := generateConnID()

	// Store client globally
	clientsMu.Lock()
	clients[connID] = client
	clientsMu.Unlock()

	// Create metadata
	metadata := make(map[string]object.Object)
	metadata["host"] = &object.String{Value: host}
	metadata["port"] = &object.Number{Value: port}
	metadata["db"] = &object.Number{Value: float64(db)}

	// Create DBConnection object
	conn := &object.DBConnection{
		ID:       connID,
		DBType:   "redis",
		Native:   client,
		Metadata: metadata,
	}

	return conn, nil
}

// Close closes a Redis connection
func Close(conn *object.DBConnection) error {
	if conn.DBType != "redis" {
		return fmt.Errorf("expected redis connection, got %s", conn.DBType)
	}

	client, ok := conn.Native.(*redis.Client)
	if !ok {
		return fmt.Errorf("invalid native connection type")
	}

	// Remove from global registry
	clientsMu.Lock()
	delete(clients, conn.ID)
	clientsMu.Unlock()

	return client.Close()
}

// String operations

// Set sets a key-value pair
func Set(conn *object.DBConnection, key string, value string, expiration time.Duration) error {
	client, ok := conn.Native.(*redis.Client)
	if !ok {
		return fmt.Errorf("invalid native connection type")
	}

	return client.Set(ctx, key, value, expiration).Err()
}

// Get gets a value by key
func Get(conn *object.DBConnection, key string) (string, error) {
	client, ok := conn.Native.(*redis.Client)
	if !ok {
		return "", fmt.Errorf("invalid native connection type")
	}

	val, err := client.Get(ctx, key).Result()
	if err == redis.Nil {
		return "", fmt.Errorf("key does not exist")
	}
	return val, err
}

// Del deletes a key
func Del(conn *object.DBConnection, key string) error {
	client, ok := conn.Native.(*redis.Client)
	if !ok {
		return fmt.Errorf("invalid native connection type")
	}

	return client.Del(ctx, key).Err()
}

// Expire sets expiration on a key
func Expire(conn *object.DBConnection, key string, seconds int) error {
	client, ok := conn.Native.(*redis.Client)
	if !ok {
		return fmt.Errorf("invalid native connection type")
	}

	return client.Expire(ctx, key, time.Duration(seconds)*time.Second).Err()
}

// List operations

// LPush pushes a value to the left of a list
func LPush(conn *object.DBConnection, key string, value string) (int64, error) {
	client, ok := conn.Native.(*redis.Client)
	if !ok {
		return 0, fmt.Errorf("invalid native connection type")
	}

	return client.LPush(ctx, key, value).Result()
}

// RPush pushes a value to the right of a list
func RPush(conn *object.DBConnection, key string, value string) (int64, error) {
	client, ok := conn.Native.(*redis.Client)
	if !ok {
		return 0, fmt.Errorf("invalid native connection type")
	}

	return client.RPush(ctx, key, value).Result()
}

// LPop pops a value from the left of a list
func LPop(conn *object.DBConnection, key string) (string, error) {
	client, ok := conn.Native.(*redis.Client)
	if !ok {
		return "", fmt.Errorf("invalid native connection type")
	}

	val, err := client.LPop(ctx, key).Result()
	if err == redis.Nil {
		return "", fmt.Errorf("list is empty")
	}
	return val, err
}

// RPop pops a value from the right of a list
func RPop(conn *object.DBConnection, key string) (string, error) {
	client, ok := conn.Native.(*redis.Client)
	if !ok {
		return "", fmt.Errorf("invalid native connection type")
	}

	val, err := client.RPop(ctx, key).Result()
	if err == redis.Nil {
		return "", fmt.Errorf("list is empty")
	}
	return val, err
}

// Hash operations

// HSet sets a hash field
func HSet(conn *object.DBConnection, key string, field string, value string) error {
	client, ok := conn.Native.(*redis.Client)
	if !ok {
		return fmt.Errorf("invalid native connection type")
	}

	return client.HSet(ctx, key, field, value).Err()
}

// HGet gets a hash field
func HGet(conn *object.DBConnection, key string, field string) (string, error) {
	client, ok := conn.Native.(*redis.Client)
	if !ok {
		return "", fmt.Errorf("invalid native connection type")
	}

	val, err := client.HGet(ctx, key, field).Result()
	if err == redis.Nil {
		return "", fmt.Errorf("field does not exist")
	}
	return val, err
}

// HGetAll gets all fields in a hash
func HGetAll(conn *object.DBConnection, key string) (map[string]string, error) {
	client, ok := conn.Native.(*redis.Client)
	if !ok {
		return nil, fmt.Errorf("invalid native connection type")
	}

	return client.HGetAll(ctx, key).Result()
}

// Sorted Sets operations

// ZAdd adds members to a sorted set with scores
func ZAdd(conn *object.DBConnection, key string, members map[string]float64) error {
	client, ok := conn.Native.(*redis.Client)
	if !ok {
		return fmt.Errorf("invalid native connection type")
	}

	// Convert to Redis Z members
	zMembers := make([]redis.Z, 0, len(members))
	for member, score := range members {
		zMembers = append(zMembers, redis.Z{
			Score:  score,
			Member: member,
		})
	}

	return client.ZAdd(ctx, key, zMembers...).Err()
}

// ZRange gets members from sorted set by index range
func ZRange(conn *object.DBConnection, key string, start, stop int64) ([]string, error) {
	client, ok := conn.Native.(*redis.Client)
	if !ok {
		return nil, fmt.Errorf("invalid native connection type")
	}

	return client.ZRange(ctx, key, start, stop).Result()
}

// ZRangeWithScores gets members with scores from sorted set
func ZRangeWithScores(conn *object.DBConnection, key string, start, stop int64) ([]redis.Z, error) {
	client, ok := conn.Native.(*redis.Client)
	if !ok {
		return nil, fmt.Errorf("invalid native connection type")
	}

	return client.ZRangeWithScores(ctx, key, start, stop).Result()
}

// ZRank gets rank of member in sorted set
func ZRank(conn *object.DBConnection, key, member string) (int64, error) {
	client, ok := conn.Native.(*redis.Client)
	if !ok {
		return 0, fmt.Errorf("invalid native connection type")
	}

	return client.ZRank(ctx, key, member).Result()
}

// ZScore gets score of member in sorted set
func ZScore(conn *object.DBConnection, key, member string) (float64, error) {
	client, ok := conn.Native.(*redis.Client)
	if !ok {
		return 0, fmt.Errorf("invalid native connection type")
	}

	return client.ZScore(ctx, key, member).Result()
}

// ZRem removes members from sorted set
func ZRem(conn *object.DBConnection, key string, members ...string) (int64, error) {
	client, ok := conn.Native.(*redis.Client)
	if !ok {
		return 0, fmt.Errorf("invalid native connection type")
	}

	// Convert to interface slice
	membersInterface := make([]interface{}, len(members))
	for i, m := range members {
		membersInterface[i] = m
	}

	return client.ZRem(ctx, key, membersInterface...).Result()
}

// Sets operations

// SAdd adds members to a set
func SAdd(conn *object.DBConnection, key string, members ...string) (int64, error) {
	client, ok := conn.Native.(*redis.Client)
	if !ok {
		return 0, fmt.Errorf("invalid native connection type")
	}

	// Convert to interface slice
	membersInterface := make([]interface{}, len(members))
	for i, m := range members {
		membersInterface[i] = m
	}

	return client.SAdd(ctx, key, membersInterface...).Result()
}

// SMembers gets all members of a set
func SMembers(conn *object.DBConnection, key string) ([]string, error) {
	client, ok := conn.Native.(*redis.Client)
	if !ok {
		return nil, fmt.Errorf("invalid native connection type")
	}

	return client.SMembers(ctx, key).Result()
}

// SIsMember checks if member is in set
func SIsMember(conn *object.DBConnection, key, member string) (bool, error) {
	client, ok := conn.Native.(*redis.Client)
	if !ok {
		return false, fmt.Errorf("invalid native connection type")
	}

	return client.SIsMember(ctx, key, member).Result()
}

// SRem removes members from set
func SRem(conn *object.DBConnection, key string, members ...string) (int64, error) {
	client, ok := conn.Native.(*redis.Client)
	if !ok {
		return 0, fmt.Errorf("invalid native connection type")
	}

	// Convert to interface slice
	membersInterface := make([]interface{}, len(members))
	for i, m := range members {
		membersInterface[i] = m
	}

	return client.SRem(ctx, key, membersInterface...).Result()
}

// SInter gets intersection of sets
func SInter(conn *object.DBConnection, keys ...string) ([]string, error) {
	client, ok := conn.Native.(*redis.Client)
	if !ok {
		return nil, fmt.Errorf("invalid native connection type")
	}

	return client.SInter(ctx, keys...).Result()
}

// SUnion gets union of sets
func SUnion(conn *object.DBConnection, keys ...string) ([]string, error) {
	client, ok := conn.Native.(*redis.Client)
	if !ok {
		return nil, fmt.Errorf("invalid native connection type")
	}

	return client.SUnion(ctx, keys...).Result()
}

// SDiff gets difference of sets
func SDiff(conn *object.DBConnection, keys ...string) ([]string, error) {
	client, ok := conn.Native.(*redis.Client)
	if !ok {
		return nil, fmt.Errorf("invalid native connection type")
	}

	return client.SDiff(ctx, keys...).Result()
}

// Counter operations

// Incr increments a key by 1
func Incr(conn *object.DBConnection, key string) (int64, error) {
	client, ok := conn.Native.(*redis.Client)
	if !ok {
		return 0, fmt.Errorf("invalid native connection type")
	}

	return client.Incr(ctx, key).Result()
}

// Decr decrements a key by 1
func Decr(conn *object.DBConnection, key string) (int64, error) {
	client, ok := conn.Native.(*redis.Client)
	if !ok {
		return 0, fmt.Errorf("invalid native connection type")
	}

	return client.Decr(ctx, key).Result()
}

// IncrBy increments a key by a specific amount
func IncrBy(conn *object.DBConnection, key string, value int64) (int64, error) {
	client, ok := conn.Native.(*redis.Client)
	if !ok {
		return 0, fmt.Errorf("invalid native connection type")
	}

	return client.IncrBy(ctx, key, value).Result()
}

// DecrBy decrements a key by a specific amount
func DecrBy(conn *object.DBConnection, key string, value int64) (int64, error) {
	client, ok := conn.Native.(*redis.Client)
	if !ok {
		return 0, fmt.Errorf("invalid native connection type")
	}

	return client.DecrBy(ctx, key, value).Result()
}

// IncrByFloat increments a key by a float value
func IncrByFloat(conn *object.DBConnection, key string, value float64) (float64, error) {
	client, ok := conn.Native.(*redis.Client)
	if !ok {
		return 0, fmt.Errorf("invalid native connection type")
	}

	return client.IncrByFloat(ctx, key, value).Result()
}

// Pub/Sub operations

// Publish publishes a message to a channel
func Publish(conn *object.DBConnection, channel, message string) (int64, error) {
	client, ok := conn.Native.(*redis.Client)
	if !ok {
		return 0, fmt.Errorf("invalid native connection type")
	}

	return client.Publish(ctx, channel, message).Result()
}

// Additional utility operations

// TTL gets time to live of a key
func TTL(conn *object.DBConnection, key string) (int64, error) {
	client, ok := conn.Native.(*redis.Client)
	if !ok {
		return 0, fmt.Errorf("invalid native connection type")
	}

	duration, err := client.TTL(ctx, key).Result()
	if err != nil {
		return 0, err
	}

	return int64(duration.Seconds()), nil
}

// Persist removes expiration from a key
func Persist(conn *object.DBConnection, key string) (bool, error) {
	client, ok := conn.Native.(*redis.Client)
	if !ok {
		return false, fmt.Errorf("invalid native connection type")
	}

	return client.Persist(ctx, key).Result()
}

// Exists checks if keys exist
func Exists(conn *object.DBConnection, keys ...string) (int64, error) {
	client, ok := conn.Native.(*redis.Client)
	if !ok {
		return 0, fmt.Errorf("invalid native connection type")
	}

	return client.Exists(ctx, keys...).Result()
}

// Keys gets all keys matching a pattern
func Keys(conn *object.DBConnection, pattern string) ([]string, error) {
	client, ok := conn.Native.(*redis.Client)
	if !ok {
		return nil, fmt.Errorf("invalid native connection type")
	}

	return client.Keys(ctx, pattern).Result()
}

// Helper functions

func extractString(config *object.Map, key string, defaultValue string) string {
	if val, ok := config.Pairs[key]; ok {
		if str, ok := val.(*object.String); ok {
			return str.Value
		}
	}
	return defaultValue
}

func extractNumber(config *object.Map, key string, defaultValue float64) float64 {
	if val, ok := config.Pairs[key]; ok {
		if num, ok := val.(*object.Number); ok {
			return num.Value
		}
	}
	return defaultValue
}

var connIDCounter int64

func generateConnID() string {
	connIDCounter++
	return fmt.Sprintf("redis-%d", connIDCounter)
}
