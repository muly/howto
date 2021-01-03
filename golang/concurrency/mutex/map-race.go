// this program is very likely to panic because of race conditions.
// Note: maps are not threadsafe. see the "map-mutex.go" for the solution using mutex
// 		fatal error: concurrent map writes
//
// run with -race flag to detect the race condition. like below
// 		go run -race map-race.go
//

package main

import (
	"fmt"
	"sync"
)

var m map[string]int

var wg = sync.WaitGroup{} // wait group is used to wait at the end of the main process for all the goroutines to complete their work

func main() {

	m = map[string]int{}

	for i := 0; i < 10; i++ {

		//fmt.Println(i)
		wg.Add(1)
		go writeToMap("A", i)
	}

	wg.Wait()
	fmt.Println(m)

}

func writeToMap(s string, i int) {
	m[s] = i

	wg.Done()
}
