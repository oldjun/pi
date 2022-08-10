package evaluator

import (
	"io"
	"pilang/object"
	"strings"
)

var fileBuiltins = map[string]*object.Builtin{
	"read": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d", len(args))
			}
			if args[0].Type() != object.FILE {
				return newError("argument to `read` must be FILE, got %s", args[0].Type())
			}
			f := args[0].(*object.File)
			if f.Reader == nil {
				return NULL
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
			text := strings.Join(lines, "\n")
			return &object.String{Value: text}
		},
	},
	"readline": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d", len(args))
			}
			if args[0].Type() != object.FILE {
				return newError("argument to `read` must be FILE, got %s", args[0].Type())
			}
			f := args[0].(*object.File)
			if f.Reader == nil {
				return NULL
			}
			line, err := f.Reader.ReadString('\n')
			line = strings.Trim(line, "\r\n")
			if err != nil {
				return NULL
			}
			return &object.String{Value: line}
		},
	},
	"readlines": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d", len(args))
			}
			if args[0].Type() != object.FILE {
				return newError("argument to `read` must be FILE, got %s", args[0].Type())
			}
			f := args[0].(*object.File)
			if f.Reader == nil {
				return NULL
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
			result := make([]object.Object, length)
			for i, line := range lines {
				result[i] = &object.String{Value: line}
			}
			return &object.List{Elements: result}
		},
	},
	"rewind": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d", len(args))
			}
			if args[0].Type() != object.FILE {
				return newError("argument to `rewind` must be FILE, got %s", args[0].Type())
			}
			f := args[0].(*object.File)
			_, err := f.Handle.Seek(0, 0)
			if err != nil {
				return FALSE
			}
			return TRUE
		},
	},
	"write": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				return newError("wrong number of arguments. got=%d", len(args))
			}
			if args[0].Type() != object.FILE {
				return newError("argument to `write` must be FILE, got %s", args[0].Type())
			}
			f := args[0].(*object.File)
			if f.Writer == nil {
				return FALSE
			}
			txt := args[1].String()
			_, err := f.Writer.Write([]byte(txt))
			if err != nil {
				return FALSE
			}
			err = f.Writer.Flush()
			if err != nil {
				return FALSE
			}
			return TRUE
		},
	},
	"close": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d", len(args))
			}
			if args[0].Type() != object.FILE {
				return newError("argument to `close` must be FILE, got %s", args[0].Type())
			}
			f := args[0].(*object.File)
			err := f.Handle.Close()
			if err != nil {
				return FALSE
			}
			return TRUE
		},
	},
}
