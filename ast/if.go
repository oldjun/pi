package ast

import (
	"github.com/oldjun/pi/token"
)

type Scenario struct {
	Condition Expression
	Block     *Block
}

type If struct {
	Expression
	Token     token.Token // The 'if' token
	Scenarios []*Scenario
}
