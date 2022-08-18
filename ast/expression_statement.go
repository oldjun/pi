package ast

import "pilang/token"

type ExpressionStatement struct {
	Statement
	Token      token.Token // the first token of the expression
	Expression Expression
}
