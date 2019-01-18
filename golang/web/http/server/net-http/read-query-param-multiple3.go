// example to demonstrate reading query parameters (with multiple values in both the below formats) using net/http library. 

// 1) example input: http://localhost:8080/customer?country=usa,india
// required output: a slice of string with values usa and india like this []string{"usa","india"}

// 2) example input: http://localhost:8080/customer?country=usa&country=india
// required output: a slice of string with values usa and india like this []string{"usa","india"}
//


package main

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

func queryParamHandler(w http.ResponseWriter, r *http.Request) {
	vars, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	var country []string

	for _, c := range vars["country"] {
		country = append(country, strings.Split(c, ",")...)
	}

	fmt.Fprintln(w, country)
}

func main() {

	http.HandleFunc("/customer", queryParamHandler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
