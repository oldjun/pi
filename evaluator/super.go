package evaluator

import (
	"pilang/ast"
	"pilang/object"
)

func evalSuper(node *ast.Super, env *object.Environment) object.Object {
	this, ok := env.Get("this")
	if ok {
		if super, ok := this.(*object.Instance).Env.Get("super"); ok {
			return super
		}
	}
	return object.NewError("runtime error: cannot call 'super' outside of scope")
}
