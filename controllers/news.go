package controllers

import (
	"fiber_exp/models"

	"github.com/gofiber/fiber/v2"
)

type NewsController struct{}

func (n *NewsController) LoadAll(c *fiber.Ctx) error {
	a := new(models.Author)
	a.FullName = "Ismayil Ismayilov"
	return c.JSON(a.Create())
}

func (n *NewsController) Load(c *fiber.Ctx) error {
	return c.SendString("Book: " + c.Params("slug"))
}

func (n *NewsController) Create(c *fiber.Ctx) error {
	return c.SendString("Book created")
}

func (n *NewsController) Delete(c *fiber.Ctx) error {
	return c.SendString("Book removed: " + c.Params("slug"))
}
