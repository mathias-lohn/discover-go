package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/PuerkitoBio/goquery"
)

func scrape(url string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Failed to fetch the URL: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Fatalf("Failed to get a valid response: %v", resp.Status)
	}

	document, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatalf("Failed to parse the HTML: %v", err)
	}
  
	document.Find("a").Each(func(index int, element *goquery.Selection) {
		href, exists := element.Attr("href")
		if exists {
			fmt.Printf("Link #%d: %s\n", index+1, href)
		}
	})
}

func main() {
	url := "https://github.com"

	scrape(url)
}
