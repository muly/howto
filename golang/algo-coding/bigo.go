package main

import (
	"fmt"
	"math"
)

// given an array, find the first element

/*
a[1,2,3]
b[1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18]


O(1)
func first(a []int)int{
	return a[0]
}

O(n)
func has1(a []int)bool{
	for _, v:= range a{
		if v==1{
			return true
			}
	}
	return false
}
*/

func main() {
	math.Floor()
	// no duplicates
	// sorted array
	aa := []int{1, 7, 9, 10, 14}
	bb := []int{1, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}

	fmt.Println(commonElements_Omxn(aa, bb))
	fmt.Println(commonElements_Ominmn(aa, bb))
}

func commonElements_Omxn(aa []int, bb []int) []int {
	bigOcnt := 0
	var out []int
	for _, a := range aa {
		for _, b := range bb {
			bigOcnt++
			if a == b {
				out = append(out, a)
			}
		}
	}

	fmt.Println("commonElements_Omxn cnt =", bigOcnt)
	return out
}

func commonElements_Ominmn(aa []int, bb []int) []int {
	bigOcnt := 0
	var out []int

	ia, ib := 0, 0

	for ia < len(aa) && ib < len(bb) {

		bigOcnt++
		if aa[ia] == bb[ib] {
			out = append(out, aa[ia])
			ia++
			ib++
		} else if aa[ia] < bb[ib] {
			ia++
		} else if bb[ib] < aa[ia] {
			ib++
		}

	}

	fmt.Println("commonElements_Omxn cnt =", bigOcnt)
	return out
}
