package services

import (
	"WebMarket/models"
	"WebMarket/repositories"
)

type ProductService interface {
	GetProducts() ([]models.Product, error)
	AddProducts(name string, price float64, amount uint) error
}

type productService struct {
	productRepo repositories.ProductRepository
}

func NewProductService(productRepo repositories.ProductRepository) ProductService {
	return &productService{productRepo: productRepo}
}

func (s *productService) GetProducts() ([]models.Product, error) {
	return s.productRepo.FindAll()
}

func (s *productService) AddProducts(name string, price float64, amount uint) error {
	product := &models.Product{Name: name, Price: float64(price), Amount: amount}
	return s.productRepo.Create(product)
}
