package evaluator

import (
	"pilang/object"
	"strings"
)

var arrayBuiltins = map[string]*object.Builtin{
	"len": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d", len(args))
			}
			if args[0].Type() != object.ARRAY {
				return newError("argument to `push` must be ARRAY, got %s", args[0].Type())
			}
			switch arg := args[0].(type) {
			case *object.List:
				return &object.Integer{Value: int64(len(arg.Elements))}
			default:
				return newError("argument to `len` not supported, got %s", args[0].Type())
			}
		},
	},
	"push": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				return newError("wrong number of arguments. got=%d", len(args))
			}
			if args[0].Type() != object.ARRAY {
				return newError("argument to `push` must be ARRAY, got %s", args[0].Type())
			}
			array := args[0].(*object.List)
			array.Elements = append(array.Elements, args[1])
			return NULL
		},
	},
	"pop": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d", len(args))
			}
			if args[0].Type() != object.ARRAY {
				return newError("argument to `push` must be ARRAY, got %s", args[0].Type())
			}
			array := args[0].(*object.List)
			if len(array.Elements) == 0 {
				return NULL
			}
			elem := array.Elements[len(array.Elements)-1]
			array.Elements = array.Elements[0 : len(array.Elements)-1]
			return elem
		},
	},
	"shift": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d", len(args))
			}
			if args[0].Type() != object.ARRAY {
				return newError("argument to `push` must be ARRAY, got %s", args[0].Type())
			}
			array := args[0].(*object.List)
			if len(array.Elements) == 0 {
				return NULL
			}
			elem := array.Elements[0]
			array.Elements = array.Elements[1:]
			return elem
		},
	},
	"join": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d", len(args))
			}
			if args[0].Type() != object.ARRAY {
				return newError("argument to `push` must be ARRAY, got %s", args[0].Type())
			}
			array := args[0].(*object.List)
			if len(array.Elements) > 0 {
				glue := ""
				if len(args) == 2 {
					glue = args[1].(*object.String).Value
				}
				length := len(array.Elements)
				newElements := make([]string, length, length)
				for k, v := range array.Elements {
					newElements[k] = v.String()
				}
				return &object.String{Value: strings.Join(newElements, glue)}
			} else {
				return &object.String{Value: ""}
			}
		},
	},
}
