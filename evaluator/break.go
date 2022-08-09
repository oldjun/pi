package evaluator

import (
	"pilang/ast"
	"pilang/object"
)

func evalBreak(node *ast.Break) object.Object {
	return BREAK
}
