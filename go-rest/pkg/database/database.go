package database

import (
	"fmt"

	"github.com/fupslot/go-rest/pkg/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const dsn = "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable"

type Database struct {
	DB *gorm.DB
}

func (p *Database) init() error {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	db.AutoMigrate(&model.User{})

	p.DB = db
	return nil
}

func (p *Database) CreateUser(username string, age uint8) {
	err := p.DB.Create(&model.User{Username: username, Age: age}).Error
	if err != nil {
		fmt.Println(err)
	}
}

func (p *Database) FindUser(username string) *model.User {
	var user *model.User
	p.DB.Find(&user, "username = ?", username)

	return user
}

func InitDb() (*Database, error) {
	d := &Database{}
	err := d.init()
	if err != nil {
		return nil, err
	}
	return d, nil
}
