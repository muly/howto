package main

import (
	"fmt"
	"os"

	"github.com/gocarina/gocsv"
)

type Contact struct {
	Name    string `csv:"name,omitempty"`
	Address string `csv:"address,omitempty"`
}

func main() {
	cs := []Contact{}
	cs = append(cs, Contact{Name: "name1", Address: `"address",1`}) // Note: double quotes and comma in the data
	cs = append(cs, Contact{Name: "name2", Address: "address,2"})

	err := writeCsv("contact-created.csv", cs)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func writeCsv(csvFile string, data interface{}) error {
	f, err := os.Create(csvFile)
	if err != nil {
		return err
	}
	defer f.Close()

	err = gocsv.MarshalFile(data, f)
	if err != nil {
		return err
	}

	return nil
}
