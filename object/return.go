package object

type Return struct {
	Value Object
}

func (r *Return) Type() Type     { return RETURN }
func (r *Return) String() string { return r.Value.String() }
