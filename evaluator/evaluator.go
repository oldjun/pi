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
	//case *ast.Decorator:
	//	return evalDecorator(node, env)
	case *ast.Call:
		return evalCall(node, env)
	case *ast.Array:
		return evalArray(node, env)
	case *ast.Hash:
		return evalHash(node, env)
	case *ast.Class:
		return evalClass(node, env)
	case *ast.This:
		return evalThis(node, env)
	case *ast.Super:
		return evalSuper(node, env)
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

func applyFunction(node object.Object, args []object.Object) object.Object {
	switch node := node.(type) {
	case *object.Function:
		extendedEnv := extendFunctionEnv(node, args)
		evaluated := Eval(node.Body, extendedEnv)
		return unwrapReturnValue(evaluated)
	case *object.Builtin:
		return node.Fn(args...)
	case *object.Class:
		obj := &object.Instance{Class: node, Env: object.NewEnvironment()}
		obj.Env.Set("this", obj)
		if node.Super != nil {
			super := &object.Instance{Class: node.Super, Env: object.NewEnvironment()}
			obj.Env.Set("super", super)
		}
		fn, ok := node.Scope.Get("__init__")
		if !ok {
			return newError("%s missing __init__ function", node.String())
		}
		fn.(*object.Function).Env.Set("this", obj)
		applyFunction(fn, args)
		fn.(*object.Function).Env.Del("this")
		return obj
	default:
		return newError("not a function: %s", node.Type())
	}
}

func extendFunctionEnv(fn *object.Function, args []object.Object) *object.Environment {
	env := object.NewEnclosedEnvironment(fn.Env)
	for paramIdx, param := range fn.Parameters {
		env.Set(param.Value, args[paramIdx])
	}
	return env
}

func unwrapReturnValue(obj object.Object) object.Object {
	if ret, ok := obj.(*object.Return); ok {
		return ret.Value
	}
	return obj
}

func castFromIntegerToFloat(obj object.Object) object.Object {
	val := obj.(*object.Integer).Value
	return &object.Float{Value: float64(val)}
}

func isTruthy(obj object.Object) bool {
	switch obj {
	case NULL:
		return false
	case TRUE:
		return true
	case FALSE:
		return false
	default:
		return true
	}
}

func newError(format string, a ...interface{}) *object.Error {
	return &object.Error{Message: fmt.Sprintf(format, a...)}
}
