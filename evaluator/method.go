package evaluator

import (
	"pilang/ast"
	"pilang/object"
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
	switch obj.(type) {
	case *object.Hash:
		hash := obj.(*object.Hash)
		return hash.Method(method.String(), args)
	case *object.List:
		list := obj.(*object.List)
		return list.Method(method.String(), args)
	case *object.String:
		str := obj.(*object.String)
		return str.Method(method.String(), args)
	case *object.File:
		file := obj.(*object.File)
		return file.Method(method.String(), args)
	case *object.Module:
		mod := obj.(*object.Module)
		if fn, ok := mod.Functions[method.String()]; ok {
			return fn(args)
		}
	case *object.Instance:
		obj := obj.(*object.Instance)
		if fn, ok := obj.Class.Scope.Get(method.String()); ok {
			fn.(*object.Function).Env.Set("this", obj)
			ret := applyFunction(fn, args)
			fn.(*object.Function).Env.Del("this")
			return ret
		}
		// walk up the chain of super instance looking for it
		super := obj.Class.Super
		for super != nil {
			if fn, ok := super.Scope.Get(method.String()); ok {
				fn.(*object.Function).Env.Set("this", obj)
				ret := applyFunction(fn, args)
				fn.(*object.Function).Env.Del("this")
				return ret
			}
			super = super.Super
		}
	}
	return object.NewError("%s does not have method '%s()'", obj.String(), method.String())
}
