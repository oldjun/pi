package ast

import "github.com/oldjun/pi/token"

type Case struct {
	Expression
	Token   token.Token  // The "case" token
	Default bool         // Is this the default branch?
	Values  []Expression // The value of the case we'll be matching against
	Body    *Block       // The block that will be evaluated if matched
}
