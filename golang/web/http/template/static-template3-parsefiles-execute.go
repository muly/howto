//simple static template read from a file, parse and execute it
package main

import (
	"net/http"
	"text/template"
)

func main() {
	tmpl, err := template.ParseFiles("stpl_single.gohtml")
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		err := tmpl.Execute(w, nil)
		if err != nil {
			panic(err)
		}
	})
	http.ListenAndServe(":8080", nil)
}
