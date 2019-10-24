//simple static template that just reads the temaplate from a constant parses and executes it
package main

import (
	"net/http"
	"text/template"
)

const stmpl = `<!DOCTYEP=html>
<html lang="en">
<head>
<meta charset="UTF-8">
<title>static-template1</title>
<body>
Hello World from Static Template1
</body>
</html>
`

func main() {
	tmpl, err := template.New("static").Parse(stmpl)
	if err != nil {
		panic(err)
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		tmpl.Execute(w, nil) //no data to pass so nill
	})
	http.ListenAndServe(":8080", nil)
}
