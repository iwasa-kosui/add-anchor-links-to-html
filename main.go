package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/PuerkitoBio/goquery"
)

func newDocumentFromFileName(filename string) (*goquery.Document, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	doc, err := goquery.NewDocumentFromReader(f)
	if err != nil {
		return nil, err
	}
	return doc, nil
}

func main() {
	var i = flag.String("i", "", "input filename")
	var o = flag.String("o", "", "output filename")
	flag.Parse()

	doc, err := newDocumentFromFileName(*i)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find("h1, h2, h3, h4, h5, h6, span, b").Each(func(i int, s *goquery.Selection) {
		shtml, err := goquery.OuterHtml(s)
		if err != nil {
			return
		}
		id, exists := s.Attr("id")
		if !exists {
			return
		}
		s.ReplaceWithHtml(fmt.Sprintf("<a href=#%s>%s</a>", id, shtml))
	})

	d, _ := doc.Html()
	if err != nil {
		log.Fatal(err)
	}

	if *o != "" {
		ioutil.WriteFile(*o, []byte(d), 0666)
		return
	}
	fmt.Println(d)
}
