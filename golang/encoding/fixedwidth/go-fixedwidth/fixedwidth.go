package main

import (
	"fmt"
	"io/ioutil"
	"os"

	fixedwidth "github.com/ianlopshire/go-fixedwidth"
)

type people struct {
	ID        int     `fixed:"1,5"`
	FirstName string  `fixed:"6,15"`
	LastName  string  `fixed:"16,25"`
	Grade     float64 `fixed:"26,30"`
}

func main() {

	f, _ := os.Open("sample.dat") //TODO: handle error

	b, _ := ioutil.ReadAll(f) //TODO: handle error

	var p []people
	err := fixedwidth.Unmarshal(b, &p)
	if err != nil {
		fmt.Println(err)
	}

	for _, r := range p {
		fmt.Println(r)
	}
}
