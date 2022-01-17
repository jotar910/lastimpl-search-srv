package models

type CommonList struct {
	TotalItems int `json:"totalItems"`
	TotalPages int `json:"totalPages"`
	Count      int `json:"count"`
	Page       int `json:"page"`
}
