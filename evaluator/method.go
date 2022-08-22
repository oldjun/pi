package evaluator

import (
	"github.com/oldjun/pi/ast"
	"github.com/oldjun/pi/object"
)

func evalMethod(node *ast.Method, env *object.Environment) object.Object {
	obj := Eval(node.Object, env)
	if isError(obj) {
		return obj
	}
	args := evalExpressions(node.Arguments, env)
	if len(args) == 1 && isError(args[0]) {
		return args[0]
	}
	return applyMethod(obj, node.Method, args)
}

func applyMethod(obj object.Object, method ast.Expression, args []object.Object) object.Object {
	switch obj := obj.(type) {
	case *object.Hash:
		return obj.Method(method.(*ast.Identifier).Value, args)
	case *object.List:
		return obj.Method(method.(*ast.Identifier).Value, args)
	case *object.String:
		return obj.Method(method.(*ast.Identifier).Value, args)
	case *object.Bytes:
		return obj.Method(method.(*ast.Identifier).Value, args)
	case *object.File:
		return obj.Method(method.(*ast.Identifier).Value, args)
	case *object.Module:
		if fn, ok := obj.Functions[method.(*ast.Identifier).Value]; ok {
			return fn(args)
		}
		return obj.Handler.Method(method.(*ast.Identifier).Value, args)
	case *object.Instance:
		if fn, ok := obj.Class.Scope.Get(method.(*ast.Identifier).Value); ok {
			fn.(*object.Function).Env.Set("this", obj)
			ret := applyFunction(fn, args)
			fn.(*object.Function).Env.Del("this")
			return ret
		}
		// walk up the chain of super instance looking for it
		super := obj.Class.Super
		for super != nil {
			if fn, ok := super.Scope.Get(method.(*ast.Identifier).Value); ok {
				fn.(*object.Function).Env.Set("this", obj)
				ret := applyFunction(fn, args)
				fn.(*object.Function).Env.Del("this")
				return ret
			}
			super = super.Super
		}
	}
	return object.NewError("%s does not have method '%s()'", obj.String(), method.(*ast.Identifier).Value)
}
