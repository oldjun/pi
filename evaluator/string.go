package evaluator

import (
	"pilang/ast"
	"pilang/object"
)

func evalString(node *ast.String) object.Object {
	return &object.String{Value: node.Value}
}
