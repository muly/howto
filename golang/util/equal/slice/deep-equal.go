package main

import (
	"fmt"
	"reflect"
)

func main() {
	a := []string{"a", "b", "c"}
	b := []string{"a", "b", "c"}
	c := []string{"c", "b", "a"}
	d := []string{"b", "c", "d"}
	e := []string{"b", "c"}

	fmt.Println(reflect.DeepEqual(a, b)) // true
	fmt.Println(reflect.DeepEqual(a, c)) // false
	fmt.Println(reflect.DeepEqual(a, d)) // false
	fmt.Println(reflect.DeepEqual(a, e)) // false
}
