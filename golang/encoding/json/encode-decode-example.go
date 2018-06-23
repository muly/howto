// example program to demonstrate the use of Encode, Decode functionality of `encoding/json` package
//// encode: from struct
//////	to stdout
//////	to json file
//////	to http ResponseWriter as json
//// decode: to struct
//////	from a json string
//////	from a json file
//////	from json data in http request

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type dog struct {
	Color string `json:"color"`
	Breed string `json:"breed"`
	Age   int    `json:"age"`
}

func main() {
	encodeToStdoutExample()
	encodeToFileExample()
	//encodeHandlerExample(w, r)

	decodeStringExample(`{"color":"brown","breed":"German Shepherd","age":5}`)
	decodeFileExample("data.json")
	//decodeHandlerExample(w, r)
}

func encodeToStdoutExample() {
	d := dog{Color: "brown", Breed: "German Shepherd", Age: 5}
	encode(d, os.Stdout)
}

func encodeToFileExample() {
	d := dog{Color: "brown", Breed: "German Shepherd", Age: 5}
	fileOut, _ := os.OpenFile("data.json", os.O_RDWR|os.O_CREATE, 0755)
	encode(d, fileOut)
}
func encodeHandlerExample(w http.ResponseWriter, r *http.Request) {
	d := dog{Color: "brown", Breed: "German Shepherd", Age: 5}
	encode(d, w)
}
func encode(i dog, o io.Writer) {
	err := json.NewEncoder(o).Encode(i) // to send as response http handler, use w, instead of Stdout
	if err != nil {
		fmt.Println(err)
	}
}

func decodeStringExample(str string) {
	stringInput := bytes.NewBuffer([]byte(str))
	fmt.Println(decode(stringInput)) // Buffer type
}
func decodeFileExample(fileName string) {
	fileInput, _ := os.Open(fileName)
	fmt.Println(decode(fileInput)) // File type
}
func decodeHandlerExample(w http.ResponseWriter, r *http.Request) {
	fmt.Println(decode(r.Body)) //io.ReadCloser type
}
func decode(i io.Reader) dog {
	output := dog{}
	err := json.NewDecoder(i).Decode(&output)
	if err != nil {
		fmt.Println(err)
	}
	return output
}
