package main

import (
	"fmt"
	"os"
    "io/ioutil"
	"time"
    "encoding/json"
    "net/url"
	"github.com/gocolly/colly"
	"strings"

)

func joinUrl(urlToJoin string) string{

	var base = "http://books.toscrape.com/catalogue/"

	u, err := url.JoinPath(base, urlToJoin)
	if err != nil {
		os.Exit(1)
	}
	return u
}


type item struct {
	BookTitle string `json:"Book Title"`
	BookUrl string `json:"Book Url"`
	BookPrice string `json:"Book Price"`
	CrawledAt time.Time `json:"Crawled at Time"`
}

func main() {
	stories := []item{}
	joinUrl("base.html")
	// Instantiate default collector
	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/41.0.2228.0 Safari/537.36"),
		colly.AllowedDomains("books.toscrape.com"),
		colly.MaxDepth(2),
		colly.Async(true),
	)

	c.OnHTML("article.product_pod", func(e *colly.HTMLElement) {
		temp := item{
		BookTitle : e.ChildText("h3"),
		BookPrice :e.ChildText("p.price_color"),
		CrawledAt: time.Now(),
		}
		temp.BookUrl = joinUrl(e.ChildAttr("a", "href"))
		stories = append(stories, temp)
	})


	// On every span tag with the class next-button
	c.OnHTML("li.next", func(h *colly.HTMLElement) {
		var pageUrl = h.ChildAttr("a", "href")
		var baseUrl string
	if strings.Contains(pageUrl, "catalogue") {
		baseUrl = "http://books.toscrape.com/"
	} else{
		baseUrl = "http://books.toscrape.com/catalogue"
	}
		t,err :=url.JoinPath(baseUrl,pageUrl )
		if err != nil {
		os.Exit(1)
	}
		fmt.Println("Visiting",t)
		c.Visit(t)
	})

	// Set max Parallelism and introduce a Random Delay
	c.Limit(&colly.LimitRule{
		Parallelism: 20,
		RandomDelay: 5 * time.Second,
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())

	})

	// Crawl all pages the user passes in
	scrape_pages := [2]string{"http://books.toscrape.com/"}
	for _, scrape_page := range scrape_pages {
		c.Visit(scrape_page)

	}

	c.Wait()
	data, err := json.MarshalIndent(stories, "", "  ")
	if err!=nil{
		fmt.Println(err)
	}
	ioutil.WriteFile("Books.json", data, 0644)
	fmt.Println(stories)

}