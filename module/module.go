package module

import "github.com/oldjun/pi/object"

var Map = map[string]*object.Module{}

func init() {
	Map["os"] = &object.Module{Name: "os", Functions: OsFunctions, Properties: OsProperties}
	Map["time"] = &object.Module{Name: "time", Functions: TimeFunctions, Properties: TimeProperties}
	Map["math"] = &object.Module{Name: "math", Functions: MathFunctions, Properties: MathProperties}
	Map["json"] = &object.Module{Name: "json", Functions: JsonFunctions, Properties: JsonProperties}
}
