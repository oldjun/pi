package object

import (
	"fmt"
	"github.com/oldjun/pi/ast"
)

type Function struct {
	Name       string
	Parameters []*ast.Identifier
	Defaults   map[string]ast.Expression
	Args       *ast.Identifier
	KwArgs     *ast.Identifier
	Body       *ast.Block
	Env        *Environment
}

func (f *Function) Type() Type { return FUNCTION }
func (f *Function) String() string {
	return fmt.Sprintf("<func:%s>", f.Name)
}
