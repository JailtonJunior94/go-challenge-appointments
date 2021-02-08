package dtos

import "time"

type UserResponse struct {
	Id           string                `json:"id"`
	Name         string                `json:"name"`
	Cpf          string                `json:"cpf"`
	Email        string                `json:"email"`
	BirthDate    time.Time             `json:"birthDate"`
	Active       bool                  `json:"active"`
	Appointments []AppointmentResponse `json:"appointments"`
}
