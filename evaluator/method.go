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
		// it might have user-defined functions
		// If so, run the user-defined function
		key := &object.String{Value: method.String()}
		pair, ok := hash.Pairs[key.HashKey()]
		if ok && pair.Value.Type() == object.FUNCTION {
			return pair.Value
		}
		if builtin, ok := hashBuiltins[method.String()]; ok {
			args = append([]object.Object{hash}, args...)
			return builtin.Fn(args...)
		}
	case *object.Array:
		array := obj.(*object.Array)
		if builtin, ok := arrayBuiltins[method.String()]; ok {
			args = append([]object.Object{array}, args...)
			return builtin.Fn(args...)
		}
	case *object.String:
		str := obj.(*object.String)
		if builtin, ok := stringBuiltins[method.String()]; ok {
			args = append([]object.Object{str}, args...)
			return builtin.Fn(args...)
		}
	case *object.File:
		file := obj.(*object.File)
		if builtin, ok := fileBuiltins[method.String()]; ok {
			args = append([]object.Object{file}, args...)
			return builtin.Fn(args...)
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
	default:
		break
	}
	return newError("%s does not have method '%s()'", obj.String(), method.String())
}
