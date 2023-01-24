package models

import (
	"time"

	"github.com/google/uuid"
)

type OrderReq struct {
	Phone    string    `json:"phone"`
	Customer string    `json:"full_name"`
	Datetime time.Time `json:"datetime"`
}

type OrderRes struct {
	ID       uuid.UUID `json:"id"`
	Phone    string    `json:"phone"`
	Customer string    `json:"full_name"`
	Datetime time.Time `json:"datetime"`
}
