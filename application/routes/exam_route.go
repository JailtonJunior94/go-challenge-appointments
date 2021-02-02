package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jailtonjunior94/go-challenge-appointments/application/handlers"
	"github.com/jailtonjunior94/go-challenge-appointments/infrastructure/facades"
)

func RegisterExamRouter(router fiber.Router) {
	f := facades.NewExamFacade()
	h := handlers.NewExamHandler(f)

	router.Get("/exams", h.GetExams)
}
