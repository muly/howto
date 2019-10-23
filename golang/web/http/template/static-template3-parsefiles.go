package main

import (
	"net/http"
	"text/template"
)

func main() {
	
	view:= template.New("") // this creates 1 template with blank name
	var err error
	view, err = view.ParseFiles("stpl2.gohtml") // this creates another template with name "stpl2.gohtml"
	if err != nil {
		panic(err)
	}

	
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		err := view.ExecuteTemplate(w, "stpl2.gohtml", nil) // Note: we need to execute the specific template
		if err != nil {
			panic(err)
		}
	})
	http.ListenAndServe(":8080", nil)
}
