package ast

import (
	"github.com/oldjun/pi/token"
)

type From struct {
	Expression
	Token       token.Token
	File        string
	Identifiers map[string]*Identifier
	Everything  bool
}
