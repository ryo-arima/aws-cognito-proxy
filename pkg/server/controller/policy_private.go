package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ryo-arima/aws-cognito-proxy/pkg/entity/model"
	"github.com/ryo-arima/aws-cognito-proxy/pkg/entity/request"
	"github.com/ryo-arima/aws-cognito-proxy/pkg/entity/response"
	"github.com/ryo-arima/aws-cognito-proxy/pkg/server/repository"
)

type PolicyControllerForPrivate interface {
	GetPolicys(c *gin.Context)
	CreatePolicy(c *gin.Context)
	UpdatePolicy(c *gin.Context)
	DeletePolicy(c *gin.Context)
}

type policyControllerForPrivate struct {
	PolicyRepository repository.PolicyRepository
}

func (policyController policyControllerForPrivate) GetPolicys(c *gin.Context) {
	var policyRequest request.PolicyRequest
	if err := c.Bind(&policyRequest); err != nil {
		c.JSON(http.StatusBadRequest, &response.PolicyResponse{Code: "SERVER_CONTROLLER_GET__FOR__001", Message: err.Error(), Policys: []response.Policy{}})
		return
	}
	res := policyController.PolicyRepository.GetPolicys()
	c.JSON(http.StatusOK, res)
	return
}


func (policyController policyControllerForPrivate) CreatePolicy(c *gin.Context) {
	var policyRequest request.PolicyRequest
	if err := c.Bind(&policyRequest); err != nil {
		c.JSON(http.StatusBadRequest, &response.PolicyResponse{Code: "SERVER_CONTROLLER_CREATE__FOR__001", Message: err.Error(), Policys: []response.Policy{}})
		return
	}
	var policyModel model.Policys
	res := policyController.PolicyRepository.CreatePolicy(policyModel)
	c.JSON(http.StatusOK, res)
	return
}


func (policyController policyControllerForPrivate) UpdatePolicy(c *gin.Context) {
	var policyRequest request.PolicyRequest
	if err := c.Bind(&policyRequest); err != nil {
		c.JSON(http.StatusBadRequest, &response.PolicyResponse{Code: "SERVER_CONTROLLER_UPDATE__FOR__001", Message: err.Error(), Policys: []response.Policy{}})
		return
	}
	var policyModel model.Policys
	res := policyController.PolicyRepository.UpdatePolicy(policyModel)
	c.JSON(http.StatusOK, res)
	return
}


func (policyController policyControllerForPrivate) DeletePolicy(c *gin.Context) {
	var policyRequest request.PolicyRequest
	if err := c.Bind(&policyRequest); err != nil {
		c.JSON(http.StatusBadRequest, &response.PolicyResponse{Code: "SERVER_CONTROLLER_DELETE__FOR__001", Message: err.Error(), Policys: []response.Policy{}})
		return
	}
	var uuid string
	res := policyController.PolicyRepository.DeletePolicy(uuid)
	c.JSON(http.StatusOK, res)
	return
}


func NewPolicyControllerForPrivate(policyRepository repository.PolicyRepository) PolicyControllerForPrivate {
	return &policyControllerForPrivate{ PolicyRepository: policyRepository}
}