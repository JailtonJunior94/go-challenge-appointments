package main

import (
	"fmt"
	"log"

	"github.com/jailtonjunior94/go-challenge-appointments/application/routes"
	"github.com/jailtonjunior94/go-challenge-appointments/infrastructure/environments"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	environments.NewApiConfig()

	app := fiber.New()

	app.Use(cors.New())
	app.Use(logger.New())

	routes.RegisterRoutes(app)

	fmt.Printf("🚀 API is running on http://localhost:%d", environments.Port)
	log.Fatal(app.Listen(fmt.Sprintf(":%v", environments.Port)))
}
