package main

import (
	"go/parser"
	"go/token"
	"testing"
)

func TestHasNotOSPackage(t *testing.T) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "test.go", "package main", parser.AllErrors)
	if err != nil {
		t.Error(err)
	}

	if hasOSPackage(f.Imports) {
		t.Fatal("failed")
	}
}

func TestHasOSPackage(t *testing.T) {
	source := `package main
	import "os"`

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "test.go", source, parser.AllErrors)
	if err != nil {
		t.Error(err)
	}

	if !hasOSPackage(f.Imports) {
		t.Fatal("failed")
	}
}
