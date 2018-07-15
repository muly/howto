// demonstrate for loop with all the possible ways to use it
// print 0-10

package main

import (
	"fmt"
)

func main() {

	//
	fmt.Println("for declare/init; check; inc{}")
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}

	//
	fmt.Println("for ;check; inc{}")
	j := 0
	for ; j <= 10; j++ {
		fmt.Println(j)
	}

	//
	fmt.Println("for ; ; inc{}")
	k := 0
	for ; ; k++ {
		if k > 10 {
			break
		}
		fmt.Println(k)
	}

	//
	fmt.Println("break")
	l := 0
	for {
		if l > 10 {
			break // exits this (inner) for-loop and continue executing the next command after the for loop
		}
		fmt.Println(l)
		l++
	}

	//
	fmt.Println("continue")
	for m := 0; m <= 10; m++ {
		if m == 5 {
			continue // skips the rest of the code in the iteration and jumps to next iteration
		}
		fmt.Println(m)
	}

	//
	// fmt.Println("for{} :infinite loop") // run with caution, use ctrl+c to kill the program to stop
	// for {
	// 	fmt.Println("hello")
	// }

}
