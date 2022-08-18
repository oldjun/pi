package ast

import (
	"pilang/token"
)

type Block struct {
	Statement
	Token      token.Token // the { token
	Statements []Statement
}
