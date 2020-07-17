package linkparser

import (
	"fmt"
	"io"
	"golang.org/x/net/html"
)


type Link struct {
	Href string
	Text string
}

func Parser(r io.Reader) ([]Link, error) {
	doc, err := html.Parse(r)
	if err!=nil {
		panic(err)
	}
	dfs(doc,"")
	return nil,nil
}

func dfs(n *html.Node,padding string) {
	msg := n.Data
	if n.Type == html.ElementNode {
		msg = "<"+msg+">"
	}
	fmt.Println(padding,msg)
	for c:=n.FirstChild;c!=nil;c=c.NextSibling {
		dfs(c,padding+" ")
	}
}


