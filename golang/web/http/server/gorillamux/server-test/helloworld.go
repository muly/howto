// demonstrate how to test the api not at handler level but at server level
package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func helloworld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!!!!!"))
}
func main() {
	h := mux.NewRouter()
	h.HandleFunc("/", helloworld)
	http.ListenAndServe(":8080", h)
}
