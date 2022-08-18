package ast

import (
	"pilang/token"
)

type For struct {
	Expression
	Token      token.Token // The 'for' token
	Identifier string      // "i"
	Starter    Expression  // i = 0
	Closer     Expression  // i++
	Condition  Expression  // i < 1
	Block      *Block      // The block executed inside the for loop
}
