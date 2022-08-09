package ast

import "pilang/token"

type String struct {
	Token token.Token
	Value string
}

func (s *String) expressionNode()      {}
func (s *String) TokenLiteral() string { return s.Token.Literal }
func (s *String) String() string       { return s.Token.Literal }
