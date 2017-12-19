/*send the client request with data in body. give examples of sending data in differnt formats listed below:
//1) form data (2 types: (1a) text, (1b) file)
2) x-www-form-urlencoded
3) raw data: json or any plain text formats
4) binary
5) multipart file attachment????? */

package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
)

func main() {

	filename := `C:\gows\src\github.com\muly\howto\golang\web\http\client\net-http\post-request-body-types-TODO.go`
	url := "http://localhost:8080/formdatafile"
	err := sendFileAsMultipartFileAttachment(filename, url)
	if err != nil {
		fmt.Println(err)
		return
	}
}

// 1b: form data file type client example
func sendFileAsMultipartFileAttachment(filename string, serverUrl string) error {

	body, contentType, err := GetMultipartFormData(filename)
	if err != nil {
		fmt.Println("GetMultipartFormData error:", err)
		return err
	}

	resp, err := http.Post(serverUrl, contentType, body)
	if err != nil {
		fmt.Println("cannot post", err)
		return err
	}

	defer resp.Body.Close()
	byte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("cannot read the responce", err)
		return err
	}
	fmt.Println(string(byte))

	return nil
}

func GetMultipartFormData(filename string) (data io.Reader, contentType string, err error) {

	body := &bytes.Buffer{}
	mw := multipart.NewWriter(body)

	w, err := mw.CreateFormFile("file", filename)
	if err != nil {
		return nil, "", err
	}

	f, err := os.Open(filename)
	if err != nil {
		return nil, "", err
	}
	_, err = io.Copy(w, f)
	if err != nil {
		return nil, "", err
	}

	err = mw.Close()
	if err != nil {
		return nil, "", err
	}

	return body, mw.FormDataContentType(), nil
}
