package module

import (
	"net"
)

type UpdConnection struct {
	Handler *net.UDPConn
}
