package object

import "fmt"

type Type string

const (
	NULL     = "NULL"
	BOOLEAN  = "BOOLEAN"
	INTEGER  = "INTEGER"
	FLOAT    = "FLOAT"
	BYTES    = "BYTES"
	STRING   = "STRING"
	RETURN   = "RETURN"
	ERROR    = "ERROR"
	FUNCTION = "FUNCTION"
	BUILTIN  = "BUILTIN"
	LIST     = "LIST"
	HASH     = "HASH"
	FILE     = "FILE"
	BREAK    = "BREAK"
	CONTINUE = "CONTINUE"
	CLASS    = "CLASS"
	INSTANCE = "INSTANCE"
	THIS     = "THIS"
	MODULE   = "MODULE"
	ORIGIN   = "ORIGIN"
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

func NewError(format string, a ...interface{}) *Error {
	return &Error{Message: fmt.Sprintf(format, a...)}
}
