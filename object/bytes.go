package object

import "bytes"

type Bytes struct {
	Value []byte
}

func (b *Bytes) Type() Type { return BYTES }
func (b *Bytes) String() string {
	idx := bytes.IndexByte(b.Value, 0)
	if idx < 0 {
		return string(b.Value)
	} else {
		return string(b.Value[0:idx])
	}
}

func (b *Bytes) Method(method string, args []Object) Object {
	switch method {
	case "str":
		return b.str(args)
	case "append":
		return b.append(args)
	}
	return nil
}

func (b *Bytes) str(args []Object) Object {
	if len(args) != 0 {
		return NewError("wrong number of arguments. bytes.str() got=%d", len(args))
	}
	return &String{Value: b.String()}
}

func (b *Bytes) append(args []Object) Object {
	if len(args) != 1 {
		return NewError("wrong number of arguments. bytes.append() got=%d", len(args))
	}
	switch arg := args[0].(type) {
	case *Integer:
		b.Value = append(b.Value, byte(arg.Value))
		return &Null{}
	case *String:
		b.Value = append(b.Value, []byte(arg.Value)...)
		return &Null{}
	case *Bytes:
		idx := bytes.IndexByte(arg.Value, 0)
		if idx < 0 {
			b.Value = append(b.Value, arg.Value...)
		} else {
			b.Value = append(b.Value, arg.Value[0:idx]...)
		}
		return &Null{}
	}
	return NewError("wrong type of arguments. bytes.append(): ", args[0].Type())
}
