package controllers

import (
	"ecommerce_api/inits"
	"ecommerce_api/models"
	"errors"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateCartItem(ctx *gin.Context) {
	var body struct {
		CartID    uint `json:"cart_id"`
		ProductID uint `json:"product_id"`
		Quantity  int  `json:"quantity"`
	}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.JSON(400, gin.H{"error": "invalid request body"})
		return
	}

	// Fetch existing cart and product
	var cart models.Cart
	if err := inits.DB.First(&cart, body.CartID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(404, gin.H{"error": "cart not found"})
		} else {
			ctx.JSON(500, gin.H{"error": err})
		}
		return
	}

	var product models.Product
	if err := inits.DB.First(&product, body.ProductID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(404, gin.H{"error": "product not found"})
		} else {
			ctx.JSON(500, gin.H{"error": err})
		}
		return
	}

	// Calculate subtotal
	subtotal := product.Price * body.Quantity

	// Create cart item
	cartItem := models.CartItem{
		CartID:    body.CartID,
		ProductID: body.ProductID,
		Quantity:  body.Quantity,
		Subtotal:  subtotal,
	}

	if err := inits.DB.Create(&cartItem).Error; err != nil {
		ctx.JSON(500, gin.H{"error": err})
		return
	}

	err := inits.DB.Preload("Cart").Preload("Product").First(&cartItem, cartItem.ID).Error
	if err != nil {
		ctx.JSON(500, gin.H{"error": err})
		return
	}

	ctx.JSON(201, gin.H{"data": cartItem})
}

func GetCartItems(ctx *gin.Context) {

	var cart_item []models.CartItem

	result := inits.DB.Preload("Cart").Preload("Product").Find(&cart_item)
	if result.Error != nil {
		ctx.JSON(500, gin.H{"error": result.Error})
		return
	}

	ctx.JSON(200, gin.H{"data": cart_item})

}

func GetCartItem(ctx *gin.Context) {

	var cart_item models.CartItem

	result := inits.DB.Preload("Cart").Preload("Product").First(&cart_item, ctx.Param("id"))
	if result.Error != nil {
		ctx.JSON(500, gin.H{"error": result.Error})
		return
	}

	ctx.JSON(200, gin.H{"data": cart_item})

}

func DeleteCartItem(ctx *gin.Context) {

	id := ctx.Param("id")

	inits.DB.Delete(&models.CartItem{}, id)

	ctx.JSON(200, gin.H{"data": "cart item has been deleted successfully"})

}
