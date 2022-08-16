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

func (m *Math) Property(property string) Object {
	switch property {
	case "pi":
		return m.pi()
	case "e":
		return m.e()
	}
	return nil
}

func (m *Math) pi() Object {
	return &Float{Value: 3.141592653589793}
}

func (m *Math) e() Object {
	return &Float{Value: 2.718281828459045}
}

func (m *Math) Method(method string, args []Object) Object {
	switch method {
	case "abs":
		return m.abs(args)
	case "min":
		return m.min(args)
	case "max":
		return m.max(args)
	case "mod":
		return m.mod(args)
	case "exp":
		return m.exp(args)
	case "trunc":
		return m.trunc(args)
	case "round":
		return m.round(args)
	case "sum":
		return m.sum(args)
	case "floor":
		return m.floor(args)
	case "ceil":
		return m.ceil(args)
	case "random":
		return m.random(args)
	case "sqrt":
		return m.sqrt(args)
	case "cbrt":
		return m.cbrt(args)
	case "pow":
		return m.pow(args)
	case "log":
		return m.log(args)
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
	case "erf":
		return m.erf(args)
	case "erfc":
		return m.erfc(args)
	}
	return newError("math module function not exists: %s", method)
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
		return &Float{Value: math.Min(float64(l.(*Integer).Value), float64(r.(*Integer).Value))}
	} else if (l.Type() == FLOAT) && (r.Type() == FLOAT) {
		return &Float{Value: math.Min(l.(*Float).Value, r.(*Float).Value)}
	} else if (l.Type() == INTEGER) && (r.Type() == FLOAT) {
		return &Float{Value: math.Min(float64(l.(*Integer).Value), r.(*Float).Value)}
	} else if (l.Type() == FLOAT) && (r.Type() == INTEGER) {
		return &Float{Value: math.Min(l.(*Float).Value, float64(r.(*Integer).Value))}
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
		return &Float{Value: math.Max(float64(l.(*Integer).Value), float64(r.(*Integer).Value))}
	} else if (l.Type() == FLOAT) && (r.Type() == FLOAT) {
		return &Float{Value: math.Max(l.(*Float).Value, r.(*Float).Value)}
	} else if (l.Type() == INTEGER) && (r.Type() == FLOAT) {
		return &Float{Value: math.Max(float64(l.(*Integer).Value), r.(*Float).Value)}
	} else if (l.Type() == FLOAT) && (r.Type() == INTEGER) {
		return &Float{Value: math.Max(l.(*Float).Value, float64(r.(*Integer).Value))}
	}
	if l.Type() != INTEGER || l.Type() != FLOAT {
		return newError("wrong type of arguments. math.max() got=%s", l.Type())
	}
	if r.Type() != INTEGER || r.Type() != FLOAT {
		return newError("wrong type of arguments. math.max() got=%s", r.Type())
	}
	return newError("wrong type of arguments. math.max()")
}

func (m *Math) mod(args []Object) Object {
	if len(args) != 2 {
		return newError("wrong number of arguments. math.mod() got=%d", len(args))
	}
	l := args[0]
	r := args[1]
	if (l.Type() == INTEGER) && (r.Type() == INTEGER) {
		return &Float{Value: math.Mod(float64(l.(*Integer).Value), float64(r.(*Integer).Value))}
	} else if (l.Type() == FLOAT) && (r.Type() == FLOAT) {
		return &Float{Value: math.Mod(l.(*Float).Value, r.(*Float).Value)}
	} else if (l.Type() == INTEGER) && (r.Type() == FLOAT) {
		return &Float{Value: math.Mod(float64(l.(*Integer).Value), r.(*Float).Value)}
	} else if (l.Type() == FLOAT) && (r.Type() == INTEGER) {
		return &Float{Value: math.Mod(l.(*Float).Value, float64(r.(*Integer).Value))}
	}
	if l.Type() != INTEGER || l.Type() != FLOAT {
		return newError("wrong type of arguments. math.mod() got=%s", l.Type())
	}
	if r.Type() != INTEGER || r.Type() != FLOAT {
		return newError("wrong type of arguments. math.mod() got=%s", r.Type())
	}
	return newError("wrong type of arguments. math.mod()")
}

func (m *Math) exp(args []Object) Object {
	if len(args) != 1 {
		return newError("wrong number of arguments. math.exp() got=%d", len(args))
	}
	switch arg := args[0].(type) {
	case *Integer:
		return &Float{Value: math.Exp(float64(arg.Value))}
	case *Float:
		return &Float{Value: math.Exp(arg.Value)}
	}
	return newError("wrong type of arguments. math.exp()")
}

func (m *Math) trunc(args []Object) Object {
	if len(args) != 1 {
		return newError("wrong number of arguments. math.trunc() got=%d", len(args))
	}
	switch arg := args[0].(type) {
	case *Integer:
		return &Float{Value: math.Trunc(float64(arg.Value))}
	case *Float:
		return &Float{Value: math.Trunc(arg.Value)}
	}
	return newError("wrong type of arguments. math.trunc()")
}

