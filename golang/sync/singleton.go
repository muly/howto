// to demonstrate singleton pattern

package main

import (
	"fmt"
	"sync"
)

type singleton struct {
	Id string
}

var instance *singleton
var once sync.Once

func GetInstance() *singleton {
	once.Do(func() {
		instance = &singleton{}
	})
	return instance
}

func main() {
	i := GetInstance()
	j := GetInstance()
	i.Id = "i"
	j.Id = "j"
	fmt.Println(*i) //{j}
	fmt.Println(*j) //{j}

	p := &singleton{}
	q := &singleton{}
	p.Id = "p"
	q.Id = "q"
	fmt.Println(*p) //{p}
	fmt.Println(*q) //{q}
}
