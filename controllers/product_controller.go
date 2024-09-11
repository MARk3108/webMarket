package controllers

import (
	"WebMarket/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

var productService services.ProductService

func InitializeProductController(service services.ProductService) {
	productService = service
}

func AddProducts(c *gin.Context) {
	var requestData struct {
		Name   string  `json:"name"`
		Price  float64 `json:"price"`
		Amount uint    `json:"amount"`
	}

	if err := c.BindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	err := productService.AddProducts(requestData.Name, requestData.Price, requestData.Amount)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, "Success")
}

func GetProducts(c *gin.Context) {
	products, err := productService.GetProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, products)
}
