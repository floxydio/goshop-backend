package routes

import (
	"goshop/database"
	"goshop/internal/handler"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Routes(routes *echo.Echo) {

	// Middleware
	routes.Use(middleware.CORS())
	routes.Use(middleware.Logger())

	// Controller
	productController := handler.ProductControllerInit(database.DbClient)
	userController := handler.UserControllerInit(database.DbClient)

	// Routes
	routes.GET("/v1/product", productController.FindDataProduct)
	routes.POST("/v1/product", productController.CreateDataProduct)
	routes.PUT("/v1/product/:id", productController.UpdateDataProduct)

	// User
	routes.POST("/v1/signup", userController.SignUp)
	routes.POST("/v1/signin", userController.SignIn)

	// Kurir
	
}
