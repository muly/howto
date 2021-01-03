// to demonstrate how to open a writer and save a file to writer instead of to disk

package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	filename := `/Users/srinivasamuly/go/src/github.com/muly/howto/golang/io/io-writer.go`

	f, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	write2Disk(f)

	write2Writer()
}

func write2Disk(f io.Reader) error {
	out, err := os.Create("temo.txt")
	if err != nil {
		return err
	}

}

func write2Writer() {}
