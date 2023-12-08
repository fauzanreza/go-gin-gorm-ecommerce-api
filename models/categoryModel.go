package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Category_Name string `json:"category_name" binding:"required"`
	Description   string `json:"description"`
}
