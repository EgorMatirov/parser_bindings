package main;

//#include "type.h"
import "C"

import "go/parser"
import "go/token"
import "fmt"

func main(){
}

//export ParseFile
func ParseFile(source string) *C.struct_File{
        fmt.Println("Parsing source...")
        f, _ := parser.ParseFile(token.NewFileSet(), "test.go", source, parser.AllErrors)
        return C.CreateFile(C.CString(f.Name.Name), nil)
}
