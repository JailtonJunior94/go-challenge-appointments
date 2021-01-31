package database

import (
	"github.com/jailtonjunior94/go-challenge-appointments/infrastructure/environments"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/jmoiron/sqlx"
)

type SqlConnection struct {
	Db *sqlx.DB
}

func NewConnection() (SqlConnection, error) {
	db, err := sqlx.Connect("sqlserver", environments.SqlConnectionString)
	if err != nil {
		return SqlConnection{}, err
	}

	if err = db.Ping(); err != nil {
		return SqlConnection{}, err
	}

	return SqlConnection{Db: db}, nil
}
