package evaluator

import (
	"github.com/oldjun/pi/ast"
	"github.com/oldjun/pi/object"
)

func evalBreak(node *ast.Break) object.Object {
	return BREAK
}
