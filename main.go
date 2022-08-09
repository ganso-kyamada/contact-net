package main

import (
	"fmt"
	"log"
	"github.com/PuerkitoBio/goquery"
)

func main() {
	get_url_info, err := goquery.NewDocument("https://www.fureai-net.city.kawasaki.jp/web/")
	if err != nil {
		log.Fatal(err)
	}

	get_url_info.Find("#userId").Each(func(i int, s *goquery.Selection) {
		fmt.Println("Hello GoLang!")
	})
}

