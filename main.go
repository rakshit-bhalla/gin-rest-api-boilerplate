package main

import (
	"github.com/gin-gonic/gin"
	"rakshit.dev/gin-rest-api-boilerplate/controllers"
	"rakshit.dev/gin-rest-api-boilerplate/db"
	"rakshit.dev/gin-rest-api-boilerplate/repositories"
	"rakshit.dev/gin-rest-api-boilerplate/routers"
	"rakshit.dev/gin-rest-api-boilerplate/services"
	"rakshit.dev/gin-rest-api-boilerplate/utils"
)

// @title Gin Rest Api Boilerplate
// @version 1.0.0
// @description Gin Rest Api Boilerplate
// @termsOfService http://swagger.io/terms/

// @contact.name Rakshit Bhalla
// @contact.url https://rakshit.dev
// @contact.email contact@rakshit.dev

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath /api/v1
// @query.collection.format multi

func main() {
	router := gin.Default()
	mongoDB := db.GetMongo()
	host := "localhost:5000"
	userRepository := repositories.NewUserRepository(mongoDB)
	userService := services.NewUserService(userRepository)
	userController := controllers.NewUserController(userService)
	routers.SetUserRoutes(router, userController)
	utils.SetSwaggerProps(host)
	routers.SetSwaggerRoutes(router)
	router.Run(host)
}
