package scraper

import (
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
)

type Content struct {
	URL      string
	Headline string
}

func GetHackerNews() ([]Content, error) {
	var allContents []Content
	getData(&allContents)
	return allContents, nil
}

func getData(allContents *[]Content) {

	// Instantiate default collector
	c := colly.NewCollector()
	extensions.RandomUserAgent(c)

	// On every a element which has href attribute call callback
	c.OnHTML("a[href].storylink", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		if len(link) > 0 && len(e.Text) > 0 {
			*allContents = append(*allContents, Content{URL: link, Headline: e.Text})
		}
	})

	// Start scraping on https://hackerspaces.org
	c.Visit("https://news.ycombinator.com/")
}
