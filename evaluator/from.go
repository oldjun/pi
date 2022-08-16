package evaluator

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"pilang/ast"
	"pilang/lexer"
	"pilang/object"
	"pilang/parser"
)

var searchPaths []string
var imported map[string]*object.Environment

func init() {
	imported = make(map[string]*object.Environment)
}

func evalFrom(node *ast.From, env *object.Environment) object.Object {
	addSearchPath(env.GetDirectory())
	filename := findFile(node.File)
	if filename == "" {
		return newError("runtime error: no file found at '%s.pi'", node.File)
	}
	var scope *object.Environment
	if hasImported(filename) {
		scope = imported[filename]
	} else {
		scope = evaluateFile(filename, env)
		addImported(filename, scope)
	}
	return importFile(node, env, scope)
}

func addSearchPath(path string) {
	searchPaths = append(searchPaths, path)
}

func findFile(name string) string {
	basename := fmt.Sprintf("%s.pi", name)
	for _, path := range searchPaths {
		file := filepath.Join(path, basename)
		if fileExists(file) {
			return file
		}
	}
	return ""
}

func fileExists(file string) bool {
	_, err := os.Stat(file)
	return err == nil
}

func addImported(path string, env *object.Environment) {
	imported[path] = env
}

func hasImported(path string) bool {
	_, ok := imported[path]
	return ok
}

func evaluateFile(file string, env *object.Environment) *object.Environment {
	source, err := ioutil.ReadFile(file)
	if err != nil {
		return nil
	}
	l := lexer.New(string(source), file)
	p := parser.New(l)
	program := p.ParseProgram()
	scope := object.NewEnvironment(env.GetDirectory())
	result := Eval(program, scope)
	if isError(result) {
		return nil
	}
	return scope
}

func importFile(node *ast.From, env *object.Environment, scope *object.Environment) object.Object {
	if node.Everything {
		for alias, value := range scope.All() {
			env.Set(alias, value)
		}
		return NULL
	}
	for alias, identifier := range node.Identifiers {
		value, ok := scope.Get(identifier.Value)
		if !ok {
			return newError("runtime error: identifier '%s' not found in module '%s.pi'", identifier.Value, node.File)
		}
		env.Set(alias, value)
	}
	return NULL
}
