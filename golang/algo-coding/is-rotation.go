package main

import "fmt"

//Assumptions:
// no duplicates

func main() {

	a := []int{1, 2, 3, 4, 5, 6, 7}
	b := []int{4, 5, 6, 7, 1, 2, 3}

	fmt.Println(isRotation(a, b))

}
func isRotation(aa []int, bb []int) bool {
	if len(aa) != len(bb) {
		return false
	}
	isrotation := false
	ai, bi := 0, 0
	for ai < len(aa) {
		if aa[ai] == bb[bi] {
			isrotation = true
			ai++
			bi++
		} else {
			isrotation = false
			bi++
			if bi == len(bb) {
				bi = 0
			}
		}
	}

	return isrotation
}
