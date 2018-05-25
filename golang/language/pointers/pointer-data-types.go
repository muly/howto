// pointer data types: demonstrate the nature of different data types: with respect to pass by value vs pass by reference

package main

import "fmt"

func main() {
	v := 10
	incv(v)
	fmt.Println("value outside func", v)
	incp(&v)
	fmt.Println("pointer value outside func", v)

	va := [1]int{10}
	fArray(va)
	fmt.Println("array outside func", va)

	vs := []int{10}
	fSlice(vs)
	fmt.Println("slice outside func", vs)

	vm := map[int]int{0: 10}
	fMap(vm)
	fmt.Println("map outside func", vm)

	vstr := "10"
	fString(vstr)
	fmt.Println("string outside func", vstr)

}

func incv(a int) { // passed by value
	a = 99
	fmt.Println("value inside func", a)
}

func incp(a *int) { // passed by reference
	*a = 99
	fmt.Println("pointer value inside func", *a)
}

func fArray(a [1]int) { // arrays are passed by value by default
	a[0] = 99
	fmt.Println("array inside func", a)
}

func fSlice(a []int) { // slices are passed by reference by default
	a[0] = 99
	fmt.Println("slice inside func", a)
}

func fMap(a map[int]int) { // maps are passed by reference by default
	a[0] = 99
	fmt.Println("map inside func", a)
}

func fString(a string) {
	a = "99"
	fmt.Println("string inside func", a)
}
