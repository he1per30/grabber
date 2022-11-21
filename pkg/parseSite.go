package pkg

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
)

func ParseSite(res *http.Response) {
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	title := doc.Find(".fi_text").Text()
	fmt.Printf("%s \n", title)
	fmt.Println("=================================================")
	// Find the review items

}
