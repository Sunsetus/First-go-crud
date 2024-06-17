package main

import (
	"github.com/Sunsetus/gocrud/initializers"
	"github.com/Sunsetus/gocrud/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}