package object

import (
	"bytes"
	"strings"
)

type List struct {
	Elements []Object
	offset   int
}

func (a *List) Type() Type { return ARRAY }
func (a *List) String() string {
	var out bytes.Buffer
	var elements []string
	for _, e := range a.Elements {
		elements = append(elements, e.String())
	}
	out.WriteString("[")
	out.WriteString(strings.Join(elements, ", "))
	out.WriteString("]")
	return out.String()
}
func (a *List) Next() (Object, Object) {
	offset := a.offset
	if len(a.Elements) > offset {
		a.offset = offset + 1
		return &Integer{Value: int64(offset)}, a.Elements[offset]
	}
	return nil, nil
}
func (a *List) Reset() {
	a.offset = 0
}
