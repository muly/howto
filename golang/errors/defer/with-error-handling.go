// demonstrate error handling inside defered anonymous function

// example inspired from https://blevesearch.com/news/Deferred-Cleanup,-Checking-Errors,-and-Potential-Problems/

package main

import (
	"fmt"
	"log"
)

func main() {
	r, err := Open("a")
	if err != nil {
		log.Fatalf("error opening 'a'\n")
	}
	// defer r.Close()

	defer func() {
		err = r.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	
	
	err = fmt.Errorf("some random error")
	if err != nil {
		log.Printf("received error %v", err)
	}
}

type Resource struct {
	name string
}

func Open(name string) (*Resource, error) {
	return &Resource{name}, nil
}

func (r *Resource) Close() error {
	log.Printf("closing %s\n", r.name)
	return fmt.Errorf("failed to close")
}
