package module

import (
	"github.com/oldjun/pi/object"
	"sync"
)

type Await struct {
	Handler *sync.WaitGroup
}

func (a *Await) Method(method string, args []object.Object) object.Object {
	switch method {
	case "add":
		return a.add(args)
	case "done":
		return a.done(args)
	case "wait":
		return a.wait(args)
	}
	return object.NewError("sync await undefined method: %s", method)
}

func (a *Await) add(args []object.Object) object.Object {
	if len(args) != 1 {
		return object.NewError("wrong number of arguments. await.add() got=%d", len(args))
	}
	switch arg := args[0].(type) {
	case *object.Integer:
		a.Handler.Add(int(arg.Value))
	default:
		return object.NewError("wrong type of arguments. await.add() got=%s", arg.Type())
	}
	return &object.Null{}
}

func (a *Await) done(args []object.Object) object.Object {
	if len(args) != 0 {
		return object.NewError("wrong number of arguments. await.done() got=%d", len(args))
	}
	a.Handler.Done()
	return &object.Null{}
}

func (a *Await) wait(args []object.Object) object.Object {
	if len(args) != 0 {
		return object.NewError("wrong number of arguments. await.wait() got=%d", len(args))
	}
	a.Handler.Wait()
	return &object.Null{}
}
