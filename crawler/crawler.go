package main

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/gocolly/colly"
)

type page struct {
	Title string
	Url   url.URL
}

func main() {
	pages := []page{}
	c := colly.NewCollector(
		// Visit only domains: hackerspaces.org, wiki.hackerspaces.org
		colly.AllowedDomains("ses.jkuat.ac.ke"),
	)

	// On every a element which has href attribute call callback
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		// Print link
		fmt.Printf("Link found: %q -> %s\n", e.Text, link)
		// Visit link found on page
		// Only those links are visited which are in AllowedDomains
		c.Visit(e.Request.AbsoluteURL(link))
	})
	c.OnHTML("head", func(e *colly.HTMLElement) {
		title := e.ChildText("title")
		title = strings.Replace(title, "\n", "", -1)
		title = strings.Replace(title, "\t", "", -1)
		fmt.Printf("Page Title: %q\n", title)
		currentPage := page{Url: *e.Request.URL, Title: title}
		pages = append(pages, currentPage)
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	// Start scraping on
	c.Visit("https://ses.jkuat.ac.ke/")
	fmt.Println(len(pages))
}
