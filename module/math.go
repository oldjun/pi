package module

import (
	"math"
	"math/rand"
	"pilang/object"
	"time"
)

// MathProperties module properties
var MathProperties = map[string]object.ModuleProperty{}

// MathFunctions module functions
var MathFunctions = map[string]object.ModuleFunction{}

func init() {
	MathProperties["pi"] = pi
	MathProperties["e"] = e
	MathFunctions["abs"] = abs
	MathFunctions["min"] = min
	MathFunctions["max"] = max
	MathFunctions["mod"] = mod
	MathFunctions["exp"] = exp
	MathFunctions["trunc"] = trunc
	MathFunctions["round"] = round
	MathFunctions["sum"] = sum
	MathFunctions["floor"] = floor
	MathFunctions["ceil"] = ceil
	MathFunctions["random"] = random
	MathFunctions["sqrt"] = sqrt
	MathFunctions["cbrt"] = cbrt
	MathFunctions["pow"] = pow
	MathFunctions["log"] = log
	MathFunctions["sin"] = sin
	MathFunctions["sinh"] = sinh
	MathFunctions["asin"] = asin
	MathFunctions["asinh"] = asinh
	MathFunctions["cos"] = cos
	MathFunctions["cosh"] = cosh
	MathFunctions["acos"] = acos
	MathFunctions["acosh"] = acosh
	MathFunctions["tan"] = tan
	MathFunctions["tanh"] = tanh
	MathFunctions["atan"] = atan
	MathFunctions["atanh"] = atanh
	MathFunctions["erf"] = erf
	MathFunctions["erfc"] = erfc
}

func pi() object.Object {
	return &object.Float{Value: 3.141592653589793}
}

func e() object.Object {
	return &object.Float{Value: 2.718281828459045}
}

func abs(args []object.Object) object.Object {
	if len(args) != 1 {
		return object.NewError("wrong number of arguments. math.abs() got=%d", len(args))
	}
	switch arg := args[0].(type) {
	case *object.Integer:
		return &object.Integer{Value: int64(math.Abs(float64(arg.Value)))}
	case *object.Float:
		return &object.Float{Value: math.Abs(arg.Value)}
	default:
		return object.NewError("wrong type of arguments. math.abs() got=%s", arg.Type())
	}
}

func min(args []object.Object) object.Object {
	if len(args) != 2 {
		return object.NewError("wrong number of arguments. math.min() got=%d", len(args))
	}
	l := args[0]
	r := args[1]
	if (l.Type() == object.INTEGER) && (r.Type() == object.INTEGER) {
		return &object.Float{Value: math.Min(float64(l.(*object.Integer).Value), float64(r.(*object.Integer).Value))}
	} else if (l.Type() == object.FLOAT) && (r.Type() == object.FLOAT) {
		return &object.Float{Value: math.Min(l.(*object.Float).Value, r.(*object.Float).Value)}
	} else if (l.Type() == object.INTEGER) && (r.Type() == object.FLOAT) {
		return &object.Float{Value: math.Min(float64(l.(*object.Integer).Value), r.(*object.Float).Value)}
	} else if (l.Type() == object.FLOAT) && (r.Type() == object.INTEGER) {
		return &object.Float{Value: math.Min(l.(*object.Float).Value, float64(r.(*object.Integer).Value))}
	}
	if l.Type() != object.INTEGER || l.Type() != object.FLOAT {
		return object.NewError("wrong type of arguments. math.min() got=%s", l.Type())
	}
	if r.Type() != object.INTEGER || r.Type() != object.FLOAT {
		return object.NewError("wrong type of arguments. math.min() got=%s", r.Type())
	}
	return object.NewError("wrong type of arguments. math.min()")
}

func max(args []object.Object) object.Object {
	if len(args) != 2 {
		return object.NewError("wrong number of arguments. math.max() got=%d", len(args))
	}
	l := args[0]
	r := args[1]
	if (l.Type() == object.INTEGER) && (r.Type() == object.INTEGER) {
		return &object.Float{Value: math.Max(float64(l.(*object.Integer).Value), float64(r.(*object.Integer).Value))}
	} else if (l.Type() == object.FLOAT) && (r.Type() == object.FLOAT) {
		return &object.Float{Value: math.Max(l.(*object.Float).Value, r.(*object.Float).Value)}
	} else if (l.Type() == object.INTEGER) && (r.Type() == object.FLOAT) {
		return &object.Float{Value: math.Max(float64(l.(*object.Integer).Value), r.(*object.Float).Value)}
	} else if (l.Type() == object.FLOAT) && (r.Type() == object.INTEGER) {
		return &object.Float{Value: math.Max(l.(*object.Float).Value, float64(r.(*object.Integer).Value))}
	}
	if l.Type() != object.INTEGER || l.Type() != object.FLOAT {
		return object.NewError("wrong type of arguments. math.max() got=%s", l.Type())
	}
	if r.Type() != object.INTEGER || r.Type() != object.FLOAT {
		return object.NewError("wrong type of arguments. math.max() got=%s", r.Type())
	}
	return object.NewError("wrong type of arguments. math.max()")
}

