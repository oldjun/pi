package object

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
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
	case "rewind":
		return f.rewind(args)
	case "write":
		return f.write(args)
	case "close":
		return f.close(args)
	}
	return nil
}

func (f *File) read(args []Object) Object {
	if len(args) != 0 {
		return newError("wrong number of arguments. file.read() got=%d", len(args))
	}
	if f.Reader == nil {
		return nil
	}
	text, _ := io.ReadAll(f.Reader)
	return &String{Value: string(text)}
}

func (f *File) readline(args []Object) Object {
	if len(args) != 0 {
		return newError("wrong number of arguments. file.readline got=%d", len(args))
	}
	if f.Reader == nil {
		return nil
	}
	line, err := f.Reader.ReadString('\n')
	line = strings.Trim(line, "\r\n")
	if err == io.EOF {
		return &Null{}
	}
	return &String{Value: line}
}

func (f *File) lines(args []Object) Object {
	if len(args) != 0 {
		return newError("wrong number of arguments. file.lines() got=%d", len(args))
	}
	if f.Reader == nil {
		return nil
	}
	var lines []string
	for {
		line, err := f.Reader.ReadString('\n')
		line = strings.Trim(line, "\r\n")
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

func (f *File) rewind(args []Object) Object {
	if len(args) != 0 {
		return newError("wrong number of arguments. file.rewind() got=%d", len(args))
	}
	_, err := f.Handle.Seek(0, 0)
	return &Boolean{Value: err == nil}
}

func (f *File) write(args []Object) Object {
	if len(args) != 1 {
		return newError("wrong number of arguments. file.write got=%d", len(args))
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
		return newError("wrong number of arguments. file.close() got=%d", len(args))
	}
	_ = f.Handle.Close()
	return &Null{}
}
