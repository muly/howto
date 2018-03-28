package main

import (
	"fmt"
	"github.com/gocarina/gocsv"
	"os"
)

type Contact struct {
	Name string `csv:"name,omitempty"`
}

func main() {
	c := []Contact{}
	err := loadCsv("contact.csv", &c)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(c)
}

func loadCsv(csvFile string, out interface{}) error {
	f, err := os.Open(csvFile)
	if err != nil {
		return err
	}
	defer f.Close()

	return gocsv.UnmarshalFile(f, out)
}
