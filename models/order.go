package models

import (
	"time"

	"github.com/google/uuid"
)

type Order struct {
	ID       uuid.UUID `json:"id"`
	Phone    string    `json:"phone"`
	Customer string    `json:"full_name"`
	Datetime time.Time `json:"datetime"`
}
