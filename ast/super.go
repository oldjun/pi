package ast

import (
	"bytes"
	"pilang/token"
)

type Super struct {
	Token token.Token
}

func (s *Super) expressionNode()      {}
func (s *Super) TokenLiteral() string { return s.Token.Literal }
func (s *Super) String() string {
	var out bytes.Buffer
	out.WriteString("super")
	return out.String()
}
