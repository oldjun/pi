package object

import (
	"fmt"
	"net"
)

type TcpListener struct {
	Handler *net.TCPListener
}

func (l *TcpListener) Type() Type { return "TCP_LISTENER" }
func (l *TcpListener) String() string {
	return fmt.Sprintf("<tcp_listener:%v>", l.Handler)
}

func (l *TcpListener) Method(method string, args []Object) Object {
	switch method {
	case "accept":
		return l.accept(args)
	}
	return NewError("tcp listener undefined method: %s", method)
}

func (l *TcpListener) accept(args []Object) Object {
	if len(args) != 0 {
		return NewError("wrong number of arguments. tcp_listener.accept() got=%d", len(args))
	}
	conn, err := l.Handler.AcceptTCP()
	if err != nil {
		return NewError("tcp_listener.accept() error got=%v", err.Error())
	}
	return &TcpConnection{Handler: conn}
}
