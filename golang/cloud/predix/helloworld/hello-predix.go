//
//
package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello predix!"))
	})
	port := os.Getenv("PORT")
	fmt.Println("running on port:", port)
	http.ListenAndServe(":"+port, nil)
}
