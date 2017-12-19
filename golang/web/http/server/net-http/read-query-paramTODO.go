//example to demonstrate reading query parameters using net/http library. example: https://localhost:8080/customer?country=usa
//TODO: need to run and test
package main

import (
  "net/http"
)

func HandleActivitiesGet(w http.ResponseWriter, r *http.Request) {
  vars, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

  country := ""
	if val, exists := vars["country"]; exists {
		country = val[0]
	}
  fmt.Fprintln(w, country)
}
