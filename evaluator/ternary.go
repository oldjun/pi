package evaluator

import (
	"github.com/oldjun/pi/ast"
	"github.com/oldjun/pi/object"
)

func evalTernary(te *ast.Ternary, env *object.Environment) object.Object {
	condition := Eval(te.Condition, env)
	if isError(condition) {
		return condition
	}
	if isTruthy(condition) {
		return Eval(te.IfTrue, env)
	}
	return Eval(te.IfFalse, env)
}
