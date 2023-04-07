package scraper

import (
    "fmt"
    "log"

    "github.com/gocolly/colly"
)

func ScrapeWebPage(url string) {
    c := colly.NewCollector()
    movieLinksCol := c.Clone()
    movieScriptsLinkCol := c.Clone()
    movieScriptCol := c.Clone()
   
    attribute := "a[href^=\"/alphabetical/" + "\"]"
    c.OnHTML(attribute, func(e *colly.HTMLElement){
        link := e.Attr("href")
        fmt.Println(e.Text)
        if err := movieLinksCol.Visit(url + link); err != nil{
            fmt.Println(err)
        }})


    attribute = "p > a[href^=\"/Movie Scripts/\"]"
    fmt.Println(attribute)

     movieLinksCol.OnHTML(attribute, func(e *colly.HTMLElement){
        link := e.Attr("href")
        fmt.Println(url + link)
        if err := movieScriptsLinkCol.Visit(url + link); err != nil{
            fmt.Println(err)
        }
    })

    attribute = "td > a[href^=\"/scripts/\"]"
    movieScriptsLinkCol.OnHTML(attribute, func(e *colly.HTMLElement){
        link := e.Attr("href")
        fmt.Println(link)
        movieScriptCol.Visit(url + link)
    })

    attribute = "td.scrtext > pre"
    movieScriptCol.OnHTML(attribute, func(e *colly.HTMLElement){
        text := e.Text
        fmt.Println(text)
    })

    c.OnScraped(func(r *colly.Response) {
        fmt.Println("Finished", r.Request.URL)
    })

    if err := c.Visit(url); err != nil{
        log.Fatalf("Error visiting web page: %s | error: %s", url, err)
    }
}
