package evaluator

import (
	"pilang/object"
)

var hashBuiltins = map[string]*object.Builtin{
	"keys": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d", len(args))
			}
			if args[0].Type() != object.HASH {
				return newError("argument to `keys` must be HASH, got %s", args[0].Type())
			}
			hash := args[0].(*object.Hash)
			pairs := hash.Pairs
			var keys []object.Object
			for _, pair := range pairs {
				key := pair.Key
				keys = append(keys, key)
			}
			return &object.List{Elements: keys}
		},
	},
	"values": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d", len(args))
			}
			if args[0].Type() != object.HASH {
				return newError("argument to `values` must be HASH, got %s", args[0].Type())
			}
			hash := args[0].(*object.Hash)
			pairs := hash.Pairs
			var values []object.Object
			for _, pair := range pairs {
				value := pair.Value
				values = append(values, value)
			}
			return &object.List{Elements: values}
		},
	},
	"get": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) > 3 {
				return newError("wrong number of arguments. got=%d", len(args))
			}
			if args[0].Type() != object.HASH {
				return newError("argument to `get` must be HASH, got %s", args[0].Type())
			}
			var key object.HashKey
			switch args[1].(type) {
			case *object.String:
				key = args[1].(*object.String).HashKey()
			case *object.Integer:
				key = args[1].(*object.Integer).HashKey()
			case *object.Boolean:
				key = args[1].(*object.Boolean).HashKey()
			default:
				return newError("argument to `get` type error, got %s", args[1].Type())
			}
			hash := args[0].(*object.Hash)
			if pair, ok := hash.Pairs[key]; ok {
				return pair.Value
			}
			switch len(args) {
			case 2:
				return NULL
			case 3:
				return args[2]
			default:
				return NULL
			}
		},
	},
	"delete": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				return newError("wrong number of arguments. got=%d", len(args))
			}
			if args[0].Type() != object.HASH {
				return newError("argument to `get` must be HASH, got %s", args[0].Type())
			}
			var key object.HashKey
			switch args[1].(type) {
			case *object.String:
				key = args[1].(*object.String).HashKey()
			case *object.Integer:
				key = args[1].(*object.Integer).HashKey()
			case *object.Boolean:
				key = args[1].(*object.Boolean).HashKey()
			default:
				return newError("argument to `get` type error, got %s", args[1].Type())
			}
			hash := args[0].(*object.Hash)
			delete(hash.Pairs, key)
			return NULL
		},
	},
}
