// demonstrate if with all the posible ways of using if

package main

import "fmt"

func main() {

	marks := 90

	// if
	if marks < 35 {
		fmt.Println("fail")
	} else if marks < 60 { // else if
		fmt.Println("pass")
	} else if marks < 75 {
		fmt.Println("pass with F")
	} else { // else
		fmt.Println("pass with D")
	}

	// nested if
	// if (cond1){
	// 	//do somethinig here
	// 	if (cond2){

	// 	}
	// }

	// if (cond1) && (cond2){

	// }
	fname := "F"
	lname := "L"
	if marks > 35 {
		fullname := fname + " " + lname
		if marks < 60 { // else if
			fmt.Println(fullname, ": pass")
		} else if marks < 75 {
			fmt.Println(fullname, ": pass with F")
		} else { // else
			fmt.Println(fullname, ": pass with D")
		}
	} else {
		fmt.Println("fail")
	}

}

/* notes:
{} are manditory
() are optional
{ should be in the same line as the if/else

Logical operators:
	|| logical or
	&& logical and
	! logical not

	Examples:
	Logical and example:
	if a==b && c==d {}

	Logical Not example
	If !(a==b){}
	If !(a==b && c==d){}


*/
