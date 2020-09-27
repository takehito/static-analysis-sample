package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
)

func hasOSPackage(importSpecs []*ast.ImportSpec) bool {
	for _, s := range importSpecs {
		if s.Path.Value == "\"os\"" {
			return true
		}
	}

	return false
}

func isOsGetenv(e *ast.CallExpr) bool {
	fun, ok := e.Fun.(*ast.SelectorExpr)
	if !ok {
		return false
	}

	x, ok := fun.X.(*ast.Ident)
	if !ok {
		return false
	}

	if x.Name != "os" || fun.Sel.Name != "Getenv" {
		return false
	}

	return true
}

func inspect() (func(ast.Node) bool, *[]ast.Expr) {
	var exp []ast.Expr
	return func(n ast.Node) bool {
		x, ok := n.(*ast.CallExpr)
		if !ok {
			return true
		}

		if isOsGetenv(x) {
			exp = append(exp, x.Args[0])
		}

		return true
	}, &exp
}

func getValue(e ast.Expr) string {
	var arg string
	switch x := e.(type) {
	case *ast.BasicLit:
		arg = x.Value
	case *ast.Ident:
		arg = "Ident"
	default:
		arg = "unknown"
	}

	return arg
}

func main() {
	file_name := os.Args[1]

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, file_name, nil, parser.AllErrors)
	if err != nil {
		log.Fatalln(err)
	}

	if !hasOSPackage(f.Imports) {
		fmt.Println("no os package")
		os.Exit(0)
	}

	inspectFunc, exp := inspect()

	ast.Inspect(f, inspectFunc)

	ast.Print(fset, exp)

	for _, v := range *exp {
		fmt.Println(getValue(v))
	}
}
