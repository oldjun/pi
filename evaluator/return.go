package evaluator

import (
	"pilang/ast"
	"pilang/object"
)

func evalReturn(node *ast.Return, env *object.Environment) object.Object {
	val := Eval(node.Value, env)
	if isError(val) {
		return val
	}
	return &object.Return{Value: val}
}
