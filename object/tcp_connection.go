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

func (c *TcpConnection) setReadBuffer(args []Object) Object {
	if len(args) != 1 {
		return NewError("wrong number of arguments. tcp_conn.set_read_buffer() got=%d", len(args))
	}
	switch arg := args[0].(type) {
	case *Integer:
		err := c.Handler.SetReadBuffer(int(arg.Value))
		if err != nil {
			return NewError("tcp_conn.set_read_buffer() error: %s", err.Error())
		}
		return &Null{}
	}
	return NewError("wrong type of arguments. tcp_conn.set_read_buffer(): %s", args[0].Type())
}

func (c *TcpConnection) setSendBuffer(args []Object) Object {
	if len(args) != 1 {
		return NewError("wrong number of arguments. tcp_conn.set_send_buffer() got=%d", len(args))
	}
	switch arg := args[0].(type) {
	case *Integer:
		err := c.Handler.SetWriteBuffer(int(arg.Value))
		if err != nil {
			return NewError("tcp_conn.set_send_buffer() error: %s", err.Error())
		}
		return &Null{}
	}
	return NewError("wrong type of arguments. tcp_conn.set_send_buffer(): %s", args[0].Type())
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

func (c *TcpConnection) setLinger(args []Object) Object {
	if len(args) != 1 {
		return NewError("wrong number of arguments. tcp_conn.set_linger() got=%d", len(args))
	}
	switch arg := args[0].(type) {
	case *Integer:
		err := c.Handler.SetLinger(int(arg.Value))
		if err != nil {
			return NewError("tcp_conn.set_linger() error: %s", err.Error())
		}
		return &Null{}
	}
	return NewError("wrong type of arguments. tcp_conn.set_linger(): %s", args[0].Type())
}

func (c *TcpConnection) setNoDelay(args []Object) Object {
	if len(args) != 1 {
		return NewError("wrong number of arguments. tcp_conn.set_no_delay() got=%d", len(args))
	}
	switch arg := args[0].(type) {
	case *Boolean:
		err := c.Handler.SetNoDelay(arg.Value)
		if err != nil {
			return NewError("tcp_conn.set_no_delay() error: %s", err.Error())
		}
		return &Null{}
	}
	return NewError("wrong type of arguments. tcp_conn.set_no_delay(): %s", args[0].Type())
}
