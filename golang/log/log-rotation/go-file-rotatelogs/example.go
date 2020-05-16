// log rotation example 

package main

import (
	"fmt"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/lestrrat/go-file-rotatelogs"
)

func main() {
	if err := setLogfile("./myapp.log"); err != nil {
		panic(err)
	}
	for i := 0; i <= 200; i++ {
		log.Println(i, "This is a test log entry to file")
		time.Sleep(2 * time.Second)
	}
}


func setLogfile(path string) error {
	logfile, err := rotatelogs.New(fmt.Sprintf("%s.generated.%s", path, "%Y-%m-%d.%H:%M:%S"),
		rotatelogs.WithLinkName("./current.log"),
		rotatelogs.WithMaxAge(time.Minute),
		rotatelogs.WithRotationTime(time.Second*60),
	)
	if err != nil {
		return err
	}
	log.SetOutput(logfile)

	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})

	return nil
}