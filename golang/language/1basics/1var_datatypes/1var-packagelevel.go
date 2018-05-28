// declaring a variable at package level

package main

import (
	"fmt"
)

// declare & initialize in seperate lines (initialized in main func)
var m int

// declare & initialize in same lines
var n int = 10

// data type inferred
var o = 10

// short hand notation (not allowed)
//p:= 10 //ERROR: syntax error: non-declaration statement outside function body
//Note := shorthand notation is not allowed outside a function

func main() {
	m = 10

	fmt.Printf("%v %v %v: %T %T %T\n", m, n, o, m, n, o)
}
