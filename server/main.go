package server

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"github.com/meanii/tgstorage/configs"
	"github.com/meanii/tgstorage/routers"
)

func Server() {
	app := fiber.New()

	app.Use(cors.New())

	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))

	// Initialize routers
	router := routers.Routers{}
	router.Init(app)

	port := fmt.Sprintf(":%s", configs.EnvConfig.Port)
	log.Infof("Server is running on http://0.0.0.0:%s", port)
	err := app.Listen(port)
	if err != nil {
		log.Fatalf("Error while listening to port %s: %s", port, err.Error())
	}
}
