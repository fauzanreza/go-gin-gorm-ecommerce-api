package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username   string `json:"username" binding:"required"`
	Email      string `json:"email" gorm:"unique"`
	Password   string `json:"password" binding:"required"`
	First_Name string `json:"first_name" binding:"required"`
	Last_Name  string `json:"last_name"`
	Phone      string `json:"phone"`
	Address    string `json:"address"`
}
