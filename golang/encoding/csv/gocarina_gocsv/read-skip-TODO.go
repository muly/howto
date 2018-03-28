// demonstrate how to read a csv file using gocsv library, skipping first n rows in the csv file
package main

import (
	"fmt"
	"os"
	//"bufio"

	"github.com/gocarina/gocsv"
)

type Contact struct {
	Name string `csv:"name,omitempty"`
}

func main() {
	c := []Contact{}
	err := loadCsv("contacts2.csv", &c, 4)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(c)
}

func loadCsv(csvFile string, out interface{}, headerAt int) error {

	f, err := os.Open(csvFile)
	if err != nil {
		return err
	}
	defer f.Close()

	//s:= bufio.NewScanner(f)

	return gocsv.UnmarshalFile(f, out)

}
