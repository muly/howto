package main

import (
	"net/http"
	"text/template"
)

func main() {
	view, err := template.ParseFiles("stpl2.gohtml")
	if err != nil {
		panic(err)
	}

	// tmpl := template.Must(template.ParseFiles("templates/homepage.html", "templates/view.html"))
	
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		err := view.Execute(w, nil)
		if err != nil {
			panic(err)
		}
	})
	http.ListenAndServe(":8080", nil)
}
