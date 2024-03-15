package main

import (
	"fmt"
	"flag"
	"github.com/gocolly/colly/v2"
)

func main() {
	wordPtr := flag.String("word", "foo", "a string")
	allowedDomain := flag.String("something", "bar", "a string")
	flag.Parse()
	c := colly.NewCollector(colly.AllowedDomains(*allowedDomain))
	
	// Find and visit all links
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		e.Request.Visit(e.Attr("href"))
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})
	//fmt.Println(flag.Args()[0])
	c.Visit(*wordPtr)
}