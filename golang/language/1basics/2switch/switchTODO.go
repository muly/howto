// demonstrate switch case

package main

import "fmt"

func main() {

	marks := 34

	switch {
	case marks < 35:
		fmt.Println("fail")
	case marks < 60:
		fmt.Println("pass")
	case marks < 75:
		fmt.Println("pass with F")
	default:
		fmt.Println("pass with D")
	}



grade := "A"
	switch  {
	case grade == "A":
		fmt.Println("pass with A grade")
	case grade == "B":
		fmt.Println("pass with B grade")
	case grade == "C":
		fmt.Println("pass with C grade")
	default:
		fmt.Println("fail")
	}

	// cleaner version 
	switch grade {
	case "A":
		fmt.Println("pass with A grade")
	case "B":
		fmt.Println("pass with B grade")
	case "C":
		fmt.Println("pass with C grade")
	default:
		fmt.Println("fail")
	}

}


/*
switch {
case bool-exp1:
	// do something here
case bool-exp1:
	// do something here
default:
	// do something here
}

switch var{
case val1:
	// do something here
case val2:
	// do something here
default:
	// do something here
}


use fallthrough to fall through. but by default go doesn't fall through


*/