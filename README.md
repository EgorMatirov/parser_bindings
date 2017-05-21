# parser_bindings
PoC library for fetching Go AST using Go parser library from C++

It's a simple C++ application that links to Go library and calls Go code for creating AST from Go code.
It demonstrates amount of partial parsing available in Go and possibility to return C structs from Go code to communication between Go and C++ worlds.

Input given to parser:
```
package main; var b int = 3 
func main(){
    a := b

}
func test(){
if
}
 func test2(){}
```

Output:

```Parsing source...
test.go:8:1: expected operand, found '}'
test.go:9:16: expected ';', found 'EOF'
test.go:9:16: expected ';', found 'EOF'
test.go:9:16: expected '{', found 'EOF'
test.go:9:16: expected '}', found 'EOF'
test.go:9:16: expected '}', found 'EOF'
Package name: main
Global declarations count: 3
1. General declaration. Count: 1
2. Function main
3. Function test
```

Note that function test2 isn't listed in output due to error in parsing (no statement after 'if')
