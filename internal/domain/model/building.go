package model

type Building struct {
	ID     uint64 `json:"id"`
	Name   string `json:"name" validate:"required"`
	City   string `json:"city" validate:"required"`
	Year   int    `json:"year" validate:"required"`
	Floors int    `json:"floors" validate:"required"`
}
