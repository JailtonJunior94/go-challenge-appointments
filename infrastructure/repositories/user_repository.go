package repositories

import (
	"database/sql"

	"github.com/jailtonjunior94/go-challenge-appointments/domain/entities"
	"github.com/jmoiron/sqlx"
)

type IUserRepository interface {
	Add(u *entities.User) (user *entities.User, err error)
	Get() (users []entities.User, err error)
}

type UserRepository struct {
	Db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) IUserRepository {
	return &UserRepository{Db: db}
}

func (r UserRepository) Add(u *entities.User) (user *entities.User, err error) {
	query := `INSERT INTO
					dbo.[User]
				VALUES
					(@id, @name, @cpf, @email, @birthDate, @createdAt, @updatedAt, @active);`

	s, err := r.Db.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer s.Close()

	result, err := s.Exec(sql.Named("id", u.ID),
		sql.Named("name", u.Name),
		sql.Named("cpf", u.Cpf),
		sql.Named("email", u.Email),
		sql.Named("birthDate", u.BirthDate),
		sql.Named("createdAt", u.CreatedAt),
		sql.Named("updatedAt", u.UpdatedAt),
		sql.Named("active", u.Active))

	if err != nil {
		return nil, err
	}

	rows, err := result.RowsAffected()
	if rows == 0 {
		return nil, err
	}

	return u, nil
}

func (r UserRepository) Get() (users []entities.User, err error) {
	query := `SELECT
				CAST([Id] AS CHAR(36)) Id,
				[Name],
				[Cpf],
				[Email],
				[BirthDate],
				[CreatedAt],
				[UpdatedAt],
				[Active]
			FROM
				[ChallengeAppointments].[dbo].[User]`

	if err := r.Db.Select(&users, query); err != nil {
		return nil, err
	}

	return users, nil
}
