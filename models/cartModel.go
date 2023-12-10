package models

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	UserID uint `json:"user_id"`
	User   User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
