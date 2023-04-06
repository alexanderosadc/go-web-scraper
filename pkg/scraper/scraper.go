package scraper

import (
	"fmt"

	"github.com/gocolly/colly"
)

func ScrapeWebPage(url string, cssSelectors []string) {
	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	for _, elem := range cssSelectors {
		c.OnHTML(elem, func(e *colly.HTMLElement) {
			fmt.Println(e.Text)
		})
	}

	c.OnScraped(func(r *colly.Response) {
		fmt.Println("Finished", r.Request.URL)
	})

	c.Visit(url)
}
