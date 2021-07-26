package routers

import (
	"fiber_exp/controllers"

	"github.com/gofiber/fiber/v2"
)

var b = new(controllers.BookController)

func BookRouter(r fiber.Router) {
	r.Get("/", b.GetBooks)
	r.Get("/:id", b.GetBook)
	r.Post("/", b.NewBook)
	r.Delete("/:id", b.DeleteBook)
}
