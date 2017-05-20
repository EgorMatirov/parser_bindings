package main

//#include "type.h"
import "C"

import "go/parser"
import "go/token"
import "go/scanner"
import "fmt"
import "os"

func main() {
}

//export ParseFile
func ParseFile(source string) *C.struct_File {
	fmt.Println("Parsing source...")
	f, err := parser.ParseFile(token.NewFileSet(), "test.go", source, parser.AllErrors)
	if err != nil {
		scanner.PrintError(os.Stdout, err)
	}
	return ConvertAstFile(f)
}
