package models

import (
	"fiber_exp/database"

	"gorm.io/gorm"
)

type News struct {
	gorm.Model
	Title   string `gorm:"unique"`
	Content string
	Read    int
}

func (n *News) LoadAll() *[]News {
	var allNews []News
	database.DBConn.Find(&allNews)
	return &allNews
}

func (n *News) Load(id interface{}) (*News, error) {
	if err := database.DBConn.First(&n, id).Error; err != nil {
		return nil, err
	}
	return n, nil
}

func (n *News) Create() *News {
	database.DBConn.Create(&n)
	return n
}

func (n *News) Delete(id interface{}) error {
	if err := database.DBConn.First(&n, id).Error; err != nil {
		return err
	}
	database.DBConn.Delete(n)
	return nil
}
