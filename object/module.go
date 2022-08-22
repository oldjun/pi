package object

import (
	"fmt"
)

type ModuleProperty func() Object
type ModuleFunction func(args []Object) Object

type Handler interface {
	Method(method string, args []Object) Object
}

type Module struct {
	Name       string
	Functions  map[string]ModuleFunction
	Properties map[string]ModuleProperty
	Handler    Handler
}

func (m *Module) Type() Type { return MODULE }
func (m *Module) String() string {
	return fmt.Sprintf("<module:%s>", m.Name)
}

func (m *Module) Method(method string, args []Object) Object {
	return m.Handler.Method(method, args)
}
