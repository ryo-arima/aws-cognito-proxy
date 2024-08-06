package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ryo-arima/aws-cognito-proxy/pkg/entity/model"
	"github.com/ryo-arima/aws-cognito-proxy/pkg/entity/request"
	"github.com/ryo-arima/aws-cognito-proxy/pkg/entity/response"
	"github.com/ryo-arima/aws-cognito-proxy/pkg/server/repository"
)

type PoliciesControllerForPrivate interface {
	GetPoliciess(c *gin.Context)
	CreatePolicies(c *gin.Context)
	UpdatePolicies(c *gin.Context)
	DeletePolicies(c *gin.Context)
}

type policiesControllerForPrivate struct {
	PoliciesRepository repository.PoliciesRepository
}

func (policiesController policiesControllerForPrivate) GetPoliciess(c *gin.Context) {
	var policiesRequest request.PoliciesRequest
	if err := c.Bind(&policiesRequest); err != nil {
		c.JSON(http.StatusBadRequest, &response.PoliciesResponse{Code: "SERVER_CONTROLLER_GET__FOR__001", Message: err.Error(), Policiess: []response.Policies{}})
		return
	}
	res := policiesController.PoliciesRepository.GetPoliciess()
	c.JSON(http.StatusOK, res)
	return
}


func (policiesController policiesControllerForPrivate) CreatePolicies(c *gin.Context) {
	var policiesRequest request.PoliciesRequest
	if err := c.Bind(&policiesRequest); err != nil {
		c.JSON(http.StatusBadRequest, &response.PoliciesResponse{Code: "SERVER_CONTROLLER_CREATE__FOR__001", Message: err.Error(), Policiess: []response.Policies{}})
		return
	}
	var policiesModel model.Policiess
	res := policiesController.PoliciesRepository.CreatePolicies(policiesModel)
	c.JSON(http.StatusOK, res)
	return
}


func (policiesController policiesControllerForPrivate) UpdatePolicies(c *gin.Context) {
	var policiesRequest request.PoliciesRequest
	if err := c.Bind(&policiesRequest); err != nil {
		c.JSON(http.StatusBadRequest, &response.PoliciesResponse{Code: "SERVER_CONTROLLER_UPDATE__FOR__001", Message: err.Error(), Policiess: []response.Policies{}})
		return
	}
	var policiesModel model.Policiess
	res := policiesController.PoliciesRepository.UpdatePolicies(policiesModel)
	c.JSON(http.StatusOK, res)
	return
}


func (policiesController policiesControllerForPrivate) DeletePolicies(c *gin.Context) {
	var policiesRequest request.PoliciesRequest
	if err := c.Bind(&policiesRequest); err != nil {
		c.JSON(http.StatusBadRequest, &response.PoliciesResponse{Code: "SERVER_CONTROLLER_DELETE__FOR__001", Message: err.Error(), Policiess: []response.Policies{}})
		return
	}
	var uuid string
	res := policiesController.PoliciesRepository.DeletePolicies(uuid)
	c.JSON(http.StatusOK, res)
	return
}


func NewPoliciesControllerForPrivate(policiesRepository repository.PoliciesRepository) PoliciesControllerForPrivate {
	return &policiesControllerForPrivate{ PoliciesRepository: policiesRepository}
}