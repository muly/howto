// slice: different ways to declare and initialize

package main

import (
	"fmt"
)

func main() {

	// 1. declare and initialize separately
	var myslice []int
	myslice = []int{100, 200, 300, 400}
	fmt.Println(myslice)

	// 2. single line init and declare
	var myslice2 []int = []int{100, 200, 300, 400}
	fmt.Println(myslice2)

	// 3. using short hand notation
	myslice3 := []int{100, 200, 300, 400}
	fmt.Println(myslice3)

	// 4. using make. see other example program

}
