package evaluator

import (
	"github.com/oldjun/pi/ast"
	"github.com/oldjun/pi/object"
)

func evalCompound(node *ast.Compound, env *object.Environment) object.Object {
	left := Eval(node.Left, env)
	if isError(left) {
		return left
	}
	right := Eval(node.Right, env)
	if isError(right) {
		return right
	}
	// multi-character operators like "+=" and "*=" are reduced to "+" or "*" for evalInfixExpression()
	op := node.Operator
	if len(op) >= 2 {
		op = op[:len(op)-1]
	}
	// get the result of the infix operation
	val := evalInfix(op, left, right)
	if isError(val) {
		return val
	}
	switch nodeLeft := node.Left.(type) {
	case *ast.Identifier:
		env.Set(nodeLeft.Value, val)
		return NULL
	case *ast.IndexExpression:
		return evalIndexAssignment(nodeLeft, val, env)
	case *ast.PropertyExpression:
		return evalPropertyAssignment(nodeLeft, val, env)
	}
	return NULL
}
