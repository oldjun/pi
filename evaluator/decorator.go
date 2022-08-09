package evaluator

import (
	"pilang/ast"
)

//func evalDecorator(node *ast.Decorator, env *object.Environment) object.Object {
//	var decorator object.Object
//	evaluated := Eval(node.Expression, env)
//	switch evaluated.(type) {
//	case *object.Function:
//		decorator = evaluated
//	default:
//		return newError("decorator '%s' is not a function", evaluated.Inspect())
//	}
//	name, ok := getDecoratedName(node.Decorated)
//	if !ok {
//		return newError("error while processing decorator: unable to find the name of the function you're trying to decorate")
//	}
//	switch decorated := node.Decorated.(type) {
//	case *ast.FunctionLiteral:
//		// Here we have a single decorator
//		fn := &object.Function{Name: name, Parameters: decorated.Parameters, Env: env, Body: decorated.Body, Decorated: decorated}
//		return applyFunction(decorator, []object.Object{fn})
//	case *ast.Decorator:
//		// Here we have a decorator of another decorator
//		// decoratorObj, _ := env.Get(node.Name)
//		// decorator := decoratorObj.(*object.Function)
//
//		// First eval the later decorator(s).
//		fnName, fn, err := doEvalDecorator(decorated, env)
//		if isError(err) {
//			return "", nil, err
//		}
//		return fnName, applyFunction(decorator, append([]object.Object{fn})), nil
//	default:
//		return "", nil, newError("a decorator must decorate a named function or another decorator")
//	}
//	// ----------------------------------
//	ident, fn, err := doEvalDecorator(node, env)
//	if isError(err) {
//		return err
//	}
//	env.Set(ident, fn)
//	return NULL
//}

// the parser should simply convert:
// @deco()
// f test() {}
// into
// f test() {}
// test = deco(test)
//func doEvalDecorator(node *ast.Decorator, env *object.Environment) (string, object.Object, object.Object) {
//	var decorator object.Object
//
//	evaluated := Eval(node.Expression, env)
//	switch evaluated.(type) {
//	case *object.Function:
//		decorator = evaluated
//	default:
//		return "", nil, newError("decorator '%s' is not a function", evaluated.Inspect())
//	}
//
//	name, ok := getDecoratedName(node.Decorated)
//	if !ok {
//		return "", nil, newError("error while processing decorator: unable to find the name of the function you're trying to decorate")
//	}
//	switch decorated := node.Decorated.(type) {
//	case *ast.FunctionLiteral:
//		// Here we have a single decorator
//		fn := &object.Function{Parameters: decorated.Parameters, Env: env, Body: decorated.Body, Name: name, Node: decorated}
//		return name, applyFunction(decorator, []object.Object{fn}), nil
//	case *ast.Decorator:
//		// Here we have a decorator of another decorator
//		// decoratorObj, _ := env.Get(node.Name)
//		// decorator := decoratorObj.(*object.Function)
//
//		// First eval the later decorator(s).
//		fnName, fn, err := doEvalDecorator(decorated, env)
//		if isError(err) {
//			return "", nil, err
//		}
//		return fnName, applyFunction(decorator, append([]object.Object{fn})), nil
//	default:
//		return "", nil, newError("a decorator must decorate a named function or another decorator")
//	}
//}

// Finds the actual name of the decorated function.
//
// Given this:
// @deco1()
// @deco2()
// f hello() {}
//
// After we evaluate the decorators, we need to
// re-assign the original function "hello", like this:
//
// hello = deco1(deco2(hello))
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
