package object

type Null struct{}

func (n *Null) Type() Type     { return NULL }
func (n *Null) String() string { return "null" }
