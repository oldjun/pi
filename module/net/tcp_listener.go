package module

import (
	"github.com/oldjun/pi/object"
	"net"
)

type TcpListener struct {
	Handler *net.TCPListener
}

func (l *TcpListener) Method(method string, args []object.Object) object.Object {
	switch method {
	case "accept":
		return l.accept(args)
	case "close":
		return l.close(args)
	}
	return object.NewError("tcp listener undefined method: %s", method)
}

func (l *TcpListener) accept(args []object.Object) object.Object {
	if len(args) != 0 {
		return object.NewError("wrong number of arguments. tcp_listener.accept() got=%d", len(args))
	}
	conn, err := l.Handler.AcceptTCP()
	if err != nil {
		return object.NewError("tcp_listener.accept() error got=%v", err.Error())
	}
	return &object.Module{Name: "connection", Handler: &TcpConnection{Handler: conn}}
}

func (l *TcpListener) close(args []object.Object) object.Object {
	if len(args) != 0 {
		return object.NewError("wrong number of arguments. tcp_listener.close() got=%d", len(args))
	}
	err := l.Handler.Close()
	if err != nil {
		return object.NewError("tcp_listener.close() error: %s", err.Error())
	}
	return &object.Null{}
}
