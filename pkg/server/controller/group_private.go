package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ryo-arima/aws-cognito-proxy/pkg/entity/model"
	"github.com/ryo-arima/aws-cognito-proxy/pkg/entity/request"
	"github.com/ryo-arima/aws-cognito-proxy/pkg/entity/response"
	"github.com/ryo-arima/aws-cognito-proxy/pkg/server/repository"
)

type GroupControllerForPrivate interface {
	GetGroups(c *gin.Context)
	CreateGroup(c *gin.Context)
	UpdateGroup(c *gin.Context)
	DeleteGroup(c *gin.Context)
}

type groupControllerForPrivate struct {
	GroupRepository repository.GroupRepository
}

func (groupController groupControllerForPrivate) GetGroups(c *gin.Context) {
	var groupRequest request.GroupRequest
	if err := c.Bind(&groupRequest); err != nil {
		c.JSON(http.StatusBadRequest, &response.GroupResponse{Code: "SERVER_CONTROLLER_GET__FOR__001", Message: err.Error(), Groups: []response.Group{}})
		return
	}
	res := groupController.GroupRepository.GetGroups()
	c.JSON(http.StatusOK, res)
	return
}


func (groupController groupControllerForPrivate) CreateGroup(c *gin.Context) {
	var groupRequest request.GroupRequest
	if err := c.Bind(&groupRequest); err != nil {
		c.JSON(http.StatusBadRequest, &response.GroupResponse{Code: "SERVER_CONTROLLER_CREATE__FOR__001", Message: err.Error(), Groups: []response.Group{}})
		return
	}
	var groupModel model.Groups
	res := groupController.GroupRepository.CreateGroup(groupModel)
	c.JSON(http.StatusOK, res)
	return
}


func (groupController groupControllerForPrivate) UpdateGroup(c *gin.Context) {
	var groupRequest request.GroupRequest
	if err := c.Bind(&groupRequest); err != nil {
		c.JSON(http.StatusBadRequest, &response.GroupResponse{Code: "SERVER_CONTROLLER_UPDATE__FOR__001", Message: err.Error(), Groups: []response.Group{}})
		return
	}
	var groupModel model.Groups
	res := groupController.GroupRepository.UpdateGroup(groupModel)
	c.JSON(http.StatusOK, res)
	return
}


func (groupController groupControllerForPrivate) DeleteGroup(c *gin.Context) {
	var groupRequest request.GroupRequest
	if err := c.Bind(&groupRequest); err != nil {
		c.JSON(http.StatusBadRequest, &response.GroupResponse{Code: "SERVER_CONTROLLER_DELETE__FOR__001", Message: err.Error(), Groups: []response.Group{}})
		return
	}
	var uuid string
	res := groupController.GroupRepository.DeleteGroup(uuid)
	c.JSON(http.StatusOK, res)
	return
}


func NewGroupControllerForPrivate(groupRepository repository.GroupRepository) GroupControllerForPrivate {
	return &groupControllerForPrivate{ GroupRepository: groupRepository}
}