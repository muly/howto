// demonstrate how to use mutex to access a shared resource from groutine

package main

import (
	"fmt"
	"sync"
	"time"
)

type counter struct {
	i int
	sync.Mutex
}

var c counter

func main() {
	for i := 0; i < 10; i++ {
		go do()
	}

	time.Sleep(5 * time.Millisecond)
	c.Lock()
	fmt.Println("do executed", c.i, "times")
	c.Unlock()
}

func do() {
	c.Lock()
	c.i++
	c.Unlock()

	time.Sleep(1 * time.Millisecond)
	//fmt.Println("done")
}

// run the code with --race flag as below to identify if there are any possible race conditions
// go run --race hello-mutex.go
