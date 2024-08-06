package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ryo-arima/aws-cognito-proxy/pkg/entity/request"
	"github.com/ryo-arima/aws-cognito-proxy/pkg/entity/response"
	"github.com/ryo-arima/aws-cognito-proxy/pkg/server/repository"
)

type PolicyControllerForPublic interface {
	GetPolicys(c *gin.Context)
}

type policyControllerForPublic struct {
	PolicyRepository repository.PolicyRepository
}

func (policyController policyControllerForPublic) GetPolicys(c *gin.Context) {
	var policyRequest request.PolicyRequest
	if err := c.Bind(&policyRequest); err != nil {
		c.JSON(http.StatusBadRequest, &response.PolicyResponse{Code: "SERVER_CONTROLLER_GET__FOR__001", Message: err.Error(), Policys: []response.Policy{}})
		return
	}
	res := policyController.PolicyRepository.GetPolicys()
	c.JSON(http.StatusOK, res)
	return
}


func NewPolicyControllerForPublic(policyRepository repository.PolicyRepository) PolicyControllerForPublic {
	return &policyControllerForPublic{ PolicyRepository: policyRepository}
}