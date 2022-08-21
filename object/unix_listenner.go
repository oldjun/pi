package object

import (
	"fmt"
	"net"
)

type UnixListener struct {
	Handler *net.UnixListener
}

func (l *UnixListener) Type() Type { return "UNIX_LISTENER" }
func (l *UnixListener) String() string {
	return fmt.Sprintf("<unix_listener:%v>", l.Handler)
}

func (l *UnixListener) Method(method string, args []Object) Object {
	switch method {
	case "accept":
		return l.accept(args)
	case "close":
		return l.close(args)
	}
	return NewError("unix listener undefined method: %s", method)
}

func (l *UnixListener) accept(args []Object) Object {
	if len(args) != 0 {
		return NewError("wrong number of arguments. unix_listener.accept() got=%d", len(args))
	}
	conn, err := l.Handler.AcceptUnix()
	if err != nil {
		return NewError("unix_listener.accept() error: %s", err.Error())
	}
	return &UnixConnection{Handler: conn}
}

func (l *UnixListener) close(args []Object) Object {
	if len(args) != 0 {
		return NewError("wrong number of arguments. unix_listener.close() got=%d", len(args))
	}
	err := l.Handler.Close()
	if err != nil {
		return NewError("unix_listener.close() error: %s", err.Error())
	}
	return &Null{}
}
