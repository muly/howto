package main

import (
	"fmt"
	"reflect"
)

type st struct {
	a int
	b []string
}

func main() {

	s1 := st{a: 1, b: []string{"hello"}}
	s2 := st{a: 1, b: []string{"hello"}}

	//fmt.Println(s1 == s2) // invalid operation: s1 == s2 (struct containing []string cannot be compared)
	fmt.Println(reflect.DeepEqual(s1, s2))
}
