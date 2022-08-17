package object

import (
	"fmt"
	"hash/fnv"
	"strings"
)

type String struct {
	Value  string
	offset int
}

func (s *String) Type() Type     { return STRING }
func (s *String) String() string { return s.Value }
func (s *String) HashKey() HashKey {
	h := fnv.New64a()
	h.Write([]byte(s.Value))
	return HashKey{Type: s.Type(), Value: h.Sum64()}
}
func (s *String) Next() (Object, Object) {
	offset := s.offset
	if len(s.Value) > offset {
		s.offset = offset + 1
		return &Integer{Value: int64(offset)}, &String{Value: string(s.Value[offset])}
	}
	return nil, nil
}
func (s *String) Reset() {
	s.offset = 0
}

func (s *String) Method(method string, args []Object) Object {
	switch method {
	case "len":
		return s.len(args)
	case "upper":
		return s.upper(args)
	case "lower":
		return s.lower(args)
	case "title":
		return s.title(args)
	case "split":
		return s.split(args)
	case "replace":
		return s.replace(args)
	case "contain":
		return s.contain(args)
	case "prefix":
		return s.prefix(args)
	case "suffix":
		return s.suffix(args)
	case "repeat":
		return s.repeat(args)
	case "trim":
		return s.trim(args)
	case "index":
		return s.index(args)
	case "format":
		return s.format(args)
	}
	return nil
}

func (s *String) len(args []Object) Object {
	if len(args) != 0 {
		return NewError("wrong number of arguments. string.len() got=%d", len(args))
	}
	return &Integer{Value: int64(len(s.Value))}
}

func (s *String) upper(args []Object) Object {
	if len(args) != 0 {
		return NewError("wrong number of arguments. string.upper() got=%d", len(args))
	}
	return &String{Value: strings.ToUpper(s.Value)}
}

func (s *String) lower(args []Object) Object {
	if len(args) != 0 {
		return NewError("wrong number of arguments. string.lower got=%d", len(args))
	}
	return &String{Value: strings.ToLower(s.Value)}
}

func (s *String) title(args []Object) Object {
	if len(args) != 0 {
		return NewError("wrong number of arguments. string.title() got=%d", len(args))
	}
	return &String{Value: strings.Title(s.Value)}
}

func (s *String) split(args []Object) Object {
	if len(args) > 1 {
		return NewError("wrong number of arguments. string.split() got=%d", len(args))
	}
	sep := " "
	if len(args) == 1 {
		sep = args[0].(*String).Value
	}
	parts := strings.Split(s.Value, sep)
	length := len(parts)
	elements := make([]Object, length, length)
	for k, v := range parts {
		elements[k] = &String{Value: v}
	}
	return &List{Elements: elements}
}

func (s *String) replace(args []Object) Object {
	if len(args) != 2 {
		return NewError("wrong number of arguments. string.replace() got=%d", len(args))
	}
	oldStr := args[0].(*String).Value
	newStr := args[1].(*String).Value
	return &String{Value: strings.Replace(s.Value, oldStr, newStr, -1)}
}

func (s *String) contain(args []Object) Object {
	if len(args) != 1 {
		return NewError("wrong number of arguments. string.contain() got=%d", len(args))
	}
	substr := args[0].(*String).Value
	return &Boolean{Value: strings.Contains(s.Value, substr)}
}

func (s *String) prefix(args []Object) Object {
	if len(args) != 1 {
		return NewError("wrong number of arguments. string.prefix() got=%d", len(args))
	}
	substr := args[0].(*String).Value
	return &Boolean{Value: strings.HasPrefix(s.Value, substr)}
}

func (s *String) suffix(args []Object) Object {
	if len(args) != 1 {
		return NewError("wrong number of arguments. string.suffix() got=%d", len(args))
	}
	substr := args[0].(*String).Value
	return &Boolean{Value: strings.HasSuffix(s.Value, substr)}
}

func (s *String) repeat(args []Object) Object {
	if len(args) != 1 {
		return NewError("wrong number of arguments. string.repeat() got=%d", len(args))
	}
	count := args[0].(*Integer).Value
	return &String{Value: strings.Repeat(s.Value, int(count))}
}

func (s *String) trim(args []Object) Object {
	if len(args) > 1 {
		return NewError("wrong number of arguments. string.trim() got=%d", len(args))
	}
	cut := " "
	if len(args) == 1 {
		cut = args[0].(*String).Value
	}
	return &String{Value: strings.Trim(s.Value, cut)}
}

func (s *String) index(args []Object) Object {
	if len(args) != 1 {
		return NewError("wrong number of arguments. string.index() got=%d", len(args))
	}
	substr := args[0].(*String).Value
	idx := strings.Index(s.Value, substr)
	return &Integer{Value: int64(idx)}
}

func (s *String) format(args []Object) Object {
	var list []interface{}
	for _, arg := range args {
		switch arg.(type) {
		case *String:
			list = append(list, arg.(*String).Value)
		case *Integer:
			list = append(list, arg.(*Integer).Value)
		case *Float:
			list = append(list, arg.(*Float).Value)
		default:
			return NewError("error occurred while calling string.format(), parameter type not support: %s", arg.String())
		}
	}
	return &String{Value: fmt.Sprintf(s.Value, list...)}
}
