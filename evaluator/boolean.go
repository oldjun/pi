package evaluator

import (
	"pilang/ast"
	"pilang/object"
)

func evalBoolean(node *ast.Boolean) object.Object {
	return toBooleanObject(node.Value)
}

func toBooleanObject(input bool) *object.Boolean {
	if input {
		return TRUE
	} else {
		return FALSE
	}
}
