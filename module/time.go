package module

import (
	"github.com/oldjun/pi/object"
	"strings"
	"time"
)

// TimeProperties module properties
var TimeProperties = map[string]object.ModuleProperty{}

// TimeFunctions module functions
var TimeFunctions = map[string]object.ModuleFunction{}

func init() {
	TimeFunctions["time"] = now
	TimeFunctions["sleep"] = sleep
	TimeFunctions["strftime"] = strftime
	TimeFunctions["strptime"] = strptime
}

func now(args []object.Object) object.Object {
	if len(args) != 0 {
		return object.NewError("wrong number of arguments. time.now() got=%d", len(args))
	}
	return &object.Float{Value: float64(time.Now().UnixNano()) / float64(time.Second)}
}

func sleep(args []object.Object) object.Object {
	if len(args) != 1 {
		return object.NewError("wrong number of arguments. time.sleep() got=%d", len(args))
	}
	switch arg := args[0].(type) {
	case *object.Integer:
		time.Sleep(time.Duration(arg.Value) * time.Millisecond)
		return &object.Null{}
	}
	return object.NewError("argument to time.sleep() not supported, got %s", args[0].Type())
}

func strftime(args []object.Object) object.Object {
	if len(args) != 2 {
		return object.NewError("wrong number of arguments. time.strftime() got=%d", len(args))
	}
	format := ""
	switch f := args[0].(type) {
	case *object.String:
		r := strings.NewReplacer("%Y", "2006", "%m", "01", "%d", "02", "%H", "15", "%M", "04", "%S", "05")
		format = r.Replace(f.Value)
	default:
		return object.NewError("argument to time.strftime() not supported, got %s", args[0].Type())
	}
	switch sec := args[1].(type) {
	case *object.Integer:
		t := time.Unix(sec.Value, 0)
		str := t.Format(format)
		return &object.String{Value: str}
	case *object.Float:
		t := time.Unix(int64(sec.Value), 0)
		str := t.Format(format)
		return &object.String{Value: str}
	}
	return object.NewError("argument to time.strftime() not supported, got %s", args[1].Type())
}

func strptime(args []object.Object) object.Object {
	if len(args) != 2 {
		return object.NewError("wrong number of arguments. time.strptime() got=%d", len(args))
	}
	str := ""
	switch s := args[0].(type) {
	case *object.String:
		str = s.Value
	default:
		return object.NewError("argument to time.strptime() not supported, got %s", args[0].Type())
	}
	format := ""
	switch f := args[1].(type) {
	case *object.String:
		r := strings.NewReplacer("%Y", "2006", "%m", "01", "%d", "02", "%H", "15", "%M", "04", "%S", "05")
		format = r.Replace(f.Value)
	default:
		return object.NewError("argument to time.strptime() not supported, got %s", args[1].Type())
	}
	t, err := time.ParseInLocation(format, str, time.Local)
	if err != nil {
		return object.NewError("time.strptime() error")
	}
	return &object.Float{Value: float64(t.UnixNano()) / float64(time.Second)}
}
