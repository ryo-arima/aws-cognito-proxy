package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ryo-arima/aws-cognito-proxy/pkg/entity/request"
	"github.com/ryo-arima/aws-cognito-proxy/pkg/entity/response"
	"github.com/ryo-arima/aws-cognito-proxy/pkg/server/repository"
)

type PoliciesControllerForPublic interface {
	GetPoliciess(c *gin.Context)
}

type policiesControllerForPublic struct {
	PoliciesRepository repository.PoliciesRepository
}

func (policiesController policiesControllerForPublic) GetPoliciess(c *gin.Context) {
	var policiesRequest request.PoliciesRequest
	if err := c.Bind(&policiesRequest); err != nil {
		c.JSON(http.StatusBadRequest, &response.PoliciesResponse{Code: "SERVER_CONTROLLER_GET__FOR__001", Message: err.Error(), Policiess: []response.Policies{}})
		return
	}
	res := policiesController.PoliciesRepository.GetPoliciess()
	c.JSON(http.StatusOK, res)
	return
}


func NewPoliciesControllerForPublic(policiesRepository repository.PoliciesRepository) PoliciesControllerForPublic {
	return &policiesControllerForPublic{ PoliciesRepository: policiesRepository}
}