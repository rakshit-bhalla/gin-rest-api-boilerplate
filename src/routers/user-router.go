package routers

import (
	"github.com/gin-gonic/gin"
	"rakshit.dev/gin-rest-api-boilerplate/src/controllers"
)

type UserController = controllers.UserController

// SetUserRoutes ...
func SetUserRoutes(router *gin.Engine, userController UserController) {
	userRouter := router.Group("/users")
	userRouter.GET("/", userController.GetUsers)
	userRouter.POST("/", userController.CreateUser)
	userRouter.GET("/:userId", userController.GetUser)
	userRouter.DELETE("/:userId", userController.DeleteUser)
}
