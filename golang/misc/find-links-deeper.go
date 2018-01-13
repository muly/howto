package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"strings"
)

func ScrapeLinks(url string) {

	doc, err := goquery.NewDocument(url)

	if err != nil {
		panic(err)
	}

	// process this part :

	// <div class="nav-collapse collapse navbar-responsive-collapse">
	//                <ul class="nav nav-pills">
	//                  <li>
	//                        <a href="/references">References</a>
	//                    </li>
	//                    <li>
	//                        <a href="/tutorials">Tutorials</a>
	//                    </li>
	//                </ul>
	// </div>

	doc.Find(".nav-collapse .nav").Each(func(i int, s *goquery.Selection) {
		Title := strings.TrimSpace(s.Find("li").Text()) // https://www.socketloop.com/tutorials/trim-white-spaces-string-golang

		// convert string to array
		Fields := strings.Fields(Title)

		// go deeper by 1 level to get the <a href=""></a>

		doc.Find(".nav-collapse .nav a").Each(func(i int, s *goquery.Selection) {
			Link, _ := s.Attr("href")
			Link = url + Link

			fmt.Printf("Title is [%s] and link is [%s]\n", Fields[i], Link)
		})

	})

}

func main() {

	ScrapeLinks("https://socketloop.com")

}
