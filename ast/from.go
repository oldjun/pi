package ast

import (
	"pilang/token"
)

type From struct {
	Expression
	Token       token.Token
	File        string
	Identifiers map[string]*Identifier
	Everything  bool
}
