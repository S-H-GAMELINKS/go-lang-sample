package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	ID int
	Name string
}

func main() {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect to database")
	}
	defer db.Close()

	db.Create(&User{ID: 1, Name: "matz"})
	db.Create(&User{ID: 2})

	var users []User

	db.Find(&users)
	for _, user := range users {
		if user.Name != "" {
			fmt.Println(user.Name)
		} else {
			fmt.Println("user name is empty")
		}
	}
	
}
