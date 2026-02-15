package mongodb

import (
	"BanglaCode/src/object"
	"context"
	"fmt"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	clients   = make(map[string]*mongo.Client)
	clientsMu sync.RWMutex
)

// Connect creates a new MongoDB connection
func Connect(config *object.Map) (*object.DBConnection, error) {
	// Extract connection parameters
	host := extractString(config, "host", "localhost")
	port := extractNumber(config, "port", 27017)
	database := extractString(config, "database", "test")
	username := extractString(config, "username", "")
	password := extractString(config, "password", "")

	// Build connection URI
	var uri string
	if username != "" && password != "" {
		uri = fmt.Sprintf("mongodb://%s:%s@%s:%d/%s",
			username, password, host, int(port), database)
	} else {
		uri = fmt.Sprintf("mongodb://%s:%d", host, int(port))
	}

	// Set client options
	clientOpts := options.Client().ApplyURI(uri)

	// Create context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Connect to MongoDB
	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		return nil, fmt.Errorf("failed to connect: %v", err)
	}

	// Ping the database
	if err := client.Ping(ctx, nil); err != nil {
		client.Disconnect(ctx)
		return nil, fmt.Errorf("failed to ping database: %v", err)
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
	metadata["database"] = &object.String{Value: database}

	// Create DBConnection object
	conn := &object.DBConnection{
		ID:       connID,
		DBType:   "mongodb",
		Native:   client,
		Metadata: metadata,
	}

	return conn, nil
}

// Close closes a MongoDB connection
func Close(conn *object.DBConnection) error {
	if conn.DBType != "mongodb" {
		return fmt.Errorf("expected mongodb connection, got %s", conn.DBType)
	}

	client, ok := conn.Native.(*mongo.Client)
	if !ok {
		return fmt.Errorf("invalid native connection type")
	}

	// Remove from global registry
	clientsMu.Lock()
	delete(clients, conn.ID)
	clientsMu.Unlock()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return client.Disconnect(ctx)
}

// GetCollection gets a MongoDB collection
func GetCollection(conn *object.DBConnection, collectionName string) (*mongo.Collection, string, error) {
	client, ok := conn.Native.(*mongo.Client)
	if !ok {
		return nil, "", fmt.Errorf("invalid native connection type")
	}

	// Get database name from metadata
	dbName := "test"
	if dbObj, ok := conn.Metadata["database"]; ok {
		if dbStr, ok := dbObj.(*object.String); ok {
			dbName = dbStr.Value
		}
	}

	collection := client.Database(dbName).Collection(collectionName)
	return collection, dbName, nil
}

// Find finds documents matching a filter
func Find(conn *object.DBConnection, collectionName string, filterMap *object.Map) (*object.DBResult, error) {
	collection, _, err := GetCollection(conn, collectionName)
	if err != nil {
		return nil, err
	}

	// Convert filter map to BSON
	filter := mapToBSON(filterMap)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Find documents
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return &object.DBResult{Error: &object.Error{Message: err.Error()}}, nil
	}
	defer cursor.Close(ctx)

	// Convert results to BanglaCode objects
	results := make([]map[string]object.Object, 0)
	for cursor.Next(ctx) {
		var doc bson.M
		if err := cursor.Decode(&doc); err != nil {
			return &object.DBResult{Error: &object.Error{Message: err.Error()}}, nil
		}

		// Convert BSON to BanglaCode map
		banglaDoc := bsonToMap(doc)
		results = append(results, banglaDoc)
	}

	if err := cursor.Err(); err != nil {
		return &object.DBResult{Error: &object.Error{Message: err.Error()}}, nil
	}

	return &object.DBResult{
		Rows:         results,
		RowsAffected: int64(len(results)),
	}, nil
}

// InsertOne inserts a single document
func InsertOne(conn *object.DBConnection, collectionName string, doc *object.Map) (*object.DBResult, error) {
	collection, _, err := GetCollection(conn, collectionName)
	if err != nil {
		return nil, err
	}

	// Convert document to BSON
	bsonDoc := mapToBSON(doc)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Insert document
	_, err = collection.InsertOne(ctx, bsonDoc)
	if err != nil {
		return &object.DBResult{Error: &object.Error{Message: err.Error()}}, nil
	}

	return &object.DBResult{
		Rows:         []map[string]object.Object{},
		RowsAffected: 1,
		LastInsertID: 0,
		Error:        nil,
	}, nil
}

// UpdateMany updates documents matching a filter
func UpdateMany(conn *object.DBConnection, collectionName string, filterMap, updateMap *object.Map) (*object.DBResult, error) {
	collection, _, err := GetCollection(conn, collectionName)
	if err != nil {
		return nil, err
	}

	// Convert maps to BSON
	filter := mapToBSON(filterMap)
	update := mapToBSON(updateMap)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Update documents
	result, err := collection.UpdateMany(ctx, filter, update)
	if err != nil {
		return &object.DBResult{Error: &object.Error{Message: err.Error()}}, nil
	}

	return &object.DBResult{
		Rows:         []map[string]object.Object{},
		RowsAffected: result.ModifiedCount,
	}, nil
}

