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
	fmt.Println(reflect.DeepEqual(s1, s2)) // true

	s3 := st{a: 1, b: []string{"hello", "world"}}
	s4 := st{a: 1, b: []string{"hello", "world"}}
	fmt.Println(reflect.DeepEqual(s3, s4)) // true

	s5 := st{a: 1, b: []string{"hello", "world"}}
	s6 := st{a: 1, b: []string{"world", "hello"}}
	fmt.Println(reflect.DeepEqual(s5, s6)) // false

}
