package models

import "github.com/google/uuid"

type Good struct {
	ID   uuid.UUID `json:"id"`
	Code string    `json:"code"`
}
