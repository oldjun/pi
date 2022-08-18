package ast

import "pilang/token"

type Identifier struct {
	Expression
	Token token.Token // the token.IDENTIFIER token
	Value string
}
