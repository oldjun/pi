package evaluator

import (
	"github.com/oldjun/pi/ast"
	"github.com/oldjun/pi/object"
)

func evalFloat(node *ast.Float) object.Object {
	return &object.Float{Value: node.Value}
}
