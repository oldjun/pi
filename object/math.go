package object

import (
	"fmt"
	"math"
	"math/rand"
	"time"
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
	case "floor":
		return m.floor(args)
	case "ceil":
		return m.ceil(args)
	case "random":
		return m.random(args)
	//case "round":
	//	return m.round(args)
	case "sqrt":
		return m.sqrt(args)
	case "cbrt":
		return m.cbrt(args)
	case "sin":
		return m.sin(args)
	case "sinh":
		return m.sinh(args)
	case "asin":
		return m.asin(args)
	case "asinh":
		return m.asinh(args)
	case "cos":
		return m.cos(args)
	case "cosh":
		return m.cosh(args)
	case "acos":
		return m.acos(args)
	case "acosh":
		return m.acosh(args)
	case "tan":
		return m.tan(args)
	case "tanh":
		return m.tanh(args)
	case "atan":
		return m.atan(args)
	case "atanh":
		return m.atanh(args)
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

func (m *Math) floor(args []Object) Object {
	if len(args) != 2 {
		return newError("wrong number of arguments. math.floor() got=%d", len(args))
	}
	switch arg := args[0].(type) {
	case *Integer:
		return &Integer{Value: arg.Value}
	case *Float:
		num := math.Floor(arg.Value)
		return &Integer{Value: int64(num)}
	}
	return newError("wrong type of arguments. math.floor()")
}

func (m *Math) ceil(args []Object) Object {
	if len(args) != 2 {
		return newError("wrong number of arguments. math.ceil() got=%d", len(args))
	}
	switch arg := args[0].(type) {
	case *Integer:
		return &Integer{Value: arg.Value}
	case *Float:
		num := math.Ceil(arg.Value)
		return &Integer{Value: int64(num)}
	}
	return newError("wrong type of arguments. math.ceil()")
}

func (m *Math) random(args []Object) Object {
	if len(args) != 2 {
		return newError("wrong number of arguments. math.random() got=%d", len(args))
	}
	l := args[0]
	r := args[1]
	if (l.Type() == INTEGER) && (r.Type() == INTEGER) {
		rand.Seed(time.Now().UnixNano())
		min := l.(*Integer).Value
		max := r.(*Integer).Value
		r := rand.Int63n(max - min)
		return &Integer{Value: r + min}
	}
	return newError("wrong type of arguments. math.random()")
}

func (m *Math) sqrt(args []Object) Object {
	if len(args) != 1 {
		return newError("wrong number of arguments. math.sqrt() got=%d", len(args))
	}
	switch arg := args[0].(type) {
	case *Integer:
		num := math.Sqrt(float64(arg.Value))
		return &Float{Value: num}
	case *Float:
		num := math.Sqrt(arg.Value)
		return &Float{Value: num}
	}
	return newError("wrong type of arguments. math.sqrt()")
}

func (m *Math) cbrt(args []Object) Object {
	if len(args) != 1 {
		return newError("wrong number of arguments. math.cbrt() got=%d", len(args))
	}
	switch arg := args[0].(type) {
	case *Integer:
		num := math.Cbrt(float64(arg.Value))
		return &Float{Value: num}
	case *Float:
		num := math.Cbrt(arg.Value)
		return &Float{Value: num}
	}
	return newError("wrong type of arguments. math.cbrt()")
}

func (m *Math) sin(args []Object) Object {
	if len(args) != 1 {
		return newError("wrong number of arguments. math.sin() got=%d", len(args))
	}
	switch arg := args[0].(type) {
	case *Integer:
		val := math.Sin(float64(arg.Value))
		return &Float{Value: val}
	case *Float:
		val := math.Sin(arg.Value)
		return &Float{Value: val}
	}
	return newError("wrong type of arguments. math.sin()")
}

func (m *Math) sinh(args []Object) Object {
	if len(args) != 1 {
		return newError("wrong number of arguments. math.sinh() got=%d", len(args))
	}
	switch arg := args[0].(type) {
	case *Integer:
		val := math.Sinh(float64(arg.Value))
		return &Float{Value: val}
	case *Float:
		val := math.Sinh(arg.Value)
		return &Float{Value: val}
	}
	return newError("wrong type of arguments. math.sinh()")
}

func (m *Math) asin(args []Object) Object {
	if len(args) != 1 {
		return newError("wrong number of arguments. math.asin() got=%d", len(args))
	}
	switch arg := args[0].(type) {
	case *Integer:
		val := math.Asin(float64(arg.Value))
		return &Float{Value: val}
	case *Float:
		val := math.Asin(arg.Value)
		return &Float{Value: val}
	}
	return newError("wrong type of arguments. math.asin()")
}

func (m *Math) asinh(args []Object) Object {
	if len(args) != 1 {
		return newError("wrong number of arguments. math.asinh() got=%d", len(args))
	}
	switch arg := args[0].(type) {
	case *Integer:
		val := math.Asinh(float64(arg.Value))
		return &Float{Value: val}
	case *Float:
		val := math.Asinh(arg.Value)
		return &Float{Value: val}
	}
	return newError("wrong type of arguments. math.asinh()")
}

func (m *Math) cos(args []Object) Object {
	if len(args) != 1 {
		return newError("wrong number of arguments. math.cos() got=%d", len(args))
	}
	switch arg := args[0].(type) {
	case *Integer:
		val := math.Cos(float64(arg.Value))
		return &Float{Value: val}
	case *Float:
		val := math.Cos(arg.Value)
		return &Float{Value: val}
	}
	return newError("wrong type of arguments. math.cos()")
}

func (m *Math) cosh(args []Object) Object {
	if len(args) != 1 {
		return newError("wrong number of arguments. math.cosh() got=%d", len(args))
	}
	switch arg := args[0].(type) {
	case *Integer:
		val := math.Cosh(float64(arg.Value))
		return &Float{Value: val}
	case *Float:
		val := math.Cosh(arg.Value)
		return &Float{Value: val}
	}
	return newError("wrong type of arguments. math.cosh()")
}

func (m *Math) acos(args []Object) Object {
	if len(args) != 1 {
		return newError("wrong number of arguments. math.acos() got=%d", len(args))
	}
	switch arg := args[0].(type) {
	case *Integer:
		val := math.Acos(float64(arg.Value))
		return &Float{Value: val}
	case *Float:
		val := math.Acos(arg.Value)
		return &Float{Value: val}
	}
	return newError("wrong type of arguments. math.acos()")
}

func (m *Math) acosh(args []Object) Object {
	if len(args) != 1 {
		return newError("wrong number of arguments. math.acosh() got=%d", len(args))
	}
	switch arg := args[0].(type) {
	case *Integer:
		val := math.Acosh(float64(arg.Value))
		return &Float{Value: val}
	case *Float:
		val := math.Acosh(arg.Value)
		return &Float{Value: val}
	}
	return newError("wrong type of arguments. math.acosh()")
}

func (m *Math) tan(args []Object) Object {
	if len(args) != 1 {
		return newError("wrong number of arguments. math.tan() got=%d", len(args))
	}
	switch arg := args[0].(type) {
	case *Integer:
		val := math.Tan(float64(arg.Value))
		return &Float{Value: val}
	case *Float:
		val := math.Tan(arg.Value)
		return &Float{Value: val}
	}
	return newError("wrong type of arguments. math.tan()")
}

func (m *Math) tanh(args []Object) Object {
	if len(args) != 1 {
		return newError("wrong number of arguments. math.tanh() got=%d", len(args))
	}
	switch arg := args[0].(type) {
	case *Integer:
		val := math.Tanh(float64(arg.Value))
		return &Float{Value: val}
	case *Float:
		val := math.Tanh(arg.Value)
		return &Float{Value: val}
	}
	return newError("wrong type of arguments. math.tanh()")
}

func (m *Math) atan(args []Object) Object {
	if len(args) != 1 {
		return newError("wrong number of arguments. math.atan() got=%d", len(args))
	}
	switch arg := args[0].(type) {
	case *Integer:
		val := math.Atan(float64(arg.Value))
		return &Float{Value: val}
	case *Float:
		val := math.Atan(arg.Value)
		return &Float{Value: val}
	}
	return newError("wrong type of arguments. math.atan()")
}

func (m *Math) atanh(args []Object) Object {
	if len(args) != 1 {
		return newError("wrong number of arguments. math.atanh() got=%d", len(args))
	}
	switch arg := args[0].(type) {
	case *Integer:
		val := math.Atanh(float64(arg.Value))
		return &Float{Value: val}
	case *Float:
		val := math.Atanh(arg.Value)
		return &Float{Value: val}
	}
	return newError("wrong type of arguments. math.atanh()")
}
