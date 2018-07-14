// demonstrate if with all the posible ways of using if

package main

import "fmt"

func main() {

	var marks int = 90

	// if
	if marks > 0 {
		fmt.Println("your marks are registered")
	}

	// if else
	if marks < 35 {
		fmt.Println("fail")
	} else if marks < 60 { // else if
		fmt.Println("pass")
	} else if marks < 75 {
		fmt.Println("pass with F")
	} else { // else
		fmt.Println("pass with D")
	}

	// Nested if
	fname := "F"
	lname := "L"
	if marks > 35 {
		fullname := fname + " " + lname
		if marks < 60 { // nested if
			fmt.Println(fullname, ": pass")
		} else if marks < 75 {
			fmt.Println(fullname, ": pass with F")
		} else {
			fmt.Println(fullname, ": pass with D")
		}
	} else {
		fmt.Println("fail")
	}

	// Logical operators
	age := 20
	height := 140
	if age > 18 && height > 150 {
		fmt.Println("you are eligible")
	} else {
		fmt.Println("sorry, you are NOT eligible")
	}
	/*
		Logical operators:
			|| logical or
			&& logical and
			! logical not
		Examples:
		Logical or example:
			if a==b || c==d {}

		Logical and example:
			if a==b && c==d {}

		Logical Not example
			If !(a==b){}
			If !(a==b && c==d){}
	*/
}

/* important notes:
{} are manditory
() are optional
{ should be in the same line as the if/else
*/
