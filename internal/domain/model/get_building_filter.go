package model

type Filter struct {
    City   string `json:"city,omitempty"` 
    Year   int    `json:"year,omitempty"`
    Floors int    `json:"floors,omitempty"`
}
