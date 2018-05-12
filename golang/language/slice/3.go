// slice: making a slice

package main

import (
	"fmt"
)

func main() {

	// make a slice. initializes to zero values
	// 1
	var myslice []int
	myslice = make([]int, 4, 6)
	fmt.Println(myslice)

	// initialize the slice as usual
	myslice[0] = 100
	myslice[1] = 200
	myslice[2] = 300
	myslice[3] = 400
	fmt.Println(myslice)

	// you can also declare & initialize in below ways
	// 2
	var myslice2 []int = make([]int, 4, 6)
	fmt.Println(myslice2)

	// 3
	var myslice3 = make([]int, 4, 6)
	fmt.Println(myslice3)

	// 4
	myslice4 := make([]int, 4, 6)
	fmt.Println(myslice4)
}
