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

func Reservation(user User) {
	num := 0
	first := true
	path := Path{
		Login:         user.Url + "/rsvWUserAttestationAction.do",
		Menu:          user.Url + "/lotWTransLotAcceptListAction.do",
		Lottery:       user.Url + "/lotWTransLotBldGrpAction.do",
		GroundMenu:    user.Url + "/lotWTransLotInstGrpAction.do",
		GroundNumMenu: user.Url + "/lotWTransLotInstSrchVacantAction.do",
		Apply:         user.Url + "/lotWInstTempLotApplyAction.do",
		Complete:      user.Url + "/lotWInstLotApplyAction.do",
	}

	c := colly.NewCollector()
	c.OnHTML("input[name=loginJKey]", func(e *colly.HTMLElement) {
		if !first {
			return
		}

		fmt.Printf("UID[%s] Login...\n", user.ID)
		first = false
		err := c.Post(path.Login, map[string]string{
			"userId":     user.ID,
			"password":   user.Password,
			"displayNo":  "pawab2000",
			"securityNo": user.SecurityNo,
			"loginJKey":  e.Attr("value"),
		})
		if err != nil {
			fmt.Printf("UID[%s] Login Error!!! %s\n", user.ID, err)
		}
	})

	c.OnHTML("input.logoutbtn", func(e *colly.HTMLElement) {
		switch e.Request.URL.String() {
		case path.Login:
			fmt.Printf("UID[%s] Login suceess!!\n", user.ID)
			visitMenuPage(c, path.Menu)
		case path.Menu:
			fmt.Printf("UID[%s] Menu\n", user.ID)
			visitLotteryPage(c, path.Lottery)
		case path.Lottery:
			fmt.Printf("UID[%s] Lottery\n", user.ID)
			visitGroundMenu(c, path.GroundMenu, user.Schedules[num])
		case path.GroundMenu:
			fmt.Printf("UID[%s] GroundMenu\n", user.ID)
			schedule := user.Schedules[num]
			if len(schedule.Places) == 2 && schedule.Places[1] != "" {
				visitGroundNumMenu(c, path.GroundNumMenu, schedule)
			} else {
				visitApply(c, path.Apply, schedule)
			}
		case path.GroundNumMenu:
			fmt.Printf("UID[%s] GroundNumMenu\n", user.ID)
			visitApply(c, path.Apply, user.Schedules[num])
		case path.Complete:
			fmt.Printf("UID[%s] Complete!!\n", user.ID)
			if len(user.Schedules) > num+1 {
				num += 1
				visitMenuPage(c, path.Menu)
			}
		}
	})

	c.OnHTML("input[name=insLotJKey]", func(e *colly.HTMLElement) {
		if len(user.Schedules) <= num+1 {
			return
		}

		fmt.Printf("UID[%s] exists insLotJKey %s\n", user.ID, e.Request.URL.String())
		visitComplete(c, path.Complete, user.Schedules[num], e.Attr("value"))
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Printf("UID[%s] Request %s\n", user.ID, r.URL)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Printf("UID[%s] Response code %d\n", user.ID, r.StatusCode)
	})

	c.Visit(user.Url)
}

func visitMenuPage(c *colly.Collector, path string) {
	c.Post(path, map[string]string{
		"displayNo":       "plwac3000",
		"selectPpsPpsdCd": "100",
		"selectPpsCd":     "100130",
	})
}

func visitLotteryPage(c *colly.Collector, path string) {
	c.Post(path, map[string]string{
		"displayNo":     "plwad1000",
		"selectClassCd": "5060010",
	})
}

func visitGroundMenu(c *colly.Collector, path string, schedule Schedule) {
	c.Post(path, map[string]string{
		"displayNo":      "plwba1000",
		"selectBldGrpCd": schedule.Places[0],
	})
}

func visitGroundNumMenu(c *colly.Collector, path string, schedule Schedule) {
	// INFO: 複数のグラウンドがあった場合（第一運動場、第二運動場など）
	c.Post(path, map[string]string{
		"displayNo":       "plwba2000",
		"selectBldGrpCd":  schedule.Places[1],
		"selectUseNumber": "50", // ここの値は固定値か要確認
	})
}

func visitApply(c *colly.Collector, path string, schedule Schedule) {
	c.Post(path, map[string]string{
		"selectFieldCnt":  "1",
		"displayNo":       "plwba3000",
		"selectUseYMD":    schedule.Date,
		"selectTimeNum":   "1",
		"maxFieldCnt":     "1",
		"selectDispStime": schedule.Start,
		"selectDispEtime": schedule.End,
		"dummy":           "",
		"dummy2":          "",
	})
}

func visitComplete(c *colly.Collector, path string, schedule Schedule, insLotJKey string) {
	c.Post(path, map[string]string{
		"displayNo":       "plwca1000",
		"applyPepopleNum": schedule.People,
		"selectApplyNo":   "-1",
		"selectHopeNo":    "1",
		"selectPpsName":   "%8F%AD%94N%96%EC%8B%85%81i%8F%AC%81E%92%86%8Aw%90%B6%81j",
		"insLotJKey":      insLotJKey,
		"dummy":           "",
	})
}
