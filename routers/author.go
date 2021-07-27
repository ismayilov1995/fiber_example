package routers

import (
	"fiber_exp/controllers"

	"github.com/gofiber/fiber/v2"
)

var a = new(controllers.AuthorController)

func AuthorRouter(r fiber.Router) {
	r.Get("/", a.LoadAll)
	r.Post("create", a.Create)
}
