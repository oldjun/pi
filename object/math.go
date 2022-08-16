package object

import (
	"fmt"
	"math"
	"math/rand"
)

type Math struct{}

func (m *Math) Type() Type { return MATH }
func (m *Math) String() string {
	return fmt.Sprintf("<math>")
}

func (m *Math) Method(method string, args []Object) Object {
	switch method {
	case "abs":
		return m.abs(args)
	case "min":
		return m.min(args)
	case "max":
		return m.max(args)
	case "random":
		return m.random(args)
	}
	return nil
}

func (m *Math) abs(args []Object) Object {
	if len(args) != 1 {
		return newError("wrong number of arguments. math.abs() got=%d", len(args))
	}
	switch arg := args[0].(type) {
	case *Integer:
		return &Integer{Value: int64(math.Abs(float64(arg.Value)))}
	case *Float:
		return &Float{Value: math.Abs(arg.Value)}
	default:
		return newError("wrong type of arguments. math.abs() got=%s", arg.Type())
	}
}

func (m *Math) min(args []Object) Object {
	if len(args) != 2 {
		return newError("wrong number of arguments. math.min() got=%d", len(args))
	}
	l := args[0]
	r := args[1]
	if (l.Type() == INTEGER) && (r.Type() == INTEGER) {
		num := math.Min(float64(l.(*Integer).Value), float64(r.(*Integer).Value))
		return &Float{Value: num}
	} else if (l.Type() == FLOAT) && (r.Type() == FLOAT) {
		num := math.Min(l.(*Float).Value, r.(*Float).Value)
		return &Float{Value: num}
	} else if (l.Type() == INTEGER) && (r.Type() == FLOAT) {
		num := math.Min(float64(l.(*Integer).Value), r.(*Float).Value)
		return &Float{Value: num}
	} else if (l.Type() == FLOAT) && (r.Type() == INTEGER) {
		num := math.Min(l.(*Float).Value, float64(r.(*Integer).Value))
		return &Float{Value: num}
	}
	if l.Type() != INTEGER || l.Type() != FLOAT {
		return newError("wrong type of arguments. math.min() got=%s", l.Type())
	}
	if r.Type() != INTEGER || r.Type() != FLOAT {
		return newError("wrong type of arguments. math.min() got=%s", r.Type())
	}
	return newError("wrong type of arguments. math.min()")
}

func (m *Math) max(args []Object) Object {
	if len(args) != 2 {
		return newError("wrong number of arguments. math.max() got=%d", len(args))
	}
	l := args[0]
	r := args[1]
	if (l.Type() == INTEGER) && (r.Type() == INTEGER) {
		num := math.Max(float64(l.(*Integer).Value), float64(r.(*Integer).Value))
		return &Float{Value: num}
	} else if (l.Type() == FLOAT) && (r.Type() == FLOAT) {
		num := math.Max(l.(*Float).Value, r.(*Float).Value)
		return &Float{Value: num}
	} else if (l.Type() == INTEGER) && (r.Type() == FLOAT) {
		num := math.Max(float64(l.(*Integer).Value), r.(*Float).Value)
		return &Float{Value: num}
	} else if (l.Type() == FLOAT) && (r.Type() == INTEGER) {
		num := math.Max(l.(*Float).Value, float64(r.(*Integer).Value))
		return &Float{Value: num}
	}
	if l.Type() != INTEGER || l.Type() != FLOAT {
		return newError("wrong type of arguments. math.max() got=%s", l.Type())
	}
	if r.Type() != INTEGER || r.Type() != FLOAT {
		return newError("wrong type of arguments. math.max() got=%s", r.Type())
	}
	return newError("wrong type of arguments. math.max()")
}

func (m *Math) random(args []Object) Object {
	if len(args) != 2 {
		return newError("wrong number of arguments. math.random() got=%d", len(args))
	}
	l := args[0]
	r := args[1]
	if (l.Type() == INTEGER) && (r.Type() == INTEGER) {
		min := l.(*Integer).Value
		max := r.(*Integer).Value
		r := rand.Int63n(max - min)
		return &Integer{Value: r + min}
	}
	return newError("wrong type of arguments. math.random()")
}
