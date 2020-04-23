package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	go someComplexFunc()
	time.Sleep(20 * time.Second) // rudimentary way to wait for the go routine to finish
}

func someComplexFunc() {
	time.Sleep(10 * time.Second)
	fmt.Println("hello world")
}
