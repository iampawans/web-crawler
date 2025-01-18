package main

import (
    "fmt"
    "net/http"
    "io/ioutil"
    "log"
    "encoding/json"
    "github.com/PuerkitoBio/goquery"
)

func startCrawl(w http.ResponseWriter, r *http.Request) {
    resp, err := http.Get("http://example.com")
    if err != nil {
        log.Fatal(err)
    }
    defer resp.Body.Close()

    body, _ := ioutil.ReadAll(resp.Body)
    doc, err := goquery.NewDocumentFromReader(string(body))
    if err != nil {
        log.Fatal(err)
    }

    doc.Find("a").Each(func(index int, item *goquery.Selection) {
        link, _ := item.Attr("href")
        fmt.Println(link)
    })
    
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]string{"status": "Crawl started!"})
}

func main() {
    http.HandleFunc("/start-crawl", startCrawl)
    fmt.Println("Crawl Service is running on port 8080...")
    http.ListenAndServe(":8080", nil)
}
