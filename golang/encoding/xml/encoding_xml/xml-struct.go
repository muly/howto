// demonstrate the following:
// xml to go struct
// go struct to xml

package main

import (
	"encoding/xml"
	"fmt"
)

type person struct {
	Name string
}

func main() {
	p := person{Name: "abc"}
	x := structToXML(p)
	fmt.Println(x)
	p = xmlToStruct(x)
	fmt.Println(p)
}
func structToXML(p person) string {
	b, err := xml.Marshal(p)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return string(b)
}

func xmlToStruct(x string) person {
	p := person{}
	err := xml.Unmarshal([]byte(x), &p)
	if err != nil {
		fmt.Println(err)
		return person{}
	}
	return p
}
