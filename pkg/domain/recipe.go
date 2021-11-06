package domain

import (
	"github.com/gofrs/uuid"
)

type Recipe struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}
