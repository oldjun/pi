package evaluator

import (
	"pilang/ast"
	"pilang/object"
)

// Property expression (x.y) evaluator.
//
// Here we have a special case, as strings
// have an .ok property when they're the result
// of a command.
//
// Else we will try to parse the property
// as an index of hash.
//
// If that doesn't work, we'll spectacularly
// give up.
func evalPropertyExpression(node *ast.PropertyExpression, env *object.Environment) object.Object {
	left := Eval(node.Object, env)
	if isError(left) {
		return left
	}
	switch left.(type) {
	case *object.Hash:
		hash := left.(*object.Hash)
		prop := node.Property.(*ast.Identifier)
		index := &object.String{Value: prop.String()}
		return evalHashIndexExpression(hash, index)
	case *object.Instance:
		obj := left.(*object.Instance)
		prop := "__prop__" + node.Property.(*ast.Identifier).String()
		if val, ok := obj.Env.Get(prop); ok {
			return val
		}
		// walk up the chain of super instance looking for it
		super, ok := obj.Env.Get("super")
		if !ok {
			return NULL
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
	}
	return newError("invalid property '%s' on type %s", node.Property.String(), left.Type())
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
		prop := "__prop__" + name.Property.String()
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
		return newError("property assignment error: %s", left.Type())
	}
	return NULL
}
