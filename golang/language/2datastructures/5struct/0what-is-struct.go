// what is a struct?
// grouping related variables (of same or different types) into entity
// struct is a user defined type

package main

import (
	"fmt"
)

func main() {

	// using variables to create related fields

	var StudentName string
	var StudentEmail string
	var StudentPhone string
	var StudentAddress string
	var StudentAge int

	StudentName = "ABC"
	StudentEmail = "abc@xyz.com"
	StudentPhone = "12345"
	StudentAddress = "St 1"
	StudentAge = 20

	fmt.Println(StudentName, StudentEmail, StudentPhone, StudentAddress, StudentAge)

	////////////////////// simplifying this using struct

	type Student struct { // group all the related fields into a struct
		Name    string
		Email   string
		Phone   string
		Address string
		Age     int
	}

	var student Student // create a variable for that struct user defined type

	student.Name = "ABC" // initialize the fields. you can also initialize them together in a single line, for later example
	student.Email = "abc@xyz.com"
	student.Phone = "12345"
	student.Address = "St 1"
	student.Age = 20

	fmt.Println(student.Name, student.Email, student.Phone, student.Address, student.Age) // you can access and print the values of each field
	fmt.Println(student)                                                                  // or print the complete struct

}
