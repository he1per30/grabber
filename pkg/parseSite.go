package pkg

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"time"
)

func ParseSite(res *http.Response, w int) chan string {
	c := make(chan string, 5)
	for i := 0; i < w; i++ {
		go func() {
			for {
				doc, err := goquery.NewDocumentFromReader(res.Body)
				if err != nil {
					log.Fatal(err)
				}
				text2 := doc.Find(".fi_text").Text()
				c <- text2
				//if text := strings.TrimSpace(doc.Find(".fi_text").Text()); text != "" {
				//	c <- text
				//}
				time.Sleep(100 * time.Millisecond)
			}
		}()

	}
	fmt.Println("Запущено потоков: ", w)
	return c
	// Find the review items

}
