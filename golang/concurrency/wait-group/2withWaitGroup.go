// demonstrate the use of wait groups

package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go msg(i)
	}
	wg.Wait()
}

func msg(i int) {
	fmt.Println(i)
	wg.Done()
}

// Output: all the numbers will be printed
// because, using wait group, we are waiting for all the go routines to complete execuing before the main function exit
