package controllers

import (
	"github.com/afriwondimu/go-gin-pra/initializers"
	"github.com/afriwondimu/go-gin-pra/models"
	"github.com/gin-gonic/gin"
)

func PostsCreate(c *gin.Context) {
	// Get data off req body
	var body struct{
		Body string
		Title string
	}
	c.Bind(&body)

	// Createa post
	post := models.Post{Title: body.Title, Body: body.Body}

	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}

	// return
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsIndex(c *gin.Context) {
	// Get the post
	var posts []models.Post
	initializers.DB.Find(&posts)

	// return
	c.JSON(200, gin.H{
		"post": posts,
	})
}

func PostsShow(c *gin.Context) {
	// GEt id off url
	id := c.Param("id")
	// Get the post
	var post models.Post
	initializers.DB.First(&post, id)

	// return
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsUpdate(c *gin.Context) {
	// GEt id off url
	id := c.Param("id")

	// Get the data off req body
	var body struct{
		Body string
		Title string
	}
	c.Bind(&body)

	// Find the post were updating
	var post models.Post
	initializers.DB.First(&post, id)

	// Update it

	initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body: body.Body,
	})
	// return
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsDelete(c *gin.Context) {
	// GEt id off url
	id := c.Param("id")

	// Delete it 

	initializers.DB.Delete(&models.Post{}, id)

	// return
	c.Status(200)

}