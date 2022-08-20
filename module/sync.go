package module

import (
	"github.com/oldjun/pi/object"
	"sync"
)

// SyncProperties module properties
var SyncProperties = map[string]object.ModuleProperty{}

// SyncFunctions module functions
var SyncFunctions = map[string]object.ModuleFunction{}

func init() {
	SyncFunctions["wait_group"] = waitGroup
}

func waitGroup(args []object.Object) object.Object {
	if len(args) != 0 {
		return object.NewError("wrong number of arguments. sync.wait_group() got=%d", len(args))
	}
	return &object.SyncWaitGroup{Value: &sync.WaitGroup{}}
}
