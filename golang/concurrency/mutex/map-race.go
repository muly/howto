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
var wg = sync.WaitGroup{}

func main() {

	m = map[string]int{}

	for i := 0; i < 10; i++ {

		//fmt.Println(i)
		wg.Add(1)
		go writeToMap(i)
	}

	wg.Wait()
	fmt.Println(m)

}

func writeToMap(i int) {
	m["A"] = i
	wg.Done()
}
