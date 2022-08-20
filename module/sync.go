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
	SyncFunctions["mutex"] = mutex
}

func await(args []object.Object) object.Object {
	if len(args) != 0 {
		return object.NewError("wrong number of arguments. sync.await() got=%d", len(args))
	}
	return &object.SyncAwait{Handler: &sync.WaitGroup{}}
}

func mutex(args []object.Object) object.Object {
	if len(args) != 0 {
		return object.NewError("wrong number of arguments. sync.mutex() got=%d", len(args))
	}
	return &object.SyncMutex{Handler: &sync.Mutex{}}
}
