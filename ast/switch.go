package ast

import "github.com/oldjun/pi/token"

type Switch struct {
	Expression
	Token token.Token // The "switch" token
	Value Expression  // The value that will be used to determine the case
	Cases []*Case     // The cases this switch statement will handle
}
