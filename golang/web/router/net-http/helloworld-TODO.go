// basic hello world web application just using net/http package
package main

import (
	"fmt"
	"net/http"
)

type clientRequest int

func (c clientRequest) ServeHTTP(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintln(w, "Hi client request")
}

func main()  {
	var c1 clientRequest
	http.ListenAndServe(":8080",c1)
}
