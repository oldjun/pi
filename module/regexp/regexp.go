package module

import (
	"github.com/oldjun/pi/object"
	"regexp"
)

type Regexp struct {
	Handler *regexp.Regexp
}

func (r *Regexp) Method(method string, args []object.Object) object.Object {
	switch method {
	case "match":
		return r.match(args)
	case "find":
		return r.find(args)
	case "index":
		return r.index(args)
	}
	return object.NewError("regexp undefined method: %s", method)
}

func (r *Regexp) match(args []object.Object) object.Object {
	if len(args) != 1 {
		return object.NewError("wrong number of arguments. regexp.match() got=%d", len(args))
	}
	switch arg := args[0].(type) {
	case *object.String:
		b := r.Handler.MatchString(arg.Value)
		return &object.Boolean{Value: b}
	}
	return object.NewError("wrong type of arguments. regexp.match(): %s", args[0].Type())
}

func (r *Regexp) find(args []object.Object) object.Object {
	if len(args) != 1 {
		return object.NewError("wrong number of arguments. regexp.find() got=%d", len(args))
	}
	switch arg := args[0].(type) {
	case *object.String:
		s := r.Handler.FindString(arg.Value)
		return &object.String{Value: s}
	}
	return object.NewError("wrong type of arguments. regexp.find(): %s", args[0].Type())
}

func (r *Regexp) index(args []object.Object) object.Object {
	if len(args) != 1 {
		return object.NewError("wrong number of arguments. regexp.index() got=%d", len(args))
	}
	switch arg := args[0].(type) {
	case *object.String:
		pos := r.Handler.FindStringIndex(arg.Value)
		if pos == nil {
			return &object.Null{}
		}
		list := &object.List{}
		list.Elements = append(list.Elements, &object.Integer{Value: int64(pos[0])})
		list.Elements = append(list.Elements, &object.Integer{Value: int64(pos[1])})
		return list
	}
	return object.NewError("wrong type of arguments. regexp.index(): %s", args[0].Type())
}
