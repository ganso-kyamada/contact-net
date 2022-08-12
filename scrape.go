package main

import (
	"fmt"

	"github.com/gocolly/colly"
)


const (
	lpath = "/rsvWUserAttestationAction.do"
	mpath = "/lotWTransLotAcceptListAction.do"
)

func Reservation(uid string, pass string, url string) {
	first := true
	login := url + lpath
	menu := url + mpath

	c := colly.NewCollector()
	c.OnHTML("input[name=loginJKey]", func(e *colly.HTMLElement) {
		if (first) {
			fmt.Println("Login...")
			first = false
			err := c.Post(login, map[string]string{
				"userId": uid,
				"password": pass,
				"displayNo": "pawab2000",
				"loginJKey": e.Attr("value"),
			})
			if err != nil {
				fmt.Println("Error!!!", err)
			}
		}
	})

	c.OnHTML("input.logoutbtn", func(e *colly.HTMLElement) {
		if (e.Request.URL.String() == login) {
			fmt.Println("Login suceess!!")
			err := c.Post(menu, map[string]string{
				"displayNo": "plwac3000",
				"selectPpsPpsdCd": "100",
				"selectPpsCd": "100130",
			})
			if err != nil {
				fmt.Println("Error!!!", err)
			}
		}

		if (e.Request.URL.String() == menu) {
			fmt.Println("TODO: Baseball field reservation")
		}
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Request", r.URL)
	})

	c.Visit(url)
}
