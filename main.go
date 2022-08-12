package main

import (
	"fmt"
	"os"
	"flag"
)

func main() {
	uid := flag.String("i", "userid", "login user id")
	pass := flag.String("p", "password", "login password")
	url := os.Getenv("URL")
	flag.Parse()

	fmt.Printf("UserId: %s, Password: %s, URL: %s\n", *uid, *pass, url)
	Reservation(*uid, *pass, url)
}
