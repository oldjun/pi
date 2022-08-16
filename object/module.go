package object

import "fmt"

type Module struct {
	Name  string
	Value Object
}

func (m *Module) Type() Type { return MODULE }
func (m *Module) String() string {
	return fmt.Sprintf("<module:%s>", m.Name)
}

var ModuleMap map[string]*Module

func init() {
	ModuleMap = make(map[string]*Module)
	ModuleMap["math"] = &Module{Name: "math", Value: &Math{}}
}
