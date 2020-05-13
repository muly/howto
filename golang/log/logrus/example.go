package main

import (
	"os"

	log "github.com/Sirupsen/logrus"
)

func main() {
	if err := SetLogfile("./sample.log"); err != nil {
        panic(err)
	}
	
	log.Println("This is a test log entry to console")
	log.Println("This is a test log entry to file")
}

func SetLogfile(path string) error {
	logfile, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	log.SetOutput(logfile)

	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})


	return nil
}
