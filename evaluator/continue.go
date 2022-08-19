package evaluator

import (
	"github.com/oldjun/pi/ast"
	"github.com/oldjun/pi/object"
)

func evalContinue(node *ast.Continue) object.Object {
	return CONTINUE
}
