package main

import (
	"fmt"
	"github.com/oldjun/pi/evaluator"
	"github.com/oldjun/pi/lexer"
	"github.com/oldjun/pi/object"
	"github.com/oldjun/pi/parser"
	"github.com/oldjun/pi/repl"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Print("missing pi file")
		return
	}
	switch os.Args[1] {
	case "repl":
		__repl()
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

func __repl() {
	fmt.Printf("Welcome to Pi Programming Language!\n")
	repl.Start(os.Stdin, os.Stdout)
	os.Exit(0)
}
