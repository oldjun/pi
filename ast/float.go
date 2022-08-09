package ast

import "pilang/token"

type Float struct {
	Token token.Token
	Value float64
}

func (f *Float) expressionNode()      {}
func (f *Float) TokenLiteral() string { return f.Token.Literal }
func (f *Float) String() string       { return f.Token.Literal }
