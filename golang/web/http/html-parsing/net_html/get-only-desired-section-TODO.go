package main

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
)

//sequence:
// Type: ElementNode
// Data: h2
// Attr.Key: id
// Attr.Val: pkg-overview

var requiredNodePath []html.Node

func init() {
	n := html.Node{}
	n.Type = html.ElementNode
	n.Data = "div"
	a := html.Attribute{}
	a.Key = "class"
	a.Val = "post-footer-line post-footer-line-1"
	n.Attr = append(n.Attr, a)
	requiredNodePath = append(requiredNodePath, n)


	n = html.Node{}
	n.Type = html.ElementNode
	n.Data = "div"
	a = html.Attribute{}
	a.Key = "class"
	a.Val = "post-author vcard"
	n.Attr = append(n.Attr, a)
	requiredNodePath = append(requiredNodePath, n)


	n = html.Node{}
	n.Type = html.ElementNode
	n.Data = "span"
	a = html.Attribute{}
	a.Key = "itemprop"
	a.Val = "name"
	n.Attr = append(n.Attr, a)
	requiredNodePath = append(requiredNodePath, n)


	n = html.Node{}
	n.Type = html.TextNode
	requiredNodePath = append(requiredNodePath, n)


	//for _, m := range requiredNodePath { // for debug purpose
	////for m := requiredNodePath.FirstChild; m != nil; m = m.NextSibling {
	//	fmt.Println(nodeTypeText(m.Type), m.Data, m.Attr)
	//}

}

func main() {
	//url := `https://godoc.org/golang.org/x/oauth2`
	//url := `https://socketloop.com/tutorials/golang-read-file`
	url := `http://goblog.qwest.io/2017/09/protobuf-for-go-quick-reference.html`

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	n, err := html.Parse(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	FindMatchingNode(n, requiredNodePath, isMatchingNode)

}

func isMatchingNode(n *html.Node, matchWith html.Node) bool {
	if n.Type == matchWith.Type && n.Data == matchWith.Data { // TODO: check only on non blank fields
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
