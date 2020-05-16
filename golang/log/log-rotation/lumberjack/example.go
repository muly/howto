// log rotation example

package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/natefinch/lumberjack"
)

var errLog *log.Logger

func main() {
	e, err := os.OpenFile("./sample.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("error opening file: %v", err)
		os.Exit(1)
	}
	errLog = log.New(e, "", log.Ldate|log.Ltime)
	errLog.SetOutput(&lumberjack.Logger{
		Filename:   "./sample.log",
		MaxSize:    1,  // megabytes after which new file is created retaining old ones
		MaxBackups: 3,  // number of backups. comment out for unlimited files
		MaxAge:     28, // in days : this deletes older logs, comment out for unlimited age
	})

	count := 100000
	for i := 0; i < count; i++ {
		errLog.Println(i, "test log entry to file")
		time.Sleep(1 * time.Millisecond)
	}
}
