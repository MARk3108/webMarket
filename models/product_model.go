package models

import (
	"github.com/jinzhu/gorm"
)

type Product struct {
	gorm.Model
	Name   string
	Price  float64
	Amount uint
}
