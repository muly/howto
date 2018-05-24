// maps basics
// TODO: cleanup and break into multiple examples if required
package main

import (
	"fmt"
)

func main() {

	var myMap map[string]string //NOTE: where KeyType may be any type that is comparable (more on this later), and ValueType may be any type at all, including another map!

	// myMap["name"] = "go class" // ERROR: panic: assignment to entry in nil map
	//NOTE: make sure to init the map before writing to it

	//..init
	myMap = map[string]string{}
	//myMap = make(map[string]string)
	myMap["name"] = "go class"

	fmt.Println(myMap)

	fmt.Println(myMap["name"]) //to retrieve a specific key value pair

	val, exists := myMap["url"]
	if exists == false {
		fmt.Println("url is missing")
	} else {
		fmt.Println(val)
	}

	myMap["url"] = "class.com"

	for k, v := range myMap {
		fmt.Println(k + ":" + v)
	}

	// key is unique
	myMap["url"] = "myclass.com"
	fmt.Println("new url:", myMap["url"])

	delete(myMap, "url")
	fmt.Println("url:", myMap["url"]) //will be blank as we deleted

}
