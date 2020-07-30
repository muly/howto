// example to demonstrate how to read query parameter example country=usa in below
// https://localhost:8080/customer?country=usa

package main

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/gorilla/mux"
)

func main() {
	h := mux.NewRouter()
	h.HandleFunc("/customer", handler)
	http.ListenAndServe(":8080", h)
}

func handler(w http.ResponseWriter, r *http.Request) {
	// vars := mux.Vars(r) // Note: mux.Vars wont work. looks like this only applies to url params, but not query params. for query params, use the the standard library method url.ParseQuery() as below
	vars, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	country := ""
	if _, ok := vars["country"]; ok {
		country = vars["country"][0]
	}

	fmt.Fprintf(w, "received country = %v", country)
}
