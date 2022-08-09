package ast

import "pilang/token"

type Continue struct {
	Token token.Token // the 'continue' token
}

func (c *Continue) statementNode()       {}
func (c *Continue) TokenLiteral() string { return c.Token.Literal }
func (c *Continue) String() string {
	return "continue"
}
