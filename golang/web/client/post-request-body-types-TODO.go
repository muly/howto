/*send the client request with data in body. give examples of sending data in differnt formats listed below:
1) form data (types: text, file)
2) x-www-form-urlencoded
3) raw data: json or any plain text formats
4) binary
5) multipart file attachment????? */

package main

import (
	"net/http"

	"bytes"
	"mime/multipart"
	"os"
	"fmt"
	"io"
	"io/ioutil"
)

func main(){

filename := "hello.txt"




	body, contentType, err := GetMultipartFormData(filename)

	url := "http://localhost:8080/formdatafile"

	resp, err:= http.Post(url,contentType, body)
	if err != nil{
		fmt.Println("cannot post",err)
return
	}


	defer resp.Body.Close()
	byte,err:=ioutil.ReadAll(resp.Body)
	if err!= nil{
		fmt.Println("cannot read the responce",err)
		return
	}
	fmt.Println(string(byte))



}

func GetMultipartFormData(filename string)(data io.Reader, contentType string, err error){

	body := &bytes.Buffer{}
	mw:=multipart.NewWriter(body)

	w, err := mw.CreateFormFile("hello", filename)
	if err != nil{
return nil, "", err
	}

	f, err := os.Open(filename)
	if err != nil{
return nil, "", err
	}
	_, err= io.Copy(w, f)
	if err != nil{
return nil,"", err
	}

	err = mw.Close()
	if err != nil{
return nil, "", err
	}

	return body, mw.FormDataContentType(), nil
}
