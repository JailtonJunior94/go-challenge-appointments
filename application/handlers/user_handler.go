package handlers

import (
	"github.com/jailtonjunior94/go-challenge-appointments/domain/dtos"
	"github.com/jailtonjunior94/go-challenge-appointments/domain/entities"
	"github.com/jailtonjunior94/go-challenge-appointments/infrastructure/repositories"

	"github.com/gofiber/fiber/v2"
)

type IUserHandler interface {
	CreateUser(c *fiber.Ctx) error
	GetUsers(c *fiber.Ctx) error
}

type UserHandler struct {
	UserRepository repositories.IUserRepository
}

func NewUserHandler(repository repositories.IUserRepository) IUserHandler {
	return &UserHandler{UserRepository: repository}
}

// Users godoc
// @Summary Users
// @Description Users
// @Accept  json
// @Produce json
// @Tags Users
// @Param UserRequest body dtos.UserRequest true "UserRequest"
// @Success 200 {object} dtos.UserResponse
// @Failure 400,404 {object} dtos.UserResponse
// @Failure 500 {object} dtos.UserResponse
// @Failure default {object} dtos.UserResponse
// @Router /api/v1/users [post]
func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
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

// Users godoc
// @Summary Users
// @Description Users
// @Accept  json
// @Produce  json
// @Tags Users
// @Success 200 {object} dtos.UserResponse
// @Failure 400,404 {object} dtos.UserResponse
// @Failure 500 {object} dtos.UserResponse
// @Failure default {object} dtos.UserResponse
// @Router /api/v1/users [get]
func (h *UserHandler) GetUsers(c *fiber.Ctx) error {
	users, err := h.UserRepository.Get()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Ocorreu um erro inesperado"})
	}

	return c.Status(fiber.StatusOK).JSON(users)
}
