// demonstrate usage of defult mux, by creating a simple web service
package main

import("net/http"
"log"
)

func main() {


  http.HandleFunc("/hello", helloHandler)

  log.Println("Listening...")
  http.ListenAndServe(":3000", nil)
}

func helloHandler(w http.ResponseWriter, r *http.Request){
w.Write([]byte("hello world"))
}
//TODO: format and test
