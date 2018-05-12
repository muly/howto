// slice basics

package main

import (
	"fmt"
)

func main() {
	// declare and initialize
	var myslice []int
	myslice = []int{100, 200, 300, 400}

	fmt.Println(myslice[2]) // accessing the value in a specific index position
	// Note: accessing the index out side the length will result in run time error
	// fmt.Println(myslice[4]) //RUNTIME ERROR: panic: runtime error: index out of range
	// myslice[4]=400 //RUNTIME ERROR: panic: runtime error: index out of range

	// len(), cap() are build in functions to get the length and capacity of the slice.
	fmt.Println("myslice", myslice, len(myslice), cap(myslice))

	fmt.Println("-------------")
	// loping through all the index positions
	for i := 0; i < len(myslice); i++ {
		fmt.Println(i, myslice[i])
	}

	fmt.Println("-------------")
	// looping using range
	for i, v := range myslice {
		fmt.Println(i, v)
	}

}
