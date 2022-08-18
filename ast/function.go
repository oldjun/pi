package ast

import (
	"pilang/token"
)

type Function struct {
	Expression
	Token      token.Token // The 'func' token
	Name       string      // name of the function
	Parameters []*Identifier
	Defaults   map[string]Expression
	Args       *Identifier
	KwArgs     *Identifier
	Body       *Block
}
