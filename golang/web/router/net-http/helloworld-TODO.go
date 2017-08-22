// basic hello world web application just using net/http package
package main

import (
	"net/http"
	"fmt"
)

func hellworld(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w,"Welcome to HelloWorld...!")
}

func main()  {
	http.HandleFunc("/",hellworld)
	http.ListenAndServe(":8080",nil)
}