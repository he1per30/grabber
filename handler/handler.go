package handler

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func ParseSiteHandler(url string, w int) {
	for i := 0; i < w; i++ {
		go func() {
			res, err := http.Get(url)
			if err != nil {
				log.Fatal(err)
			}
			defer res.Body.Close()
			if res.StatusCode != 200 {
				log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
			}
		}()
	}

	for i := 0; i < 5; i++ {

		fmt.Println("====================")
		time.Sleep(1 * time.Second)
	}
}
