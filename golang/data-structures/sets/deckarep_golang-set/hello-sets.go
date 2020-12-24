package main

import (
	"fmt"
	mapset "github.com/deckarep/golang-set"
)

func main() {

	s1 := mapset.NewSetFromSlice([]interface{}{"Delhi", "Gurgaon", "Noida", "Greater Noida", "Ghaziabad", "Manesar", "Faridabad",
		"Sonipat", "Meerut", "Chandigarh", "Ambala", "Bharatpur", "Mathura", "Rishikesh", "Dehradun",
		"Haridwar", "Mussoorie"})

	fmt.Println(s1.Contains("Delhi"))
	fmt.Println(s1.Contains("XYZ"))

}