// DeleteMany deletes documents matching a filter
func DeleteMany(conn *object.DBConnection, collectionName string, filterMap *object.Map) (*object.DBResult, error) {
	collection, _, err := GetCollection(conn, collectionName)
	if err != nil {
		return nil, err
	}

	// Convert filter to BSON
	filter := mapToBSON(filterMap)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Delete documents
	result, err := collection.DeleteMany(ctx, filter)
	if err != nil {
		return &object.DBResult{Error: &object.Error{Message: err.Error()}}, nil
	}

	return &object.DBResult{
		Rows:         []map[string]object.Object{},
		RowsAffected: result.DeletedCount,
	}, nil
}

// Aggregate executes an aggregation pipeline
func Aggregate(conn *object.DBConnection, collectionName string, pipeline *object.Array) (*object.DBResult, error) {
	collection, _, err := GetCollection(conn, collectionName)
	if err != nil {
		return nil, err
	}

	// Convert pipeline array to BSON array
	bsonPipeline := make([]interface{}, len(pipeline.Elements))
	for i, stage := range pipeline.Elements {
		bsonPipeline[i] = objectToBSON(stage)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Execute aggregation
	cursor, err := collection.Aggregate(ctx, bsonPipeline)
	if err != nil {
		return &object.DBResult{Error: &object.Error{Message: err.Error()}}, nil
	}
	defer cursor.Close(ctx)

	// Convert results
	results := make([]map[string]object.Object, 0)
	for cursor.Next(ctx) {
		var doc bson.M
		if err := cursor.Decode(&doc); err != nil {
			return &object.DBResult{Error: &object.Error{Message: err.Error()}}, nil
		}
		results = append(results, bsonToMap(doc))
	}

	if err := cursor.Err(); err != nil {
		return &object.DBResult{Error: &object.Error{Message: err.Error()}}, nil
	}

	return &object.DBResult{
		Rows:         results,
		RowsAffected: int64(len(results)),
	}, nil
}

// FindOne finds a single document
func FindOne(conn *object.DBConnection, collectionName string, filterMap *object.Map) (*object.DBResult, error) {
	collection, _, err := GetCollection(conn, collectionName)
	if err != nil {
		return nil, err
	}

	filter := mapToBSON(filterMap)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var doc bson.M
	err = collection.FindOne(ctx, filter).Decode(&doc)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return &object.DBResult{
				Rows:         []map[string]object.Object{},
				RowsAffected: 0,
			}, nil
		}
		return &object.DBResult{Error: &object.Error{Message: err.Error()}}, nil
	}

	return &object.DBResult{
		Rows:         []map[string]object.Object{bsonToMap(doc)},
		RowsAffected: 1,
	}, nil
}

// Count counts documents matching a filter
func Count(conn *object.DBConnection, collectionName string, filterMap *object.Map) (int64, error) {
	collection, _, err := GetCollection(conn, collectionName)
	if err != nil {
		return 0, err
	}

	filter := mapToBSON(filterMap)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	count, err := collection.CountDocuments(ctx, filter)
	if err != nil {
		return 0, err
	}

	return count, nil
}

// Distinct gets distinct values for a field
func Distinct(conn *object.DBConnection, collectionName string, field string, filterMap *object.Map) (*object.Array, error) {
	collection, _, err := GetCollection(conn, collectionName)
	if err != nil {
		return nil, err
	}

	filter := mapToBSON(filterMap)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	values, err := collection.Distinct(ctx, field, filter)
	if err != nil {
		return nil, err
	}

	// Convert to BanglaCode array
	elements := make([]object.Object, len(values))
	for i, val := range values {
		elements[i] = bsonToObject(val)
	}

	return &object.Array{Elements: elements}, nil
}

