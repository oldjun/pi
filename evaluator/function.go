package evaluator

import (
	"github.com/oldjun/pi/ast"
	"github.com/oldjun/pi/object"
)

func evalFunction(node *ast.Function, env *object.Environment) object.Object {
	function := &object.Function{
		Name:       node.Name,
		Parameters: node.Parameters,
		Defaults:   node.Defaults,
		Args:       node.Args,
		KwArgs:     node.KwArgs,
		Body:       node.Body,
		Env:        env,
	}
	if node.Name != "" {
		env.Set(node.Name, function)
	}
	return function
}
