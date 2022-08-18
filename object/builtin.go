package object

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type BuiltinFunction func(args []Object) Object

type Builtin struct {
	Fn BuiltinFunction
}

func (b *Builtin) Type() Type     { return BUILTIN }
func (b *Builtin) String() string { return "builtin function" }

var Builtins = map[string]*Builtin{}

func init() {
	Builtins["len"] = &Builtin{Fn: lenFunction}
	Builtins["type"] = &Builtin{Fn: typeFunction}
	Builtins["open"] = &Builtin{Fn: openFunction}
	Builtins["exit"] = &Builtin{Fn: exitFunction}
	Builtins["print"] = &Builtin{Fn: printFunction}
	Builtins["printf"] = &Builtin{Fn: printfFunction}
	Builtins["sprintf"] = &Builtin{Fn: sprintfFunction}
}

func lenFunction(args []Object) Object {
	if len(args) != 1 {
		return NewError("wrong number of arguments. len() got=%d", len(args))
	}
	switch arg := args[0].(type) {
	case *String:
		return &Integer{Value: int64(len(arg.Value))}
	case *List:
		return &Integer{Value: int64(len(arg.Elements))}
	}
	return NewError("argument to `len` not supported, got %s", args[0].Type())
}

func typeFunction(args []Object) Object {
	if len(args) != 1 {
		return NewError("wrong number of arguments. type() got=%d", len(args))
	}
	return &String{Value: string(args[0].Type())}
}

func printFunction(args []Object) Object {
	var arr []string
	for _, arg := range args {
		arr = append(arr, arg.String())
	}
	str := strings.Join(arr, " ")
	print(str + "\n")
	return &Null{}
}

func printfFunction(args []Object) Object {
	format := args[0].(*String).Value
	var a []interface{}
	for _, arg := range args[1:] {
		switch arg.(type) {
		case *String:
			a = append(a, arg.(*String).Value)
		case *Integer:
			a = append(a, arg.(*Integer).Value)
		case *Float:
			a = append(a, arg.(*Float).Value)
		default:
			return NewError("error occurred while calling 'printf', parameter type not support: %s", arg.String())
		}
	}
	str := fmt.Sprintf(format, a...)
	print(str + "\n")
	return &Null{}
}

func sprintfFunction(args []Object) Object {
	format := args[0].(*String).Value
	var a []interface{}
	for _, arg := range args[1:] {
		switch arg.(type) {
		case *String:
			a = append(a, arg.(*String).Value)
		case *Integer:
			a = append(a, arg.(*Integer).Value)
		case *Float:
			a = append(a, arg.(*Float).Value)
		default:
			return NewError("error occurred while calling 'sprintf', parameter type not support: %s", arg.String())
		}
	}
	str := fmt.Sprintf(format, a...)
	return &String{Value: str}
}

func openFunction(args []Object) Object {
	if len(args) > 2 {
		return NewError("wrong number of arguments. open() got=%d", len(args))
	}
	filename := args[0].(*String).Value
	mode := os.O_RDONLY
	if len(args) == 2 {
		fileMode := args[1].(*String).Value
		switch fileMode {
		case "r":
			mode = os.O_RDONLY
		case "w":
			mode = os.O_WRONLY
			err := os.Remove(filename)
			if err != nil {
				return &Null{}
			}
		case "a":
			mode = os.O_APPEND
		default:
			return NewError("file mode error. got=%s", fileMode)
		}
	}
	file, err := os.OpenFile(filename, os.O_CREATE|mode, 0644)
	if err != nil {
		return &Null{}
	}
	var reader *bufio.Reader
	var writer *bufio.Writer
	if mode == os.O_RDONLY {
		reader = bufio.NewReader(file)
	} else {
		writer = bufio.NewWriter(file)
	}
	return &File{Filename: filename, Reader: reader, Writer: writer, Handle: file}
}

func exitFunction(args []Object) Object {
	if len(args) == 0 {
		os.Exit(0)
	}
	if len(args) != 1 {
		return NewError("wrong number of arguments. exit() got=%d", len(args))
	}
	switch arg := args[0].(type) {
	case *Integer:
		os.Exit(int(arg.Value))
	}
	return NewError("wrong type of arguments. exit() got=%s", args[0].Type())
}
