package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello")
	a := []int{1, 3, 4, 6, 7, 9, 10}
	b := []int{1, 2, 4, 5, 9, 10}
	fmt.Println(commonOmxn(a, b))

	//a1:=[]int{9,3,9,3,2,3}
	//fmt.Println(mostFrequent(a1))
}
