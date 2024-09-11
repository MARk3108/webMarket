package models

import (
	"github.com/jinzhu/gorm"
)

type Cart struct {
	gorm.Model
	UserID   uint
	Products []CartProduct `gorm:"foreignKey:CartID"`
}
