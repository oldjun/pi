package evaluator

import (
	"pilang/object"
	"strings"
)

func evalInfix(operator string, left, right object.Object) object.Object {
	switch {
	case left.Type() == object.INTEGER && right.Type() == object.INTEGER:
		return evalIntegerInfixExpression(operator, left, right)
	case left.Type() == object.FLOAT && right.Type() == object.FLOAT:
		return evalFloatInfixExpression(operator, left, right)
	case left.Type() == object.INTEGER && right.Type() == object.FLOAT:
		return evalFloatInfixExpression(operator, castFromIntegerToFloat(left), right)
	case left.Type() == object.FLOAT && right.Type() == object.INTEGER:
		return evalFloatInfixExpression(operator, left, castFromIntegerToFloat(right))
	case left.Type() == object.STRING && right.Type() == object.STRING:
		return evalStringInfixExpression(operator, left, right)
	case left.Type() == object.STRING && right.Type() == object.INTEGER:
		return evalStringIntegerInfixExpression(operator, left, right)
	case operator == "in":
		return evalInExpression(left, right)
	case operator == "==":
		return toBooleanObject(left == right)
	case operator == "!=":
		return toBooleanObject(left != right)
	case operator == "and" || operator == "or" || operator == "&&" || operator == "||":
		return evalBooleanInfixExpression(operator, left, right)
	case left.Type() != right.Type():
		return newError("type mismatch: %s %s %s",
			left.Type(), operator, right.Type())
	default:
		return newError("unknown operator: %s %s %s",
			left.Type(), operator, right.Type())
	}
}

func evalBooleanInfixExpression(operator string, left, right object.Object) object.Object {
	leftVal := left.(*object.Boolean).Value
	rightVal := right.(*object.Boolean).Value
	switch {
	case operator == "and" || operator == "&&":
		if leftVal && rightVal {
			return TRUE
		} else {
			return FALSE
		}
	case operator == "or" || operator == "||":
		if leftVal || rightVal {
			return TRUE
		} else {
			return FALSE
		}
	default:
		return newError("unknown operator: %s %s %s",
			left.Type(), operator, right.Type())
	}
}

func evalIntegerInfixExpression(operator string, left, right object.Object) object.Object {
	leftVal := left.(*object.Integer).Value
	rightVal := right.(*object.Integer).Value
	switch operator {
	case "+":
		return &object.Integer{Value: leftVal + rightVal}
	case "-":
		return &object.Integer{Value: leftVal - rightVal}
	case "*":
		return &object.Integer{Value: leftVal * rightVal}
	case "/":
		return &object.Integer{Value: leftVal / rightVal}
	case "%":
		return &object.Integer{Value: leftVal % rightVal}
	case "&":
		return &object.Integer{Value: leftVal & rightVal}
	case "|":
		return &object.Integer{Value: leftVal | rightVal}
	case "^":
		return &object.Integer{Value: leftVal ^ rightVal}
	case "<<":
		return &object.Integer{Value: leftVal << rightVal}
	case ">>":
		return &object.Integer{Value: leftVal >> rightVal}
	case "<":
		return toBooleanObject(leftVal < rightVal)
	case ">":
		return toBooleanObject(leftVal > rightVal)
	case "<=":
		return toBooleanObject(leftVal <= rightVal)
	case ">=":
		return toBooleanObject(leftVal >= rightVal)
	case "==":
		return toBooleanObject(leftVal == rightVal)
	case "!=":
		return toBooleanObject(leftVal != rightVal)
	default:
		return newError("unknown operator: %s %s %s",
			left.Type(), operator, right.Type())
	}
}

func evalInExpression(left, right object.Object) object.Object {
	switch right.(type) {
	case *object.String:
		return evalInStringExpression(left, right)
	case *object.List:
		return evalInArrayExpression(left, right)
	case *object.Hash:
		return evalInHashExpression(left, right)
	default:
		return FALSE
	}
}

func evalInStringExpression(left, right object.Object) object.Object {
	if left.Type() != object.STRING {
		return FALSE
	}
	leftVal := left.(*object.String)
	rightVal := right.(*object.String)
	found := strings.Contains(rightVal.Value, leftVal.Value)
	return toBooleanObject(found)
}

func evalInHashExpression(left, right object.Object) object.Object {
	leftVal, ok := left.(object.Hashable)
	if !ok {
		return newError("unusable as hash key: %s", left.Type())
	}
	key := leftVal.HashKey()
	rightVal := right.(*object.Hash).Pairs
	_, ok = rightVal[key]
	return toBooleanObject(ok)
}

func evalInArrayExpression(left, right object.Object) object.Object {
	rightVal := right.(*object.List)
	switch leftVal := left.(type) {
	case *object.String:
		for _, v := range rightVal.Elements {
			if v.Type() == object.STRING {
				elem := v.(*object.String)
				if elem.Value == leftVal.Value {
					return TRUE
				}
			}
		}
	case *object.Integer:
		for _, v := range rightVal.Elements {
			if v.Type() == object.INTEGER {
				elem := v.(*object.Integer)
				if elem.Value == leftVal.Value {
					return TRUE
				}
			}
		}
	case *object.Float:
		for _, v := range rightVal.Elements {
			if v.Type() == object.FLOAT {
				elem := v.(*object.Float)
				if elem.Value == leftVal.Value {
					return TRUE
				}
			}
		}
	}
	return FALSE
}

func evalFloatInfixExpression(operator string, left, right object.Object) object.Object {
	leftVal := left.(*object.Float).Value
	rightVal := right.(*object.Float).Value
	switch operator {
	case "+":
		return &object.Float{Value: leftVal + rightVal}
	case "-":
		return &object.Float{Value: leftVal - rightVal}
	case "*":
		return &object.Float{Value: leftVal * rightVal}
	case "/":
		return &object.Float{Value: leftVal / rightVal}
	case "<":
		return toBooleanObject(leftVal < rightVal)
	case ">":
		return toBooleanObject(leftVal > rightVal)
	case "<=":
		return toBooleanObject(leftVal <= rightVal)
	case ">=":
		return toBooleanObject(leftVal >= rightVal)
	case "==":
		return toBooleanObject(leftVal == rightVal)
	case "!=":
		return toBooleanObject(leftVal != rightVal)
	default:
		return newError("unknown operator: %s %s %s",
			left.Type(), operator, right.Type())
	}
}

func evalStringInfixExpression(operator string, left, right object.Object) object.Object {
	leftVal := left.(*object.String).Value
	rightVal := right.(*object.String).Value
	switch operator {
	case "+":
		return &object.String{Value: leftVal + rightVal}
	case "==":
		return &object.Boolean{Value: leftVal == rightVal}
	case "!=":
		return &object.Boolean{Value: leftVal != rightVal}
	case ">":
		return &object.Boolean{Value: leftVal > rightVal}
	case "<":
		return &object.Boolean{Value: leftVal < rightVal}
	case "in":
		return evalInStringExpression(left, right)
	default:
		return newError("unknown operator: %s %s %s",
			left.Type(), operator, right.Type())
	}
}

func evalStringIntegerInfixExpression(operator string, left, right object.Object) object.Object {
	leftVal := left.(*object.String).Value
	rightVal := right.(*object.Integer).Value
	switch operator {
	case "*":
		val := strings.Repeat(leftVal, int(rightVal))
		return &object.String{Value: val}
	default:
		return newError("unknown operator: %s %s %s",
			left.Type(), operator, right.Type())
	}
}
