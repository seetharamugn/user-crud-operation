package main

import (
	"go-crud/initializers"
	"go-crud/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectionDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.User{})
}
