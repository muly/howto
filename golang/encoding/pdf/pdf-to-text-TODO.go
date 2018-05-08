package main

import (
	"fmt"
	"github.com/sajari/docconv"
	"os"
)

func main() {
	f, err := os.Open("journal.pdf")
	if err != nil {
		fmt.Println(err)
		return
	}

	text, m, err := docconv.ConvertPDF(f)
	if err != nil {
		fmt.Println("error converting pdf to text:", err)
		return
	}

	fmt.Println(m) // Question: what is this map for?
	fmt.Println("///////////////////////////////////////////////////////////////////////")
	fmt.Println(text)

}

// Got the below errors: project's github read me has mention about dependencies. need to review it to understand how to use this package to convert pdf to text.
//2017/11/21 23:28:25 pdftotext: exec: "pdftotext": executable file not found in $PATH
//2017/11/21 23:28:25 pdfinfo: exec: "pdfinfo": executable file not found in $PATH
