package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name        string `json:"name" binding:"required"`
	Price       int    `json:"price"`
	Description string `json:"description"`
	Stock       int    `json:"stock" binding:"required"`
	Image_Url   string `json:"image_url"`
}
