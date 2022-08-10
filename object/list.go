package object

import (
	"bytes"
	"strings"
)

type List struct {
	Elements []Object
	offset   int
}

func (l *List) Type() Type { return ARRAY }
func (l *List) String() string {
	var out bytes.Buffer
	var elements []string
	for _, e := range l.Elements {
		elements = append(elements, e.String())
	}
	out.WriteString("[")
	out.WriteString(strings.Join(elements, ", "))
	out.WriteString("]")
	return out.String()
}
func (l *List) Next() (Object, Object) {
	offset := l.offset
	if len(l.Elements) > offset {
		l.offset = offset + 1
		return &Integer{Value: int64(offset)}, l.Elements[offset]
	}
	return nil, nil
}
func (l *List) Reset() {
	l.offset = 0
}

func (l *List) Method(method string, args []Object) Object {
	switch method {
	case "len":
		return l.len(args)
	case "push":
		return l.push(args)
	case "pop":
		return l.pop(args)
	case "shift":
		return l.shift(args)
	case "join":
		return l.join(args)
	default:
		return nil
	}
}

func (l *List) len(args []Object) Object {
	if len(args) != 0 {
		return newError("wrong number of arguments. list.len() got=%d", len(args))
	}
	return &Integer{Value: int64(len(l.Elements))}
}

func (l *List) push(args []Object) Object {
	if len(args) != 1 {
		return newError("wrong number of arguments. list.push() got=%d", len(args))
	}
	l.Elements = append(l.Elements, args[0])
	return nil
}

func (l *List) pop(args []Object) Object {
	if len(args) != 0 {
		return newError("wrong number of arguments. list.pop() got=%d", len(args))
	}
	if len(l.Elements) == 0 {
		return nil
	}
	elem := l.Elements[len(l.Elements)-1]
	l.Elements = l.Elements[0 : len(l.Elements)-1]
	return elem
}

func (l *List) shift(args []Object) Object {
	if len(args) != 0 {
		return newError("wrong number of arguments. list.shift() got=%d", len(args))
	}
	if len(l.Elements) == 0 {
		return nil
	}
	elem := l.Elements[0]
	l.Elements = l.Elements[1:]
	return elem
}

func (l *List) join(args []Object) Object {
	if len(args) > 1 {
		return newError("wrong number of arguments. list.join() got=%d", len(args))
	}
	if len(l.Elements) > 0 {
		glue := ""
		if len(args) == 1 {
			glue = args[0].(*String).Value
		}
		length := len(l.Elements)
		newElements := make([]string, length, length)
		for k, v := range l.Elements {
			newElements[k] = v.String()
		}
		return &String{Value: strings.Join(newElements, glue)}
	} else {
		return &String{Value: ""}
	}
}
