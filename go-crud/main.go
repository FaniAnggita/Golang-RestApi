package main

import (
	"github.com/FaniAnggita/go-crud/controllers"
	"github.com/FaniAnggita/go-crud/initializers"
	"github.com/FaniAnggita/go-crud/models"
	"github.com/gin-gonic/gin"
)

func init(){
	initializers.LoadEnvVariables()
	initializers.ConnectTODB()
	
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
	r := gin.Default()

	// Route
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("/create", controllers.PostsCreate)
	r.GET("/post", controllers.PostIndex)
	r.GET("/post/:id", controllers.PostbyID)
	r.PUT("/post/:id", controllers.PostUpdate)
	r.DELETE("/post/:id", controllers.PostDelete)

	r.Run() // listen and serve on 0.0.0.0:8080
}