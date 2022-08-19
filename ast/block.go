package ast

import (
	"github.com/oldjun/pi/token"
)

type Block struct {
	Statement
	Token      token.Token // the { token
	Statements []Statement
}
