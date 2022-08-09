package ast

import (
	"bytes"
	"pilang/token"
)

type IndexExpression struct {
	Token   token.Token // The [ token
	Left    Expression
	Index   Expression
	IsRange bool       // whether the expression is a range [1:10]
	End     Expression // the end of the range, if the expression is a range
}

func (ie *IndexExpression) expressionNode()      {}
func (ie *IndexExpression) TokenLiteral() string { return ie.Token.Literal }
func (ie *IndexExpression) String() string {
	var out bytes.Buffer
	out.WriteString(ie.Left.String())
	out.WriteString("[")
	if ie.IsRange {
		start := ""
		if ie.Index != nil {
			start = ie.Index.String()
		}
		end := ""
		if ie.End != nil {
			end = ie.End.String()
		}
		out.WriteString(start + ":" + end)
	} else {
		out.WriteString(ie.Index.String())
	}
	out.WriteString("]")
	return out.String()
}