func mod(args []object.Object) object.Object {
	if len(args) != 2 {
		return object.NewError("wrong number of arguments. math.mod() got=%d", len(args))
	}
	l := args[0]
	r := args[1]
	if (l.Type() == object.INTEGER) && (r.Type() == object.INTEGER) {
		return &object.Float{Value: math.Mod(float64(l.(*object.Integer).Value), float64(r.(*object.Integer).Value))}
	} else if (l.Type() == object.FLOAT) && (r.Type() == object.FLOAT) {
		return &object.Float{Value: math.Mod(l.(*object.Float).Value, r.(*object.Float).Value)}
	} else if (l.Type() == object.INTEGER) && (r.Type() == object.FLOAT) {
		return &object.Float{Value: math.Mod(float64(l.(*object.Integer).Value), r.(*object.Float).Value)}
	} else if (l.Type() == object.FLOAT) && (r.Type() == object.INTEGER) {
		return &object.Float{Value: math.Mod(l.(*object.Float).Value, float64(r.(*object.Integer).Value))}
	}
	if l.Type() != object.INTEGER || l.Type() != object.FLOAT {
		return object.NewError("wrong type of arguments. math.mod() got=%s", l.Type())
	}
	if r.Type() != object.INTEGER || r.Type() != object.FLOAT {
		return object.NewError("wrong type of arguments. math.mod() got=%s", r.Type())
	}
	return object.NewError("wrong type of arguments. math.mod()")
}

func exp(args []object.Object) object.Object {
	if len(args) != 1 {
		return object.NewError("wrong number of arguments. math.exp() got=%d", len(args))
	}
	switch arg := args[0].(type) {
	case *object.Integer:
		return &object.Float{Value: math.Exp(float64(arg.Value))}
	case *object.Float:
		return &object.Float{Value: math.Exp(arg.Value)}
	}
	return object.NewError("wrong type of arguments. math.exp()")
}

func trunc(args []object.Object) object.Object {
	if len(args) != 1 {
		return object.NewError("wrong number of arguments. math.trunc() got=%d", len(args))
	}
	switch arg := args[0].(type) {
	case *object.Integer:
		return &object.Float{Value: math.Trunc(float64(arg.Value))}
	case *object.Float:
		return &object.Float{Value: math.Trunc(arg.Value)}
	}
	return object.NewError("wrong type of arguments. math.trunc()")
}

func round(args []object.Object) object.Object {
	if len(args) != 1 {
		return object.NewError("wrong number of arguments. math.round() got=%d", len(args))
	}
	switch arg := args[0].(type) {
	case *object.Integer:
		return &object.Float{Value: math.Round(float64(arg.Value))}
	case *object.Float:
		return &object.Float{Value: math.Round(arg.Value)}
	}
	return object.NewError("wrong type of arguments. math.round()")
}

func sum(args []object.Object) object.Object {
	if len(args) < 1 {
		return object.NewError("wrong number of arguments. math.sum() got=%d", len(args))
	}
	sum := 0.0
	for _, arg := range args {
		switch val := arg.(type) {
		case *object.Integer:
			sum += float64(val.Value)
		case *object.Float:
			sum += val.Value
		default:
			return object.NewError("wrong type of arguments. math.sum()")
		}
	}
	return &object.Float{Value: sum}
}

func floor(args []object.Object) object.Object {
	if len(args) != 1 {
		return object.NewError("wrong number of arguments. math.floor() got=%d", len(args))
	}
	switch arg := args[0].(type) {
	case *object.Integer:
		return &object.Integer{Value: arg.Value}
	case *object.Float:
		return &object.Integer{Value: int64(math.Floor(arg.Value))}
	}
	return object.NewError("wrong type of arguments. math.floor()")
}

func ceil(args []object.Object) object.Object {
	if len(args) != 1 {
		return object.NewError("wrong number of arguments. math.ceil() got=%d", len(args))
	}
	switch arg := args[0].(type) {
	case *object.Integer:
		return &object.Integer{Value: arg.Value}
	case *object.Float:
		return &object.Integer{Value: int64(math.Ceil(arg.Value))}
	}
	return object.NewError("wrong type of arguments. math.ceil()")
}

