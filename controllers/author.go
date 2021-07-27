package controllers

import (
	"fiber_exp/models"

	"github.com/gofiber/fiber/v2"
)

type AuthorController struct{}

func (a *AuthorController) LoadAll(c *fiber.Ctx) error {
	author := new(models.Author)
	return c.JSON(author.LoadAll())
}

func (a *AuthorController) Create(c *fiber.Ctx) error {
	author := new(models.Author)
	c.BodyParser(&author)
	if authorResponse, err := author.Create(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	} else {
		return c.JSON(authorResponse)
	}
}
