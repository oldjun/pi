package evaluator

import (
	"github.com/oldjun/pi/ast"
	"github.com/oldjun/pi/object"
)

func evalIf(ie *ast.If, env *object.Environment) object.Object {
	for _, scenario := range ie.Scenarios {
		condition := Eval(scenario.Condition, env)
		if isError(condition) {
			return condition
		}
		if isTruthy(condition) {
			return Eval(scenario.Consequence, env)
		}
	}
	return NULL
}
