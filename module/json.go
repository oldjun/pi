package module

import (
	"encoding/json"
	"github.com/oldjun/pi/object"
)

// JsonProperties module properties
var JsonProperties = map[string]object.ModuleProperty{}

// JsonFunctions module functions
var JsonFunctions = map[string]object.ModuleFunction{}

func init() {
	JsonFunctions["decode"] = decode
	JsonFunctions["encode"] = encode
}

func decode(args []object.Object) object.Object {
	var i interface{}
	input := args[0].(*object.String).Value
	err := json.Unmarshal([]byte(input), &i)
	if err != nil {
		return object.NewError("Error while parsing json: %s", err)
	}
	return interfaceToObject(i)
}

func encode(args []object.Object) object.Object {
	if len(args) != 1 {
		return object.NewError("wrong number of arguments. json.encode() got=%d", len(args))
	}
	switch arg := args[0].(type) {
	case *object.List:
		return &object.String{Value: arg.String()}
	case *object.Hash:
		return &object.String{Value: arg.String()}
	}
	return object.NewError("wrong type of arguments. json.encode() got=%s", args[0].Type())
}

func interfaceToObject(i interface{}) object.Object {
	switch v := i.(type) {
	case map[string]interface{}:
		hash := &object.Hash{}
		hash.Pairs = make(map[object.HashKey]object.HashPair)
		for key, val := range v {
			pair := object.HashPair{
				Key:   &object.String{Value: key},
				Value: interfaceToObject(val),
			}
			hash.Pairs[pair.Key.(object.Hashable).HashKey()] = pair
		}
		return hash
	case []interface{}:
		list := &object.List{}
		for _, elem := range v {
			list.Elements = append(list.Elements, interfaceToObject(elem))
		}
		return list
	case string:
		return &object.String{Value: v}
	case int64:
		return &object.Integer{Value: v}
	case float64:
		return &object.Float{Value: v}
	case bool:
		if v {
			return &object.Boolean{Value: true}
		} else {
			return &object.Boolean{Value: false}
		}
	}
	return &object.Null{}
}
