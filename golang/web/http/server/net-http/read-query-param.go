//example to demonstrate reading query parameters using net/http library. 
// example: https://localhost:8080/?country=usa

package main

import (
	"net/http"
	"net/url"
)

func queryParamHandler(w http.ResponseWriter, r *http.Request) {
	vars, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	country := ""
	if val, exists := vars["country"]; exists {
		country = val[0]
	}
	w.Write([]byte(country))
}

func main() {

	http.HandleFunc("/", queryParamHandler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
