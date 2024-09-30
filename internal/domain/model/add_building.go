package model

type AddBuilding struct {
	Name   string `json:"name" binding:"required"`
	City   string `json:"city" binding:"required"`
	Year   int    `json:"year"`
	Floors int    `json:"floors"`
}
