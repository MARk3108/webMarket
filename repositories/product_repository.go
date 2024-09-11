package repositories

import (
	"WebMarket/config"
	"WebMarket/models"
)

type ProductRepository interface {
	FindAll() ([]models.Product, error)
	FindByID(id uint) (*models.Product, error)
	Create(product *models.Product) error
	UpdateAmount(newAmount, id uint) error
}

type productRepository struct{}

func NewProductRepository() ProductRepository {
	return &productRepository{}
}

func (r *productRepository) FindAll() ([]models.Product, error) {
	var products []models.Product
	err := config.DB.Find(&products).Error
	return products, err
}

func (r *productRepository) FindByID(id uint) (*models.Product, error) {
	var product models.Product
	err := config.DB.First(&product, id).Error
	return &product, err
}

func (r *productRepository) Create(product *models.Product) error {
	return config.DB.Create(product).Error
}

func (r *productRepository) UpdateAmount(newAmount, productId uint) error {
	var product models.Product
	err := config.DB.First(&product, productId).Error
	if err != nil {
		return err
	}
	product.Amount = newAmount
	result := config.DB.Save(&product)
	return result.Error
}
