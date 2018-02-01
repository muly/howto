// write HTTP client to make a get request to an existing web service, example google.com search
// note: https://medium.com/@nate510/don-t-use-go-s-default-http-client-4804cb19f779

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {

	v := url.Values{}
	v.Add("q", "golang")
	reqURL := fmt.Sprintf("https://google.com/search?%s", v.Encode())

	r, _ := http.NewRequest("GET", reqURL, nil)
	fmt.Println(r)

	client := http.Client{}
	resp, _ := client.Do(r)
	fmt.Println(resp.Status)

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(b))
}
