package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ryo-arima/aws-cognito-proxy/pkg/entity/model"
	"github.com/ryo-arima/aws-cognito-proxy/pkg/entity/request"
	"github.com/ryo-arima/aws-cognito-proxy/pkg/entity/response"
	"github.com/ryo-arima/aws-cognito-proxy/pkg/server/repository"
)

type RoleControllerForPrivate interface {
	GetRoles(c *gin.Context)
	CreateRole(c *gin.Context)
	UpdateRole(c *gin.Context)
	DeleteRole(c *gin.Context)
}

type roleControllerForPrivate struct {
	RoleRepository repository.RoleRepository
}

func (roleController roleControllerForPrivate) GetRoles(c *gin.Context) {
	var roleRequest request.RoleRequest
	if err := c.Bind(&roleRequest); err != nil {
		c.JSON(http.StatusBadRequest, &response.RoleResponse{Code: "SERVER_CONTROLLER_GET__FOR__001", Message: err.Error(), Roles: []response.Role{}})
		return
	}
	res := roleController.RoleRepository.GetRoles()
	c.JSON(http.StatusOK, res)
	return
}


func (roleController roleControllerForPrivate) CreateRole(c *gin.Context) {
	var roleRequest request.RoleRequest
	if err := c.Bind(&roleRequest); err != nil {
		c.JSON(http.StatusBadRequest, &response.RoleResponse{Code: "SERVER_CONTROLLER_CREATE__FOR__001", Message: err.Error(), Roles: []response.Role{}})
		return
	}
	var roleModel model.Roles
	res := roleController.RoleRepository.CreateRole(roleModel)
	c.JSON(http.StatusOK, res)
	return
}


func (roleController roleControllerForPrivate) UpdateRole(c *gin.Context) {
	var roleRequest request.RoleRequest
	if err := c.Bind(&roleRequest); err != nil {
		c.JSON(http.StatusBadRequest, &response.RoleResponse{Code: "SERVER_CONTROLLER_UPDATE__FOR__001", Message: err.Error(), Roles: []response.Role{}})
		return
	}
	var roleModel model.Roles
	res := roleController.RoleRepository.UpdateRole(roleModel)
	c.JSON(http.StatusOK, res)
	return
}


func (roleController roleControllerForPrivate) DeleteRole(c *gin.Context) {
	var roleRequest request.RoleRequest
	if err := c.Bind(&roleRequest); err != nil {
		c.JSON(http.StatusBadRequest, &response.RoleResponse{Code: "SERVER_CONTROLLER_DELETE__FOR__001", Message: err.Error(), Roles: []response.Role{}})
		return
	}
	var uuid string
	res := roleController.RoleRepository.DeleteRole(uuid)
	c.JSON(http.StatusOK, res)
	return
}


func NewRoleControllerForPrivate(roleRepository repository.RoleRepository) RoleControllerForPrivate {
	return &roleControllerForPrivate{ RoleRepository: roleRepository}
}