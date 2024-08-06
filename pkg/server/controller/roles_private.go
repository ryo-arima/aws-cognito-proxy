package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ryo-arima/aws-cognito-proxy/pkg/entity/model"
	"github.com/ryo-arima/aws-cognito-proxy/pkg/entity/request"
	"github.com/ryo-arima/aws-cognito-proxy/pkg/entity/response"
	"github.com/ryo-arima/aws-cognito-proxy/pkg/server/repository"
)

type RolesControllerForPrivate interface {
	GetRoless(c *gin.Context)
	CreateRoles(c *gin.Context)
	UpdateRoles(c *gin.Context)
	DeleteRoles(c *gin.Context)
}

type rolesControllerForPrivate struct {
	RolesRepository repository.RolesRepository
}

func (rolesController rolesControllerForPrivate) GetRoless(c *gin.Context) {
	var rolesRequest request.RolesRequest
	if err := c.Bind(&rolesRequest); err != nil {
		c.JSON(http.StatusBadRequest, &response.RolesResponse{Code: "SERVER_CONTROLLER_GET__FOR__001", Message: err.Error(), Roless: []response.Roles{}})
		return
	}
	res := rolesController.RolesRepository.GetRoless()
	c.JSON(http.StatusOK, res)
	return
}


func (rolesController rolesControllerForPrivate) CreateRoles(c *gin.Context) {
	var rolesRequest request.RolesRequest
	if err := c.Bind(&rolesRequest); err != nil {
		c.JSON(http.StatusBadRequest, &response.RolesResponse{Code: "SERVER_CONTROLLER_CREATE__FOR__001", Message: err.Error(), Roless: []response.Roles{}})
		return
	}
	var rolesModel model.Roless
	res := rolesController.RolesRepository.CreateRoles(rolesModel)
	c.JSON(http.StatusOK, res)
	return
}


func (rolesController rolesControllerForPrivate) UpdateRoles(c *gin.Context) {
	var rolesRequest request.RolesRequest
	if err := c.Bind(&rolesRequest); err != nil {
		c.JSON(http.StatusBadRequest, &response.RolesResponse{Code: "SERVER_CONTROLLER_UPDATE__FOR__001", Message: err.Error(), Roless: []response.Roles{}})
		return
	}
	var rolesModel model.Roless
	res := rolesController.RolesRepository.UpdateRoles(rolesModel)
	c.JSON(http.StatusOK, res)
	return
}


func (rolesController rolesControllerForPrivate) DeleteRoles(c *gin.Context) {
	var rolesRequest request.RolesRequest
	if err := c.Bind(&rolesRequest); err != nil {
		c.JSON(http.StatusBadRequest, &response.RolesResponse{Code: "SERVER_CONTROLLER_DELETE__FOR__001", Message: err.Error(), Roless: []response.Roles{}})
		return
	}
	var uuid string
	res := rolesController.RolesRepository.DeleteRoles(uuid)
	c.JSON(http.StatusOK, res)
	return
}


func NewRolesControllerForPrivate(rolesRepository repository.RolesRepository) RolesControllerForPrivate {
	return &rolesControllerForPrivate{ RolesRepository: rolesRepository}
}