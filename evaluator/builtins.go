package evaluator

import (
	"bufio"
	"fmt"
	"os"
	"pilang/object"
	"strings"
	"time"
)

var builtins = map[string]*object.Builtin{
	"len": {
		Fn: lenFunction,
	},
	"type": {
		Fn: typeFunction,
	},
	"exit": {
		Fn: exitFunction,
	},
	"sleep": {
		Fn: sleepFunction,
	},
	"time": {
		Fn: timeFunction,
	},
	"print": {
		Fn: printFunction,
	},
	"printf": {
		Fn: printfFunction,
	},
	"sprintf": {
		Fn: sprintfFunction,
	},
	"open": {
		Fn: openFunction,
	},
}

func lenFunction(args []object.Object) object.Object {
	if len(args) != 1 {
		return object.NewError("wrong number of arguments. got=%d", len(args))
	}
	switch arg := args[0].(type) {
	case *object.String:
		return &object.Integer{Value: int64(len(arg.Value))}
	case *object.List:
		return &object.Integer{Value: int64(len(arg.Elements))}
	case *object.Hash:
		return &object.Integer{Value: int64(len(arg.Pairs))}
	default:
		return object.NewError("argument to `len` not supported, got %s", args[0].Type())
	}
}

func typeFunction(args []object.Object) object.Object {
	if len(args) != 1 {
		return object.NewError("wrong number of arguments. got=%d", len(args))
	}
	return &object.String{Value: string(args[0].Type())}
}

func exitFunction(args []object.Object) object.Object {
	if len(args) != 1 {
		return object.NewError("wrong number of arguments. exit() got=%d", len(args))
	}
	if args[0].Type() != object.INTEGER {
		return object.NewError("argument to `exit` must be INTEGER, got=%s", args[0].Type())
	}
	os.Exit(int(args[0].(*object.Integer).Value))
	return NULL
}

func sleepFunction(args []object.Object) object.Object {
	if len(args) != 1 {
		return object.NewError("wrong number of arguments. got=%d", len(args))
	}
	switch arg := args[0].(type) {
	case *object.Integer:
		time.Sleep(time.Duration(arg.Value) * time.Millisecond)
		return NULL
	default:
		return object.NewError("argument to `sleep` not supported, got %s", args[0].Type())
	}
}

func timeFunction(args []object.Object) object.Object {
	if len(args) != 0 {
		return object.NewError("wrong number of arguments. got=%d", len(args))
	}
	return &object.Integer{Value: time.Now().UnixNano() / 1000000}
}

func printFunction(args []object.Object) object.Object {
	var arr []string
	for _, arg := range args {
		arr = append(arr, arg.String())
	}
	str := strings.Join(arr, " ")
	print(str + "\n")
	return NULL
}

func printfFunction(args []object.Object) object.Object {
	format := args[0].(*object.String).Value
	var a []interface{}
	for _, arg := range args[1:] {
		switch arg.(type) {
		case *object.String:
			a = append(a, arg.(*object.String).Value)
		case *object.Integer:
			a = append(a, arg.(*object.Integer).Value)
		case *object.Float:
			a = append(a, arg.(*object.Float).Value)
		case *object.Boolean:
			a = append(a, arg.(*object.Boolean).Value)
		default:
			a = append(a, arg.String())
		}
	}
	str := fmt.Sprintf(format, a...)
	print(str + "\n")
	return NULL
}

func sprintfFunction(args []object.Object) object.Object {
	format := args[0].(*object.String).Value
	var a []interface{}
	for _, arg := range args[1:] {
		switch arg.(type) {
		case *object.String:
			a = append(a, arg.(*object.String).Value)
		case *object.Integer:
			a = append(a, arg.(*object.Integer).Value)
		case *object.Float:
			a = append(a, arg.(*object.Float).Value)
		case *object.Boolean:
			a = append(a, arg.(*object.Boolean).Value)
		default:
			a = append(a, arg.String())
		}
	}
	str := fmt.Sprintf(format, a...)
	return &object.String{Value: str}
}

func openFunction(args []object.Object) object.Object {
	if len(args) > 2 {
		return object.NewError("wrong number of arguments. got=%d", len(args))
	}
	filename := args[0].(*object.String).Value
	mode := os.O_RDONLY
	if len(args) == 2 {
		fileMode := args[1].(*object.String).Value
		switch fileMode {
		case "r":
			mode = os.O_RDONLY
		case "w":
			mode = os.O_WRONLY
			err := os.Remove(filename)
			if err != nil {
				return NULL
			}
		case "a":
			mode = os.O_APPEND
		default:
			return object.NewError("file mode error. got=%s", fileMode)
		}
	}
	file, err := os.OpenFile(filename, os.O_CREATE|mode, 0644)
	if err != nil {
		return NULL
	}
	var reader *bufio.Reader
	var writer *bufio.Writer
	if mode == os.O_RDONLY {
		reader = bufio.NewReader(file)
	} else {
		writer = bufio.NewWriter(file)
	}
	return &object.File{Filename: filename, Reader: reader, Writer: writer, Handle: file}
}
