package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ryo-arima/aws-cognito-proxy/pkg/entity/model"
	"github.com/ryo-arima/aws-cognito-proxy/pkg/entity/request"
	"github.com/ryo-arima/aws-cognito-proxy/pkg/entity/response"
	"github.com/ryo-arima/aws-cognito-proxy/pkg/server/repository"
)

type GroupsControllerForPrivate interface {
	GetGroupss(c *gin.Context)
	CreateGroups(c *gin.Context)
	UpdateGroups(c *gin.Context)
	DeleteGroups(c *gin.Context)
}

type groupsControllerForPrivate struct {
	GroupsRepository repository.GroupsRepository
}

func (groupsController groupsControllerForPrivate) GetGroupss(c *gin.Context) {
	var groupsRequest request.GroupsRequest
	if err := c.Bind(&groupsRequest); err != nil {
		c.JSON(http.StatusBadRequest, &response.GroupsResponse{Code: "SERVER_CONTROLLER_GET__FOR__001", Message: err.Error(), Groupss: []response.Groups{}})
		return
	}
	res := groupsController.GroupsRepository.GetGroupss()
	c.JSON(http.StatusOK, res)
	return
}


func (groupsController groupsControllerForPrivate) CreateGroups(c *gin.Context) {
	var groupsRequest request.GroupsRequest
	if err := c.Bind(&groupsRequest); err != nil {
		c.JSON(http.StatusBadRequest, &response.GroupsResponse{Code: "SERVER_CONTROLLER_CREATE__FOR__001", Message: err.Error(), Groupss: []response.Groups{}})
		return
	}
	var groupsModel model.Groupss
	res := groupsController.GroupsRepository.CreateGroups(groupsModel)
	c.JSON(http.StatusOK, res)
	return
}


func (groupsController groupsControllerForPrivate) UpdateGroups(c *gin.Context) {
	var groupsRequest request.GroupsRequest
	if err := c.Bind(&groupsRequest); err != nil {
		c.JSON(http.StatusBadRequest, &response.GroupsResponse{Code: "SERVER_CONTROLLER_UPDATE__FOR__001", Message: err.Error(), Groupss: []response.Groups{}})
		return
	}
	var groupsModel model.Groupss
	res := groupsController.GroupsRepository.UpdateGroups(groupsModel)
	c.JSON(http.StatusOK, res)
	return
}


func (groupsController groupsControllerForPrivate) DeleteGroups(c *gin.Context) {
	var groupsRequest request.GroupsRequest
	if err := c.Bind(&groupsRequest); err != nil {
		c.JSON(http.StatusBadRequest, &response.GroupsResponse{Code: "SERVER_CONTROLLER_DELETE__FOR__001", Message: err.Error(), Groupss: []response.Groups{}})
		return
	}
	var uuid string
	res := groupsController.GroupsRepository.DeleteGroups(uuid)
	c.JSON(http.StatusOK, res)
	return
}


func NewGroupsControllerForPrivate(groupsRepository repository.GroupsRepository) GroupsControllerForPrivate {
	return &groupsControllerForPrivate{ GroupsRepository: groupsRepository}
}