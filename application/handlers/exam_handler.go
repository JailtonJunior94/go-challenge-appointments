package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jailtonjunior94/go-challenge-appointments/infrastructure/facades"
)

type IExamHandler interface {
	GetExams(c *fiber.Ctx) error
}

type ExamHandler struct {
	ExamFacade facades.IExamFacade
}

func NewExamHandler(facade facades.IExamFacade) IExamHandler {
	return &ExamHandler{ExamFacade: facade}
}

func (e ExamHandler) GetExams(c *fiber.Ctx) error {
	result, err := e.ExamFacade.GetExams()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Ocorreu um erro inesperado"})
	}

	return c.Status(fiber.StatusOK).JSON(result)
}
