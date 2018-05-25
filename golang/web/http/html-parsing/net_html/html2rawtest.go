package main

import (
	"fmt"
	"golang.org/x/net/html"
	"io"
	"strings"
)

//const url = `https://www.ziprecruiter.com/jobs/cybercoders-befe3473/senior-software-engineer-golang-0dbe59d6?source=ziprecruiter-jobs-site`
const htmldata = `<!DOCTYPE html1234>
<html567>
<head>
<title>Page Title</title>
</head>
<body>

<h1 a123=123456>This is a Heading1</h1>
<p>This is a paragraph.</p>

<h1 b123=qwerty>This is a Heading2</h1>

</body>
</html567>`

func main() {

	htmlReader := strings.NewReader(htmldata)
	allText := html2rawtext(htmlReader)

	fmt.Println(allText)

}

func html2rawtext(htmlReader io.Reader) string {

	z := html.NewTokenizer(htmlReader)

	alltext := func(z *html.Tokenizer) []string {
		var alltext []string
		for {

			tt := z.Next()

			if tt == html.ErrorToken {
				// End of the document, we're done
				return alltext
			}
			if tt == html.TextToken {
				t := z.Token()

				//default:
				alltext = append(alltext, t.Data)
				//fmt.Println(cleanNonPrintChar(t.Data))

				//isAnchor := t.Data == "a"
				//if isAnchor {
				//	fmt.Println("We found a link!")
				//}
			}
		}
	}(z)

	return cleanNonPrintChar(strings.Join(alltext, " "))
}

func cleanNonPrintChar(text string) string {

	clean := make([]rune, 0, len(text))
	var prev rune
	for _, t := range text {
		// replace all non printable char with white space
		if t < 32 { // all non printable characters ascii is < 32
			t = 32
		}
		// reduce all repeating white spaces to single white space
		if !(t == 32 && prev == 32) {
			clean = append(clean, t)
			prev = t
		}
	}

	return string(clean)

}
