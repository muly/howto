package main

import (
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", helloHandler)
	http.ListenAndServe(":8080", nil)

}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("request received")
	w.Write([]byte("hello world"))
}
