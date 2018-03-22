// demonstrate a simple webapp with docker

// https://medium.com/travis-on-docker/how-to-dockerize-your-go-golang-app-542af15c27a2

package main

import (
	"net/http"
)

func main() {

	http.HandleFunc("/", helloHandler)
	http.ListenAndServe(":8080", nil)

}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
}

// docker build to generate image, and tag it with the specified tag (using -t flag in the below example)
// 		docker build -t hellowebapp .
// docker run to create the container from the specified image (using tag in below example).
// and -p flag to indicate the outside port to insdide docker port mapping
// 		docker run --rm -p 9000:8080 hellowebapp
// in browser, run the below url to verify
//		http://localhost:9000/
