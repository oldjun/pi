package evaluator

import (
	"pilang/ast"
	"pilang/object"
)

func evalPropertyExpression(node *ast.PropertyExpression, env *object.Environment) object.Object {
	left := Eval(node.Object, env)
	if isError(left) {
		return left
	}
	switch left.(type) {
	case *object.Hash:
		obj := left.(*object.Hash)
		prop := node.Property.(*ast.Identifier)
		index := &object.String{Value: prop.String()}
		return evalHashIndexExpression(obj, index)
	case *object.Instance:
		obj := left.(*object.Instance)
		prop := node.Property.(*ast.Identifier).String()
		if val, ok := obj.Env.Get(prop); ok {
			return val
		}
		// walk up the chain of super instance looking for it
		super, ok := obj.Env.Get("super")
		if !ok {
			break
		}
		for {
			if val, ok := super.(*object.Instance).Env.Get(prop); ok {
				return val
			}
			super, ok = super.(*object.Instance).Env.Get("super")
			if !ok {
				break
			}
		}
	case *object.Module:
		mod := left.(*object.Module)
		prop := node.Property.(*ast.Identifier).String()
		if val, ok := mod.Properties[prop]; ok {
			return val()
		}
	}
	return object.NewError("invalid property '%s' on type %s", node.Property.String(), left.String())
}

func evalPropertyAssignment(name *ast.PropertyExpression, val object.Object, env *object.Environment) object.Object {
	left := Eval(name.Object, env)
	if isError(left) {
		return left
	}
	switch left.(type) {
	case *object.Hash:
		hash := left.(*object.Hash)
		prop := &object.String{Value: name.Property.(*ast.Identifier).Value}
		hashKey := prop.HashKey()
		hash.Pairs[hashKey] = object.HashPair{Key: prop, Value: val}
	case *object.Instance:
		obj := left.(*object.Instance)
		prop := name.Property.String()
		if _, ok := obj.Env.Get(prop); ok {
			obj.Env.Set(prop, val)
			return NULL
		}
		// walk up the chain of super instance looking for it
		super, ok := obj.Env.Get("super")
		if !ok {
			obj.Env.Set(prop, val)
			return NULL
		}
		for {
			if _, ok := super.(*object.Instance).Env.Get(prop); ok {
				super.(*object.Instance).Env.Set(prop, val)
				return NULL
			}
			super, ok = super.(*object.Instance).Env.Get("super")
			if !ok {
				break
			}
		}
		obj.Env.Set(prop, val)
	default:
		return object.NewError("property assign error: %s", left.Type())
	}
	return NULL
}
