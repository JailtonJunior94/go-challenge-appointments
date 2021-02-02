package entities

import (
	"database/sql"
	"time"

	uuid "github.com/satori/go.uuid"
)

type Entity struct {
	ID        string       `db:"Id"`
	CreatedAt time.Time    `db:"CreatedAt"`
	UpdatedAt sql.NullTime `db:"UpdatedAt"`
	Active    bool         `db:"Active"`
}

func (e *Entity) NewEntity() {
	e.ID = uuid.NewV4().String()
	e.CreatedAt = time.Now()
	e.Active = true
}

func (e *Entity) ChangeUpdatedAt() {
	e.UpdatedAt.Time = time.Now()
}

func (e *Entity) ChangeStatus(status bool) {
	e.Active = status
}
