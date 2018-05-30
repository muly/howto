// Array basics:

package main

import (
	"fmt"
)

func main() {
	var myarray [4]int           // declare
	myarray = [4]int{1, 2, 3, 5} // initialize
	myarray[0] = 100             // set value in one index position

	// access each element using index, starting at 0 to len-1
	fmt.Println("myarray", myarray[0], myarray[1], myarray[2], myarray[3])

	// Note: cannot access index locations beyond the length specified in declaration
	//
	// myarray[4]= 9 // compiler error: invalid array index 4 (out of bounds for 4-element array)
	// fmt.Println(myarray[5]) // compiler error: invalid array index 5 (out of bounds for 4-element array)

	// len(), cap() are build in functions to get the length and capacity of the array.
	// In case of an array, these two will be the same
	fmt.Println("myarray", myarray, len(myarray), cap(myarray))

	fmt.Println("-------------")
	// looping through all the index positions
	for i := 0; i < len(myarray); i++ {
		fmt.Println(i, myarray[i])
	}

	fmt.Println("-------------")
	// looping using range
	for i, v := range myarray {
		fmt.Println(i, v)
	}

	// Note: array of one size is not compatible with array of another type
	var myarray2 [5]int
	//myarray2 = myarray //cannot use myarray (type [4]int) as type [5]int in assignment
	fmt.Println("myarray2", myarray2, len(myarray2), cap(myarray2))

	// declaring and initialiazing
	myarray3 := [5]int{10, 11, 12, 13, 14}
	fmt.Println("myarray3", myarray3, len(myarray3), cap(myarray3))

	// declaring and initialiazing without mentioning the size
	myarray4 := [...]int{10, 11, 12, 13, 14}
	fmt.Println("myarray4:", myarray4, len(myarray4), cap(myarray4))

	// declaring with a size but initialiazing with less elements.
	// Note: the other elements will be initialized with the zero value of the respective type
	myarray5 := [6]int{10, 11, 12, 13, 14}
	fmt.Println("myarray5", myarray5, len(myarray5), cap(myarray5))

	// declaring with a size but initialiazing with more elements.
	//
	// myarray6 := [3]int{10, 11, 12, 13, 14} //COMPILER ERROR: array index 3 out of bounds [0:3]
	// fmt.Println("myarray6",myarray6, len(myarray6), cap(myarray6))

}
