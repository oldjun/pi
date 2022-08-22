package module

import (
	"github.com/oldjun/pi/object"
	"sync"
)

type Mutex struct {
	Handler *sync.Mutex
}

func (m *Mutex) Method(method string, args []object.Object) object.Object {
	switch method {
	case "lock":
		return m.lock(args)
	case "unlock":
		return m.unlock(args)
	}
	return object.NewError("sync mutex undefined method: %s", method)
}

func (m *Mutex) lock(args []object.Object) object.Object {
	if len(args) != 0 {
		return object.NewError("wrong number of arguments. mutex.lock() got=%d", len(args))
	}
	m.Handler.Lock()
	return &object.Null{}
}

func (m *Mutex) unlock(args []object.Object) object.Object {
	if len(args) != 0 {
		return object.NewError("wrong number of arguments. mutex.unlock() got=%d", len(args))
	}
	m.Handler.Unlock()
	return &object.Null{}
}
