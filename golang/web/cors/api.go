// cors example

package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*") // ******** VERY IMPORTANT ********
		w.Write([]byte(r.URL.Query().Get("name") + `, hello from api`))
	})
	http.ListenAndServe(":8080", nil)
}

/*
Note: without the Access-Control-Allow-Origin header in the response, the browsers will not alow other web applications to access the data returned by the api. it will result in an error something like below

Access to XMLHttpRequest at 'http://localhost:8080/?name=your_name' from origin 'null' has been blocked by CORS policy: No 'Access-Control-Allow-Origin' header is present on the requested resource.
*/

