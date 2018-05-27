// demonstrate for loop with all the possible ways to use it
// print 0-10

package main

import (
	"fmt"
)

func main() {

	// default syntax
	// for declare/init; check; inc{}

	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}

	j := 0
	for ; j <= 10; j++ {
		fmt.Println(j)
	}

	k := 0
	for ; ; k++ {
		if k > 10 {
			break
		}
		fmt.Println(k)
	}

	l := 0
	for {
		if l > 10 {
			break
		}
		fmt.Println(l)
		l++
	}

	// infinite loop, run with caution, use ctrl+c to kill the program to stop
	// for{
	// 	fmt.Println("hello")
	// }

}
