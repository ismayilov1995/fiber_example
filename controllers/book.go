package controllers

import (
	"fiber_exp/database"
	"fiber_exp/models"

	"github.com/gofiber/fiber/v2"
)

type BookController struct{}

func (b *BookController) GetBooks(c *fiber.Ctx) error {
	db := database.DBConn
	var books []models.Book
	db.Find(&books)
	return c.JSON(books)
}

func (b *BookController) GetBook(c *fiber.Ctx) error {
	db := database.DBConn
	var book models.Book
	db.Find(&book, c.Params(":id"))
	return c.JSON(book)
}

func (b *BookController) NewBook(c *fiber.Ctx) error {
	db := database.DBConn
	var book models.Book
	if err := c.BodyParser(&book); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}
	db.Create(&book)
	return c.JSON(book)
}

func (b *BookController) DeleteBook(c *fiber.Ctx) error {
	db := database.DBConn
	var book models.Book
	db.First(&book, c.Params("id"))
	if book.Title == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "There is no book contains that id"})
	}
	db.Delete(&book)
	return c.JSON(fiber.Map{"message": "Selected book is removed"})
}
