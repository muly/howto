package main

import (
	"fmt"
)

func main() {

forOutside:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if j > 2 {
				break forOutside // will break to outer loop, otherwise "break" will only break from innermost for loop where that break is called
			}
			fmt.Println(i, j)
		}
	}
}
