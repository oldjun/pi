package ast

import (
	"pilang/token"
)

type While struct {
	Expression
	Token       token.Token // The 'while' token
	Condition   Expression
	Consequence *Block
}
