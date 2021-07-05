package main


import (
	"fmt"

	"github.com/google/go-cmp/cmp"
)


type st struct {
	A int
	B string
}

func main() {

	s1 := st{A: 1, B: "hello"}
	s2 := st{A: 1, B: "hello"}

	fmt.Println(cmp.Equal(s1, s2))
}
