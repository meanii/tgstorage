package routers

import (
	"github.com/gofiber/fiber/v2"

	"github.com/meanii/tgstorage/controllers"
)

type Routers struct {
	app  *fiber.App
	root fiber.Router
}

func (r *Routers) Init(app *fiber.App) {
	r.app = app
	r.root = r.app.Group("/tgstorage")
	r.welcome() // Welcome message
	r.AppController()
}

// AppController is a controller for the root path
func (r *Routers) AppController() {
	app := controllers.Controller{}
	app.Init(r.root)
	app.Upload()
}

func (r *Routers) welcome() fiber.Router {
	return r.root.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to tgstorage 🦉")
	})
}
