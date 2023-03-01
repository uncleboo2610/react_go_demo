package controllers

import (
	config "first-app/configuration"
	"first-app/models"
	"fmt"

	"github.com/gin-gonic/gin"
)

func PostCreate(c *gin.Context) {
	var reqBody struct {
		Title string
		Body  string
		ID    int
		Name  string
	}

	c.Bind(&reqBody)

	post := models.Post{
		Title: reqBody.Title,
		Body:  reqBody.Body,
		User:  models.User{ID: reqBody.ID, Name: reqBody.Name}}

	result := config.DB.Create(&post)
	fmt.Println(result)
	c.JSON(200, post)
}

func PostsRead(c *gin.Context) {
	// Get all records
	var posts []models.Post
	config.DB.Preload("User").Find(&posts)
	c.JSON(200, posts)
}

func PostShow(c *gin.Context) {
	id := c.Param("id")

	var posts []models.Post
	config.DB.Preload("User").Find(&posts, id)
	c.JSON(200, posts)
}

func PostUpdate(c *gin.Context) {
	var reqBody struct {
		Title string
		Body  string
	}
	c.Bind(&reqBody)
	id := c.Param("id")

	var posts []models.Post
	config.DB.Preload("User").Find(&posts, id)
	config.DB.Model(&posts).Updates(models.Post{Title: reqBody.Title, Body: reqBody.Body})
	c.JSON(200, gin.H{
		"post": posts,
	})
}

func UserUpdate(c *gin.Context) {
	var reqBody struct {
		Name string
	}
	c.Bind(&reqBody)
	id := c.Param("id")

	var users []models.User
	config.DB.First(&users, id)
	config.DB.Model(&users).Updates(models.User{Name: reqBody.Name})
	c.JSON(200, gin.H{
		"user": users,
	})
}

func PostDelete(c *gin.Context) {
	id := c.Param("id")
	var posts []models.Post
	config.DB.Select("User").Delete(&posts, id)
}
