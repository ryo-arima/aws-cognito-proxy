package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ryo-arima/aws-cognito-proxy/pkg/entity/request"
	"github.com/ryo-arima/aws-cognito-proxy/pkg/entity/response"
	"github.com/ryo-arima/aws-cognito-proxy/pkg/server/repository"
)

type GroupsControllerForPublic interface {
	GetGroupss(c *gin.Context)
}

type groupsControllerForPublic struct {
	GroupsRepository repository.GroupsRepository
}

func (groupsController groupsControllerForPublic) GetGroupss(c *gin.Context) {
	var groupsRequest request.GroupsRequest
	if err := c.Bind(&groupsRequest); err != nil {
		c.JSON(http.StatusBadRequest, &response.GroupsResponse{Code: "SERVER_CONTROLLER_GET__FOR__001", Message: err.Error(), Groupss: []response.Groups{}})
		return
	}
	res := groupsController.GroupsRepository.GetGroupss()
	c.JSON(http.StatusOK, res)
	return
}


func NewGroupsControllerForPublic(groupsRepository repository.GroupsRepository) GroupsControllerForPublic {
	return &groupsControllerForPublic{ GroupsRepository: groupsRepository}
}