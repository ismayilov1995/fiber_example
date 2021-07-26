package models

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open(sqlite.Open("author.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("author connection successfuly")
	db.AutoMigrate(&Author{})
	fmt.Println("author migrated")
}

type Author struct {
	gorm.Model
	FullName string `json:"fullName"`
}

func (a *Author) Create() *Author {
	db.Create(&a)
	return a
}
