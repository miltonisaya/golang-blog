package controllers

import (
	"blog/inits"
	"blog/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreatePost(ctx *gin.Context) {
	var body struct {
		Title  string `json:"title"`
		Body   string `json:"body"`
		Likes  int    `json:"likes"`
		Draft  bool   `json:"draft"`
		Author string `json:"author"`
		UserID uint   `json:'user_id'`
	}

	user, exist := ctx.Get("user")
	if !exist {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}
	body.UserID = user.(models.User).ID

	ctx.BindJSON(&body)
	post := models.Post{Title: body.Title, Body: body.Body, Likes: body.Likes, Author: body.Author, UserID: body.UserID}
	fmt.Println(post)
	result := inits.DB.Create(&post)
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"data": post})
}

func GetPosts(ctx *gin.Context) {
	var posts models.Post

	result := inits.DB.Find(&posts)
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
	}

	ctx.JSON(http.StatusOK, gin.H{"data": posts})
}

func GetPost(ctx *gin.Context) {
	var post models.Post

	result := inits.DB.First(&post, ctx.Param("id"))
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": post})
}

func UpdatePost(ctx *gin.Context) {
	var body struct {
		Title  string `json:"title"`
		Body   string `json:"body"`
		Likes  int    `json:"likes"`
		Draft  bool   `json:"draft"`
		Author string `json:"author"`
	}

	ctx.BindJSON(&body)
	var post models.Post
	result := inits.DB.First(&post, ctx.Param("id"))
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	inits.DB.Model(&post).Updates(models.Post{Title: body.Title, Body: body.Body, Likes: body.Likes, Author: body.Author})
	ctx.JSON(http.StatusOK, gin.H{"data": post})
}

func DeletePost(ctx *gin.Context) {
	id := ctx.Param("id")
	inits.DB.Delete(&models.Post{}, id)
	ctx.JSON(http.StatusOK, gin.H{"data": "Post has been deleted successfully"})
}
