package evaluator

import (
	"pilang/ast"
	"pilang/object"
)

// for x = 0; x < 10; x++ {}
func evalFor(fe *ast.For, env *object.Environment) object.Object {
	// Let's figure out if the for loop is using a variable that's
	// already been declared. If so, let's keep it aside for now.
	obj, ok := env.Get(fe.Identifier)

	// Final cleanup: we remove the x from the environment. If
	// it was already declared before the for loop, we restore
	// it to its original value
	defer func() {
		if ok {
			env.Set(fe.Identifier, obj)
		}
	}()

	// Eval the starter (x = 0)
	err := Eval(fe.Starter, env)
	if isError(err) {
		return err
	}

	// When for is while...
	for true {
		// Evaluate the for condition
		evaluated := Eval(fe.Condition, env)
		if isError(evaluated) {
			return evaluated
		}
		if !isTruthy(evaluated) {
			break
		}
		// If truthy, execute the block and the closer
		res := Eval(fe.Block, env)
		if isError(res) {
			return res
		}
		if res.Type() == object.BREAK {
			break
		}
		if res.Type() == object.CONTINUE {
			err = Eval(fe.Closer, env)
			if isError(err) {
				return err
			}
			continue
		}
		if res.Type() == object.RETURN {
			return res
		}
		err = Eval(fe.Closer, env)
		if isError(err) {
			return err
		}
	}
	return NULL
}

// for k,v in [1,2,3] {}
func evalForIn(fie *ast.ForIn, env *object.Environment) object.Object {
	iterable := Eval(fie.Iterable, env)
	// If "k" and "v" were already declared, let's keep them aside...
	existingKeyIdentifier, okk := env.Get(fie.Key)
	existingValueIdentifier, okv := env.Get(fie.Value)
	// ...so that we can restore them after the for loop is over
	defer func() {
		if okk {
			env.Set(fie.Key, existingKeyIdentifier)
		}
		if okv {
			env.Set(fie.Value, existingValueIdentifier)
		}
	}()

	switch i := iterable.(type) {
	case object.Iterable:
		defer func() {
			i.Reset()
		}()
		return loopIterable(i.Next, env, fie)
	default:
		return object.NewError("%s is a %s, not an iterable, cannot be used in for loop", i.String(), i.Type())
	}
}

// This function iterates over an iterable
// represented by the next() function: everytime
// we call it, a new kv pair is popped from the iterable
func loopIterable(next func() (object.Object, object.Object), env *object.Environment, fi *ast.ForIn) object.Object {
	// Let's get the first kv pair out
	k, v := next()

	// Let's keep going until there are no more kv pairs
	for k != nil && v != nil {
		// set the special k v variables in the
		// environment
		env.Set(fi.Key, k)
		env.Set(fi.Value, v)
		res := Eval(fi.Block, env)
		if isError(res) {
			return res
		}
		if res.Type() == object.BREAK {
			break
		}
		if res.Type() == object.CONTINUE {
			k, v = next()
			continue
		}
		if res.Type() == object.RETURN {
			return res
		}

		// Let's pull the next kv pair
		k, v = next()
	}
	return NULL
}
