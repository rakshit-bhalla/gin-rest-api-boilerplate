package utils

import (
	"rakshit.dev/gin-rest-api-boilerplate/errors"
	"rakshit.dev/gin-rest-api-boilerplate/models"
)

type Response = models.Response
type Error = errors.Error

func GetResponse(data interface{}, err *Error) Response {
	if err != nil {
		return Response{
			Success: false,
			Data:    nil,
			Error:   err.Msg,
		}
	}
	return Response{
		Success: true,
		Data:    data,
		Error:   "",
	}
}
