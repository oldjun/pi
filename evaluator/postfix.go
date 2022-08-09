package evaluator

import (
	"pilang/ast"
	"pilang/object"
)

func evalPostfix(node *ast.Postfix, env *object.Environment) object.Object {
	left := Eval(node.Left, env)
	if isError(left) {
		return left
	}
	val := left.(*object.Integer)
	switch node.Operator {
	case "++":
		val = &object.Integer{Value: val.Value + 1}
	case "--":
		val = &object.Integer{Value: val.Value - 1}
	default:
		return nil
	}
	switch nodeLeft := node.Left.(type) {
	case *ast.Identifier:
		env.Set(nodeLeft.Value, val)
		return NULL
	case *ast.IndexExpression:
		// support assignment to index expression: h["a"]++
		return evalIndexAssignment(nodeLeft, val, env)
	case *ast.PropertyExpression:
		// support assignment to hash property: h.a++
		return evalPropertyAssignment(nodeLeft, val, env)
	}
	// otherwise
	env.Set(node.Left.String(), val)
	return NULL
}
