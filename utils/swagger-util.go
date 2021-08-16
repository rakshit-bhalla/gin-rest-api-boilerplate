package utils

import (
	"rakshit.dev/gin-rest-api-boilerplate/docs"
)

func SetSwaggerProps(host string) {
	docs.SwaggerInfo.Host = host
}
