package main

import (
	"ecommerce_api/inits"

	"ecommerce_api/models"
)

func init() {
	inits.LoadEnv()
	inits.DBInit()
}

func main() {
	inits.DB.AutoMigrate(
		&models.Product{},
		&models.Category{},
		&models.User{},
		&models.Cart{},
		&models.CartItem{},
		&models.Order{},
		&models.OrderItem{},
	)

}
