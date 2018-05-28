// declaring a constant at package level

package main

import (
	"fmt"
)

// declare and initialize in the same line
const n int = 10

// data type inferred
const o = 10

func main() {
	fmt.Printf("%v %v: %T %T\n", n, o, n, o)
}
