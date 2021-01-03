// demonstrate the usage of multiple template files and serve a single page
package main

import (
	"net/http"
	"text/template"
)

func main() {
	tmpl := template.Must(template.ParseFiles("stpl_compose.gohtml", "stpl_body.gohtml", "stpl_header.gohtml"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		err := tmpl.Execute(w, nil)
		if err != nil {
			panic(err)
		}
	})
	http.ListenAndServe(":8080", nil)
}
