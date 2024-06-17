package controllers

import (
	"github.com/Sunsetus/gocrud/initializers"
	"github.com/Sunsetus/gocrud/models"
	"github.com/gin-gonic/gin"
)

func PostsCreate(c *gin.Context) {
	//Get a data off req body
	var body struct{
		Body	string
		Tittle	string
	}

	c.Bind(&body)

	//Create a post
	post := models.Post{Tittle: body.Tittle, Body: body.Body}

	result := initializers.DB.Create(&post)

	if result.Error != nil{
		c.Status(400)
		return
	}
	
	//Return it
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsIndex (c *gin.Context){
	var posts []models.Post
	initializers.DB.Find(&posts)

	//Respond
	c.JSON(200, gin.H{ 
		"post": posts,
	})
}

func PostsShow (c *gin.Context) {
	// Get ID of URL
	id :=c.Param("id")

	//Get the post
	var post models.Post
	initializers.DB.First(&post, id)

	//Respond with them
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsUpdate(c *gin.Context){
	// Get id off the URL
	id :=c.Param("id")

	// Get the date off request body

	var body struct{
		Body	string
		Tittle	string
	}

	c.Bind(&body)

	// Find the post were updating 
	var post models.Post
	initializers.DB.First(&post, id)

	// Update it
	initializers.DB.Model(&post).Updates(models.Post{
		Tittle: body.Tittle,
		Body: body.Body,
	})

	// Respond with it
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsDelete(c *gin.Context){
	// Get id off the URL
	id :=c.Param("id")

	// Delete post
	initializers.DB.Delete(&models.Post{}, id)

	//Respond
	c.Status(200)
}

