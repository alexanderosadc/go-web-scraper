package main

import "github.com/alexanderosadc/go-web-app/pkg/scraper"

func main() {
	var selectors :=[]string{
		
	}

	scraper.ScrapeWebPage("https://quotes.toscrape.com/", selectors)
}
