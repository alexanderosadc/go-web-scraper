package scraper

import (
	"fmt"
	"log"
	"os"

	"github.com/gocolly/colly"
)

func ScrapeWebPage(url string) {
    c := colly.NewCollector()
    attributes := []string{"a[href^=\"/alphabetical/" + "\"]", "p > a[href^=\"/Movie Scripts/\"]", "td > a[href^=\"/scripts/\"]"}
    var nextCol *colly.Collector

    for _, attribute := range(attributes){
        nextCol = crawl(attribute, url)
        fmt.Printf("after crawl %#v \n", nextCol)
    }

    if nextCol == nil{
        log.Println("nextCol == nil")
        return
    }

	attribute := "td.scrtext > pre"
	nextCol.OnHTML(attribute, func(e *colly.HTMLElement) {
		_ = e.DOM.Find("b").Remove()
		res, err := e.DOM.Html()
		if err != nil {
			fmt.Println(err)
		}

		f, err := os.Create("." + e.Request.URL.Path)
		if err != nil {
			fmt.Println(err)
		}
		defer f.Close()

		nrByte, err := f.Write([]byte(res))
		if err != nil {
			fmt.Println(err)
		}

		fmt.Printf("nr of bytes written:%b, filename:%s\n", nrByte, e.Request.URL.Path)
	})

	if err := c.Visit(url); err != nil {
		log.Fatalf("Error visiting web page: %s | error: %s", url, err)
	}
}

func crawl(attribute string, url string) (nextCol *colly.Collector) {
	c := colly.NewCollector()
	c.OnHTML(attribute, func(e *colly.HTMLElement) {
		link := e.Attr("href")
		fmt.Println(e.Text)
		nextCol = c.Clone()
		if err := nextCol.Visit(url + link); err != nil {
			log.Println(err)
            nextCol = nil
            return
		}
	})
	return
}
