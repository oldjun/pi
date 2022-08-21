package object

import (
	"bytes"
	"fmt"
	"net"
)

type UnixConnection struct {
	Handler *net.UnixConn
}

func (c *UnixConnection) Type() Type { return "UNIX_CONNECTION" }
func (c *UnixConnection) String() string {
	return fmt.Sprintf("<unix_connection:%v>", c.Handler)
}

func (c *UnixConnection) Method(method string, args []Object) Object {
	switch method {
	case "read":
		return c.read(args)
	case "send":
		return c.send(args)
	case "close":
		return c.close(args)
	case "set_send_buffer":
		return c.setSendBuffer(args)
	case "set_read_buffer":
		return c.setReadBuffer(args)
	}
	return NewError("unix_conn undefined method: %s", method)
}

func (c *UnixConnection) read(args []Object) Object {
	if len(args) != 1 {
		return NewError("wrong number of arguments. unix_conn.read() got=%d", len(args))
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
	return NewError("wrong type of arguments. unix_conn.read(): %s", args[0].Type())
}

func (c *UnixConnection) send(args []Object) Object {
	if len(args) != 1 {
		return NewError("wrong number of arguments. unix_conn.send() got=%d", len(args))
	}
	switch arg := args[0].(type) {
	case *Bytes:
		data := arg.Value
		size, err := c.Handler.Write(data)
		if err != nil {
			return NewError("unix_conn.send() error: %s", err.Error())
		}
		return &Integer{Value: int64(size)}
	case *String:
		data := []byte(arg.Value)
		size, err := c.Handler.Write(data)
		if err != nil {
			return NewError("unix_conn.send() error: %s", err.Error())
		}
		return &Integer{Value: int64(size)}
	}
	return NewError("wrong type of arguments. unix_conn.send(): %s", args[0].Type())
}

func (c *UnixConnection) close(args []Object) Object {
	if len(args) != 0 {
		return NewError("wrong number of arguments. unix_conn.close() got=%d", len(args))
	}
	err := c.Handler.Close()
	if err != nil {
		return NewError("unix_conn.close() error: %s", err.Error())
	}
	return &Null{}
}

func (c *UnixConnection) setReadBuffer(args []Object) Object {
	if len(args) != 1 {
		return NewError("wrong number of arguments. unix_conn.set_read_buffer() got=%d", len(args))
	}
	switch arg := args[0].(type) {
	case *Integer:
		err := c.Handler.SetReadBuffer(int(arg.Value))
		if err != nil {
			return NewError("unix_conn.set_read_buffer() error: %s", err.Error())
		}
		return &Null{}
	}
	return NewError("wrong type of arguments. unix_conn.set_read_buffer(): %s", args[0].Type())
}

func (c *UnixConnection) setSendBuffer(args []Object) Object {
	if len(args) != 1 {
		return NewError("wrong number of arguments. unix_conn.set_send_buffer() got=%d", len(args))
	}
	switch arg := args[0].(type) {
	case *Integer:
		err := c.Handler.SetWriteBuffer(int(arg.Value))
		if err != nil {
			return NewError("unix_conn.set_send_buffer() error: %s", err.Error())
		}
		return &Null{}
	}
	return NewError("wrong type of arguments. unix_conn.set_send_buffer(): %s", args[0].Type())
}
