package routes

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/jailtonjunior94/go-challenge-appointments/application/handlers"
	"github.com/jailtonjunior94/go-challenge-appointments/infrastructure/database"
	"github.com/jailtonjunior94/go-challenge-appointments/infrastructure/repositories"
)

func RegisterUserRouter(router fiber.Router) {
	sql, err := database.NewConnection()
	if err != nil {
		log.Fatal(err)
	}

	r := repositories.NewUserRepository(sql.Db)
	h := handlers.NewUserHandler(r)

	router.Post("/users", h.CreateUser)
	router.Get("/users", h.GetUsers)
}