func random(args []object.Object) object.Object {
	if len(args) != 2 {
		return object.NewError("wrong number of arguments. math.random() got=%d", len(args))
	}
	l := args[0]
	r := args[1]
	if (l.Type() == object.INTEGER) && (r.Type() == object.INTEGER) {
		rand.Seed(time.Now().UnixNano())
		min := l.(*object.Integer).Value
		max := r.(*object.Integer).Value
		r := rand.Int63n(max - min)
		return &object.Integer{Value: r + min}
	}
	return object.NewError("wrong type of arguments. math.random()")
}

func sqrt(args []object.Object) object.Object {
	if len(args) != 1 {
		return object.NewError("wrong number of arguments. math.sqrt() got=%d", len(args))
	}
	switch arg := args[0].(type) {
	case *object.Integer:
		return &object.Float{Value: math.Sqrt(float64(arg.Value))}
	case *object.Float:
		return &object.Float{Value: math.Sqrt(arg.Value)}
	}
	return object.NewError("wrong type of arguments. math.sqrt()")
}

func cbrt(args []object.Object) object.Object {
	if len(args) != 1 {
		return object.NewError("wrong number of arguments. math.cbrt() got=%d", len(args))
	}
	switch arg := args[0].(type) {
	case *object.Integer:
		return &object.Float{Value: math.Cbrt(float64(arg.Value))}
	case *object.Float:
		return &object.Float{Value: math.Cbrt(arg.Value)}
	}
	return object.NewError("wrong type of arguments. math.cbrt()")
}

func pow(args []object.Object) object.Object {
	if len(args) != 2 {
		return object.NewError("wrong number of arguments. math.pow() got=%d", len(args))
	}
	l := args[0]
	r := args[1]
	if (l.Type() == object.INTEGER) && (r.Type() == object.INTEGER) {
		return &object.Float{Value: math.Pow(float64(l.(*object.Integer).Value), float64(r.(*object.Integer).Value))}
	} else if (l.Type() == object.FLOAT) && (r.Type() == object.FLOAT) {
		return &object.Float{Value: math.Pow(l.(*object.Float).Value, r.(*object.Float).Value)}
	} else if (l.Type() == object.INTEGER) && (r.Type() == object.FLOAT) {
		return &object.Float{Value: math.Pow(float64(l.(*object.Integer).Value), r.(*object.Float).Value)}
	} else if (l.Type() == object.FLOAT) && (r.Type() == object.INTEGER) {
		return &object.Float{Value: math.Pow(l.(*object.Float).Value, float64(r.(*object.Integer).Value))}
	}
	if l.Type() != object.INTEGER || l.Type() != object.FLOAT {
		return object.NewError("wrong type of arguments. math.pow() got=%s", l.Type())
	}
	if r.Type() != object.INTEGER || r.Type() != object.FLOAT {
		return object.NewError("wrong type of arguments. math.pow() got=%s", r.Type())
	}
	return object.NewError("wrong type of arguments. math.pow()")
}

func log(args []object.Object) object.Object {
	if len(args) != 1 {
		return object.NewError("wrong number of arguments. math.log() got=%d", len(args))
	}
	switch arg := args[0].(type) {
	case *object.Integer:
		return &object.Float{Value: math.Log(float64(arg.Value))}
	case *object.Float:
		return &object.Float{Value: math.Log(arg.Value)}
	}
	return object.NewError("wrong type of arguments. math.log()")
}

func sin(args []object.Object) object.Object {
	if len(args) != 1 {
		return object.NewError("wrong number of arguments. math.sin() got=%d", len(args))
	}
	switch arg := args[0].(type) {
	case *object.Integer:
		return &object.Float{Value: math.Sin(float64(arg.Value))}
	case *object.Float:
		return &object.Float{Value: math.Sin(arg.Value)}
	}
	return object.NewError("wrong type of arguments. math.sin()")
}

func sinh(args []object.Object) object.Object {
	if len(args) != 1 {
		return object.NewError("wrong number of arguments. math.sinh() got=%d", len(args))
	}
	switch arg := args[0].(type) {
	case *object.Integer:
		return &object.Float{Value: math.Sinh(float64(arg.Value))}
	case *object.Float:
		return &object.Float{Value: math.Sinh(arg.Value)}
	}
	return object.NewError("wrong type of arguments. math.sinh()")
}

func asin(args []object.Object) object.Object {
	if len(args) != 1 {
		return object.NewError("wrong number of arguments. math.asin() got=%d", len(args))
	}
	switch arg := args[0].(type) {
	case *object.Integer:
		return &object.Float{Value: math.Asin(float64(arg.Value))}
	case *object.Float:
		return &object.Float{Value: math.Asin(arg.Value)}
	}
	return object.NewError("wrong type of arguments. math.asin()")
}

func asinh(args []object.Object) object.Object {
	if len(args) != 1 {
		return object.NewError("wrong number of arguments. math.asinh() got=%d", len(args))
	}
	switch arg := args[0].(type) {
	case *object.Integer:
		return &object.Float{Value: math.Asinh(float64(arg.Value))}
	case *object.Float:
		return &object.Float{Value: math.Asinh(arg.Value)}
	}
	return object.NewError("wrong type of arguments. math.asinh()")
}

