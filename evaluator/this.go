package evaluator

import (
	"github.com/oldjun/pi/ast"
	"github.com/oldjun/pi/object"
)

func evalThis(node *ast.This, env *object.Environment) object.Object {
	if this, ok := env.Get("this"); ok {
		return this
	}
	return object.NewError("runtime error: cannot call 'this' outside of scope")
}
