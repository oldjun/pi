package ast

type Node interface{}

type Statement interface {
	Node
}

type Expression interface {
	Node
}
