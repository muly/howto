// to demonstrate that maps are reference types

package main

import "fmt"

func modify(m map[string]int) {
	m["foo"] = 50
	m["bar"] = 20
}

func main() {
	fmt.Println("Hello, playground")

	original := make(map[string]int)

	original["foo"] = 5
	original["foobar"] = 100

	modify(original)

	fmt.Println(original) // map[foo:50 foobar:100 bar:20]
}
