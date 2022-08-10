package evaluator

import (
	"pilang/ast"
	"pilang/object"
)

func evalSwitch(node *ast.Switch, env *object.Environment) object.Object {
	obj := Eval(node.Value, env)
	for _, opt := range node.Cases {
		// skipping the default-case, which we'll handle later.
		if opt.Default {
			continue
		}
		for _, val := range opt.Values {
			out := Eval(val, env)
			if (obj.Type() == out.Type()) && (obj.String() == out.String()) {
				return evalBlock(opt.Body, env)
			}
		}
	}
	// no match ? Handle default if present
	for _, opt := range node.Cases {
		if opt.Default {
			return evalBlock(opt.Body, env)
		}
	}
	return nil
}
