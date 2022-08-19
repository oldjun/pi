package module

import (
	"github.com/oldjun/pi/object"
	"os"
	"runtime"
)

// OsProperties module properties
var OsProperties = map[string]object.ModuleProperty{}

// OsFunctions module functions
var OsFunctions = map[string]object.ModuleFunction{}

func init() {
	OsProperties["name"] = name
	OsProperties["hostname"] = hostname
	OsProperties["args"] = args
	OsFunctions["exit"] = exit
	OsFunctions["getwd"] = getwd
}

func name() object.Object {
	return &object.String{Value: runtime.GOOS}
}

func hostname() object.Object {
	hostname, ok := os.Hostname()
	if ok == nil {
		return &object.String{Value: hostname}
	}
	return &object.String{Value: ""}
}

func args() object.Object {
	list := &object.List{}
	for _, val := range os.Args {
		list.Elements = append(list.Elements, &object.String{Value: val})
	}
	return list
}

func exit(args []object.Object) object.Object {
	if len(args) != 1 {
		return object.NewError("wrong number of arguments. os.exit() got=%d", len(args))
	}
	switch arg := args[0].(type) {
	case *object.Integer:
		os.Exit(int(arg.Value))
	}
	return object.NewError("wrong type of arguments. os.exit() got=%s", args[0].Type())
}

func getwd(args []object.Object) object.Object {
	if len(args) != 0 {
		return object.NewError("wrong number of arguments. os.getwd() got=%d", len(args))
	}
	path, ok := os.Getwd()
	if ok != nil {
		return object.NewError("error os.getwd() got=%s", ok.Error())
	}
	return &object.String{Value: path}
}
