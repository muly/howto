//demonstrate example of http test using the http server code and the handler test (in _test.go)
package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/hello", handler)

	http.ListenAndServe(":8085", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}
