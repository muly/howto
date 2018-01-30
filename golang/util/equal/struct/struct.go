package main

import (
	"fmt"
)

type st struct {
	a int
	b string
}

func main() {

	s1 := st{a: 1, b: "hello"}
	s2 := st{a: 1, b: "hello"}

	fmt.Println(s1 == s2)

}
