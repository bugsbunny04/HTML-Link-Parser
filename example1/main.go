package main

import (
	"flag"
	link "linkparser"
	"os"
	"fmt"
)

func main() {
	filename := flag.String("file", "ex1.html", "File that contains the HTML to be parsed")
	f, err := os.Open(*filename)
	if err != nil {
		panic(err)
	}

	links, err := link.Parser(f)
	fmt.Printf("%+v",links)

}
