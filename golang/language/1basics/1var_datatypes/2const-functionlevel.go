// declaring a constant at function level

package main

import (
	"fmt"
)

func main() {

	// declare & initialize in seperate lines (not allowed)
	//const i int // ERROR: const declaration cannot have type without expression
	//i = 10

	// declare and initialize in the same line
	const j int = 10

	// data type inferred
	const k = 10

	// short hand notation (not allowed)
	// l:= 10 //this will be a variable, not a constant

	fmt.Printf("%v %v: %T %T\n", j, k, j, k)

	// modifying the value of a constant is not allowed
	// k=11 //ERROR: cannot assign to k

}
