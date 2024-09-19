package main

import (
	"github.com/Anand-02/golang-learning-1/controllers"
	"github.com/gin-gonic/gin"
	"github.com/Anand-02/golang-learning-1/models"
	
)

func main() {

	router := gin.Default()

	models.ConnectDB()

	router.POST("/posts", controllers.CreatePost)
	router.GET("/posts", controllers.ReadPost)
	router.GET("/posts/:id", controllers.FindPost)
	router.PUT("/posts/:id", controllers.UpdatePost)
	router.DELETE("/posts/:id", controllers.DeletePost)
	

	router.Run("localhost:3000")
}
