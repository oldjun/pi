package module

import "pilang/object"

var Map = map[string]*object.Module{}

func init() {
	Map["os"] = &object.Module{Name: "os", Functions: OsFunctions, Properties: OsProperties}
	Map["math"] = &object.Module{Name: "math", Functions: MathFunctions, Properties: MathProperties}
}
