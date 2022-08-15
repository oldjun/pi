package ast

import (
	"bytes"
	"pilang/token"
	"strings"
)

type Function struct {
	Token      token.Token // The 'func' token
	Name       string      // name of the function
	Parameters []*Identifier
	Defaults   map[string]Expression
	Args       *Identifier
	KwArgs     *Identifier
	Body       *Block
}

func (f *Function) expressionNode()      {}
func (f *Function) TokenLiteral() string { return f.Token.Literal }
func (f *Function) String() string {
	var out bytes.Buffer
	var params []string
	for _, p := range f.Parameters {
		params = append(params, p.String())
	}
	out.WriteString(f.TokenLiteral())
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") ")
	out.WriteString(f.Body.String())
	return out.String()
}
