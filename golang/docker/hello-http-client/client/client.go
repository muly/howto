// write HTTP client to make a get request to an existing web service, example google.com search
// note: https://medium.com/@nate510/don-t-use-go-s-default-http-client-4804cb19f779

package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	log.Println("starting client...")
	serverURL := "http://127.0.0.1:8080"
	log.Println("server url:", serverURL)
	r, err := http.NewRequest("GET", serverURL, nil)
	if err != nil {
		log.Println("http.NewRequest error:", err)
		return
	}
	log.Println(r)

	client := http.Client{}
	resp, err := client.Do(r)
	if err != nil {
		log.Println("client.Do error:", err)
		return
	}
	log.Println(resp.Status)

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("ioutil.ReadAll error:", err)
		return
	}
	log.Println(string(b))
}