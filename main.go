package main

import (
	"fiber_exp/database"
	"fiber_exp/models"
	"fiber_exp/routers"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func initDB() {
	var err error
	database.DBConn, err = gorm.Open(sqlite.Open("book.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("database connection successfuly")
	database.DBConn.AutoMigrate(&models.Book{})
	fmt.Println("database migrated")
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
