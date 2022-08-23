package evaluator

import (
	"github.com/oldjun/pi/ast"
	"github.com/oldjun/pi/object"
)

func evalBlock(block *ast.Block, env *object.Environment) object.Object {
	var result object.Object
	var deferred []*ast.Defer
	for _, statement := range block.Statements {
		switch stmt := statement.(type) {
		case *ast.Defer:
			deferred = append(deferred, stmt)
			continue
		}
		result = Eval(statement, env)
		if result != nil {
			rt := result.Type()
			if rt == object.RETURN || rt == object.BREAK || rt == object.CONTINUE || rt == object.ERROR {
				return result
			}
		}
	}
	for _, statement := range deferred {
		Eval(statement.Call, env)
	}
	return result
}
