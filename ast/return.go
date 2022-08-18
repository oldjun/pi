package ast

import (
	"bytes"
	"pilang/token"
)

type Return struct {
	Token token.Token // the 'return' token
	Value Expression
}

func (r *Return) statementNode()       {}
func (r *Return) TokenLiteral() string { return r.Token.Literal }
func (r *Return) String() string {
	var out bytes.Buffer
	out.WriteString(r.Token.Literal + " ")
	if r.Value != nil {
		out.WriteString(r.Value.String())
	}
	out.WriteString(";")
	return out.String()
}
