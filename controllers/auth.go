package controllers

import "github.com/gofiber/fiber/v2"

type AuthController struct{}

func (a *AuthController) Login(c *fiber.Ctx) error {
	return c.SendString("welcome asshole")
}

func (a *AuthController) Signup(c *fiber.Ctx) error {
	return c.SendString("welcome to gang")
}
