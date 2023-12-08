package main

import (
	"ecommerce_api/controllers"
	"ecommerce_api/inits"

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
	r.POST("/categories", controllers.CreateCategory)
	r.GET("/categories", controllers.GetCategorys)
	r.GET("/categories/:id", controllers.GetCategory)
	r.PUT("/categories/:id", controllers.UpdateCategory)
	r.DELETE("/categories/:id", controllers.DeleteCategory)

	r.Run()
}
