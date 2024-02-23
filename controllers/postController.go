package controllers

import (
	"example.com/m/v2/initializers"
	"example.com/m/v2/models"
	"github.com/gin-gonic/gin"
)

func PostCreate (c *gin.Context) {

	// Get data from body
	var body struct {
		Body string
		Title string
	}

	c.Bind(&body)

	post := models.Post{ Title: body.Title, Body: body.Body }
	result := initializers.DB.Create(&post)
	if result.Error != nil {
		c.JSON(500, gin.H{
			"error": result.Error,
		})
		return
	}

	if result.RowsAffected == 0 {
		c.JSON(500, gin.H{
			"error": "Could not create post",
		})
		return
	}

	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostIndex (c *gin.Context) {
	var posts []models.Post
	result := initializers.DB.Find(&posts)
	if result.Error != nil {
		c.JSON(500, gin.H{
			"error": result.Error,
		})
		return
	}

	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func PostShow (c *gin.Context) {
	var post models.Post
	result := initializers.DB.First(&post, c.Param("id"))
	if result.Error != nil {
		c.JSON(500, gin.H{
			"error": result.Error,
		})
		return
	}

	c.JSON(200, gin.H{
		"post": post,
	})
}


func PostUpdate(c *gin.Context) {
	var post models.Post
	result := initializers.DB.First(&post, c.Param("id"))
	if result.Error != nil {
		c.JSON(500, gin.H{
			"error": result.Error,
		})
		return
	}

	var body struct {
		Body string
		Title string
	}

	c.Bind(&body)

	post.Title = body.Title
	post.Body = body.Body

	result = initializers.DB.Save(&post)
	if result.Error != nil {
		c.JSON(500, gin.H{
			"error": result.Error,
		})
		return
	}

	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostDelete(c *gin.Context) {
	var post models.Post
	result := initializers.DB.Delete(&post, c.Param("id"))
	if result.Error != nil {
		c.JSON(500, gin.H{
			"error": result.Error,
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Post deleted",
	})
}