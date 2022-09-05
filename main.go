package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

type Schedule struct {
	Date   string
	Start  string
	End    string
	People string
	Places []string
}

type User struct {
	ID         string
	Password   string
	SecurityNo string
	Url        string
	Schedules  []Schedule
}

func main() {
	url := os.Getenv("URL")

	schedulesFile, err := os.Open("schedules.csv")
	if err != nil {
		fmt.Println(err)
	}
	defer schedulesFile.Close()

	schedulesFileReader := csv.NewReader(schedulesFile)
	schedulesRows, err := schedulesFileReader.ReadAll()
	if err != nil {
		fmt.Println(err)
	}

	usersFile, err := os.Open("users.csv")
	if err != nil {
		fmt.Println(err)
	}
	defer usersFile.Close()

	usersFileReader := csv.NewReader(usersFile)
	usersRows, err := usersFileReader.ReadAll()
	if err != nil {
		fmt.Println(err)
	}

	for i, u := range usersRows {
		if i == 0 {
			continue
		}

		var schedules []Schedule
		for j, s := range schedulesRows {
			if j == 0 {
				continue
			}
			schedule := Schedule{
				Date:   s[0],
				Start:  s[1],
				End:    s[2],
				People: s[3],
				Places: s[4:],
			}
			schedules = append(schedules, schedule)
		}

		user := User{
			ID:         u[0],
			Password:   u[1],
			SecurityNo: u[2],
			Url:        url,
			Schedules:  schedules,
		}

		Reservation(user)
	}
}
