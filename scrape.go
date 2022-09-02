package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

type Path struct {
	Login         string
	Menu          string
	Lottery       string
	GroundMenu    string
	GroundNumMenu string
	Apply         string
	Complete      string
}

func Reservation(uid string, pass string, url string) {
	first := true
	path := Path{
		Login:         url + "/rsvWUserAttestationAction.do",
		Menu:          url + "/lotWTransLotAcceptListAction.do",
		Lottery:       url + "/lotWTransLotBldGrpAction.do",
		GroundMenu:    url + "/lotWTransLotInstGrpAction.do",
		GroundNumMenu: url + "/lotWTransLotInstSrchVacantAction.do",
		Apply:         url + "/lotWInstTempLotApplyAction.do",
		Complete:      url + "/lotWInstLotApplyAction.do",
	}

	c := colly.NewCollector()
	c.OnHTML("input[name=loginJKey]", func(e *colly.HTMLElement) {
		if !first {
			return
		}

		fmt.Printf("UID[%s] Login...\n", uid)
		first = false
		err := c.Post(path.Login, map[string]string{
			"userId":    uid,
			"password":  pass,
			"displayNo": "pawab2000",
			"loginJKey": e.Attr("value"),
		})
		if err != nil {
			fmt.Printf("UID[%s] Login Error!!! %s\n", uid, err)
		}
	})

	c.OnHTML("input.logoutbtn", func(e *colly.HTMLElement) {
		switch e.Request.URL.String() {
		case path.Login:
			fmt.Printf("UID[%s] Login suceess!!\n", uid)
			c.Post(path.Menu, map[string]string{
				"displayNo":       "plwac3000",
				"selectPpsPpsdCd": "100",
				"selectPpsCd":     "100130",
			})
		case path.Menu:
			fmt.Printf("UID[%s] Menu\n", uid)
			c.Post(path.Lottery, map[string]string{
				"displayNo":      "plwad1000",
				"selectClassCd": "5060010",
			})
		case path.Lottery:
			fmt.Printf("UID[%s] Lottery\n", uid)
			c.Post(path.GroundMenu, map[string]string{
				"displayNo":      "plwba1000",
				"selectBldGrpCd": "5450010",
			})
			// INFO: 複数のグラウンドがあった場合（第一運動場、第二運動場など）
			// err := c.Post(path.GroundNumMenu, map[string]string{
			// 	"displayNo":      "plwba2000",
			// 	"selectBldGrpCd": "5110030",
			// 	"selectUseNumber": "50", // ここの値は固定値か要確認
			// })
		case path.GroundMenu, path.GroundNumMenu:
			fmt.Printf("UID[%s] GroundMenu\n", uid)
			c.Post(path.Apply, map[string]string{
				"selectFieldCnt":  "1",
				"displayNo":       "plwba3000",
				"selectUseYMD":    "20221002",
				"selectTimeNum":   "1",
				"maxFieldCnt":     "1",
				"selectDispStime": "800",
				"selectDispEtime": "1000",
				"dummy":           "",
				"dummy2":          "",
			})
		case path.Complete:
			fmt.Printf("UID[%s] Complete!!\n", uid)
		}
	})

	c.OnHTML("input[name=insLotJKey]", func(e *colly.HTMLElement) {
		fmt.Printf("exists insLotJKey %s\n", e.Request.URL.String())
		c.Post(path.Complete, map[string]string{
			"displayNo":       "plwca1000",
			"applyPepopleNum": "15", // 人数
			"selectApplyNo":   "-1",
			"selectHopeNo":    "1",
			"selectPpsName":   "%8F%AD%94N%96%EC%8B%85%81i%8F%AC%81E%92%86%8Aw%90%B6%81j",
			"insLotJKey": e.Attr("value"),
			"dummy": "",
		})
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Printf("UID[%s] Request %s\n", uid, r.URL)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Printf("UID[%s] Response code %d\n", uid, r.StatusCode)
	})

	c.Visit(url)
}
