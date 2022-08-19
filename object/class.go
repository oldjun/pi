package object

import (
	"fmt"
	"github.com/oldjun/pi/ast"
)

type Class struct {
	Name  *ast.Identifier
	Super *Class
	Env   *Environment
	Scope *Environment
}

func (c *Class) Type() Type { return CLASS }
func (c *Class) String() string {
	return fmt.Sprintf("<class:%s>", c.Name.Value)
}
