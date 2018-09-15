// struct fields can have complex datatype
package main

import (
	"fmt"
)

type address struct {
	line1, line2              string
	city, state, zip, country string
}

type customer struct {
	fname, lname string
	emails       []string
	alsoEmails   map[string]string
	homeAddress  address
}

func main() {

	c := customer{}
	c.fname = "abc"
	c.lname = "xyz"
	c.emails = []string{"abc@email.com", "email2"}
	c.alsoEmails = map[string]string{"email": "abc@email.com"}
	c.homeAddress = address{city: "SFO", country: "usa"}

	fmt.Println(c)

	for _, v := range c.emails {
		fmt.Println(v)
	}

}
