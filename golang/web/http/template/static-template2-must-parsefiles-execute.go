//simple static template read from a file, parse and execute it
package main

import (
	"net/http"
	"text/template"
)

func main() {
	tmpl := template.Must(template.ParseFiles("stpl_single.gohtml"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err := tmpl.Execute(w, nil)
		if err != nil {
			panic(err)
		}
	})
	http.ListenAndServe(":8080", nil)
}
