package controllers

import (
	"ecommerce_api/inits"
	"ecommerce_api/models"
	"fmt"

	"github.com/gin-gonic/gin"
)

func CreateCart(ctx *gin.Context) {

	var body struct {
		UserID uint
	}

	ctx.BindJSON(&body)

	user, exists := ctx.Get("user")

	if !exists {
		ctx.JSON(500, gin.H{"error": "user not found"})
		return
	}
	body.UserID = user.(models.User).ID
	cart := models.Cart{UserID: body.UserID}

	fmt.Println(cart)
	result := inits.DB.Create(&cart)
	if result.Error != nil {
		ctx.JSON(500, gin.H{"error": result.Error})
		return
	}

	err := inits.DB.Preload("User").First(&cart, cart.ID).Error
	if err != nil {
		ctx.JSON(500, gin.H{"error": err})
		return
	}

	ctx.JSON(200, gin.H{"data": cart})

}

func GetCarts(ctx *gin.Context) {

	var cart []models.Cart

	result := inits.DB.Preload("User").Find(&cart)
	if result.Error != nil {
		ctx.JSON(500, gin.H{"error": result.Error})
		return
	}

	ctx.JSON(200, gin.H{"data": cart})

}

func GetCart(ctx *gin.Context) {

	var cart models.Cart

	result := inits.DB.Preload("User").First(&cart, ctx.Param("id"))
	if result.Error != nil {
		ctx.JSON(500, gin.H{"error": result.Error})
		return
	}

	ctx.JSON(200, gin.H{"data": cart})

}

func DeleteCart(ctx *gin.Context) {

	id := ctx.Param("id")

	inits.DB.Delete(&models.Cart{}, id)

	ctx.JSON(200, gin.H{"data": "cart has been deleted successfully"})

}
