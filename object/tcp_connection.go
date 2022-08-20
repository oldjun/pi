package object

import (
	"fmt"
	"net"
)

type TcpConnection struct {
	Handler *net.TCPConn
}

func (c *TcpConnection) Type() Type { return "TCP_CONNECTION" }
func (c *TcpConnection) String() string {
	return fmt.Sprintf("<tcp_connection:%v>", c.Handler)
}

func (c *TcpConnection) Method(method string, args []Object) Object {
	switch method {
	case "read":
		return c.read(args)
	case "send":
		return c.send(args)
	}
	return NewError("tcp_conn undefined method: %s", method)
}

func (c *TcpConnection) read(args []Object) Object {
	if len(args) != 1 {
		return NewError("wrong number of arguments. tcp_conn.read() got=%d", len(args))
	}
	switch arg := args[0].(type) {
	case *Integer:
		size := arg.Value
		buffer := make([]byte, size)
		_, err := c.Handler.Read(buffer)
		if err != nil {
			return NewError("tcp_conn.read() error: %s", err.Error())
		}
		return &String{Value: string(buffer)}
	}
	return NewError("wrong type of arguments. tcp_conn.read(): %s", args[0].Type())
}

func (c *TcpConnection) send(args []Object) Object {
	if len(args) != 1 {
		return NewError("wrong number of arguments. tcp_conn.send() got=%d", len(args))
	}
	data := []byte(args[0].(*String).Value)
	size, err := c.Handler.Write(data)
	if err != nil {
		return NewError("tcp_conn.send() error: %s", err.Error())
	}
	return &Integer{Value: int64(size)}
}
