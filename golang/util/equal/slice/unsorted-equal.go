// given 2 slices which are not sorted, find if they both have same elements (need not be in the same order)

//TODO: below logic need not be the most performant one. it works good for small size of the slices
package main

import "fmt"

func main() {
	a := []string{"a", "b", "c"}
	b := []string{"a", "b", "c"} // same data case
	c := []string{"c", "b", "a"} // same data different order case
	d := []string{"b", "c", "d"} // different data, same length case
	e := []string{"b", "c"}      // different data, different lenght case
	f := []string{}              // empty case

	a1 := []string{"a", "a", "b"}
	a2 := []string{"a", "a", "b"} // same data case with duplicates
	a3 := []string{"a", "b", "a"} //
	g := []string{"a", "b", "b"}
	h := []string{"a", "b"}

	a4 := []string{}
	a5 := []string{} // both empty case

	fmt.Println(sliceStringComp(a, b))   // true
	fmt.Println(sliceStringComp(a, c))   // true
	fmt.Println(sliceStringComp(a, d))   // false
	fmt.Println(sliceStringComp(a, e))   // false
	fmt.Println(sliceStringComp(a, f))   // false
	fmt.Println(sliceStringComp(a1, g))  // false
	fmt.Println(sliceStringComp(a1, h))  // false
	fmt.Println(sliceStringComp(a1, a2)) // true
	fmt.Println(sliceStringComp(a1, a3)) // true
	fmt.Println(sliceStringComp(a4, a5)) // true
}

func sliceStringComp(aa []string, bb []string) bool {
	if len(aa) != len(bb) {
		return false
	}

	ma := sliceString2map(aa)
	mb := sliceString2map(bb)

	for a, acnt := range ma {
		if bcnt, exists := mb[a]; !exists || acnt != bcnt {
			return false
		}

	}
	return true
}

func sliceString2map(aa []string) map[string]int {
	m := map[string]int{}
	for _, a := range aa {
		m[a]++
	}
	return m
}
