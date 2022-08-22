package module

import (
	"github.com/oldjun/pi/object"
)

var Map = map[string]*object.Module{}

func init() {
	Map["os"] = &object.Module{Name: "os", Functions: OsFunctions, Properties: OsProperties}
	Map["net"] = &object.Module{Name: "net", Functions: NetFunctions, Properties: NetProperties}
	Map["time"] = &object.Module{Name: "time", Functions: TimeFunctions, Properties: TimeProperties}
	Map["math"] = &object.Module{Name: "math", Functions: MathFunctions, Properties: MathProperties}
	Map["json"] = &object.Module{Name: "json", Functions: JsonFunctions, Properties: JsonProperties}
	Map["sync"] = &object.Module{Name: "sync", Functions: SyncFunctions, Properties: SyncProperties}
	Map["regexp"] = &object.Module{Name: "regexp", Functions: RegexpFunctions, Properties: RegexpProperties}
}
