package ast

import "github.com/oldjun/pi/token"

type Continue struct {
	Statement
	Token token.Token // the 'continue' token
}
