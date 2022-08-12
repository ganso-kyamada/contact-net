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
		if first {
			fmt.Printf("UID[%s] Login...\n", uid)
			first = false
			err := c.Post(login, map[string]string{
				"userId":    uid,
				"password":  pass,
				"displayNo": "pawab2000",
				"loginJKey": e.Attr("value"),
			})
			if err != nil {
				fmt.Printf("UID[%s] Error!!! %s\n", uid, err)
			}
		}
	})

	c.OnHTML("input.logoutbtn", func(e *colly.HTMLElement) {
		if e.Request.URL.String() == login {
			fmt.Printf("UID[%s] Login suceess!!\n", uid)
			err := c.Post(menu, map[string]string{
				"displayNo":       "plwac3000",
				"selectPpsPpsdCd": "100",
				"selectPpsCd":     "100130",
			})
			if err != nil {
				fmt.Printf("UID[%s] Error!!! %s\n", uid, err)
			}
		}

		if e.Request.URL.String() == menu {
			fmt.Printf("UID[%s] TODO: Baseball field reservation\n", uid)
		}
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Printf("UID[%s] Request %s\n", uid, r.URL)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Printf("UID[%s] Response code %d\n", uid, r.StatusCode)
	})

	c.Visit(url)
}
