package object

import (
	"fmt"
	"pilang/ast"
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

/*
func (f *Function) String() string {
	var out bytes.Buffer
	var params []string
	for _, p := range f.Parameters {
		if e, ok := f.Defaults[p.String()]; ok {
			params = append(params, p.String()+"="+e.String())
		} else {
			params = append(params, p.String())
		}
	}
	if f.Args != nil {
		params = append(params, token.ASTERISK+f.Args.String())
	}
	if f.KwArgs != nil {
		params = append(params, token.ASTERISK_ASTERISK+f.KwArgs.String())
	}
	out.WriteString("func")
	if f.Name != "" {
		out.WriteString(" " + f.Name)
	}
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") ")
	out.WriteString("{ ")
	out.WriteString(f.Body.String())
	out.WriteString(" }")
	return out.String()
}
*/
