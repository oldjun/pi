package ast

import (
	"bytes"
	"pilang/token"
	"strings"
)

type From struct {
	Token       token.Token
	File        string
	Identifiers map[string]*Identifier
	Everything  bool
}

func (f *From) expressionNode()      {}
func (f *From) TokenLiteral() string { return f.Token.Literal }
func (f *From) String() string {
	var out bytes.Buffer
	out.WriteString(" from " + f.File)
	out.WriteString("import ")
	var pairs []string
	for key, val := range f.Identifiers {
		if key != val.Value {
			pairs = append(pairs, val.Value+" as "+key)
		} else {
			pairs = append(pairs, key)
		}
	}
	out.WriteString(strings.Join(pairs, ", "))
	return out.String()
}
