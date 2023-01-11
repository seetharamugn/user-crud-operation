package main

import (
	"go-crud/initializers"
	"go-crud/routers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectionDB()
}

func main() {
	r := routers.SetupRouter()
	r.Run()
}
