package server

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"os"
)

func Server() {
	app := fiber.New()

	app.Use(cors.New()) // Enable CORS

	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	})) // Enable logging

	port := fmt.Sprintf(":%s", os.Getenv("PORT"))
	log.Infof("Server is running on port %s", port)
	err := app.Listen(port)
	if err != nil {
		log.Fatalf("Error while listening to port %s: %s", port, err.Error())
	}

}
