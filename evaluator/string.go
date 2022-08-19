package evaluator

import (
	"github.com/oldjun/pi/ast"
	"github.com/oldjun/pi/object"
)

func evalString(node *ast.String) object.Object {
	return &object.String{Value: node.Value}
}
