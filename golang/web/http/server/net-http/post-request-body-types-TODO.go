//accept the request to server with data in body in various ways listed below:
//1a) form data: text
//1b) form data: file
//2) x-www-form-urlencoded
//3.1) raw data: plain text
//3.2) raw data: json
//4) binary
//5) multipart file attachment?????
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func main() {

	http.HandleFunc("/rawdatajson", rawDataJsonHandler)
	http.HandleFunc("/formdatafile", handlerFormDataFileType)
	http.ListenAndServe(":8080", nil)
}

// 1a) form data: text

// 1b) form data: file (or multipart example?)
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
// input json data example: {"name":"golang"}
func rawDataJsonHandler(w http.ResponseWriter, r *http.Request) {
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

//4) binary

//5) multipart file attachment?????
