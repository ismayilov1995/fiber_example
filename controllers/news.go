package controllers

import (
	"fiber_exp/models"

	"github.com/gofiber/fiber/v2"
)

type NewsController struct{}

func (n *NewsController) LoadAll(c *fiber.Ctx) error {
	news := new(models.News)
	return c.JSON(news.LoadAll())
}

func (n *NewsController) Load(c *fiber.Ctx) error {
	news := new(models.News)
	if news, err := news.Load(c.Params("slug")); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	} else {
		return c.JSON(news)
	}
}

func (n *NewsController) Create(c *fiber.Ctx) error {
	news := new(models.News)
	c.BodyParser(&news)
	return c.JSON(news.Create())
}

func (n *NewsController) Delete(c *fiber.Ctx) error {
	news := new(models.News)
	if err := news.Delete(c.Params("slug")); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "deleted successfuly"})
}
