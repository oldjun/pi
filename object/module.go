package object

import (
	"fmt"
)

type ModuleProperty func() Object
type ModuleFunction func(args []Object) Object

type Module struct {
	Name       string
	Functions  map[string]ModuleFunction
	Properties map[string]ModuleProperty
}

func (m *Module) Type() Type { return MODULE }
func (m *Module) String() string {
	return fmt.Sprintf("<module:%s>", m.Name)
}
