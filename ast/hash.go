package ast

import (
	"bytes"
	"pilang/token"
	"strings"
)

type Hash struct {
	token.Token // the '{' token
	Pairs       map[Expression]Expression
}

func (h *Hash) expressionNode()      {}
func (h *Hash) TokenLiteral() string { return h.Token.Literal }
func (h *Hash) String() string {
	var out bytes.Buffer
	var pairs []string
	for key, value := range h.Pairs {
		pairs = append(pairs, key.String()+":"+value.String())
	}
	out.WriteString("{")
	out.WriteString(strings.Join(pairs, ", "))
	out.WriteString("}")
	return out.String()
}
