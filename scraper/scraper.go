package scraper

import (
	"fmt"
	"github.com/gocolly/colly"
)

type Content struct {
	url      string
	headline string
	text     string
}

func GetHackerNews() ([]Content, error) {
	var allContents []Content
	fmt.Println("running...")
	getData()
	fmt.Println("done")
	return allContents, nil
}

func getData() {
	c := colly.NewCollector()

	// Find and visit all links
	c.OnHTML("a", func(e *colly.HTMLElement) {
		e.Request.Visit(e.Attr("href"))
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.Visit("http://go-colly.org/")
}
