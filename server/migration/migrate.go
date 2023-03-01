package main

import (
	config "first-app/configuration"
	"first-app/models"
)

func init() {
	config.ConnectToDB()
}

func main() {
	config.DB.AutoMigrate(&models.Post{})
	config.DB.AutoMigrate(&models.User{})
	config.DB.AutoMigrate(&models.Auth{})
}
