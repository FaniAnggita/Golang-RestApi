package controllers

import (
	"net/http"

	"github.com/FaniAnggita/go-crud/initializers"
	"github.com/FaniAnggita/go-crud/models"
	"github.com/gin-gonic/gin"
)

func PostsCreate(c *gin.Context) {
	// GET DATA
	var body struct {
		Title string `json:"title" binding:"required"`
		Body  string `json:"body" binding:"required"`
	}

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if Title or Body is empty
	if body.Title == "" || body.Body == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Title and Body cannot be empty"})
		return
	}

	// CREATE A POST
	post := models.Post{Title: body.Title, Body: body.Body}

	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	// RETURN RESPONSE
	c.JSON(http.StatusOK, gin.H{
		"post": post,
	})
}

func PostIndex(c  *gin.Context) {
	var posts []models.Post
	initializers.DB.Find(&posts)

	c.JSON(http.StatusOK, gin.H{
		"post": posts,
	})

}

func PostbyID(c  *gin.Context) {
	id := c.Param("id")
	var post []models.Post
	initializers.DB.Find(&post, id)
	
	c.JSON(http.StatusOK, gin.H{
		"post": post,
	})
}

func PostUpdate(c *gin.Context) {
	// Get ID from URL Params
	id := c.Param("id")

	// Get the data from req body
	var body struct {
		Body  string `json:"body"`
		Title string `json:"title"`
	}

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Find the post we're updating
	var post models.Post
	if err := initializers.DB.First(&post, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	// Update
	if err := initializers.DB.Model(&post).Updates(models.Post{Title: body.Title, Body: body.Body}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update post"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"updated": post,
	})
}

func PostDelete(c *gin.Context) {
	id := c.Param("id")

	// Mencari posting dengan ID yang diberikan
	var post models.Post
	if err := initializers.DB.First(&post, id).Error; err != nil {
		// Jika posting tidak ditemukan, kembalikan pesan kesalahan
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "Post not found",
			"message": err.Error(),
		})
		return
	}

	// Menghapus posting
	result := initializers.DB.Delete(&post)
	if result.Error != nil {
		// Jika gagal menghapus posting, kembalikan pesan kesalahan
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to delete post",
			"message": result.Error.Error(),
		})
		return
	}

	// Jika berhasil menghapus posting, kembalikan pesan sukses
	c.JSON(http.StatusOK, gin.H{
		"message": "Post Deleted!",
	})
}


