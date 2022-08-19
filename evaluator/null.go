package evaluator

import (
	"github.com/oldjun/pi/ast"
	"github.com/oldjun/pi/object"
)

func evalNull(node *ast.Null) object.Object {
	return NULL
}
