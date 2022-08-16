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
	case *object.Math:
		obj := left.(*object.Math)
		prop := node.Property.(*ast.Identifier).String()
		val := obj.Property(prop)
		if val != nil {
			return val
		}
	}
	return newError("invalid property '%s' on type %s", node.Property.String(), left.String())
}

func evalPropertyAssignment(name *ast.PropertyExpression, val object.Object, env *object.Environment) object.Object {
	left := Eval(name.Object, env)
	if isError(left) {
		return left
	}
	switch left.(type) {
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
		return newError("property assign error: %s", left.Type())
	}
	return NULL
}
