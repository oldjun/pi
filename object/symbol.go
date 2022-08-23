package object

import "fmt"

type Handler interface {
	Method(method string, args []Object) Object
}

type Symbol struct {
	Name    string
	Handler Handler
}

func (s *Symbol) Type() Type { return SYMBOL }
func (s *Symbol) String() string {
	return fmt.Sprintf("<symbol:%s>", s.Name)
}

func (s *Symbol) Method(method string, args []Object) Object {
	return s.Handler.Method(method, args)
}
