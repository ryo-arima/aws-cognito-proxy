package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ryo-arima/aws-cognito-proxy/pkg/entity/model"
	"github.com/ryo-arima/aws-cognito-proxy/pkg/entity/request"
	"github.com/ryo-arima/aws-cognito-proxy/pkg/entity/response"
	"github.com/ryo-arima/aws-cognito-proxy/pkg/server/repository"
)

type UsersControllerForPrivate interface {
	GetUserss(c *gin.Context)
	CreateUsers(c *gin.Context)
	UpdateUsers(c *gin.Context)
	DeleteUsers(c *gin.Context)
}

type usersControllerForPrivate struct {
	UsersRepository repository.UsersRepository
}

func (usersController usersControllerForPrivate) GetUserss(c *gin.Context) {
	var usersRequest request.UsersRequest
	if err := c.Bind(&usersRequest); err != nil {
		c.JSON(http.StatusBadRequest, &response.UsersResponse{Code: "SERVER_CONTROLLER_GET__FOR__001", Message: err.Error(), Userss: []response.Users{}})
		return
	}
	res := usersController.UsersRepository.GetUserss()
	c.JSON(http.StatusOK, res)
	return
}


func (usersController usersControllerForPrivate) CreateUsers(c *gin.Context) {
	var usersRequest request.UsersRequest
	if err := c.Bind(&usersRequest); err != nil {
		c.JSON(http.StatusBadRequest, &response.UsersResponse{Code: "SERVER_CONTROLLER_CREATE__FOR__001", Message: err.Error(), Userss: []response.Users{}})
		return
	}
	var usersModel model.Userss
	res := usersController.UsersRepository.CreateUsers(usersModel)
	c.JSON(http.StatusOK, res)
	return
}


func (usersController usersControllerForPrivate) UpdateUsers(c *gin.Context) {
	var usersRequest request.UsersRequest
	if err := c.Bind(&usersRequest); err != nil {
		c.JSON(http.StatusBadRequest, &response.UsersResponse{Code: "SERVER_CONTROLLER_UPDATE__FOR__001", Message: err.Error(), Userss: []response.Users{}})
		return
	}
	var usersModel model.Userss
	res := usersController.UsersRepository.UpdateUsers(usersModel)
	c.JSON(http.StatusOK, res)
	return
}


func (usersController usersControllerForPrivate) DeleteUsers(c *gin.Context) {
	var usersRequest request.UsersRequest
	if err := c.Bind(&usersRequest); err != nil {
		c.JSON(http.StatusBadRequest, &response.UsersResponse{Code: "SERVER_CONTROLLER_DELETE__FOR__001", Message: err.Error(), Userss: []response.Users{}})
		return
	}
	var uuid string
	res := usersController.UsersRepository.DeleteUsers(uuid)
	c.JSON(http.StatusOK, res)
	return
}


func NewUsersControllerForPrivate(usersRepository repository.UsersRepository) UsersControllerForPrivate {
	return &usersControllerForPrivate{ UsersRepository: usersRepository}
}