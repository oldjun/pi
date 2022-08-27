package object

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type BuiltinFunction func(args []Object) Object

type Builtin struct {
	Name string
	Fn   BuiltinFunction
}

func (b *Builtin) Type() Type { return BUILTIN }
func (b *Builtin) String() string {
	return fmt.Sprintf("<builtin:%s>", b.Name)
}

var Builtins = map[string]*Builtin{}

func init() {
	Builtins["len"] = &Builtin{Name: "len", Fn: lenFunction}
	Builtins["int"] = &Builtin{Name: "int", Fn: intFunction}
	Builtins["float"] = &Builtin{Name: "float", Fn: floatFunction}
	Builtins["str"] = &Builtin{Name: "str", Fn: strFunction}
	Builtins["type"] = &Builtin{Name: "type", Fn: typeFunction}
	Builtins["open"] = &Builtin{Name: "open", Fn: openFunction}
	Builtins["exit"] = &Builtin{Name: "exit", Fn: exitFunction}
	Builtins["bytes"] = &Builtin{Name: "bytes", Fn: bytesFunction}
	Builtins["print"] = &Builtin{Name: "print", Fn: printFunction}
	Builtins["printf"] = &Builtin{Name: "printf", Fn: printfFunction}
	Builtins["sprintf"] = &Builtin{Name: "sprintf", Fn: sprintfFunction}
}

func lenFunction(args []Object) Object {
	if len(args) != 1 {
		return NewError("wrong number of arguments. len() got=%d", len(args))
	}
	switch arg := args[0].(type) {
	case *String:
		return &Integer{Value: int64(len(arg.Value))}
	case *Bytes:
		return &Integer{Value: int64(len(arg.Value))}
	case *List:
		return &Integer{Value: int64(len(arg.Elements))}
	case *Hash:
		return &Integer{Value: int64(len(arg.Pairs))}
	}
	return NewError("argument to `len` not supported, got %s", args[0].Type())
}

func intFunction(args []Object) Object {
	if len(args) != 1 {
		return NewError("wrong number of arguments. int() got=%d", len(args))
	}
	switch arg := args[0].(type) {
	case *Integer:
		return &Integer{Value: arg.Value}
	case *Float:
		return &Integer{Value: int64(arg.Value)}
	case *String:
		val, ok := strconv.ParseInt(arg.Value, 10, 64)
		if ok != nil {
			return NewError("argument to `int` error, got %s", arg.String())
		}
		return &Integer{Value: val}
	}
	return NewError("argument to `int` not supported, got %s", args[0].Type())
}

func floatFunction(args []Object) Object {
	if len(args) != 1 {
		return NewError("wrong number of arguments. float() got=%d", len(args))
	}
	switch arg := args[0].(type) {
	case *Integer:
		return &Float{Value: float64(arg.Value)}
	case *Float:
		return &Float{Value: arg.Value}
	case *String:
		val, ok := strconv.ParseFloat(arg.Value, 64)
		if ok != nil {
			return NewError("argument to `float` error, got %s", arg.String())
		}
		return &Float{Value: val}
	}
	return NewError("argument to `float` not supported, got %s", args[0].Type())
}

func strFunction(args []Object) Object {
	if len(args) != 1 {
		return NewError("wrong number of arguments. str() got=%d", len(args))
	}
	switch arg := args[0].(type) {
	case *Integer:
		return &String{Value: arg.String()}
	case *Float:
		return &String{Value: arg.String()}
	case *String:
		return &String{Value: arg.Value}
	case *Bytes:
		return &String{Value: arg.String()}
	}
	return NewError("argument to `str()` not supported, got %s", args[0].Type())
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

func bytesFunction(args []Object) Object {
	if len(args) == 0 {
		return &Bytes{}
	}
	if len(args) != 1 {
		return NewError("wrong number of arguments. bytes() got=%d", len(args))
	}
	switch arg := args[0].(type) {
	case *String:
		bytes := &Bytes{}
		bytes.Value = append(bytes.Value, []byte(arg.Value)...)
		return bytes
	}
	return NewError("wrong type of arguments. bytes(): %s", args[0].Type())
}
