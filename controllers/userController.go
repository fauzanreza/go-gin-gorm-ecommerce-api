package controllers

import (
	"ecommerce_api/inits"
	"ecommerce_api/models"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func Signup(ctx *gin.Context) {
	var body struct {
		Username   string
		Email      string
		Password   string
		First_Name string
		Last_Name  string
		Phone      string
		Address    string
	}

	if ctx.BindJSON(&body) != nil {
		ctx.JSON(400, gin.H{"error": "bad request"})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
		ctx.JSON(500, gin.H{"error": err})
		return
	}

	user := models.User{Username: body.Username, Email: body.Email, Password: string(hash), First_Name: body.First_Name, Last_Name: body.Last_Name, Phone: body.Phone, Address: body.Address}

	result := inits.DB.Create(&user)

	if result.Error != nil {
		ctx.JSON(500, gin.H{"error": result.Error})
		return
	}

	ctx.JSON(200, gin.H{"data": user})

}

func Login(ctx *gin.Context) {
	var body struct {
		Email    string
		Password string
	}

	if ctx.BindJSON(&body) != nil {
		ctx.JSON(400, gin.H{"error": "bad request"})
		return
	}

	var user models.User

	result := inits.DB.Where("email = ?", body.Email).First(&user)

	if result.Error != nil {
		ctx.JSON(500, gin.H{"error": "User not found"})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if err != nil {
		ctx.JSON(401, gin.H{"error": "unauthorized"})
		return
	}

	// generate jwt token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		ctx.JSON(500, gin.H{"error": "error signing token"})
		return
	}

	ctx.SetSameSite(http.SameSiteLaxMode)
	ctx.SetCookie("Authorization", tokenString, 3600*24*30, "", "localhost", false, true)
	ctx.JSON(200, gin.H{"data": "You are login successfully", "token": tokenString})
}

func GetUsers(ctx *gin.Context) {

	var user []models.User

	result := inits.DB.Find(&user)
	if result.Error != nil {
		ctx.JSON(500, gin.H{"error": result.Error})
		return
	}

	ctx.JSON(200, gin.H{"data": user})

}

func Logout(ctx *gin.Context) {
	ctx.SetSameSite(http.SameSiteLaxMode)
	ctx.SetCookie("Authorization", "", -1, "", "localhost", false, true)
	ctx.JSON(200, gin.H{"data": "You are logged out!"})
}
