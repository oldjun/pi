package ast

import (
	"pilang/token"
)

type Scenario struct {
	Condition   Expression
	Consequence *Block
}

type If struct {
	Expression
	Token     token.Token // The 'if' token
	Scenarios []*Scenario
}
