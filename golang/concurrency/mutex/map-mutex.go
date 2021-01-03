// demonstrate how to use mutex, in this case to safegaurd a map from race conditions. run the program to see the result
// 		go run map-mutex.go
//
// run using -race flag like below. you will note that the there is no race condition, unlike the other example: map-race.go
// 		go run -race map-mutex.go

package main

import (
	"fmt"
	"sync"
)

var m map[string]int
var mutex sync.Mutex

var wg = sync.WaitGroup{} // wait group is used to wait at the end of the main process for all the goroutines to complete their work

func main() {
	m = map[string]int{}

	for i := 0; i < 10; i++ {

		fmt.Println(i)
		wg.Add(1)
		go writeToMap("A", i)
	}

	wg.Wait()
	fmt.Println(m)

}

func writeToMap(s string, i int) {
	mutex.Lock()
	m[s] = i
	mutex.Unlock()

	wg.Done()
}
