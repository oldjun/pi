package module

import (
	"github.com/oldjun/pi/object"
	"net"
)

// NetProperties module properties
var NetProperties = map[string]object.ModuleProperty{}

// NetFunctions module functions
var NetFunctions = map[string]object.ModuleFunction{}

func init() {
	NetFunctions["listen"] = listen
	NetFunctions["connect"] = connect
}

func listen(args []object.Object) object.Object {
	if len(args) != 2 {
		return object.NewError("wrong number of arguments. net.listen() got=%d", len(args))
	}
	network := args[0].(*object.String).Value
	address := args[1].(*object.String).Value
	addr, err := net.ResolveTCPAddr(network, address)
	if err != nil {
		return object.NewError("net.resolve tcp addr error: %s", err.Error())
	}
	switch network {
	case "tcp", "tcp4", "tcp6":
		listener, err := net.ListenTCP(network, addr)
		if err != nil {
			return object.NewError("net.listen error: %s", err.Error())
		}
		return &object.TcpListener{Handler: listener}
	}
	return object.NewError("net.listen network error: %s", network)
}

func connect(args []object.Object) object.Object {
	if len(args) != 2 {
		return object.NewError("wrong number of arguments. net.connect() got=%d", len(args))
	}
	network := args[0].(*object.String).Value
	address := args[1].(*object.String).Value
	addr, err := net.ResolveTCPAddr(network, address)
	if err != nil {
		return object.NewError("net.resolve tcp addr error: %s", err.Error())
	}
	switch network {
	case "tcp", "tcp4", "tcp6":
		conn, err := net.DialTCP(network, nil, addr)
		if err != nil {
			return object.NewError("net.connect error: %s", err.Error())
		}
		return &object.TcpConnection{Handler: conn}
	}
	return object.NewError("net.connect network error: %s", network)
}
