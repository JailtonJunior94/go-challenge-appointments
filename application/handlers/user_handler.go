package handlers

import (
	"github.com/gofiber/fiber/v2"

	"github.com/jailtonjunior94/go-challenge-appointments/domain/dtos"
	"github.com/jailtonjunior94/go-challenge-appointments/domain/entities"
	"github.com/jailtonjunior94/go-challenge-appointments/infrastructure/repositories"
)

type IUserHandler interface {
	CreateUser(c *fiber.Ctx) error
}

type UserHandler struct {
	UserRepository repositories.IUserRepository
}

func NewUserHandler(repository repositories.IUserRepository) IUserHandler {
	return &UserHandler{UserRepository: repository}
}

func (h UserHandler) CreateUser(c *fiber.Ctx) error {
	request := new(dtos.UserRequest)
	if err := c.BodyParser(request); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"message": "Unprocessable Entity"})
	}

	if err := request.IsValid(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	user := entities.User{}
	user.NewUser(request.Name, request.Cpf, request.Email, request.BirthDate)

	result, err := h.UserRepository.Add(&user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Ocorreu um erro inesperado"})
	}

	response := dtos.UserResponse{
		Id:        result.ID,
		Name:      result.Name,
		Cpf:       result.Cpf,
		Email:     result.Email,
		BirthDate: result.BirthDate,
		Active:    result.Active,
	}

	return c.Status(fiber.StatusCreated).JSON(response)
}
