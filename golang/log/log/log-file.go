package main

import (
	"log"
	"os"
)

func main() {
	f, err := os.OpenFile("logfile", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()

	log.Println("This is a test log entry to console")
	log.SetOutput(f)
	log.Println("This is a test log entry to file")
}
