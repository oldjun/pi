package ast

import "pilang/token"

type Continue struct {
	Statement
	Token token.Token // the 'continue' token
}
