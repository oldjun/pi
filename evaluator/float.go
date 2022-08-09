package evaluator

import (
	"pilang/ast"
	"pilang/object"
)

func evalFloat(node *ast.Float) object.Object {
	return &object.Float{Value: node.Value}
}
