package models

type StatusRequest struct{}

type StatusResponse struct {
	Status string `json:"status"`
}
