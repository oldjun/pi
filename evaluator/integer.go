package evaluator

import (
	"pilang/ast"
	"pilang/object"
)

func evalInteger(node *ast.Integer) object.Object {
	return &object.Integer{Value: node.Value}
}
