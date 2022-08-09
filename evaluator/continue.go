package evaluator

import (
	"pilang/ast"
	"pilang/object"
)

func evalContinue(node *ast.Continue) object.Object {
	return CONTINUE
}
