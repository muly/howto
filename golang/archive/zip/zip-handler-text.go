// demonstrate creating a zip file with a map of filename string to []bytes as input data

package main

import (
	"archive/zip"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", zipit)
	http.ListenAndServe(":8080", nil)
}

func zipit(w http.ResponseWriter, r *http.Request) {
	// Create a new zip archive
	writer := zip.NewWriter(w)

	// Add some files to the archive.
	files := map[string][]byte{}

	files["readme.txt"] = []byte("This archive contains some text files.")
	files["gopher.txt"] = []byte("Gopher names:\nGeorge\nGeoffrey\nGonzo")
	files["todo.txt"] = []byte("Get animal handling licence.\nWrite more examples.")

	for fileName, fileBody := range files {
		f, err := writer.Create(fileName)
		if err != nil {
			log.Fatal(err)
		}
		_, err = f.Write(fileBody)
		if err != nil {
			log.Fatal(err)
		}
	}
	// Make sure to check the error on Close.
	err := writer.Close()
	if err != nil {
		log.Fatal(err)
	}
}
