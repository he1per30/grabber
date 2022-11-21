package handler

import (
	"grabber/pkg"
	"log"
	"net/http"
)

var (
	WORKERS       int    = 2            //кол-во "потоков"
	REPORT_PERIOD int    = 10           //частота отчетов (сек)
	DUP_TO_STOP   int    = 500          //максимум повторов до останова
	HASH_FILE     string = "hash.bin"   //файл с хешами
	QUOTES_FILE   string = "quotes.txt" //файл с цитатами
)

func ParseSiteHandler(url string) {
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	go pkg.ParseSite(res)
}
