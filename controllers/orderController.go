package controllers

import (
	"ecommerce_api/inits"
	"ecommerce_api/models"
	"errors"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateOrder(ctx *gin.Context) {
	// Get logged-in user
	user, ok := ctx.Get("user")
	if !ok {
		ctx.JSON(401, gin.H{"error": "unauthorized"})
		return
	}

	userID := user.(models.User).ID // Get the user ID from the retrieved user object

	// Validate order details
	var orderDetails struct {
		CartID          uint   `json:"cart_id" binding:"required"`
		ShippingAddress string `json:"shipping_address" binding:"required"`
		PaymentMethod   string `json:"payment_method" binding:"required"`
	}

	if err := ctx.BindJSON(&orderDetails); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// Fetch cart and verify ownership
	var cart models.Cart
	if err := inits.DB.Where("user_id = ? AND id = ?", userID, orderDetails.CartID).First(&cart).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(404, gin.H{"error": "cart not found or invalid user"})
		} else {
			ctx.JSON(500, gin.H{"error": err})
		}
		return
	}

	// Fetch all cart items
	var cartItems []models.CartItem
	if err := inits.DB.Where("cart_id = ?", cart.ID).Find(&cartItems).Error; err != nil {
		ctx.JSON(500, gin.H{"error": err})
		return
	}

	// Calculate total price
	totalPrice := 0
	for _, item := range cartItems {
		totalPrice += item.Subtotal
	}

	// Create order
	order := models.Order{
		UserID:          userID,
		OrderItems:      cartItems,
		ShippingAddress: orderDetails.ShippingAddress,
		PaymentMethod:   orderDetails.PaymentMethod,
		TotalPrice:      totalPrice,
		Status:          "pending", // adjust initial order status
	}

	// Save order and update cart items (mark as ordered)
	if err := inits.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&order).Error; err != nil {
			return err
		}
		for _, item := range cartItems {
			item.OrderID = &order.ID                     // Set the order ID for each item
			if err := tx.Save(&item).Error; err != nil { // Use Save instead of Update
				return err
			}
		}
		return nil
	}); err != nil {
		ctx.JSON(500, gin.H{"error": err})
		return
	}

	// Clear cart after order creation (optional, adjust as needed)
	if err := inits.DB.Where("id = ?", cart.ID).Delete(&cart).Error; err != nil {
		ctx.JSON(500, gin.H{"error": err})
		return
	}

	// Respond with order information
	ctx.JSON(201, gin.H{"data": order})
}

func GetOrders(ctx *gin.Context) {

	var order []models.Order

	result := inits.DB.Find(&order)
	if result.Error != nil {
		ctx.JSON(500, gin.H{"error": result.Error})
		return
	}

	ctx.JSON(200, gin.H{"data": order})

}

func GetOrder(ctx *gin.Context) {

	var order models.Order

	result := inits.DB.First(&order, ctx.Param("id"))
	if result.Error != nil {
		ctx.JSON(500, gin.H{"error": result.Error})
		return
	}

	ctx.JSON(200, gin.H{"data": order})

}

func DeleteOrder(ctx *gin.Context) {

	id := ctx.Param("id")

	inits.DB.Delete(&models.Order{}, id)

	ctx.JSON(200, gin.H{"data": "cart item has been deleted successfully"})

}
