package ast

import (
	"github.com/oldjun/pi/token"
)

type ForIn struct {
	Expression
	Token    token.Token // The 'for' token
	Key      string
	Value    string
	Iterable Expression // An expression that should return an iterable ([1, 2, 3] or x in 1..10)
	Block    *Block     // The block executed inside the for loop
}
