package evaluator

import (
	"github.com/oldjun/pi/ast"
	"github.com/oldjun/pi/object"
)

func evalPostfix(node *ast.Postfix, env *object.Environment) object.Object {
	left := Eval(node.Left, env)
	if isError(left) {
		return left
	}
	var val object.Object
	switch left := left.(type) {
	case *object.Integer:
		switch node.Operator {
		case "++":
			val = &object.Integer{Value: left.Value + 1}
		case "--":
			val = &object.Integer{Value: left.Value - 1}
		default:
			return object.NewError("postfix operator not support: %s", node.Operator)
		}
	case *object.Float:
		switch node.Operator {
		case "++":
			val = &object.Float{Value: left.Value + 1}
		case "--":
			val = &object.Float{Value: left.Value - 1}
		default:
			return object.NewError("postfix operator not support: %s", node.Operator)
		}
	default:
		return object.NewError("postfix not support: %s", left.Type())
	}
	switch nodeLeft := node.Left.(type) {
	case *ast.Identifier:
		env.Set(nodeLeft.Value, val)
		return NULL
	case *ast.IndexExpression:
		return evalIndexAssignment(nodeLeft, val, env)
	case *ast.PropertyExpression:
		return evalPropertyAssignment(nodeLeft, val, env)
	}
	return NULL
}
