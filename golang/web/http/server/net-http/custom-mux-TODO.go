// demonstrate usage of custom mux, by creating a simple web service
package main

import("net/http"
"log"
)


func main() {
  mux := http.NewServeMux()

  mux.HandleFunc("/hello", helloHandler)

  log.Println("Listening...")
  http.ListenAndServe(":3000", mux)// notice that we are passing the mux object (2nd parameter)
}

func helloHandler(w http.ResponseWriter, r *http.Request){
w.Write([]byte("hello world"))
}
//TODO: format and test
