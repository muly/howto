package main

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func helloworld(w http.ResponseWriter, r *http.Request) {
	msg := ""
	for i := 0; i < 10000; i++ {
		msg = msg + "Hello gzip middleware!\n"
	}
	w.Write([]byte(msg))
}
func main() {
	h := mux.NewRouter()
	h.HandleFunc("/", helloworld)
	log.Fatal(http.ListenAndServe(":8080", handlers.CompressHandler(h)))  // with compression middleware
	// log.Fatal(http.ListenAndServe(":8080", h)) // without compression middleware
}

////  manual testing
// run the application: using one of the methods below: go run, or docker run or docker-compose
	// go run gzip.go
	// or 
	// docker build -t gzip-image .
	// docker run --rm -p 8080:8080  gzip-image
	// or
	// docker build -t gzip-image .
	// docker-compose up
// run the api, see curl.txt file

////  integration testing
// go test -v -timeout 30s -run ^TestEndpointDockerfile$ github.com/muly/howto/middleware/compression/gorilla
// go test -v -timeout 30s -run ^TestEndpointDockerCompose$ github.com/muly/howto/middleware/compression/gorilla
