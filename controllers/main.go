package controllers

import (
	"github.com/gofiber/fiber/v2"
)

type Controller struct {
	router fiber.Router
}

func (c *Controller) Init(router fiber.Router) {
	c.router = router
}

func (c *Controller) Upload() fiber.Router {
	return c.router.Get("/upload", func(c *fiber.Ctx) error {
		return c.SendString("Upload")
	})
}
