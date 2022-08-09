package object

type Builtin struct {
	Fn BuiltinFunction
}

func (b *Builtin) Type() Type     { return BUILTIN }
func (b *Builtin) String() string { return "builtin function" }
