//
//TODO: need to add notes and format. may be change the scenrio if required

package main

import (
	"fmt"
)

type customer struct {
	fname, lname, email string
}

func (s customer) fullname() string {
	return fmt.Sprintf("%s %s", s.fname, s.lname)
}

func (s customer) sendEmail() {
	fmt.Println("email sent sucessfully")
}

func main() {

	c := customer{fname: "F", lname: "L", email: "E"}

	fmt.Println(c.fullname())

	c.sendEmail()

}
