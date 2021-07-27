package models

import (
	"fiber_exp/database"

	"gorm.io/gorm"
)

type News struct {
	gorm.Model
	Title   string `gorm:"unique;not null;"`
	Content string `gorm:"not null"`
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
	n.Read = n.Read + 1
	database.DBConn.Save(&n)
	return n, nil
}

func (n *News) Create() (*News, error) {
	if res := database.DBConn.Create(&n); res.Error != nil {
		return nil, res.Error
	}
	return n, nil
}

func (n *News) Delete(id interface{}) error {
	if err := database.DBConn.First(&n, id).Error; err != nil {
		return err
	}
	database.DBConn.Delete(n)
	return nil
}
