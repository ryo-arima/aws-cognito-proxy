package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ryo-arima/aws-cognito-proxy/pkg/entity/request"
	"github.com/ryo-arima/aws-cognito-proxy/pkg/entity/response"
	"github.com/ryo-arima/aws-cognito-proxy/pkg/server/repository"
)

type UsersControllerForPublic interface {
	GetUserss(c *gin.Context)
}

type usersControllerForPublic struct {
	UsersRepository repository.UsersRepository
}

func (usersController usersControllerForPublic) GetUserss(c *gin.Context) {
	var usersRequest request.UsersRequest
	if err := c.Bind(&usersRequest); err != nil {
		c.JSON(http.StatusBadRequest, &response.UsersResponse{Code: "SERVER_CONTROLLER_GET__FOR__001", Message: err.Error(), Userss: []response.Users{}})
		return
	}
	res := usersController.UsersRepository.GetUserss()
	c.JSON(http.StatusOK, res)
	return
}


func NewUsersControllerForPublic(usersRepository repository.UsersRepository) UsersControllerForPublic {
	return &usersControllerForPublic{ UsersRepository: usersRepository}
}