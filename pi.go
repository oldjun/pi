package main

import (
	"fmt"
	"os"
	"path/filepath"
	"pilang/evaluator"
	"pilang/lexer"
	"pilang/object"
	"pilang/parser"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Print("missing pi file")
		return
	}
	file := os.Args[1]
	input, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}
	directory, _ := os.Getwd()
	l := lexer.New(string(input), filepath.Join(directory, file))
	p := parser.New(l)
	program := p.ParseProgram()
	env := object.NewEnvironment(directory)
	evaluated := evaluator.Eval(program, env)
	if evaluated != nil {
		if evaluated.Type() != object.NULL {
			fmt.Print(evaluated.String())
		}
	}
}
