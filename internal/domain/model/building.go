package model

type Building struct {
	ID     uint64
	Name   string `json:"name" binding:"required"`
	City   string `json:"city" binding:"required"`
	Year   int    `json:"year" binding:"required"`
	Floors int    `json:"floors" validate:"required"`
}
