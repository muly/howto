// demonstrate how to write a simple webserver using github.com/julienschmidt/httprouter router (see hello-web.go) along with handler unit test (see hello-web_test.go)

package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "hello world!\n")
}

func main() {
	router := httprouter.New()
	router.GET("/hello", helloHandler)

	log.Fatal(http.ListenAndServe(":8080", router))
}
