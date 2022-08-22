package module

import (
	"github.com/oldjun/pi/object"
	"net"
)

type UnixListener struct {
	Handler *net.UnixListener
}

func (l *UnixListener) Method(method string, args []object.Object) object.Object {
	switch method {
	case "accept":
		return l.accept(args)
	case "close":
		return l.close(args)
	}
	return object.NewError("unix listener undefined method: %s", method)
}

func (l *UnixListener) accept(args []object.Object) object.Object {
	if len(args) != 0 {
		return object.NewError("wrong number of arguments. unix_listener.accept() got=%d", len(args))
	}
	conn, err := l.Handler.AcceptUnix()
	if err != nil {
		return object.NewError("unix_listener.accept() error: %s", err.Error())
	}
	return &object.Module{Name: "connection", Handler: &UnixConnection{Handler: conn}}
}

func (l *UnixListener) close(args []object.Object) object.Object {
	if len(args) != 0 {
		return object.NewError("wrong number of arguments. unix_listener.close() got=%d", len(args))
	}
	err := l.Handler.Close()
	if err != nil {
		return object.NewError("unix_listener.close() error: %s", err.Error())
	}
	return &object.Null{}
}
