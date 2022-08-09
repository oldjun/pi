package object

import "fmt"

type Instance struct {
	Class *Class
	Env   *Environment
}

func (i *Instance) Type() Type { return INSTANCE }
func (i *Instance) String() string {
	return fmt.Sprintf("<instance:%s>", i.Class.Name.Value)
}
