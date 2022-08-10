package object

import "fmt"

type Type string

type BuiltinFunction func(args []Object) Object

const (
	INTEGER  = "INTEGER"
	BOOLEAN  = "BOOLEAN"
	FLOAT    = "FLOAT"
	NULL     = "NULL"
	RETURN   = "RETURN"
	ERROR    = "ERROR"
	FUNCTION = "FUNCTION"
	STRING   = "STRING"
	BUILTIN  = "BUILTIN"
	ARRAY    = "ARRAY"
	HASH     = "HASH"
	FILE     = "FILE"
	BREAK    = "BREAK"
	CONTINUE = "CONTINUE"
	CLASS    = "CLASS"
	INSTANCE = "INSTANCE"
	THIS     = "THIS"
)

type Object interface {
	Type() Type
	String() string
}

type Hashable interface {
	HashKey() HashKey
}

type Iterable interface {
	Next() (Object, Object)
	Reset()
}

func newError(format string, a ...interface{}) *Error {
	return &Error{Message: fmt.Sprintf(format, a...)}
}
