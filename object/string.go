package object

import "hash/fnv"

type String struct {
	Value string
}

func (s *String) Type() Type     { return STRING }
func (s *String) String() string { return s.Value }
func (s *String) HashKey() HashKey {
	h := fnv.New64a()
	h.Write([]byte(s.Value))
	return HashKey{Type: s.Type(), Value: h.Sum64()}
}
