package module

import (
	"bytes"
	"github.com/oldjun/pi/object"
	"net"
	"time"
)

type TcpConnection struct {
	Handler *net.TCPConn
}

func (c *TcpConnection) Method(method string, args []object.Object) object.Object {
	switch method {
	case "read":
		return c.read(args)
	case "send":
		return c.send(args)
	case "close":
		return c.close(args)
	case "set_read_buffer":
		return c.setReadBuffer(args)
	case "set_send_buffer":
		return c.setSendBuffer(args)
	case "set_keep_alive":
		return c.setKeepAlive(args)
	case "set_keep_alive_time":
		return c.setKeepAliveTime(args)
	case "set_linger":
		return c.setLinger(args)
	case "set_no_delay":
		return c.setNoDelay(args)
	}
	return object.NewError("tcp_conn undefined method: %s", method)
}

func (c *TcpConnection) read(args []object.Object) object.Object {
	if len(args) != 1 {
		return object.NewError("wrong number of arguments. tcp_conn.read() got=%d", len(args))
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
	return object.NewError("wrong type of arguments. tcp_conn.read(): %s", args[0].Type())
}

func (c *TcpConnection) send(args []object.Object) object.Object {
	if len(args) != 1 {
		return object.NewError("wrong number of arguments. tcp_conn.send() got=%d", len(args))
	}
	switch arg := args[0].(type) {
	case *object.Bytes:
		data := arg.Value
		size, err := c.Handler.Write(data)
		if err != nil {
			return object.NewError("tcp_conn.send() error: %s", err.Error())
		}
		return &object.Integer{Value: int64(size)}
	case *object.String:
		data := []byte(arg.Value)
		size, err := c.Handler.Write(data)
		if err != nil {
			return object.NewError("tcp_conn.send() error: %s", err.Error())
		}
		return &object.Integer{Value: int64(size)}
	}
	return object.NewError("wrong type of arguments. tcp_conn.send(): %s", args[0].Type())
}

func (c *TcpConnection) close(args []object.Object) object.Object {
	if len(args) != 0 {
		return object.NewError("wrong number of arguments. tcp_conn.close() got=%d", len(args))
	}
	err := c.Handler.Close()
	if err != nil {
		return object.NewError("tcp_conn.close() error: %s", err.Error())
	}
	return &object.Null{}
}

func (c *TcpConnection) setReadBuffer(args []object.Object) object.Object {
	if len(args) != 1 {
		return object.NewError("wrong number of arguments. tcp_conn.set_read_buffer() got=%d", len(args))
	}
	switch arg := args[0].(type) {
	case *object.Integer:
		err := c.Handler.SetReadBuffer(int(arg.Value))
		if err != nil {
			return object.NewError("tcp_conn.set_read_buffer() error: %s", err.Error())
		}
		return &object.Null{}
	}
	return object.NewError("wrong type of arguments. tcp_conn.set_read_buffer(): %s", args[0].Type())
}

func (c *TcpConnection) setSendBuffer(args []object.Object) object.Object {
	if len(args) != 1 {
		return object.NewError("wrong number of arguments. tcp_conn.set_send_buffer() got=%d", len(args))
	}
	switch arg := args[0].(type) {
	case *object.Integer:
		err := c.Handler.SetWriteBuffer(int(arg.Value))
		if err != nil {
			return object.NewError("tcp_conn.set_send_buffer() error: %s", err.Error())
		}
		return &object.Null{}
	}
	return object.NewError("wrong type of arguments. tcp_conn.set_send_buffer(): %s", args[0].Type())
}

func (c *TcpConnection) setKeepAlive(args []object.Object) object.Object {
	if len(args) != 1 {
		return object.NewError("wrong number of arguments. tcp_conn.set_keep_alive() got=%d", len(args))
	}
	switch arg := args[0].(type) {
	case *object.Boolean:
		err := c.Handler.SetKeepAlive(arg.Value)
		if err != nil {
			return object.NewError("tcp_conn.set_keep_alive() error: %s", err.Error())
		}
		return &object.Null{}
	}
	return object.NewError("wrong type of arguments. tcp_conn.set_keep_alive(): %s", args[0].Type())
}

func (c *TcpConnection) setKeepAliveTime(args []object.Object) object.Object {
	if len(args) != 1 {
		return object.NewError("wrong number of arguments. tcp_conn.set_keep_alive_time() got=%d", len(args))
	}
	switch arg := args[0].(type) {
	case *object.Integer:
		err := c.Handler.SetKeepAlivePeriod(time.Second * time.Duration(arg.Value))
		if err != nil {
			return object.NewError("tcp_conn.set_keep_alive_time() error: %s", err.Error())
		}
		return &object.Null{}
	}
	return object.NewError("wrong type of arguments. tcp_conn.set_keep_alive_time(): %s", args[0].Type())
}

func (c *TcpConnection) setLinger(args []object.Object) object.Object {
	if len(args) != 1 {
		return object.NewError("wrong number of arguments. tcp_conn.set_linger() got=%d", len(args))
	}
	switch arg := args[0].(type) {
	case *object.Integer:
		err := c.Handler.SetLinger(int(arg.Value))
		if err != nil {
			return object.NewError("tcp_conn.set_linger() error: %s", err.Error())
		}
		return &object.Null{}
	}
	return object.NewError("wrong type of arguments. tcp_conn.set_linger(): %s", args[0].Type())
}

func (c *TcpConnection) setNoDelay(args []object.Object) object.Object {
	if len(args) != 1 {
		return object.NewError("wrong number of arguments. tcp_conn.set_no_delay() got=%d", len(args))
	}
	switch arg := args[0].(type) {
	case *object.Boolean:
		err := c.Handler.SetNoDelay(arg.Value)
		if err != nil {
			return object.NewError("tcp_conn.set_no_delay() error: %s", err.Error())
		}
		return &object.Null{}
	}
	return object.NewError("wrong type of arguments. tcp_conn.set_no_delay(): %s", args[0].Type())
}
