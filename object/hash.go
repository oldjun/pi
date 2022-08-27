package object

import (
	"bytes"
	"fmt"
	"sort"
	"strings"
)

type HashKey struct {
	Type  Type
	Value uint64
}

type HashPair struct {
	Key   Object
	Value Object
}

type Hash struct {
	Pairs  map[HashKey]HashPair
	offset int
}

func (h *Hash) Type() Type { return HASH }
func (h *Hash) String() string {
	var out bytes.Buffer
	var pairs []string
	for _, pair := range h.Pairs {
		key := ""
		switch k := pair.Key.(type) {
		case *String:
			key = "\"" + k.Value + "\""
		default:
			key = k.String()
		}
		val := ""
		switch v := pair.Value.(type) {
		case *String:
			val = "\"" + v.Value + "\""
		default:
			val = v.String()
		}
		pairs = append(pairs, fmt.Sprintf("%s: %s", key, val))
	}
	out.WriteString("{")
	out.WriteString(strings.Join(pairs, ", "))
	out.WriteString("}")
	return out.String()
}
func (h *Hash) Next() (Object, Object) {
	idx := 0
	hash := make(map[string]HashPair)
	var keys []string
	for _, v := range h.Pairs {
		hash[v.Key.String()] = v
		keys = append(keys, v.Key.String())
	}
	sort.Strings(keys)

	for _, k := range keys {
		if h.offset == idx {
			h.offset += 1
			return hash[k].Key, hash[k].Value
		}
		idx += 1
	}
	return nil, nil
}
func (h *Hash) Reset() {
	h.offset = 0
}

func (h *Hash) Method(method string, args []Object) Object {
	switch method {
	case "len":
		return h.len(args)
	case "keys":
		return h.keys(args)
	case "values":
		return h.values(args)
	case "has":
		return h.has(args)
	case "get":
		return h.get(args)
	case "set":
		return h.set(args)
	case "copy":
		return h.copy(args)
	case "update":
		return h.update(args)
	case "delete":
		return h.delete(args)
	case "clear":
		return h.clear(args)
	case "json":
		return h.json(args)
	}
	return nil
}

func (h *Hash) len(args []Object) Object {
	if len(args) != 0 {
		return NewError("wrong number of arguments. hash.len() got=%d", len(args))
	}
	return &Integer{Value: int64(len(h.Pairs))}
}

func (h *Hash) keys(args []Object) Object {
	if len(args) != 0 {
		return NewError("wrong number of arguments. hash.keys() got=%d", len(args))
	}
	pairs := h.Pairs
	var keys []Object
	for _, pair := range pairs {
		key := pair.Key
		keys = append(keys, key)
	}
	return &List{Elements: keys}
}

func (h *Hash) values(args []Object) Object {
	if len(args) != 0 {
		return NewError("wrong number of arguments. hash.values() got=%d", len(args))
	}
	pairs := h.Pairs
	var values []Object
	for _, pair := range pairs {
		value := pair.Value
		values = append(values, value)
	}
	return &List{Elements: values}
}

func (h *Hash) has(args []Object) Object {
	if len(args) != 1 {
		return NewError("wrong number of arguments. hash.has() got=%d", len(args))
	}
	var key HashKey
	switch arg := args[0].(type) {
	case *String:
		key = arg.HashKey()
	case *Integer:
		key = arg.HashKey()
	case *Boolean:
		key = arg.HashKey()
	default:
		return NewError("argument to hash.has() type error, got %s", arg.Type())
	}
	if _, ok := h.Pairs[key]; ok {
		return &Boolean{Value: true}
	} else {
		return &Boolean{Value: false}
	}
}

func (h *Hash) get(args []Object) Object {
	if len(args) > 2 {
		return NewError("wrong number of arguments. hash.get() got=%d", len(args))
	}
	var key HashKey
	switch arg := args[0].(type) {
	case *String:
		key = arg.HashKey()
	case *Integer:
		key = arg.HashKey()
	case *Boolean:
		key = arg.HashKey()
	default:
		return NewError("argument to hash.get() type error, got %s", arg.Type())
	}
	if pair, ok := h.Pairs[key]; ok {
		return pair.Value
	}
	switch len(args) {
	case 1:
		return &Null{}
	case 2:
		return args[1]
	}
	return nil
}

func (h *Hash) set(args []Object) Object {
	if len(args) != 2 {
		return NewError("wrong number of arguments. hash.set() got=%d", len(args))
	}
	var key HashKey
	switch arg := args[0].(type) {
	case *String:
		key = arg.HashKey()
	case *Integer:
		key = arg.HashKey()
	case *Boolean:
		key = arg.HashKey()
	default:
		return NewError("argument to hash.set() type error, got %s", arg.Type())
	}
	h.Pairs[key] = HashPair{Key: args[0], Value: args[1]}
	return nil
}

func (h *Hash) copy(args []Object) Object {
	if len(args) != 0 {
		return NewError("wrong number of arguments. hash.copy() got=%d", len(args))
	}
	pairs := make(map[HashKey]HashPair)
	for k, v := range h.Pairs {
		pairs[k] = v
	}
	return &Hash{Pairs: pairs}
}

func (h *Hash) update(args []Object) Object {
	if len(args) != 1 {
		return NewError("wrong number of arguments. hash.update() got=%d", len(args))
	}
	switch arg := args[0].(type) {
	case *Hash:
		for k, v := range arg.Pairs {
			h.Pairs[k] = v
		}
	default:
		return NewError("argument to hash.update() type error, got %s", args[0].Type())
	}
	return nil
}

func (h *Hash) delete(args []Object) Object {
	if len(args) != 1 {
		return NewError("wrong number of arguments. hash.delete() got=%d", len(args))
	}
	var key HashKey
	switch arg := args[0].(type) {
	case *String:
		key = arg.HashKey()
	case *Integer:
		key = arg.HashKey()
	case *Boolean:
		key = arg.HashKey()
	default:
		return NewError("argument to hash.delete() type error, got %s", arg.Type())
	}
	delete(h.Pairs, key)
	return nil
}

func (h *Hash) clear(args []Object) Object {
	if len(args) != 0 {
		return NewError("wrong number of arguments. hash.clear() got=%d", len(args))
	}
	h.Pairs = make(map[HashKey]HashPair)
	return nil
}

func (h *Hash) json(args []Object) Object {
	if len(args) != 0 {
		return NewError("wrong number of arguments. hash.json() got=%d", len(args))
	}
	return &String{Value: h.String()}
}
