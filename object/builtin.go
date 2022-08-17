package object

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

type Builtin struct {
	Fn BuiltinFunction
}

func (b *Builtin) Type() Type     { return BUILTIN }
func (b *Builtin) String() string { return "builtin function" }

func (b *Builtin) Method(method string, args []Object) Object {
	switch method {
	case "len":
		return b.len(args)
	case "type":
		return b.typeof(args)
	case "sleep":
		return b.sleep(args)
	case "time":
		return b.time(args)
	case "print":
		return b.print(args)
	case "printf":
		return b.printf(args)
	case "sprintf":
		return b.sprintf(args)
	case "open":
		return b.open(args)
	case "exit":
		return b.exit(args)
	}
	return nil
}

func (b *Builtin) len(args []Object) Object {
	if len(args) != 1 {
		return newError("wrong number of arguments. len() got=%d", len(args))
	}
	switch arg := args[0].(type) {
	case *String:
		return &Integer{Value: int64(len(arg.Value))}
	case *List:
		return &Integer{Value: int64(len(arg.Elements))}
	}
	return newError("argument to `len` not supported, got %s", args[0].Type())
}

func (b *Builtin) typeof(args []Object) Object {
	if len(args) != 1 {
		return newError("wrong number of arguments. type() got=%d", len(args))
	}
	return &String{Value: string(args[0].Type())}
}

func (b *Builtin) sleep(args []Object) Object {
	if len(args) != 1 {
		return NewError("wrong number of arguments. sleep() got=%d", len(args))
	}
	switch arg := args[0].(type) {
	case *Integer:
		time.Sleep(time.Duration(arg.Value) * time.Millisecond)
		return &Null{}
	}
	return newError("argument to `sleep` not supported, got %s", args[0].Type())
}

func (b *Builtin) time(args []Object) Object {
	if len(args) != 0 {
		return newError("wrong number of arguments. time() got=%d", len(args))
	}
	return &Integer{Value: time.Now().UnixNano() / 1000000}
}

func (b *Builtin) print(args []Object) Object {
	var arr []string
	for _, arg := range args {
		arr = append(arr, arg.String())
	}
	str := strings.Join(arr, " ")
	print(str + "\n")
	return &Null{}
}

func (b *Builtin) printf(args []Object) Object {
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
			return newError("error occurred while calling 'printf', parameter type not support: %s", arg.String())
		}
	}
	str := fmt.Sprintf(format, a...)
	print(str + "\n")
	return &Null{}
}

func (b *Builtin) sprintf(args []Object) Object {
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
			return newError("error occurred while calling 'sprintf', parameter type not support: %s", arg.String())
		}
	}
	str := fmt.Sprintf(format, a...)
	return &String{Value: str}
}

func (b *Builtin) open(args []Object) Object {
	if len(args) > 2 {
		return newError("wrong number of arguments. open() got=%d", len(args))
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
			return newError("file mode error. got=%s", fileMode)
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

func (b *Builtin) exit(args []Object) Object {
	if len(args) != 2 {
		return newError("wrong number of arguments. exit() got=%d", len(args))
	}
	switch arg := args[0].(type) {
	case *Integer:
		os.Exit(int(arg.Value))
	}
	return NewError("wrong type of arguments. exit() got=%s", args[0].Type())
}
