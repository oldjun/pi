package ast

import (
	"bytes"
	"pilang/token"
	"strings"
)

type Import struct {
	Token       token.Token
	Identifiers map[string]*Identifier
}

func (i *Import) expressionNode()      {}
func (i *Import) TokenLiteral() string { return i.Token.Literal }
func (i *Import) String() string {
	var out bytes.Buffer
	out.WriteString("import ")
	var arr []string
	for key, val := range i.Identifiers {
		if key != val.Value {
			arr = append(arr, val.Value+" as "+key)
		} else {
			arr = append(arr, val.Value)
		}
	}
	out.WriteString(strings.Join(arr, ", "))
	return out.String()
}