// FindWithOptions finds documents with sort, limit, skip, projection
func FindWithOptions(conn *object.DBConnection, collectionName string, filterMap *object.Map, opts *object.Map) (*object.DBResult, error) {
	collection, _, err := GetCollection(conn, collectionName)
	if err != nil {
		return nil, err
	}

	filter := mapToBSON(filterMap)

	// Build find options
	findOpts := options.Find()

	// Sort
	if sortObj, ok := opts.Pairs["sort"]; ok {
		if sortMap, ok := sortObj.(*object.Map); ok {
			findOpts.SetSort(mapToBSON(sortMap))
		}
	}

	// Limit
	if limitObj, ok := opts.Pairs["limit"]; ok {
		if limitNum, ok := limitObj.(*object.Number); ok {
			findOpts.SetLimit(int64(limitNum.Value))
		}
	}

	// Skip
	if skipObj, ok := opts.Pairs["skip"]; ok {
		if skipNum, ok := skipObj.(*object.Number); ok {
			findOpts.SetSkip(int64(skipNum.Value))
		}
	}

	// Projection (field selection)
	if projObj, ok := opts.Pairs["projection"]; ok {
		if projMap, ok := projObj.(*object.Map); ok {
			findOpts.SetProjection(mapToBSON(projMap))
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, filter, findOpts)
	if err != nil {
		return &object.DBResult{Error: &object.Error{Message: err.Error()}}, nil
	}
	defer cursor.Close(ctx)

	// Convert results
	results := make([]map[string]object.Object, 0)
	for cursor.Next(ctx) {
		var doc bson.M
		if err := cursor.Decode(&doc); err != nil {
			return &object.DBResult{Error: &object.Error{Message: err.Error()}}, nil
		}
		results = append(results, bsonToMap(doc))
	}

	if err := cursor.Err(); err != nil {
		return &object.DBResult{Error: &object.Error{Message: err.Error()}}, nil
	}

	return &object.DBResult{
		Rows:         results,
		RowsAffected: int64(len(results)),
	}, nil
}

// CreateIndex creates an index on a collection
func CreateIndex(conn *object.DBConnection, collectionName string, keys *object.Map, unique bool) error {
	collection, _, err := GetCollection(conn, collectionName)
	if err != nil {
		return err
	}

	indexKeys := mapToBSON(keys)

	indexModel := mongo.IndexModel{
		Keys:    indexKeys,
		Options: options.Index().SetUnique(unique),
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err = collection.Indexes().CreateOne(ctx, indexModel)
	return err
}

// InsertMany inserts multiple documents
func InsertMany(conn *object.DBConnection, collectionName string, docs *object.Array) (*object.DBResult, error) {
	collection, _, err := GetCollection(conn, collectionName)
	if err != nil {
		return nil, err
	}

	// Convert documents to BSON
	bsonDocs := make([]interface{}, len(docs.Elements))
	for i, doc := range docs.Elements {
		bsonDocs[i] = objectToBSON(doc)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	result, err := collection.InsertMany(ctx, bsonDocs)
	if err != nil {
		return &object.DBResult{Error: &object.Error{Message: err.Error()}}, nil
	}

	return &object.DBResult{
		Rows:         []map[string]object.Object{},
		RowsAffected: int64(len(result.InsertedIDs)),
	}, nil
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
	return fmt.Sprintf("mongodb-%d", connIDCounter)
}

// mapToBSON converts a BanglaCode Map to BSON
func mapToBSON(banglaMap *object.Map) bson.M {
	bsonMap := bson.M{}
	for key, value := range banglaMap.Pairs {
		bsonMap[key] = objectToBSON(value)
	}
	return bsonMap
}

// objectToBSON converts a BanglaCode object to BSON value
func objectToBSON(obj object.Object) interface{} {
	switch o := obj.(type) {
	case *object.Number:
		return o.Value
	case *object.String:
		return o.Value
	case *object.Boolean:
		return o.Value
	case *object.Null:
		return nil
	case *object.Array:
		arr := make([]interface{}, len(o.Elements))
		for i, elem := range o.Elements {
			arr[i] = objectToBSON(elem)
		}
		return arr
	case *object.Map:
		return mapToBSON(o)
	default:
		return o.Inspect()
	}
}

// bsonToMap converts BSON to BanglaCode Map
func bsonToMap(bsonDoc bson.M) map[string]object.Object {
	banglaMap := make(map[string]object.Object)
	for key, value := range bsonDoc {
		banglaMap[key] = bsonToObject(value)
	}
	return banglaMap
}

// bsonToObject converts BSON value to BanglaCode object
func bsonToObject(value interface{}) object.Object {
	if value == nil {
		return object.NULL
	}

	switch v := value.(type) {
	case int:
		return &object.Number{Value: float64(v)}
	case int32:
		return &object.Number{Value: float64(v)}
	case int64:
		return &object.Number{Value: float64(v)}
	case float32:
		return &object.Number{Value: float64(v)}
	case float64:
		return &object.Number{Value: v}
	case string:
		return &object.String{Value: v}
	case bool:
		return object.NativeBoolToBooleanObject(v)
	case []interface{}:
		elements := make([]object.Object, len(v))
		for i, elem := range v {
			elements[i] = bsonToObject(elem)
		}
		return &object.Array{Elements: elements}
	case bson.M:
		return &object.Map{Pairs: bsonToMap(v)}
	case map[string]interface{}:
		pairs := make(map[string]object.Object)
		for k, val := range v {
			pairs[k] = bsonToObject(val)
		}
		return &object.Map{Pairs: pairs}
	default:
		return &object.String{Value: fmt.Sprintf("%v", v)}
	}
}
