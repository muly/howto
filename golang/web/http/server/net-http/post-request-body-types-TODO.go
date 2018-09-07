//accept the request to server with data in body in various ways listed below:
//1a) form data: text
//1b) form data: file
//2) x-www-form-urlencoded
//3.1) raw data: plain text
//3.2) raw data: json
//4) binary
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func main() {

	//TODO: need to run and test these handlers
	http.HandleFunc("/formdatatext", handlerFormDataText)
	http.HandleFunc("/formdatafile", handlerFormDataFileType) // use form/multi form iption to upload a json file
	http.HandleFunc("/rawdatajson", rBodyJsonHandler) // paste the json content 
	http.HandleFunc("/binaryjson", rBodyJsonHandler) // attach the json file with the request
	http.ListenAndServe(":8080", nil)
}

// 1a) form data: text
func handlerFormDataText(w http.ResponseWriter, r *http.Request) {
// TODO
}

// 1b) form data: file (also as multipart file)
// NOTE: tested with client code "github.com\muly\howto\golang\web\http\client\net-http\post-request-body-types-TODO.go"
// But not tested with Postman or some other REST client.
func handlerFormDataFileType(w http.ResponseWriter, r *http.Request) {

	fmt.Println("received request")

	r.ParseMultipartForm(32 << 30)
	dataFile, _, err := r.FormFile("file") // Here file is the key specified in the form while sending the file in form
	if err != nil {
		fmt.Fprintln(w, "r.formfile error:"+err.Error())
		return
	}
	if dataFile == nil {
		fmt.Fprintln(w, "file NOT received")
		return
	}
	defer dataFile.Close()

	processFile(dataFile)

	fmt.Fprintln(w, "file received")
}

func processFile(f io.Reader) error {
	b, err := ioutil.ReadAll(f)
	if err != nil {
		return err
	}

	fmt.Println(string(b))
	return nil
}

//2) x-www-form-urlencoded

//3.1) raw data: plain text

//3.2) raw data: json
// example input json data: {"name":"golang"}
//4) binary
// example input json data in the file: {"name":"golang"}
func rBodyJsonHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	
	d := struct {
		Name string `json:"name"`
	}{}

	err := json.NewDecoder(r.Body).Decode(&d)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "name is %v", d.Name)
}

