package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/gocolly/colly"
	"github.com/gocolly/redisstorage"
)

// func main() {

// 	fName := "data.csv"
//     file, err := os.Create(fName)
//     if err != nil {
//         log.Fatalf("Could not create file, err: %q", err)
//         return
//     }
//     defer file.Close()

//     writer := csv.NewWriter(file)
//     defer writer.Flush()
//     c := colly.NewCollector(
//         colly.AllowedDomains("w3schools.com"),
//     )
// 	if p, err := proxy.RoundRobinProxySwitcher(
// 		"socks5://127.0.0.1:1337",
// 		"socks5://127.0.0.1:1338",
// 		"http://127.0.0.1:8080",
// 	); err == nil {
// 		c.SetProxyFunc(p)
// 	}


func main() {
    fName := "data.csv"
    file, err := os.Create(fName)
    if err != nil {
        log.Fatalf("Could not create file, err: %q", err)
        return
    }
    defer file.Close()

    writer := csv.NewWriter(file)
    defer writer.Flush()
		storage := &redisstorage.Storage{
    Address:  "127.0.0.1:6379",
    Password: "",
    DB:       0,
    Prefix:   "job01",
	}

	err := c.SetStorage(storage)
	if err != nil {
		panic(err)
	}

    c := colly.NewCollector()
    c.OnHTML("table#customers", func(e *colly.HTMLElement) {
        e.ForEach("tr", func(_ int, el *colly.HTMLElement) {
            writer.Write([]string{
                el.ChildText("td:nth-child(1)"),
                el.ChildText("td:nth-child(2)"),
                el.ChildText("td:nth-child(3)"),
            })
        })
        fmt.Println("Scraping Complete")
    })
    c.Visit("https://www.w3schools.com/html/html_tables.asp")
}