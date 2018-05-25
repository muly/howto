package main

import (
	"fmt"
	"golang.org/x/net/html"
	"io"
	"net/http"
	"os"
)

func prep() ([]html.Token, io.Reader, error) {
	rtp := initFileGoblogDate()
	//rtp:= initFileGoblogAuthor()
	b, err := inputfromFileGoblog()

	//rtp:= initFileExample1()
	//b, err := inputfromFileExample1()

	return rtp, b, err
}

func main() {
	rtp, b, err := prep()
	if err != nil {
		fmt.Println(err)
		return
	}

	//n, err := html.Parse(resp.Body)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//
	//FindMatchingNode(n, requiredNodePath, isMatchingNode)

	t := crawl(b, rtp)
	if t.Type != html.ErrorToken {
		fmt.Println("output ::::::::::::::::", t.Type, t.Data, t.Attr)
	} else {
		fmt.Println("output :::::::::::::::: not found")
	}

}

func crawl(b io.Reader, rtp []html.Token) html.Token {
	z := html.NewTokenizer(b)
	for i, m := range rtp {
		fmt.Println("compare with", m.Type, m.Data, m.Attr)
		for {
			tt := z.Next()
			if tt == html.ErrorToken {
				return html.Token{}
			}
			t := z.Token()
			//if t.Data =="div" {
			//	fmt.Println(t.Type, t.Data, t.Attr)
			//}
			fmt.Println("	match", t.Type, t.Data, t.Attr)
			if isMatchingToken(t, m) {
				fmt.Println("	matched *************")
				if i == len(rtp)-1 {
					return t
				}
				break
			}
		}
	}
	return html.Token{}
}

func isMatchingToken(t html.Token, m html.Token) bool {
	if m.Type != html.ErrorToken {
		if m.Type != t.Type {
			return false
		}
	}
	if m.Data != "" {
		if m.Data != t.Data {
			return false
		}
	}
	if len(m.Attr) > 0 {
		return isMatchingAttr(m.Attr, t.Attr)
	}
	return true
}

func isMatchingAttr(aa []html.Attribute, bb []html.Attribute) bool {
	if len(aa) != len(bb) {
		return false
	}

	ma := sliceAttribute2map(aa)
	mb := sliceAttribute2map(bb)
	//fmt.Println("		ma",ma)
	fmt.Println("		mb", mb)

	for a, acnt := range ma {
		if bcnt, exists := mb[a]; !exists || acnt != bcnt {
			return false
		}
	}
	return true
}

func sliceAttribute2map(aa []html.Attribute) map[html.Attribute]int {
	m := map[html.Attribute]int{}
	for _, a := range aa {
		m[a]++
	}
	return m
}

func initFileGoblogDate() []html.Token { // init for url `http://goblog.qwest.io/2017/09/protobuf-for-go-quick-reference.html`
	tokenPath := []html.Token{}
	n := html.Token{}
	a := html.Attribute{}

	n = html.Token{}
	n.Type = html.StartTagToken
	n.Data = "div"
	a = html.Attribute{}
	a.Key = "class"
	a.Val = "post-footer-line post-footer-line-1"
	n.Attr = append(n.Attr, a)
	tokenPath = append(tokenPath, n)

	n = html.Token{}
	n.Type = html.StartTagToken
	n.Data = "span"
	a = html.Attribute{}
	a.Key = "class"
	a.Val = "post-timestamp"
	n.Attr = append(n.Attr, a)
	tokenPath = append(tokenPath, n)

	// abbr class='published' itemprop='datePublished'
	//TODO: pub-date is not retrieved correctly. need to debug to find if the issue is with the below data or a code bug.
	n = html.Token{}
	n.Type = html.StartTagToken
	n.Data = "abbr"
	a = html.Attribute{}
	a.Key = "class"
	a.Val = "published"
	n.Attr = append(n.Attr, a)
	a = html.Attribute{}
	a.Key = "itemprop"
	a.Val = "datePublished"
	n.Attr = append(n.Attr, a)
	tokenPath = append(tokenPath, n)

	return tokenPath
}

