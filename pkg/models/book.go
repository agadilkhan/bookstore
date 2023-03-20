package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Title    string  `json:"title"`
	Descript string  `json:"descript"`
	Cost     float64 `json:"cost"`
}
