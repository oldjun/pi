package object

import (
	"bytes"
	"fmt"
	"net"
	"time"
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
	case "close":
		return c.close(args)
	case "set_keep_alive":
		return c.setKeepAlive(args)
	case "set_keep_alive_time":
		return c.setKeepAliveTime(args)
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
		buf := make([]byte, size)
		_, err := c.Handler.Read(buf)
		if err != nil {
			return &Bytes{}
		}
		idx := bytes.IndexByte(buf, 0)
		return &Bytes{Value: buf[0:idx]}
	}
	return NewError("wrong type of arguments. tcp_conn.read(): %s", args[0].Type())
}

func (c *TcpConnection) send(args []Object) Object {
	if len(args) != 1 {
		return NewError("wrong number of arguments. tcp_conn.send() got=%d", len(args))
	}
	switch arg := args[0].(type) {
	case *Bytes:
		data := arg.Value
		size, err := c.Handler.Write(data)
		if err != nil {
			return NewError("tcp_conn.send() error: %s", err.Error())
		}
		return &Integer{Value: int64(size)}
	case *String:
		data := []byte(arg.Value)
		size, err := c.Handler.Write(data)
		if err != nil {
			return NewError("tcp_conn.send() error: %s", err.Error())
		}
		return &Integer{Value: int64(size)}
	}
	return NewError("wrong type of arguments. tcp_conn.send(): %s", args[0].Type())
}

func (c *TcpConnection) close(args []Object) Object {
	if len(args) != 0 {
		return NewError("wrong number of arguments. tcp_conn.close() got=%d", len(args))
	}
	err := c.Handler.Close()
	if err != nil {
		return NewError("tcp_conn.close() error: %s", err.Error())
	}
	return &Null{}
}

func (c *TcpConnection) setKeepAlive(args []Object) Object {
	if len(args) != 1 {
		return NewError("wrong number of arguments. tcp_conn.set_keep_alive() got=%d", len(args))
	}
	switch arg := args[0].(type) {
	case *Boolean:
		err := c.Handler.SetKeepAlive(arg.Value)
		if err != nil {
			return NewError("tcp_conn.set_keep_alive() error: %s", err.Error())
		}
		return &Null{}
	}
	return NewError("wrong type of arguments. tcp_conn.set_keep_alive(): %s", args[0].Type())
}

func (c *TcpConnection) setKeepAliveTime(args []Object) Object {
	if len(args) != 1 {
		return NewError("wrong number of arguments. tcp_conn.set_keep_alive_time() got=%d", len(args))
	}
	switch arg := args[0].(type) {
	case *Integer:
		err := c.Handler.SetKeepAlivePeriod(time.Second * time.Duration(arg.Value))
		if err != nil {
			return NewError("tcp_conn.set_keep_alive_time() error: %s", err.Error())
		}
		return &Null{}
	}
	return NewError("wrong type of arguments. tcp_conn.set_keep_alive_time(): %s", args[0].Type())
}
