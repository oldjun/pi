package object

import (
	"bytes"
	"pilang/ast"
	"strings"
)

type Function struct {
	Name       string
	Parameters []*ast.Identifier
	Body       *ast.Block
	Decorated  *ast.Function
	Env        *Environment
}

func (f *Function) Type() Type { return FUNCTION }
func (f *Function) String() string {
	var out bytes.Buffer
	var params []string
	for _, p := range f.Parameters {
		params = append(params, p.String())
	}
	out.WriteString("function")
	if f.Name != "" {
		out.WriteString(" " + f.Name)
	}
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") ")
	out.WriteString(f.Body.String())
	return out.String()
}
