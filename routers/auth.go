package routers

import (
	"fiber_exp/controllers"

	"github.com/gofiber/fiber/v2"
)

var a = new(controllers.AuthController)

func AuthRouter(r fiber.Router) {
	r.Post("/login", a.Login)
	r.Post("/signup", a.Signup)
}
