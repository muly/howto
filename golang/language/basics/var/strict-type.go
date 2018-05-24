// go is strictly typed language
package main

import (
	"fmt"
)

func main() {
	var i int8 = 9
	var j int

	//j= i //ERROR: cannot use i (type int8) as type int in assignment
	j = int(i) //type cast

	fmt.Println(i, j)
}