func (m *Math) round(args []Object) Object {
	if len(args) != 1 {
		return newError("wrong number of arguments. math.round() got=%d", len(args))
	}
	switch arg := args[0].(type) {
	case *Integer:
		return &Float{Value: math.Round(float64(arg.Value))}
	case *Float:
		return &Float{Value: math.Round(arg.Value)}
	}
	return newError("wrong type of arguments. math.round()")
}

func (m *Math) sum(args []Object) Object {
	if len(args) < 1 {
		return newError("wrong number of arguments. math.sum() got=%d", len(args))
	}
	sum := 0.0
	for _, arg := range args {
		switch val := arg.(type) {
		case *Integer:
			sum += float64(val.Value)
		case *Float:
			sum += val.Value
		default:
			return newError("wrong type of arguments. math.sum()")
		}
	}
	return &Float{Value: sum}
}

func (m *Math) floor(args []Object) Object {
	if len(args) != 1 {
		return newError("wrong number of arguments. math.floor() got=%d", len(args))
	}
	switch arg := args[0].(type) {
	case *Integer:
		return &Integer{Value: arg.Value}
	case *Float:
		return &Integer{Value: int64(math.Floor(arg.Value))}
	}
	return newError("wrong type of arguments. math.floor()")
}

func (m *Math) ceil(args []Object) Object {
	if len(args) != 1 {
		return newError("wrong number of arguments. math.ceil() got=%d", len(args))
	}
	switch arg := args[0].(type) {
	case *Integer:
		return &Integer{Value: arg.Value}
	case *Float:
		return &Integer{Value: int64(math.Ceil(arg.Value))}
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
		return &Float{Value: math.Sqrt(float64(arg.Value))}
	case *Float:
		return &Float{Value: math.Sqrt(arg.Value)}
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

func (m *Math) pow(args []Object) Object {
	if len(args) != 2 {
		return newError("wrong number of arguments. math.pow() got=%d", len(args))
	}
	l := args[0]
	r := args[1]
	if (l.Type() == INTEGER) && (r.Type() == INTEGER) {
		return &Float{Value: math.Pow(float64(l.(*Integer).Value), float64(r.(*Integer).Value))}
	} else if (l.Type() == FLOAT) && (r.Type() == FLOAT) {
		return &Float{Value: math.Pow(l.(*Float).Value, r.(*Float).Value)}
	} else if (l.Type() == INTEGER) && (r.Type() == FLOAT) {
		return &Float{Value: math.Pow(float64(l.(*Integer).Value), r.(*Float).Value)}
	} else if (l.Type() == FLOAT) && (r.Type() == INTEGER) {
		return &Float{Value: math.Pow(l.(*Float).Value, float64(r.(*Integer).Value))}
	}
	if l.Type() != INTEGER || l.Type() != FLOAT {
		return newError("wrong type of arguments. math.pow() got=%s", l.Type())
	}
	if r.Type() != INTEGER || r.Type() != FLOAT {
		return newError("wrong type of arguments. math.pow() got=%s", r.Type())
	}
	return newError("wrong type of arguments. math.pow()")
}

func (m *Math) log(args []Object) Object {
	if len(args) != 1 {
		return newError("wrong number of arguments. math.log() got=%d", len(args))
	}
	switch arg := args[0].(type) {
	case *Integer:
		return &Float{Value: math.Log(float64(arg.Value))}
	case *Float:
		return &Float{Value: math.Log(arg.Value)}
	}
	return newError("wrong type of arguments. math.log()")
}

func (m *Math) sin(args []Object) Object {
	if len(args) != 1 {
		return newError("wrong number of arguments. math.sin() got=%d", len(args))
	}
	switch arg := args[0].(type) {
	case *Integer:
		return &Float{Value: math.Sin(float64(arg.Value))}
	case *Float:
		return &Float{Value: math.Sin(arg.Value)}
	}
	return newError("wrong type of arguments. math.sin()")
}

func (m *Math) sinh(args []Object) Object {
	if len(args) != 1 {
		return newError("wrong number of arguments. math.sinh() got=%d", len(args))
	}
	switch arg := args[0].(type) {
	case *Integer:
		return &Float{Value: math.Sinh(float64(arg.Value))}
	case *Float:
		return &Float{Value: math.Sinh(arg.Value)}
	}
	return newError("wrong type of arguments. math.sinh()")
}

func (m *Math) asin(args []Object) Object {
	if len(args) != 1 {
		return newError("wrong number of arguments. math.asin() got=%d", len(args))
	}
	switch arg := args[0].(type) {
	case *Integer:
		return &Float{Value: math.Asin(float64(arg.Value))}
	case *Float:
		return &Float{Value: math.Asin(arg.Value)}
	}
	return newError("wrong type of arguments. math.asin()")
}

func (m *Math) asinh(args []Object) Object {
	if len(args) != 1 {
		return newError("wrong number of arguments. math.asinh() got=%d", len(args))
	}
	switch arg := args[0].(type) {
	case *Integer:
		return &Float{Value: math.Asinh(float64(arg.Value))}
	case *Float:
		return &Float{Value: math.Asinh(arg.Value)}
	}
	return newError("wrong type of arguments. math.asinh()")
}

func (m *Math) cos(args []Object) Object {
	if len(args) != 1 {
		return newError("wrong number of arguments. math.cos() got=%d", len(args))
	}
	switch arg := args[0].(type) {
	case *Integer:
		return &Float{Value: math.Cos(float64(arg.Value))}
	case *Float:
		return &Float{Value: math.Cos(arg.Value)}
	}
	return newError("wrong type of arguments. math.cos()")
}

func (m *Math) cosh(args []Object) Object {
	if len(args) != 1 {
		return newError("wrong number of arguments. math.cosh() got=%d", len(args))
	}
	switch arg := args[0].(type) {
	case *Integer:
		return &Float{Value: math.Cosh(float64(arg.Value))}
	case *Float:
		return &Float{Value: math.Cosh(arg.Value)}
	}
	return newError("wrong type of arguments. math.cosh()")
}

func (m *Math) acos(args []Object) Object {
	if len(args) != 1 {
		return newError("wrong number of arguments. math.acos() got=%d", len(args))
	}
	switch arg := args[0].(type) {
	case *Integer:
		return &Float{Value: math.Acos(float64(arg.Value))}
	case *Float:
		return &Float{Value: math.Acos(arg.Value)}
	}
	return newError("wrong type of arguments. math.acos()")
}

func (m *Math) acosh(args []Object) Object {
	if len(args) != 1 {
		return newError("wrong number of arguments. math.acosh() got=%d", len(args))
	}
	switch arg := args[0].(type) {
	case *Integer:
		return &Float{Value: math.Acosh(float64(arg.Value))}
	case *Float:
		return &Float{Value: math.Acosh(arg.Value)}
	}
	return newError("wrong type of arguments. math.acosh()")
}

func (m *Math) tan(args []Object) Object {
	if len(args) != 1 {
		return newError("wrong number of arguments. math.tan() got=%d", len(args))
	}
	switch arg := args[0].(type) {
	case *Integer:
		return &Float{Value: math.Tan(float64(arg.Value))}
	case *Float:
		return &Float{Value: math.Tan(arg.Value)}
	}
	return newError("wrong type of arguments. math.tan()")
}

func (m *Math) tanh(args []Object) Object {
	if len(args) != 1 {
		return newError("wrong number of arguments. math.tanh() got=%d", len(args))
	}
	switch arg := args[0].(type) {
	case *Integer:
		return &Float{Value: math.Tanh(float64(arg.Value))}
	case *Float:
		return &Float{Value: math.Tanh(arg.Value)}
	}
	return newError("wrong type of arguments. math.tanh()")
}

func (m *Math) atan(args []Object) Object {
	if len(args) != 1 {
		return newError("wrong number of arguments. math.atan() got=%d", len(args))
	}
	switch arg := args[0].(type) {
	case *Integer:
		return &Float{Value: math.Atan(float64(arg.Value))}
	case *Float:
		return &Float{Value: math.Atan(arg.Value)}
	}
	return newError("wrong type of arguments. math.atan()")
}

func (m *Math) atanh(args []Object) Object {
	if len(args) != 1 {
		return newError("wrong number of arguments. math.atanh() got=%d", len(args))
	}
	switch arg := args[0].(type) {
	case *Integer:
		return &Float{Value: math.Atanh(float64(arg.Value))}
	case *Float:
		return &Float{Value: math.Atanh(arg.Value)}
	}
	return newError("wrong type of arguments. math.atanh()")
}

func (m *Math) erf(args []Object) Object {
	if len(args) != 1 {
		return newError("wrong number of arguments. math.erf() got=%d", len(args))
	}
	switch arg := args[0].(type) {
	case *Integer:
		return &Float{Value: math.Erf(float64(arg.Value))}
	case *Float:
		return &Float{Value: math.Erf(arg.Value)}
	}
	return newError("wrong type of arguments. math.erf()")
}

func (m *Math) erfc(args []Object) Object {
	if len(args) != 1 {
		return newError("wrong number of arguments. math.erfc() got=%d", len(args))
	}
	switch arg := args[0].(type) {
	case *Integer:
		return &Float{Value: math.Erfc(float64(arg.Value))}
	case *Float:
		return &Float{Value: math.Erfc(arg.Value)}
	}
	return newError("wrong type of arguments. math.erfc()")
}
