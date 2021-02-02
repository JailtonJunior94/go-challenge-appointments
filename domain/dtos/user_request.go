package dtos

import (
	"errors"
	"time"
)

type UserRequest struct {
	Name      string    `json:"name"`
	Cpf       string    `json:"cpf"`
	Email     string    `json:"email"`
	BirthDate time.Time `json:"birthDate"`
}

func (user *UserRequest) IsValid() error {
	if user.Name == "" {
		return errors.New("O Nome é obrigatório")
	}

	if user.Cpf == "" {
		return errors.New("O CPF é obrigatório")
	}

	if user.Email == "" {
		return errors.New("O E-mail é obrigatório")
	}

	return nil
}