func initFileGoblogAuthor() []html.Token { // init for url `http://goblog.qwest.io/2017/09/protobuf-for-go-quick-reference.html`
	tokenPath := []html.Token{}
	n := html.Token{}
	a := html.Attribute{}

	n = html.Token{}
	n.Type = html.StartTagToken
	n.Data = "div"
	a = html.Attribute{}
	a.Key = "class"
	a.Val = "post-footer-line post-footer-line-1"
	n.Attr = append(n.Attr, a)
	tokenPath = append(tokenPath, n)

	n = html.Token{}
	n.Type = html.StartTagToken
	n.Data = "span"
	a = html.Attribute{}
	a.Key = "class"
	a.Val = "post-author vcard"
	n.Attr = append(n.Attr, a)
	tokenPath = append(tokenPath, n)

	n = html.Token{}
	n.Type = html.StartTagToken
	n.Data = "span"
	a = html.Attribute{}
	a.Key = "itemprop"
	a.Val = "name"
	n.Attr = append(n.Attr, a)
	tokenPath = append(tokenPath, n)

	n = html.Token{}
	n.Type = html.TextToken
	tokenPath = append(tokenPath, n)

	return tokenPath
}

func inputfromUrlGoblog() (io.Reader, error) {
	//url := `https://godoc.org/golang.org/x/oauth2`
	//url := `https://socketloop.com/tutorials/golang-read-file`
	url := `http://goblog.qwest.io/2017/09/protobuf-for-go-quick-reference.html`

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	return resp.Body, nil
}

func inputfromFileGoblog() (io.Reader, error) {
	filename := `C:\gows\src\github.com\muly\howto\golang\web\http\html-parsing\net_html\exampleGoBlog.html`
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	return f, nil
}

func inputfromFileExample1() (io.Reader, error) {
	filename := `C:\gows\src\github.com\muly\howto\golang\web\http\html-parsing\net_html\example1.html`
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	return f, nil
}

func initFileExample1() []html.Token { // for the example1.html
	requiredTokenPath := []html.Token{}
	n := html.Token{}
	a := html.Attribute{}

	n = html.Token{}
	n.Type = html.StartTagToken
	n.Data = "div"
	a = html.Attribute{}
	a.Key = "style"
	a.Val = "color:#0000FF"
	n.Attr = append(n.Attr, a)
	requiredTokenPath = append(requiredTokenPath, n)

	n = html.Token{}
	n.Type = html.StartTagToken
	n.Data = "h2"
	a = html.Attribute{}
	a.Key = "class"
	a.Val = "city"
	n.Attr = append(n.Attr, a)
	requiredTokenPath = append(requiredTokenPath, n)

	n = html.Token{}
	n.Type = html.TextToken
	requiredTokenPath = append(requiredTokenPath, n)

	return requiredTokenPath
}

/*
func FindMatchingNode(root *html.Node, match []html.Node, walkF WalkFunc) {
	for i, m := range match {
		//for m := match.FirstChild; m != nil; m = m.NextSibling {
		//fmt.Println(nodeTypeText(m.Type), m.Data, m.Attr)
		var f func(*html.Node) *html.Node
		f = func(n *html.Node) *html.Node {
			//fmt.Println(nodeTypeText(n.Type), n.Data, n.Attr)
			if walkF(n, m) {
				fmt.Println(i, nodeTypeText(n.Type), n.Data, n.Attr)
				return n
			}

			for c := n.FirstChild; c != nil; c = c.NextSibling {
				f(c)
			}
			return nil
		}
		root = f(root)
	}
}

func isMatchingNode(n *html.Node, matchWith html.Node) bool {
	if n.Type == matchWith.Type && n.Data == matchWith.Data {
		matchCnt := 0
		for _, a := range matchWith.Attr {
			for _, w := range n.Attr {
				if a.Key == w.Key && a.Val == w.Val {
					matchCnt++
				}
			}
		}
		if matchCnt != len(matchWith.Attr) {
			return false
		}
		return true
	}
	return false
}

type WalkFunc func(*html.Node, html.Node) bool


func nodeTypeText(nodeType html.NodeType) string {
	switch nodeType {
	case html.ErrorNode:
		return "ErrorNode"
	case html.TextNode:
		return "TextNode"
	case html.DocumentNode:
		return "DocumentNode"
	case html.ElementNode:
		return "ElementNode"
	case html.CommentNode:
		return "CommentNode"
	case html.DoctypeNode:
		return "DoctypeNode"
	}
	return ""

}


var requiredNodePath []html.Node

*/
