package main

import (
	"fmt"
	"os"
	"pilang/repl"
)

func main() {
	fmt.Printf("Welcome to Pi Programming Language!\n")
	repl.Start(os.Stdin, os.Stdout)
}
