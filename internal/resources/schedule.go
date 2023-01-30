package resources

import (
	"encoding/csv"
	"os"
)

type Schedule struct {
	Date   string   // 抽選日時
	Start  string   // 開始時間
	End    string   // 終了時間
	People string   // 人数
	Places []string // グラウンド番号
}

func GetSchedules() (schedules []Schedule, error error) {
	schedulesFile, err := os.Open("schedules.csv")
	if err != nil {
		return schedules, err
	}
	defer schedulesFile.Close()

	schedulesFileReader := csv.NewReader(schedulesFile)
	schedulesRows, err := schedulesFileReader.ReadAll()
	if err != nil {
		return schedules, err
	}

	for i, s := range schedulesRows {
		if i == 0 {
			continue
		}
		schedule := Schedule{
			Date:   s[0],
			Start:  s[1],
			End:    s[2],
			People: s[3],
		}
		for _, p := range s[4:] {
			if p != "" {
				schedule.Places = append(schedule.Places, p)
			}
		}
		schedules = append(schedules, schedule)
	}

	return schedules, nil
}
