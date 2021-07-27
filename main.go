package main

import (
	"fiber_exp/database"
	"fiber_exp/models"
	"fiber_exp/routers"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func initDB() {
	var err error
	dsn := "host=localhost user=postgres password=7090698 dbname=vnews port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	database.DBConn, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("database connection successfuly")
	database.DBConn.AutoMigrate(&models.News{})
}

func main() {
	app := fiber.New()
	initDB()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	routers.ApiRouters(app)

	app.Listen(":3000")
}
