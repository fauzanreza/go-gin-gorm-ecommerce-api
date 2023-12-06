package controllers

import (
	"ecommerce_api/inits"
	"ecommerce_api/models"
	"fmt"

	"github.com/gin-gonic/gin"
)

func CreateProduct(ctx *gin.Context) {

	var body struct {
		Name        string
		Price       int
		Description string
		Stock       int
		Image_Url   string
	}

	ctx.BindJSON(&body)

	product := models.Product{Name: body.Name, Price: body.Price, Description: body.Description, Stock: body.Stock, Image_Url: body.Image_Url}

	fmt.Println(product)
	result := inits.DB.Create(&product)
	if result.Error != nil {
		ctx.JSON(500, gin.H{"error": result.Error})
		return
	}

	ctx.JSON(200, gin.H{"data": product})

}

func GetProducts(ctx *gin.Context) {

	var product []models.Product

	result := inits.DB.Find(&product)
	if result.Error != nil {
		ctx.JSON(500, gin.H{"error": result.Error})
		return
	}

	ctx.JSON(200, gin.H{"data": product})

}

func GetProduct(ctx *gin.Context) {

	var product models.Product

	result := inits.DB.First(&product, ctx.Param("id"))
	if result.Error != nil {
		ctx.JSON(500, gin.H{"error": result.Error})
		return
	}

	ctx.JSON(200, gin.H{"data": product})

}

func UpdateProduct(ctx *gin.Context) {

	var body struct {
		Name        string
		Price       int
		Description string
		Stock       int
		Image_Url   string
	}

	ctx.BindJSON(&body)

	var product models.Product

	result := inits.DB.First(&product, ctx.Param("id"))
	if result.Error != nil {
		ctx.JSON(500, gin.H{"error": result.Error})
		return
	}

	inits.DB.Model(&product).Updates(models.Product{Name: body.Name, Price: body.Price, Description: body.Description, Stock: body.Stock, Image_Url: body.Image_Url})

	ctx.JSON(200, gin.H{"data": product})

}

func DeleteProduct(ctx *gin.Context) {

	id := ctx.Param("id")

	inits.DB.Delete(&models.Product{}, id)

	ctx.JSON(200, gin.H{"data": "product has been deleted successfully"})

}
