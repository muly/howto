// demonstrate how to read a csv file using gocsv library, skipping first n rows in the csv file

package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"os"

	"github.com/gocarina/gocsv"
)

type Contact struct {
	Id   int    `csv:"id"`
	Name string `csv:"name"`
}

func main() {
	trim := readcsv("contact2.csv", 3) //TODO: records to be ignored in the csv file should have the same number of fields as the actual data. otherwise this program is breaking
	//trim := readcsv("contact3.csv", 0)

	c := []Contact{}
	err := loadCsv2(trim, &c)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(c)
}

func loadCsv2(csvdata io.Reader, out interface{}) error {
	return gocsv.Unmarshal(csvdata, out)
}

func readcsv(csvFile string, skip int) io.Reader {
	f, err := os.Open(csvFile)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer f.Close()

	r := csv.NewReader(f)

	records := [][]string{}

	for i := 0; ; i++ {

		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
			return nil
		}
		if i < skip {
			continue
		}
		//fmt.Println(record)
		records = append(records, record)

	}
	//fmt.Println(records)

	ob := &bytes.Buffer{}

	out := csv.NewWriter(ob)

	err = out.WriteAll(records)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	//fmt.Println(ob)

	return ob

}
