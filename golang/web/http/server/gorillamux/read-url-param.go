// example to demonstrate how to read parameter from the url. example id (1) from the below
// http://localhost:8080/customer/{id}
// http://localhost:8080/customer/1

package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	h := mux.NewRouter()
	h.HandleFunc("/customer/{id}", handler)
	http.ListenAndServe(":8080", h)
}

func handler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	fmt.Fprintf(w, "received 
	id = %v", id)
}
