// demonstrate method override:
// when child type composes a parent type
// then all parent methods are accesible by the child type.
// Now, we can now add a new method to child type
// with the same signature as of a method in the parent type,
// thereby overriding the functionality

package main

import (
	"fmt"
)

type mytype1 struct{}

func (m mytype1) Print() {
	fmt.Println("mytype1")
}

type mytype2 struct {
	mytype1 // composition
}

//func (m mytype2) Print() { // uncomment this function to over ride the parent type's method
//	fmt.Println("mytype2")
//}

func main() {
	a := mytype2{}
	a.Print()
}
