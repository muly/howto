//example to demonstrate reading query parameters using net/http library. 
// example: http://localhost:8080/?country=usa&state=ca&zip=12345

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
	
	// method 3: no need to parse the form first
	phone:= r.FormValue("phone")
	
	// method 4:
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	} 
	zip := r.Form.Get("zip") // this gets the first value only (in case if there are multiple zip query params like http://localhost:8080/?country=usa&state=ca&zip=45678&zip=12345)

	w.Write([]byte(country+","+state+","+zip+","+phone))
}

func main() {

	http.HandleFunc("/", queryParamHandler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
