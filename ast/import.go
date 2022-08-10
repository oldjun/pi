package ast

import (
	"bytes"
	"pilang/token"
	"strings"
)

type Import struct {
	Token       token.Token
	File        string
	Identifiers map[string]*Identifier
	Everything  bool
}

func (i *Import) expressionNode()      {}
func (i *Import) TokenLiteral() string { return i.Token.Literal }
func (i *Import) String() string {
	var out bytes.Buffer
	out.WriteString("import ")
	var pairs []string
	for key, _ := range i.Identifiers {
		pairs = append(pairs, key)
	}
	out.WriteString(strings.Join(pairs, ", "))
	out.WriteString(" from " + i.File)
	return out.String()
}
