package api

import (
	"net/http"

	log "github.com/sirupsen/logrus"
)

func Helloworld(w http.ResponseWriter, r *http.Request) {

	log.WithFields(log.Fields{"pkg": "api"}).Info("api request received")
	w.Write([]byte("Hello World!"))
}
