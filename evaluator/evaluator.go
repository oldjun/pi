package evaluator

import (
	"fmt"
	"pilang/ast"
	"pilang/object"
)

var (
	NULL     = &object.Null{}
	TRUE     = &object.Boolean{Value: true}
	FALSE    = &object.Boolean{Value: false}
	BREAK    = &object.Break{}
	CONTINUE = &object.Continue{}
)

func Eval(node ast.Node, env *object.Environment) object.Object {
	switch node := node.(type) {
	case *ast.Program:
		return evalProgram(node.Statements, env)
	case *ast.ExpressionStatement:
		return Eval(node.Expression, env)
	case *ast.Null:
		return evalNull(node)
	case *ast.Integer:
		return evalInteger(node)
	case *ast.Float:
		return evalFloat(node)
	case *ast.Boolean:
		return evalBoolean(node)
	case *ast.String:
		return evalString(node)
	case *ast.Prefix:
		return evalPrefix(node, env)
	case *ast.Infix:
		left := Eval(node.Left, env)
		if isError(left) {
			return left
		}
		right := Eval(node.Right, env)
		if isError(right) {
			return right
		}
		return evalInfix(node.Operator, left, right)
	case *ast.Postfix:
		return evalPostfix(node, env)
	case *ast.Block:
		return evalBlock(node, env)
	case *ast.If:
		return evalIf(node, env)
	case *ast.While:
		return evalWhile(node, env)
	case *ast.For:
		return evalFor(node, env)
	case *ast.ForIn:
		return evalForIn(node, env)
	case *ast.Identifier:
		return evalIdentifier(node, env)
	case *ast.Assign:
		return evalAssign(node, env)
	case *ast.Compound:
		return evalCompound(node, env)
	case *ast.IndexExpression:
		return evalIndexExpression(node, env)
	case *ast.IndexAssignment:
		val := Eval(node.Value, env)
		if isError(val) {
			return val
		}
		return evalIndexAssignment(node.Name, val, env)
	case *ast.PropertyExpression:
		return evalPropertyExpression(node, env)
	case *ast.PropertyAssignment:
		val := Eval(node.Value, env)
		if isError(val) {
			return val
		}
		return evalPropertyAssignment(node.Name, val, env)
	case *ast.Method:
		return evalMethod(node, env)
	case *ast.Ternary:
		return evalTernary(node, env)
	case *ast.Return:
		return evalReturn(node, env)
	case *ast.Function:
		return evalFunction(node, env)
	case *ast.Decorator:
		return evalDecorator(node, env)
	case *ast.Call:
		return evalCall(node, env)
	case *ast.List:
		return evalList(node, env)
	case *ast.Hash:
		return evalHash(node, env)
	case *ast.Class:
		return evalClass(node, env)
	case *ast.This:
		return evalThis(node, env)
	case *ast.Super:
		return evalSuper(node, env)
	case *ast.From:
		return evalFrom(node, env)
	case *ast.Switch:
		return evalSwitch(node, env)
	case *ast.Break:
		return evalBreak(node)
	case *ast.Continue:
		return evalContinue(node)
	}
	return NULL
}

func isError(obj object.Object) bool {
	if obj != nil {
		return obj.Type() == object.ERROR
	}
	return false
}

func castFromIntegerToFloat(obj object.Object) object.Object {
	val := obj.(*object.Integer).Value
	return &object.Float{Value: float64(val)}
}

func isTruthy(obj object.Object) bool {
	switch obj := obj.(type) {
	case *object.Null:
		return false
	case *object.Boolean:
		return obj.Value
	}
	return true
}

func newError(format string, a ...interface{}) *object.Error {
	return &object.Error{Message: fmt.Sprintf(format, a...)}
}
