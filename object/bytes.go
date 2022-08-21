package object

type Bytes struct {
	Value  []byte
	offset int
}

func (b *Bytes) Type() Type { return BYTES }
func (b *Bytes) String() string {
	return string(b.Value)
}

func (b *Bytes) Method(method string, args []Object) Object {
	switch method {
	case "str":
		return b.str(args)
	}
	return nil
}

func (b *Bytes) str(args []Object) Object {
	if len(args) != 0 {
		return NewError("wrong number of arguments. bytes.str() got=%d", len(args))
	}
	return &String{Value: b.String()}
}
