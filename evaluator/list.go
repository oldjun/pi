package evaluator

import (
	"pilang/ast"
	"pilang/object"
)

func evalList(node *ast.List, env *object.Environment) object.Object {
	elements := evalExpressions(node.Elements, env)
	if len(elements) == 1 && isError(elements[0]) {
		return elements[0]
	}
	return &object.List{Elements: elements}
}
