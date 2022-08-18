package evaluator

import (
	"pilang/ast"
	"pilang/object"
)

func evalIndexExpression(node *ast.IndexExpression, env *object.Environment) object.Object {
	left := Eval(node.Left, env)
	if isError(left) {
		return left
	}
	index := Eval(node.Index, env)
	if isError(index) {
		return index
	}
	end := Eval(node.End, env)
	if isError(end) {
		return end
	}
	switch {
	case left.Type() == object.LIST && index.Type() == object.INTEGER:
		return evalListIndexExpression(left, index, end, node.IsRange)
	case left.Type() == object.HASH && index.Type() == object.STRING:
		return evalHashIndexExpression(left, index)
	case left.Type() == object.STRING && index.Type() == object.INTEGER:
		return evalStringIndexExpression(left, index, end, node.IsRange)
	default:
		return object.NewError("index operator not supported: %s on %s", index.String(), left.Type())
	}
}

func evalListIndexExpression(left, index, end object.Object, isRange bool) object.Object {
	leftObject := left.(*object.List)
	idx := index.(*object.Integer).Value
	max := int64(len(leftObject.Elements) - 1)
	if idx < -(max+1) || idx > max {
		return NULL
	}
	if isRange {
		if idx < 0 {
			idx += max + 1
		}
		if end == NULL {
			endIdx := max + 1
			if idx >= endIdx {
				return &object.List{Elements: []object.Object{}}
			}
			return &object.List{Elements: leftObject.Elements[idx:endIdx]}
		} else {
			endObj, ok := end.(*object.Integer)
			if ok {
				endIdx := endObj.Value
				if endIdx < -(max+1) || endIdx > max {
					return object.NewError(`index out of range: got %s`, end.String())
				}
				if endIdx < 0 {
					endIdx += max + 1
				}
				if idx >= endIdx {
					return &object.List{Elements: []object.Object{}}
				}
				return &object.List{Elements: leftObject.Elements[idx:endIdx]}
			} else {
				return object.NewError(`index range can only be numerical: got %s (type %s)`, end.String(), end.Type())
			}
		}
	} else {
		if idx < 0 {
			idx += max + 1
		}
		return leftObject.Elements[idx]
	}
}

func evalHashIndexExpression(hash, index object.Object) object.Object {
	hashObject := hash.(*object.Hash)
	key, ok := index.(object.Hashable)
	if !ok {
		return object.NewError("unusable as hash key: %s", index.Type())
	}
	pair, ok := hashObject.Pairs[key.HashKey()]
	if !ok {
		return NULL
	}
	return pair.Value
}

func evalStringIndexExpression(left, index, end object.Object, isRange bool) object.Object {
	leftObject := left.(*object.String)
	idx := index.(*object.Integer).Value
	max := int64(len(leftObject.Value) - 1)
	if idx < -(max+1) || idx > max {
		return NULL
	}
	if isRange {
		if idx < 0 {
			idx += max + 1
		}
		if end == NULL {
			endIdx := max + 1
			if idx >= endIdx {
				return &object.String{Value: ""}
			}
			return &object.String{Value: leftObject.Value[idx:endIdx]}
		} else {
			endObj, ok := end.(*object.Integer)
			if ok {
				endIdx := endObj.Value
				if endIdx < -(max+1) || endIdx > max {
					return object.NewError(`index out of range: got %s`, end.String())
				}
				if endIdx < 0 {
					endIdx += max + 1
				}
				if idx >= endIdx {
					return &object.String{Value: ""}
				}
				return &object.String{Value: leftObject.Value[idx:endIdx]}
			} else {
				return object.NewError(`index range can only be numerical: got %s (type %s)`, end.String(), end.Type())
			}
		}
	} else {
		if idx < 0 {
			idx += max + 1
		}
		return &object.String{Value: string(leftObject.Value[idx])}
	}
}

func evalIndexAssignment(name *ast.IndexExpression, val object.Object, env *object.Environment) object.Object {
	left := Eval(name.Left, env)
	if isError(left) {
		return left
	}
	index := Eval(name.Index, env)
	if isError(index) {
		return index
	}
	switch left.(type) {
	case *object.List:
		array := left.(*object.List)
		idx := int(index.(*object.Integer).Value)
		if idx < 0 {
			return object.NewError("index out of range: %d", idx)
		}
		if idx >= len(array.Elements) {
			// expand the array by appending null objects
			for i := len(array.Elements); i <= idx; i++ {
				array.Elements = append(array.Elements, NULL)
			}
		}
		array.Elements[idx] = val
	case *object.Hash:
		hash := left.(*object.Hash)
		key, ok := index.(object.Hashable)
		if !ok {
			return object.NewError("unusable as hash key: %s", index.Type())
		}
		hashKey := key.HashKey()
		hash.Pairs[hashKey] = object.HashPair{Key: index, Value: val}
	default:
		return object.NewError("index assignment error: %s", left.Type())
	}
	return NULL
}
