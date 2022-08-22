package module

import (
	"bytes"
	"github.com/oldjun/pi/object"
	"net"
)

type UnixConnection struct {
	Handler *net.UnixConn
}

func (c *UnixConnection) Method(method string, args []object.Object) object.Object {
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
	return object.NewError("unix_conn undefined method: %s", method)
}

func (c *UnixConnection) read(args []object.Object) object.Object {
	if len(args) != 1 {
		return object.NewError("wrong number of arguments. unix_conn.read() got=%d", len(args))
	}
	switch arg := args[0].(type) {
	case *object.Integer:
		size := arg.Value
		buf := make([]byte, size)
		_, err := c.Handler.Read(buf)
		if err != nil {
			return &object.Bytes{}
		}
		idx := bytes.IndexByte(buf, 0)
		return &object.Bytes{Value: buf[0:idx]}
	}
	return object.NewError("wrong type of arguments. unix_conn.read(): %s", args[0].Type())
}

func (c *UnixConnection) send(args []object.Object) object.Object {
	if len(args) != 1 {
		return object.NewError("wrong number of arguments. unix_conn.send() got=%d", len(args))
	}
	switch arg := args[0].(type) {
	case *object.Bytes:
		data := arg.Value
		size, err := c.Handler.Write(data)
		if err != nil {
			return object.NewError("unix_conn.send() error: %s", err.Error())
		}
		return &object.Integer{Value: int64(size)}
	case *object.String:
		data := []byte(arg.Value)
		size, err := c.Handler.Write(data)
		if err != nil {
			return object.NewError("unix_conn.send() error: %s", err.Error())
		}
		return &object.Integer{Value: int64(size)}
	}
	return object.NewError("wrong type of arguments. unix_conn.send(): %s", args[0].Type())
}

func (c *UnixConnection) close(args []object.Object) object.Object {
	if len(args) != 0 {
		return object.NewError("wrong number of arguments. unix_conn.close() got=%d", len(args))
	}
	err := c.Handler.Close()
	if err != nil {
		return object.NewError("unix_conn.close() error: %s", err.Error())
	}
	return &object.Null{}
}

func (c *UnixConnection) setReadBuffer(args []object.Object) object.Object {
	if len(args) != 1 {
		return object.NewError("wrong number of arguments. unix_conn.set_read_buffer() got=%d", len(args))
	}
	switch arg := args[0].(type) {
	case *object.Integer:
		err := c.Handler.SetReadBuffer(int(arg.Value))
		if err != nil {
			return object.NewError("unix_conn.set_read_buffer() error: %s", err.Error())
		}
		return &object.Null{}
	}
	return object.NewError("wrong type of arguments. unix_conn.set_read_buffer(): %s", args[0].Type())
}

func (c *UnixConnection) setSendBuffer(args []object.Object) object.Object {
	if len(args) != 1 {
		return object.NewError("wrong number of arguments. unix_conn.set_send_buffer() got=%d", len(args))
	}
	switch arg := args[0].(type) {
	case *object.Integer:
		err := c.Handler.SetWriteBuffer(int(arg.Value))
		if err != nil {
			return object.NewError("unix_conn.set_send_buffer() error: %s", err.Error())
		}
		return &object.Null{}
	}
	return object.NewError("wrong type of arguments. unix_conn.set_send_buffer(): %s", args[0].Type())
}
