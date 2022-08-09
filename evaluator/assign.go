package evaluator

import (
	"pilang/ast"
	"pilang/object"
)

func evalAssign(node *ast.Assign, env *object.Environment) object.Object {
	val := Eval(node.Value, env)
	if isError(val) {
		return val
	}
	env.Set(node.Name.Value, val)
	return NULL
}
