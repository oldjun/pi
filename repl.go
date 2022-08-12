package main

import (
	"fmt"
	"os"
	"pilang/repl"
)

func main2() {
	fmt.Printf("Welcome to Pi Programming Language!\n")
	repl.Start(os.Stdin, os.Stdout)
}
