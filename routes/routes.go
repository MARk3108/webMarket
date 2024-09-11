package routes

import (
	"WebMarket/controllers"
	"WebMarket/middlewares"
	"WebMarket/repositories"
	"WebMarket/services"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	// Initialize repositories
	userRepo := repositories.NewUserRepository()
	cartRepo := repositories.NewCartRepository()
	productRepo := repositories.NewProductRepository()

	// Initialize services
	authService := services.NewAuthService(userRepo)
	cartService := services.NewCartService(cartRepo, productRepo)
	productService := services.NewProductService(productRepo)

	// Initialize controllers
	controllers.InitializeAuthController(authService)
	controllers.InitializeCartController(cartService)
	controllers.InitializeProductController(productService)

	api := r.Group("/api")
	{
		api.POST("/register", controllers.RegisterUser)
		api.POST("/login", controllers.LoginUser)
		api.GET("/products", controllers.GetProducts)
		api.POST("/products", middlewares.AuthMiddleware(), middlewares.AdminMiddleware(), controllers.AddProducts)
		api.GET("/cart", middlewares.AuthMiddleware(), controllers.GetCart)
		api.POST("/cart", middlewares.AuthMiddleware(), controllers.AddToCart)
	}
}
