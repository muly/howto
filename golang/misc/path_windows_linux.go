package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func main() {
	var path string = `./path_windows_linux.go`

	read(path)

}

func read(path string) {
	// open file
	file, err := os.Open(path)
	if err != nil {
		log.Fatalln("error opening file\n", err)
	}
	defer file.Close()

	// read output
	rd, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatalln("error reading file\n", err)
	}

	// print file
	fmt.Println(string(rd))
}
