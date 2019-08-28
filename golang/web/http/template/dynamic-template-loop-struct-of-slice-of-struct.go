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

type persons struct {
	Tenant string
	Data   []person
}

const dtmpl = `
<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<title>dynamic-templateloop</title>
</head>
<body>
	Tenant: {{ .Tenant }}<br/>
	{{range .Data}}
	Name :{{.Fname}} {{.Mname}}.{{.Lname}}<br/>
	Age :{{ .Age }}<br/>
	{{end}}
</body>
</html>
`

func main() {
	p1 := person{"Praveen", "Kumar", "K", 36}
	p2 := person{"Srinivasa", "Reddy", "M", 36}
	p3 := person{"Mahesh", "Reddy", "M", 36}
	data := persons{Tenant: "tenant1", Data: []person{p1, p2, p3}}

	tmpl, err := template.New("dynamic-loop").Parse(dtmpl)
	if err != nil {
		panic(err)
	}

	//var tpl bytes.Buffer
	//tmpl.Execute(&tpl, data)
	//fmt.Println(tpl.String())
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, data)
	})
	http.ListenAndServe(":8080", nil)
}
