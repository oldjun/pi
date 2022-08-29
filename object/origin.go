package object

import "fmt"

type Handler interface {
	Method(method string, args []Object) Object
}

type Origin struct {
	Name    string
	Handler Handler
}

func (o *Origin) Type() Type { return ORIGIN }
func (o *Origin) String() string {
	return fmt.Sprintf("<origin:%s>", o.Name)
}

func (o *Origin) Method(method string, args []Object) Object {
	return o.Handler.Method(method, args)
}
