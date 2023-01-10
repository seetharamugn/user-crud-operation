package main

import (
	"go-crud/controllers"
	"go-crud/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectionDB()
}

func main() {
	r := gin.Default()
	//create user api
	r.POST("/users", controllers.UsersCreate)
	//fetch all user with pagination api
	r.GET("/users", controllers.FetchAll)
	//getby id user api
	r.GET("/users/:id", controllers.FindByID)
	//update user api
	r.PUT("/users/:id", controllers.UsersUpdate)
	//delete user api
	r.DELETE("/users/:id", controllers.UsersDelete)
	r.Run() // listen and serve on 0.0.0.0:8080
}
