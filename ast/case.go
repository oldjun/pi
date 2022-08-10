package ast

import "pilang/token"

type Case struct {
	Token   token.Token  // The "case" token
	Default bool         // Is this the default branch?
	Values  []Expression // The value of the case we'll be matching against
	Body    *Block       // The block that will be evaluated if matched
}

func (c *Case) expressionNode()      {}
func (c *Case) TokenLiteral() string { return c.Token.Literal }
func (c *Case) String() string {
	return "case"
}
