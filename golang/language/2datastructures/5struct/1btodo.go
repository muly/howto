// struct defined inside a func
package main

import (
	"fmt"
)

func main() {

	type customer struct {
		fname, lname, email string
	}

	var c1 customer // declare

	c1.fname = "abc" // initialize one field at a time
	c1.lname = "xyz"

	c1 = customer{fname: "abc", lname: "xyz"} // initialize whole struct

	fmt.Println(c1.fname) // access one field at a time
	fmt.Println(c1)       // access the complete struct

	var c2 customer = customer{fname: "abc", lname: "xyz"} // declare and init at the same time

	var c3 = customer{fname: "abc", lname: "xyz"} // declare and init at the same time

	c4 := customer{fname: "abc", lname: "xyz"} // declare and init at the same time

	fmt.Println(c2, c3, c4)
}
