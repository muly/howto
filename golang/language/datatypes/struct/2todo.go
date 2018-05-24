// struct fields can have complex datatype
//TODO: need to add more (self explanatory) notes here
package main

import (
	"fmt"
)

type customer struct {
	fname, lname string
	emails       []string
}

func main() {

	c := customer{}
	c.fname = "F"
	c.lname = "L"
	c.emails = []string{"email1", "email2"}

	fmt.Println(c)
	
	for _, v := range c.emails{
		fmt.Println(v)
	}

}

