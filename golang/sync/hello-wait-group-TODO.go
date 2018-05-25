// Demonstrate the usage of wait groups to syncronize the go routines
// so that the main function waits for all the go routines to be executed completely before the main function exits.

package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go do()
		wg.Wait()
		fmt.Println("Hello go")
	}

}
func do() {
	fmt.Println("doing something")
	time.Sleep(1 * time.Second)
	fmt.Println("done something")
	defer wg.Done()
}
