package controllers

import (
	"WebMarket/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

var cartService services.CartService

func InitializeCartController(service services.CartService) {
	cartService = service
}

func GetCart(c *gin.Context) {
	userID := c.MustGet("user_id").(uint)
	cart, err := cartService.GetCart(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, cart)
}

func AddToCart(c *gin.Context) {
	userID := c.MustGet("user_id").(uint)
	var requestData struct {
		ProductID     uint `json:"product_id"`
		ProductAmount uint `json:"product_amount"`
	}
	if err := c.BindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := cartService.AddToCart(userID, requestData.ProductID, requestData.ProductAmount); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "Product added to cart"})
}
