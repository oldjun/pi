package object

import "fmt"

type Integer struct {
	Value int64
}

func (i *Integer) Type() Type { return INTEGER }
func (i *Integer) String() string {
	return fmt.Sprintf("%d", i.Value)
}
func (i *Integer) HashKey() HashKey {
	return HashKey{Type: i.Type(), Value: uint64(i.Value)}
}
