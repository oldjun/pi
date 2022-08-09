package object

import "strconv"

type Float struct {
	Value float64
}

func (f *Float) Type() Type { return FLOAT }
func (f *Float) String() string {
	return strconv.FormatFloat(f.Value, 'f', -1, 64)
}