func cos(args []object.Object) object.Object {
	if len(args) != 1 {
		return object.NewError("wrong number of arguments. math.cos() got=%d", len(args))
	}
	switch arg := args[0].(type) {
	case *object.Integer:
		return &object.Float{Value: math.Cos(float64(arg.Value))}
	case *object.Float:
		return &object.Float{Value: math.Cos(arg.Value)}
	}
	return object.NewError("wrong type of arguments. math.cos()")
}

func cosh(args []object.Object) object.Object {
	if len(args) != 1 {
		return object.NewError("wrong number of arguments. math.cosh() got=%d", len(args))
	}
	switch arg := args[0].(type) {
	case *object.Integer:
		return &object.Float{Value: math.Cosh(float64(arg.Value))}
	case *object.Float:
		return &object.Float{Value: math.Cosh(arg.Value)}
	}
	return object.NewError("wrong type of arguments. math.cosh()")
}

func acos(args []object.Object) object.Object {
	if len(args) != 1 {
		return object.NewError("wrong number of arguments. math.acos() got=%d", len(args))
	}
	switch arg := args[0].(type) {
	case *object.Integer:
		return &object.Float{Value: math.Acos(float64(arg.Value))}
	case *object.Float:
		return &object.Float{Value: math.Acos(arg.Value)}
	}
	return object.NewError("wrong type of arguments. math.acos()")
}

func acosh(args []object.Object) object.Object {
	if len(args) != 1 {
		return object.NewError("wrong number of arguments. math.acosh() got=%d", len(args))
	}
	switch arg := args[0].(type) {
	case *object.Integer:
		return &object.Float{Value: math.Acosh(float64(arg.Value))}
	case *object.Float:
		return &object.Float{Value: math.Acosh(arg.Value)}
	}
	return object.NewError("wrong type of arguments. math.acosh()")
}

func tan(args []object.Object) object.Object {
	if len(args) != 1 {
		return object.NewError("wrong number of arguments. math.tan() got=%d", len(args))
	}
	switch arg := args[0].(type) {
	case *object.Integer:
		return &object.Float{Value: math.Tan(float64(arg.Value))}
	case *object.Float:
		return &object.Float{Value: math.Tan(arg.Value)}
	}
	return object.NewError("wrong type of arguments. math.tan()")
}

func tanh(args []object.Object) object.Object {
	if len(args) != 1 {
		return object.NewError("wrong number of arguments. math.tanh() got=%d", len(args))
	}
	switch arg := args[0].(type) {
	case *object.Integer:
		return &object.Float{Value: math.Tanh(float64(arg.Value))}
	case *object.Float:
		return &object.Float{Value: math.Tanh(arg.Value)}
	}
	return object.NewError("wrong type of arguments. math.tanh()")
}

func atan(args []object.Object) object.Object {
	if len(args) != 1 {
		return object.NewError("wrong number of arguments. math.atan() got=%d", len(args))
	}
	switch arg := args[0].(type) {
	case *object.Integer:
		return &object.Float{Value: math.Atan(float64(arg.Value))}
	case *object.Float:
		return &object.Float{Value: math.Atan(arg.Value)}
	}
	return object.NewError("wrong type of arguments. math.atan()")
}

func atanh(args []object.Object) object.Object {
	if len(args) != 1 {
		return object.NewError("wrong number of arguments. math.atanh() got=%d", len(args))
	}
	switch arg := args[0].(type) {
	case *object.Integer:
		return &object.Float{Value: math.Atanh(float64(arg.Value))}
	case *object.Float:
		return &object.Float{Value: math.Atanh(arg.Value)}
	}
	return object.NewError("wrong type of arguments. math.atanh()")
}

func erf(args []object.Object) object.Object {
	if len(args) != 1 {
		return object.NewError("wrong number of arguments. math.erf() got=%d", len(args))
	}
	switch arg := args[0].(type) {
	case *object.Integer:
		return &object.Float{Value: math.Erf(float64(arg.Value))}
	case *object.Float:
		return &object.Float{Value: math.Erf(arg.Value)}
	}
	return object.NewError("wrong type of arguments. math.erf()")
}

func erfc(args []object.Object) object.Object {
	if len(args) != 1 {
		return object.NewError("wrong number of arguments. math.erfc() got=%d", len(args))
	}
	switch arg := args[0].(type) {
	case *object.Integer:
		return &object.Float{Value: math.Erfc(float64(arg.Value))}
	case *object.Float:
		return &object.Float{Value: math.Erfc(arg.Value)}
	}
	return object.NewError("wrong type of arguments. math.erfc()")
}
