package controllers

import (
	"fmt"
	"net/http"

	"rakshit.dev/gin-rest-api-boilerplate/src/errors"
	"rakshit.dev/gin-rest-api-boilerplate/src/models"
	"rakshit.dev/gin-rest-api-boilerplate/src/services"
	"rakshit.dev/gin-rest-api-boilerplate/src/utils"

	"github.com/gin-gonic/gin"
)

type User = models.User
type UserReq = models.UserReq
type Response = models.Response

// UserController ...
type UserController interface {
	GetUser(c *gin.Context)
	GetUsers(c *gin.Context)
	DeleteUser(c *gin.Context)
	CreateUser(c *gin.Context)
}

type userController struct {
	userService services.UserService
}

// NewUserController ...
func NewUserController(userService services.UserService) UserController {
	return &userController{
		userService: userService,
	}
}

// GetUser ... Get the user by userId
// @Summary Get one user
// @Description get user by UserID
// @Tags user-controller
// @Param userId path string true "UserID"
// @Success 200 {object} models.Response{data=models.User}
// @Failure 404,500 {object} models.Response
// @Router /{userId} [get]
func (u *userController) GetUser(c *gin.Context) {
	userID := c.Param("userId")
	user, err := u.userService.GetUser(userID)
	if status := errors.GetHttpStatus(err); nil != status {
		c.IndentedJSON(*status, utils.GetResponse(nil, err))
		return
	}
	if nil != err {
		fmt.Printf("File: %s  Function: %s Line: %d Error: %s", err.File, err.Func, err.Line, err.Msg)
		err := errors.CreateError("internal server error")
		c.IndentedJSON(http.StatusInternalServerError, utils.GetResponse(nil, &err))
		return
	}
	c.IndentedJSON(http.StatusOK, utils.GetResponse(user, nil))
}

// GetUsers ... Get all users
// @Summary Get all users
// @Description get all users
// @Tags user-controller
// @Success 200 {object} models.Response{data=[]models.User}
// @Failure 500 {object} models.Response
// @Router / [get]
func (u *userController) GetUsers(c *gin.Context) {
	users, err := u.userService.GetUsers()
	if nil != err {
		fmt.Printf("File: %s  Function: %s Line: %d Error: %s", err.File, err.Func, err.Line, err.Msg)
		err := errors.CreateError("internal server error")
		c.IndentedJSON(http.StatusInternalServerError, utils.GetResponse(nil, &err))
		return
	}
	c.IndentedJSON(http.StatusOK, utils.GetResponse(users, nil))
}

// DeleteUser ... Delete the user by userId
// @Summary Delete one user
// @Description delete user by UserID
// @Tags user-controller
// @Param userId path string true "UserID"
// @Success 200 {object} models.Response{data=models.User}
// @Failure 404,500 {object} models.Response
// @Router /{userId} [delete]
func (u *userController) DeleteUser(c *gin.Context) {
	userID := c.Param("userId")
	user, err := u.userService.DeleteUser(userID)
	if status := errors.GetHttpStatus(err); nil != status {
		c.IndentedJSON(*status, utils.GetResponse(nil, err))
		return
	}
	if nil != err {
		fmt.Printf("File: %s  Function: %s Line: %d Error: %s", err.File, err.Func, err.Line, err.Msg)
		err := errors.CreateError("internal server error")
		c.IndentedJSON(http.StatusInternalServerError, utils.GetResponse(nil, &err))
		return
	}
	c.IndentedJSON(http.StatusOK, utils.GetResponse(user, nil))
}

// CreateUser ... Create User
// @Summary Create new user
// @Description Create new user based on parameters
// @Tags user-controller
// @Accept json
// @Param user body models.UserReq true "User Data"
// @Success 201 {object}  models.Response{data=models.User}
// @Failure 400,500 {object}  models.Response
// @Router / [post]
func (u *userController) CreateUser(c *gin.Context) {
	var input UserReq
	if err := c.ShouldBindJSON(&input); err != nil {
		err := errors.CreateError(err.Error())
		fmt.Printf("File: %s  Function: %s Line: %d Error: %s", err.File, err.Func, err.Line, err.Msg)
		err = errors.CreateError("bad params")
		c.IndentedJSON(http.StatusBadRequest, utils.GetResponse(nil, &err))
		return
	}
	user, err := u.userService.CreateUser(input)
	if nil != err {
		fmt.Printf("File: %s  Function: %s Line: %d Error: %s", err.File, err.Func, err.Line, err.Msg)
		err := errors.CreateError("internal server error")
		c.IndentedJSON(http.StatusInternalServerError, utils.GetResponse(nil, &err))
		return
	}
	c.IndentedJSON(http.StatusCreated, utils.GetResponse(user, nil))
}
