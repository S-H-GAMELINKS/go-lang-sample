package main

import (
	"os"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	ID int
	Name string
	Posts []Post
}

type Post struct {
	ID int
	Name string
	UserID uint
}

func main() {
	_, err := os.Stat("./test.db")
	if err == nil {
		fmt.Println("Remove database")
		os.Remove("./test.db")
	}
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect to database")
	}
	defer db.Close()

	db.AutoMigrate(&User{}, &Post{})

	db.Create(&User{ID: 1, Name: "matz"})
	db.Create(&Post{ID: 1, Name: "Ruby", UserID: 1})

	var users []User
	var posts []Post
	var updatePosts []Post
	var updateUsersAfterPosts []Post

	db.Preload("Posts").Find(&users)
	db.Find(&posts)
	for _, user := range users {
		if user.Name != "" {
			fmt.Println(user.Name)
			fmt.Println(user.Posts)

			for _, post := range posts {
				fmt.Println(post.Name)
				db.Model(&post).Select("name").Update(Post{Name: "Go"})
			}
			
			db.Find(&updatePosts)

			for _, post := range updatePosts {
				fmt.Println(post.Name)
			}

			db.Model(&user).Select("name").Update(User{Name: "nobu"})

			db.Find(&updateUsersAfterPosts)

			for _, post := range updateUsersAfterPosts {
				fmt.Println(post.Name)
			}
		} else {
			fmt.Println("user name is empty")
		}
	}
}
