//TODO: simple static template read from a file, parse and execute it
package main

import (
	"net/http"
	"text/template"
)

var tmpl *template.Template

func main() {
	tmpl = template.Must(template.ParseFiles("stpl2.gohtml"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, nil)
	})
	http.ListenAndServe(":8080", nil)
}
