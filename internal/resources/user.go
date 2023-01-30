package resources

import (
	"encoding/csv"
	"os"
)

type User struct {
	ID         string // ID
	Password   string // パスワード
	SecurityNo string // セキュリティNo.
}

func GetUsers() (users []User, error error) {
	usersFile, err := os.Open("users.csv")
	if err != nil {
		return users, err
	}
	defer usersFile.Close()

	usersFileReader := csv.NewReader(usersFile)
	usersRows, err := usersFileReader.ReadAll()
	if err != nil {
		return users, err
	}

	for i, u := range usersRows {
		if i == 0 {
			continue
		}
		user := User{
			ID:         u[0],
			Password:   u[1],
			SecurityNo: u[2],
		}
		users = append(users, user)
	}
	return users, nil
}
