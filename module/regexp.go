package module

import (
	module "github.com/oldjun/pi/module/regexp"
	"github.com/oldjun/pi/object"
	"regexp"
)

// RegexpProperties module properties
var RegexpProperties = map[string]object.ModuleProperty{}

// RegexpFunctions module functions
var RegexpFunctions = map[string]object.ModuleFunction{}

func init() {
	RegexpFunctions["compile"] = compile
	RegexpFunctions["match"] = match
	RegexpFunctions["find"] = find
	RegexpFunctions["index"] = index
}

func compile(args []object.Object) object.Object {
	if len(args) != 1 {
		return object.NewError("wrong number of arguments. regexp.compile() got=%d", len(args))
	}
	switch arg := args[0].(type) {
	case *object.String:
		re, err := regexp.Compile(arg.Value)
		if err != nil {
			return object.NewError("regexp.compile() error: %s", err.Error())
		}
		return &object.Module{Name: "regexp", Handler: &module.Regexp{Handler: re}}
	}
	return object.NewError("wrong type of arguments. regexp.compile(): %s", args[0].Type())
}

func match(args []object.Object) object.Object {
	if len(args) != 2 {
		return object.NewError("wrong number of arguments. regexp.match() got=%d", len(args))
	}
	pattern := ""
	switch arg := args[0].(type) {
	case *object.String:
		pattern = arg.Value
	default:
		return object.NewError("wrong type of arguments. regexp.match(): %s", args[0].Type())
	}
	target := ""
	switch arg := args[1].(type) {
	case *object.String:
		target = arg.Value
	default:
		return object.NewError("wrong type of arguments. regexp.match(): %s", args[1].Type())
	}
	matched, err := regexp.MatchString(pattern, target)
	if err != nil {
		return object.NewError("regexp.match() error: %s", err.Error())
	}
	return &object.Boolean{Value: matched}
}

func find(args []object.Object) object.Object {
	if len(args) != 2 {
		return object.NewError("wrong number of arguments. regexp.find() got=%d", len(args))
	}
	pattern := ""
	switch arg := args[0].(type) {
	case *object.String:
		pattern = arg.Value
	default:
		return object.NewError("wrong type of arguments. regexp.find(): %s", args[0].Type())
	}
	target := ""
	switch arg := args[1].(type) {
	case *object.String:
		target = arg.Value
	default:
		return object.NewError("wrong type of arguments. regexp.find(): %s", args[1].Type())
	}
	re, err := regexp.Compile(pattern)
	if err != nil {
		return object.NewError("regexp.find() error: %s", err.Error())
	}
	str := re.FindString(target)
	return &object.String{Value: str}
}

func index(args []object.Object) object.Object {
	if len(args) != 2 {
		return object.NewError("wrong number of arguments. regexp.index() got=%d", len(args))
	}
	pattern := ""
	switch arg := args[0].(type) {
	case *object.String:
		pattern = arg.Value
	default:
		return object.NewError("wrong type of arguments. regexp.index(): %s", args[0].Type())
	}
	target := ""
	switch arg := args[1].(type) {
	case *object.String:
		target = arg.Value
	default:
		return object.NewError("wrong type of arguments. regexp.index(): %s", args[1].Type())
	}
	re, err := regexp.Compile(pattern)
	if err != nil {
		return object.NewError("regexp.index() error: %s", err.Error())
	}
	pos := re.FindStringIndex(target)
	if pos == nil {
		return &object.Null{}
	}
	list := &object.List{}
	list.Elements = append(list.Elements, &object.Integer{Value: int64(pos[0])})
	list.Elements = append(list.Elements, &object.Integer{Value: int64(pos[1])})
	return list
}
