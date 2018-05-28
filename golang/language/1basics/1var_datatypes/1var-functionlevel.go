// declaring a variable at function level

package main

import (
	"fmt"
)

func main() {

	// declare & initialize in seperate lines
	var i int
	i = 10

	// declare & initialize in same lines
	var j int = 10

	// data type inferred
	var k = 10

	// short hand notation
	l := 10

	fmt.Printf("%v %v %v %v: %T %T %T %T\n", i, j, k, l, i, j, k, l)
}
