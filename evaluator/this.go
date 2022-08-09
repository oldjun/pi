package evaluator

import (
	"pilang/ast"
	"pilang/object"
)

func evalThis(node *ast.This, env *object.Environment) object.Object {
	if this, ok := env.Get("this"); ok {
		return this
	}
	return newError("runtime error: cannot call 'this' outside of scope")
}
