//accept the request to server with data in body in various ways listed below:
//1) form data (types: text, file)
//2) x-www-form-urlencoded
//3) raw data: json or any plain text formats
//4) binary
//5) multipart file attachment?????
package main

import (
	"fmt"
	"net/http"
	"io"
	"bytes"
	"os"
)

func main() {
	http.HandleFunc("/formdatafile", handlerFormDataFileType)
	http.ListenAndServe(":8080", nil)
}

//1b) form data file type example
func handlerFormDataFileType(w http.ResponseWriter, r *http.Request) {
	var buf bytes.Buffer

	file,header,err:=r.FormFile("hello")
	if err!= nil {
		fmt.Fprintln(w,"cannot get the file",err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	name:= header.Filename
	_,err=io.Copy(&buf,file)
	if err!= nil {
		fmt.Fprintln(w,"cannot copy the file",err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	contents:=buf.String()
	fmt.Println(contents)
	_,err=os.Create(name)
	if err!= nil {
		fmt.Fprintln(w,"cannot create the file",err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	downfile,err:=os.OpenFile(name,os.O_WRONLY,os.ModeAppend)
	if err!= nil {
		fmt.Fprintln(w,"cannot open the file",err)

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_,err=downfile.WriteString(contents)
	if err!= nil {
		fmt.Fprintln(w,"cannot write into the file",err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()
	fmt.Fprintln(w, "file received")
}
