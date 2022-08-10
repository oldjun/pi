package main

import (
	"fmt"
	"os"
	"pilang/evaluator"
	"pilang/lexer"
	"pilang/object"
	"pilang/parser"
)

func main() {
	file := "demo.pi"
	//file := os.Args[1]
	input, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}
	l := lexer.New(string(input), file)
	p := parser.New(l)
	program := p.ParseProgram()
	env := object.NewEnvironment()
	evaluated := evaluator.Eval(program, env)
	if evaluated != nil {
		if evaluated.Type() != object.NULL {
			fmt.Print(evaluated.String())
		}
	}
}
