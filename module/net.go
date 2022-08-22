package module

import (
	module "github.com/oldjun/pi/module/net"
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
	switch network {
	case "tcp":
		addr, err := net.ResolveTCPAddr(network, address)
		if err != nil {
			return object.NewError("net.resolve tcp addr error: %s", err.Error())
		}
		listener, err := net.ListenTCP(network, addr)
		if err != nil {
			return object.NewError("net.listen error: %s", err.Error())
		}
		return &object.Module{Name: "listener", Handler: &module.TcpListener{Handler: listener}}
	case "unix":
		addr, err := net.ResolveUnixAddr(network, address)
		if err != nil {
			return object.NewError("net.resolve unix addr error: %s", err.Error())
		}
		listener, err := net.ListenUnix(network, addr)
		if err != nil {
			return object.NewError("net.listen error: %s", err.Error())
		}
		return &object.Module{Name: "listener", Handler: &module.UnixListener{Handler: listener}}
	}
	return object.NewError("net.listen network type error: %s", network)
}

func connect(args []object.Object) object.Object {
	if len(args) != 2 {
		return object.NewError("wrong number of arguments. net.connect() got=%d", len(args))
	}
	network := args[0].(*object.String).Value
	address := args[1].(*object.String).Value
	switch network {
	case "tcp":
		addr, err := net.ResolveTCPAddr(network, address)
		if err != nil {
			return object.NewError("net.resolve tcp addr error: %s", err.Error())
		}
		conn, err := net.DialTCP(network, nil, addr)
		if err != nil {
			return object.NewError("net.connect error: %s", err.Error())
		}
		return &object.Module{Name: "connection", Handler: &module.TcpConnection{Handler: conn}}
	case "unix":
		addr, err := net.ResolveUnixAddr(network, address)
		if err != nil {
			return object.NewError("net.resolve unix addr error: %s", err.Error())
		}
		conn, err := net.DialUnix(network, nil, addr)
		if err != nil {
			return object.NewError("net.connect error: %s", err.Error())
		}
		return &object.Module{Name: "connection", Handler: &module.UnixConnection{Handler: conn}}
	}
	return object.NewError("net.connect network error: %s", network)
}
