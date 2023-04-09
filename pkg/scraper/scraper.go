package scraper

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/gocolly/colly"
)

func ScrapeWebPage(url string) {
	attributes := []string{"a[href^=\"/alphabetical/" + "\"]", "p > a[href^=\"/Movie Scripts/\"]", "td > a[href^=\"/scripts/\"]"}
	c := colly.NewCollector()
	movieLinksCol := c.Clone()
	movieScriptsLinkCol := c.Clone()
	movieScriptCol := c.Clone()

	crawl(attributes[0], url, c, movieLinksCol)
	crawl(attributes[1], url, movieLinksCol, movieScriptsLinkCol)
	crawl(attributes[2], url, movieScriptsLinkCol, movieScriptCol)

	attribute := "td.scrtext > pre"
	movieScriptCol.OnHTML(attribute, func(e *colly.HTMLElement) {
        htmlDoc := e.DOM;
		_ = htmlDoc.Find("b").Remove()
		_, err := e.DOM.Html()
		if err != nil {
			fmt.Println(err)
		}

        log.Println(e.Request.URL.Path)
		f, err := os.Create("." + e.Request.URL.Path)
		if err != nil {
			fmt.Println(err)
		}
		defer f.Close()

		nrByte, err := f.Write([]byte(htmlDoc.Text()))
		if err != nil {
			fmt.Println(err)
		}

		fmt.Printf("nr of bytes written:%b, filename:%s\n", nrByte, e.Request.URL.Path)
	})

	if err := c.Visit(url); err != nil {
		log.Fatalf("Error visiting web page: %s | error: %s", url, err)
	}
}

func crawl(attribute string, url string, curCol *colly.Collector, nextCol *colly.Collector) (err error) {
	fmt.Printf("crawl started %s", curCol)
	curCol.OnHTML(attribute, func(e *colly.HTMLElement) {
		link := e.Attr("href")
		if nextCol == nil {
			errMsg := "crawler cannot move deeper,nextCol  = nil"
			log.Println(errMsg)
			err = errors.New(errMsg)
			return
		}

		if err := nextCol.Visit(url + link); err != nil {
			log.Println(err)
			return
		}
	})
	return
}
