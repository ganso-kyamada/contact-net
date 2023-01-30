package main

import (
	"fmt"
	"ganso-kyamada/contact-net/internal/handlers"
	"ganso-kyamada/contact-net/internal/resources"
	"os"
)

func main() {
	fmt.Println("START")
	url := os.Getenv("URL")
	if url == "" {
		fmt.Println("Please set enviroment \"URL\"")
		return
	}
	path := resources.GetPath(url)
	fmt.Println("INFO: GetPath")

	schedules, err := resources.GetSchedules()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("INFO: GetSchedules")
	fmt.Println(schedules)

	users, err := resources.GetUsers()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("INFO: GetUsers")
	fmt.Println(users)

	for _, user := range users {
		for _, schedule := range schedules {
			handlers.Reservation(path, user, schedule)
		}
	}
	fmt.Println("END")
}
