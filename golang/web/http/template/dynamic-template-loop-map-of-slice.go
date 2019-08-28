package main

import (
	"net/http"
	"text/template"
)

type person struct {
	Fname string
	Mname string
	Lname string
	Age   int
}

const dtmpl = `
<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<title>dynamic-templateloop</title>
</head>
<body>
	{{range $key, $value := .}}
	Tenant: {{ $key }}<br/>
	{{range $value}}
	Name :{{.Fname}} {{.Mname}}.{{.Lname}}<br/>
	Age :{{ .Age }}<br/>
	{{end}}
	{{end}}
</body>
</html>
`

func main() {
	p1 := person{"Praveen", "Kumar", "K", 36}
	p2 := person{"Srinivasa", "Reddy", "M", 36}
	p3 := person{"Mahesh", "Reddy", "M", 36}
	persons := map[string][]person{}
	persons["tenant1"] = []person{p1, p2}
	persons["tenant2"] = []person{p3}

	tmpl, err := template.New("dynamic-loop").Parse(dtmpl)
	if err != nil {
		panic(err)
	}

	//var tpl bytes.Buffer
	//tmpl.Execute(&tpl, persons)
	//fmt.Println(tpl.String())
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, persons)
	})
	http.ListenAndServe(":8080", nil)
}
