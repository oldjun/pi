package object

type Scope struct {
	Env  *Environment
	Self Object
}

func (s *Scope) Type() Type { return SCOPE }
func (s *Scope) String() string {
	return "scope"
}
