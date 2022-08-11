package evaluator

import (
	"pilang/ast"
	"pilang/object"
)

func evalNull(node *ast.Null) object.Object {
	return NULL
}
