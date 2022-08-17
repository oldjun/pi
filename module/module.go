package module

import "pilang/object"

var Map = map[string]*object.Module{}

func init() {
	Map["math"] = &object.Module{Name: "math", Functions: MathFunctions, Properties: MathProperties}
}
