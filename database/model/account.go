package model

import (
	"github.com/google/uuid"
	"time"
)

type Account struct {
	ID        uuid.UUID `db:"id"`
	Name      string    `db:"name"`
	Balance   int64     `db:"balance"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
