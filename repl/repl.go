package repl

import (
	"bufio"
	"fmt"
	"io"
	"pilang/evaluator"
	"pilang/lexer"
	"pilang/object"
	"pilang/parser"
)

const PROMPT = ">> "

//func Start(in io.Reader, out io.Writer) {
//	scanner := bufio.NewScanner(in)
//	for {
//		fmt.Printf(PROMPT)
//		scanned := scanner.Scan()
//		if !scanned {
//			return
//		}
//		line := scanner.Text()
//		l := lexer.New(line)
//		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
//			fmt.Printf("%+v\n", tok)
//		}
//	}
//}

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment("")
	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		line := scanner.Text()
		l := lexer.New(line, "")
		p := parser.New(l)
		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}
		evaluated := evaluator.Eval(program, env)
		if evaluated != nil {
			if evaluated.Type() != object.NULL {
				io.WriteString(out, evaluated.String())
				io.WriteString(out, "\n")
			}
		}
	}
}

func printParserErrors(out io.Writer, errors []string) {
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
