package controllers

import (
	"ecommerce_api/inits"
	"ecommerce_api/models"
	"fmt"

	"github.com/gin-gonic/gin"
)

func CreateCategory(ctx *gin.Context) {

	var body struct {
		Category_Name string
		Description   string
	}

	ctx.BindJSON(&body)

	category := models.Category{Category_Name: body.Category_Name, Description: body.Description}

	fmt.Println(category)
	result := inits.DB.Create(&category)
	if result.Error != nil {
		ctx.JSON(500, gin.H{"error": result.Error})
		return
	}

	ctx.JSON(200, gin.H{"data": category})

}

func GetCategorys(ctx *gin.Context) {

	var category []models.Category

	result := inits.DB.Find(&category)
	if result.Error != nil {
		ctx.JSON(500, gin.H{"error": result.Error})
		return
	}

	ctx.JSON(200, gin.H{"data": category})

}

func GetCategory(ctx *gin.Context) {

	var category models.Category

	result := inits.DB.First(&category, ctx.Param("id"))
	if result.Error != nil {
		ctx.JSON(500, gin.H{"error": result.Error})
		return
	}

	ctx.JSON(200, gin.H{"data": category})

}

func UpdateCategory(ctx *gin.Context) {

	var body struct {
		Category_Name string
		Description   string
	}

	ctx.BindJSON(&body)

	var category models.Category

	result := inits.DB.First(&category, ctx.Param("id"))
	if result.Error != nil {
		ctx.JSON(500, gin.H{"error": result.Error})
		return
	}

	inits.DB.Model(&category).Updates(models.Category{Category_Name: body.Category_Name, Description: body.Description})

	ctx.JSON(200, gin.H{"data": category})

}

func DeleteCategory(ctx *gin.Context) {

	id := ctx.Param("id")

	inits.DB.Delete(&models.Category{}, id)

	ctx.JSON(200, gin.H{"data": "category has been deleted successfully"})

}
