package models

import "gorm.io/gorm"

type CartItem struct {
	gorm.Model
	CartID    uint    `json:"cart_id"`
	Cart      Cart    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	ProductID uint    `json:"product_id"`
	Product   Product `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Quantity  int     `json:"quantity"`
	Subtotal  int     `json:"subtotal"`
	OrderID   *uint   `gorm:"foreignKey:OrderID;constraint:fk_cart_items_order;sql:-"`
}
