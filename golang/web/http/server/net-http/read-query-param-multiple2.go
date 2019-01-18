// example to demonstrate reading query parameters (with multiple values in the below format) using net/http library.
// example input: http://localhost:8080/customer?country=usa&country=india
// required output: a slice of string with values usa and india like this []string{"usa","india"}
//

package main

import (
	"fmt"
	"net/http"
	"net/url"
)

func queryParamHandler(w http.ResponseWriter, r *http.Request) {
	vars, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	var country []string
	country = vars["country"]

	fmt.Fprintln(w, country)
}

func main() {

	http.HandleFunc("/customer", queryParamHandler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
