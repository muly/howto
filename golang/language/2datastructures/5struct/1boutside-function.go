// struct defined outside a func
// scope: within the package
package main

import (
	"fmt"
)

type customer struct {
	fname, lname, email string
}

func main() {

	var c1 customer // declare

	c1.fname = "abc" // initialize one field at a time
	c1.lname = "xyz"

	c1 = customer{fname: "abc", lname: "xyz"} // initialize whole struct

	c1 = customer{"abc", "xyz", "abc@email.com"} // initialize whole struct without having to mention the field

	fmt.Println(c1.fname) // access one field at a time
	fmt.Println(c1)       // access the complete struct

	//
	var c2 customer = customer{fname: "abc", lname: "xyz"} // declare and init at the same time

	var c3 = customer{fname: "abc", lname: "xyz"} // declare and init at the same time

	c4 := customer{fname: "abc", lname: "xyz"} // declare and init at the same time

	fmt.Println(c2, c3, c4)
}
