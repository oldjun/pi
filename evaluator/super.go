package evaluator

import (
	"pilang/ast"
	"pilang/object"
)

func evalSuper(node *ast.Super, env *object.Environment) object.Object {
	if super, ok := env.Get("super"); ok {
		return super
	}
	return newError("runtime error: cannot call 'super' outside of scope")
}
