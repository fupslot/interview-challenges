package main

import (
	"fmt"
	"os"

	"github.com/fupslot/go-rest/pkg/database"
	"github.com/fupslot/go-rest/pkg/model"
)

func Seeding(db *database.Database) {
	db.CreateUser("John Doe", 32)
	db.CreateUser("Alice Doe", 45)
	db.CreateUser("Bob Lazar", 38)
	db.CreateUser("Larry Moor", 64)
}

func main() {
	db, err := database.InitDb()
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
	// db.DB.AutoMigrate(model.User{})
	// db.DB.AutoMigrate(model.Comment{})

	db.DB.Create(&model.User{
		Username: "John Doe", Comments: []model.Comment{
			{Text: "Hello World"},
		},
	})

	// user := model.User{Username: "Oleg"}
	// db.DB.Create(&user)

	// fmt.Printf("user id: %d", user.ID)

	// db.DB.Delete(&model.User{}, 1)

	// Seeding(db)
	// user := db.FindUser("John Doe")

	// if user != nil {
	// 	fmt.Printf("user id: %d", user.ID)
	// 	if user.Age == 0 {
	// 		db.DB.Model(&user).Update("Age", 25)
	// 	}
	// }
}
