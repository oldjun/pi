package evaluator

import (
	"github.com/oldjun/pi/ast"
	"github.com/oldjun/pi/object"
)

func evalInteger(node *ast.Integer) object.Object {
	return &object.Integer{Value: node.Value}
}
