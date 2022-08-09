package object

import "fmt"

type Boolean struct {
	Value bool
}

func (b *Boolean) Type() Type { return BOOLEAN }
func (b *Boolean) String() string {
	return fmt.Sprintf("%t", b.Value)
}
func (b *Boolean) HashKey() HashKey {
	var value uint64
	if b.Value {
		value = 1
	} else {
		value = 0
	}
	return HashKey{Type: b.Type(), Value: value}
}
