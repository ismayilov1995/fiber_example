package models

import (
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
}

type Author struct {
	gorm.Model
	FullName string `json:"fullName"`
}

func (a *Author) Create() *Author {
	db.Create(&a)
	return a
}
