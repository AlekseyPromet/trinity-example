package models

import (
	"time"

	"github.com/google/uuid"
)

type OrderStatus string

const (
	OrderStatusOk  OrderStatus = "Created"
	OrderStatusUp  OrderStatus = "Updated"
	OrderStatusBad OrderStatus = "Not find"
	OrderStatusDel OrderStatus = "Deleted"
)

type ErrorOrder struct {
	Code  int    `json:"code"`
	Cause string `json:"cause"`
}

type OrderReq struct {
	Phone    string    `json:"phone"`
	Customer string    `json:"full_name"`
	Datetime time.Time `json:"datetime"`
	Goods    []string  `json:"goods"`
}

type OrderRes struct {
	ID       uuid.UUID   `json:"id"`
	Status   OrderStatus `json:"status"`
	Datetime time.Time   `json:"datetime"`
	Error    ErrorOrder  `json:"error"`
}
