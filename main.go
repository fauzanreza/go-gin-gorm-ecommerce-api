package main

import (
	"ecommerce_api/controllers"
	"ecommerce_api/inits"
	"ecommerce_api/middlewares"

	"github.com/gin-gonic/gin"
)

func init() {
	inits.LoadEnv()
	inits.DBInit()
}

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello Ojan!",
		})
	})

	// Product Routes
	r.POST("/products", controllers.CreateProduct)
	r.GET("/products", controllers.GetProducts)
	r.GET("/products/:id", controllers.GetProduct)
	r.PUT("/products/:id", controllers.UpdateProduct)
	r.DELETE("/products/:id", controllers.DeleteProduct)

	// Category Routes
	r.POST("/categories", middlewares.RequireAuth, controllers.CreateCategory)
	r.GET("/categories", middlewares.RequireAuth, controllers.GetCategorys)
	r.GET("/categories/:id", middlewares.RequireAuth, controllers.GetCategory)
	r.PUT("/categories/:id", middlewares.RequireAuth, controllers.UpdateCategory)
	r.DELETE("/categories/:id", middlewares.RequireAuth, controllers.DeleteCategory)

	//User Routes
	r.POST("/user", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.GET("/users", middlewares.RequireAuth, controllers.GetUsers)
	r.GET("/logout", controllers.Logout)

	// Cart Routes
	r.POST("/cart", middlewares.RequireAuth, controllers.CreateCart)
	r.GET("/cart", middlewares.RequireAuth, controllers.GetCarts)
	r.GET("/cart/:id", middlewares.RequireAuth, controllers.GetCart)
	r.DELETE("/cart/:id", middlewares.RequireAuth, controllers.DeleteCart)

	r.Run()
}
