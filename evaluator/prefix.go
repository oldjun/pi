package evaluator

import (
	"pilang/ast"
	"pilang/object"
)

func evalPrefix(node *ast.Prefix, env *object.Environment) object.Object {
	operator := node.Operator
	right := Eval(node.Right, env)
	if isError(right) {
		return right
	}
	switch operator {
	case "!":
		return evalBangOperatorExpression(right)
	case "not":
		return evalBangOperatorExpression(right)
	case "-":
		return evalMinusPrefixOperatorExpression(right)
	case "~":
		return evalTildePrefixOperatorExpression(right)
	default:
		return newError("unknown operator: %s%s", operator, right.Type())
	}
}

func evalBangOperatorExpression(right object.Object) object.Object {
	switch right := right.(type) {
	case *object.Null:
		return &object.Boolean{Value: true}
	case *object.Boolean:
		return &object.Boolean{Value: !right.Value}
	default:
		return &object.Boolean{Value: false}
	}
}

func evalMinusPrefixOperatorExpression(right object.Object) object.Object {
	if right.Type() != object.INTEGER {
		return newError("unknown operator: -%s", right.Type())
	}
	value := right.(*object.Integer).Value
	return &object.Integer{Value: -value}
}

func evalTildePrefixOperatorExpression(right object.Object) object.Object {
	if right.Type() != object.INTEGER {
		return newError("unknown operator: ~%s", right.Type())
	}
	value := right.(*object.Integer).Value
	return &object.Integer{Value: ^value}
}
