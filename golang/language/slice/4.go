// slice: slicing an array

package main

import (
	"fmt"
)

func main() {
	var myarray [4]int

	myarray = [4]int{0, 10, 20, 30}

	myslice := myarray[1:3] // myarray[i:j] i left boundary (inclusive). j is right boundary (excluding, i.e included till j-1)
	fmt.Println("myslice", myslice, len(myslice), cap(myslice))

	myslice1 := myarray[1:] // right default is len. myarray[1:len(myarray)] or myarray[1:4]
	fmt.Println("myslice1", myslice1, len(myslice1), cap(myslice1))

	myslice3 := myarray[:3] // left default is 0 index
	fmt.Println("myslice3", myslice3, len(myslice3), cap(myslice3))

}
