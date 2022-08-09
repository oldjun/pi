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
	switch right {
	case TRUE:
		return FALSE
	case FALSE:
		return TRUE
	case NULL:
		return TRUE
	default:
		return FALSE
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
