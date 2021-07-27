package routers

import (
	"github.com/gofiber/fiber/v2"
)

func ApiRouters(app fiber.Router) {
	a := app.Group("api/v1")
	NewsRouter(a.Group("news"))
	AuthorRouter(a.Group("author"))
	BookRouter(a.Group("book"))
}
