package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	UserID          uint       `json:"user_id"`
	User            User       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	OrderItems      []CartItem `json:"cart_items"`
	ShippingAddress string     `json:"shipping_address"`
	PaymentMethod   string     `json:"payment_method"`
	TotalPrice      int        `json:"total_price"`
	Status          string     `json:"status"`
}

type OrderItem struct {
	OrderID   uint    `gorm:"foreignKey:OrderID;constraint:fk_order_items_order" json:"order_id"`
	ProductID uint    `json:"product_id"`
	Product   Product `gorm:"foreignKey:ProductID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Quantity  int     `json:"quantity"`
	Subtotal  int     `json:"subtotal"`
}
