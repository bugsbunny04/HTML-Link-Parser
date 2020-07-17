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
	nodes := linknodes(doc)
	for _, node := range nodes {
		fmt.Println(node)
	}
	return nil,nil
}

func linknodes(n *html.Node) []*html.Node {
	if n.Type==html.ElementNode && n.Data=="a" {
		return []*html.Node{n}
	}
	var ret []*html.Node
	for c:=n.FirstChild;c!=nil;c=c.NextSibling {
		ret = append(ret , linknodes(c)...)
	}
	return ret
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


