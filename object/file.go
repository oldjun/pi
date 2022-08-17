package object

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type File struct {
	Filename string
	Reader   *bufio.Reader
	Writer   *bufio.Writer
	Handle   *os.File
}

func (f *File) Type() Type { return FILE }
func (f *File) String() string {
	return fmt.Sprintf("<file:%s>", f.Filename)
}

func (f *File) Method(method string, args []Object) Object {
	switch method {
	case "read":
		return f.read(args)
	case "readline":
		return f.readline(args)
	case "lines":
		return f.lines(args)
	case "seek":
		return f.seek(args)
	case "write":
		return f.write(args)
	case "close":
		return f.close(args)
	}
	return nil
}

func (f *File) read(args []Object) Object {
	if len(args) > 1 {
		return NewError("wrong number of arguments. file.read() got=%d", len(args))
	}
	if f.Reader == nil {
		return nil
	}
	if len(args) == 0 {
		txt, _ := io.ReadAll(f.Reader)
		return &String{Value: string(txt)}
	}
	switch arg := args[0].(type) {
	case *Integer:
		size := arg.Value
		buf := make([]byte, size)
		_, err := io.ReadFull(f.Reader, buf)
		if err != nil {
			return nil
		}
		return &String{Value: string(buf)}
	default:
		return NewError("wrong type of arguments. file.read() got=%s", arg.Type())
	}
}

func (f *File) readline(args []Object) Object {
	if len(args) != 0 {
		return NewError("wrong number of arguments. file.readline got=%d", len(args))
	}
	if f.Reader == nil {
		return nil
	}
	line, _ := f.Reader.ReadString('\n')
	return &String{Value: line}
}

func (f *File) lines(args []Object) Object {
	if len(args) != 0 {
		return NewError("wrong number of arguments. file.lines() got=%d", len(args))
	}
	if f.Reader == nil {
		return nil
	}
	var lines []string
	for {
		line, err := f.Reader.ReadString('\n')
		lines = append(lines, line)
		if err == io.EOF {
			break
		}
	}
	length := len(lines)
	result := make([]Object, length)
	for i, line := range lines {
		result[i] = &String{Value: line}
	}
	return &List{Elements: result}
}

func (f *File) seek(args []Object) Object {
	if len(args) != 2 {
		return NewError("wrong number of arguments. file.seek() got=%d", len(args))
	}
	offset := args[0].(*Integer).Value
	whence := args[1].(*Integer).Value
	_, err := f.Handle.Seek(offset, int(whence))
	return &Boolean{Value: err == nil}
}

func (f *File) write(args []Object) Object {
	if len(args) != 1 {
		return NewError("wrong number of arguments. file.write() got=%d", len(args))
	}
	if f.Writer == nil {
		return nil
	}
	text := args[0].String()
	count, err := f.Writer.Write([]byte(text))
	if err != nil {
		return nil
	}
	err = f.Writer.Flush()
	return &Integer{Value: int64(count)}
}

func (f *File) close(args []Object) Object {
	if len(args) != 0 {
		return NewError("wrong number of arguments. file.close() got=%d", len(args))
	}
	_ = f.Handle.Close()
	return &Null{}
}
