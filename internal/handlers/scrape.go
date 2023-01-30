package handlers

import (
	"fmt"
	"ganso-kyamada/contact-net/internal/resources"

	"github.com/gocolly/colly"
)

func Reservation(path resources.Path, user resources.User, schedule resources.Schedule) {
	c := colly.NewCollector()
	c.OnHTML("input[name=loginJKey]", func(e *colly.HTMLElement) {
		fmt.Printf("UID[%s] Login...\n", user.ID)
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
			visitGroundMenu(c, path.GroundMenu, schedule)
		case path.GroundMenu:
			fmt.Printf("UID[%s] GroundMenu\n", user.ID)
			// INFO: 複数のグラウンドがあった場合（第一運動場、第二運動場など）
			if len(schedule.Places) > 1 {
				visitGroundNumMenu(c, path.GroundNumMenu, schedule)
			} else {
				visitApply(c, path.Apply, schedule)
			}
		case path.GroundNumMenu:
			fmt.Printf("UID[%s] GroundNumMenu\n", user.ID)
			visitApply(c, path.Apply, schedule)
		case path.Complete:
			fmt.Printf("UID[%s] Complete!!\n", user.ID)
		}
	})

	c.OnHTML("input[name=insLotJKey]", func(e *colly.HTMLElement) {
		fmt.Printf("UID[%s] exists insLotJKey %s\n", user.ID, e.Request.URL.String())
		visitComplete(c, path.Complete, schedule, e.Attr("value"))
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Printf("UID[%s] Request %s\n", user.ID, r.URL)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Printf("UID[%s] Response code %d\n", user.ID, r.StatusCode)
	})
	c.Visit(path.Url)
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

func visitGroundMenu(c *colly.Collector, path string, schedule resources.Schedule) {
	c.Post(path, map[string]string{
		"displayNo":      "plwba1000",
		"selectBldGrpCd": schedule.Places[0],
	})
}

func visitGroundNumMenu(c *colly.Collector, path string, schedule resources.Schedule) {
	c.Post(path, map[string]string{
		"displayNo":       "plwba2000",
		"selectBldGrpCd":  schedule.Places[1],
		"selectUseNumber": "50", // ここの値は固定値か要確認
	})
}

func visitApply(c *colly.Collector, path string, schedule resources.Schedule) {
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

func visitComplete(c *colly.Collector, path string, schedule resources.Schedule, insLotJKey string) {
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
