package utils

import (
	"rakshit.dev/gin-rest-api-boilerplate/src/configs"
	"rakshit.dev/gin-rest-api-boilerplate/src/docs"
)

func SetSwaggerProps() {
	docs.SwaggerInfo.Host = configs.APIHost
}
