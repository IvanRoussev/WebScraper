package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"log"
)

func Scraper() {

	c := colly.NewCollector()

	site := "https://scrapeme.live/shop/"

	err := c.Visit(site)
	if err != nil {
		log.Fatalf("could not visit site: %v: error: %v", site, err)
	}

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting: ", r.URL)
	})

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong: ", err)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Page Visited: ", r.Request.URL)
	})

	c.OnHTML("a", func(e *colly.HTMLElement) {
		fmt.Printf("%v\n", e.Attr("href"))
	})

	c.OnScraped(func(r *colly.Response) {
		fmt.Println(r.Request.URL, " scraped!")
	})

	fmt.Println("Hello World")
}
