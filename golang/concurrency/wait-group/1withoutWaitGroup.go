// demonstrate the problems because of lack of synching mechanism like wait groups

package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 10; i++ {
		go msg(i)
	}
}

func msg(i int) {
	fmt.Println(i)
}

// Output: all the numbers will not be printed
// because the go routines will not have a chance to execute before the main function ends
// you need to introduce some sort of delay at the end of the main function so that it waits for some time before exiting.
// this can be achieved using wait group as shown in next example "2withWaitGroup.go"
