package controllers

import (
	"fmt"
	"net/http"

	"github.com/Anand-02/golang-learning-1/models"
	"github.com/gin-gonic/gin"
)

type PostInput struct {
	Title  string `json:"title" binding:"required"`
	Body   string `json:"body" binding:"required"`
	Author string `json:"author" binding:"required"`
}

func CreatePost(ctx *gin.Context) {
	var input PostInput

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	post := models.BlogItem{Title: input.Title, Author: input.Author, Body: input.Body}
	models.DB.Create(&post)

	ctx.JSON(http.StatusOK, gin.H{"data": post})
}

func ReadPost(ctx *gin.Context) {
	var blogposts []models.BlogItem

	models.DB.Find(&blogposts)

	ctx.JSON(http.StatusOK, gin.H{"data": blogposts})
}

func FindPost(ctx *gin.Context) {
	var reqpost models.BlogItem

	if err := models.DB.Where("id = ?", ctx.Param("id")).First(&reqpost).Error; err != nil {
		fmt.Println("Invalid ID. Requested blog does not exist.")
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": reqpost})
}

func UpdatePost(ctx *gin.Context){
	id := ctx.Param("id")

	var blogToUpdate models.BlogItem

	if err := models.DB.Where("id = ?", id).First(&blogToUpdate).Error; err != nil {
		fmt.Println("Invalid ID. Requested blog does not exist.")
		return
	}

	var input PostInput

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&blogToUpdate).Updates(input)

	ctx.JSON(http.StatusOK, gin.H{"data": blogToUpdate})
}

func DeletePost(ctx *gin.Context){
	id := ctx.Param("id")

	var blogToDelete models.BlogItem

	if err := models.DB.Where("id = ?", id).First(&blogToDelete).Error; err != nil {
		fmt.Println("Invalid ID. Requested blog does not exist.")
		return
	}

	models.DB.Delete(&blogToDelete);

	ctx.JSON(http.StatusOK, gin.H{"data": "Successfully Deleted!"})
}