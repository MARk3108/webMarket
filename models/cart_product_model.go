package models

type CartProduct struct {
	ID        uint `gorm:"primary_key;autoIncrement"`
	CartID    uint
	ProductID uint
	Amount    uint
}
