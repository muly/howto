// demonstrate creating a zip file from a static list of files available on the server

package main

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", zipit)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("error:", err)
	}
}

func zipit(w http.ResponseWriter, r *http.Request) {
	fname := "file.txt"
	f, err := os.Open(fname)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data, err := ioutil.ReadAll(f)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	buf, err := zipfile(fname, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Disposition", "attachment; filename=file.zip")
	io.Copy(w, buf)
	return
}

func zipfile(fname string, data []byte) (*bytes.Buffer, error) {
	var b []byte
	buf := bytes.NewBuffer(b)
	zwriter := zip.NewWriter(buf)

	fwriter, err := zwriter.Create(fname)
	if err != nil {
		return nil, err
	}
	_, err = fwriter.Write(data)
	if err != nil {
		return nil, err
	}
	err = zwriter.Close()
	if err != nil {
		return nil, err
	}

	return buf, nil

}
