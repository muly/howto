// demonstrating that marshalling of []struct and []map returns the same json string.

package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type myStruct struct {
	Cola string `json:"cola"`
	Colb string `json:"colb"`
	Colc string `json:"colc"`
}

func main() {

	m := []myStruct{}
	s := []myStruct{}

	// convert the map to json and then to struct
	err := json.Unmarshal(mapjson(), &m)
	if err != nil {
		fmt.Println("m1:", err)
		return
	}

	// convert the struct to json and back to struct
	err = json.Unmarshal(structjson(), &s)
	if err != nil {
		fmt.Println("s1:", err)
		return
	}

	// compare 
	if reflect.DeepEqual(m, s) {
		fmt.Println("same", m, s) // this will be the output as they both result in the same json string
	} else {
		fmt.Println("not same", m, s)
	}

}

func structjson() []byte {
	var data []myStruct

	row1 := myStruct{
		Cola: "val-a1",
		Colb: "val-b1",
		Colc: "val-c1",
	}
	data = append(data, row1)

	row2 := myStruct{
		Cola: "val-a2",
		Colb: "val-b2",
		Colc: "val-c2",
	}
	data = append(data, row2)

	d, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return d
}

func mapjson() []byte {
	var data []map[string]string

	row1 := map[string]string{}
	row1["cola"] = "val-a1"
	row1["colb"] = "val-b1"
	row1["colc"] = "val-c1"
	data = append(data, row1)

	row2 := map[string]string{}
	row2["cola"] = "val-a2"
	row2["colb"] = "val-b2"
	row2["colc"] = "val-c2"
	data = append(data, row2)

	d, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return d
}
