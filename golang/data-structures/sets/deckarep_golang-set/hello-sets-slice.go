// example of "sets" operation on slice using mapset library
// - remove duplicates
// - check for existance of a value

package main

import (
	"fmt"
	mapset "github.com/deckarep/golang-set"
)

func main() {
	// initialiazing with some data. can also initialize with empty
	s1 := mapset.NewSetFromSlice([]interface{}{"Delhi", "Gurgaon"})

	// adding more values
	s1.Add("Noida")
	// adding duplicate values
	s1.Add("Noida")

	// convert to slice
	fmt.Println(s1.ToSlice()) // duplicates will be removed

	// check for existance of a value
	fmt.Println(s1.Contains("Delhi"))
	fmt.Println(s1.Contains("XYZ"))
}
