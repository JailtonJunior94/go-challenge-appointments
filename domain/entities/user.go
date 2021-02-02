package entities

import "time"

type User struct {
	Entity
	Name         string    `db:"Name"`
	Cpf          string    `db:"Cpf"`
	Email        string    `db:"Email"`
	BirthDate    time.Time `db:"BirthDate"`
	Appointments []Appointment
}

func (u *User) NewUser(name, cpf, email string, birthDate time.Time) {
	u.Name = name
	u.Cpf = cpf
	u.Email = email
	u.BirthDate = birthDate
	u.Entity.NewEntity()
}

func (u *User) Update(name string, birthDate time.Time) {
	u.Name = name
	u.BirthDate = birthDate
	u.Entity.ChangeUpdatedAt()
}
