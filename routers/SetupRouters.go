package routers

import (
	"go-crud/controllers"
	"go-crud/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.MaxMultipartMemory = 8 << 20 // 8 MiB
	v1 := r.Group("/v1")
	{
		v1.POST("/users", middleware.NewUserMiddleware, controllers.UsersCreate)
		v1.GET("/users", controllers.FetchAll)
		v1.GET("/users/:id", middleware.GetUserMiddleware, controllers.FindByID)
		v1.PUT("/users/:id", middleware.UpdateUserMiddleware, controllers.UsersUpdate)
		v1.DELETE("/users/:id", middleware.GetUserMiddleware, controllers.UsersDelete)
	}

	return r
}
