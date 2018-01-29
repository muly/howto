package main

import (
	"encoding/json"
	"fmt"
)

type JsonSt struct {
	First string
	Last  string
}

func ConvJson(j string) JsonSt {
	data := JsonSt{}
	json.Unmarshal([]byte(j), &data)
	return data
}

func main() {
	jString := `{"First":"Jame","Last":"Bond"}`
	fmt.Println(ConvJson(jString))
}
