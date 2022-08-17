package evaluator

import (
	"pilang/ast"
	"pilang/object"
)

func evalDecorator(node *ast.Decorator, env *object.Environment) object.Object {
	identifier, fn, err := doEvalDecorator(node, env)
	if isError(err) {
		return err
	}
	env.Set(identifier, fn)
	return NULL
}

// the parser should simply convert:
//
// @decor()
// func test() {}
//
// into
//
// func test() {}
// test = decor(test)
func doEvalDecorator(node *ast.Decorator, env *object.Environment) (string, object.Object, object.Object) {
	var decorator object.Object
	evaluated := Eval(node.Expression, env)
	switch evaluated.(type) {
	case *object.Function:
		decorator = evaluated
	default:
		return "", nil, object.NewError("decorator '%s' is not a function", evaluated.String())
	}
	name, ok := getDecoratedName(node.Decorated)
	if !ok {
		return "", nil, object.NewError("error while processing decorator: unable to find the name of the function you're trying to decorate")
	}
	switch decorated := node.Decorated.(type) {
	case *ast.Function:
		// Here we have a single decorator
		fn := &object.Function{Name: name, Parameters: decorated.Parameters, Env: env, Body: decorated.Body}
		return name, applyFunction(decorator, []object.Object{fn}), nil
	case *ast.Decorator:
		// First eval the later decorator(s).
		fnName, fn, err := doEvalDecorator(decorated, env)
		if isError(err) {
			return "", nil, err
		}
		return fnName, applyFunction(decorator, append([]object.Object{fn})), nil
	default:
		return "", nil, object.NewError("a decorator must decorate a named function or another decorator")
	}
}

// Finds the actual name of the decorated function.
//
// Given this:
// @decor1()
// @decor2()
// f hello() {}
//
// After we evaluate the decorators, we need to
// re-assign the original function "hello", like this:
//
// hello = decor1(decor2(hello))
//
// This function traverses the decorators and finds the
// name of the function we have to re-assign.
func getDecoratedName(decorated ast.Expression) (string, bool) {
	switch d := decorated.(type) {
	case *ast.Function:
		return d.Name, true
	case *ast.Decorator:
		return getDecoratedName(d.Decorated)
	}
	return "", false
}
