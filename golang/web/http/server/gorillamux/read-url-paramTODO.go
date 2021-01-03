// example to demonstrate how to read parameter from the url. example id (1) from the below
// https://localhost:8080/customer/{id}
// https://localhost:8080/customer/1

//TODO: test the handler

package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	fmt.Fprintf(w, "received id = %v", id)
}
