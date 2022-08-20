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
	SyncFunctions["await"] = await
}

func await(args []object.Object) object.Object {
	if len(args) != 0 {
		return object.NewError("wrong number of arguments. sync.await() got=%d", len(args))
	}
	return &object.SyncAwait{Value: &sync.WaitGroup{}}
}
