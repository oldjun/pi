package evaluator

import (
	"pilang/ast"
	"pilang/object"
)

func evalBlock(block *ast.Block, env *object.Environment) object.Object {
	var result object.Object
	for _, statement := range block.Statements {
		result = Eval(statement, env)
		if result != nil {
			rt := result.Type()
			if rt == object.RETURN || rt == object.BREAK || rt == object.CONTINUE || rt == object.ERROR {
				return result
			}
		}
	}
	return result
}
