package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ryo-arima/aws-cognito-proxy/pkg/entity/request"
	"github.com/ryo-arima/aws-cognito-proxy/pkg/entity/response"
	"github.com/ryo-arima/aws-cognito-proxy/pkg/server/repository"
)

type RolesControllerForPublic interface {
	GetRoless(c *gin.Context)
}

type rolesControllerForPublic struct {
	RolesRepository repository.RolesRepository
}

func (rolesController rolesControllerForPublic) GetRoless(c *gin.Context) {
	var rolesRequest request.RolesRequest
	if err := c.Bind(&rolesRequest); err != nil {
		c.JSON(http.StatusBadRequest, &response.RolesResponse{Code: "SERVER_CONTROLLER_GET__FOR__001", Message: err.Error(), Roless: []response.Roles{}})
		return
	}
	res := rolesController.RolesRepository.GetRoless()
	c.JSON(http.StatusOK, res)
	return
}


func NewRolesControllerForPublic(rolesRepository repository.RolesRepository) RolesControllerForPublic {
	return &rolesControllerForPublic{ RolesRepository: rolesRepository}
}