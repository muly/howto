// Array basics:

package main

import (
	"fmt"
)

func main() {

	var i int
	var j int

	i = 10
	j = 20
	i = j

	//Note that one variable of a type can be passed/assigned to different variable of same type
	fmt.Println(i, j) // 20 20

	// but that is not the same case with arrays
	// because array size is part of the array type
	var myarray [4]int           // declare
	myarray = [4]int{1, 2, 3, 5} // initialize

	// Note: array of one size is not compatible with array of another type
	var myarray2 [5]int
	//myarray2 = myarray //ERROR: cannot use myarray (type [4]int) as type [5]int in assignment
	fmt.Println("myarray2", myarray2, len(myarray2), cap(myarray2))

}
