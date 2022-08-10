package evaluator

import (
	"fmt"
	"pilang/object"
	"strings"
)

var stringBuiltins = map[string]*object.Builtin{
	"len": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d", len(args))
			}
			switch arg := args[0].(type) {
			case *object.String:
				return &object.Integer{Value: int64(len(arg.Value))}
			default:
				return newError("argument to `len` not supported, got %s", args[0].Type())
			}
		},
	},
	"upper": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d", len(args))
			}
			switch arg := args[0].(type) {
			case *object.String:
				return &object.String{Value: strings.ToUpper(arg.Value)}
			default:
				return newError("argument to `upper` not supported, got %s", args[0].Type())
			}
		},
	},
	"lower": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d", len(args))
			}
			switch arg := args[0].(type) {
			case *object.String:
				return &object.String{Value: strings.ToLower(arg.Value)}
			default:
				return newError("argument to `lower` not supported, got %s", args[0].Type())
			}
		},
	},
	"title": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d", len(args))
			}
			switch arg := args[0].(type) {
			case *object.String:
				return &object.String{Value: strings.Title(arg.Value)}
			default:
				return newError("argument to `title` not supported, got %s", args[0].Type())
			}
		},
	},
	"split": {
		Fn: func(args ...object.Object) object.Object {
			sep := " "
			if len(args) > 1 {
				sep = args[1].(*object.String).Value
			}
			switch arg := args[0].(type) {
			case *object.String:
				parts := strings.Split(arg.Value, sep)
				length := len(parts)
				elements := make([]object.Object, length, length)
				for k, v := range parts {
					elements[k] = &object.String{Value: v}
				}
				return &object.List{Elements: elements}
			default:
				return newError("argument to `split` not supported, got %s", args[0].Type())
			}
		},
	},
	"replace": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 3 {
				return newError("wrong number of arguments. got=%d", len(args))
			}
			switch arg := args[0].(type) {
			case *object.String:
				oldStr := args[1].(*object.String).Value
				newStr := args[2].(*object.String).Value
				return &object.String{Value: strings.Replace(arg.Value, oldStr, newStr, -1)}
			default:
				return newError("argument to `replace` not supported, got %s", args[0].Type())
			}
		},
	},
	"contain": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				return newError("wrong number of arguments. got=%d", len(args))
			}
			switch arg := args[0].(type) {
			case *object.String:
				substr := args[1].(*object.String).Value
				return &object.Boolean{Value: strings.Contains(arg.Value, substr)}
			default:
				return newError("argument to `contain` not supported, got %s", args[0].Type())
			}
		},
	},
	"prefix": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				return newError("wrong number of arguments. got=%d", len(args))
			}
			switch arg := args[0].(type) {
			case *object.String:
				substr := args[1].(*object.String).Value
				return &object.Boolean{Value: strings.HasPrefix(arg.Value, substr)}
			default:
				return newError("argument to `prefix` not supported, got %s", args[0].Type())
			}
		},
	},
	"suffix": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				return newError("wrong number of arguments. got=%d", len(args))
			}
			switch arg := args[0].(type) {
			case *object.String:
				substr := args[1].(*object.String).Value
				return &object.Boolean{Value: strings.HasSuffix(arg.Value, substr)}
			default:
				return newError("argument to `suffix` not supported, got %s", args[0].Type())
			}
		},
	},
	"repeat": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				return newError("wrong number of arguments. got=%d", len(args))
			}
			switch arg := args[0].(type) {
			case *object.String:
				count := args[1].(*object.Integer).Value
				return &object.String{Value: strings.Repeat(arg.Value, int(count))}
			default:
				return newError("argument to `repeat` not supported, got %s", args[0].Type())
			}
		},
	},
	"trim": {
		Fn: func(args ...object.Object) object.Object {
			switch arg := args[0].(type) {
			case *object.String:
				cut := " "
				if len(args) == 2 {
					cut = args[1].(*object.String).Value
				}
				return &object.String{Value: strings.Trim(arg.Value, cut)}
			default:
				return newError("argument to `trim` not supported, got %s", args[0].Type())
			}
		},
	},
	"index": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				return newError("wrong number of arguments. got=%d", len(args))
			}
			switch arg := args[0].(type) {
			case *object.String:
				substr := args[1].(*object.String).Value
				idx := strings.Index(arg.Value, substr)
				if idx < 0 {
					return NULL
				}
				return &object.Integer{Value: int64(idx)}
			default:
				return newError("argument to `index` not supported, got %s", args[0].Type())
			}
		},
	},
	"format": {
		Fn: func(args ...object.Object) object.Object {
			switch arg := args[0].(type) {
			case *object.String:
				var list []interface{}
				for _, s := range args[1:] {
					list = append(list, s.String())
				}
				return &object.String{Value: fmt.Sprintf(arg.Value, list...)}
			default:
				return newError("argument to `format` not supported, got %s", args[0].Type())
			}
		},
	},
}
