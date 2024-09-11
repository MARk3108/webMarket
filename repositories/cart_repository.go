package repositories

import (
	"WebMarket/config"
	"WebMarket/models"
)

type CartRepository interface {
	Create(cart *models.Cart) error
	FindByUserID(userID uint) (*models.Cart, error)
	FindCartProduct(cartID, productID uint) (*models.CartProduct, error)
	AddProductToCart(cartProduct *models.CartProduct) error
	SaveCartProduct(cartProduct *models.CartProduct) error
}

type cartRepository struct{}

func NewCartRepository() CartRepository {
	return &cartRepository{}
}

func (r *cartRepository) Create(cart *models.Cart) error {
	return config.DB.Create(cart).Error
}

func (r *cartRepository) FindByUserID(userID uint) (*models.Cart, error) {
	var cart models.Cart
	err := config.DB.Preload("Products").Where("user_id = ?", userID).First(&cart).Error
	return &cart, err
}

func (r *cartRepository) FindCartProduct(cartID uint, productID uint) (*models.CartProduct, error) {
	var cartProduct models.CartProduct
	err := config.DB.Where("cart_id = ? AND product_id = ?", cartID, productID).First(&cartProduct).Error
	return &cartProduct, err
}

func (r *cartRepository) AddProductToCart(cartProduct *models.CartProduct) error {
	return config.DB.Create(cartProduct).Error
}

func (r *cartRepository) SaveCartProduct(cartProduct *models.CartProduct) error {
	return config.DB.Save(cartProduct).Error
}
