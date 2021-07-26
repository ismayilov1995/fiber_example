package routers

import (
	"fiber_exp/controllers"

	"github.com/gofiber/fiber/v2"
)

var n = new(controllers.NewsController)

func NewsRouter(r fiber.Router) {
	r.Get("/", n.LoadAll)
	r.Get("create", n.Create)
	r.Get(":slug", n.Load)
	r.Get("delete/:slug", n.Delete)
}
