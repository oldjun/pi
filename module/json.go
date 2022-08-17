package module

import "pilang/object"

// JsonProperties module properties
var JsonProperties = map[string]object.ModuleProperty{}

// JsonFunctions module functions
var JsonFunctions = map[string]object.ModuleFunction{}

func init() {
	JsonFunctions["decode"] = decode
	JsonFunctions["encode"] = encode
}

func decode(args []object.Object) object.Object {
	return &object.Null{}
}

func encode(args []object.Object) object.Object {
	return &object.Null{}
}
