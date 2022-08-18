package evaluator

import (
	"pilang/ast"
	"pilang/object"
)

func evalCall(node *ast.Call, env *object.Environment) object.Object {
	obj := Eval(node.Function, env)
	if isError(obj) {
		return obj
	}
	var args []object.Object
	switch obj.(type) {
	case *object.Function:
		fn := obj.(*object.Function)
		// if decorated function is a class method, env should carry `this`
		if this, ok := env.Get("this"); ok {
			obj.(*object.Function).Env.Set("this", this)
		}
		args = evalArgumentExpressions(node, fn, env)
	case *object.Class:
		fn, ok := obj.(*object.Class).Scope.Get("__init__")
		if !ok {
			return object.NewError("class has not __init__ function")
		}
		args = evalArgumentExpressions(node, fn.(*object.Function), env)
	default:
		args = evalExpressions(node.Arguments, env)
	}
	if len(args) == 1 && isError(args[0]) {
		return args[0]
	}
	return applyFunction(obj, args)
}

func evalArgumentExpressions(node *ast.Call, fn *object.Function, env *object.Environment) []object.Object {
	argsList := &object.List{}
	argsHash := &object.Hash{}
	argsHash.Pairs = make(map[object.HashKey]object.HashPair)
	for _, exp := range node.Arguments {
		switch e := exp.(type) {
		case *ast.Assign:
			val := Eval(e.Value, env)
			if isError(val) {
				return []object.Object{val}
			}
			var keyHash object.HashKey
			key := &object.String{Value: e.Name.Value}
			keyHash = key.HashKey()
			pair := object.HashPair{Key: key, Value: val}
			argsHash.Pairs[keyHash] = pair
		default:
			evaluated := Eval(e, env)
			if isError(evaluated) {
				return []object.Object{evaluated}
			}
			argsList.Elements = append(argsList.Elements, evaluated)
		}
	}

	var result []object.Object
	params := make(map[string]bool)
	for _, exp := range fn.Parameters {
		params[exp.Value] = true
		if len(argsList.Elements) > 0 {
			result = append(result, argsList.Elements[0])
			argsList.Elements = argsList.Elements[1:]
		} else {
			keyParam := &object.String{Value: exp.Value}
			keyParamHash := keyParam.HashKey()
			if valParam, ok := argsHash.Pairs[keyParamHash]; ok {
				result = append(result, valParam.Value)
				delete(argsHash.Pairs, keyParamHash)
			} else {
				if _e, _ok := fn.Defaults[exp.Value]; _ok {
					evaluated := Eval(_e, env)
					if isError(evaluated) {
						return []object.Object{evaluated}
					}
					result = append(result, evaluated)
				} else {
					return []object.Object{object.NewError("function %s missing parameter %s", fn.Name, exp.Value)}
				}
			}
		}
	}

	for _, pair := range argsHash.Pairs {
		if _, ok := params[pair.Key.(*object.String).Value]; ok {
			return []object.Object{object.NewError("func got multiple values for argument %s", pair.Key.(*object.String).Value)}
		}
	}

	if fn.Args != nil {
		result = append(result, argsList)
	} else {
		if len(argsList.Elements) > 0 {
			return []object.Object{object.NewError("function args parameters error: %s", fn.Name)}
		}
	}
	if fn.KwArgs != nil {
		result = append(result, argsHash)
	} else {
		if len(argsHash.Pairs) > 0 {
			return []object.Object{object.NewError("function kwargs parameters error: %s", fn.Name)}
		}
	}
	return result
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
			return object.NewError("class %s missing __init__ function", node.Name.Value)
		}
		fn.(*object.Function).Env.Set("this", obj)
		applyFunction(fn, args)
		fn.(*object.Function).Env.Del("this")
		return obj
	default:
		return object.NewError("not a function: %s", node.Type())
	}
}

func extendFunctionEnv(fn *object.Function, args []object.Object) *object.Environment {
	env := object.NewEnclosedEnvironment(fn.Env)
	for idx, param := range fn.Parameters {
		env.Set(param.Value, args[idx])
	}
	if fn.Args != nil && fn.KwArgs != nil {
		env.Set(fn.Args.Value, args[len(args)-2])
		env.Set(fn.KwArgs.Value, args[len(args)-1])
	} else if fn.Args != nil && fn.KwArgs == nil {
		env.Set(fn.Args.Value, args[len(args)-1])
	} else if fn.Args == nil && fn.KwArgs != nil {
		env.Set(fn.KwArgs.Value, args[len(args)-1])
	}
	return env
}

func unwrapReturnValue(obj object.Object) object.Object {
	if ret, ok := obj.(*object.Return); ok {
		return ret.Value
	}
	return obj
}
