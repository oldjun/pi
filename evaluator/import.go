package evaluator

import (
	"github.com/oldjun/pi/ast"
	"github.com/oldjun/pi/module"
	"github.com/oldjun/pi/object"
)

func evalImport(node *ast.Import, env *object.Environment) object.Object {
	for key, val := range node.Identifiers {
		if mod, ok := module.Map[val.Value]; ok {
			env.Set(key, mod)
		} else {
			return object.NewError("module not exists: %s", val.Value)
		}
	}
	return NULL
}
