package main

import (
	"fmt"
)

func main() {
	var i int = 1       // int int8 int16 int32 int64
	var f float32 = 1.1 // float32 float64
	var s string = "hello"
	var b bool = true

	fmt.Printf("%d %f %s %v\n", i, f, s, b)
}

// References:
// basic types: https://tour.golang.org/basics/11
