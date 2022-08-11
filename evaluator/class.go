package evaluator

import (
	"pilang/ast"
	"pilang/object"
)

func evalClass(node *ast.Class, env *object.Environment) object.Object {
	//class := &object.Class{
	//	Name:  node.Name,
	//	Super: nil,
	//	Env:   env,
	//	Scope: object.NewEnvironment(env.GetDirectory()),
	//}
	class := &object.Class{
		Name:  node.Name,
		Super: nil,
		Env:   env,
		Scope: object.NewEnclosedEnvironment(env),
	}

	// super
	if node.Super != nil {
		identifier, ok := env.Get(node.Super.Value)
		if !ok {
			newError("runtime error: identifier '%s' not found", node.Super.Value)
		}
		super, ok := identifier.(*object.Class)
		if !ok {
			newError("runtime error: referenced identifier in extends not a class, got=%T", super)
		}
		class.Super = super
	}

	// Create a new scope for this class
	Eval(node.Body, class.Scope)
	env.Set(node.Name.Value, class)
	return class
}
