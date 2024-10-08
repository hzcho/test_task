package model

type Building struct {
	ID     uint64 `json:"id"`
	Name   string `json:"name" binding:"required"`
	City   string `json:"city" binding:"required"`
	Year   int    `json:"year"`
	Floors int    `json:"floors"`
}
