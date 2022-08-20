package evaluator

import (
	"github.com/oldjun/pi/ast"
	"github.com/oldjun/pi/object"
)

func evalAsync(node *ast.Async, env *object.Environment) object.Object {
	switch call := node.Call.(type) {
	case *ast.Call:
		scope := object.NewEnvironment(env.GetDirectory())
		for key, val := range env.All() {
			scope.Set(key, val)
		}
		go func(scope *object.Environment) {
			Eval(call, scope)
		}(scope)
	default:
		return object.NewError("async only support function")
	}
	return NULL
}
