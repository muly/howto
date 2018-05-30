// demonstrate for loop with all the possible ways to use it
// print 0-10

package main

import (
	"fmt"
)

func main() {

	// default syntax
	// for declare/init; check; inc{}

	fmt.Println("i")
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}

	fmt.Println("j")
	j := 0
	for ; j <= 10; j++ {
		fmt.Println(j)
	}

	fmt.Println("k")
	k := 0
	for ; ; k++ {
		if k > 10 {
			break
		}
		fmt.Println(k)
	}

	fmt.Println("l")
	l := 0
	for {
		if l > 10 {
			break // exits this (inner) for-loop and continue executing the next command after the for loop
		}
		fmt.Println(l)
		l++
	}

	fmt.Println("m")
	for m:=0; m<=10; m++{
		if m == 5 {
			continue // skips the rest of the iteration and jumps to next iteration
		}
		fmt.Println(m)
	}

	// infinite loop, run with caution, use ctrl+c to kill the program to stop
	// for{
	// 	fmt.Println("hello")
	// }

}
