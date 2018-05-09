package main

import (
	"fmt"
	"net/http"

	"github.com/jung-kurt/gofpdf"
)

func main() {
	http.HandleFunc("/pdf", gofpdfHandler)
	http.ListenAndServe(":8080", nil)
}

func gofpdfHandler(w http.ResponseWriter, r *http.Request) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(40, 10, "Hello, world")

	w.Header().Set("Content-Disposition", "attachment; filename=hello2.pdf")
	w.Header().Set("Content-Type", "application/pdf")
	err := pdf.Output(w)
	if err != nil {
		fmt.Println(err)
		return
	}
}
