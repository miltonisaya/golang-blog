package controllers

import (
	"blog/inits"
	"blog/models"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(ctx *gin.Context) {
	var body struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if ctx.BindJSON(&body) != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request"})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{Name: body.Name, Email: body.Email, Password: string(hash)}
	result := inits.DB.Create(&user)
	if result.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"user": user})
}

func Login(ctx *gin.Context) {
	var body struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if ctx.BindJSON(&body) != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request"})
		return
	}

	var user models.User
	result := inits.DB.Where("email = ?", body.Email).First(&user)

	if result.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Wrong Password"})
		return
	}

	//Generate JWT password
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  user.ID,
		"exp": time.Now().Add(time.Hour * time.Duration(24)).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.SetSameSite(http.SameSiteLaxMode)
	ctx.SetCookie("Authorization", tokenString, 3600*24*30, "", "localhost", false, true)
}

func GetUsers(ctx *gin.Context) {
	var users []models.User
	err := inits.DB.Model(&models.User{}).Preload("Posts").Find(&users).Error

	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error getting users"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"users": users})
}

func Validate(ctx *gin.Context) {
	user, err := ctx.Get("user")
	if err != false {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": "You are logged in", "user": user})
}
