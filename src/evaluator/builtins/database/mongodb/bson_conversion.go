package mongodb

import (
	"BanglaCode/src/object"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
)

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
