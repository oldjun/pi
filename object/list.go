package object

import (
	"bytes"
	"strings"
)

type List struct {
	Elements []Object
	offset   int
}

func (l *List) Type() Type { return LIST }
func (l *List) String() string {
	var out bytes.Buffer
	var elements []string
	for _, e := range l.Elements {
		switch e := e.(type) {
		case *String:
			elements = append(elements, "\""+e.String()+"\"")
		default:
			elements = append(elements, e.String())
		}
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
	case "append":
		return l.append(args)
	case "pop":
		return l.pop(args)
	case "shift":
		return l.shift(args)
	case "insert":
		return l.insert(args)
	case "remove":
		return l.remove(args)
	case "index":
		return l.index(args)
	case "extend":
		return l.extend(args)
	case "join":
		return l.join(args)
	case "json":
		return l.json(args)
	}
	return nil
}

func (l *List) len(args []Object) Object {
	if len(args) != 0 {
		return NewError("wrong number of arguments. list.len() got=%d", len(args))
	}
	return &Integer{Value: int64(len(l.Elements))}
}

func (l *List) append(args []Object) Object {
	if len(args) != 1 {
		return NewError("wrong number of arguments. list.push() got=%d", len(args))
	}
	l.Elements = append(l.Elements, args[0])
	return nil
}

func (l *List) pop(args []Object) Object {
	if len(args) != 0 {
		return NewError("wrong number of arguments. list.pop() got=%d", len(args))
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
		return NewError("wrong number of arguments. list.shift() got=%d", len(args))
	}
	if len(l.Elements) == 0 {
		return nil
	}
	elem := l.Elements[0]
	l.Elements = l.Elements[1:]
	return elem
}

func (l *List) insert(args []Object) Object {
	if len(args) != 2 {
		return NewError("wrong number of arguments. list.insert() got=%d", len(args))
	}
	switch arg := args[0].(type) {
	case *Integer:
		idx := int(arg.Value)
		if idx > len(l.Elements) {
			return NewError("out of range. list.insert() got=%d", idx)
		}
		var elements []Object
		for _, elem := range l.Elements[0:idx] {
			elements = append(elements, elem)
		}
		elements = append(elements, args[1])
		for _, elem := range l.Elements[idx:] {
			elements = append(elements, elem)
		}
		l.Elements = elements
		return nil
	default:
		return NewError("wrong type of arguments. list.extend() got=%s", arg.Type())
	}
}

func (l *List) remove(args []Object) Object {
	if len(args) != 1 {
		return NewError("wrong number of arguments. list.remove() got=%d", len(args))
	}
	switch arg := args[0].(type) {
	case *Integer:
		idx := int(arg.Value)
		if idx >= len(l.Elements) {
			return NewError("out of range. list.remove() got=%d", idx)
		}
		var elements []Object
		for _, elem := range l.Elements[0:idx] {
			elements = append(elements, elem)
		}
		for _, elem := range l.Elements[idx+1:] {
			elements = append(elements, elem)
		}
		l.Elements = elements
		return nil
	default:
		return NewError("wrong type of arguments. list.extend() got=%s", arg.Type())
	}
}

func (l *List) index(args []Object) Object {
	if len(args) != 1 {
		return NewError("wrong number of arguments. list.index() got=%d", len(args))
	}
	for idx, elem := range l.Elements {
		if (elem.Type() == args[0].Type()) && (elem.String() == args[0].String()) {
			return &Integer{Value: int64(idx)}
		}
	}
	return &Integer{Value: -1}
}

func (l *List) extend(args []Object) Object {
	if len(args) != 1 {
		return NewError("wrong number of arguments. list.extend() got=%d", len(args))
	}
	switch arg := args[0].(type) {
	case *List:
		for _, elem := range arg.Elements {
			l.Elements = append(l.Elements, elem)
		}
		return nil
	default:
		return NewError("wrong type of arguments. list.extend() got=%s", arg.Type())
	}
}

func (l *List) join(args []Object) Object {
	if len(args) > 1 {
		return NewError("wrong number of arguments. list.join() got=%d", len(args))
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

func (l *List) json(args []Object) Object {
	if len(args) != 0 {
		return NewError("wrong number of arguments. list.json() got=%d", len(args))
	}
	return &String{Value: l.String()}
}
