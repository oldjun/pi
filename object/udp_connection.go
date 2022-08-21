package object

import (
	"fmt"
	"net"
)

type UpdConnection struct {
	Handler *net.UDPConn
}

func (c *UpdConnection) Type() Type { return "UDP_CONNECTION" }
func (c *UpdConnection) String() string {
	return fmt.Sprintf("<upd_connection:%v>", c.Handler)
}
