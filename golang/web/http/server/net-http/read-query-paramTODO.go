//example to demonstrate reading query parameters using net/http library. 
// example: http://localhost:8080/?country=usa&state=ca

package main

import (
	"net/http"
	"net/url"
)

func queryParamHandler(w http.ResponseWriter, r *http.Request) {
	// method 1: 
	vars, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	country := ""
	if val, exists := vars["country"]; exists {
		country = val[0]
	}
	
	// method 2: 
	vars2 := r.URL.Query()
	state := ""
	if val, exists := vars2["state"]; exists {
		state = val[0]
	}
	
	//TODO: any other methods? from other libraries, which support without having to worry about index (like vars[0])
	
	w.Write([]byte(country+","+state))
}

func main() {

	http.HandleFunc("/", queryParamHandler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
