package evaluator

import (
	"pilang/ast"
	"pilang/object"
)

func evalCall(node *ast.Call, env *object.Environment) object.Object {
	function := Eval(node.Function, env)
	if isError(function) {
		return function
	}
	// if decorated function is a class method, env should carry `this`
	if this, ok := env.Get("this"); ok {
		switch function.(type) {
		case *object.Function:
			function.(*object.Function).Env.Set("this", this)
		}
	}
	args := evalExpressions(node.Arguments, env)
	if len(args) == 1 && isError(args[0]) {
		return args[0]
	}
	return applyFunction(function, args)
}

func applyFunction(node object.Object, args []object.Object) object.Object {
	switch node := node.(type) {
	case *object.Function:
		extendedEnv := extendFunctionEnv(node, args)
		evaluated := Eval(node.Body, extendedEnv)
		return unwrapReturnValue(evaluated)
	case *object.Builtin:
		return node.Fn(args)
	case *object.Class:
		obj := &object.Instance{Class: node, Env: object.NewEnvironment(node.Env.GetDirectory())}
		obj.Env.Set("this", obj)
		if node.Super != nil {
			super := &object.Instance{Class: node.Super, Env: object.NewEnvironment(node.Env.GetDirectory())}
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
