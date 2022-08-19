package evaluator

import (
	"github.com/oldjun/pi/ast"
	"github.com/oldjun/pi/object"
)

func evalWhile(we *ast.While, env *object.Environment) object.Object {
	condition := Eval(we.Condition, env)
	if isError(condition) {
		return condition
	}
	if isTruthy(condition) {
		evaluated := Eval(we.Consequence, env)
		if isError(evaluated) {
			return evaluated
		}
		if evaluated.Type() == object.BREAK {
			return evaluated
		}
		evalWhile(we, env)
	}
	return NULL
}
