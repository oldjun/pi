package ast

import "pilang/token"

type Switch struct {
	Token token.Token // The "switch" token
	Value Expression  // The value that will be used to determine the case
	Cases []*Case     // The cases this switch statement will handle
}

func (s *Switch) expressionNode()      {}
func (s *Switch) TokenLiteral() string { return s.Token.Literal }
func (s *Switch) String() string {
	return "switch"
}
