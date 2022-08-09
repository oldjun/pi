package object

type Break struct{}

func (b *Break) Type() Type     { return BREAK }
func (b *Break) String() string { return "break" }
