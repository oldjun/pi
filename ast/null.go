package ast

import "pilang/token"

type Null struct {
	Expression
	Token token.Token // the 'null' token
}
