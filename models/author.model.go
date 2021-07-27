package models

import (
	"fiber_exp/database"

	"gorm.io/gorm"
)

type Author struct {
	gorm.Model
	FullName string `gorm:"not null;"`
	Email    string `gorm:"unique;not null;"`
	Avatar   string
	Raiting  int
}

func (a *Author) Create() (*Author, error) {
	if res := database.DBConn.Create(&a); res.Error != nil {
		return nil, res.Error
	}
	return a, nil
}

func (a *Author) LoadAll() *[]Author {
	var authors *[]Author
	database.DBConn.Find(&authors)
	return authors
}
