package services

import (
	"WebMarket/models"
	"WebMarket/repositories"
	"errors"
)

type CartService interface {
	GetCart(userID uint) (*models.Cart, error)
	AddToCart(userID uint, productID uint, productAmount uint) error
}

type cartService struct {
	cartRepo    repositories.CartRepository
	productRepo repositories.ProductRepository
}

func NewCartService(cartRepo repositories.CartRepository, productRepo repositories.ProductRepository) CartService {
	return &cartService{cartRepo: cartRepo, productRepo: productRepo}
}

func (s *cartService) GetCart(userID uint) (*models.Cart, error) {
	return s.cartRepo.FindByUserID(userID)
}

func (s *cartService) AddToCart(userID uint, productID uint, productAmount uint) error {
	cart, err := s.cartRepo.FindByUserID(userID)
	if cart.UserID != userID {
		cart = &models.Cart{UserID: userID}
		if err := s.cartRepo.Create(cart); err != nil {
			return err
		}
	} else {
		if err != nil {
			return err
		}
	}

	product, err := s.productRepo.FindByID(productID)
	if err != nil {
		return err
	}

	if product.Amount < productAmount {
		return errors.New("requested amount is more than available in stock")
	}

	cartProduct, err := s.cartRepo.FindCartProduct(cart.ID, productID)
	if err != nil {
		cartProduct = &models.CartProduct{
			CartID:    cart.ID,
			ProductID: product.ID,
			Amount:    productAmount,
		}
		if err := s.cartRepo.AddProductToCart(cartProduct); err != nil {
			return err
		}
	} else {
		cartProduct.Amount += productAmount
		if err := s.cartRepo.SaveCartProduct(cartProduct); err != nil {
			return err
		}
	}
	newAmount := product.Amount - productAmount
	if err := s.productRepo.UpdateAmount(newAmount, productID); err != nil {
		return err
	}

	return nil
}
