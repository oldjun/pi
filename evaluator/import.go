package evaluator

import (
	"pilang/ast"
	"pilang/module"
	"pilang/object"
)

func evalImport(node *ast.Import, env *object.Environment) object.Object {
	for key, val := range node.Identifiers {
		if mod, ok := module.Map[val.Value]; ok {
			env.Set(key, mod)
		} else {
			return newError("module not exists: %s", val.Value)
		}
	}
	return NULL
}
