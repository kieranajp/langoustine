package domain

import (
	"github.com/gofrs/uuid"
)

type Recipe struct {
	ID   uuid.UUID `json:"id" db:"uuid"`
	Name string    `json:"name"`
}
