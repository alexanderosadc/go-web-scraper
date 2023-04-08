package scraper

import (
    "fmt"
    "log"
    "os"

    "github.com/gocolly/colly"
)

func ScrapeWebPage(url string) {
    colly.CacheDir("./cache/")
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
     movieLinksCol.OnHTML(attribute, func(e *colly.HTMLElement){
        link := e.Attr("href")
        if err := movieScriptsLinkCol.Visit(url + link); err != nil{
            fmt.Println(err)
        }
    })

    attribute = "td > a[href^=\"/scripts/\"]"
    movieScriptsLinkCol.OnHTML(attribute, func(e *colly.HTMLElement){
        link := e.Attr("href")
        if err := movieScriptCol.Visit(url + link); err != nil{
            fmt.Println(err)
        }
    })

    attribute = "td.scrtext > pre"
    movieScriptCol.OnHTML(attribute, func(e *colly.HTMLElement){
        _ = e.DOM.Find("b").Remove()
        res, err :=e.DOM.Html()
        if err != nil{
            fmt.Println(err)
        }

        f, err := os.Create("." + e.Request.URL.Path)
        if err != nil{
            fmt.Println(err) 
        }
        defer f.Close()

        nrByte, err := f.Write([]byte(res))
        if err != nil{
            fmt.Println(err)
        }

        fmt.Printf("nr of bytes written:%b, filename:%s\n", nrByte, e.Request.URL.Path)
    })

    if err := c.Visit(url); err != nil{
        log.Fatalf("Error visiting web page: %s | error: %s", url, err)
    }
}
