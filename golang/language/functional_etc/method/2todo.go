//
//TODO: what new concept does this example demonstrate than the 1.go? accordingly delete this or add the notes
package main

import (
	"fmt"
)

type customer struct {
	fname, lname, email string
	billingAddress      address
	shippingAddress     address
}

type address struct {
	line1, street, city, state, country, zip string
}

func (a address) printAddress() {
	fmt.Println(a.line1 + " " + a.street) // TODO: make it formatted full address
}

func main() {
	c := customer{}

	c.billingAddress = address{line1: "billing address line 1"}

	fmt.Println(c.billingAddress.line1)
	c.billingAddress.printAddress()

}
