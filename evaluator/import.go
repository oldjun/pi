package evaluator

import (
	"pilang/ast"
	"pilang/object"
)

func evalImport(node *ast.Import, env *object.Environment) object.Object {
	for key, val := range node.Identifiers {
		if module, ok := object.ModuleMap[val.Value]; ok {
			env.Set(key, module.Value)
		} else {
			return newError("module not exists: %s", val.Value)
		}
	}
	return NULL
}
